// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLWebSocketDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSQLWebSocketDomainRequest
	GetDBClusterId() *string
	SetModule(v string) *DescribeSQLWebSocketDomainRequest
	GetModule() *string
	SetRegionId(v string) *DescribeSQLWebSocketDomainRequest
	GetRegionId() *string
}

type DescribeSQLWebSocketDomainRequest struct {
	// The cluster ID.
	//
	// > Call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) API to query the details of all clusters in your account, including cluster IDs.
	//
	// example:
	//
	// amv-bp1lw6g669zpi660
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The application module name.
	//
	// - `SQLWebSocket`: The module for SQL development.
	//
	// - `Assistant`: The module for the intelligent assistant.
	//
	// example:
	//
	// Assistant
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The region ID.
	//
	// > Call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) API to query the region IDs supported by AnalyticDB for MySQL.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSQLWebSocketDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLWebSocketDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLWebSocketDomainRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSQLWebSocketDomainRequest) GetModule() *string {
	return s.Module
}

func (s *DescribeSQLWebSocketDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSQLWebSocketDomainRequest) SetDBClusterId(v string) *DescribeSQLWebSocketDomainRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLWebSocketDomainRequest) SetModule(v string) *DescribeSQLWebSocketDomainRequest {
	s.Module = &v
	return s
}

func (s *DescribeSQLWebSocketDomainRequest) SetRegionId(v string) *DescribeSQLWebSocketDomainRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSQLWebSocketDomainRequest) Validate() error {
	return dara.Validate(s)
}
