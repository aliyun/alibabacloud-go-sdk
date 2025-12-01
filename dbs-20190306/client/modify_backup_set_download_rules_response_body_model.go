// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupSetDownloadRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *ModifyBackupSetDownloadRulesResponseBody
	GetBackupPlanId() *string
	SetErrCode(v string) *ModifyBackupSetDownloadRulesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyBackupSetDownloadRulesResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyBackupSetDownloadRulesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyBackupSetDownloadRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyBackupSetDownloadRulesResponseBody
	GetSuccess() *bool
}

type ModifyBackupSetDownloadRulesResponseBody struct {
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

func (s ModifyBackupSetDownloadRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupSetDownloadRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupSetDownloadRulesResponseBody) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ModifyBackupSetDownloadRulesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyBackupSetDownloadRulesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyBackupSetDownloadRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyBackupSetDownloadRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBackupSetDownloadRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyBackupSetDownloadRulesResponseBody) SetBackupPlanId(v string) *ModifyBackupSetDownloadRulesResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesResponseBody) SetErrCode(v string) *ModifyBackupSetDownloadRulesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesResponseBody) SetErrMessage(v string) *ModifyBackupSetDownloadRulesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesResponseBody) SetHttpStatusCode(v int32) *ModifyBackupSetDownloadRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesResponseBody) SetRequestId(v string) *ModifyBackupSetDownloadRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesResponseBody) SetSuccess(v bool) *ModifyBackupSetDownloadRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
