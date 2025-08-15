// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDisasterRecoveryCheckpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SyncDisasterRecoveryCheckpointResponseBody
	GetCode() *string
	SetData(v bool) *SyncDisasterRecoveryCheckpointResponseBody
	GetData() *bool
	SetDynamicCode(v string) *SyncDisasterRecoveryCheckpointResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *SyncDisasterRecoveryCheckpointResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *SyncDisasterRecoveryCheckpointResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SyncDisasterRecoveryCheckpointResponseBody
	GetMessage() *string
	SetRequestId(v string) *SyncDisasterRecoveryCheckpointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncDisasterRecoveryCheckpointResponseBody
	GetSuccess() *bool
}

type SyncDisasterRecoveryCheckpointResponseBody struct {
	// Error Code
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Result Data
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Dynamic Error Code
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// HTTP Status Code
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error Message
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 7358418D-83BD-507A-8079-611C63E05674
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Success or Not
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncDisasterRecoveryCheckpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncDisasterRecoveryCheckpointResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) GetCode() *string {
	return s.Code
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) GetData() *bool {
	return s.Data
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) SetCode(v string) *SyncDisasterRecoveryCheckpointResponseBody {
	s.Code = &v
	return s
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) SetData(v bool) *SyncDisasterRecoveryCheckpointResponseBody {
	s.Data = &v
	return s
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) SetDynamicCode(v string) *SyncDisasterRecoveryCheckpointResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) SetDynamicMessage(v string) *SyncDisasterRecoveryCheckpointResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) SetHttpStatusCode(v int32) *SyncDisasterRecoveryCheckpointResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) SetMessage(v string) *SyncDisasterRecoveryCheckpointResponseBody {
	s.Message = &v
	return s
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) SetRequestId(v string) *SyncDisasterRecoveryCheckpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) SetSuccess(v bool) *SyncDisasterRecoveryCheckpointResponseBody {
	s.Success = &v
	return s
}

func (s *SyncDisasterRecoveryCheckpointResponseBody) Validate() error {
	return dara.Validate(s)
}
