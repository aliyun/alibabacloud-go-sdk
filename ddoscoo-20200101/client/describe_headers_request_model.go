// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHeadersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeHeadersRequest
	GetDomain() *string
	SetResourceGroupId(v string) *DescribeHeadersRequest
	GetResourceGroupId() *string
}

type DescribeHeadersRequest struct {
	// The domain name that you want to query.
	//
	// > You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all the domain names that are added to Anti-DDoS Pro or Anti-DDoS Premium.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-aek3cmuvpia****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeHeadersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHeadersRequest) GoString() string {
	return s.String()
}

func (s *DescribeHeadersRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeHeadersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeHeadersRequest) SetDomain(v string) *DescribeHeadersRequest {
	s.Domain = &v
	return s
}

func (s *DescribeHeadersRequest) SetResourceGroupId(v string) *DescribeHeadersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHeadersRequest) Validate() error {
	return dara.Validate(s)
}
