// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBISchemaIndexDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBISchemaIndexDeleteResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBISchemaIndexDeleteResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBISchemaIndexDeleteResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBISchemaIndexDeleteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBISchemaIndexDeleteResponseBody
	GetSuccess() *bool
}

type ChatBISchemaIndexDeleteResponseBody struct {
	// example:
	//
	// true/false
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// null
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// null
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BADFB626-FB3E-5220-806B-ACF5A625109A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChatBISchemaIndexDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBISchemaIndexDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBISchemaIndexDeleteResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBISchemaIndexDeleteResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBISchemaIndexDeleteResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBISchemaIndexDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBISchemaIndexDeleteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBISchemaIndexDeleteResponseBody) SetData(v interface{}) *ChatBISchemaIndexDeleteResponseBody {
	s.Data = v
	return s
}

func (s *ChatBISchemaIndexDeleteResponseBody) SetErrCode(v string) *ChatBISchemaIndexDeleteResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBISchemaIndexDeleteResponseBody) SetErrMessage(v string) *ChatBISchemaIndexDeleteResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBISchemaIndexDeleteResponseBody) SetRequestId(v string) *ChatBISchemaIndexDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBISchemaIndexDeleteResponseBody) SetSuccess(v bool) *ChatBISchemaIndexDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBISchemaIndexDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
