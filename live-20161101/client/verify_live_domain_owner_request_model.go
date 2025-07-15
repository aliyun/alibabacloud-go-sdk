// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyLiveDomainOwnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *VerifyLiveDomainOwnerRequest
	GetDomainName() *string
	SetOwnerId(v int64) *VerifyLiveDomainOwnerRequest
	GetOwnerId() *int64
	SetRegionId(v string) *VerifyLiveDomainOwnerRequest
	GetRegionId() *string
	SetVerifyType(v string) *VerifyLiveDomainOwnerRequest
	GetVerifyType() *string
}

type VerifyLiveDomainOwnerRequest struct {
	// The domain name for which you want to verify the ownership. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The verification method. Valid values:
	//
	// 	- dnsCheck: DNS record verification
	//
	// 	- fileCheck: file verification
	//
	// This parameter is required.
	//
	// example:
	//
	// dnsCheck
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s VerifyLiveDomainOwnerRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyLiveDomainOwnerRequest) GoString() string {
	return s.String()
}

func (s *VerifyLiveDomainOwnerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *VerifyLiveDomainOwnerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *VerifyLiveDomainOwnerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *VerifyLiveDomainOwnerRequest) GetVerifyType() *string {
	return s.VerifyType
}

func (s *VerifyLiveDomainOwnerRequest) SetDomainName(v string) *VerifyLiveDomainOwnerRequest {
	s.DomainName = &v
	return s
}

func (s *VerifyLiveDomainOwnerRequest) SetOwnerId(v int64) *VerifyLiveDomainOwnerRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyLiveDomainOwnerRequest) SetRegionId(v string) *VerifyLiveDomainOwnerRequest {
	s.RegionId = &v
	return s
}

func (s *VerifyLiveDomainOwnerRequest) SetVerifyType(v string) *VerifyLiveDomainOwnerRequest {
	s.VerifyType = &v
	return s
}

func (s *VerifyLiveDomainOwnerRequest) Validate() error {
	return dara.Validate(s)
}
