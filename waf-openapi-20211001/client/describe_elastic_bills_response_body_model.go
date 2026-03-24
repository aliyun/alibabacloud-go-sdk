// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticBillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBills(v []*DescribeElasticBillsResponseBodyBills) *DescribeElasticBillsResponseBody
	GetBills() []*DescribeElasticBillsResponseBodyBills
	SetRequestId(v string) *DescribeElasticBillsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeElasticBillsResponseBody
	GetTotalCount() *int64
}

type DescribeElasticBillsResponseBody struct {
	// The list of bills for the on-demand WAF instance.
	Bills []*DescribeElasticBillsResponseBodyBills `json:"Bills,omitempty" xml:"Bills,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6FBF08CB-8691-5B65-BBF8-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeElasticBillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticBillsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticBillsResponseBody) GetBills() []*DescribeElasticBillsResponseBodyBills {
	return s.Bills
}

func (s *DescribeElasticBillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticBillsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeElasticBillsResponseBody) SetBills(v []*DescribeElasticBillsResponseBodyBills) *DescribeElasticBillsResponseBody {
	s.Bills = v
	return s
}

func (s *DescribeElasticBillsResponseBody) SetRequestId(v string) *DescribeElasticBillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticBillsResponseBody) SetTotalCount(v int64) *DescribeElasticBillsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeElasticBillsResponseBody) Validate() error {
	if s.Bills != nil {
		for _, item := range s.Bills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeElasticBillsResponseBodyBills struct {
	// The total number of SCUs.
	//
	// example:
	//
	// 50
	Cu *float64 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The end time of the bill. This value is a UNIX timestamp that is in UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1717084800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of SCUs consumed by WAF features.
	//
	// example:
	//
	// 30
	FunctionCu *float64 `json:"FunctionCu,omitempty" xml:"FunctionCu,omitempty"`
	// The start time of the bill. This value is a UNIX timestamp that is in UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1665484616000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The number of security capacity units (SCUs) consumed by traffic processing.
	//
	// example:
	//
	// 20
	TrafficCu *float64 `json:"TrafficCu,omitempty" xml:"TrafficCu,omitempty"`
}

func (s DescribeElasticBillsResponseBodyBills) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticBillsResponseBodyBills) GoString() string {
	return s.String()
}

func (s *DescribeElasticBillsResponseBodyBills) GetCu() *float64 {
	return s.Cu
}

func (s *DescribeElasticBillsResponseBodyBills) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeElasticBillsResponseBodyBills) GetFunctionCu() *float64 {
	return s.FunctionCu
}

func (s *DescribeElasticBillsResponseBodyBills) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeElasticBillsResponseBodyBills) GetTrafficCu() *float64 {
	return s.TrafficCu
}

func (s *DescribeElasticBillsResponseBodyBills) SetCu(v float64) *DescribeElasticBillsResponseBodyBills {
	s.Cu = &v
	return s
}

func (s *DescribeElasticBillsResponseBodyBills) SetEndTime(v int64) *DescribeElasticBillsResponseBodyBills {
	s.EndTime = &v
	return s
}

func (s *DescribeElasticBillsResponseBodyBills) SetFunctionCu(v float64) *DescribeElasticBillsResponseBodyBills {
	s.FunctionCu = &v
	return s
}

func (s *DescribeElasticBillsResponseBodyBills) SetStartTime(v int64) *DescribeElasticBillsResponseBodyBills {
	s.StartTime = &v
	return s
}

func (s *DescribeElasticBillsResponseBodyBills) SetTrafficCu(v float64) *DescribeElasticBillsResponseBodyBills {
	s.TrafficCu = &v
	return s
}

func (s *DescribeElasticBillsResponseBodyBills) Validate() error {
	return dara.Validate(s)
}
