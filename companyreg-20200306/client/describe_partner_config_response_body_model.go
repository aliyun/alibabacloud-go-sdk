// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePartnerConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContact(v string) *DescribePartnerConfigResponseBody
	GetContact() *string
	SetPartnerCode(v string) *DescribePartnerConfigResponseBody
	GetPartnerCode() *string
	SetPartnerName(v string) *DescribePartnerConfigResponseBody
	GetPartnerName() *string
	SetRequestId(v string) *DescribePartnerConfigResponseBody
	GetRequestId() *string
}

type DescribePartnerConfigResponseBody struct {
	Contact *string `json:"Contact,omitempty" xml:"Contact,omitempty"`
	// example:
	//
	// jinsan
	PartnerCode *string `json:"PartnerCode,omitempty" xml:"PartnerCode,omitempty"`
	PartnerName *string `json:"PartnerName,omitempty" xml:"PartnerName,omitempty"`
	// example:
	//
	// 8179A0B3-A5D3-52F4-8845-F0ABC3CC6783
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePartnerConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePartnerConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePartnerConfigResponseBody) GetContact() *string {
	return s.Contact
}

func (s *DescribePartnerConfigResponseBody) GetPartnerCode() *string {
	return s.PartnerCode
}

func (s *DescribePartnerConfigResponseBody) GetPartnerName() *string {
	return s.PartnerName
}

func (s *DescribePartnerConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePartnerConfigResponseBody) SetContact(v string) *DescribePartnerConfigResponseBody {
	s.Contact = &v
	return s
}

func (s *DescribePartnerConfigResponseBody) SetPartnerCode(v string) *DescribePartnerConfigResponseBody {
	s.PartnerCode = &v
	return s
}

func (s *DescribePartnerConfigResponseBody) SetPartnerName(v string) *DescribePartnerConfigResponseBody {
	s.PartnerName = &v
	return s
}

func (s *DescribePartnerConfigResponseBody) SetRequestId(v string) *DescribePartnerConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePartnerConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
