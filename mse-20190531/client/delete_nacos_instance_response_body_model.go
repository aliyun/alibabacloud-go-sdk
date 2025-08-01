// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNacosInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteNacosInstanceResponseBody
	GetCode() *int32
	SetData(v string) *DeleteNacosInstanceResponseBody
	GetData() *string
	SetDynamicMessage(v string) *DeleteNacosInstanceResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *DeleteNacosInstanceResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *DeleteNacosInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteNacosInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteNacosInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteNacosInstanceResponseBody
	GetSuccess() *bool
}

type DeleteNacosInstanceResponseBody struct {
	// The status code. The value 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the node.
	//
	// example:
	//
	// OK
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8BD1E58D-0755-42AC-A599-E6B55112****
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

func (s DeleteNacosInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNacosInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNacosInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteNacosInstanceResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteNacosInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteNacosInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteNacosInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteNacosInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteNacosInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNacosInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteNacosInstanceResponseBody) SetCode(v int32) *DeleteNacosInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteNacosInstanceResponseBody) SetData(v string) *DeleteNacosInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteNacosInstanceResponseBody) SetDynamicMessage(v string) *DeleteNacosInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteNacosInstanceResponseBody) SetErrorCode(v string) *DeleteNacosInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteNacosInstanceResponseBody) SetHttpStatusCode(v int32) *DeleteNacosInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteNacosInstanceResponseBody) SetMessage(v string) *DeleteNacosInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteNacosInstanceResponseBody) SetRequestId(v string) *DeleteNacosInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNacosInstanceResponseBody) SetSuccess(v bool) *DeleteNacosInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteNacosInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
