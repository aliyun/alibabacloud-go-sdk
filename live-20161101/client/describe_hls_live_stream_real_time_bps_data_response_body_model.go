// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHlsLiveStreamRealTimeBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeHlsLiveStreamRealTimeBpsDataResponseBody
	GetRequestId() *string
	SetTime(v string) *DescribeHlsLiveStreamRealTimeBpsDataResponseBody
	GetTime() *string
	SetUsageData(v []*DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageData) *DescribeHlsLiveStreamRealTimeBpsDataResponseBody
	GetUsageData() []*DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageData
}

type DescribeHlsLiveStreamRealTimeBpsDataResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 17C16B18-D3EA-4809-9CC3-8A2CBE14BC7B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timestamp for which the data was queried.
	//
	// example:
	//
	// 2018-08-08T00:00:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The usage data.
	UsageData []*DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBody) GetTime() *string {
	return s.Time
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBody) GetUsageData() []*DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageData {
	return s.UsageData
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBody) SetRequestId(v string) *DescribeHlsLiveStreamRealTimeBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBody) SetTime(v string) *DescribeHlsLiveStreamRealTimeBpsDataResponseBody {
	s.Time = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBody) SetUsageData(v []*DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageData) *DescribeHlsLiveStreamRealTimeBpsDataResponseBody {
	s.UsageData = v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBody) Validate() error {
	if s.UsageData != nil {
		for _, item := range s.UsageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageData struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Details about the statistics on each HLS stream under the domain name.
	StreamInfos []*DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfos `json:"StreamInfos,omitempty" xml:"StreamInfos,omitempty" type:"Repeated"`
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageData) GoString() string {
	return s.String()
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageData) GetStreamInfos() []*DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfos {
	return s.StreamInfos
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageData) SetDomainName(v string) *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageData {
	s.DomainName = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageData) SetStreamInfos(v []*DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfos) *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageData {
	s.StreamInfos = v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageData) Validate() error {
	if s.StreamInfos != nil {
		for _, item := range s.StreamInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfos struct {
	// The statistics on the HLS stream.
	Infos []*DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfosInfos `json:"Infos,omitempty" xml:"Infos,omitempty" type:"Repeated"`
	// The name of the stream.
	//
	// example:
	//
	// /live/sport.m3u8
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfos) GoString() string {
	return s.String()
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfos) GetInfos() []*DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfosInfos {
	return s.Infos
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfos) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfos) SetInfos(v []*DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfosInfos) *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfos {
	s.Infos = v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfos) SetStreamName(v string) *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfos {
	s.StreamName = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfos) Validate() error {
	if s.Infos != nil {
		for _, item := range s.Infos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfosInfos struct {
	// The bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 11440.88
	DownFlow *float32 `json:"DownFlow,omitempty" xml:"DownFlow,omitempty"`
	// The number of online users.
	//
	// example:
	//
	// 1
	Online *float32 `json:"Online,omitempty" xml:"Online,omitempty"`
	// The bitrate.
	//
	// example:
	//
	// 1028
	Rate *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfosInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfosInfos) GoString() string {
	return s.String()
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfosInfos) GetDownFlow() *float32 {
	return s.DownFlow
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfosInfos) GetOnline() *float32 {
	return s.Online
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfosInfos) GetRate() *string {
	return s.Rate
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfosInfos) SetDownFlow(v float32) *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfosInfos {
	s.DownFlow = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfosInfos) SetOnline(v float32) *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfosInfos {
	s.Online = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfosInfos) SetRate(v string) *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfosInfos {
	s.Rate = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseBodyUsageDataStreamInfosInfos) Validate() error {
	return dara.Validate(s)
}
