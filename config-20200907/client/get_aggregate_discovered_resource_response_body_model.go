// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateDiscoveredResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiscoveredResourceDetail(v *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) *GetAggregateDiscoveredResourceResponseBody
	GetDiscoveredResourceDetail() *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail
	SetRequestId(v string) *GetAggregateDiscoveredResourceResponseBody
	GetRequestId() *string
}

type GetAggregateDiscoveredResourceResponseBody struct {
	// The details of the resource.
	DiscoveredResourceDetail *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail `json:"DiscoveredResourceDetail,omitempty" xml:"DiscoveredResourceDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E4D71ACE-6B0A-46E0-8352-56952378CC7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateDiscoveredResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateDiscoveredResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateDiscoveredResourceResponseBody) GetDiscoveredResourceDetail() *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	return s.DiscoveredResourceDetail
}

func (s *GetAggregateDiscoveredResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateDiscoveredResourceResponseBody) SetDiscoveredResourceDetail(v *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) *GetAggregateDiscoveredResourceResponseBody {
	s.DiscoveredResourceDetail = v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBody) SetRequestId(v string) *GetAggregateDiscoveredResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBody) Validate() error {
	if s.DiscoveredResourceDetail != nil {
		if err := s.DiscoveredResourceDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 100931896542****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the zone in which the resource resides.
	//
	// example:
	//
	// cn-hangzhou-h
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// The compliance evaluation result of the resource. Valid values:
	//
	// 	- COMPLIANT: The resource is evaluated as compliant.
	//
	// 	- NON_COMPLIANT: The resource is evaluated as non-compliant.
	//
	// 	- NOT_APPLICABLE: The rule does not apply to the resource.
	//
	// 	- INSUFFICIENT_DATA: No data is available.
	//
	// 	- IGNORED: The resource is ignored during compliance evaluation.
	//
	// example:
	//
	// COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The configuration of the resource.
	//
	// example:
	//
	// {\\"AccessControlList\\":{\\"Grant\\":\\"private\\"},\\"ServerSideEncryptionRule\\":{\\"SSEAlgorithm\\":\\"None\\"},\\"Comment\\":\\"\\",\\"CreationDate\\":\\"2021-06-29T10:05:12.000Z\\",\\"Owner\\":{\\"DisplayName\\":\\"100931896542****\\",\\"ID\\":\\"100931896542****\\"},\\"StorageClass\\":\\"Standard\\",\\"DataRedundancyType\\":\\"LRS\\",\\"AllowEmptyReferer\\":\\"true\\",\\"Name\\":\\"new-bucket\\",\\"BucketPolicy\\":{\\"LogPrefix\\":\\"\\",\\"LogBucket\\":\\"\\"},\\"ExtranetEndpoint\\":\\"oss-cn-hangzhou.aliyuncs.com\\",\\"IntranetEndpoint\\":\\"oss-cn-hangzhou-internal.aliyuncs.com\\",\\"Location\\":\\"oss-cn-hangzhou\\"}
	Configuration *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The timestamp when the resource was created.
	//
	// example:
	//
	// 1624961112000
	ResourceCreationTime *int64 `json:"ResourceCreationTime,omitempty" xml:"ResourceCreationTime,omitempty"`
	// Indicates whether the resource was deleted. Valid values:
	//
	// 	- 1: The resource was not deleted.
	//
	// 	- 0: The resource was deleted.
	//
	// example:
	//
	// 1
	ResourceDeleted *int32 `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// new-bucket
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// new-bucket
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The status of the resource. The value of this parameter varies based on the resource type and may be empty.
	//
	// 	- If the ResourceType parameter is set to ACS::ECS::Instance, the resource is an ECS instance that has a specific state. In this case, the valid values of this parameter are Running and Stopped.
	//
	// 	- If the ResourceType parameter is ACS::OSS::Bucket, the resource is an Object Storage Service (OSS) bucket that is not in a specific state. In this case, this parameter is left empty.
	//
	// example:
	//
	// offline
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ACS::OSS::BucketACS::CDN::Domain
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	//
	// example:
	//
	// {\\"\\"hc\\"\\":[\\"\\"value2\\"\\"]}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-t4n7pokxxxxxxxxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// vpc-t4nhheyvay74fp7n0hxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GoString() string {
	return s.String()
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetConfiguration() *string {
	return s.Configuration
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetRegion() *string {
	return s.Region
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceCreationTime() *int64 {
	return s.ResourceCreationTime
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceDeleted() *int32 {
	return s.ResourceDeleted
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetTags() *string {
	return s.Tags
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetVpcId() *string {
	return s.VpcId
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetAccountId(v int64) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.AccountId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetAvailabilityZone(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.AvailabilityZone = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetComplianceType(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ComplianceType = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetConfiguration(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Configuration = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetRegion(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Region = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceCreationTime(v int64) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceCreationTime = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceDeleted(v int32) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceDeleted = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceId(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceName(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceName = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceStatus(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceStatus = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceType(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceType = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetTags(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Tags = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetVSwitchId(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.VSwitchId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetVpcId(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.VpcId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) Validate() error {
	return dara.Validate(s)
}
