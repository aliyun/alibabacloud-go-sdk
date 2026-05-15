// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSkillGroupResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSkillGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSkillGroupResponseBody
	GetSuccess() *bool
}

type UpdateSkillGroupResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSkillGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSkillGroupResponseBody) SetCode(v string) *UpdateSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSkillGroupResponseBody) SetMessage(v string) *UpdateSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSkillGroupResponseBody) SetRequestId(v string) *UpdateSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillGroupResponseBody) SetSuccess(v bool) *UpdateSkillGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
