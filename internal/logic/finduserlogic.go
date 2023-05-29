package logic

import (
	"context"

	"github.com/andres06-hub/users/internal/svc"
	"github.com/andres06-hub/users/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *users.FindUserRequest) (*users.Response, error) {
	// todo: add your logic here and delete this line

	return &users.Response{}, nil
}
