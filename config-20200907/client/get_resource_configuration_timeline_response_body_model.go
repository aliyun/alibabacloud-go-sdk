// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceConfigurationTimelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResourceConfigurationTimelineResponseBody
	GetRequestId() *string
	SetResourceConfigurationTimeline(v *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) *GetResourceConfigurationTimelineResponseBody
	GetResourceConfigurationTimeline() *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline
}

type GetResourceConfigurationTimelineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ED9CD1B3-286C-4E05-A765-5E1E0B9BC2AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configuration timeline of the resource.
	ResourceConfigurationTimeline *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline `json:"ResourceConfigurationTimeline,omitempty" xml:"ResourceConfigurationTimeline,omitempty" type:"Struct"`
}

func (s GetResourceConfigurationTimelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceConfigurationTimelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationTimelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceConfigurationTimelineResponseBody) GetResourceConfigurationTimeline() *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	return s.ResourceConfigurationTimeline
}

func (s *GetResourceConfigurationTimelineResponseBody) SetRequestId(v string) *GetResourceConfigurationTimelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBody) SetResourceConfigurationTimeline(v *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) *GetResourceConfigurationTimelineResponseBody {
	s.ResourceConfigurationTimeline = v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBody) Validate() error {
	if s.ResourceConfigurationTimeline != nil {
		if err := s.ResourceConfigurationTimeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline struct {
	// The configuration changes on the configuration timeline.
	ConfigurationList []*GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList `json:"ConfigurationList,omitempty" xml:"ConfigurationList,omitempty" type:"Repeated"`
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
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) String() string {
	return dara.Prettify(s)
}

func (s GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) GetConfigurationList() []*GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	return s.ConfigurationList
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) GetNextToken() *string {
	return s.NextToken
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) SetConfigurationList(v []*GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	s.ConfigurationList = v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) SetMaxResults(v int32) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	s.MaxResults = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) SetNextToken(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	s.NextToken = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) Validate() error {
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

type GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 100931896542****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// The time when the resource change snapshot was recorded. Unit: milliseconds.
	//
	// example:
	//
	// 1624961156000
	CaptureTime *string `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	// The details of the resource changes that trigger the compliance evaluation.
	//
	// example:
	//
	// {\\"AccessControlList\\":[null,{\\"Grant\\":\\"private\\"}],\\"ServerSideEncryptionRule\\":[null,{\\"SSEAlgorithm\\":\\"None\\"}],\\"CreationDate\\":[null,\\"2021-06-29T10:05:12.000Z\\"],\\"Owner\\":[null,{\\"DisplayName\\":\\"100931896542****\\",\\"ID\\":\\"100931896542****\\"}],\\"BucketPolicy\\":[null,{\\"LogPrefix\\":\\"\\",\\"LogBucket\\":\\"\\"}],\\"StorageClass\\":[null,\\"Standard\\"],\\"ExtranetEndpoint\\":[null,\\"oss-cn-hangzhou.aliyuncs.com\\"],\\"DataRedundancyType\\":[null,\\"LRS\\"],\\"AllowEmptyReferer\\":[null,\\"true\\"],\\"IntranetEndpoint\\":[null,\\"oss-cn-hangzhou-internal.aliyuncs.com\\"],\\"Name\\":[null,\\"new-bucket\\"],\\"Location\\":[null,\\"oss-cn-hangzhou\\"]}
	ConfigurationDiff *string `json:"ConfigurationDiff,omitempty" xml:"ConfigurationDiff,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The details of each resource that is associated with the current resource, including the region ID, resource relationship, resource ID, and resource type.
	//
	// example:
	//
	// ""
	Relationship *string `json:"Relationship,omitempty" xml:"Relationship,omitempty"`
	// The change records of the resource relationship.
	//
	// example:
	//
	// ""
	RelationshipDiff *string `json:"RelationshipDiff,omitempty" xml:"RelationshipDiff,omitempty"`
	// The time when the resource was created. Unit: milliseconds.
	//
	// example:
	//
	// 1624961112000
	ResourceCreateTime *string `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	// The type of the resource change event. Valid values:
	//
	// 	- DISCOVERED: A resource is created.
	//
	// 	- DISCOVERED_REVISED: A resource is created by periodic remediation tasks.
	//
	// 	- MODIFY: A resource is modified.
	//
	// 	- MODIFY_REVISED: A resource is modified by periodic remediation tasks.
	//
	// 	- REMOVE: A resource is deleted.
	//
	// >
	//
	// 	- To ensure the integrity of resources, periodic remediation tasks are run to check data and generate events that indicate the creation of new resources. Such events are infrequent.
	//
	// 	- The time when a resource change event is generated by a periodic remediation task is considered as the detection time of Cloud Config. The detection time is later than the time when the resource is modified.
	//
	// example:
	//
	// DISCOVERED
	ResourceEventType *string `json:"ResourceEventType,omitempty" xml:"ResourceEventType,omitempty"`
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

func (s GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) String() string {
	return dara.Prettify(s)
}

func (s GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetCaptureTime() *string {
	return s.CaptureTime
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetConfigurationDiff() *string {
	return s.ConfigurationDiff
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetRegion() *string {
	return s.Region
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetRelationship() *string {
	return s.Relationship
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetRelationshipDiff() *string {
	return s.RelationshipDiff
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetResourceCreateTime() *string {
	return s.ResourceCreateTime
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetResourceEventType() *string {
	return s.ResourceEventType
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GetTags() *string {
	return s.Tags
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetAccountId(v int64) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.AccountId = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetAvailabilityZone(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.AvailabilityZone = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetCaptureTime(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.CaptureTime = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetConfigurationDiff(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ConfigurationDiff = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetRegion(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.Region = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetRelationship(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.Relationship = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetRelationshipDiff(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.RelationshipDiff = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceCreateTime(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceCreateTime = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceEventType(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceEventType = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceId(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceId = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceName(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceName = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceType(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceType = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetTags(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.Tags = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) Validate() error {
	return dara.Validate(s)
}
