// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSchemasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSchemasRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeSchemasRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSchemasRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeSchemasRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSchemasRequest
	GetResourceOwnerId() *int64
}

type DescribeSchemasRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSchemasRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSchemasRequest) GoString() string {
	return s.String()
}

func (s *DescribeSchemasRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSchemasRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSchemasRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSchemasRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSchemasRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSchemasRequest) SetDBClusterId(v string) *DescribeSchemasRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSchemasRequest) SetOwnerAccount(v string) *DescribeSchemasRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSchemasRequest) SetOwnerId(v int64) *DescribeSchemasRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSchemasRequest) SetResourceOwnerAccount(v string) *DescribeSchemasRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSchemasRequest) SetResourceOwnerId(v int64) *DescribeSchemasRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSchemasRequest) Validate() error {
	return dara.Validate(s)
}
