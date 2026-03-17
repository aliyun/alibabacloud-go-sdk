// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDpiSignaturesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDpiGroupId(v string) *ListDpiSignaturesRequest
	GetDpiGroupId() *string
	SetDpiSignatureIds(v []*string) *ListDpiSignaturesRequest
	GetDpiSignatureIds() []*string
	SetDpiSignatureNames(v []*string) *ListDpiSignaturesRequest
	GetDpiSignatureNames() []*string
	SetMaxResults(v int32) *ListDpiSignaturesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDpiSignaturesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListDpiSignaturesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListDpiSignaturesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListDpiSignaturesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListDpiSignaturesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListDpiSignaturesRequest
	GetResourceOwnerId() *int64
}

type ListDpiSignaturesRequest struct {
	// The ID of the application group.
	//
	// example:
	//
	// 20
	DpiGroupId *string `json:"DpiGroupId,omitempty" xml:"DpiGroupId,omitempty"`
	// example:
	//
	// 235
	DpiSignatureIds []*string `json:"DpiSignatureIds,omitempty" xml:"DpiSignatureIds,omitempty" type:"Repeated"`
	// example:
	//
	// EdgeCast
	DpiSignatureNames []*string `json:"DpiSignatureNames,omitempty" xml:"DpiSignatureNames,omitempty" type:"Repeated"`
	// The number of entries to return on each page.
	//
	// Valid values: **1*	- to **100**.
	//
	// Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to query the next page.
	//
	// example:
	//
	// caeba0bbb2****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the application or application group belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListDpiSignaturesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDpiSignaturesRequest) GoString() string {
	return s.String()
}

func (s *ListDpiSignaturesRequest) GetDpiGroupId() *string {
	return s.DpiGroupId
}

func (s *ListDpiSignaturesRequest) GetDpiSignatureIds() []*string {
	return s.DpiSignatureIds
}

func (s *ListDpiSignaturesRequest) GetDpiSignatureNames() []*string {
	return s.DpiSignatureNames
}

func (s *ListDpiSignaturesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDpiSignaturesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDpiSignaturesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListDpiSignaturesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListDpiSignaturesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDpiSignaturesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListDpiSignaturesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListDpiSignaturesRequest) SetDpiGroupId(v string) *ListDpiSignaturesRequest {
	s.DpiGroupId = &v
	return s
}

func (s *ListDpiSignaturesRequest) SetDpiSignatureIds(v []*string) *ListDpiSignaturesRequest {
	s.DpiSignatureIds = v
	return s
}

func (s *ListDpiSignaturesRequest) SetDpiSignatureNames(v []*string) *ListDpiSignaturesRequest {
	s.DpiSignatureNames = v
	return s
}

func (s *ListDpiSignaturesRequest) SetMaxResults(v int32) *ListDpiSignaturesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDpiSignaturesRequest) SetNextToken(v string) *ListDpiSignaturesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDpiSignaturesRequest) SetOwnerAccount(v string) *ListDpiSignaturesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListDpiSignaturesRequest) SetOwnerId(v int64) *ListDpiSignaturesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDpiSignaturesRequest) SetRegionId(v string) *ListDpiSignaturesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDpiSignaturesRequest) SetResourceOwnerAccount(v string) *ListDpiSignaturesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListDpiSignaturesRequest) SetResourceOwnerId(v int64) *ListDpiSignaturesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListDpiSignaturesRequest) Validate() error {
	return dara.Validate(s)
}
