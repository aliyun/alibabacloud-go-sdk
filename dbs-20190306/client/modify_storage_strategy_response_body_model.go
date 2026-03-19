// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStorageStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *ModifyStorageStrategyResponseBody
	GetBackupPlanId() *string
	SetErrCode(v string) *ModifyStorageStrategyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyStorageStrategyResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyStorageStrategyResponseBody
	GetHttpStatusCode() *int32
	SetNeedPrecheck(v bool) *ModifyStorageStrategyResponseBody
	GetNeedPrecheck() *bool
	SetRequestId(v string) *ModifyStorageStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyStorageStrategyResponseBody
	GetSuccess() *bool
}

type ModifyStorageStrategyResponseBody struct {
	// Backup plan ID.
	//
	// example:
	//
	// dbsqdss5tmh****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// Error code.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// findValidDBSJob error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Indicates whether this modification triggers a precheck. Return values:
	//
	// - **true**: A precheck is triggered. Manually call the [StartBackupPlan](https://help.aliyun.com/document_detail/2869818.html) API to start the backup plan.
	//
	// - **false**: No precheck is triggered.
	//
	// example:
	//
	// false
	NeedPrecheck *bool `json:"NeedPrecheck,omitempty" xml:"NeedPrecheck,omitempty"`
	// Request ID.
	//
	// example:
	//
	// E995F91F-6F89-503B-9F7D-502F58FD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Return values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyStorageStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyStorageStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyStorageStrategyResponseBody) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ModifyStorageStrategyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyStorageStrategyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyStorageStrategyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyStorageStrategyResponseBody) GetNeedPrecheck() *bool {
	return s.NeedPrecheck
}

func (s *ModifyStorageStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyStorageStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyStorageStrategyResponseBody) SetBackupPlanId(v string) *ModifyStorageStrategyResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyStorageStrategyResponseBody) SetErrCode(v string) *ModifyStorageStrategyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyStorageStrategyResponseBody) SetErrMessage(v string) *ModifyStorageStrategyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyStorageStrategyResponseBody) SetHttpStatusCode(v int32) *ModifyStorageStrategyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyStorageStrategyResponseBody) SetNeedPrecheck(v bool) *ModifyStorageStrategyResponseBody {
	s.NeedPrecheck = &v
	return s
}

func (s *ModifyStorageStrategyResponseBody) SetRequestId(v string) *ModifyStorageStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyStorageStrategyResponseBody) SetSuccess(v bool) *ModifyStorageStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyStorageStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
