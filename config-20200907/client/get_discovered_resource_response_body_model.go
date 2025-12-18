// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiscoveredResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiscoveredResourceDetail(v *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) *GetDiscoveredResourceResponseBody
	GetDiscoveredResourceDetail() *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail
	SetRequestId(v string) *GetDiscoveredResourceResponseBody
	GetRequestId() *string
}

type GetDiscoveredResourceResponseBody struct {
	// The details of the resource.
	DiscoveredResourceDetail *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail `json:"DiscoveredResourceDetail,omitempty" xml:"DiscoveredResourceDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E4D71ACE-6B0A-46E0-8352-56952378CC7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDiscoveredResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDiscoveredResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceResponseBody) GetDiscoveredResourceDetail() *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	return s.DiscoveredResourceDetail
}

func (s *GetDiscoveredResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDiscoveredResourceResponseBody) SetDiscoveredResourceDetail(v *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) *GetDiscoveredResourceResponseBody {
	s.DiscoveredResourceDetail = v
	return s
}

func (s *GetDiscoveredResourceResponseBody) SetRequestId(v string) *GetDiscoveredResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDiscoveredResourceResponseBody) Validate() error {
	if s.DiscoveredResourceDetail != nil {
		if err := s.DiscoveredResourceDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDiscoveredResourceResponseBodyDiscoveredResourceDetail struct {
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
	// 	- NOT_APPLICABLE: The rule does not apply to the resources.
	//
	// 	- INSUFFICIENT_DATA: No data is available.
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
	// if can be null:
	// true
	//
	// example:
	//
	// vsw-t4n7pokxxxxxxxxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// vpc-t4nhheyvay74fp7n0hxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) String() string {
	return dara.Prettify(s)
}

func (s GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetConfiguration() *string {
	return s.Configuration
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetRegion() *string {
	return s.Region
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceCreationTime() *int64 {
	return s.ResourceCreationTime
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceDeleted() *int32 {
	return s.ResourceDeleted
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetTags() *string {
	return s.Tags
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GetVpcId() *string {
	return s.VpcId
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetAccountId(v int64) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.AccountId = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetAvailabilityZone(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.AvailabilityZone = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetComplianceType(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ComplianceType = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetConfiguration(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Configuration = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetRegion(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Region = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceCreationTime(v int64) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceCreationTime = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceDeleted(v int32) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceDeleted = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceId(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceId = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceName(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceName = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceStatus(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceStatus = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceType(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceType = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetTags(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Tags = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetVSwitchId(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.VSwitchId = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetVpcId(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.VpcId = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) Validate() error {
	return dara.Validate(s)
}
