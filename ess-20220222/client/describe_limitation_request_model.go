// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLimitationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeLimitationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeLimitationRequest
	GetResourceOwnerAccount() *string
}

type DescribeLimitationRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DescribeLimitationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLimitationRequest) GoString() string {
	return s.String()
}

func (s *DescribeLimitationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLimitationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLimitationRequest) SetOwnerId(v int64) *DescribeLimitationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLimitationRequest) SetResourceOwnerAccount(v string) *DescribeLimitationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLimitationRequest) Validate() error {
	return dara.Validate(s)
}
