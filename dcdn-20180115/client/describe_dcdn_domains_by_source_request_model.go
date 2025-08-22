// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainsBySourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSources(v string) *DescribeDcdnDomainsBySourceRequest
	GetSources() *string
}

type DescribeDcdnDomainsBySourceRequest struct {
	// The list of origin servers. Separate origin servers with commas (,). You can specify a maximum of 20 origin servers. Fuzzy match is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.org
	Sources *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
}

func (s DescribeDcdnDomainsBySourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainsBySourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainsBySourceRequest) GetSources() *string {
	return s.Sources
}

func (s *DescribeDcdnDomainsBySourceRequest) SetSources(v string) *DescribeDcdnDomainsBySourceRequest {
	s.Sources = &v
	return s
}

func (s *DescribeDcdnDomainsBySourceRequest) Validate() error {
	return dara.Validate(s)
}
