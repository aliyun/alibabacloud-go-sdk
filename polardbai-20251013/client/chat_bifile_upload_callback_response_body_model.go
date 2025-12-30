// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIFileUploadCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIFileUploadCallbackResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIFileUploadCallbackResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIFileUploadCallbackResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIFileUploadCallbackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIFileUploadCallbackResponseBody
	GetSuccess() *bool
}

type ChatBIFileUploadCallbackResponseBody struct {
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

func (s ChatBIFileUploadCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIFileUploadCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIFileUploadCallbackResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIFileUploadCallbackResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIFileUploadCallbackResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIFileUploadCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIFileUploadCallbackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIFileUploadCallbackResponseBody) SetData(v interface{}) *ChatBIFileUploadCallbackResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIFileUploadCallbackResponseBody) SetErrCode(v string) *ChatBIFileUploadCallbackResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIFileUploadCallbackResponseBody) SetErrMessage(v string) *ChatBIFileUploadCallbackResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIFileUploadCallbackResponseBody) SetRequestId(v string) *ChatBIFileUploadCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIFileUploadCallbackResponseBody) SetSuccess(v bool) *ChatBIFileUploadCallbackResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIFileUploadCallbackResponseBody) Validate() error {
	return dara.Validate(s)
}
