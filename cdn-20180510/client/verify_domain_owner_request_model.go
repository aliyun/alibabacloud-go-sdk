// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyDomainOwnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *VerifyDomainOwnerRequest
	GetDomainName() *string
	SetVerifyType(v string) *VerifyDomainOwnerRequest
	GetVerifyType() *string
}

type VerifyDomainOwnerRequest struct {
	// The domain name of which you want to verify the ownership. You can specify only one domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The verification method. Valid values:
	//
	// 	- **dnsCheck**: by DNS record
	//
	// 	- **fileCheck**: by verification file
	//
	// This parameter is required.
	//
	// example:
	//
	// dnsCheck
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s VerifyDomainOwnerRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyDomainOwnerRequest) GoString() string {
	return s.String()
}

func (s *VerifyDomainOwnerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *VerifyDomainOwnerRequest) GetVerifyType() *string {
	return s.VerifyType
}

func (s *VerifyDomainOwnerRequest) SetDomainName(v string) *VerifyDomainOwnerRequest {
	s.DomainName = &v
	return s
}

func (s *VerifyDomainOwnerRequest) SetVerifyType(v string) *VerifyDomainOwnerRequest {
	s.VerifyType = &v
	return s
}

func (s *VerifyDomainOwnerRequest) Validate() error {
	return dara.Validate(s)
}
