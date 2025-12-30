// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigUpdateEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIConfigUpdateEntryResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIConfigUpdateEntryResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIConfigUpdateEntryResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIConfigUpdateEntryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIConfigUpdateEntryResponseBody
	GetSuccess() *bool
}

type ChatBIConfigUpdateEntryResponseBody struct {
	// example:
	//
	// true/false
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 错误码
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// 错误信息
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 983863E2-****-1D15-A4AE-3468CD75383D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChatBIConfigUpdateEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigUpdateEntryResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIConfigUpdateEntryResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIConfigUpdateEntryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIConfigUpdateEntryResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIConfigUpdateEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIConfigUpdateEntryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIConfigUpdateEntryResponseBody) SetData(v interface{}) *ChatBIConfigUpdateEntryResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIConfigUpdateEntryResponseBody) SetErrCode(v string) *ChatBIConfigUpdateEntryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIConfigUpdateEntryResponseBody) SetErrMessage(v string) *ChatBIConfigUpdateEntryResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIConfigUpdateEntryResponseBody) SetRequestId(v string) *ChatBIConfigUpdateEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIConfigUpdateEntryResponseBody) SetSuccess(v bool) *ChatBIConfigUpdateEntryResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIConfigUpdateEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
