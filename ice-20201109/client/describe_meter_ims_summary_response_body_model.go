// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterImsSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeMeterImsSummaryResponseBodyData) *DescribeMeterImsSummaryResponseBody
	GetData() []*DescribeMeterImsSummaryResponseBodyData
	SetRequestId(v string) *DescribeMeterImsSummaryResponseBody
	GetRequestId() *string
}

type DescribeMeterImsSummaryResponseBody struct {
	// The usage statistics of IMS.
	Data []*DescribeMeterImsSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BEA98A0C-7870-15FE-B96F-8880BB600A2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterImsSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsSummaryResponseBody) GetData() []*DescribeMeterImsSummaryResponseBodyData {
	return s.Data
}

func (s *DescribeMeterImsSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMeterImsSummaryResponseBody) SetData(v []*DescribeMeterImsSummaryResponseBodyData) *DescribeMeterImsSummaryResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterImsSummaryResponseBody) SetRequestId(v string) *DescribeMeterImsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMeterImsSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMeterImsSummaryResponseBodyData struct {
	// The duration of video editing.
	//
	// example:
	//
	// 8722
	EditingDuration *string `json:"EditingDuration,omitempty" xml:"EditingDuration,omitempty"`
	// The duration of live editing.
	//
	// example:
	//
	// 2000
	LiveEditDuration *string `json:"LiveEditDuration,omitempty" xml:"LiveEditDuration,omitempty"`
	// The duration of live stream recording.
	//
	// example:
	//
	// 100
	LiveRecordDuration *string `json:"LiveRecordDuration,omitempty" xml:"LiveRecordDuration,omitempty"`
	// The number of live stream snapshots.
	//
	// example:
	//
	// 100
	LiveSnapshotCount *string `json:"LiveSnapshotCount,omitempty" xml:"LiveSnapshotCount,omitempty"`
	// The duration of live stream transcoding.
	//
	// example:
	//
	// 12356
	LiveTranscodeDuration *int64 `json:"LiveTranscodeDuration,omitempty" xml:"LiveTranscodeDuration,omitempty"`
	// The duration of AI processing.
	//
	// example:
	//
	// 0
	MpsAiDuration *int64 `json:"MpsAiDuration,omitempty" xml:"MpsAiDuration,omitempty"`
	// The duration of video-on-demand (VOD) transcoding.
	//
	// example:
	//
	// 17337
	MpsTranscodeDuration *int64 `json:"MpsTranscodeDuration,omitempty" xml:"MpsTranscodeDuration,omitempty"`
	// The duration of audio and video enhancement.
	//
	// example:
	//
	// 300
	MpsTranscodeUHDDuration *int64 `json:"MpsTranscodeUHDDuration,omitempty" xml:"MpsTranscodeUHDDuration,omitempty"`
}

func (s DescribeMeterImsSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsSummaryResponseBodyData) GetEditingDuration() *string {
	return s.EditingDuration
}

func (s *DescribeMeterImsSummaryResponseBodyData) GetLiveEditDuration() *string {
	return s.LiveEditDuration
}

func (s *DescribeMeterImsSummaryResponseBodyData) GetLiveRecordDuration() *string {
	return s.LiveRecordDuration
}

func (s *DescribeMeterImsSummaryResponseBodyData) GetLiveSnapshotCount() *string {
	return s.LiveSnapshotCount
}

func (s *DescribeMeterImsSummaryResponseBodyData) GetLiveTranscodeDuration() *int64 {
	return s.LiveTranscodeDuration
}

func (s *DescribeMeterImsSummaryResponseBodyData) GetMpsAiDuration() *int64 {
	return s.MpsAiDuration
}

func (s *DescribeMeterImsSummaryResponseBodyData) GetMpsTranscodeDuration() *int64 {
	return s.MpsTranscodeDuration
}

func (s *DescribeMeterImsSummaryResponseBodyData) GetMpsTranscodeUHDDuration() *int64 {
	return s.MpsTranscodeUHDDuration
}

func (s *DescribeMeterImsSummaryResponseBodyData) SetEditingDuration(v string) *DescribeMeterImsSummaryResponseBodyData {
	s.EditingDuration = &v
	return s
}

func (s *DescribeMeterImsSummaryResponseBodyData) SetLiveEditDuration(v string) *DescribeMeterImsSummaryResponseBodyData {
	s.LiveEditDuration = &v
	return s
}

func (s *DescribeMeterImsSummaryResponseBodyData) SetLiveRecordDuration(v string) *DescribeMeterImsSummaryResponseBodyData {
	s.LiveRecordDuration = &v
	return s
}

func (s *DescribeMeterImsSummaryResponseBodyData) SetLiveSnapshotCount(v string) *DescribeMeterImsSummaryResponseBodyData {
	s.LiveSnapshotCount = &v
	return s
}

func (s *DescribeMeterImsSummaryResponseBodyData) SetLiveTranscodeDuration(v int64) *DescribeMeterImsSummaryResponseBodyData {
	s.LiveTranscodeDuration = &v
	return s
}

func (s *DescribeMeterImsSummaryResponseBodyData) SetMpsAiDuration(v int64) *DescribeMeterImsSummaryResponseBodyData {
	s.MpsAiDuration = &v
	return s
}

func (s *DescribeMeterImsSummaryResponseBodyData) SetMpsTranscodeDuration(v int64) *DescribeMeterImsSummaryResponseBodyData {
	s.MpsTranscodeDuration = &v
	return s
}

func (s *DescribeMeterImsSummaryResponseBodyData) SetMpsTranscodeUHDDuration(v int64) *DescribeMeterImsSummaryResponseBodyData {
	s.MpsTranscodeUHDDuration = &v
	return s
}

func (s *DescribeMeterImsSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
