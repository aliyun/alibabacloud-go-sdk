// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRtcRecordUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeLiveRtcRecordUsageDataResponseBody
	GetAppId() *string
	SetAudioSummaryDuration(v float64) *DescribeLiveRtcRecordUsageDataResponseBody
	GetAudioSummaryDuration() *float64
	SetData(v []*DescribeLiveRtcRecordUsageDataResponseBodyData) *DescribeLiveRtcRecordUsageDataResponseBody
	GetData() []*DescribeLiveRtcRecordUsageDataResponseBodyData
	SetRecordMode(v string) *DescribeLiveRtcRecordUsageDataResponseBody
	GetRecordMode() *string
	SetRequestId(v string) *DescribeLiveRtcRecordUsageDataResponseBody
	GetRequestId() *string
	SetTotalSummaryDuration(v float64) *DescribeLiveRtcRecordUsageDataResponseBody
	GetTotalSummaryDuration() *float64
	SetV1080SummaryDuration(v float64) *DescribeLiveRtcRecordUsageDataResponseBody
	GetV1080SummaryDuration() *float64
	SetV480SummaryDuration(v float64) *DescribeLiveRtcRecordUsageDataResponseBody
	GetV480SummaryDuration() *float64
	SetV720SummaryDuration(v float64) *DescribeLiveRtcRecordUsageDataResponseBody
	GetV720SummaryDuration() *float64
}

type DescribeLiveRtcRecordUsageDataResponseBody struct {
	// example:
	//
	// 7fd341b3-c6fa-43e1-96b8-7295a1dd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 20
	AudioSummaryDuration *float64                                          `json:"AudioSummaryDuration,omitempty" xml:"AudioSummaryDuration,omitempty"`
	Data                 []*DescribeLiveRtcRecordUsageDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	RecordMode *string `json:"RecordMode,omitempty" xml:"RecordMode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 16A96B9A-F2*************B92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 150
	TotalSummaryDuration *float64 `json:"TotalSummaryDuration,omitempty" xml:"TotalSummaryDuration,omitempty"`
	// example:
	//
	// 10
	V1080SummaryDuration *float64 `json:"V1080SummaryDuration,omitempty" xml:"V1080SummaryDuration,omitempty"`
	// example:
	//
	// 30
	V480SummaryDuration *float64 `json:"V480SummaryDuration,omitempty" xml:"V480SummaryDuration,omitempty"`
	// example:
	//
	// 40
	V720SummaryDuration *float64 `json:"V720SummaryDuration,omitempty" xml:"V720SummaryDuration,omitempty"`
}

func (s DescribeLiveRtcRecordUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRtcRecordUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) GetAudioSummaryDuration() *float64 {
	return s.AudioSummaryDuration
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) GetData() []*DescribeLiveRtcRecordUsageDataResponseBodyData {
	return s.Data
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) GetRecordMode() *string {
	return s.RecordMode
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) GetTotalSummaryDuration() *float64 {
	return s.TotalSummaryDuration
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) GetV1080SummaryDuration() *float64 {
	return s.V1080SummaryDuration
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) GetV480SummaryDuration() *float64 {
	return s.V480SummaryDuration
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) GetV720SummaryDuration() *float64 {
	return s.V720SummaryDuration
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) SetAppId(v string) *DescribeLiveRtcRecordUsageDataResponseBody {
	s.AppId = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) SetAudioSummaryDuration(v float64) *DescribeLiveRtcRecordUsageDataResponseBody {
	s.AudioSummaryDuration = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) SetData(v []*DescribeLiveRtcRecordUsageDataResponseBodyData) *DescribeLiveRtcRecordUsageDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) SetRecordMode(v string) *DescribeLiveRtcRecordUsageDataResponseBody {
	s.RecordMode = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) SetRequestId(v string) *DescribeLiveRtcRecordUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) SetTotalSummaryDuration(v float64) *DescribeLiveRtcRecordUsageDataResponseBody {
	s.TotalSummaryDuration = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) SetV1080SummaryDuration(v float64) *DescribeLiveRtcRecordUsageDataResponseBody {
	s.V1080SummaryDuration = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) SetV480SummaryDuration(v float64) *DescribeLiveRtcRecordUsageDataResponseBody {
	s.V480SummaryDuration = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) SetV720SummaryDuration(v float64) *DescribeLiveRtcRecordUsageDataResponseBody {
	s.V720SummaryDuration = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponseBody) Validate() error {
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

type DescribeLiveRtcRecordUsageDataResponseBodyData struct {
	// example:
	//
	// 20
	AudioDuration *float64 `json:"AudioDuration,omitempty" xml:"AudioDuration,omitempty"`
	// example:
	//
	// 2022-10-10T20:00:00Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 150
	TotalDuration *float64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// example:
	//
	// 10
	V1080Duration *float64 `json:"V1080Duration,omitempty" xml:"V1080Duration,omitempty"`
	// example:
	//
	// 30
	V480Duration *float64 `json:"V480Duration,omitempty" xml:"V480Duration,omitempty"`
	// example:
	//
	// 40
	V720Duration *float64 `json:"V720Duration,omitempty" xml:"V720Duration,omitempty"`
}

func (s DescribeLiveRtcRecordUsageDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRtcRecordUsageDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLiveRtcRecordUsageDataResponseBodyData) GetAudioDuration() *float64 {
	return s.AudioDuration
}

func (s *DescribeLiveRtcRecordUsageDataResponseBodyData) GetTimestamp() *string {
	return s.Timestamp
}

func (s *DescribeLiveRtcRecordUsageDataResponseBodyData) GetTotalDuration() *float64 {
	return s.TotalDuration
}

func (s *DescribeLiveRtcRecordUsageDataResponseBodyData) GetV1080Duration() *float64 {
	return s.V1080Duration
}

func (s *DescribeLiveRtcRecordUsageDataResponseBodyData) GetV480Duration() *float64 {
	return s.V480Duration
}

func (s *DescribeLiveRtcRecordUsageDataResponseBodyData) GetV720Duration() *float64 {
	return s.V720Duration
}

func (s *DescribeLiveRtcRecordUsageDataResponseBodyData) SetAudioDuration(v float64) *DescribeLiveRtcRecordUsageDataResponseBodyData {
	s.AudioDuration = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponseBodyData) SetTimestamp(v string) *DescribeLiveRtcRecordUsageDataResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponseBodyData) SetTotalDuration(v float64) *DescribeLiveRtcRecordUsageDataResponseBodyData {
	s.TotalDuration = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponseBodyData) SetV1080Duration(v float64) *DescribeLiveRtcRecordUsageDataResponseBodyData {
	s.V1080Duration = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponseBodyData) SetV480Duration(v float64) *DescribeLiveRtcRecordUsageDataResponseBodyData {
	s.V480Duration = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponseBodyData) SetV720Duration(v float64) *DescribeLiveRtcRecordUsageDataResponseBodyData {
	s.V720Duration = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}
