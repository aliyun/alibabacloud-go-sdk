// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawSkillDetailRequest
	GetApplicationId() *string
	SetSlug(v string) *DescribePolarClawSkillDetailRequest
	GetSlug() *string
}

type DescribePolarClawSkillDetailRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The Skill identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// alibacloud-rds-copilot
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
}

func (s DescribePolarClawSkillDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillDetailRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawSkillDetailRequest) GetSlug() *string {
	return s.Slug
}

func (s *DescribePolarClawSkillDetailRequest) SetApplicationId(v string) *DescribePolarClawSkillDetailRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawSkillDetailRequest) SetSlug(v string) *DescribePolarClawSkillDetailRequest {
	s.Slug = &v
	return s
}

func (s *DescribePolarClawSkillDetailRequest) Validate() error {
	return dara.Validate(s)
}
