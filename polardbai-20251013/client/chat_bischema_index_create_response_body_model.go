// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBISchemaIndexCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBISchemaIndexCreateResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBISchemaIndexCreateResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBISchemaIndexCreateResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBISchemaIndexCreateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBISchemaIndexCreateResponseBody
	GetSuccess() *bool
}

type ChatBISchemaIndexCreateResponseBody struct {
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

func (s ChatBISchemaIndexCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBISchemaIndexCreateResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBISchemaIndexCreateResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBISchemaIndexCreateResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBISchemaIndexCreateResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBISchemaIndexCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBISchemaIndexCreateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBISchemaIndexCreateResponseBody) SetData(v interface{}) *ChatBISchemaIndexCreateResponseBody {
	s.Data = v
	return s
}

func (s *ChatBISchemaIndexCreateResponseBody) SetErrCode(v string) *ChatBISchemaIndexCreateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBISchemaIndexCreateResponseBody) SetErrMessage(v string) *ChatBISchemaIndexCreateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBISchemaIndexCreateResponseBody) SetRequestId(v string) *ChatBISchemaIndexCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBISchemaIndexCreateResponseBody) SetSuccess(v bool) *ChatBISchemaIndexCreateResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBISchemaIndexCreateResponseBody) Validate() error {
	return dara.Validate(s)
}
