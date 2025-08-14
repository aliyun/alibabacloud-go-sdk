// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantRulesToCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGrantRules(v *DescribeGrantRulesToCenResponseBodyGrantRules) *DescribeGrantRulesToCenResponseBody
	GetGrantRules() *DescribeGrantRulesToCenResponseBodyGrantRules
	SetMaxResults(v int64) *DescribeGrantRulesToCenResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeGrantRulesToCenResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeGrantRulesToCenResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeGrantRulesToCenResponseBody
	GetTotalCount() *int64
}

type DescribeGrantRulesToCenResponseBody struct {
	// The permissions that are granted to the CEN instance.
	GrantRules *DescribeGrantRulesToCenResponseBodyGrantRules `json:"GrantRules,omitempty" xml:"GrantRules,omitempty" type:"Struct"`
	// 	- If no value is specified for **MaxResults**, query results are returned in one batch. The value of **MaxResults*	- indicates the total number of entries.
	//
	// 	- If a value is specified for **MaxResults**, it indicates that you need to query results in batches. The value of **MaxResults*	- in the response indicates the number of entries in the current batch.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- was returned in the previous query, specify the value to obtain the next set of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4906B209-8613-5C19-9CC9-B7A3FFDA731C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGrantRulesToCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToCenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenResponseBody) GetGrantRules() *DescribeGrantRulesToCenResponseBodyGrantRules {
	return s.GrantRules
}

func (s *DescribeGrantRulesToCenResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeGrantRulesToCenResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeGrantRulesToCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGrantRulesToCenResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeGrantRulesToCenResponseBody) SetGrantRules(v *DescribeGrantRulesToCenResponseBodyGrantRules) *DescribeGrantRulesToCenResponseBody {
	s.GrantRules = v
	return s
}

func (s *DescribeGrantRulesToCenResponseBody) SetMaxResults(v int64) *DescribeGrantRulesToCenResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBody) SetNextToken(v string) *DescribeGrantRulesToCenResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBody) SetRequestId(v string) *DescribeGrantRulesToCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBody) SetTotalCount(v int64) *DescribeGrantRulesToCenResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGrantRulesToCenResponseBodyGrantRules struct {
	GrantRule []*DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule `json:"GrantRule,omitempty" xml:"GrantRule,omitempty" type:"Repeated"`
}

func (s DescribeGrantRulesToCenResponseBodyGrantRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToCenResponseBodyGrantRules) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRules) GetGrantRule() []*DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	return s.GrantRule
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRules) SetGrantRule(v []*DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) *DescribeGrantRulesToCenResponseBodyGrantRules {
	s.GrantRule = v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRules) Validate() error {
	return dara.Validate(s)
}

type DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule struct {
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-nye53d7p3hzyu4****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the Alibaba Cloud account to which the CEN instance belongs.
	//
	// example:
	//
	// 1210123456123456
	CenOwnerId *int64 `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The ID of the network instance.
	//
	// example:
	//
	// vpc-bp1rgeww9mdstuuar****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the network instance belongs.
	//
	// example:
	//
	// 1250123456123456
	ChildInstanceOwnerId *int64 `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	// The region ID of the network instance.
	//
	// example:
	//
	// cn-hangzhou
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**: VPC
	//
	// 	- **VBR**: VBR
	//
	// 	- **CCN**: CCN instance
	//
	// 	- **VPN**: IPsec-VPN connection
	//
	// example:
	//
	// VPC
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	// The time when the permissions were granted to the CEN instance.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-01-24T16:27Z
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The entity that pays the fees of the network instance. Valid values:
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

func (s DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) GetCenId() *string {
	return s.CenId
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) GetCenOwnerId() *int64 {
	return s.CenOwnerId
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) GetChildInstanceOwnerId() *int64 {
	return s.ChildInstanceOwnerId
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetCenId(v string) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.CenId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetCenOwnerId(v int64) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.CenOwnerId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetChildInstanceId(v string) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetChildInstanceOwnerId(v int64) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetChildInstanceRegionId(v string) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetChildInstanceType(v string) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetCreateTime(v int64) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.CreateTime = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) SetOrderType(v string) *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule {
	s.OrderType = &v
	return s
}

func (s *DescribeGrantRulesToCenResponseBodyGrantRulesGrantRule) Validate() error {
	return dara.Validate(s)
}
