// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWebSocketSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateWebSocketSignResponseBody
	GetCode() *string
	SetData(v string) *GenerateWebSocketSignResponseBody
	GetData() *string
	SetHttpStatusCode(v int64) *GenerateWebSocketSignResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *GenerateWebSocketSignResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateWebSocketSignResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateWebSocketSignResponseBody
	GetSuccess() *bool
}

type GenerateWebSocketSignResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// dnthF_oinHg7JMJDmKqex3UxDD7
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateWebSocketSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateWebSocketSignResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateWebSocketSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateWebSocketSignResponseBody) GetData() *string {
	return s.Data
}

func (s *GenerateWebSocketSignResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GenerateWebSocketSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateWebSocketSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateWebSocketSignResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateWebSocketSignResponseBody) SetCode(v string) *GenerateWebSocketSignResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetData(v string) *GenerateWebSocketSignResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetHttpStatusCode(v int64) *GenerateWebSocketSignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetMessage(v string) *GenerateWebSocketSignResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetRequestId(v string) *GenerateWebSocketSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetSuccess(v bool) *GenerateWebSocketSignResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) Validate() error {
	return dara.Validate(s)
}
