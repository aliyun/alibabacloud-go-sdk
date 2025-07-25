// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsThreatStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribePdnsThreatStatisticResponseBodyData) *DescribePdnsThreatStatisticResponseBody
	GetData() []*DescribePdnsThreatStatisticResponseBodyData
	SetRequestId(v string) *DescribePdnsThreatStatisticResponseBody
	GetRequestId() *string
}

type DescribePdnsThreatStatisticResponseBody struct {
	Data      []*DescribePdnsThreatStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePdnsThreatStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsThreatStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePdnsThreatStatisticResponseBody) GetData() []*DescribePdnsThreatStatisticResponseBodyData {
	return s.Data
}

func (s *DescribePdnsThreatStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePdnsThreatStatisticResponseBody) SetData(v []*DescribePdnsThreatStatisticResponseBodyData) *DescribePdnsThreatStatisticResponseBody {
	s.Data = v
	return s
}

func (s *DescribePdnsThreatStatisticResponseBody) SetRequestId(v string) *DescribePdnsThreatStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePdnsThreatStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePdnsThreatStatisticResponseBodyData struct {
	DohTotalCount *int64  `json:"DohTotalCount,omitempty" xml:"DohTotalCount,omitempty"`
	ThreatLevel   *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	ThreatType    *string `json:"ThreatType,omitempty" xml:"ThreatType,omitempty"`
	Timestamp     *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TotalCount    *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UdpTotalCount *int64  `json:"UdpTotalCount,omitempty" xml:"UdpTotalCount,omitempty"`
}

func (s DescribePdnsThreatStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsThreatStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePdnsThreatStatisticResponseBodyData) GetDohTotalCount() *int64 {
	return s.DohTotalCount
}

func (s *DescribePdnsThreatStatisticResponseBodyData) GetThreatLevel() *string {
	return s.ThreatLevel
}

func (s *DescribePdnsThreatStatisticResponseBodyData) GetThreatType() *string {
	return s.ThreatType
}

func (s *DescribePdnsThreatStatisticResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribePdnsThreatStatisticResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePdnsThreatStatisticResponseBodyData) GetUdpTotalCount() *int64 {
	return s.UdpTotalCount
}

func (s *DescribePdnsThreatStatisticResponseBodyData) SetDohTotalCount(v int64) *DescribePdnsThreatStatisticResponseBodyData {
	s.DohTotalCount = &v
	return s
}

func (s *DescribePdnsThreatStatisticResponseBodyData) SetThreatLevel(v string) *DescribePdnsThreatStatisticResponseBodyData {
	s.ThreatLevel = &v
	return s
}

func (s *DescribePdnsThreatStatisticResponseBodyData) SetThreatType(v string) *DescribePdnsThreatStatisticResponseBodyData {
	s.ThreatType = &v
	return s
}

func (s *DescribePdnsThreatStatisticResponseBodyData) SetTimestamp(v int64) *DescribePdnsThreatStatisticResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *DescribePdnsThreatStatisticResponseBodyData) SetTotalCount(v int64) *DescribePdnsThreatStatisticResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribePdnsThreatStatisticResponseBodyData) SetUdpTotalCount(v int64) *DescribePdnsThreatStatisticResponseBodyData {
	s.UdpTotalCount = &v
	return s
}

func (s *DescribePdnsThreatStatisticResponseBodyData) Validate() error {
	return dara.Validate(s)
}
