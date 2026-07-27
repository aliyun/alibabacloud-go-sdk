// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKgRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateKgRelationResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateKgRelationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateKgRelationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateKgRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateKgRelationResponseBody
	GetSuccess() *bool
}

type UpdateKgRelationResponseBody struct {
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

func (s UpdateKgRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKgRelationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKgRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateKgRelationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateKgRelationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateKgRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKgRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateKgRelationResponseBody) SetCode(v string) *UpdateKgRelationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateKgRelationResponseBody) SetHttpStatusCode(v int32) *UpdateKgRelationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateKgRelationResponseBody) SetMessage(v string) *UpdateKgRelationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateKgRelationResponseBody) SetRequestId(v string) *UpdateKgRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKgRelationResponseBody) SetSuccess(v bool) *UpdateKgRelationResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateKgRelationResponseBody) Validate() error {
	return dara.Validate(s)
}
