// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *StartBackupPlanResponseBody
	GetBackupPlanId() *string
	SetCreatedFullBackupsetId(v string) *StartBackupPlanResponseBody
	GetCreatedFullBackupsetId() *string
	SetErrCode(v string) *StartBackupPlanResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *StartBackupPlanResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *StartBackupPlanResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *StartBackupPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartBackupPlanResponseBody
	GetSuccess() *bool
}

type StartBackupPlanResponseBody struct {
	// The ID of the backup schedule.
	//
	// example:
	//
	// dbsqdss5tm****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the full backup set.
	//
	// example:
	//
	// 1h7toien3****
	CreatedFullBackupsetId *string `json:"CreatedFullBackupsetId,omitempty" xml:"CreatedFullBackupsetId,omitempty"`
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
	// D13761C3-9971-5C02-B789-3F3CD159****
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

func (s StartBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *StartBackupPlanResponseBody) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *StartBackupPlanResponseBody) GetCreatedFullBackupsetId() *string {
	return s.CreatedFullBackupsetId
}

func (s *StartBackupPlanResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *StartBackupPlanResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *StartBackupPlanResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartBackupPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartBackupPlanResponseBody) SetBackupPlanId(v string) *StartBackupPlanResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *StartBackupPlanResponseBody) SetCreatedFullBackupsetId(v string) *StartBackupPlanResponseBody {
	s.CreatedFullBackupsetId = &v
	return s
}

func (s *StartBackupPlanResponseBody) SetErrCode(v string) *StartBackupPlanResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartBackupPlanResponseBody) SetErrMessage(v string) *StartBackupPlanResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartBackupPlanResponseBody) SetHttpStatusCode(v int32) *StartBackupPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartBackupPlanResponseBody) SetRequestId(v string) *StartBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartBackupPlanResponseBody) SetSuccess(v bool) *StartBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *StartBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
