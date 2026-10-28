// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakeCacheSizeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeLakeCacheSizeRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeLakeCacheSizeRequest
	GetRegionId() *string
}

type DescribeLakeCacheSizeRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the cluster ID of a Data Warehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp10b6646l07akdt
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the instance. The value must be the same as the region of the instance specified by DBClusterId.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLakeCacheSizeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakeCacheSizeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLakeCacheSizeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeLakeCacheSizeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLakeCacheSizeRequest) SetDBClusterId(v string) *DescribeLakeCacheSizeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLakeCacheSizeRequest) SetRegionId(v string) *DescribeLakeCacheSizeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLakeCacheSizeRequest) Validate() error {
	return dara.Validate(s)
}
