// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseClusterPublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ReleaseClusterPublicConnectionRequest
	GetDBClusterId() *string
	SetEngine(v string) *ReleaseClusterPublicConnectionRequest
	GetEngine() *string
	SetResourceGroupName(v string) *ReleaseClusterPublicConnectionRequest
	GetResourceGroupName() *string
}

type ReleaseClusterPublicConnectionRequest struct {
	// <props="china">The ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database engine. Valid values:
	//
	// - **AnalyticDB*	- (Default): AnalyticDB for MySQL.
	//
	// - **ClickHouse**: a wide table engine.
	//
	// example:
	//
	// Clickhouse
	Engine            *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s ReleaseClusterPublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseClusterPublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseClusterPublicConnectionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ReleaseClusterPublicConnectionRequest) GetEngine() *string {
	return s.Engine
}

func (s *ReleaseClusterPublicConnectionRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ReleaseClusterPublicConnectionRequest) SetDBClusterId(v string) *ReleaseClusterPublicConnectionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) SetEngine(v string) *ReleaseClusterPublicConnectionRequest {
	s.Engine = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) SetResourceGroupName(v string) *ReleaseClusterPublicConnectionRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
