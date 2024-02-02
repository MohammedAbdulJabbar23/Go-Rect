package routes


import (
  "github.com/gin-gonic/gin"
  "net/http"
  "example.com/event-booking/models"
  "strconv"
  "fmt"
)


func getEvents(context *gin.Context) {
    events,err := models.GetAllEvents();
    if err != nil {
      context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later"});
      return;
    }
    context.JSON(http.StatusOK, events);
}

func createEvent(context *gin.Context) {

  var event models.Event 
    err := context.ShouldBindJSON(&event);
    if err != nil {
      context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse request data"});
      return;
    }

    userId := context.GetInt64("userId")
    event.UserID = userId;
    err = event.Save();
    if err != nil {
      context.JSON(http.StatusInternalServerError, gin.H{"message":"could not create event. try again later"});
      return;
    }
    context.JSON(http.StatusCreated, gin.H{"message": "Event Created!","event":event});
}
func getEvent(context *gin.Context) {
  eventId, err := strconv.ParseInt(context.Param("id"), 10, 64);
  fmt.Println(eventId)
  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event id"});
    return;
  }
  event, err := models.GetEventByID(eventId);
  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message":"Could not find event id"});
    return;
  }
  context.JSON(http.StatusOK, event);
}

func updateEvent(context *gin.Context) {
  eventId, err := strconv.ParseInt(context.Param("id"),10,64);
  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message":"Could not update event"});
    return;
  }
  event, err := models.GetEventByID(eventId);
  userId := context.GetInt64("userId");
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message":"could not find event"});
    return;
  }
  if event.UserID != userId {
    context.JSON(http.StatusUnauthorized, gin.H{"message":"Not authorized to update event"});
    return;
  }
  var updatedEvent models.Event;
  err = context.ShouldBindJSON(&updatedEvent);
  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message":"could not parse request data"});
    return;
  }
  updatedEvent.ID = eventId;
  err = updatedEvent.Update();
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"});
    return;
  }
  context.JSON(http.StatusOK, gin.H{"message": "Event Updated successfully"});
}

func deleteEvent(context *gin.Context) {
  eventId, err := strconv.ParseInt(context.Param("id"),10,64);
  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message":"could not parse event id."})
    return;
  }
  event, err := models.GetEventByID(eventId);
  if err != nil {
    context.JSON(http.StatusInternalServerError,gin.H{"message": "could not parse id"});
    return;
  }

  userId := context.GetInt64("userId");
  if event.UserID != userId {
    context.JSON(http.StatusUnauthorized, gin.H{"message":"Not authorized to update event"});
    return;
  }
  err = event.Delete();
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the event."});
  return;
  }
  context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"});
}
