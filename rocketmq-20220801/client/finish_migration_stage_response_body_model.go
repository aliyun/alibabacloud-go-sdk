// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishMigrationStageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FinishMigrationStageResponseBody
	GetCode() *string
	SetData(v bool) *FinishMigrationStageResponseBody
	GetData() *bool
	SetDynamicCode(v string) *FinishMigrationStageResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *FinishMigrationStageResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *FinishMigrationStageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *FinishMigrationStageResponseBody
	GetMessage() *string
	SetRequestId(v string) *FinishMigrationStageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FinishMigrationStageResponseBody
	GetSuccess() *bool
}

type FinishMigrationStageResponseBody struct {
	// example:
	//
	// InvalidConsumerGroupId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 16425867-C948-5A0C-9A24-5259727BE727
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FinishMigrationStageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FinishMigrationStageResponseBody) GoString() string {
	return s.String()
}

func (s *FinishMigrationStageResponseBody) GetCode() *string {
	return s.Code
}

func (s *FinishMigrationStageResponseBody) GetData() *bool {
	return s.Data
}

func (s *FinishMigrationStageResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *FinishMigrationStageResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *FinishMigrationStageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *FinishMigrationStageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FinishMigrationStageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FinishMigrationStageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FinishMigrationStageResponseBody) SetCode(v string) *FinishMigrationStageResponseBody {
	s.Code = &v
	return s
}

func (s *FinishMigrationStageResponseBody) SetData(v bool) *FinishMigrationStageResponseBody {
	s.Data = &v
	return s
}

func (s *FinishMigrationStageResponseBody) SetDynamicCode(v string) *FinishMigrationStageResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *FinishMigrationStageResponseBody) SetDynamicMessage(v string) *FinishMigrationStageResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *FinishMigrationStageResponseBody) SetHttpStatusCode(v int32) *FinishMigrationStageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FinishMigrationStageResponseBody) SetMessage(v string) *FinishMigrationStageResponseBody {
	s.Message = &v
	return s
}

func (s *FinishMigrationStageResponseBody) SetRequestId(v string) *FinishMigrationStageResponseBody {
	s.RequestId = &v
	return s
}

func (s *FinishMigrationStageResponseBody) SetSuccess(v bool) *FinishMigrationStageResponseBody {
	s.Success = &v
	return s
}

func (s *FinishMigrationStageResponseBody) Validate() error {
	return dara.Validate(s)
}
