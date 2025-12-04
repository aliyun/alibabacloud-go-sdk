// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSLBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteSLBRequest
	GetDBClusterId() *string
	SetProduct(v string) *DeleteSLBRequest
	GetProduct() *string
}

type DeleteSLBRequest struct {
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to view cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-uf6bnitmve5n0****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Product     *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s DeleteSLBRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSLBRequest) GoString() string {
	return s.String()
}

func (s *DeleteSLBRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteSLBRequest) GetProduct() *string {
	return s.Product
}

func (s *DeleteSLBRequest) SetDBClusterId(v string) *DeleteSLBRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteSLBRequest) SetProduct(v string) *DeleteSLBRequest {
	s.Product = &v
	return s
}

func (s *DeleteSLBRequest) Validate() error {
	return dara.Validate(s)
}
