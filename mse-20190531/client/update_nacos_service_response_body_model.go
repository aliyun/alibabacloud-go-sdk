// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacosServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateNacosServiceResponseBody
	GetCode() *int32
	SetData(v string) *UpdateNacosServiceResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateNacosServiceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateNacosServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateNacosServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateNacosServiceResponseBody
	GetSuccess() *bool
}

type UpdateNacosServiceResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 1
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data.
	//
	// example:
	//
	// 6
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
	// 4E9FDCFE-0738-493B-B801-82BDFBCB****
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

func (s UpdateNacosServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacosServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNacosServiceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateNacosServiceResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateNacosServiceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateNacosServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateNacosServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNacosServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateNacosServiceResponseBody) SetCode(v int32) *UpdateNacosServiceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNacosServiceResponseBody) SetData(v string) *UpdateNacosServiceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateNacosServiceResponseBody) SetHttpStatusCode(v int32) *UpdateNacosServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateNacosServiceResponseBody) SetMessage(v string) *UpdateNacosServiceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNacosServiceResponseBody) SetRequestId(v string) *UpdateNacosServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNacosServiceResponseBody) SetSuccess(v bool) *UpdateNacosServiceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNacosServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
