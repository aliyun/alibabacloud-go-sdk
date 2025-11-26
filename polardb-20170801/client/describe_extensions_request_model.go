// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExtensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeExtensionsRequest
	GetDBClusterId() *string
	SetDBName(v string) *DescribeExtensionsRequest
	GetDBName() *string
	SetExtensionName(v string) *DescribeExtensionsRequest
	GetExtensionName() *string
	SetOwnerAccount(v string) *DescribeExtensionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeExtensionsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeExtensionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeExtensionsRequest
	GetResourceOwnerId() *int64
}

type DescribeExtensionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// song
	DBName        *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ExtensionName *string `json:"ExtensionName,omitempty" xml:"ExtensionName,omitempty"`
	// example:
	//
	// test@example.com
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// example:
	//
	// 1234567890123456
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// test@example.com
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// example:
	//
	// 1234567890123456
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeExtensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExtensionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeExtensionsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeExtensionsRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeExtensionsRequest) GetExtensionName() *string {
	return s.ExtensionName
}

func (s *DescribeExtensionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeExtensionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeExtensionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeExtensionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeExtensionsRequest) SetDBClusterId(v string) *DescribeExtensionsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeExtensionsRequest) SetDBName(v string) *DescribeExtensionsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeExtensionsRequest) SetExtensionName(v string) *DescribeExtensionsRequest {
	s.ExtensionName = &v
	return s
}

func (s *DescribeExtensionsRequest) SetOwnerAccount(v string) *DescribeExtensionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeExtensionsRequest) SetOwnerId(v int64) *DescribeExtensionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeExtensionsRequest) SetResourceOwnerAccount(v string) *DescribeExtensionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeExtensionsRequest) SetResourceOwnerId(v int64) *DescribeExtensionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeExtensionsRequest) Validate() error {
	return dara.Validate(s)
}
