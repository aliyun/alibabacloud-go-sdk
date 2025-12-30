// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIConfigCreateResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIConfigCreateResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIConfigCreateResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIConfigCreateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIConfigCreateResponseBody
	GetSuccess() *bool
}

type ChatBIConfigCreateResponseBody struct {
	// example:
	//
	// true
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

func (s ChatBIConfigCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigCreateResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIConfigCreateResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIConfigCreateResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIConfigCreateResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIConfigCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIConfigCreateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIConfigCreateResponseBody) SetData(v interface{}) *ChatBIConfigCreateResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIConfigCreateResponseBody) SetErrCode(v string) *ChatBIConfigCreateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIConfigCreateResponseBody) SetErrMessage(v string) *ChatBIConfigCreateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIConfigCreateResponseBody) SetRequestId(v string) *ChatBIConfigCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIConfigCreateResponseBody) SetSuccess(v bool) *ChatBIConfigCreateResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIConfigCreateResponseBody) Validate() error {
	return dara.Validate(s)
}
