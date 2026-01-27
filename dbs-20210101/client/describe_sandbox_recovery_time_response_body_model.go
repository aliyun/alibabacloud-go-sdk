// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSandboxRecoveryTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSandboxRecoveryTimeResponseBody
	GetCode() *string
	SetData(v *DescribeSandboxRecoveryTimeResponseBodyData) *DescribeSandboxRecoveryTimeResponseBody
	GetData() *DescribeSandboxRecoveryTimeResponseBodyData
	SetErrCode(v string) *DescribeSandboxRecoveryTimeResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeSandboxRecoveryTimeResponseBody
	GetErrMessage() *string
	SetMessage(v string) *DescribeSandboxRecoveryTimeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSandboxRecoveryTimeResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSandboxRecoveryTimeResponseBody
	GetSuccess() *string
}

type DescribeSandboxRecoveryTimeResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// Param.NotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *DescribeSandboxRecoveryTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F1888AC-1138-4995-B9FE-D2734F61C058
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSandboxRecoveryTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSandboxRecoveryTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSandboxRecoveryTimeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSandboxRecoveryTimeResponseBody) GetData() *DescribeSandboxRecoveryTimeResponseBodyData {
	return s.Data
}

func (s *DescribeSandboxRecoveryTimeResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeSandboxRecoveryTimeResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeSandboxRecoveryTimeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSandboxRecoveryTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSandboxRecoveryTimeResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetCode(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetData(v *DescribeSandboxRecoveryTimeResponseBodyData) *DescribeSandboxRecoveryTimeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetErrCode(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetErrMessage(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetMessage(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetRequestId(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetSuccess(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSandboxRecoveryTimeResponseBodyData struct {
	// The backup schedule of the sandbox instance.
	//
	// example:
	//
	// 1hxxxx8xxxxxa
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The beginning of the time range during which the sandbox instance can be restored. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-08-01T12:01:01Z
	RecoveryBeginTime *string `json:"RecoveryBeginTime,omitempty" xml:"RecoveryBeginTime,omitempty"`
	// The end of the time range during which the sandbox instance can be restored. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-08-02T12:01:01Z
	RecoveryEndTime *string `json:"RecoveryEndTime,omitempty" xml:"RecoveryEndTime,omitempty"`
}

func (s DescribeSandboxRecoveryTimeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSandboxRecoveryTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSandboxRecoveryTimeResponseBodyData) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribeSandboxRecoveryTimeResponseBodyData) GetRecoveryBeginTime() *string {
	return s.RecoveryBeginTime
}

func (s *DescribeSandboxRecoveryTimeResponseBodyData) GetRecoveryEndTime() *string {
	return s.RecoveryEndTime
}

func (s *DescribeSandboxRecoveryTimeResponseBodyData) SetBackupPlanId(v string) *DescribeSandboxRecoveryTimeResponseBodyData {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBodyData) SetRecoveryBeginTime(v string) *DescribeSandboxRecoveryTimeResponseBodyData {
	s.RecoveryBeginTime = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBodyData) SetRecoveryEndTime(v string) *DescribeSandboxRecoveryTimeResponseBodyData {
	s.RecoveryEndTime = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
