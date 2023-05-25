package logic

import (
	"context"

	"github.com/andres06-hub/users/internal/svc"
	"github.com/andres06-hub/users/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *users.Request) (*users.Response, error) {
	// todo: add your logic here and delete this line

	return &users.Response{}, nil
}
