// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPlanNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *ModifyBackupPlanNameResponseBody
	GetBackupPlanId() *string
	SetErrCode(v string) *ModifyBackupPlanNameResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyBackupPlanNameResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyBackupPlanNameResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyBackupPlanNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyBackupPlanNameResponseBody
	GetSuccess() *bool
}

type ModifyBackupPlanNameResponseBody struct {
	// The ID of the backup schedule.
	//
	// example:
	//
	// dbstooi0XXXX
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

func (s ModifyBackupPlanNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPlanNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanNameResponseBody) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ModifyBackupPlanNameResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyBackupPlanNameResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyBackupPlanNameResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyBackupPlanNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBackupPlanNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyBackupPlanNameResponseBody) SetBackupPlanId(v string) *ModifyBackupPlanNameResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupPlanNameResponseBody) SetErrCode(v string) *ModifyBackupPlanNameResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyBackupPlanNameResponseBody) SetErrMessage(v string) *ModifyBackupPlanNameResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyBackupPlanNameResponseBody) SetHttpStatusCode(v int32) *ModifyBackupPlanNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyBackupPlanNameResponseBody) SetRequestId(v string) *ModifyBackupPlanNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupPlanNameResponseBody) SetSuccess(v bool) *ModifyBackupPlanNameResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBackupPlanNameResponseBody) Validate() error {
	return dara.Validate(s)
}
