// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRetryStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateInstanceRetryStrategyResponseBody
	GetCode() *int32
	SetData(v bool) *UpdateInstanceRetryStrategyResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateInstanceRetryStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateInstanceRetryStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateInstanceRetryStrategyResponseBody
	GetSuccess() *bool
}

type UpdateInstanceRetryStrategyResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceRetryStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRetryStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRetryStrategyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateInstanceRetryStrategyResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateInstanceRetryStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateInstanceRetryStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceRetryStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateInstanceRetryStrategyResponseBody) SetCode(v int32) *UpdateInstanceRetryStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceRetryStrategyResponseBody) SetData(v bool) *UpdateInstanceRetryStrategyResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceRetryStrategyResponseBody) SetMessage(v string) *UpdateInstanceRetryStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceRetryStrategyResponseBody) SetRequestId(v string) *UpdateInstanceRetryStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceRetryStrategyResponseBody) SetSuccess(v bool) *UpdateInstanceRetryStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceRetryStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
