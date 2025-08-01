// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNacosServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateNacosServiceResponseBody
	GetCode() *int32
	SetData(v string) *CreateNacosServiceResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateNacosServiceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateNacosServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateNacosServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateNacosServiceResponseBody
	GetSuccess() *bool
}

type CreateNacosServiceResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the node.
	//
	// example:
	//
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5B170A0D-2C5D-4CF8-B808-03966B86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateNacosServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNacosServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNacosServiceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateNacosServiceResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateNacosServiceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateNacosServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateNacosServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNacosServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateNacosServiceResponseBody) SetCode(v int32) *CreateNacosServiceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateNacosServiceResponseBody) SetData(v string) *CreateNacosServiceResponseBody {
	s.Data = &v
	return s
}

func (s *CreateNacosServiceResponseBody) SetHttpStatusCode(v int32) *CreateNacosServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateNacosServiceResponseBody) SetMessage(v string) *CreateNacosServiceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateNacosServiceResponseBody) SetRequestId(v string) *CreateNacosServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNacosServiceResponseBody) SetSuccess(v bool) *CreateNacosServiceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateNacosServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
