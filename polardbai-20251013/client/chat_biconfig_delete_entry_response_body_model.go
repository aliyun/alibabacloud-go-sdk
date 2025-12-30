// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigDeleteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIConfigDeleteEntryResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIConfigDeleteEntryResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIConfigDeleteEntryResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIConfigDeleteEntryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIConfigDeleteEntryResponseBody
	GetSuccess() *bool
}

type ChatBIConfigDeleteEntryResponseBody struct {
	// example:
	//
	// true/false
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 具体错误码
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// 具体错误信息
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

func (s ChatBIConfigDeleteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigDeleteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIConfigDeleteEntryResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIConfigDeleteEntryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIConfigDeleteEntryResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIConfigDeleteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIConfigDeleteEntryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIConfigDeleteEntryResponseBody) SetData(v interface{}) *ChatBIConfigDeleteEntryResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIConfigDeleteEntryResponseBody) SetErrCode(v string) *ChatBIConfigDeleteEntryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIConfigDeleteEntryResponseBody) SetErrMessage(v string) *ChatBIConfigDeleteEntryResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIConfigDeleteEntryResponseBody) SetRequestId(v string) *ChatBIConfigDeleteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIConfigDeleteEntryResponseBody) SetSuccess(v bool) *ChatBIConfigDeleteEntryResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIConfigDeleteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
