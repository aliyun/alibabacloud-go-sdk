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
	// The end time of the policy.
	//
	// The time must be in UTC and in the ISO 8601 format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-03T19:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The site IDs to associate with the policy. Use a comma (,) to separate multiple IDs.
	//
	// > This parameter is deprecated. We recommend using the `SiteIds` parameter instead. If the `SiteIds` parameter is specified, the `Objects` parameter is ignored. You must specify a value for either the `Objects` or `SiteIds` parameter.
	//
	// example:
	//
	// 123456****,123457****
	Objects *string `json:"Objects,omitempty" xml:"Objects,omitempty"`
	// To obtain the policy ID, call the [DescribeCustomScenePolicies](https://help.aliyun.com/document_detail/2850508.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The site IDs to associate with the policy. Use a comma (,) to separate multiple IDs.
	//
	// example:
	//
	// 123456****,123457****
	SiteIds *string `json:"SiteIds,omitempty" xml:"SiteIds,omitempty"`
	// The start time of the policy.
	//
	// The time must be in UTC and in the ISO 8601 format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-03T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the template. Valid value:
	//
	// - **promotion**: major promotion
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
