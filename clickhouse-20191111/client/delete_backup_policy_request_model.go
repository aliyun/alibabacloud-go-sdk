// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteBackupPolicyRequest
	GetClientToken() *string
	SetDBClusterId(v string) *DeleteBackupPolicyRequest
	GetDBClusterId() *string
	SetMaxResults(v int32) *DeleteBackupPolicyRequest
	GetMaxResults() *int32
	SetProduct(v string) *DeleteBackupPolicyRequest
	GetProduct() *string
	SetResourceOwnerId(v int64) *DeleteBackupPolicyRequest
	GetResourceOwnerId() *int64
}

type DeleteBackupPolicyRequest struct {
	// A client token used to ensure the idempotence of the request. The value must be a string that contains a maximum of 64 ASCII characters and cannot contain non-ASCII characters.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all clusters in the destination region, including cluster IDs.
	//
	// example:
	//
	// cc-bp1t9lbb7a4z7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The number of records to return on each page. Valid values: 1 to **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The product name.
	//
	// example:
	//
	// clickhouse
	Product         *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteBackupPolicyRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteBackupPolicyRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DeleteBackupPolicyRequest) GetProduct() *string {
	return s.Product
}

func (s *DeleteBackupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteBackupPolicyRequest) SetClientToken(v string) *DeleteBackupPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteBackupPolicyRequest) SetDBClusterId(v string) *DeleteBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteBackupPolicyRequest) SetMaxResults(v int32) *DeleteBackupPolicyRequest {
	s.MaxResults = &v
	return s
}

func (s *DeleteBackupPolicyRequest) SetProduct(v string) *DeleteBackupPolicyRequest {
	s.Product = &v
	return s
}

func (s *DeleteBackupPolicyRequest) SetResourceOwnerId(v int64) *DeleteBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
