// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelApThirdAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *DelApThirdAppResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *DelApThirdAppResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *DelApThirdAppResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *DelApThirdAppResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DelApThirdAppResponseBody
	GetRequestId() *string
}

type DelApThirdAppResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DelApThirdAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DelApThirdAppResponseBody) GoString() string {
	return s.String()
}

func (s *DelApThirdAppResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *DelApThirdAppResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *DelApThirdAppResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DelApThirdAppResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DelApThirdAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DelApThirdAppResponseBody) SetData(v map[string]interface{}) *DelApThirdAppResponseBody {
	s.Data = v
	return s
}

func (s *DelApThirdAppResponseBody) SetErrorCode(v int32) *DelApThirdAppResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DelApThirdAppResponseBody) SetErrorMessage(v string) *DelApThirdAppResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DelApThirdAppResponseBody) SetIsSuccess(v bool) *DelApThirdAppResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DelApThirdAppResponseBody) SetRequestId(v string) *DelApThirdAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DelApThirdAppResponseBody) Validate() error {
	return dara.Validate(s)
}
