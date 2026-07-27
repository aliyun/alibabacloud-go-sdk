// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKgRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteKgRelationResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteKgRelationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteKgRelationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteKgRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteKgRelationResponseBody
	GetSuccess() *bool
}

type DeleteKgRelationResponseBody struct {
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

func (s DeleteKgRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKgRelationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKgRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteKgRelationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteKgRelationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteKgRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKgRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteKgRelationResponseBody) SetCode(v string) *DeleteKgRelationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteKgRelationResponseBody) SetHttpStatusCode(v int32) *DeleteKgRelationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteKgRelationResponseBody) SetMessage(v string) *DeleteKgRelationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteKgRelationResponseBody) SetRequestId(v string) *DeleteKgRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKgRelationResponseBody) SetSuccess(v bool) *DeleteKgRelationResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteKgRelationResponseBody) Validate() error {
	return dara.Validate(s)
}
