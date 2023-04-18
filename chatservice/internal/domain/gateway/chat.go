package gateway

import (
	"context"

	"github.com/yrllanio/fclx/chatservice/internal/domain/entity"
)

type ChatGateway interface {
	CreteChat(ctx context.Context, chat *entity.Chat) error
	FindChatByID(ctx context.Context, chatID string) (*entity.Chat, error)
	SaveChat(ctx context.Context, chat *entity.Chat) error
}
