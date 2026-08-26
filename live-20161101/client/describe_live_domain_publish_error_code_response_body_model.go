// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainPublishErrorCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeLiveDomainPublishErrorCodeResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeLiveDomainPublishErrorCodeResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainPublishErrorCodeResponseBody
	GetEndTime() *string
	SetRealTimeCodeData(v []*DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeData) *DescribeLiveDomainPublishErrorCodeResponseBody
	GetRealTimeCodeData() []*DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeData
	SetRequestId(v string) *DescribeLiveDomainPublishErrorCodeResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveDomainPublishErrorCodeResponseBody
	GetStartTime() *string
}

type DescribeLiveDomainPublishErrorCodeResponseBody struct {
	// The time granularity of the queried data. Unit: seconds. Default value: 60.
	//
	// example:
	//
	// 60
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The ingest domain.
	//
	// example:
	//
	// example1.aliyundoc.com,example2.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2016-06-29T09:10:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The status code distribution data for each time interval.
	RealTimeCodeData []*DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeData `json:"RealTimeCodeData,omitempty" xml:"RealTimeCodeData,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BC858082-736F-4A25-867B-E5B67C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2016-06-29T09:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainPublishErrorCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPublishErrorCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBody) GetRealTimeCodeData() []*DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeData {
	return s.RealTimeCodeData
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBody) SetDataInterval(v string) *DescribeLiveDomainPublishErrorCodeResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBody) SetDomainName(v string) *DescribeLiveDomainPublishErrorCodeResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBody) SetEndTime(v string) *DescribeLiveDomainPublishErrorCodeResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBody) SetRealTimeCodeData(v []*DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeData) *DescribeLiveDomainPublishErrorCodeResponseBody {
	s.RealTimeCodeData = v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBody) SetRequestId(v string) *DescribeLiveDomainPublishErrorCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBody) SetStartTime(v string) *DescribeLiveDomainPublishErrorCodeResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBody) Validate() error {
	if s.RealTimeCodeData != nil {
		for _, item := range s.RealTimeCodeData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeData struct {
	// The list of status code distribution data.
	CodeData []*DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeDataCodeData `json:"CodeData,omitempty" xml:"CodeData,omitempty" type:"Repeated"`
	// The start time of the time interval. The time is in the yyyy-MM-ddTHH:mm:ssZ format (UTC).
	//
	// example:
	//
	// 2016-06-29T09:01:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeData) GetCodeData() []*DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeDataCodeData {
	return s.CodeData
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeData) SetCodeData(v []*DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeDataCodeData) *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeData {
	s.CodeData = v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeData) SetTimeStamp(v string) *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeData) Validate() error {
	if s.CodeData != nil {
		for _, item := range s.CodeData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeDataCodeData struct {
	// The response code. Valid values:
	//
	// - 3: read data timeout.
	//
	// - 4: write data error.
	//
	// - 6: write data timeout.
	//
	// - 200: success.
	//
	// - 500: internal system unknown error.
	//
	// - 501: stream ingest exception.
	//
	// - 502: signaling took too long.
	//
	// - 401: stream ingest parameter exception.
	//
	// - 403: stream ingest authentication failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of responses.
	//
	// example:
	//
	// 20
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The proportion of responses.
	//
	// example:
	//
	// 66.04
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeDataCodeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeDataCodeData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeDataCodeData) GetCode() *string {
	return s.Code
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeDataCodeData) GetCount() *string {
	return s.Count
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeDataCodeData) GetProportion() *string {
	return s.Proportion
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeDataCodeData) SetCode(v string) *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeDataCodeData {
	s.Code = &v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeDataCodeData) SetCount(v string) *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeDataCodeData {
	s.Count = &v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeDataCodeData) SetProportion(v string) *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeDataCodeData {
	s.Proportion = &v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeResponseBodyRealTimeCodeDataCodeData) Validate() error {
	return dara.Validate(s)
}
