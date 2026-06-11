// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v string) *UpdateSkillLabelsRequest
	GetLabels() *string
	SetNamespaceId(v string) *UpdateSkillLabelsRequest
	GetNamespaceId() *string
	SetSkillName(v string) *UpdateSkillLabelsRequest
	GetSkillName() *string
}

type UpdateSkillLabelsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"latest":"0.0.2","stable":"0.0.1"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// customer-service-skill
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
}

func (s UpdateSkillLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillLabelsRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillLabelsRequest) GetLabels() *string {
	return s.Labels
}

func (s *UpdateSkillLabelsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateSkillLabelsRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *UpdateSkillLabelsRequest) SetLabels(v string) *UpdateSkillLabelsRequest {
	s.Labels = &v
	return s
}

func (s *UpdateSkillLabelsRequest) SetNamespaceId(v string) *UpdateSkillLabelsRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateSkillLabelsRequest) SetSkillName(v string) *UpdateSkillLabelsRequest {
	s.SkillName = &v
	return s
}

func (s *UpdateSkillLabelsRequest) Validate() error {
	return dara.Validate(s)
}
