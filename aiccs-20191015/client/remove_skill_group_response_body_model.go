// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveSkillGroupResponseBody
	GetCode() *string
	SetMessage(v string) *RemoveSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveSkillGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveSkillGroupResponseBody
	GetSuccess() *bool
}

type RemoveSkillGroupResponseBody struct {
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

func (s RemoveSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveSkillGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveSkillGroupResponseBody) SetCode(v string) *RemoveSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveSkillGroupResponseBody) SetMessage(v string) *RemoveSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveSkillGroupResponseBody) SetRequestId(v string) *RemoveSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSkillGroupResponseBody) SetSuccess(v bool) *RemoveSkillGroupResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
