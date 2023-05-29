package logic

import (
	"context"

	"github.com/andres06-hub/users/internal/svc"
	"github.com/andres06-hub/users/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *users.CreateUserRequest) (*users.Response, error) {
	// todo: add your logic here and delete this line

	return &users.Response{}, nil
}
