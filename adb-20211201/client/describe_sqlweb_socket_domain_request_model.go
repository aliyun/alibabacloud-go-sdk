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
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// example:
	//
	// amv-bp1lw6g669zpi660
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Module      *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The region ID
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
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
