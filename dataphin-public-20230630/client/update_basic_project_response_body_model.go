// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateBasicProjectResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateBasicProjectResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateBasicProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateBasicProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateBasicProjectResponseBody
	GetSuccess() *bool
}

type UpdateBasicProjectResponseBody struct {
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

func (s UpdateBasicProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBasicProjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateBasicProjectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateBasicProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateBasicProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBasicProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBasicProjectResponseBody) SetCode(v string) *UpdateBasicProjectResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateBasicProjectResponseBody) SetHttpStatusCode(v int32) *UpdateBasicProjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateBasicProjectResponseBody) SetMessage(v string) *UpdateBasicProjectResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBasicProjectResponseBody) SetRequestId(v string) *UpdateBasicProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBasicProjectResponseBody) SetSuccess(v bool) *UpdateBasicProjectResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBasicProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
