// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateProjectMemberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateProjectMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateProjectMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateProjectMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateProjectMemberResponseBody
	GetSuccess() *bool
}

type UpdateProjectMemberResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateProjectMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateProjectMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateProjectMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateProjectMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProjectMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateProjectMemberResponseBody) SetCode(v string) *UpdateProjectMemberResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProjectMemberResponseBody) SetHttpStatusCode(v int32) *UpdateProjectMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateProjectMemberResponseBody) SetMessage(v string) *UpdateProjectMemberResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateProjectMemberResponseBody) SetRequestId(v string) *UpdateProjectMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectMemberResponseBody) SetSuccess(v bool) *UpdateProjectMemberResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateProjectMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
