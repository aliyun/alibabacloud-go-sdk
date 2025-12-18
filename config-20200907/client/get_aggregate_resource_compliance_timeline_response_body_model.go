// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceComplianceTimelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAggregateResourceComplianceTimelineResponseBody
	GetRequestId() *string
	SetResourceComplianceTimeline(v *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) *GetAggregateResourceComplianceTimelineResponseBody
	GetResourceComplianceTimeline() *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline
}

type GetAggregateResourceComplianceTimelineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8D53A78F-1EB8-4264-A554-72F07E34FAE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the compliance timeline.
	ResourceComplianceTimeline *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline `json:"ResourceComplianceTimeline,omitempty" xml:"ResourceComplianceTimeline,omitempty" type:"Struct"`
}

func (s GetAggregateResourceComplianceTimelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceTimelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceTimelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateResourceComplianceTimelineResponseBody) GetResourceComplianceTimeline() *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	return s.ResourceComplianceTimeline
}

func (s *GetAggregateResourceComplianceTimelineResponseBody) SetRequestId(v string) *GetAggregateResourceComplianceTimelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBody) SetResourceComplianceTimeline(v *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) *GetAggregateResourceComplianceTimelineResponseBody {
	s.ResourceComplianceTimeline = v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBody) Validate() error {
	if s.ResourceComplianceTimeline != nil {
		if err := s.ResourceComplianceTimeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline struct {
	// The status of the resource. The parameter value varies based on the resource type and may be left empty. Examples:
	//
	// 	- If the value of the ResourceType parameter is ACS::ECS::Instance, the resource is an Elastic Compute Service (ECS) instance that has a specific state. In this case, the valid values of this parameter are Running and Stopped.
	//
	// 	- If the value of the ResourceType parameter is ACS::OSS::Bucket, the resource is an Object Storage Service (OSS) bucket that is not in a specific state. In this case, this parameter is empty.
	ComplianceList []*GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList `json:"ComplianceList,omitempty" xml:"ComplianceList,omitempty" type:"Repeated"`
	// The maximum number of entries returned for a single request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 5OVS5J4I1/UKTkHV5oNs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) GetComplianceList() []*GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	return s.ComplianceList
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) GetNextToken() *string {
	return s.NextToken
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) SetComplianceList(v []*GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	s.ComplianceList = v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) SetMaxResults(v int32) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	s.MaxResults = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) SetNextToken(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	s.NextToken = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) Validate() error {
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

type GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 100931896542****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the zone in which the resource resides.
	//
	// example:
	//
	// cn-hangzhou-f
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// The timestamp when the compliance evaluation was recorded. Unit: milliseconds.
	//
	// example:
	//
	// 1625200295276
	CaptureTime *int64 `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	// The information about the rules that evaluated the resource and the compliance evaluation result.
	//
	// example:
	//
	// {
	//
	//     "AccessControlList": {
	//
	//         "Grant": "private"
	//
	//     },
	//
	//     "ServerSideEncryptionRule": {
	//
	//         "SSEAlgorithm": "None"
	//
	//     },
	//
	//     "Comment": "",
	//
	//     "Owner": {
	//
	//         "DisplayName": "126672004088****",
	//
	//         "ID": "126672004088****"
	//
	//     },
	//
	//     "ResourceGroupId": "rg-acfmy4dfoau****",
	//
	//     "StorageClass": "Standard",
	//
	//     "DataRedundancyType": "LRS",
	//
	//     "AllowEmptyReferer": "true",
	//
	//     "Name": "test-private-oss****",
	//
	//     "BucketPolicy": {
	//
	//         "LogPrefix": "",
	//
	//         "LogBucket": ""
	//
	//     },
	//
	//     "ExtranetEndpoint": "oss-cn-hangzhou.aliyuncs.com",
	//
	//     "IntranetEndpoint": "oss-cn-hangzhou-internal.aliyuncs.com",
	//
	//     "Location": "oss-cn-hangzhou"
	//
	// }
	Configuration *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The details of the resource change that triggered the compliance evaluation.
	//
	// example:
	//
	// {
	//
	//     "ResourceGroupId": [
	//
	//         "rg-acfmy4dfoau****",
	//
	//         "rg-a3dmy4lksta****"
	//
	//     ]
	//
	// }
	ConfigurationDiff *string `json:"ConfigurationDiff,omitempty" xml:"ConfigurationDiff,omitempty"`
	// The ID of the region in which the resource resides.
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
	// The ID of the resource.
	//
	// example:
	//
	// test-private-oss****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// test-private-oss****
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The status of the resource. The parameter value varies based on the resource type and may be left empty. Examples:
	//
	// 	- If the ResourceType parameter is set to ACS::ECS::Instance, the resource is an Elastic Compute Service (ECS) instance that has a specific state. In this case, the valid values of this parameter are Running and Stopped.
	//
	// 	- If the ResourceType parameter is set to ACS::OSS::Bucket, the resource is an OSS bucket that does not have a specific state. In this case, this parameter is left empty.
	//
	// example:
	//
	// Running
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ACS::OSS::Bucket
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	//
	// example:
	//
	// {\\"\\"hc\\"\\":[\\"\\"value2\\"\\"]}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetAccountId() *string {
	return s.AccountId
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetCaptureTime() *int64 {
	return s.CaptureTime
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetConfiguration() *string {
	return s.Configuration
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetConfigurationDiff() *string {
	return s.ConfigurationDiff
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetRegion() *string {
	return s.Region
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetResourceCreateTime() *int64 {
	return s.ResourceCreateTime
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GetTags() *string {
	return s.Tags
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetAccountId(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.AccountId = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetAvailabilityZone(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.AvailabilityZone = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetCaptureTime(v int64) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.CaptureTime = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetConfiguration(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.Configuration = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetConfigurationDiff(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ConfigurationDiff = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetRegion(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.Region = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceCreateTime(v int64) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceCreateTime = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceId(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceId = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceName(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceName = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceStatus(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceStatus = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceType(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceType = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetTags(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.Tags = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) Validate() error {
	return dara.Validate(s)
}
