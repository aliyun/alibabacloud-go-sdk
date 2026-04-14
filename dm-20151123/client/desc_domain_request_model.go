// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v int32) *DescDomainRequest
	GetDomainId() *int32
	SetOwnerId(v int64) *DescDomainRequest
	GetOwnerId() *int64
	SetRequireRealTimeDnsRecords(v bool) *DescDomainRequest
	GetRequireRealTimeDnsRecords() *bool
	SetResourceOwnerAccount(v string) *DescDomainRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescDomainRequest
	GetResourceOwnerId() *int64
}

type DescDomainRequest struct {
	// Domain ID. Can be obtained through QueryDomainByParam.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13464
	DomainId *int32 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	OwnerId  *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Determines whether to perform real-time DNS resolution
	//
	// example:
	//
	// true
	RequireRealTimeDnsRecords *bool   `json:"RequireRealTimeDnsRecords,omitempty" xml:"RequireRealTimeDnsRecords,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DescDomainRequest) GoString() string {
	return s.String()
}

func (s *DescDomainRequest) GetDomainId() *int32 {
	return s.DomainId
}

func (s *DescDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescDomainRequest) GetRequireRealTimeDnsRecords() *bool {
	return s.RequireRealTimeDnsRecords
}

func (s *DescDomainRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescDomainRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescDomainRequest) SetDomainId(v int32) *DescDomainRequest {
	s.DomainId = &v
	return s
}

func (s *DescDomainRequest) SetOwnerId(v int64) *DescDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DescDomainRequest) SetRequireRealTimeDnsRecords(v bool) *DescDomainRequest {
	s.RequireRealTimeDnsRecords = &v
	return s
}

func (s *DescDomainRequest) SetResourceOwnerAccount(v string) *DescDomainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescDomainRequest) SetResourceOwnerId(v int64) *DescDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescDomainRequest) Validate() error {
	return dara.Validate(s)
}
