// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvisorChecksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *DescribeAdvisorChecksRequest
	GetLanguage() *string
	SetProduct(v string) *DescribeAdvisorChecksRequest
	GetProduct() *string
}

type DescribeAdvisorChecksRequest struct {
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s DescribeAdvisorChecksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorChecksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeAdvisorChecksRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeAdvisorChecksRequest) SetLanguage(v string) *DescribeAdvisorChecksRequest {
	s.Language = &v
	return s
}

func (s *DescribeAdvisorChecksRequest) SetProduct(v string) *DescribeAdvisorChecksRequest {
	s.Product = &v
	return s
}

func (s *DescribeAdvisorChecksRequest) Validate() error {
	return dara.Validate(s)
}
