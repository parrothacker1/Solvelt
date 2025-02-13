package routers

import (
	"net/http"

	"github.com/parrothacker1/Solvelt/users/handlers"
	"github.com/parrothacker1/Solvelt/users/utils/middlewares"
)

var TeamRouter *Router

func init() {
  TeamRouter = NewRouter()
  TeamRouter.Handle(http.MethodPost,"/",handlers.CreateTeam,middlewares.AuthMiddleware)
  TeamRouter.Handle(http.MethodPut,"/",handlers.UpdateTeam,middlewares.AuthMiddleware,middlewares.TeamLeaderMiddleware)
  TeamRouter.Handle(http.MethodDelete,"/",handlers.DeleteTeam,middlewares.AuthMiddleware,middlewares.TeamLeaderMiddleware)
  TeamRouter.Handle(http.MethodPost,"/",handlers.JoinTeam,middlewares.AuthMiddleware)
  TeamRouter.Handle(http.MethodGet,"/me",handlers.GetTeam,middlewares.AuthMiddleware)
}
