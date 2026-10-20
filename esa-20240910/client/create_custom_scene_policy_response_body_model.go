// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomScenePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *CreateCustomScenePolicyResponseBody
	GetEndTime() *string
	SetName(v string) *CreateCustomScenePolicyResponseBody
	GetName() *string
	SetObjects(v []*string) *CreateCustomScenePolicyResponseBody
	GetObjects() []*string
	SetPolicyId(v int64) *CreateCustomScenePolicyResponseBody
	GetPolicyId() *int64
	SetRequestId(v string) *CreateCustomScenePolicyResponseBody
	GetRequestId() *string
	SetSiteIds(v string) *CreateCustomScenePolicyResponseBody
	GetSiteIds() *string
	SetStartTime(v string) *CreateCustomScenePolicyResponseBody
	GetStartTime() *string
	SetTemplate(v string) *CreateCustomScenePolicyResponseBody
	GetTemplate() *string
}

type CreateCustomScenePolicyResponseBody struct {
	// The end time of the policy.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2021-11-07T18:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The policy name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of IDs of the associated websites.
	//
	// > **[Deprecated]*	- This parameter is deprecated. Use the SiteIds parameter instead.
	Objects []*string `json:"Objects,omitempty" xml:"Objects,omitempty" type:"Repeated"`
	// The ID of the policy.
	//
	// example:
	//
	// 1
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDs of the associated websites. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// 7096621098****, 7096621099****
	SiteIds *string `json:"SiteIds,omitempty" xml:"SiteIds,omitempty"`
	// The start time of the policy.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2021-11-07T17:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// promotion
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s CreateCustomScenePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomScenePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomScenePolicyResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateCustomScenePolicyResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateCustomScenePolicyResponseBody) GetObjects() []*string {
	return s.Objects
}

func (s *CreateCustomScenePolicyResponseBody) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *CreateCustomScenePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomScenePolicyResponseBody) GetSiteIds() *string {
	return s.SiteIds
}

func (s *CreateCustomScenePolicyResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateCustomScenePolicyResponseBody) GetTemplate() *string {
	return s.Template
}

func (s *CreateCustomScenePolicyResponseBody) SetEndTime(v string) *CreateCustomScenePolicyResponseBody {
	s.EndTime = &v
	return s
}

func (s *CreateCustomScenePolicyResponseBody) SetName(v string) *CreateCustomScenePolicyResponseBody {
	s.Name = &v
	return s
}

func (s *CreateCustomScenePolicyResponseBody) SetObjects(v []*string) *CreateCustomScenePolicyResponseBody {
	s.Objects = v
	return s
}

func (s *CreateCustomScenePolicyResponseBody) SetPolicyId(v int64) *CreateCustomScenePolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreateCustomScenePolicyResponseBody) SetRequestId(v string) *CreateCustomScenePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomScenePolicyResponseBody) SetSiteIds(v string) *CreateCustomScenePolicyResponseBody {
	s.SiteIds = &v
	return s
}

func (s *CreateCustomScenePolicyResponseBody) SetStartTime(v string) *CreateCustomScenePolicyResponseBody {
	s.StartTime = &v
	return s
}

func (s *CreateCustomScenePolicyResponseBody) SetTemplate(v string) *CreateCustomScenePolicyResponseBody {
	s.Template = &v
	return s
}

func (s *CreateCustomScenePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
