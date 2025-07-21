// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescAccountSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescAccountSummaryRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescAccountSummaryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescAccountSummaryRequest
	GetResourceOwnerId() *int64
}

type DescAccountSummaryRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescAccountSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescAccountSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescAccountSummaryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescAccountSummaryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescAccountSummaryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescAccountSummaryRequest) SetOwnerId(v int64) *DescAccountSummaryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescAccountSummaryRequest) SetResourceOwnerAccount(v string) *DescAccountSummaryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescAccountSummaryRequest) SetResourceOwnerId(v int64) *DescAccountSummaryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescAccountSummaryRequest) Validate() error {
	return dara.Validate(s)
}
