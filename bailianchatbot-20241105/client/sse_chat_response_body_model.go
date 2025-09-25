// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSseChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SseChatResponseBody
	GetCode() *string
	SetData(v interface{}) *SseChatResponseBody
	GetData() interface{}
	SetMessage(v string) *SseChatResponseBody
	GetMessage() *string
	SetRequestId(v string) *SseChatResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SseChatResponseBody
	GetSuccess() *bool
}

type SseChatResponseBody struct {
	// example:
	//
	// 200
	Code *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 30D68C4C-4512-58A7-A328-568015B39A02
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SseChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SseChatResponseBody) GoString() string {
	return s.String()
}

func (s *SseChatResponseBody) GetCode() *string {
	return s.Code
}

func (s *SseChatResponseBody) GetData() interface{} {
	return s.Data
}

func (s *SseChatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SseChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SseChatResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SseChatResponseBody) SetCode(v string) *SseChatResponseBody {
	s.Code = &v
	return s
}

func (s *SseChatResponseBody) SetData(v interface{}) *SseChatResponseBody {
	s.Data = v
	return s
}

func (s *SseChatResponseBody) SetMessage(v string) *SseChatResponseBody {
	s.Message = &v
	return s
}

func (s *SseChatResponseBody) SetRequestId(v string) *SseChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *SseChatResponseBody) SetSuccess(v bool) *SseChatResponseBody {
	s.Success = &v
	return s
}

func (s *SseChatResponseBody) Validate() error {
	return dara.Validate(s)
}
