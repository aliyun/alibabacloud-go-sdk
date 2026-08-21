// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyVodDomainOwnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *VerifyVodDomainOwnerRequest
	GetDomainName() *string
	SetOwnerId(v int64) *VerifyVodDomainOwnerRequest
	GetOwnerId() *int64
	SetVerifyType(v string) *VerifyVodDomainOwnerRequest
	GetVerifyType() *string
}

type VerifyVodDomainOwnerRequest struct {
	// The domain name to verify. Only a single domain name can be verified at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The domain name ownership verification method. Valid values:
	//
	// - **dnsCheck**: DNS resolution verification.
	//
	// - **fileCheck**: File verification.
	//
	// This parameter is required.
	//
	// example:
	//
	// dnsCheck
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s VerifyVodDomainOwnerRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyVodDomainOwnerRequest) GoString() string {
	return s.String()
}

func (s *VerifyVodDomainOwnerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *VerifyVodDomainOwnerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *VerifyVodDomainOwnerRequest) GetVerifyType() *string {
	return s.VerifyType
}

func (s *VerifyVodDomainOwnerRequest) SetDomainName(v string) *VerifyVodDomainOwnerRequest {
	s.DomainName = &v
	return s
}

func (s *VerifyVodDomainOwnerRequest) SetOwnerId(v int64) *VerifyVodDomainOwnerRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyVodDomainOwnerRequest) SetVerifyType(v string) *VerifyVodDomainOwnerRequest {
	s.VerifyType = &v
	return s
}

func (s *VerifyVodDomainOwnerRequest) Validate() error {
	return dara.Validate(s)
}
