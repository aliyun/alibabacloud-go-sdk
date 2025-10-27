// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogStoreKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogStoreName(v string) *DescribeLogStoreKeysRequest
	GetLogStoreName() *string
	SetOwnerAccount(v string) *DescribeLogStoreKeysRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLogStoreKeysRequest
	GetOwnerId() *int64
	SetProjectName(v string) *DescribeLogStoreKeysRequest
	GetProjectName() *string
	SetRegionId(v string) *DescribeLogStoreKeysRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeLogStoreKeysRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLogStoreKeysRequest
	GetResourceOwnerId() *int64
}

type DescribeLogStoreKeysRequest struct {
	// The name of the Logstore.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-hcl2
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the Simple Log Service project.
	//
	// This parameter is required.
	//
	// example:
	//
	// nbgame-point
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLogStoreKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogStoreKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreKeysRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeLogStoreKeysRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLogStoreKeysRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLogStoreKeysRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeLogStoreKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLogStoreKeysRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLogStoreKeysRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLogStoreKeysRequest) SetLogStoreName(v string) *DescribeLogStoreKeysRequest {
	s.LogStoreName = &v
	return s
}

func (s *DescribeLogStoreKeysRequest) SetOwnerAccount(v string) *DescribeLogStoreKeysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLogStoreKeysRequest) SetOwnerId(v int64) *DescribeLogStoreKeysRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLogStoreKeysRequest) SetProjectName(v string) *DescribeLogStoreKeysRequest {
	s.ProjectName = &v
	return s
}

func (s *DescribeLogStoreKeysRequest) SetRegionId(v string) *DescribeLogStoreKeysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogStoreKeysRequest) SetResourceOwnerAccount(v string) *DescribeLogStoreKeysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLogStoreKeysRequest) SetResourceOwnerId(v int64) *DescribeLogStoreKeysRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLogStoreKeysRequest) Validate() error {
	return dara.Validate(s)
}
