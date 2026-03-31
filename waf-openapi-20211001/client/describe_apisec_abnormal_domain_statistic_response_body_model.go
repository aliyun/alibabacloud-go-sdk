// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecAbnormalDomainStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeApisecAbnormalDomainStatisticResponseBodyData) *DescribeApisecAbnormalDomainStatisticResponseBody
	GetData() []*DescribeApisecAbnormalDomainStatisticResponseBodyData
	SetRequestId(v string) *DescribeApisecAbnormalDomainStatisticResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeApisecAbnormalDomainStatisticResponseBody
	GetTotalCount() *int64
}

type DescribeApisecAbnormalDomainStatisticResponseBody struct {
	// The response parameters.
	Data []*DescribeApisecAbnormalDomainStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Id of the request.
	//
	// example:
	//
	// 66A98669-CC6E-4F3E-80A6-3014***B11AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisecAbnormalDomainStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecAbnormalDomainStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBody) GetData() []*DescribeApisecAbnormalDomainStatisticResponseBodyData {
	return s.Data
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBody) SetData(v []*DescribeApisecAbnormalDomainStatisticResponseBodyData) *DescribeApisecAbnormalDomainStatisticResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBody) SetRequestId(v string) *DescribeApisecAbnormalDomainStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBody) SetTotalCount(v int64) *DescribeApisecAbnormalDomainStatisticResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApisecAbnormalDomainStatisticResponseBodyData struct {
	// The number of APIs.
	//
	// example:
	//
	// 10
	ApiCount *int64 `json:"ApiCount,omitempty" xml:"ApiCount,omitempty"`
	// The domain name.
	//
	// example:
	//
	// ba.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The number of high-level risks.
	//
	// example:
	//
	// 12
	High *int64 `json:"High,omitempty" xml:"High,omitempty"`
	// The number of low-level risks.
	//
	// example:
	//
	// 4
	Low *int64 `json:"Low,omitempty" xml:"Low,omitempty"`
	// The number of medium-level risks.
	//
	// example:
	//
	// 9
	Medium *int64 `json:"Medium,omitempty" xml:"Medium,omitempty"`
}

func (s DescribeApisecAbnormalDomainStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecAbnormalDomainStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBodyData) GetApiCount() *int64 {
	return s.ApiCount
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBodyData) GetDomain() *string {
	return s.Domain
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBodyData) GetHigh() *int64 {
	return s.High
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBodyData) GetLow() *int64 {
	return s.Low
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBodyData) GetMedium() *int64 {
	return s.Medium
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBodyData) SetApiCount(v int64) *DescribeApisecAbnormalDomainStatisticResponseBodyData {
	s.ApiCount = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBodyData) SetDomain(v string) *DescribeApisecAbnormalDomainStatisticResponseBodyData {
	s.Domain = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBodyData) SetHigh(v int64) *DescribeApisecAbnormalDomainStatisticResponseBodyData {
	s.High = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBodyData) SetLow(v int64) *DescribeApisecAbnormalDomainStatisticResponseBodyData {
	s.Low = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBodyData) SetMedium(v int64) *DescribeApisecAbnormalDomainStatisticResponseBodyData {
	s.Medium = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticResponseBodyData) Validate() error {
	return dara.Validate(s)
}
