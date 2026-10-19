// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiModelKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdbInstanceName(v string) *CreateMultiModelKnowledgeBaseRequest
	GetAdbInstanceName() *string
	SetDBClusterId(v string) *CreateMultiModelKnowledgeBaseRequest
	GetDBClusterId() *string
	SetDbClusterAcu(v int32) *CreateMultiModelKnowledgeBaseRequest
	GetDbClusterAcu() *int32
	SetDescription(v string) *CreateMultiModelKnowledgeBaseRequest
	GetDescription() *string
	SetLakeStorageBucketName(v string) *CreateMultiModelKnowledgeBaseRequest
	GetLakeStorageBucketName() *string
	SetRegionId(v string) *CreateMultiModelKnowledgeBaseRequest
	GetRegionId() *string
	SetResourceAcuMax(v int32) *CreateMultiModelKnowledgeBaseRequest
	GetResourceAcuMax() *int32
	SetResourceAcuMin(v int32) *CreateMultiModelKnowledgeBaseRequest
	GetResourceAcuMin() *int32
	SetVSwitchId(v string) *CreateMultiModelKnowledgeBaseRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateMultiModelKnowledgeBaseRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateMultiModelKnowledgeBaseRequest
	GetZoneId() *string
}

type CreateMultiModelKnowledgeBaseRequest struct {
	// **[Deprecated]**
	//
	// example:
	//
	// None
	AdbInstanceName *string `json:"AdbInstanceName,omitempty" xml:"AdbInstanceName,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The number of ACUs for the AnalyticDB for MySQL instance resources.
	//
	// example:
	//
	// 24
	DbClusterAcu *int32 `json:"DbClusterAcu,omitempty" xml:"DbClusterAcu,omitempty"`
	// The description of the multimodal knowledge base instance.
	//
	// example:
	//
	// None
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// None
	LakeStorageBucketName *string `json:"LakeStorageBucketName,omitempty" xml:"LakeStorageBucketName,omitempty"`
	// The region ID.
	//
	// > You can call the DescribeRegions operation to query the region ID of a specified Data Lakehouse Edition cluster.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The maximum number of ACUs for the multimodal knowledge base resources.
	//
	// example:
	//
	// 128
	ResourceAcuMax *int32 `json:"ResourceAcuMax,omitempty" xml:"ResourceAcuMax,omitempty"`
	// The minimum number of ACUs for the multimodal knowledge base resources.
	//
	// example:
	//
	// 32
	ResourceAcuMin *int32 `json:"ResourceAcuMin,omitempty" xml:"ResourceAcuMin,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp1ak7shaq72ywiah****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-uf66oesmrqge1t2gs****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone to which the instance belongs.
	//
	// example:
	//
	// cn-shanghai-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateMultiModelKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiModelKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *CreateMultiModelKnowledgeBaseRequest) GetAdbInstanceName() *string {
	return s.AdbInstanceName
}

func (s *CreateMultiModelKnowledgeBaseRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateMultiModelKnowledgeBaseRequest) GetDbClusterAcu() *int32 {
	return s.DbClusterAcu
}

func (s *CreateMultiModelKnowledgeBaseRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMultiModelKnowledgeBaseRequest) GetLakeStorageBucketName() *string {
	return s.LakeStorageBucketName
}

func (s *CreateMultiModelKnowledgeBaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMultiModelKnowledgeBaseRequest) GetResourceAcuMax() *int32 {
	return s.ResourceAcuMax
}

func (s *CreateMultiModelKnowledgeBaseRequest) GetResourceAcuMin() *int32 {
	return s.ResourceAcuMin
}

func (s *CreateMultiModelKnowledgeBaseRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateMultiModelKnowledgeBaseRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateMultiModelKnowledgeBaseRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateMultiModelKnowledgeBaseRequest) SetAdbInstanceName(v string) *CreateMultiModelKnowledgeBaseRequest {
	s.AdbInstanceName = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseRequest) SetDBClusterId(v string) *CreateMultiModelKnowledgeBaseRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseRequest) SetDbClusterAcu(v int32) *CreateMultiModelKnowledgeBaseRequest {
	s.DbClusterAcu = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseRequest) SetDescription(v string) *CreateMultiModelKnowledgeBaseRequest {
	s.Description = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseRequest) SetLakeStorageBucketName(v string) *CreateMultiModelKnowledgeBaseRequest {
	s.LakeStorageBucketName = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseRequest) SetRegionId(v string) *CreateMultiModelKnowledgeBaseRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseRequest) SetResourceAcuMax(v int32) *CreateMultiModelKnowledgeBaseRequest {
	s.ResourceAcuMax = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseRequest) SetResourceAcuMin(v int32) *CreateMultiModelKnowledgeBaseRequest {
	s.ResourceAcuMin = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseRequest) SetVSwitchId(v string) *CreateMultiModelKnowledgeBaseRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseRequest) SetVpcId(v string) *CreateMultiModelKnowledgeBaseRequest {
	s.VpcId = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseRequest) SetZoneId(v string) *CreateMultiModelKnowledgeBaseRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
