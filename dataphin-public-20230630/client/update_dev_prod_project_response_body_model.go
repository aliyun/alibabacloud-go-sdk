// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDevProdProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDevProdProjectResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateDevProdProjectResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateDevProdProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDevProdProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDevProdProjectResponseBody
	GetSuccess() *bool
}

type UpdateDevProdProjectResponseBody struct {
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

func (s UpdateDevProdProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDevProdProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDevProdProjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDevProdProjectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateDevProdProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDevProdProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDevProdProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDevProdProjectResponseBody) SetCode(v string) *UpdateDevProdProjectResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDevProdProjectResponseBody) SetHttpStatusCode(v int32) *UpdateDevProdProjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDevProdProjectResponseBody) SetMessage(v string) *UpdateDevProdProjectResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDevProdProjectResponseBody) SetRequestId(v string) *UpdateDevProdProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDevProdProjectResponseBody) SetSuccess(v bool) *UpdateDevProdProjectResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDevProdProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
