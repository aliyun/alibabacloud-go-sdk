// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *StopBackupPlanResponseBody
	GetBackupPlanId() *string
	SetErrCode(v string) *StopBackupPlanResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *StopBackupPlanResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *StopBackupPlanResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *StopBackupPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopBackupPlanResponseBody
	GetSuccess() *bool
}

type StopBackupPlanResponseBody struct {
	// The ID of the backup schedule.
	//
	// example:
	//
	// dbs1h****usfa
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

func (s StopBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *StopBackupPlanResponseBody) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *StopBackupPlanResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *StopBackupPlanResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *StopBackupPlanResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StopBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopBackupPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopBackupPlanResponseBody) SetBackupPlanId(v string) *StopBackupPlanResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *StopBackupPlanResponseBody) SetErrCode(v string) *StopBackupPlanResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StopBackupPlanResponseBody) SetErrMessage(v string) *StopBackupPlanResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StopBackupPlanResponseBody) SetHttpStatusCode(v int32) *StopBackupPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopBackupPlanResponseBody) SetRequestId(v string) *StopBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopBackupPlanResponseBody) SetSuccess(v bool) *StopBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *StopBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
