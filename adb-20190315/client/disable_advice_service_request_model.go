// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAdviceServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DisableAdviceServiceRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DisableAdviceServiceRequest
	GetRegionId() *string
}

type DisableAdviceServiceRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf67culrr26q2****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableAdviceServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableAdviceServiceRequest) GoString() string {
	return s.String()
}

func (s *DisableAdviceServiceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DisableAdviceServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableAdviceServiceRequest) SetDBClusterId(v string) *DisableAdviceServiceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DisableAdviceServiceRequest) SetRegionId(v string) *DisableAdviceServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DisableAdviceServiceRequest) Validate() error {
	return dara.Validate(s)
}
