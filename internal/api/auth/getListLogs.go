package authAPI

import (
	"context"
	"github.com/semho/chat-microservices/auth/internal/converter"
	desc "github.com/semho/chat-microservices/auth/pkg/auth_v1"
	"log"
)

func (i *Implementation) GetListLogs(ctx context.Context, req *desc.GetListLogsRequest) (*desc.LogsResponse, error) {
	log.Printf("get logs")

	pageNumber := req.GetPageNumber()
	pageSize := req.GetPageSize()

	log.Printf("pageSize: %d", pageSize)

	if pageNumber == 0 {
		pageNumber = 1
	}

	if pageSize == 0 {
		pageSize = 5
	}

	listLogs, err := i.authService.GetListLogs(ctx, pageNumber, pageSize)
	if err != nil {
		return nil, err
	}

	return converter.ToLogFromService(listLogs), nil
}
