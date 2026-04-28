// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditLogExportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskId(v string) *AuditLogExportResponseBody
	GetAsyncTaskId() *string
}

type AuditLogExportResponseBody struct {
	// The ID of the asynchronous task used to export audit logs.
	//
	// example:
	//
	// 4221bf6e6ab43c255edc4463bf3a6f5f5d31****
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
}

func (s AuditLogExportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuditLogExportResponseBody) GoString() string {
	return s.String()
}

func (s *AuditLogExportResponseBody) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *AuditLogExportResponseBody) SetAsyncTaskId(v string) *AuditLogExportResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *AuditLogExportResponseBody) Validate() error {
	return dara.Validate(s)
}
