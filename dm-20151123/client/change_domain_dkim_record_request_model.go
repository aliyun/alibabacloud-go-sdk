// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDomainDkimRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDkimRsaLength(v int32) *ChangeDomainDkimRecordRequest
	GetDkimRsaLength() *int32
	SetDomain(v string) *ChangeDomainDkimRecordRequest
	GetDomain() *string
	SetOwnerId(v int64) *ChangeDomainDkimRecordRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ChangeDomainDkimRecordRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ChangeDomainDkimRecordRequest
	GetResourceOwnerId() *int64
}

type ChangeDomainDkimRecordRequest struct {
	DkimRsaLength        *int32  `json:"DkimRsaLength,omitempty" xml:"DkimRsaLength,omitempty"`
	Domain               *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ChangeDomainDkimRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeDomainDkimRecordRequest) GoString() string {
	return s.String()
}

func (s *ChangeDomainDkimRecordRequest) GetDkimRsaLength() *int32 {
	return s.DkimRsaLength
}

func (s *ChangeDomainDkimRecordRequest) GetDomain() *string {
	return s.Domain
}

func (s *ChangeDomainDkimRecordRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChangeDomainDkimRecordRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ChangeDomainDkimRecordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ChangeDomainDkimRecordRequest) SetDkimRsaLength(v int32) *ChangeDomainDkimRecordRequest {
	s.DkimRsaLength = &v
	return s
}

func (s *ChangeDomainDkimRecordRequest) SetDomain(v string) *ChangeDomainDkimRecordRequest {
	s.Domain = &v
	return s
}

func (s *ChangeDomainDkimRecordRequest) SetOwnerId(v int64) *ChangeDomainDkimRecordRequest {
	s.OwnerId = &v
	return s
}

func (s *ChangeDomainDkimRecordRequest) SetResourceOwnerAccount(v string) *ChangeDomainDkimRecordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ChangeDomainDkimRecordRequest) SetResourceOwnerId(v int64) *ChangeDomainDkimRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ChangeDomainDkimRecordRequest) Validate() error {
	return dara.Validate(s)
}
