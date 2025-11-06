// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckIntlFixPriceDomainStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *CheckIntlFixPriceDomainStatusRequest
	GetDomain() *string
}

type CheckIntlFixPriceDomainStatusRequest struct {
	// example:
	//
	// appp16.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s CheckIntlFixPriceDomainStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckIntlFixPriceDomainStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckIntlFixPriceDomainStatusRequest) GetDomain() *string {
	return s.Domain
}

func (s *CheckIntlFixPriceDomainStatusRequest) SetDomain(v string) *CheckIntlFixPriceDomainStatusRequest {
	s.Domain = &v
	return s
}

func (s *CheckIntlFixPriceDomainStatusRequest) Validate() error {
	return dara.Validate(s)
}
