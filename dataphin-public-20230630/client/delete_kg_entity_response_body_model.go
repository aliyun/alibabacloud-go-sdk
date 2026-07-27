// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKgEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteKgEntityResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteKgEntityResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteKgEntityResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteKgEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteKgEntityResponseBody
	GetSuccess() *bool
}

type DeleteKgEntityResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteKgEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKgEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKgEntityResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteKgEntityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteKgEntityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteKgEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKgEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteKgEntityResponseBody) SetCode(v string) *DeleteKgEntityResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteKgEntityResponseBody) SetHttpStatusCode(v int32) *DeleteKgEntityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteKgEntityResponseBody) SetMessage(v string) *DeleteKgEntityResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteKgEntityResponseBody) SetRequestId(v string) *DeleteKgEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKgEntityResponseBody) SetSuccess(v bool) *DeleteKgEntityResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteKgEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
