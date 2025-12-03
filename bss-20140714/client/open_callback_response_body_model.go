// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OpenCallbackResponseBody
	GetCode() *string
	SetData(v string) *OpenCallbackResponseBody
	GetData() *string
	SetMessage(v string) *OpenCallbackResponseBody
	GetMessage() *string
	SetRequestId(v string) *OpenCallbackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OpenCallbackResponseBody
	GetSuccess() *bool
}

type OpenCallbackResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OpenCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *OpenCallbackResponseBody) GetCode() *string {
	return s.Code
}

func (s *OpenCallbackResponseBody) GetData() *string {
	return s.Data
}

func (s *OpenCallbackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OpenCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenCallbackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OpenCallbackResponseBody) SetCode(v string) *OpenCallbackResponseBody {
	s.Code = &v
	return s
}

func (s *OpenCallbackResponseBody) SetData(v string) *OpenCallbackResponseBody {
	s.Data = &v
	return s
}

func (s *OpenCallbackResponseBody) SetMessage(v string) *OpenCallbackResponseBody {
	s.Message = &v
	return s
}

func (s *OpenCallbackResponseBody) SetRequestId(v string) *OpenCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenCallbackResponseBody) SetSuccess(v bool) *OpenCallbackResponseBody {
	s.Success = &v
	return s
}

func (s *OpenCallbackResponseBody) Validate() error {
	return dara.Validate(s)
}
