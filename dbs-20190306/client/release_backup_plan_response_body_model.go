// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *ReleaseBackupPlanResponseBody
	GetBackupPlanId() *string
	SetErrCode(v string) *ReleaseBackupPlanResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ReleaseBackupPlanResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ReleaseBackupPlanResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ReleaseBackupPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReleaseBackupPlanResponseBody
	GetSuccess() *bool
}

type ReleaseBackupPlanResponseBody struct {
	// The ID of the backup schedule.
	//
	// example:
	//
	// dbstooi01****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// findValidDBSJob error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code returned.
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
	// Indicates whether the request succeeded. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseBackupPlanResponseBody) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ReleaseBackupPlanResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ReleaseBackupPlanResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ReleaseBackupPlanResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ReleaseBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseBackupPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReleaseBackupPlanResponseBody) SetBackupPlanId(v string) *ReleaseBackupPlanResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *ReleaseBackupPlanResponseBody) SetErrCode(v string) *ReleaseBackupPlanResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ReleaseBackupPlanResponseBody) SetErrMessage(v string) *ReleaseBackupPlanResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ReleaseBackupPlanResponseBody) SetHttpStatusCode(v int32) *ReleaseBackupPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReleaseBackupPlanResponseBody) SetRequestId(v string) *ReleaseBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseBackupPlanResponseBody) SetSuccess(v bool) *ReleaseBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *ReleaseBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
