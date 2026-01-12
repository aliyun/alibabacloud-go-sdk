// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateProjectResponseBody
	GetCode() *string
	SetErrorName(v string) *UpdateProjectResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *UpdateProjectResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *UpdateProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateProjectResponseBody
	GetSuccess() *bool
}

type UpdateProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateProjectResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *UpdateProjectResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateProjectResponseBody) SetCode(v string) *UpdateProjectResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProjectResponseBody) SetErrorName(v string) *UpdateProjectResponseBody {
	s.ErrorName = &v
	return s
}

func (s *UpdateProjectResponseBody) SetHttpCode(v int32) *UpdateProjectResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateProjectResponseBody) SetMessage(v string) *UpdateProjectResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectResponseBody) SetSuccess(v bool) *UpdateProjectResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
