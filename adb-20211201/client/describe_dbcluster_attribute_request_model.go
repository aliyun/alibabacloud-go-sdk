// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterAttributeRequest
	GetDBClusterId() *string
}

type DescribeDBClusterAttributeRequest struct {
	// <props="china">The ID of an Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to obtain the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-wz9509beptiz****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeDBClusterAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterAttributeRequest) SetDBClusterId(v string) *DescribeDBClusterAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) Validate() error {
	return dara.Validate(s)
}
