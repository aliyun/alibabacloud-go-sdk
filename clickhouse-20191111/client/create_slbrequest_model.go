// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSLBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateSLBRequest
	GetDBClusterId() *string
	SetProduct(v string) *CreateSLBRequest
	GetProduct() *string
	SetResourceOwnerId(v int64) *CreateSLBRequest
	GetResourceOwnerId() *int64
}

type CreateSLBRequest struct {
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to obtain the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-uf6bnitmve5n0****
	DBClusterId     *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Product         *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateSLBRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSLBRequest) GoString() string {
	return s.String()
}

func (s *CreateSLBRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateSLBRequest) GetProduct() *string {
	return s.Product
}

func (s *CreateSLBRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSLBRequest) SetDBClusterId(v string) *CreateSLBRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateSLBRequest) SetProduct(v string) *CreateSLBRequest {
	s.Product = &v
	return s
}

func (s *CreateSLBRequest) SetResourceOwnerId(v int64) *CreateSLBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSLBRequest) Validate() error {
	return dara.Validate(s)
}
