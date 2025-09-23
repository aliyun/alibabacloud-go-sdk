// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimpleDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainList(v []*string) *DescribeSimpleDomainsResponseBody
	GetDomainList() []*string
	SetRequestId(v string) *DescribeSimpleDomainsResponseBody
	GetRequestId() *string
}

type DescribeSimpleDomainsResponseBody struct {
	DomainList []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSimpleDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimpleDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSimpleDomainsResponseBody) GetDomainList() []*string {
	return s.DomainList
}

func (s *DescribeSimpleDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSimpleDomainsResponseBody) SetDomainList(v []*string) *DescribeSimpleDomainsResponseBody {
	s.DomainList = v
	return s
}

func (s *DescribeSimpleDomainsResponseBody) SetRequestId(v string) *DescribeSimpleDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSimpleDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}
