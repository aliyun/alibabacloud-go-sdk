// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRDDomainConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeRDDomainConfigRequest
	GetDomainName() *string
	SetFunctionName(v string) *DescribeRDDomainConfigRequest
	GetFunctionName() *string
}

type DescribeRDDomainConfigRequest struct {
	// The accelerated domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the feature. Default value: source_group.
	//
	// example:
	//
	// source_group
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s DescribeRDDomainConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDDomainConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeRDDomainConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeRDDomainConfigRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeRDDomainConfigRequest) SetDomainName(v string) *DescribeRDDomainConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeRDDomainConfigRequest) SetFunctionName(v string) *DescribeRDDomainConfigRequest {
	s.FunctionName = &v
	return s
}

func (s *DescribeRDDomainConfigRequest) Validate() error {
	return dara.Validate(s)
}
