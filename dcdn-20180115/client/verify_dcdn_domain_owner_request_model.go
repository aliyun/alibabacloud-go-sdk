// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyDcdnDomainOwnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *VerifyDcdnDomainOwnerRequest
	GetDomainName() *string
	SetVerifyType(v string) *VerifyDcdnDomainOwnerRequest
	GetVerifyType() *string
}

type VerifyDcdnDomainOwnerRequest struct {
	// The domain name of which you want to verify the ownership. You can specify only one domain name in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// **example**.com
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

func (s VerifyDcdnDomainOwnerRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyDcdnDomainOwnerRequest) GoString() string {
	return s.String()
}

func (s *VerifyDcdnDomainOwnerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *VerifyDcdnDomainOwnerRequest) GetVerifyType() *string {
	return s.VerifyType
}

func (s *VerifyDcdnDomainOwnerRequest) SetDomainName(v string) *VerifyDcdnDomainOwnerRequest {
	s.DomainName = &v
	return s
}

func (s *VerifyDcdnDomainOwnerRequest) SetVerifyType(v string) *VerifyDcdnDomainOwnerRequest {
	s.VerifyType = &v
	return s
}

func (s *VerifyDcdnDomainOwnerRequest) Validate() error {
	return dara.Validate(s)
}
