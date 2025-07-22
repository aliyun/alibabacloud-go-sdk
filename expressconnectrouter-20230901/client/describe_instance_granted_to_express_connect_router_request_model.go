// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceGrantedToExpressConnectRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallerType(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest
	GetCallerType() *string
	SetClientToken(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest
	GetClientToken() *string
	SetDryRun(v bool) *DescribeInstanceGrantedToExpressConnectRouterRequest
	GetDryRun() *bool
	SetEcrId(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest
	GetEcrId() *string
	SetInstanceId(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest
	GetInstanceId() *string
	SetInstanceOwnerId(v int64) *DescribeInstanceGrantedToExpressConnectRouterRequest
	GetInstanceOwnerId() *int64
	SetInstanceRegionId(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest
	GetInstanceRegionId() *string
	SetInstanceType(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest
	GetInstanceType() *string
	SetMaxResults(v int32) *DescribeInstanceGrantedToExpressConnectRouterRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest
	GetResourceGroupId() *string
	SetTagModels(v []*DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) *DescribeInstanceGrantedToExpressConnectRouterRequest
	GetTagModels() []*DescribeInstanceGrantedToExpressConnectRouterRequestTagModels
}

type DescribeInstanceGrantedToExpressConnectRouterRequest struct {
	// The type of the user account. Valid values:
	//
	// 	- **sub**: a Resource Access Management (RAM) user.
	//
	// 	- **parent**: an Alibaba Cloud account.
	//
	// example:
	//
	// OTHER
	CallerType *string `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the network instance.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the network instance.
	//
	// example:
	//
	// 129845258050****
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
	// The region ID of the network instance.
	//
	// example:
	//
	// cn-hangzhou
	InstanceRegionId *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VBR**
	//
	// 	- **VPC**
	//
	// example:
	//
	// VBR
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum number of entries to read. Valid values: 1 to 2147483647. Default value: 20.
	//
	// example:
	//
	// 6
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// gAAAAABkyGzFbZR2NnxnyVk0EiL7F3qMBtBim8Es0mugRT3qb8yEHAMaHGanzuaHUmiEq9QRmok0RgxJAINBOJZa5KPjopEu_Q==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group to which the network instance belongs.
	//
	// example:
	//
	// rg-aek2tsvbnfe****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	TagModels []*DescribeInstanceGrantedToExpressConnectRouterRequestTagModels `json:"TagModels,omitempty" xml:"TagModels,omitempty" type:"Repeated"`
}

func (s DescribeInstanceGrantedToExpressConnectRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGrantedToExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) GetCallerType() *string {
	return s.CallerType
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) GetInstanceOwnerId() *int64 {
	return s.InstanceOwnerId
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) GetInstanceRegionId() *string {
	return s.InstanceRegionId
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) GetTagModels() []*DescribeInstanceGrantedToExpressConnectRouterRequestTagModels {
	return s.TagModels
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetCallerType(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.CallerType = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetClientToken(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetDryRun(v bool) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetEcrId(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetInstanceId(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetInstanceOwnerId(v int64) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.InstanceOwnerId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetInstanceRegionId(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.InstanceRegionId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetInstanceType(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetMaxResults(v int32) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetNextToken(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetResourceGroupId(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetTagModels(v []*DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.TagModels = v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceGrantedToExpressConnectRouterRequestTagModels struct {
	// The tag key. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value. You can specify up to 20 tag values. The tag value cannot be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) SetTagKey(v string) *DescribeInstanceGrantedToExpressConnectRouterRequestTagModels {
	s.TagKey = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) SetTagValue(v string) *DescribeInstanceGrantedToExpressConnectRouterRequestTagModels {
	s.TagValue = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) Validate() error {
	return dara.Validate(s)
}
