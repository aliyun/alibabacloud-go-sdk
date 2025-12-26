// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressList(v []*DescribeAddressesResponseBodyAddressList) *DescribeAddressesResponseBody
	GetAddressList() []*DescribeAddressesResponseBodyAddressList
	SetMaxResults(v int32) *DescribeAddressesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAddressesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeAddressesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeAddressesResponseBody
	GetTotalCount() *int64
}

type DescribeAddressesResponseBody struct {
	AddressList []*DescribeAddressesResponseBodyAddressList `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 8D8EBFB7-E1EB-5236-952A-092EDC72***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAddressesResponseBody) GetAddressList() []*DescribeAddressesResponseBodyAddressList {
	return s.AddressList
}

func (s *DescribeAddressesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAddressesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAddressesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAddressesResponseBody) SetAddressList(v []*DescribeAddressesResponseBodyAddressList) *DescribeAddressesResponseBody {
	s.AddressList = v
	return s
}

func (s *DescribeAddressesResponseBody) SetMaxResults(v int32) *DescribeAddressesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeAddressesResponseBody) SetNextToken(v string) *DescribeAddressesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAddressesResponseBody) SetRequestId(v string) *DescribeAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAddressesResponseBody) SetTotalCount(v int64) *DescribeAddressesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAddressesResponseBody) Validate() error {
	if s.AddressList != nil {
		for _, item := range s.AddressList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAddressesResponseBodyAddressList struct {
	// example:
	//
	// 1.1.1.1
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// 1760408233000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 12345678
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeAddressesResponseBodyAddressList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddressesResponseBodyAddressList) GoString() string {
	return s.String()
}

func (s *DescribeAddressesResponseBodyAddressList) GetAddress() *string {
	return s.Address
}

func (s *DescribeAddressesResponseBodyAddressList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeAddressesResponseBodyAddressList) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeAddressesResponseBodyAddressList) SetAddress(v string) *DescribeAddressesResponseBodyAddressList {
	s.Address = &v
	return s
}

func (s *DescribeAddressesResponseBodyAddressList) SetGmtModified(v int64) *DescribeAddressesResponseBodyAddressList {
	s.GmtModified = &v
	return s
}

func (s *DescribeAddressesResponseBodyAddressList) SetRuleId(v int64) *DescribeAddressesResponseBodyAddressList {
	s.RuleId = &v
	return s
}

func (s *DescribeAddressesResponseBodyAddressList) Validate() error {
	return dara.Validate(s)
}
