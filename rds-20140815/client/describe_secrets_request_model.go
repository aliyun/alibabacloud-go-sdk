// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecretsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeSecretsRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *DescribeSecretsRequest
	GetClientToken() *string
	SetDbInstanceId(v string) *DescribeSecretsRequest
	GetDbInstanceId() *string
	SetEngine(v string) *DescribeSecretsRequest
	GetEngine() *string
	SetOwnerAccount(v string) *DescribeSecretsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSecretsRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeSecretsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSecretsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeSecretsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSecretsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeSecretsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSecretsRequest
	GetResourceOwnerId() *int64
}

type DescribeSecretsRequest struct {
	// The language of the text within the response. Valid values:
	//
	// 	- **zh-CN**: Chinese
	//
	// 	- **en-US**: English
	//
	// > The default value is **en-US**.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// example:
	//
	// rm-xjkljjxxxxx
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The database engine of the database.
	//
	// > Only MySQL is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	Engine       *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Valid values: any non-zero positive integer.
	//
	// > The default value is 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the DescribeDBInstanceAttribute operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmxypivk***
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSecretsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecretsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecretsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeSecretsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeSecretsRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DescribeSecretsRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeSecretsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSecretsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSecretsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSecretsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSecretsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecretsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSecretsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSecretsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSecretsRequest) SetAcceptLanguage(v string) *DescribeSecretsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeSecretsRequest) SetClientToken(v string) *DescribeSecretsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSecretsRequest) SetDbInstanceId(v string) *DescribeSecretsRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeSecretsRequest) SetEngine(v string) *DescribeSecretsRequest {
	s.Engine = &v
	return s
}

func (s *DescribeSecretsRequest) SetOwnerAccount(v string) *DescribeSecretsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSecretsRequest) SetOwnerId(v int64) *DescribeSecretsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSecretsRequest) SetPageNumber(v int64) *DescribeSecretsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSecretsRequest) SetPageSize(v int64) *DescribeSecretsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSecretsRequest) SetRegionId(v string) *DescribeSecretsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecretsRequest) SetResourceGroupId(v string) *DescribeSecretsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSecretsRequest) SetResourceOwnerAccount(v string) *DescribeSecretsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSecretsRequest) SetResourceOwnerId(v int64) *DescribeSecretsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSecretsRequest) Validate() error {
	return dara.Validate(s)
}
