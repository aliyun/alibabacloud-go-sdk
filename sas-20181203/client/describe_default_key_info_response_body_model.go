// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefaultKeyInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainList(v []*string) *DescribeDefaultKeyInfoResponseBody
	GetDomainList() []*string
	SetNames(v string) *DescribeDefaultKeyInfoResponseBody
	GetNames() *string
	SetRequestId(v string) *DescribeDefaultKeyInfoResponseBody
	GetRequestId() *string
}

type DescribeDefaultKeyInfoResponseBody struct {
	// The domain names.
	DomainList []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	// The company name.
	//
	// example:
	//
	// Test
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDefaultKeyInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefaultKeyInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefaultKeyInfoResponseBody) GetDomainList() []*string {
	return s.DomainList
}

func (s *DescribeDefaultKeyInfoResponseBody) GetNames() *string {
	return s.Names
}

func (s *DescribeDefaultKeyInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefaultKeyInfoResponseBody) SetDomainList(v []*string) *DescribeDefaultKeyInfoResponseBody {
	s.DomainList = v
	return s
}

func (s *DescribeDefaultKeyInfoResponseBody) SetNames(v string) *DescribeDefaultKeyInfoResponseBody {
	s.Names = &v
	return s
}

func (s *DescribeDefaultKeyInfoResponseBody) SetRequestId(v string) *DescribeDefaultKeyInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefaultKeyInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
