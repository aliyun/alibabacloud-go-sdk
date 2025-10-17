// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDbClusterAttributeZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDbClusterAttributeZonalRequest
	GetDBClusterId() *string
	SetDescribeType(v string) *DescribeDbClusterAttributeZonalRequest
	GetDescribeType() *string
	SetOwnerAccount(v string) *DescribeDbClusterAttributeZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDbClusterAttributeZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDbClusterAttributeZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDbClusterAttributeZonalRequest
	GetResourceOwnerId() *int64
}

type DescribeDbClusterAttributeZonalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// AI
	DescribeType         *string `json:"DescribeType,omitempty" xml:"DescribeType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDbClusterAttributeZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbClusterAttributeZonalRequest) GoString() string {
	return s.String()
}

func (s *DescribeDbClusterAttributeZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDbClusterAttributeZonalRequest) GetDescribeType() *string {
	return s.DescribeType
}

func (s *DescribeDbClusterAttributeZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDbClusterAttributeZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDbClusterAttributeZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDbClusterAttributeZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDbClusterAttributeZonalRequest) SetDBClusterId(v string) *DescribeDbClusterAttributeZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalRequest) SetDescribeType(v string) *DescribeDbClusterAttributeZonalRequest {
	s.DescribeType = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalRequest) SetOwnerAccount(v string) *DescribeDbClusterAttributeZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalRequest) SetOwnerId(v int64) *DescribeDbClusterAttributeZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalRequest) SetResourceOwnerAccount(v string) *DescribeDbClusterAttributeZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalRequest) SetResourceOwnerId(v int64) *DescribeDbClusterAttributeZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalRequest) Validate() error {
	return dara.Validate(s)
}
