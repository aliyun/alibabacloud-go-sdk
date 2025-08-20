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
}

type DescribeLakeCacheSizeRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp10b6646l07akdt
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
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

func (s *DescribeLakeCacheSizeRequest) SetDBClusterId(v string) *DescribeLakeCacheSizeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLakeCacheSizeRequest) Validate() error {
	return dara.Validate(s)
}
