// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableBackupLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *DisableBackupLogResponseBody
	GetBackupPlanId() *string
	SetErrCode(v string) *DisableBackupLogResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DisableBackupLogResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DisableBackupLogResponseBody
	GetHttpStatusCode() *int32
	SetNeedPrecheck(v bool) *DisableBackupLogResponseBody
	GetNeedPrecheck() *bool
	SetRequestId(v string) *DisableBackupLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableBackupLogResponseBody
	GetSuccess() *bool
}

type DisableBackupLogResponseBody struct {
	// The backup schedule ID.
	//
	// example:
	//
	// dbstooi01xxxx
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
	// findValidDBSJob error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Indicates whether a precheck is triggered. Valid values:
	//
	// 	- true: A precheck is triggered. You must call the [StartBackupPlan](https://help.aliyun.com/document_detail/2869816.html) operation to start the backup schedule.
	//
	// 	- false: No precheck is triggered.
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

func (s DisableBackupLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableBackupLogResponseBody) GoString() string {
	return s.String()
}

func (s *DisableBackupLogResponseBody) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DisableBackupLogResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DisableBackupLogResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DisableBackupLogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DisableBackupLogResponseBody) GetNeedPrecheck() *bool {
	return s.NeedPrecheck
}

func (s *DisableBackupLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableBackupLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableBackupLogResponseBody) SetBackupPlanId(v string) *DisableBackupLogResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *DisableBackupLogResponseBody) SetErrCode(v string) *DisableBackupLogResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DisableBackupLogResponseBody) SetErrMessage(v string) *DisableBackupLogResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DisableBackupLogResponseBody) SetHttpStatusCode(v int32) *DisableBackupLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableBackupLogResponseBody) SetNeedPrecheck(v bool) *DisableBackupLogResponseBody {
	s.NeedPrecheck = &v
	return s
}

func (s *DisableBackupLogResponseBody) SetRequestId(v string) *DisableBackupLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableBackupLogResponseBody) SetSuccess(v bool) *DisableBackupLogResponseBody {
	s.Success = &v
	return s
}

func (s *DisableBackupLogResponseBody) Validate() error {
	return dara.Validate(s)
}
