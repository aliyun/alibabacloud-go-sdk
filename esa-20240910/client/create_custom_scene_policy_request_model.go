// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomScenePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *CreateCustomScenePolicyRequest
	GetEndTime() *string
	SetName(v string) *CreateCustomScenePolicyRequest
	GetName() *string
	SetObjects(v string) *CreateCustomScenePolicyRequest
	GetObjects() *string
	SetSiteIds(v string) *CreateCustomScenePolicyRequest
	GetSiteIds() *string
	SetStartTime(v string) *CreateCustomScenePolicyRequest
	GetStartTime() *string
	SetTemplate(v string) *CreateCustomScenePolicyRequest
	GetTemplate() *string
}

type CreateCustomScenePolicyRequest struct {
	// This parameter is required.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Objects *string `json:"Objects,omitempty" xml:"Objects,omitempty"`
	SiteIds *string `json:"SiteIds,omitempty" xml:"SiteIds,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s CreateCustomScenePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomScenePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomScenePolicyRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateCustomScenePolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreateCustomScenePolicyRequest) GetObjects() *string {
	return s.Objects
}

func (s *CreateCustomScenePolicyRequest) GetSiteIds() *string {
	return s.SiteIds
}

func (s *CreateCustomScenePolicyRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateCustomScenePolicyRequest) GetTemplate() *string {
	return s.Template
}

func (s *CreateCustomScenePolicyRequest) SetEndTime(v string) *CreateCustomScenePolicyRequest {
	s.EndTime = &v
	return s
}

func (s *CreateCustomScenePolicyRequest) SetName(v string) *CreateCustomScenePolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomScenePolicyRequest) SetObjects(v string) *CreateCustomScenePolicyRequest {
	s.Objects = &v
	return s
}

func (s *CreateCustomScenePolicyRequest) SetSiteIds(v string) *CreateCustomScenePolicyRequest {
	s.SiteIds = &v
	return s
}

func (s *CreateCustomScenePolicyRequest) SetStartTime(v string) *CreateCustomScenePolicyRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCustomScenePolicyRequest) SetTemplate(v string) *CreateCustomScenePolicyRequest {
	s.Template = &v
	return s
}

func (s *CreateCustomScenePolicyRequest) Validate() error {
	return dara.Validate(s)
}
