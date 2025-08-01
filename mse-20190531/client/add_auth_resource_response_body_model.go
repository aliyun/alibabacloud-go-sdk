// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuthResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddAuthResourceResponseBody
	GetCode() *int32
	SetData(v int64) *AddAuthResourceResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *AddAuthResourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddAuthResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddAuthResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddAuthResourceResponseBody
	GetSuccess() *bool
}

type AddAuthResourceResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data structure.
	//
	// example:
	//
	// 24
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4BBCF560-4DD7-5DBD-B849-CCB135BBBAB7
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

func (s AddAuthResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAuthResourceResponseBody) GoString() string {
	return s.String()
}

func (s *AddAuthResourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddAuthResourceResponseBody) GetData() *int64 {
	return s.Data
}

func (s *AddAuthResourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddAuthResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddAuthResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAuthResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddAuthResourceResponseBody) SetCode(v int32) *AddAuthResourceResponseBody {
	s.Code = &v
	return s
}

func (s *AddAuthResourceResponseBody) SetData(v int64) *AddAuthResourceResponseBody {
	s.Data = &v
	return s
}

func (s *AddAuthResourceResponseBody) SetHttpStatusCode(v int32) *AddAuthResourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddAuthResourceResponseBody) SetMessage(v string) *AddAuthResourceResponseBody {
	s.Message = &v
	return s
}

func (s *AddAuthResourceResponseBody) SetRequestId(v string) *AddAuthResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAuthResourceResponseBody) SetSuccess(v bool) *AddAuthResourceResponseBody {
	s.Success = &v
	return s
}

func (s *AddAuthResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
