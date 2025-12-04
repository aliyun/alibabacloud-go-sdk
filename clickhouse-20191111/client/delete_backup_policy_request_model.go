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
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// cc-bp1t9lbb7a4z7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
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
