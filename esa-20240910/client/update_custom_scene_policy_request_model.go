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
	// The time when the policy expires.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-03T19:00:00Z
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
	// 123456****
	Objects *string `json:"Objects,omitempty" xml:"Objects,omitempty"`
	// The policy ID, which can be obtained by calling the [DescribeCustomScenePolicies](https://help.aliyun.com/document_detail/2850508.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PolicyId *int64  `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	SiteIds  *string `json:"SiteIds,omitempty" xml:"SiteIds,omitempty"`
	// The time when the policy takes effect.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-03T16:00:00Z
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
