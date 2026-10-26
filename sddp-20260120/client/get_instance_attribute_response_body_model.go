// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentKernelVersion(v string) *GetInstanceAttributeResponseBody
	GetCurrentKernelVersion() *string
	SetEngine(v string) *GetInstanceAttributeResponseBody
	GetEngine() *string
	SetEngineVersion(v string) *GetInstanceAttributeResponseBody
	GetEngineVersion() *string
	SetErrorCode(v string) *GetInstanceAttributeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetInstanceAttributeResponseBody
	GetErrorMessage() *string
	SetKmsEncryptionSupported(v bool) *GetInstanceAttributeResponseBody
	GetKmsEncryptionSupported() *bool
	SetMaintainEndTime(v int64) *GetInstanceAttributeResponseBody
	GetMaintainEndTime() *int64
	SetMaintainStartTime(v int64) *GetInstanceAttributeResponseBody
	GetMaintainStartTime() *int64
	SetRequestId(v string) *GetInstanceAttributeResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetInstanceAttributeResponseBody
	GetStatus() *string
}

type GetInstanceAttributeResponseBody struct {
	// The current kernel version of the instance.
	//
	// example:
	//
	// rds_20220731
	CurrentKernelVersion *string `json:"CurrentKernelVersion,omitempty" xml:"CurrentKernelVersion,omitempty"`
	// The database engine type.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The major version of the database engine.
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The business error code for the record or operation. This parameter can be empty when no error occurs. You should evaluate this parameter together with ErrorMessage.
	//
	// example:
	//
	// data_masking_not_running
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The business error message for the record or operation.
	//
	// example:
	//
	// The instance is not in the running state
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance supports using KMS keys for column encryption.
	KmsEncryptionSupported *bool `json:"KmsEncryptionSupported,omitempty" xml:"KmsEncryptionSupported,omitempty"`
	// The end time of the maintenance window. The value is a UNIX timestamp in milliseconds that have elapsed since 1970-01-01 00:00:00 UTC. The returned value combines the UTC time of the maintenance window with the date of the query day.
	//
	// example:
	//
	// 1788570000000
	MaintainEndTime *int64 `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window. The value is a UNIX timestamp in milliseconds that have elapsed since 1970-01-01 00:00:00 UTC. The returned value combines the UTC time of the maintenance window with the date of the query day.
	//
	// example:
	//
	// 1788566400000
	MaintainStartTime *int64 `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The request ID, which can be used to locate this API call.
	//
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The current running status of the instance.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceAttributeResponseBody) GetCurrentKernelVersion() *string {
	return s.CurrentKernelVersion
}

func (s *GetInstanceAttributeResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *GetInstanceAttributeResponseBody) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *GetInstanceAttributeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetInstanceAttributeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetInstanceAttributeResponseBody) GetKmsEncryptionSupported() *bool {
	return s.KmsEncryptionSupported
}

func (s *GetInstanceAttributeResponseBody) GetMaintainEndTime() *int64 {
	return s.MaintainEndTime
}

func (s *GetInstanceAttributeResponseBody) GetMaintainStartTime() *int64 {
	return s.MaintainStartTime
}

func (s *GetInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceAttributeResponseBody) SetCurrentKernelVersion(v string) *GetInstanceAttributeResponseBody {
	s.CurrentKernelVersion = &v
	return s
}

func (s *GetInstanceAttributeResponseBody) SetEngine(v string) *GetInstanceAttributeResponseBody {
	s.Engine = &v
	return s
}

func (s *GetInstanceAttributeResponseBody) SetEngineVersion(v string) *GetInstanceAttributeResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *GetInstanceAttributeResponseBody) SetErrorCode(v string) *GetInstanceAttributeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetInstanceAttributeResponseBody) SetErrorMessage(v string) *GetInstanceAttributeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetInstanceAttributeResponseBody) SetKmsEncryptionSupported(v bool) *GetInstanceAttributeResponseBody {
	s.KmsEncryptionSupported = &v
	return s
}

func (s *GetInstanceAttributeResponseBody) SetMaintainEndTime(v int64) *GetInstanceAttributeResponseBody {
	s.MaintainEndTime = &v
	return s
}

func (s *GetInstanceAttributeResponseBody) SetMaintainStartTime(v int64) *GetInstanceAttributeResponseBody {
	s.MaintainStartTime = &v
	return s
}

func (s *GetInstanceAttributeResponseBody) SetRequestId(v string) *GetInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceAttributeResponseBody) SetStatus(v string) *GetInstanceAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
