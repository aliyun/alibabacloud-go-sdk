// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveOutboundTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeSensitiveOutboundTrendResponseBodyData) *DescribeSensitiveOutboundTrendResponseBody
	GetData() []*DescribeSensitiveOutboundTrendResponseBodyData
	SetRequestId(v string) *DescribeSensitiveOutboundTrendResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeSensitiveOutboundTrendResponseBody
	GetTotalCount() *int64
}

type DescribeSensitiveOutboundTrendResponseBody struct {
	// The information records involved in cross-border data transfer.
	Data []*DescribeSensitiveOutboundTrendResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C1823E96-EF4B-5BD2-9E02-1D18****3ED8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 7
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSensitiveOutboundTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveOutboundTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveOutboundTrendResponseBody) GetData() []*DescribeSensitiveOutboundTrendResponseBodyData {
	return s.Data
}

func (s *DescribeSensitiveOutboundTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSensitiveOutboundTrendResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSensitiveOutboundTrendResponseBody) SetData(v []*DescribeSensitiveOutboundTrendResponseBodyData) *DescribeSensitiveOutboundTrendResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSensitiveOutboundTrendResponseBody) SetRequestId(v string) *DescribeSensitiveOutboundTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSensitiveOutboundTrendResponseBody) SetTotalCount(v int64) *DescribeSensitiveOutboundTrendResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSensitiveOutboundTrendResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSensitiveOutboundTrendResponseBodyData struct {
	// The total number of personal information records.
	//
	// example:
	//
	// 672
	InfoCount *int64 `json:"InfoCount,omitempty" xml:"InfoCount,omitempty"`
	// The total number of personal information records involved in cross-border data transfer.
	//
	// example:
	//
	// 541
	InfoOutboundCount *int64 `json:"InfoOutboundCount,omitempty" xml:"InfoOutboundCount,omitempty"`
	// The total number of sensitive information records involved in cross-border data transfer.
	//
	// example:
	//
	// 378
	SensitiveOutboundCount *int64 `json:"SensitiveOutboundCount,omitempty" xml:"SensitiveOutboundCount,omitempty"`
	// The time of cross-border data transfer. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1672502400
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeSensitiveOutboundTrendResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveOutboundTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveOutboundTrendResponseBodyData) GetInfoCount() *int64 {
	return s.InfoCount
}

func (s *DescribeSensitiveOutboundTrendResponseBodyData) GetInfoOutboundCount() *int64 {
	return s.InfoOutboundCount
}

func (s *DescribeSensitiveOutboundTrendResponseBodyData) GetSensitiveOutboundCount() *int64 {
	return s.SensitiveOutboundCount
}

func (s *DescribeSensitiveOutboundTrendResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeSensitiveOutboundTrendResponseBodyData) SetInfoCount(v int64) *DescribeSensitiveOutboundTrendResponseBodyData {
	s.InfoCount = &v
	return s
}

func (s *DescribeSensitiveOutboundTrendResponseBodyData) SetInfoOutboundCount(v int64) *DescribeSensitiveOutboundTrendResponseBodyData {
	s.InfoOutboundCount = &v
	return s
}

func (s *DescribeSensitiveOutboundTrendResponseBodyData) SetSensitiveOutboundCount(v int64) *DescribeSensitiveOutboundTrendResponseBodyData {
	s.SensitiveOutboundCount = &v
	return s
}

func (s *DescribeSensitiveOutboundTrendResponseBodyData) SetTimestamp(v int64) *DescribeSensitiveOutboundTrendResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *DescribeSensitiveOutboundTrendResponseBodyData) Validate() error {
	return dara.Validate(s)
}
