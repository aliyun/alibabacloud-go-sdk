// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainStagingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainStagingConfigRequest
	GetDomainName() *string
	SetFunctionNames(v string) *DescribeDcdnDomainStagingConfigRequest
	GetFunctionNames() *string
}

type DescribeDcdnDomainStagingConfigRequest struct {
	// The accelerated domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The names of the features to query. You can separate multiple features with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// aliauth
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
}

func (s DescribeDcdnDomainStagingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainStagingConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainStagingConfigRequest) GetFunctionNames() *string {
	return s.FunctionNames
}

func (s *DescribeDcdnDomainStagingConfigRequest) SetDomainName(v string) *DescribeDcdnDomainStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainStagingConfigRequest) SetFunctionNames(v string) *DescribeDcdnDomainStagingConfigRequest {
	s.FunctionNames = &v
	return s
}

func (s *DescribeDcdnDomainStagingConfigRequest) Validate() error {
	return dara.Validate(s)
}
