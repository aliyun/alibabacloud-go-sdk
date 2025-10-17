// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyJobStepCheckpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyJobStepCheckpointResponseBody
	GetCode() *string
	SetDynamicMessage(v string) *ModifyJobStepCheckpointResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *ModifyJobStepCheckpointResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyJobStepCheckpointResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyJobStepCheckpointResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyJobStepCheckpointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyJobStepCheckpointResponseBody
	GetSuccess() *bool
}

type ModifyJobStepCheckpointResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// can not find env: zbyk-pre
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// C306C198-7807-409D-930A-D6CE6C32****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyJobStepCheckpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyJobStepCheckpointResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyJobStepCheckpointResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyJobStepCheckpointResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyJobStepCheckpointResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyJobStepCheckpointResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyJobStepCheckpointResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyJobStepCheckpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyJobStepCheckpointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyJobStepCheckpointResponseBody) SetCode(v string) *ModifyJobStepCheckpointResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyJobStepCheckpointResponseBody) SetDynamicMessage(v string) *ModifyJobStepCheckpointResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyJobStepCheckpointResponseBody) SetErrCode(v string) *ModifyJobStepCheckpointResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyJobStepCheckpointResponseBody) SetErrMessage(v string) *ModifyJobStepCheckpointResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyJobStepCheckpointResponseBody) SetHttpStatusCode(v int32) *ModifyJobStepCheckpointResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyJobStepCheckpointResponseBody) SetRequestId(v string) *ModifyJobStepCheckpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyJobStepCheckpointResponseBody) SetSuccess(v bool) *ModifyJobStepCheckpointResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyJobStepCheckpointResponseBody) Validate() error {
	return dara.Validate(s)
}
