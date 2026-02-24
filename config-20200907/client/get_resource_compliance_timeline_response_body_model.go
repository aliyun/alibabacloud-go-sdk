// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceComplianceTimelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResourceComplianceTimelineResponseBody
	GetRequestId() *string
	SetResourceComplianceTimeline(v *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) *GetResourceComplianceTimelineResponseBody
	GetResourceComplianceTimeline() *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline
}

type GetResourceComplianceTimelineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8D53A78F-1EB8-4264-A554-72F07E34FAE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The compliance evaluation history of the resource.
	ResourceComplianceTimeline *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline `json:"ResourceComplianceTimeline,omitempty" xml:"ResourceComplianceTimeline,omitempty" type:"Struct"`
}

func (s GetResourceComplianceTimelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceTimelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceTimelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceComplianceTimelineResponseBody) GetResourceComplianceTimeline() *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	return s.ResourceComplianceTimeline
}

func (s *GetResourceComplianceTimelineResponseBody) SetRequestId(v string) *GetResourceComplianceTimelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBody) SetResourceComplianceTimeline(v *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) *GetResourceComplianceTimelineResponseBody {
	s.ResourceComplianceTimeline = v
	return s
}

func (s *GetResourceComplianceTimelineResponseBody) Validate() error {
	if s.ResourceComplianceTimeline != nil {
		if err := s.ResourceComplianceTimeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline struct {
	// The list of compliance evaluation history records.
	ComplianceList []*GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList `json:"ComplianceList,omitempty" xml:"ComplianceList,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that you use to retrieve the next page of results.
	//
	// example:
	//
	// 5OVS5J4I1/UKTkHV5oNs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) GetComplianceList() []*GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	return s.ComplianceList
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) GetNextToken() *string {
	return s.NextToken
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) SetComplianceList(v []*GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	s.ComplianceList = v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) SetMaxResults(v int32) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	s.MaxResults = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) SetNextToken(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	s.NextToken = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) Validate() error {
	if s.ComplianceList != nil {
		for _, item := range s.ComplianceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 100931896542****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The zone where the resource resides.
	//
	// example:
	//
	// cn-hangzhou-f
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// The timestamp when the compliance evaluation of the resource was recorded. Unit: milliseconds.
	//
	// example:
	//
	// 1625200295276
	CaptureTime *int64 `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	// The list of rules associated with the resource and the compliance details of the rules.
	//
	// example:
	//
	// {\\"Compliance\\":{\\"complianceType\\":\\"COMPLIANT\\",\\"count\\":1},\\"ConfigRuleList\\":[{\\"accountId\\":100931896542****,\\"configRuleId\\":\\"cr-9524626622af003d****\\",\\"configRuleArn\\":\\"acs:config::100931896542****:rule/cr-9524626622af003d****\\",\\"configRuleName\\":\\"OSS存储空间ACL禁止公共读写\\",\\"complianceType\\":\\"COMPLIANT\\",\\"riskLevel\\":1,\\"annotation\\":\\"\\",\\"invokingEventMessageType\\":\\"ScheduledNotification\\"}]}
	Configuration *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The details of the resource change that triggered this evaluation.
	//
	// example:
	//
	// {\\"OSS存储空间ACL禁止公共读写\\":[{\\"accountId\\":100931896542****,\\"configRuleId\\":\\"cr-965f626622af003d****\\",\\"configRuleArn\\":\\"acs:config::100931896542****:rule/cr-965f626622af003d****\\",\\"configRuleName\\":\\"OSS存储空间ACL禁止公共读写\\",\\"complianceType\\":\\"COMPLIANT\\",\\"riskLevel\\":1,\\"annotation\\":\\"\\",\\"invokingEventMessageType\\":\\"ScheduledNotification\\"},{}]}
	ConfigurationDiff *string `json:"ConfigurationDiff,omitempty" xml:"ConfigurationDiff,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The timestamp when the resource was created. Unit: milliseconds.
	//
	// example:
	//
	// 1624961112000
	ResourceCreateTime *int64 `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// new-bucket
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// new-bucket
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource status. The resource status is defined by each Alibaba Cloud service. This parameter can be empty. For example:
	//
	// - If the resource type is \\`ACS::ECS::Instance\\`, this parameter can be \\`Running\\` or \\`Stopped\\` because ECS instances are stateful.
	//
	// - If the resource type is \\`ACS::OSS::Bucket\\`, this parameter is empty because OSS buckets are stateless.
	//
	// example:
	//
	// null
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::OSS::Bucket
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The resource tags.
	//
	// example:
	//
	// {\\"\\"hc\\"\\":[\\"\\"value2\\"\\"]}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetAccountId() *string {
	return s.AccountId
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetCaptureTime() *int64 {
	return s.CaptureTime
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetConfiguration() *string {
	return s.Configuration
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetConfigurationDiff() *string {
	return s.ConfigurationDiff
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetRegion() *string {
	return s.Region
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetResourceCreateTime() *int64 {
	return s.ResourceCreateTime
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetTags() *string {
	return s.Tags
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetAccountId(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.AccountId = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetAvailabilityZone(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.AvailabilityZone = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetCaptureTime(v int64) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.CaptureTime = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetConfiguration(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.Configuration = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetConfigurationDiff(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ConfigurationDiff = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetRegion(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.Region = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceCreateTime(v int64) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceCreateTime = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceId(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceId = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceName(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceName = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceStatus(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceStatus = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceType(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceType = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetTags(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.Tags = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) Validate() error {
	return dara.Validate(s)
}
