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
	// The policy end time.
	//
	// The time must be in UTC and in ISO 8601 format: `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-07T18:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The policy name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The IDs of the sites to associate with the policy. Separate multiple site IDs with commas.
	//
	// > This parameter is deprecated. Use `SiteIds` instead. If `SiteIds` is specified, the value of this parameter is ignored. To prevent ambiguity, specify a value for either `SiteIds` or `Objects`.
	//
	// example:
	//
	// 7096621098****,7096621099****
	Objects *string `json:"Objects,omitempty" xml:"Objects,omitempty"`
	// The IDs of the sites to associate with the policy. Separate multiple site IDs with commas.
	//
	// example:
	//
	// 7096621098****,7096621099****
	SiteIds *string `json:"SiteIds,omitempty" xml:"SiteIds,omitempty"`
	// The policy start time.
	//
	// The time must be in UTC and in ISO 8601 format: `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-07T17:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The template name. Valid value:
	//
	// - **promotion**: a policy for major events.
	//
	// This parameter is required.
	//
	// example:
	//
	// promotion
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
