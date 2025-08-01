// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNacosInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateNacosInstanceResponseBody
	GetCode() *int32
	SetData(v string) *CreateNacosInstanceResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateNacosInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateNacosInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateNacosInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateNacosInstanceResponseBody
	GetSuccess() *bool
}

type CreateNacosInstanceResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result returned.
	//
	// example:
	//
	// ok
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9e78a671-4b9b-4dd4-99c1-0b9da87d3dec
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateNacosInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNacosInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNacosInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateNacosInstanceResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateNacosInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateNacosInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateNacosInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNacosInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateNacosInstanceResponseBody) SetCode(v int32) *CreateNacosInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateNacosInstanceResponseBody) SetData(v string) *CreateNacosInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *CreateNacosInstanceResponseBody) SetHttpStatusCode(v int32) *CreateNacosInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateNacosInstanceResponseBody) SetMessage(v string) *CreateNacosInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateNacosInstanceResponseBody) SetRequestId(v string) *CreateNacosInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNacosInstanceResponseBody) SetSuccess(v bool) *CreateNacosInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateNacosInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
