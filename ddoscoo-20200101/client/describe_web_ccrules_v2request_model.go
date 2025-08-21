// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCCRulesV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeWebCCRulesV2Request
	GetDomain() *string
	SetOffset(v string) *DescribeWebCCRulesV2Request
	GetOffset() *string
	SetOwner(v string) *DescribeWebCCRulesV2Request
	GetOwner() *string
	SetPageSize(v string) *DescribeWebCCRulesV2Request
	GetPageSize() *string
}

type DescribeWebCCRulesV2Request struct {
	// The domain name of the website that you want to add to the Anti-DDoS Proxy instance for protection.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The number of entries that you want the system to skip before the system returns entries. Default value: **0**.
	//
	// example:
	//
	// 0
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The method used to create the rule. Valid values:
	//
	// 	- **manual*	- (default): manually created.
	//
	// 	- **clover**: automatically created.
	//
	// example:
	//
	// manual
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The number of entries per page. Maximum value: **20**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeWebCCRulesV2Request) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCCRulesV2Request) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesV2Request) GetDomain() *string {
	return s.Domain
}

func (s *DescribeWebCCRulesV2Request) GetOffset() *string {
	return s.Offset
}

func (s *DescribeWebCCRulesV2Request) GetOwner() *string {
	return s.Owner
}

func (s *DescribeWebCCRulesV2Request) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeWebCCRulesV2Request) SetDomain(v string) *DescribeWebCCRulesV2Request {
	s.Domain = &v
	return s
}

func (s *DescribeWebCCRulesV2Request) SetOffset(v string) *DescribeWebCCRulesV2Request {
	s.Offset = &v
	return s
}

func (s *DescribeWebCCRulesV2Request) SetOwner(v string) *DescribeWebCCRulesV2Request {
	s.Owner = &v
	return s
}

func (s *DescribeWebCCRulesV2Request) SetPageSize(v string) *DescribeWebCCRulesV2Request {
	s.PageSize = &v
	return s
}

func (s *DescribeWebCCRulesV2Request) Validate() error {
	return dara.Validate(s)
}
