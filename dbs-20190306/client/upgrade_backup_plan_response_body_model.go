// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *UpgradeBackupPlanResponseBody
	GetBackupPlanId() *string
	SetErrCode(v string) *UpgradeBackupPlanResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpgradeBackupPlanResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *UpgradeBackupPlanResponseBody
	GetHttpStatusCode() *int32
	SetOrderId(v string) *UpgradeBackupPlanResponseBody
	GetOrderId() *string
	SetRequestId(v string) *UpgradeBackupPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeBackupPlanResponseBody
	GetSuccess() *bool
}

type UpgradeBackupPlanResponseBody struct {
	// The ID of the backup schedule.
	//
	// example:
	//
	// dbstooi01****
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
	// 2056157****
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

func (s UpgradeBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeBackupPlanResponseBody) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *UpgradeBackupPlanResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpgradeBackupPlanResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpgradeBackupPlanResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpgradeBackupPlanResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *UpgradeBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeBackupPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeBackupPlanResponseBody) SetBackupPlanId(v string) *UpgradeBackupPlanResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *UpgradeBackupPlanResponseBody) SetErrCode(v string) *UpgradeBackupPlanResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpgradeBackupPlanResponseBody) SetErrMessage(v string) *UpgradeBackupPlanResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpgradeBackupPlanResponseBody) SetHttpStatusCode(v int32) *UpgradeBackupPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpgradeBackupPlanResponseBody) SetOrderId(v string) *UpgradeBackupPlanResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeBackupPlanResponseBody) SetRequestId(v string) *UpgradeBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeBackupPlanResponseBody) SetSuccess(v bool) *UpgradeBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
