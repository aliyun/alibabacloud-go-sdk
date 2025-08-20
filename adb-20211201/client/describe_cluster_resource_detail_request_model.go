// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterResourceDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeClusterResourceDetailRequest
	GetDBClusterId() *string
}

type DescribeClusterResourceDetailRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1jj9xqft1po****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeClusterResourceDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResourceDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceDetailRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeClusterResourceDetailRequest) SetDBClusterId(v string) *DescribeClusterResourceDetailRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeClusterResourceDetailRequest) Validate() error {
	return dara.Validate(s)
}
