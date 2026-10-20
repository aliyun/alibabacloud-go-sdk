// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomScenePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *UpdateCustomScenePolicyResponseBody
	GetEndTime() *string
	SetName(v string) *UpdateCustomScenePolicyResponseBody
	GetName() *string
	SetObjects(v []*string) *UpdateCustomScenePolicyResponseBody
	GetObjects() []*string
	SetPolicyId(v int64) *UpdateCustomScenePolicyResponseBody
	GetPolicyId() *int64
	SetRequestId(v string) *UpdateCustomScenePolicyResponseBody
	GetRequestId() *string
	SetSiteIds(v string) *UpdateCustomScenePolicyResponseBody
	GetSiteIds() *string
	SetStartTime(v string) *UpdateCustomScenePolicyResponseBody
	GetStartTime() *string
	SetTemplate(v string) *UpdateCustomScenePolicyResponseBody
	GetTemplate() *string
}

type UpdateCustomScenePolicyResponseBody struct {
	// The end time of the policy.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2023-04-03T19:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The policy name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of associated website IDs.
	//
	// > **[Deprecated]*	- We recommend that you read the value of SiteIds instead.
	Objects []*string `json:"Objects,omitempty" xml:"Objects,omitempty" type:"Repeated"`
	// The ID of the policy.
	//
	// example:
	//
	// 1
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDs of the websites to associate. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// 123456****,123457****
	SiteIds *string `json:"SiteIds,omitempty" xml:"SiteIds,omitempty"`
	// The start time of the policy.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2023-04-03T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the template. Valid values:
	//
	// - **promotion**: Important Activity.
	//
	// example:
	//
	// promotion
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s UpdateCustomScenePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomScenePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomScenePolicyResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateCustomScenePolicyResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateCustomScenePolicyResponseBody) GetObjects() []*string {
	return s.Objects
}

func (s *UpdateCustomScenePolicyResponseBody) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *UpdateCustomScenePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomScenePolicyResponseBody) GetSiteIds() *string {
	return s.SiteIds
}

func (s *UpdateCustomScenePolicyResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateCustomScenePolicyResponseBody) GetTemplate() *string {
	return s.Template
}

func (s *UpdateCustomScenePolicyResponseBody) SetEndTime(v string) *UpdateCustomScenePolicyResponseBody {
	s.EndTime = &v
	return s
}

func (s *UpdateCustomScenePolicyResponseBody) SetName(v string) *UpdateCustomScenePolicyResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateCustomScenePolicyResponseBody) SetObjects(v []*string) *UpdateCustomScenePolicyResponseBody {
	s.Objects = v
	return s
}

func (s *UpdateCustomScenePolicyResponseBody) SetPolicyId(v int64) *UpdateCustomScenePolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *UpdateCustomScenePolicyResponseBody) SetRequestId(v string) *UpdateCustomScenePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomScenePolicyResponseBody) SetSiteIds(v string) *UpdateCustomScenePolicyResponseBody {
	s.SiteIds = &v
	return s
}

func (s *UpdateCustomScenePolicyResponseBody) SetStartTime(v string) *UpdateCustomScenePolicyResponseBody {
	s.StartTime = &v
	return s
}

func (s *UpdateCustomScenePolicyResponseBody) SetTemplate(v string) *UpdateCustomScenePolicyResponseBody {
	s.Template = &v
	return s
}

func (s *UpdateCustomScenePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
