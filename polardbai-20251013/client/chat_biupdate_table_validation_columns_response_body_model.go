// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIUpdateTableValidationColumnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIUpdateTableValidationColumnsResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIUpdateTableValidationColumnsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIUpdateTableValidationColumnsResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIUpdateTableValidationColumnsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIUpdateTableValidationColumnsResponseBody
	GetSuccess() *bool
}

type ChatBIUpdateTableValidationColumnsResponseBody struct {
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

func (s ChatBIUpdateTableValidationColumnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIUpdateTableValidationColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIUpdateTableValidationColumnsResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIUpdateTableValidationColumnsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIUpdateTableValidationColumnsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIUpdateTableValidationColumnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIUpdateTableValidationColumnsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIUpdateTableValidationColumnsResponseBody) SetData(v interface{}) *ChatBIUpdateTableValidationColumnsResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIUpdateTableValidationColumnsResponseBody) SetErrCode(v string) *ChatBIUpdateTableValidationColumnsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIUpdateTableValidationColumnsResponseBody) SetErrMessage(v string) *ChatBIUpdateTableValidationColumnsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIUpdateTableValidationColumnsResponseBody) SetRequestId(v string) *ChatBIUpdateTableValidationColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIUpdateTableValidationColumnsResponseBody) SetSuccess(v bool) *ChatBIUpdateTableValidationColumnsResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIUpdateTableValidationColumnsResponseBody) Validate() error {
	return dara.Validate(s)
}
