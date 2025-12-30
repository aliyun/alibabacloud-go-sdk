// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIFileUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIFileUploadResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIFileUploadResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIFileUploadResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIFileUploadResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIFileUploadResponseBody
	GetSuccess() *bool
}

type ChatBIFileUploadResponseBody struct {
	// example:
	//
	// 返回viper所需信息
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

func (s ChatBIFileUploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIFileUploadResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIFileUploadResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIFileUploadResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIFileUploadResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIFileUploadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIFileUploadResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIFileUploadResponseBody) SetData(v interface{}) *ChatBIFileUploadResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIFileUploadResponseBody) SetErrCode(v string) *ChatBIFileUploadResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIFileUploadResponseBody) SetErrMessage(v string) *ChatBIFileUploadResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIFileUploadResponseBody) SetRequestId(v string) *ChatBIFileUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIFileUploadResponseBody) SetSuccess(v bool) *ChatBIFileUploadResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIFileUploadResponseBody) Validate() error {
	return dara.Validate(s)
}
