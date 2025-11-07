// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationGroupBillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationGroupConsume(v []*DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) *DescribeApplicationGroupBillResponseBody
	GetApplicationGroupConsume() []*DescribeApplicationGroupBillResponseBodyApplicationGroupConsume
	SetMaxResults(v int32) *DescribeApplicationGroupBillResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeApplicationGroupBillResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeApplicationGroupBillResponseBody
	GetRequestId() *string
}

type DescribeApplicationGroupBillResponseBody struct {
	// The consume of application group.
	ApplicationGroupConsume []*DescribeApplicationGroupBillResponseBodyApplicationGroupConsume `json:"ApplicationGroupConsume,omitempty" xml:"ApplicationGroupConsume,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// ""
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E897A1AB-4701-5B71-93AD-38FD703763A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApplicationGroupBillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationGroupBillResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationGroupBillResponseBody) GetApplicationGroupConsume() []*DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	return s.ApplicationGroupConsume
}

func (s *DescribeApplicationGroupBillResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeApplicationGroupBillResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeApplicationGroupBillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationGroupBillResponseBody) SetApplicationGroupConsume(v []*DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) *DescribeApplicationGroupBillResponseBody {
	s.ApplicationGroupConsume = v
	return s
}

func (s *DescribeApplicationGroupBillResponseBody) SetMaxResults(v int32) *DescribeApplicationGroupBillResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBody) SetNextToken(v string) *DescribeApplicationGroupBillResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBody) SetRequestId(v string) *DescribeApplicationGroupBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBody) Validate() error {
	if s.ApplicationGroupConsume != nil {
		for _, item := range s.ApplicationGroupConsume {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationGroupBillResponseBodyApplicationGroupConsume struct {
	// The amount consumed by the instance.
	//
	// example:
	//
	// 18.88
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2023-06-10T06:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The currency unit.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-0jl781czrhqey0s5zpsj
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// test-
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ALIYUN::ECS::INSTANCE
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Optimization suggestions.
	//
	// example:
	//
	// 1
	Optimization *string `json:"Optimization,omitempty" xml:"Optimization,omitempty"`
	// The peak type.
	//
	// example:
	//
	// WHITE
	PeakType *string `json:"PeakType,omitempty" xml:"PeakType,omitempty"`
	// The performance of the data synchronization instance.
	//
	// example:
	//
	// "{\\"mem\\":\\"6.82\\",\\"cpu\\":\\"0.55\\"}"
	Performance *string `json:"Performance,omitempty" xml:"Performance,omitempty"`
	// The status of instance.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) GoString() string {
	return s.String()
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) GetAmount() *float32 {
	return s.Amount
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) GetOptimization() *string {
	return s.Optimization
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) GetPeakType() *string {
	return s.PeakType
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) GetPerformance() *string {
	return s.Performance
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) GetStatus() *string {
	return s.Status
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetAmount(v float32) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.Amount = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetCreationTime(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.CreationTime = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetCurrency(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.Currency = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetInstanceId(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.InstanceId = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetInstanceName(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.InstanceName = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetInstanceType(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.InstanceType = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetOptimization(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.Optimization = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetPeakType(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.PeakType = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetPerformance(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.Performance = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetStatus(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.Status = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) Validate() error {
	return dara.Validate(s)
}
