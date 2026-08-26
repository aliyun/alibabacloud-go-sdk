// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBType(v string) *CreateKnowledgeSpaceRequest
	GetDBType() *string
	SetDescription(v string) *CreateKnowledgeSpaceRequest
	GetDescription() *string
	SetEmbeddingDimension(v int32) *CreateKnowledgeSpaceRequest
	GetEmbeddingDimension() *int32
	SetEmbeddingModel(v string) *CreateKnowledgeSpaceRequest
	GetEmbeddingModel() *string
	SetEnforceAcl(v bool) *CreateKnowledgeSpaceRequest
	GetEnforceAcl() *bool
	SetLLMModel(v string) *CreateKnowledgeSpaceRequest
	GetLLMModel() *string
	SetName(v string) *CreateKnowledgeSpaceRequest
	GetName() *string
	SetOSSAccessKey(v string) *CreateKnowledgeSpaceRequest
	GetOSSAccessKey() *string
	SetOSSBucket(v string) *CreateKnowledgeSpaceRequest
	GetOSSBucket() *string
	SetOSSSecretKey(v string) *CreateKnowledgeSpaceRequest
	GetOSSSecretKey() *string
	SetRegionId(v string) *CreateKnowledgeSpaceRequest
	GetRegionId() *string
	SetRerankModel(v string) *CreateKnowledgeSpaceRequest
	GetRerankModel() *string
	SetSecurityGroupId(v string) *CreateKnowledgeSpaceRequest
	GetSecurityGroupId() *string
	SetShardingSize(v int32) *CreateKnowledgeSpaceRequest
	GetShardingSize() *int32
	SetShardingStrategy(v string) *CreateKnowledgeSpaceRequest
	GetShardingStrategy() *string
	SetVSwitchId(v string) *CreateKnowledgeSpaceRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateKnowledgeSpaceRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateKnowledgeSpaceRequest
	GetZoneId() *string
}

type CreateKnowledgeSpaceRequest struct {
	// The database engine type.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The description of the knowledge space. The description can be up to 512 characters in length.
	//
	// example:
	//
	// testDesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The vector dimensions.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1536
	EmbeddingDimension *int32 `json:"EmbeddingDimension,omitempty" xml:"EmbeddingDimension,omitempty"`
	// The name of the embedding model.
	//
	// This parameter is required.
	//
	// example:
	//
	// text-embedding-v4
	EmbeddingModel *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
	// Specifies whether to enable ACL-based authentication for the knowledge space.
	//
	// example:
	//
	// false
	EnforceAcl *bool `json:"EnforceAcl,omitempty" xml:"EnforceAcl,omitempty"`
	// The name of the large language model.
	//
	// example:
	//
	// qwen3.6-plus
	LLMModel *string `json:"LLMModel,omitempty" xml:"LLMModel,omitempty"`
	// The name of the knowledge space. The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// OSS AK
	//
	// This parameter is required.
	//
	// example:
	//
	// ******
	OSSAccessKey *string `json:"OSSAccessKey,omitempty" xml:"OSSAccessKey,omitempty"`
	// The name of an existing OSS bucket in the same region.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-bucket
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	// OSS SK
	//
	// This parameter is required.
	//
	// example:
	//
	// ******
	OSSSecretKey *string `json:"OSSSecretKey,omitempty" xml:"OSSSecretKey,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the reranking model.
	//
	// example:
	//
	// qwen3-rerank
	RerankModel *string `json:"RerankModel,omitempty" xml:"RerankModel,omitempty"`
	// The security group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-********************
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The shard size, in tokens.
	//
	// This parameter is required.
	//
	// example:
	//
	// 512
	ShardingSize *int32 `json:"ShardingSize,omitempty" xml:"ShardingSize,omitempty"`
	// The sharding strategy. Valid values:
	//
	// - hierarchical (default)
	//
	// - hybrid
	//
	// This parameter is required.
	//
	// example:
	//
	// hierarchical
	ShardingStrategy *string `json:"ShardingStrategy,omitempty" xml:"ShardingStrategy,omitempty"`
	// The vSwitch for automatic creation of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-*********************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC for automatic creation of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-*************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The active zone for automatic creation of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateKnowledgeSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeSpaceRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeSpaceRequest) GetDBType() *string {
	return s.DBType
}

func (s *CreateKnowledgeSpaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeSpaceRequest) GetEmbeddingDimension() *int32 {
	return s.EmbeddingDimension
}

func (s *CreateKnowledgeSpaceRequest) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *CreateKnowledgeSpaceRequest) GetEnforceAcl() *bool {
	return s.EnforceAcl
}

func (s *CreateKnowledgeSpaceRequest) GetLLMModel() *string {
	return s.LLMModel
}

func (s *CreateKnowledgeSpaceRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeSpaceRequest) GetOSSAccessKey() *string {
	return s.OSSAccessKey
}

func (s *CreateKnowledgeSpaceRequest) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *CreateKnowledgeSpaceRequest) GetOSSSecretKey() *string {
	return s.OSSSecretKey
}

func (s *CreateKnowledgeSpaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateKnowledgeSpaceRequest) GetRerankModel() *string {
	return s.RerankModel
}

func (s *CreateKnowledgeSpaceRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateKnowledgeSpaceRequest) GetShardingSize() *int32 {
	return s.ShardingSize
}

func (s *CreateKnowledgeSpaceRequest) GetShardingStrategy() *string {
	return s.ShardingStrategy
}

func (s *CreateKnowledgeSpaceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateKnowledgeSpaceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateKnowledgeSpaceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateKnowledgeSpaceRequest) SetDBType(v string) *CreateKnowledgeSpaceRequest {
	s.DBType = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetDescription(v string) *CreateKnowledgeSpaceRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetEmbeddingDimension(v int32) *CreateKnowledgeSpaceRequest {
	s.EmbeddingDimension = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetEmbeddingModel(v string) *CreateKnowledgeSpaceRequest {
	s.EmbeddingModel = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetEnforceAcl(v bool) *CreateKnowledgeSpaceRequest {
	s.EnforceAcl = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetLLMModel(v string) *CreateKnowledgeSpaceRequest {
	s.LLMModel = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetName(v string) *CreateKnowledgeSpaceRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetOSSAccessKey(v string) *CreateKnowledgeSpaceRequest {
	s.OSSAccessKey = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetOSSBucket(v string) *CreateKnowledgeSpaceRequest {
	s.OSSBucket = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetOSSSecretKey(v string) *CreateKnowledgeSpaceRequest {
	s.OSSSecretKey = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetRegionId(v string) *CreateKnowledgeSpaceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetRerankModel(v string) *CreateKnowledgeSpaceRequest {
	s.RerankModel = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetSecurityGroupId(v string) *CreateKnowledgeSpaceRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetShardingSize(v int32) *CreateKnowledgeSpaceRequest {
	s.ShardingSize = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetShardingStrategy(v string) *CreateKnowledgeSpaceRequest {
	s.ShardingStrategy = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetVSwitchId(v string) *CreateKnowledgeSpaceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetVpcId(v string) *CreateKnowledgeSpaceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) SetZoneId(v string) *CreateKnowledgeSpaceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateKnowledgeSpaceRequest) Validate() error {
	return dara.Validate(s)
}
