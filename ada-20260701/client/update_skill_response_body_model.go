// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateSkillResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateSkillResponseBody
	GetRequestId() *string
	SetSkillId(v string) *UpdateSkillResponseBody
	GetSkillId() *string
	SetSuccess(v bool) *UpdateSkillResponseBody
	GetSuccess() *bool
	SetUpdatedAt(v int64) *UpdateSkillResponseBody
	GetUpdatedAt() *int64
}

type UpdateSkillResponseBody struct {
	// The name of the skill.
	//
	// example:
	//
	// code-review
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The skill ID.
	//
	// example:
	//
	// skill_example123
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// Indicates whether the operation is successful. A value of `true` is returned when the skill update and optional bundle replacement are fully submitted and confirmed by read-back.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The time when the skill was updated. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1760000200000
	UpdatedAt *int64 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s UpdateSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSkillResponseBody) GetSkillId() *string {
	return s.SkillId
}

func (s *UpdateSkillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSkillResponseBody) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *UpdateSkillResponseBody) SetName(v string) *UpdateSkillResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateSkillResponseBody) SetRequestId(v string) *UpdateSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillResponseBody) SetSkillId(v string) *UpdateSkillResponseBody {
	s.SkillId = &v
	return s
}

func (s *UpdateSkillResponseBody) SetSuccess(v bool) *UpdateSkillResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSkillResponseBody) SetUpdatedAt(v int64) *UpdateSkillResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
