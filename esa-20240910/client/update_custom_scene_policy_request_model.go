// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomScenePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *UpdateCustomScenePolicyRequest
	GetEndTime() *string
	SetName(v string) *UpdateCustomScenePolicyRequest
	GetName() *string
	SetObjects(v string) *UpdateCustomScenePolicyRequest
	GetObjects() *string
	SetPolicyId(v int64) *UpdateCustomScenePolicyRequest
	GetPolicyId() *int64
	SetSiteIds(v string) *UpdateCustomScenePolicyRequest
	GetSiteIds() *string
	SetStartTime(v string) *UpdateCustomScenePolicyRequest
	GetStartTime() *string
	SetTemplate(v string) *UpdateCustomScenePolicyRequest
	GetTemplate() *string
}

type UpdateCustomScenePolicyRequest struct {
	// This parameter is required.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Objects *string `json:"Objects,omitempty" xml:"Objects,omitempty"`
	// This parameter is required.
	PolicyId *int64  `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	SiteIds  *string `json:"SiteIds,omitempty" xml:"SiteIds,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s UpdateCustomScenePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomScenePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomScenePolicyRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateCustomScenePolicyRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCustomScenePolicyRequest) GetObjects() *string {
	return s.Objects
}

func (s *UpdateCustomScenePolicyRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *UpdateCustomScenePolicyRequest) GetSiteIds() *string {
	return s.SiteIds
}

func (s *UpdateCustomScenePolicyRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateCustomScenePolicyRequest) GetTemplate() *string {
	return s.Template
}

func (s *UpdateCustomScenePolicyRequest) SetEndTime(v string) *UpdateCustomScenePolicyRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateCustomScenePolicyRequest) SetName(v string) *UpdateCustomScenePolicyRequest {
	s.Name = &v
	return s
}

func (s *UpdateCustomScenePolicyRequest) SetObjects(v string) *UpdateCustomScenePolicyRequest {
	s.Objects = &v
	return s
}

func (s *UpdateCustomScenePolicyRequest) SetPolicyId(v int64) *UpdateCustomScenePolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateCustomScenePolicyRequest) SetSiteIds(v string) *UpdateCustomScenePolicyRequest {
	s.SiteIds = &v
	return s
}

func (s *UpdateCustomScenePolicyRequest) SetStartTime(v string) *UpdateCustomScenePolicyRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateCustomScenePolicyRequest) SetTemplate(v string) *UpdateCustomScenePolicyRequest {
	s.Template = &v
	return s
}

func (s *UpdateCustomScenePolicyRequest) Validate() error {
	return dara.Validate(s)
}
