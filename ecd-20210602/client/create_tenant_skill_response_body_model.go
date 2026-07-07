// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTenantSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTenantSkillResponseBody
	GetRequestId() *string
	SetSkillId(v string) *CreateTenantSkillResponseBody
	GetSkillId() *string
}

type CreateTenantSkillResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The unique ID of the skill.
	//
	// example:
	//
	// s-04rj8mzqj1fu****
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

func (s CreateTenantSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTenantSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTenantSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTenantSkillResponseBody) GetSkillId() *string {
	return s.SkillId
}

func (s *CreateTenantSkillResponseBody) SetRequestId(v string) *CreateTenantSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTenantSkillResponseBody) SetSkillId(v string) *CreateTenantSkillResponseBody {
	s.SkillId = &v
	return s
}

func (s *CreateTenantSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
