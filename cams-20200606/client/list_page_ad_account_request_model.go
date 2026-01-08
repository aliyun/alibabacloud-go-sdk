// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPageAdAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *ListPageAdAccountRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *ListPageAdAccountRequest
	GetOwnerId() *int64
	SetPageId(v string) *ListPageAdAccountRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *ListPageAdAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListPageAdAccountRequest
	GetResourceOwnerId() *int64
}

type ListPageAdAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cams-**
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 230***
	PageId               *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListPageAdAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPageAdAccountRequest) GoString() string {
	return s.String()
}

func (s *ListPageAdAccountRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListPageAdAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListPageAdAccountRequest) GetPageId() *string {
	return s.PageId
}

func (s *ListPageAdAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListPageAdAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListPageAdAccountRequest) SetCustSpaceId(v string) *ListPageAdAccountRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListPageAdAccountRequest) SetOwnerId(v int64) *ListPageAdAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *ListPageAdAccountRequest) SetPageId(v string) *ListPageAdAccountRequest {
	s.PageId = &v
	return s
}

func (s *ListPageAdAccountRequest) SetResourceOwnerAccount(v string) *ListPageAdAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListPageAdAccountRequest) SetResourceOwnerId(v int64) *ListPageAdAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListPageAdAccountRequest) Validate() error {
	return dara.Validate(s)
}
