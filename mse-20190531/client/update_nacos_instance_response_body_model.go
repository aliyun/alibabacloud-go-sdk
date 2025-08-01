// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacosInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateNacosInstanceResponseBody
	GetCode() *int32
	SetData(v string) *UpdateNacosInstanceResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateNacosInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateNacosInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateNacosInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateNacosInstanceResponseBody
	GetSuccess() *bool
}

type UpdateNacosInstanceResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the modification.
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
	// The returned message.
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
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNacosInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacosInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNacosInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateNacosInstanceResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateNacosInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateNacosInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateNacosInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNacosInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateNacosInstanceResponseBody) SetCode(v int32) *UpdateNacosInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNacosInstanceResponseBody) SetData(v string) *UpdateNacosInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateNacosInstanceResponseBody) SetHttpStatusCode(v int32) *UpdateNacosInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateNacosInstanceResponseBody) SetMessage(v string) *UpdateNacosInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNacosInstanceResponseBody) SetRequestId(v string) *UpdateNacosInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNacosInstanceResponseBody) SetSuccess(v bool) *UpdateNacosInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNacosInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
