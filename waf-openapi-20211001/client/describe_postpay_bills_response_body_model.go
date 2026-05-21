// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayBillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBillDetail(v []*DescribePostpayBillsResponseBodyBillDetail) *DescribePostpayBillsResponseBody
	GetBillDetail() []*DescribePostpayBillsResponseBodyBillDetail
	SetMaxResults(v int32) *DescribePostpayBillsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePostpayBillsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribePostpayBillsResponseBody
	GetRequestId() *string
}

type DescribePostpayBillsResponseBody struct {
	BillDetail []*DescribePostpayBillsResponseBodyBillDetail `json:"BillDetail,omitempty" xml:"BillDetail,omitempty" type:"Repeated"`
	// example:
	//
	// 24
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 0FBBDE11-C35F-531B-96BA-64CA****C875
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePostpayBillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayBillsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePostpayBillsResponseBody) GetBillDetail() []*DescribePostpayBillsResponseBodyBillDetail {
	return s.BillDetail
}

func (s *DescribePostpayBillsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePostpayBillsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePostpayBillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePostpayBillsResponseBody) SetBillDetail(v []*DescribePostpayBillsResponseBodyBillDetail) *DescribePostpayBillsResponseBody {
	s.BillDetail = v
	return s
}

func (s *DescribePostpayBillsResponseBody) SetMaxResults(v int32) *DescribePostpayBillsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribePostpayBillsResponseBody) SetNextToken(v string) *DescribePostpayBillsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribePostpayBillsResponseBody) SetRequestId(v string) *DescribePostpayBillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePostpayBillsResponseBody) Validate() error {
	if s.BillDetail != nil {
		for _, item := range s.BillDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePostpayBillsResponseBodyBillDetail struct {
	// example:
	//
	// {\\"aiWhiteListTemplateCount\\":4,\\"apisecResourceCount\\":2,\\"botAppTemplateCount\\":1,\\"botWebTemplateCount\\":4,\\"ccRuleCount\\":1,\\"customAclAdvanceRuleCount\\":4,\\"customResponseRuleCount\\":5,\\"dlpRuleCount\\":1,\\"gslb\\":1,\\"instanceFee\\":1,\\"ipv6\\":1,\\"nonPort\\":1,\\"qps\\":0,\\"regionBlockRuleCount\\":1,\\"threatIntelligenceTemplateCount\\":1,\\"wafBaseTemplateCount\\":4}
	ChargeData *string `json:"ChargeData,omitempty" xml:"ChargeData,omitempty"`
	// example:
	//
	// 51
	Cu *string `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// example:
	//
	// 1779123599
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 30
	FunctionCu *string `json:"FunctionCu,omitempty" xml:"FunctionCu,omitempty"`
	// example:
	//
	// 1779120000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 21
	TrafficCu *string `json:"TrafficCu,omitempty" xml:"TrafficCu,omitempty"`
}

func (s DescribePostpayBillsResponseBodyBillDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayBillsResponseBodyBillDetail) GoString() string {
	return s.String()
}

func (s *DescribePostpayBillsResponseBodyBillDetail) GetChargeData() *string {
	return s.ChargeData
}

func (s *DescribePostpayBillsResponseBodyBillDetail) GetCu() *string {
	return s.Cu
}

func (s *DescribePostpayBillsResponseBodyBillDetail) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePostpayBillsResponseBodyBillDetail) GetFunctionCu() *string {
	return s.FunctionCu
}

func (s *DescribePostpayBillsResponseBodyBillDetail) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePostpayBillsResponseBodyBillDetail) GetTrafficCu() *string {
	return s.TrafficCu
}

func (s *DescribePostpayBillsResponseBodyBillDetail) SetChargeData(v string) *DescribePostpayBillsResponseBodyBillDetail {
	s.ChargeData = &v
	return s
}

func (s *DescribePostpayBillsResponseBodyBillDetail) SetCu(v string) *DescribePostpayBillsResponseBodyBillDetail {
	s.Cu = &v
	return s
}

func (s *DescribePostpayBillsResponseBodyBillDetail) SetEndTime(v int64) *DescribePostpayBillsResponseBodyBillDetail {
	s.EndTime = &v
	return s
}

func (s *DescribePostpayBillsResponseBodyBillDetail) SetFunctionCu(v string) *DescribePostpayBillsResponseBodyBillDetail {
	s.FunctionCu = &v
	return s
}

func (s *DescribePostpayBillsResponseBodyBillDetail) SetStartTime(v int64) *DescribePostpayBillsResponseBodyBillDetail {
	s.StartTime = &v
	return s
}

func (s *DescribePostpayBillsResponseBodyBillDetail) SetTrafficCu(v string) *DescribePostpayBillsResponseBodyBillDetail {
	s.TrafficCu = &v
	return s
}

func (s *DescribePostpayBillsResponseBodyBillDetail) Validate() error {
	return dara.Validate(s)
}
