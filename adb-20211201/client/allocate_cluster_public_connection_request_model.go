// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateClusterPublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *AllocateClusterPublicConnectionRequest
	GetConnectionStringPrefix() *string
	SetDBClusterId(v string) *AllocateClusterPublicConnectionRequest
	GetDBClusterId() *string
	SetEngine(v string) *AllocateClusterPublicConnectionRequest
	GetEngine() *string
	SetResourceGroupName(v string) *AllocateClusterPublicConnectionRequest
	GetResourceGroupName() *string
}

type AllocateClusterPublicConnectionRequest struct {
	// Specifies the prefix of the public endpoint. The prefix must meet the following requirements:
	//
	// - It must start with a lowercase letter and can contain lowercase letters, digits, and hyphens (-).
	//
	// - It can be up to 30 characters in length.
	//
	// example:
	//
	// test12
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// <props="china">The ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1z5d2q71is2****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database engine. Valid values:
	//
	// - **AnalyticDB*	- (default): AnalyticDB for MySQL engine.
	//
	// - **Clickhouse**: wide table engine.
	//
	// example:
	//
	// Clickhouse
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// test_group
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s AllocateClusterPublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateClusterPublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateClusterPublicConnectionRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *AllocateClusterPublicConnectionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *AllocateClusterPublicConnectionRequest) GetEngine() *string {
	return s.Engine
}

func (s *AllocateClusterPublicConnectionRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *AllocateClusterPublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateClusterPublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetDBClusterId(v string) *AllocateClusterPublicConnectionRequest {
	s.DBClusterId = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetEngine(v string) *AllocateClusterPublicConnectionRequest {
	s.Engine = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetResourceGroupName(v string) *AllocateClusterPublicConnectionRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
