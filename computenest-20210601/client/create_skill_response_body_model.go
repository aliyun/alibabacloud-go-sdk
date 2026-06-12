// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSkillResponseBody
	GetRequestId() *string
	SetSkillId(v string) *CreateSkillResponseBody
	GetSkillId() *string
}

type CreateSkillResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A17F9930-E2DC-5E87-B6D6-B0BCD2B00834
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the Skill.
	//
	// example:
	//
	// s-051j4ot2aerr5dyc4
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

func (s CreateSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSkillResponseBody) GetSkillId() *string {
	return s.SkillId
}

func (s *CreateSkillResponseBody) SetRequestId(v string) *CreateSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillResponseBody) SetSkillId(v string) *CreateSkillResponseBody {
	s.SkillId = &v
	return s
}

func (s *CreateSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
