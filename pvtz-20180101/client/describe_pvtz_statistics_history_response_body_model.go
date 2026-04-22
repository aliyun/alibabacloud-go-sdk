// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePvtzStatisticsHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribePvtzStatisticsHistoryResponseBodyData) *DescribePvtzStatisticsHistoryResponseBody
	GetData() []*DescribePvtzStatisticsHistoryResponseBodyData
	SetRequestId(v string) *DescribePvtzStatisticsHistoryResponseBody
	GetRequestId() *string
}

type DescribePvtzStatisticsHistoryResponseBody struct {
	Data []*DescribePvtzStatisticsHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 46973D4C-E3E4-4ABA-9190-9A9DE406C7E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePvtzStatisticsHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsHistoryResponseBody) GetData() []*DescribePvtzStatisticsHistoryResponseBodyData {
	return s.Data
}

func (s *DescribePvtzStatisticsHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePvtzStatisticsHistoryResponseBody) SetData(v []*DescribePvtzStatisticsHistoryResponseBodyData) *DescribePvtzStatisticsHistoryResponseBody {
	s.Data = v
	return s
}

func (s *DescribePvtzStatisticsHistoryResponseBody) SetRequestId(v string) *DescribePvtzStatisticsHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryResponseBody) Validate() error {
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

type DescribePvtzStatisticsHistoryResponseBodyData struct {
	// example:
	//
	// 29
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// wget At8P8tza.popscan.xaliyun.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 100
	FailCount *int64 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// example:
	//
	// UDP、TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// A
	Qtype *string `json:"Qtype,omitempty" xml:"Qtype,omitempty"`
	// example:
	//
	// 95
	Ratio *int64 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// example:
	//
	// 5
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// 1687190400000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// cn-zhangjiakou-share.log.aliyuncs.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribePvtzStatisticsHistoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) GetCount() *int64 {
	return s.Count
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) GetFailCount() *int64 {
	return s.FailCount
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) GetQtype() *string {
	return s.Qtype
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) GetRatio() *int64 {
	return s.Ratio
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) SetCount(v int64) *DescribePvtzStatisticsHistoryResponseBodyData {
	s.Count = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) SetDomainName(v string) *DescribePvtzStatisticsHistoryResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) SetFailCount(v int64) *DescribePvtzStatisticsHistoryResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) SetProtocol(v string) *DescribePvtzStatisticsHistoryResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) SetQtype(v string) *DescribePvtzStatisticsHistoryResponseBodyData {
	s.Qtype = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) SetRatio(v int64) *DescribePvtzStatisticsHistoryResponseBodyData {
	s.Ratio = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) SetSuccessCount(v int64) *DescribePvtzStatisticsHistoryResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) SetTimestamp(v int64) *DescribePvtzStatisticsHistoryResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) SetZoneName(v string) *DescribePvtzStatisticsHistoryResponseBodyData {
	s.ZoneName = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
