// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *RenewBackupPlanResponseBody
	GetBackupPlanId() *string
	SetErrCode(v string) *RenewBackupPlanResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *RenewBackupPlanResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *RenewBackupPlanResponseBody
	GetHttpStatusCode() *int32
	SetOrderId(v string) *RenewBackupPlanResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RenewBackupPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RenewBackupPlanResponseBody
	GetSuccess() *bool
}

type RenewBackupPlanResponseBody struct {
	// The ID of the backup schedule.
	//
	// example:
	//
	// dbstooi01e****
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
	// 20202****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D6E068C3-25BC-455A-85FE-45F0B22ECB1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenewBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *RenewBackupPlanResponseBody) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *RenewBackupPlanResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RenewBackupPlanResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *RenewBackupPlanResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RenewBackupPlanResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewBackupPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RenewBackupPlanResponseBody) SetBackupPlanId(v string) *RenewBackupPlanResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *RenewBackupPlanResponseBody) SetErrCode(v string) *RenewBackupPlanResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RenewBackupPlanResponseBody) SetErrMessage(v string) *RenewBackupPlanResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *RenewBackupPlanResponseBody) SetHttpStatusCode(v int32) *RenewBackupPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RenewBackupPlanResponseBody) SetOrderId(v string) *RenewBackupPlanResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewBackupPlanResponseBody) SetRequestId(v string) *RenewBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewBackupPlanResponseBody) SetSuccess(v bool) *RenewBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *RenewBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
