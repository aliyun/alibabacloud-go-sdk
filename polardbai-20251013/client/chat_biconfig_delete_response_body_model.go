// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIConfigDeleteResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIConfigDeleteResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIConfigDeleteResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIConfigDeleteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIConfigDeleteResponseBody
	GetSuccess() *bool
}

type ChatBIConfigDeleteResponseBody struct {
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

func (s ChatBIConfigDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIConfigDeleteResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIConfigDeleteResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIConfigDeleteResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIConfigDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIConfigDeleteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIConfigDeleteResponseBody) SetData(v interface{}) *ChatBIConfigDeleteResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIConfigDeleteResponseBody) SetErrCode(v string) *ChatBIConfigDeleteResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIConfigDeleteResponseBody) SetErrMessage(v string) *ChatBIConfigDeleteResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIConfigDeleteResponseBody) SetRequestId(v string) *ChatBIConfigDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIConfigDeleteResponseBody) SetSuccess(v bool) *ChatBIConfigDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIConfigDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
