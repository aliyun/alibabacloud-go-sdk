// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainStagingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeCdnDomainStagingConfigRequest
	GetDomainName() *string
	SetFunctionNames(v string) *DescribeCdnDomainStagingConfigRequest
	GetFunctionNames() *string
}

type DescribeCdnDomainStagingConfigRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The list of feature names. Separate multiple values with commas (,). For more information, see [A list of features](https://help.aliyun.com/document_detail/388460.html).
	//
	// example:
	//
	// aliauth
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
}

func (s DescribeCdnDomainStagingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainStagingConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnDomainStagingConfigRequest) GetFunctionNames() *string {
	return s.FunctionNames
}

func (s *DescribeCdnDomainStagingConfigRequest) SetDomainName(v string) *DescribeCdnDomainStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigRequest) SetFunctionNames(v string) *DescribeCdnDomainStagingConfigRequest {
	s.FunctionNames = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigRequest) Validate() error {
	return dara.Validate(s)
}
