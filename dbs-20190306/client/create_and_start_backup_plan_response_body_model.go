// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndStartBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *CreateAndStartBackupPlanResponseBody
	GetBackupPlanId() *string
	SetCreateBackupSet(v bool) *CreateAndStartBackupPlanResponseBody
	GetCreateBackupSet() *bool
	SetErrCode(v string) *CreateAndStartBackupPlanResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateAndStartBackupPlanResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CreateAndStartBackupPlanResponseBody
	GetHttpStatusCode() *int32
	SetOrderId(v string) *CreateAndStartBackupPlanResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateAndStartBackupPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAndStartBackupPlanResponseBody
	GetSuccess() *bool
}

type CreateAndStartBackupPlanResponseBody struct {
	// The backup schedule ID.
	//
	// example:
	//
	// dbs1hvb0wwweusfa
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// Indicates whether a backup is performed immediately after the backup schedule is configured. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CreateBackupSet *bool `json:"CreateBackupSet,omitempty" xml:"CreateBackupSet,omitempty"`
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
	// 2056157***
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F1888AC-1138-4995-B9FE-D2734F61C058
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAndStartBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAndStartBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAndStartBackupPlanResponseBody) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *CreateAndStartBackupPlanResponseBody) GetCreateBackupSet() *bool {
	return s.CreateBackupSet
}

func (s *CreateAndStartBackupPlanResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateAndStartBackupPlanResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateAndStartBackupPlanResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAndStartBackupPlanResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateAndStartBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAndStartBackupPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAndStartBackupPlanResponseBody) SetBackupPlanId(v string) *CreateAndStartBackupPlanResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *CreateAndStartBackupPlanResponseBody) SetCreateBackupSet(v bool) *CreateAndStartBackupPlanResponseBody {
	s.CreateBackupSet = &v
	return s
}

func (s *CreateAndStartBackupPlanResponseBody) SetErrCode(v string) *CreateAndStartBackupPlanResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateAndStartBackupPlanResponseBody) SetErrMessage(v string) *CreateAndStartBackupPlanResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateAndStartBackupPlanResponseBody) SetHttpStatusCode(v int32) *CreateAndStartBackupPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAndStartBackupPlanResponseBody) SetOrderId(v string) *CreateAndStartBackupPlanResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateAndStartBackupPlanResponseBody) SetRequestId(v string) *CreateAndStartBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAndStartBackupPlanResponseBody) SetSuccess(v bool) *CreateAndStartBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAndStartBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
