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
	// The time when the policy expires.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
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
	// The IDs of the websites that you want to associate with the policy. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// 7096621098****
	Objects *string `json:"Objects,omitempty" xml:"Objects,omitempty"`
	SiteIds *string `json:"SiteIds,omitempty" xml:"SiteIds,omitempty"`
	// The time when the policy takes effect.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-07T17:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the policy template. Valid value:
	//
	// 	- **promotion**: major events.
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
