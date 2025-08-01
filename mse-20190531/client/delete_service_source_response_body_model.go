// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteServiceSourceResponseBody
	GetCode() *int32
	SetHttpStatusCode(v int32) *DeleteServiceSourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteServiceSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteServiceSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteServiceSourceResponseBody
	GetSuccess() *bool
}

type DeleteServiceSourceResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 1
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2CEF593F-D60C-5449-9E98-15CA6ECD9189
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

func (s DeleteServiceSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceSourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteServiceSourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteServiceSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteServiceSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteServiceSourceResponseBody) SetCode(v int32) *DeleteServiceSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteServiceSourceResponseBody) SetHttpStatusCode(v int32) *DeleteServiceSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteServiceSourceResponseBody) SetMessage(v string) *DeleteServiceSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceSourceResponseBody) SetRequestId(v string) *DeleteServiceSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceSourceResponseBody) SetSuccess(v bool) *DeleteServiceSourceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteServiceSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
