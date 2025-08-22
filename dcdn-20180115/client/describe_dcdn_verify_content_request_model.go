// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnVerifyContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnVerifyContentRequest
	GetDomainName() *string
}

type DescribeDcdnVerifyContentRequest struct {
	// The domain name for which you want to query the ownership verification content. You can specify only one domain name in one request.
	//
	// This parameter is required.
	//
	// example:
	//
	// **example**.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDcdnVerifyContentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnVerifyContentRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnVerifyContentRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnVerifyContentRequest) SetDomainName(v string) *DescribeDcdnVerifyContentRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnVerifyContentRequest) Validate() error {
	return dara.Validate(s)
}
