// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *ModifyBackupStrategyResponseBody
	GetBackupPlanId() *string
	SetErrCode(v string) *ModifyBackupStrategyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyBackupStrategyResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyBackupStrategyResponseBody
	GetHttpStatusCode() *int32
	SetNeedPrecheck(v bool) *ModifyBackupStrategyResponseBody
	GetNeedPrecheck() *bool
	SetRequestId(v string) *ModifyBackupStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyBackupStrategyResponseBody
	GetSuccess() *bool
}

type ModifyBackupStrategyResponseBody struct {
	// The ID of the backup schedule.
	//
	// example:
	//
	// dbstooi01XXXX
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
	// Indicates whether a precheck is triggered. If the value of this parameter is true, you must start the backup schedule by calling the StartBackupPlan operation.
	//
	// example:
	//
	// false
	NeedPrecheck *bool `json:"NeedPrecheck,omitempty" xml:"NeedPrecheck,omitempty"`
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

func (s ModifyBackupStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupStrategyResponseBody) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ModifyBackupStrategyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyBackupStrategyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyBackupStrategyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyBackupStrategyResponseBody) GetNeedPrecheck() *bool {
	return s.NeedPrecheck
}

func (s *ModifyBackupStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBackupStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyBackupStrategyResponseBody) SetBackupPlanId(v string) *ModifyBackupStrategyResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupStrategyResponseBody) SetErrCode(v string) *ModifyBackupStrategyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyBackupStrategyResponseBody) SetErrMessage(v string) *ModifyBackupStrategyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyBackupStrategyResponseBody) SetHttpStatusCode(v int32) *ModifyBackupStrategyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyBackupStrategyResponseBody) SetNeedPrecheck(v bool) *ModifyBackupStrategyResponseBody {
	s.NeedPrecheck = &v
	return s
}

func (s *ModifyBackupStrategyResponseBody) SetRequestId(v string) *ModifyBackupStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupStrategyResponseBody) SetSuccess(v bool) *ModifyBackupStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBackupStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
