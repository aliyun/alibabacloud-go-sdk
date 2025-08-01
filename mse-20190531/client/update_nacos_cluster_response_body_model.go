// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacosClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateNacosClusterResponseBody
	GetCode() *int32
	SetData(v string) *UpdateNacosClusterResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateNacosClusterResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateNacosClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateNacosClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateNacosClusterResponseBody
	GetSuccess() *bool
}

type UpdateNacosClusterResponseBody struct {
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

func (s UpdateNacosClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacosClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNacosClusterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateNacosClusterResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateNacosClusterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateNacosClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateNacosClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNacosClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateNacosClusterResponseBody) SetCode(v int32) *UpdateNacosClusterResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNacosClusterResponseBody) SetData(v string) *UpdateNacosClusterResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateNacosClusterResponseBody) SetHttpStatusCode(v int32) *UpdateNacosClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateNacosClusterResponseBody) SetMessage(v string) *UpdateNacosClusterResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNacosClusterResponseBody) SetRequestId(v string) *UpdateNacosClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNacosClusterResponseBody) SetSuccess(v bool) *UpdateNacosClusterResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNacosClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
