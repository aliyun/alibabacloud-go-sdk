// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIFileTemplateDownloadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIFileTemplateDownloadResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIFileTemplateDownloadResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIFileTemplateDownloadResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIFileTemplateDownloadResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIFileTemplateDownloadResponseBody
	GetSuccess() *bool
}

type ChatBIFileTemplateDownloadResponseBody struct {
	// example:
	//
	// http url
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

func (s ChatBIFileTemplateDownloadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIFileTemplateDownloadResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIFileTemplateDownloadResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIFileTemplateDownloadResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIFileTemplateDownloadResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIFileTemplateDownloadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIFileTemplateDownloadResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIFileTemplateDownloadResponseBody) SetData(v interface{}) *ChatBIFileTemplateDownloadResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIFileTemplateDownloadResponseBody) SetErrCode(v string) *ChatBIFileTemplateDownloadResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIFileTemplateDownloadResponseBody) SetErrMessage(v string) *ChatBIFileTemplateDownloadResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIFileTemplateDownloadResponseBody) SetRequestId(v string) *ChatBIFileTemplateDownloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIFileTemplateDownloadResponseBody) SetSuccess(v bool) *ChatBIFileTemplateDownloadResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIFileTemplateDownloadResponseBody) Validate() error {
	return dara.Validate(s)
}
