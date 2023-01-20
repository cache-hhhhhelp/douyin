// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	interaction "douyin/internal/handler/interaction"
	message "douyin/internal/handler/message"
	relation "douyin/internal/handler/relation"
	user "douyin/internal/handler/user"
	video "douyin/internal/handler/video"
	"douyin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: user.GetUserHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: video.ListVideoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/feed"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/action",
				Handler: video.UploadVideoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: video.ListVideoByUserHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/publish"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/action",
				Handler: interaction.FavoriteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: interaction.ListFavoriteHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/favorite"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/action",
				Handler: interaction.CommentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: interaction.ListCommentHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/comment"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/relation/action",
				Handler: relation.FollowHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/relation/follow/list",
				Handler: relation.ListFollowHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/relation/follower/list",
				Handler: relation.ListFollowerHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/relation/friend/list",
				Handler: relation.ListFriendHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/message/action",
				Handler: message.SendMessageHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/message/chat",
				Handler: message.ListMessageHandler(serverCtx),
			},
		},
	)
}
