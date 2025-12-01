// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *ModifyBackupObjectsResponseBody
	GetBackupPlanId() *string
	SetErrCode(v string) *ModifyBackupObjectsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyBackupObjectsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyBackupObjectsResponseBody
	GetHttpStatusCode() *int32
	SetNeedPrecheck(v bool) *ModifyBackupObjectsResponseBody
	GetNeedPrecheck() *bool
	SetRequestId(v string) *ModifyBackupObjectsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyBackupObjectsResponseBody
	GetSuccess() *bool
}

type ModifyBackupObjectsResponseBody struct {
	// The ID of the backup schedule.
	//
	// example:
	//
	// dbs1h****usfa
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The error code returned if the request failed.
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
	// Indicates whether a precheck is triggered. If true is returned, you must call the [StartBackupPlan](https://help.aliyun.com/document_detail/2869816.html) operation to start the backup schedule.
	//
	// example:
	//
	// true
	NeedPrecheck *bool `json:"NeedPrecheck,omitempty" xml:"NeedPrecheck,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D6E068C3-25BC-455A-85FE-45F0B22ECB1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyBackupObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupObjectsResponseBody) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ModifyBackupObjectsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyBackupObjectsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyBackupObjectsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyBackupObjectsResponseBody) GetNeedPrecheck() *bool {
	return s.NeedPrecheck
}

func (s *ModifyBackupObjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBackupObjectsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyBackupObjectsResponseBody) SetBackupPlanId(v string) *ModifyBackupObjectsResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupObjectsResponseBody) SetErrCode(v string) *ModifyBackupObjectsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyBackupObjectsResponseBody) SetErrMessage(v string) *ModifyBackupObjectsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyBackupObjectsResponseBody) SetHttpStatusCode(v int32) *ModifyBackupObjectsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyBackupObjectsResponseBody) SetNeedPrecheck(v bool) *ModifyBackupObjectsResponseBody {
	s.NeedPrecheck = &v
	return s
}

func (s *ModifyBackupObjectsResponseBody) SetRequestId(v string) *ModifyBackupObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupObjectsResponseBody) SetSuccess(v bool) *ModifyBackupObjectsResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBackupObjectsResponseBody) Validate() error {
	return dara.Validate(s)
}
