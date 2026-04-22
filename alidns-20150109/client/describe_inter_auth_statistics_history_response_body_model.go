// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInterAuthStatisticsHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeInterAuthStatisticsHistoryResponseBodyData) *DescribeInterAuthStatisticsHistoryResponseBody
	GetData() []*DescribeInterAuthStatisticsHistoryResponseBodyData
	SetRequestId(v string) *DescribeInterAuthStatisticsHistoryResponseBody
	GetRequestId() *string
}

type DescribeInterAuthStatisticsHistoryResponseBody struct {
	Data []*DescribeInterAuthStatisticsHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInterAuthStatisticsHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsHistoryResponseBody) GetData() []*DescribeInterAuthStatisticsHistoryResponseBodyData {
	return s.Data
}

func (s *DescribeInterAuthStatisticsHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInterAuthStatisticsHistoryResponseBody) SetData(v []*DescribeInterAuthStatisticsHistoryResponseBodyData) *DescribeInterAuthStatisticsHistoryResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryResponseBody) SetRequestId(v string) *DescribeInterAuthStatisticsHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryResponseBody) Validate() error {
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

type DescribeInterAuthStatisticsHistoryResponseBodyData struct {
	// example:
	//
	// 20
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// A
	Qtype *string `json:"Qtype,omitempty" xml:"Qtype,omitempty"`
	// example:
	//
	// 100
	Ratio *int64 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// example:
	//
	// 1706716800000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// mt2.cn
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeInterAuthStatisticsHistoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsHistoryResponseBodyData) GetCount() *int64 {
	return s.Count
}

func (s *DescribeInterAuthStatisticsHistoryResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeInterAuthStatisticsHistoryResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeInterAuthStatisticsHistoryResponseBodyData) GetQtype() *string {
	return s.Qtype
}

func (s *DescribeInterAuthStatisticsHistoryResponseBodyData) GetRatio() *int64 {
	return s.Ratio
}

func (s *DescribeInterAuthStatisticsHistoryResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeInterAuthStatisticsHistoryResponseBodyData) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeInterAuthStatisticsHistoryResponseBodyData) SetCount(v int64) *DescribeInterAuthStatisticsHistoryResponseBodyData {
	s.Count = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryResponseBodyData) SetDomainName(v string) *DescribeInterAuthStatisticsHistoryResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryResponseBodyData) SetProtocol(v string) *DescribeInterAuthStatisticsHistoryResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryResponseBodyData) SetQtype(v string) *DescribeInterAuthStatisticsHistoryResponseBodyData {
	s.Qtype = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryResponseBodyData) SetRatio(v int64) *DescribeInterAuthStatisticsHistoryResponseBodyData {
	s.Ratio = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryResponseBodyData) SetTimestamp(v int64) *DescribeInterAuthStatisticsHistoryResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryResponseBodyData) SetZoneName(v string) *DescribeInterAuthStatisticsHistoryResponseBodyData {
	s.ZoneName = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
