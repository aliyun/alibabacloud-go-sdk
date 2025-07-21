// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v int32) *CheckDomainRequest
	GetDomainId() *int32
	SetOwnerId(v int64) *CheckDomainRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CheckDomainRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckDomainRequest
	GetResourceOwnerId() *int64
}

type CheckDomainRequest struct {
	// Domain ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 153345
	DomainId             *int32  `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDomainRequest) GoString() string {
	return s.String()
}

func (s *CheckDomainRequest) GetDomainId() *int32 {
	return s.DomainId
}

func (s *CheckDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckDomainRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckDomainRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckDomainRequest) SetDomainId(v int32) *CheckDomainRequest {
	s.DomainId = &v
	return s
}

func (s *CheckDomainRequest) SetOwnerId(v int64) *CheckDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckDomainRequest) SetResourceOwnerAccount(v string) *CheckDomainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckDomainRequest) SetResourceOwnerId(v int64) *CheckDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckDomainRequest) Validate() error {
	return dara.Validate(s)
}
