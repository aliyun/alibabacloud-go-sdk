// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNacosServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteNacosServiceResponseBody
	GetCode() *int32
	SetData(v string) *DeleteNacosServiceResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteNacosServiceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteNacosServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteNacosServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteNacosServiceResponseBody
	GetSuccess() *bool
}

type DeleteNacosServiceResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The deletion result.
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

func (s DeleteNacosServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNacosServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNacosServiceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteNacosServiceResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteNacosServiceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteNacosServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteNacosServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNacosServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteNacosServiceResponseBody) SetCode(v int32) *DeleteNacosServiceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteNacosServiceResponseBody) SetData(v string) *DeleteNacosServiceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteNacosServiceResponseBody) SetHttpStatusCode(v int32) *DeleteNacosServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteNacosServiceResponseBody) SetMessage(v string) *DeleteNacosServiceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteNacosServiceResponseBody) SetRequestId(v string) *DeleteNacosServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNacosServiceResponseBody) SetSuccess(v bool) *DeleteNacosServiceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteNacosServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
