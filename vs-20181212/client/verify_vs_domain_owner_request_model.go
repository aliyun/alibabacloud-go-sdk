// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyVsDomainOwnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *VerifyVsDomainOwnerRequest
	GetDomainName() *string
	SetOwnerId(v int64) *VerifyVsDomainOwnerRequest
	GetOwnerId() *int64
	SetVerifyType(v string) *VerifyVsDomainOwnerRequest
	GetVerifyType() *string
}

type VerifyVsDomainOwnerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dnsCheck
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s VerifyVsDomainOwnerRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyVsDomainOwnerRequest) GoString() string {
	return s.String()
}

func (s *VerifyVsDomainOwnerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *VerifyVsDomainOwnerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *VerifyVsDomainOwnerRequest) GetVerifyType() *string {
	return s.VerifyType
}

func (s *VerifyVsDomainOwnerRequest) SetDomainName(v string) *VerifyVsDomainOwnerRequest {
	s.DomainName = &v
	return s
}

func (s *VerifyVsDomainOwnerRequest) SetOwnerId(v int64) *VerifyVsDomainOwnerRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyVsDomainOwnerRequest) SetVerifyType(v string) *VerifyVsDomainOwnerRequest {
	s.VerifyType = &v
	return s
}

func (s *VerifyVsDomainOwnerRequest) Validate() error {
	return dara.Validate(s)
}
