// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveProjectMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveProjectMemberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RemoveProjectMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemoveProjectMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveProjectMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveProjectMemberResponseBody
	GetSuccess() *bool
}

type RemoveProjectMemberResponseBody struct {
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

func (s RemoveProjectMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveProjectMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveProjectMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveProjectMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveProjectMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveProjectMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveProjectMemberResponseBody) SetCode(v string) *RemoveProjectMemberResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveProjectMemberResponseBody) SetHttpStatusCode(v int32) *RemoveProjectMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveProjectMemberResponseBody) SetMessage(v string) *RemoveProjectMemberResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveProjectMemberResponseBody) SetRequestId(v string) *RemoveProjectMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveProjectMemberResponseBody) SetSuccess(v bool) *RemoveProjectMemberResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveProjectMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
