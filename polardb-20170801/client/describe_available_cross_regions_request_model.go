// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableCrossRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAvailableCrossRegionsRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeAvailableCrossRegionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAvailableCrossRegionsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeAvailableCrossRegionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAvailableCrossRegionsRequest
	GetResourceOwnerId() *int64
}

type DescribeAvailableCrossRegionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-xxxxxxxxxxx
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAvailableCrossRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableCrossRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCrossRegionsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAvailableCrossRegionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAvailableCrossRegionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAvailableCrossRegionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAvailableCrossRegionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAvailableCrossRegionsRequest) SetDBClusterId(v string) *DescribeAvailableCrossRegionsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAvailableCrossRegionsRequest) SetOwnerAccount(v string) *DescribeAvailableCrossRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAvailableCrossRegionsRequest) SetOwnerId(v int64) *DescribeAvailableCrossRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableCrossRegionsRequest) SetResourceOwnerAccount(v string) *DescribeAvailableCrossRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableCrossRegionsRequest) SetResourceOwnerId(v int64) *DescribeAvailableCrossRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableCrossRegionsRequest) Validate() error {
	return dara.Validate(s)
}
