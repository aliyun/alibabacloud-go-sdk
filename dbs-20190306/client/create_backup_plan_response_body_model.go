// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *CreateBackupPlanResponseBody
	GetBackupPlanId() *string
	SetErrCode(v string) *CreateBackupPlanResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateBackupPlanResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CreateBackupPlanResponseBody
	GetHttpStatusCode() *int32
	SetOrderId(v string) *CreateBackupPlanResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateBackupPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBackupPlanResponseBody
	GetSuccess() *bool
}

type CreateBackupPlanResponseBody struct {
	// The ID of the backup schedule.
	//
	// example:
	//
	// dbsrhahrsu2****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The error code.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 21437345592****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7BCF6D62-885F-5A4A-91A1-679760E7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanResponseBody) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *CreateBackupPlanResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateBackupPlanResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateBackupPlanResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBackupPlanResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBackupPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBackupPlanResponseBody) SetBackupPlanId(v string) *CreateBackupPlanResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetErrCode(v string) *CreateBackupPlanResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetErrMessage(v string) *CreateBackupPlanResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetHttpStatusCode(v int32) *CreateBackupPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetOrderId(v string) *CreateBackupPlanResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetRequestId(v string) *CreateBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetSuccess(v bool) *CreateBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
