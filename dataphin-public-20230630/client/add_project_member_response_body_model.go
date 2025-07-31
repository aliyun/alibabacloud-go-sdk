// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddProjectMemberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddProjectMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddProjectMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddProjectMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddProjectMemberResponseBody
	GetSuccess() *bool
}

type AddProjectMemberResponseBody struct {
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

func (s AddProjectMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddProjectMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddProjectMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddProjectMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddProjectMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddProjectMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddProjectMemberResponseBody) SetCode(v string) *AddProjectMemberResponseBody {
	s.Code = &v
	return s
}

func (s *AddProjectMemberResponseBody) SetHttpStatusCode(v int32) *AddProjectMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddProjectMemberResponseBody) SetMessage(v string) *AddProjectMemberResponseBody {
	s.Message = &v
	return s
}

func (s *AddProjectMemberResponseBody) SetRequestId(v string) *AddProjectMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddProjectMemberResponseBody) SetSuccess(v bool) *AddProjectMemberResponseBody {
	s.Success = &v
	return s
}

func (s *AddProjectMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
