// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpfilterListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetIpfilterListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetIpfilterListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetIpfilterListRequest
	GetResourceOwnerId() *int64
}

type GetIpfilterListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetIpfilterListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIpfilterListRequest) GoString() string {
	return s.String()
}

func (s *GetIpfilterListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetIpfilterListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetIpfilterListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetIpfilterListRequest) SetOwnerId(v int64) *GetIpfilterListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetIpfilterListRequest) SetResourceOwnerAccount(v string) *GetIpfilterListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetIpfilterListRequest) SetResourceOwnerId(v int64) *GetIpfilterListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetIpfilterListRequest) Validate() error {
	return dara.Validate(s)
}
