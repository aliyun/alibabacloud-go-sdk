// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecEventDomainStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeApisecEventDomainStatisticResponseBodyData) *DescribeApisecEventDomainStatisticResponseBody
	GetData() []*DescribeApisecEventDomainStatisticResponseBodyData
	SetRequestId(v string) *DescribeApisecEventDomainStatisticResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeApisecEventDomainStatisticResponseBody
	GetTotalCount() *int64
}

type DescribeApisecEventDomainStatisticResponseBody struct {
	// The response parameters.
	Data []*DescribeApisecEventDomainStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Id of the request.
	//
	// example:
	//
	// 66A98669-*******-80A6-3014697B11AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisecEventDomainStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecEventDomainStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecEventDomainStatisticResponseBody) GetData() []*DescribeApisecEventDomainStatisticResponseBodyData {
	return s.Data
}

func (s *DescribeApisecEventDomainStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecEventDomainStatisticResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeApisecEventDomainStatisticResponseBody) SetData(v []*DescribeApisecEventDomainStatisticResponseBodyData) *DescribeApisecEventDomainStatisticResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApisecEventDomainStatisticResponseBody) SetRequestId(v string) *DescribeApisecEventDomainStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticResponseBody) SetTotalCount(v int64) *DescribeApisecEventDomainStatisticResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticResponseBody) Validate() error {
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

type DescribeApisecEventDomainStatisticResponseBodyData struct {
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
	// a.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The number of high-risk security events.
	//
	// example:
	//
	// 10
	High *int64 `json:"High,omitempty" xml:"High,omitempty"`
	// The number of low-risk security events.
	//
	// example:
	//
	// 2
	Low *int64 `json:"Low,omitempty" xml:"Low,omitempty"`
	// The number of medium-risk security events.
	//
	// example:
	//
	// 6
	Medium *int64 `json:"Medium,omitempty" xml:"Medium,omitempty"`
}

func (s DescribeApisecEventDomainStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecEventDomainStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApisecEventDomainStatisticResponseBodyData) GetApiCount() *int64 {
	return s.ApiCount
}

func (s *DescribeApisecEventDomainStatisticResponseBodyData) GetDomain() *string {
	return s.Domain
}

func (s *DescribeApisecEventDomainStatisticResponseBodyData) GetHigh() *int64 {
	return s.High
}

func (s *DescribeApisecEventDomainStatisticResponseBodyData) GetLow() *int64 {
	return s.Low
}

func (s *DescribeApisecEventDomainStatisticResponseBodyData) GetMedium() *int64 {
	return s.Medium
}

func (s *DescribeApisecEventDomainStatisticResponseBodyData) SetApiCount(v int64) *DescribeApisecEventDomainStatisticResponseBodyData {
	s.ApiCount = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticResponseBodyData) SetDomain(v string) *DescribeApisecEventDomainStatisticResponseBodyData {
	s.Domain = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticResponseBodyData) SetHigh(v int64) *DescribeApisecEventDomainStatisticResponseBodyData {
	s.High = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticResponseBodyData) SetLow(v int64) *DescribeApisecEventDomainStatisticResponseBodyData {
	s.Low = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticResponseBodyData) SetMedium(v int64) *DescribeApisecEventDomainStatisticResponseBodyData {
	s.Medium = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticResponseBodyData) Validate() error {
	return dara.Validate(s)
}
