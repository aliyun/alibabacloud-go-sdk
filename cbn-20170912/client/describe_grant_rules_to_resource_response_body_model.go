// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantRulesToResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGrantRules(v []*DescribeGrantRulesToResourceResponseBodyGrantRules) *DescribeGrantRulesToResourceResponseBody
	GetGrantRules() []*DescribeGrantRulesToResourceResponseBodyGrantRules
	SetMaxResults(v int32) *DescribeGrantRulesToResourceResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeGrantRulesToResourceResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeGrantRulesToResourceResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeGrantRulesToResourceResponseBody
	GetTotalCount() *int32
}

type DescribeGrantRulesToResourceResponseBody struct {
	// A list of permission records.
	GrantRules []*DescribeGrantRulesToResourceResponseBodyGrantRules `json:"GrantRules,omitempty" xml:"GrantRules,omitempty" type:"Repeated"`
	// - If the **MaxResults*	- parameter was not included in the request, this field indicates the total number of entries.
	//
	// - If the **MaxResults*	- parameter was included in the request, this field indicates the number of entries on the current page.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to use in your next request to retrieve a new page of results. Valid values:
	//
	// - If the **NextToken*	- parameter is empty, no more results are available.
	//
	// - If a value is returned for **NextToken**, use it for the next request to get the next page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C6E5992C-A57B-5A6C-9B26-568074DC68BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGrantRulesToResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToResourceResponseBody) GetGrantRules() []*DescribeGrantRulesToResourceResponseBodyGrantRules {
	return s.GrantRules
}

func (s *DescribeGrantRulesToResourceResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeGrantRulesToResourceResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeGrantRulesToResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGrantRulesToResourceResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeGrantRulesToResourceResponseBody) SetGrantRules(v []*DescribeGrantRulesToResourceResponseBodyGrantRules) *DescribeGrantRulesToResourceResponseBody {
	s.GrantRules = v
	return s
}

func (s *DescribeGrantRulesToResourceResponseBody) SetMaxResults(v int32) *DescribeGrantRulesToResourceResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeGrantRulesToResourceResponseBody) SetNextToken(v string) *DescribeGrantRulesToResourceResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeGrantRulesToResourceResponseBody) SetRequestId(v string) *DescribeGrantRulesToResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGrantRulesToResourceResponseBody) SetTotalCount(v int32) *DescribeGrantRulesToResourceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGrantRulesToResourceResponseBody) Validate() error {
	if s.GrantRules != nil {
		for _, item := range s.GrantRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGrantRulesToResourceResponseBodyGrantRules struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-44m0p68spvlrqq****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the main account that owns the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// 1250123456123456
	CenOwnerId *int64 `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The timestamp when the permission was granted. The value is in UTC and follows the ISO 8601 standard: `YYYY-MM-DDThh:mmZ`.
	//
	// example:
	//
	// 2024-01-24T16:46Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// PayByCenOwner
	EffectiveOrderType *string `json:"EffectiveOrderType,omitempty" xml:"EffectiveOrderType,omitempty"`
	// The payer for the network instance. Valid values:
	//
	// - **PayByCenOwner**: The CEN instance owner pays the Transit Router connection and data processing fees for the network instance.
	//
	// - **PayByResourceOwner**: The network instance owner pays the Transit Router connection and data processing fees for the network instance.
	//
	// example:
	//
	// PayByCenOwner
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s DescribeGrantRulesToResourceResponseBodyGrantRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToResourceResponseBodyGrantRules) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToResourceResponseBodyGrantRules) GetCenId() *string {
	return s.CenId
}

func (s *DescribeGrantRulesToResourceResponseBodyGrantRules) GetCenOwnerId() *int64 {
	return s.CenOwnerId
}

func (s *DescribeGrantRulesToResourceResponseBodyGrantRules) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGrantRulesToResourceResponseBodyGrantRules) GetEffectiveOrderType() *string {
	return s.EffectiveOrderType
}

func (s *DescribeGrantRulesToResourceResponseBodyGrantRules) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeGrantRulesToResourceResponseBodyGrantRules) SetCenId(v string) *DescribeGrantRulesToResourceResponseBodyGrantRules {
	s.CenId = &v
	return s
}

func (s *DescribeGrantRulesToResourceResponseBodyGrantRules) SetCenOwnerId(v int64) *DescribeGrantRulesToResourceResponseBodyGrantRules {
	s.CenOwnerId = &v
	return s
}

func (s *DescribeGrantRulesToResourceResponseBodyGrantRules) SetCreateTime(v string) *DescribeGrantRulesToResourceResponseBodyGrantRules {
	s.CreateTime = &v
	return s
}

func (s *DescribeGrantRulesToResourceResponseBodyGrantRules) SetEffectiveOrderType(v string) *DescribeGrantRulesToResourceResponseBodyGrantRules {
	s.EffectiveOrderType = &v
	return s
}

func (s *DescribeGrantRulesToResourceResponseBodyGrantRules) SetOrderType(v string) *DescribeGrantRulesToResourceResponseBodyGrantRules {
	s.OrderType = &v
	return s
}

func (s *DescribeGrantRulesToResourceResponseBodyGrantRules) Validate() error {
	return dara.Validate(s)
}
