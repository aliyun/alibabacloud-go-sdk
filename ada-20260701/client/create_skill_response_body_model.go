// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateSkillResponseBody
	GetName() *string
	SetRequestId(v string) *CreateSkillResponseBody
	GetRequestId() *string
	SetSkillId(v string) *CreateSkillResponseBody
	GetSkillId() *string
	SetSuccess(v bool) *CreateSkillResponseBody
	GetSuccess() *bool
	SetUpdatedAt(v int64) *CreateSkillResponseBody
	GetUpdatedAt() *int64
}

type CreateSkillResponseBody struct {
	// The unique identifier of the Skill.
	//
	// example:
	//
	// code-review
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID, which is used to locate and troubleshoot the request.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Skill ID。
	//
	// example:
	//
	// skill_example123
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// Indicates whether the Skill and its body or bundle are fully created and readable. A value of true is returned upon success. Business failures are returned as error responses.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The time when the Skill was last updated after creation. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1760000100000
	UpdatedAt *int64 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s CreateSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSkillResponseBody) GetSkillId() *string {
	return s.SkillId
}

func (s *CreateSkillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSkillResponseBody) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *CreateSkillResponseBody) SetName(v string) *CreateSkillResponseBody {
	s.Name = &v
	return s
}

func (s *CreateSkillResponseBody) SetRequestId(v string) *CreateSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillResponseBody) SetSkillId(v string) *CreateSkillResponseBody {
	s.SkillId = &v
	return s
}

func (s *CreateSkillResponseBody) SetSuccess(v bool) *CreateSkillResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSkillResponseBody) SetUpdatedAt(v int64) *CreateSkillResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *CreateSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
