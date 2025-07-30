// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterUsedUtilizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeClusterUsedUtilizationRequest
	GetAccountId() *string
	SetClientToken(v string) *DescribeClusterUsedUtilizationRequest
	GetClientToken() *string
	SetDedicatedClusterId(v string) *DescribeClusterUsedUtilizationRequest
	GetDedicatedClusterId() *string
	SetDtsJobId(v string) *DescribeClusterUsedUtilizationRequest
	GetDtsJobId() *string
	SetEnv(v string) *DescribeClusterUsedUtilizationRequest
	GetEnv() *string
	SetMetricType(v string) *DescribeClusterUsedUtilizationRequest
	GetMetricType() *string
	SetOwnerID(v string) *DescribeClusterUsedUtilizationRequest
	GetOwnerID() *string
	SetRegionId(v string) *DescribeClusterUsedUtilizationRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeClusterUsedUtilizationRequest
	GetResourceGroupId() *string
	SetSecurityToken(v string) *DescribeClusterUsedUtilizationRequest
	GetSecurityToken() *string
}

type DescribeClusterUsedUtilizationRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter is discontinued.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. **The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the cluster. If the **MetricType*	- parameter is set to **CLUSTER**, enter the ID of the exclusive cluster. Otherwise, set this parameter to a node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// nodeid
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// The ID of the data migration or synchronization task.
	//
	// example:
	//
	// k2gm967v16f****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The cluster environment. Default value: **ALIYUN**.
	//
	// example:
	//
	// ALIYUN
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// Specifies whether to query the metrics of the cluster or a node. Default value: CLUSTER. Valid values:
	//
	// 	- **CLUSTER**: query the metrics of the cluster.
	//
	// 	- **NODE**: query the metrics of a node.
	//
	// example:
	//
	// NODE
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	OwnerID    *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// The ID of the region in which the Data Transmission Service (DTS) instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeClusterUsedUtilizationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterUsedUtilizationRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterUsedUtilizationRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeClusterUsedUtilizationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeClusterUsedUtilizationRequest) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *DescribeClusterUsedUtilizationRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeClusterUsedUtilizationRequest) GetEnv() *string {
	return s.Env
}

func (s *DescribeClusterUsedUtilizationRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeClusterUsedUtilizationRequest) GetOwnerID() *string {
	return s.OwnerID
}

func (s *DescribeClusterUsedUtilizationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterUsedUtilizationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeClusterUsedUtilizationRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeClusterUsedUtilizationRequest) SetAccountId(v string) *DescribeClusterUsedUtilizationRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeClusterUsedUtilizationRequest) SetClientToken(v string) *DescribeClusterUsedUtilizationRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeClusterUsedUtilizationRequest) SetDedicatedClusterId(v string) *DescribeClusterUsedUtilizationRequest {
	s.DedicatedClusterId = &v
	return s
}

func (s *DescribeClusterUsedUtilizationRequest) SetDtsJobId(v string) *DescribeClusterUsedUtilizationRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeClusterUsedUtilizationRequest) SetEnv(v string) *DescribeClusterUsedUtilizationRequest {
	s.Env = &v
	return s
}

func (s *DescribeClusterUsedUtilizationRequest) SetMetricType(v string) *DescribeClusterUsedUtilizationRequest {
	s.MetricType = &v
	return s
}

func (s *DescribeClusterUsedUtilizationRequest) SetOwnerID(v string) *DescribeClusterUsedUtilizationRequest {
	s.OwnerID = &v
	return s
}

func (s *DescribeClusterUsedUtilizationRequest) SetRegionId(v string) *DescribeClusterUsedUtilizationRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterUsedUtilizationRequest) SetResourceGroupId(v string) *DescribeClusterUsedUtilizationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClusterUsedUtilizationRequest) SetSecurityToken(v string) *DescribeClusterUsedUtilizationRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeClusterUsedUtilizationRequest) Validate() error {
	return dara.Validate(s)
}
