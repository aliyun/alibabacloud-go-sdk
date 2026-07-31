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
	// The prefix of the public connection address.
	//
	// - It must begin with a lowercase letter and can contain only lowercase letters, digits, and hyphens (-).
	//
	// - It must be no more than 30 characters long.
	//
	// example:
	//
	// test12
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// <props="china">The cluster ID of an Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The cluster ID of a Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1z5d2q71is2****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database engine. Valid values:
	//
	// - **AnalyticDB*	- (default): the AnalyticDB for MySQL engine.
	//
	// - **Clickhouse**: the wide table engine.
	//
	// example:
	//
	// Clickhouse
	Engine            *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
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
