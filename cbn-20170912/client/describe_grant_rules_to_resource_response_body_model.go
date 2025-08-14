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
	// The permissions that are granted to the CEN instance.
	GrantRules []*DescribeGrantRulesToResourceResponseBodyGrantRules `json:"GrantRules,omitempty" xml:"GrantRules,omitempty" type:"Repeated"`
	// 	- If no value is specified for **MaxResults**, query results are returned in one batch. The value of **MaxResults*	- indicates the total number of entries.
	//
	// 	- If a value is specified for **MaxResults**, query results are returned in batches. The value of **MaxResults*	- in the response indicates the number of entries in the current batch.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If the **NextToken*	- parameter is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
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
	return dara.Validate(s)
}

type DescribeGrantRulesToResourceResponseBodyGrantRules struct {
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-44m0p68spvlrqq****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the Alibaba Cloud account to which the CEN instance belongs.
	//
	// example:
	//
	// 1250123456123456
	CenOwnerId *int64 `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The timestamp when the permissions were granted. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-01-24T16:46Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The entity that pays the fees of the network instance. Valid values: Valid values:
	//
	// 	- **PayByCenOwner**: The fees of the connections and data forwarding on the transit router are paid by the Alibaba Cloud account to which the CEN instance belongs.
	//
	// 	- **PayByResourceOwner**: The fees of the connections and data forwarding on the transit router are paid by the Alibaba Cloud account to which the network instance belongs.
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

func (s *DescribeGrantRulesToResourceResponseBodyGrantRules) SetOrderType(v string) *DescribeGrantRulesToResourceResponseBodyGrantRules {
	s.OrderType = &v
	return s
}

func (s *DescribeGrantRulesToResourceResponseBodyGrantRules) Validate() error {
	return dara.Validate(s)
}
