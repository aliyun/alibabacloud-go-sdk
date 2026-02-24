// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceConfigurationTimelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAggregateResourceConfigurationTimelineResponseBody
	GetRequestId() *string
	SetResourceConfigurationTimeline(v *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) *GetAggregateResourceConfigurationTimelineResponseBody
	GetResourceConfigurationTimeline() *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline
}

type GetAggregateResourceConfigurationTimelineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ED9CD1B3-286C-4E05-A765-5E1E0B9BC2AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configuration history of the resource.
	ResourceConfigurationTimeline *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline `json:"ResourceConfigurationTimeline,omitempty" xml:"ResourceConfigurationTimeline,omitempty" type:"Struct"`
}

func (s GetAggregateResourceConfigurationTimelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceConfigurationTimelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceConfigurationTimelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateResourceConfigurationTimelineResponseBody) GetResourceConfigurationTimeline() *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	return s.ResourceConfigurationTimeline
}

func (s *GetAggregateResourceConfigurationTimelineResponseBody) SetRequestId(v string) *GetAggregateResourceConfigurationTimelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBody) SetResourceConfigurationTimeline(v *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) *GetAggregateResourceConfigurationTimelineResponseBody {
	s.ResourceConfigurationTimeline = v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBody) Validate() error {
	if s.ResourceConfigurationTimeline != nil {
		if err := s.ResourceConfigurationTimeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline struct {
	// A list of configuration changes.
	ConfigurationList []*GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList `json:"ConfigurationList,omitempty" xml:"ConfigurationList,omitempty" type:"Repeated"`
	// The maximum number of entries returned on each page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to start the next query.
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) GetConfigurationList() []*GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	return s.ConfigurationList
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) GetNextToken() *string {
	return s.NextToken
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) SetConfigurationList(v []*GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	s.ConfigurationList = v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) SetMaxResults(v int32) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	s.MaxResults = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) SetNextToken(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	s.NextToken = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) Validate() error {
	if s.ConfigurationList != nil {
		for _, item := range s.ConfigurationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList struct {
	// The ID of the Alibaba Cloud account that owns the resource.
	//
	// example:
	//
	// 100931896542****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The zone.
	//
	// example:
	//
	// cn-hangzhou-h
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// The time when the resource change snapshot was recorded. This is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1624961156000
	CaptureTime *string `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	// The details of the resource changes that triggered the evaluation.
	//
	// example:
	//
	// {\\"AccessControlList\\":[null,{\\"Grant\\":\\"private\\"}],\\"ServerSideEncryptionRule\\":[null,{\\"SSEAlgorithm\\":\\"None\\"}],\\"CreationDate\\":[null,\\"2021-06-29T10:05:12.000Z\\"],\\"Owner\\":[null,{\\"DisplayName\\":\\"100931896542****\\",\\"ID\\":\\"100931896542****\\"}],\\"BucketPolicy\\":[null,{\\"LogPrefix\\":\\"\\",\\"LogBucket\\":\\"\\"}],\\"StorageClass\\":[null,\\"Standard\\"],\\"ExtranetEndpoint\\":[null,\\"oss-cn-hangzhou.aliyuncs.com\\"],\\"DataRedundancyType\\":[null,\\"LRS\\"],\\"AllowEmptyReferer\\":[null,\\"true\\"],\\"IntranetEndpoint\\":[null,\\"oss-cn-hangzhou-internal.aliyuncs.com\\"],\\"Name\\":[null,\\"new-bucket\\"],\\"Location\\":[null,\\"oss-cn-hangzhou\\"]}
	ConfigurationDiff *string `json:"ConfigurationDiff,omitempty" xml:"ConfigurationDiff,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The time when the resource was created. This is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1624961112000
	ResourceCreateTime *string `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	// The type of the resource change event. Valid values:
	//
	// - DISCOVERED: The resource is discovered by Cloud Config.
	//
	// - DISCOVERED_REVISED: The resource is discovered by Cloud Config through a periodic remediation task.
	//
	// - MODIFY: The resource is modified.
	//
	// - MODIFY_REVISED: The resource is modified, as detected by a periodic remediation task.
	//
	// - REMOVE: The resource is deleted.
	//
	// > 	- To ensure data integrity, Cloud Config periodically runs remediation tasks to align data. This process may generate resource discovery events. These events occur infrequently.
	//
	// >
	//
	// > 	- The time when a resource event is generated by a periodic remediation task is the discovery time of the task. This time is later than the actual time of the resource change.
	//
	// example:
	//
	// DISCOVERED
	ResourceEventType *string `json:"ResourceEventType,omitempty" xml:"ResourceEventType,omitempty"`
	// The ID of the resource.
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
	// The resource type.
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

func (s GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetCaptureTime() *string {
	return s.CaptureTime
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetConfigurationDiff() *string {
	return s.ConfigurationDiff
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetRegion() *string {
	return s.Region
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetResourceCreateTime() *string {
	return s.ResourceCreateTime
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetResourceEventType() *string {
	return s.ResourceEventType
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetTags() *string {
	return s.Tags
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetAccountId(v int64) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.AccountId = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetAvailabilityZone(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.AvailabilityZone = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetCaptureTime(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.CaptureTime = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetConfigurationDiff(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ConfigurationDiff = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetRegion(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.Region = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceCreateTime(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceCreateTime = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceEventType(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceEventType = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceId(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceId = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceName(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceName = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceType(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceType = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetTags(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.Tags = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) Validate() error {
	return dara.Validate(s)
}
