// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfiguredDestinationIPResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDestinations(v []*DescribeConfiguredDestinationIPResponseBodyDestinations) *DescribeConfiguredDestinationIPResponseBody
	GetDestinations() []*DescribeConfiguredDestinationIPResponseBodyDestinations
	SetPageNo(v int32) *DescribeConfiguredDestinationIPResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeConfiguredDestinationIPResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeConfiguredDestinationIPResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeConfiguredDestinationIPResponseBody
	GetTotalCount() *int32
}

type DescribeConfiguredDestinationIPResponseBody struct {
	Destinations []*DescribeConfiguredDestinationIPResponseBodyDestinations `json:"Destinations,omitempty" xml:"Destinations,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// F0AE91F8-E6C5-50D4-983F-FC53672****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 32
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeConfiguredDestinationIPResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfiguredDestinationIPResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfiguredDestinationIPResponseBody) GetDestinations() []*DescribeConfiguredDestinationIPResponseBodyDestinations {
	return s.Destinations
}

func (s *DescribeConfiguredDestinationIPResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeConfiguredDestinationIPResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeConfiguredDestinationIPResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConfiguredDestinationIPResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeConfiguredDestinationIPResponseBody) SetDestinations(v []*DescribeConfiguredDestinationIPResponseBodyDestinations) *DescribeConfiguredDestinationIPResponseBody {
	s.Destinations = v
	return s
}

func (s *DescribeConfiguredDestinationIPResponseBody) SetPageNo(v int32) *DescribeConfiguredDestinationIPResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeConfiguredDestinationIPResponseBody) SetPageSize(v int32) *DescribeConfiguredDestinationIPResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeConfiguredDestinationIPResponseBody) SetRequestId(v string) *DescribeConfiguredDestinationIPResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfiguredDestinationIPResponseBody) SetTotalCount(v int32) *DescribeConfiguredDestinationIPResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeConfiguredDestinationIPResponseBody) Validate() error {
	if s.Destinations != nil {
		for _, item := range s.Destinations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeConfiguredDestinationIPResponseBodyDestinations struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 1.1.1.1
	DestinationIP  *string `json:"DestinationIP,omitempty" xml:"DestinationIP,omitempty"`
	DestinationISP *string `json:"DestinationISP,omitempty" xml:"DestinationISP,omitempty"`
	// example:
	//
	// cn-shenzhen
	DestinationRegion *string `json:"DestinationRegion,omitempty" xml:"DestinationRegion,omitempty"`
	// example:
	//
	// 1534408189
	OperationTime *int32 `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
}

func (s DescribeConfiguredDestinationIPResponseBodyDestinations) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfiguredDestinationIPResponseBodyDestinations) GoString() string {
	return s.String()
}

func (s *DescribeConfiguredDestinationIPResponseBodyDestinations) GetComment() *string {
	return s.Comment
}

func (s *DescribeConfiguredDestinationIPResponseBodyDestinations) GetDestinationIP() *string {
	return s.DestinationIP
}

func (s *DescribeConfiguredDestinationIPResponseBodyDestinations) GetDestinationISP() *string {
	return s.DestinationISP
}

func (s *DescribeConfiguredDestinationIPResponseBodyDestinations) GetDestinationRegion() *string {
	return s.DestinationRegion
}

func (s *DescribeConfiguredDestinationIPResponseBodyDestinations) GetOperationTime() *int32 {
	return s.OperationTime
}

func (s *DescribeConfiguredDestinationIPResponseBodyDestinations) SetComment(v string) *DescribeConfiguredDestinationIPResponseBodyDestinations {
	s.Comment = &v
	return s
}

func (s *DescribeConfiguredDestinationIPResponseBodyDestinations) SetDestinationIP(v string) *DescribeConfiguredDestinationIPResponseBodyDestinations {
	s.DestinationIP = &v
	return s
}

func (s *DescribeConfiguredDestinationIPResponseBodyDestinations) SetDestinationISP(v string) *DescribeConfiguredDestinationIPResponseBodyDestinations {
	s.DestinationISP = &v
	return s
}

func (s *DescribeConfiguredDestinationIPResponseBodyDestinations) SetDestinationRegion(v string) *DescribeConfiguredDestinationIPResponseBodyDestinations {
	s.DestinationRegion = &v
	return s
}

func (s *DescribeConfiguredDestinationIPResponseBodyDestinations) SetOperationTime(v int32) *DescribeConfiguredDestinationIPResponseBodyDestinations {
	s.OperationTime = &v
	return s
}

func (s *DescribeConfiguredDestinationIPResponseBodyDestinations) Validate() error {
	return dara.Validate(s)
}
