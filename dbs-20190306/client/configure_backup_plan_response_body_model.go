// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *ConfigureBackupPlanResponseBody
	GetBackupPlanId() *string
	SetErrCode(v string) *ConfigureBackupPlanResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ConfigureBackupPlanResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ConfigureBackupPlanResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ConfigureBackupPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConfigureBackupPlanResponseBody
	GetSuccess() *bool
}

type ConfigureBackupPlanResponseBody struct {
	// The ID of the backup schedule.
	//
	// example:
	//
	// dbstooi01ex****
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
	// F1FB49D4-B504-47F1-9F43-D7EAB33F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigureBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureBackupPlanResponseBody) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ConfigureBackupPlanResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ConfigureBackupPlanResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ConfigureBackupPlanResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ConfigureBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigureBackupPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConfigureBackupPlanResponseBody) SetBackupPlanId(v string) *ConfigureBackupPlanResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *ConfigureBackupPlanResponseBody) SetErrCode(v string) *ConfigureBackupPlanResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureBackupPlanResponseBody) SetErrMessage(v string) *ConfigureBackupPlanResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureBackupPlanResponseBody) SetHttpStatusCode(v int32) *ConfigureBackupPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ConfigureBackupPlanResponseBody) SetRequestId(v string) *ConfigureBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureBackupPlanResponseBody) SetSuccess(v bool) *ConfigureBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
