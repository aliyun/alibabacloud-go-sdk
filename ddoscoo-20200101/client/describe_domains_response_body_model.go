// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*string) *DescribeDomainsResponseBody
	GetDomains() []*string
	SetRequestId(v string) *DescribeDomainsResponseBody
	GetRequestId() *string
}

type DescribeDomainsResponseBody struct {
	// An array consisting of details of the domain name for which the forwarding rules are configured.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// F908E959-ADA8-4D7B-8A05-FF2F67F50964
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBody) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainsResponseBody) SetDomains(v []*string) *DescribeDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDomainsResponseBody) SetRequestId(v string) *DescribeDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}
