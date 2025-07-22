// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudNotesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeCloudNotesResponseBodyItems) *DescribeCloudNotesResponseBody
	GetItems() []*DescribeCloudNotesResponseBodyItems
	SetPageNo(v int32) *DescribeCloudNotesResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeCloudNotesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCloudNotesResponseBody
	GetRequestId() *string
	SetTotalCnt(v int32) *DescribeCloudNotesResponseBody
	GetTotalCnt() *int32
}

type DescribeCloudNotesResponseBody struct {
	Items []*DescribeCloudNotesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 154EF5DE-3D08-1F2C-A482-281F78D74B7C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCnt *int32 `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
}

func (s DescribeCloudNotesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudNotesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudNotesResponseBody) GetItems() []*DescribeCloudNotesResponseBodyItems {
	return s.Items
}

func (s *DescribeCloudNotesResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeCloudNotesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudNotesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudNotesResponseBody) GetTotalCnt() *int32 {
	return s.TotalCnt
}

func (s *DescribeCloudNotesResponseBody) SetItems(v []*DescribeCloudNotesResponseBodyItems) *DescribeCloudNotesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeCloudNotesResponseBody) SetPageNo(v int32) *DescribeCloudNotesResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeCloudNotesResponseBody) SetPageSize(v int32) *DescribeCloudNotesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudNotesResponseBody) SetRequestId(v string) *DescribeCloudNotesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudNotesResponseBody) SetTotalCnt(v int32) *DescribeCloudNotesResponseBody {
	s.TotalCnt = &v
	return s
}

func (s *DescribeCloudNotesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudNotesResponseBodyItems struct {
	// example:
	//
	// cloudNote/ksvxxppi/88_12/autoChapters_1724914365173.json
	AutoChaptersFilePath *string `json:"AutoChaptersFilePath,omitempty" xml:"AutoChaptersFilePath,omitempty"`
	// example:
	//
	// sample-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// testchannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// cloudNote/ksvxxppi/88_12/customPrompt_1724914365173.json
	CustomPromptFilePath *string `json:"CustomPromptFilePath,omitempty" xml:"CustomPromptFilePath,omitempty"`
	// example:
	//
	// cloudNote/ksvxxppi/88_12/meetingAssistance_1724914365173.json
	MeetingAssistanceFilePath *string `json:"MeetingAssistanceFilePath,omitempty" xml:"MeetingAssistanceFilePath,omitempty"`
	// example:
	//
	// 1
	Region *int32 `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// cloudNote/ksvxxppi/88_12/serviceInspection_1724914365173.json
	ServiceInspectionFilePath *string `json:"ServiceInspectionFilePath,omitempty" xml:"ServiceInspectionFilePath,omitempty"`
	// example:
	//
	// 1731939816837
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// example:
	//
	// cloudNote/ksvxxppi/88_12/summarization_1724914365173.json
	SummarizationFilePath *string `json:"SummarizationFilePath,omitempty" xml:"SummarizationFilePath,omitempty"`
	// example:
	//
	// test001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// cloudNote/ksvxxppi/88_12/textPolish_1724914365173.json
	TextPolishFilePath *string `json:"TextPolishFilePath,omitempty" xml:"TextPolishFilePath,omitempty"`
	// example:
	//
	// cloudNote/ksvxxppi/88_12/transcription_1724914365173.json
	TranscriptionFilePath *string `json:"TranscriptionFilePath,omitempty" xml:"TranscriptionFilePath,omitempty"`
	// example:
	//
	// 1
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeCloudNotesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudNotesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeCloudNotesResponseBodyItems) GetAutoChaptersFilePath() *string {
	return s.AutoChaptersFilePath
}

func (s *DescribeCloudNotesResponseBodyItems) GetBucket() *string {
	return s.Bucket
}

func (s *DescribeCloudNotesResponseBodyItems) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeCloudNotesResponseBodyItems) GetCustomPromptFilePath() *string {
	return s.CustomPromptFilePath
}

func (s *DescribeCloudNotesResponseBodyItems) GetMeetingAssistanceFilePath() *string {
	return s.MeetingAssistanceFilePath
}

func (s *DescribeCloudNotesResponseBodyItems) GetRegion() *int32 {
	return s.Region
}

func (s *DescribeCloudNotesResponseBodyItems) GetServiceInspectionFilePath() *string {
	return s.ServiceInspectionFilePath
}

func (s *DescribeCloudNotesResponseBodyItems) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeCloudNotesResponseBodyItems) GetSummarizationFilePath() *string {
	return s.SummarizationFilePath
}

func (s *DescribeCloudNotesResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeCloudNotesResponseBodyItems) GetTextPolishFilePath() *string {
	return s.TextPolishFilePath
}

func (s *DescribeCloudNotesResponseBodyItems) GetTranscriptionFilePath() *string {
	return s.TranscriptionFilePath
}

func (s *DescribeCloudNotesResponseBodyItems) GetVendor() *int32 {
	return s.Vendor
}

func (s *DescribeCloudNotesResponseBodyItems) SetAutoChaptersFilePath(v string) *DescribeCloudNotesResponseBodyItems {
	s.AutoChaptersFilePath = &v
	return s
}

func (s *DescribeCloudNotesResponseBodyItems) SetBucket(v string) *DescribeCloudNotesResponseBodyItems {
	s.Bucket = &v
	return s
}

func (s *DescribeCloudNotesResponseBodyItems) SetChannelId(v string) *DescribeCloudNotesResponseBodyItems {
	s.ChannelId = &v
	return s
}

func (s *DescribeCloudNotesResponseBodyItems) SetCustomPromptFilePath(v string) *DescribeCloudNotesResponseBodyItems {
	s.CustomPromptFilePath = &v
	return s
}

func (s *DescribeCloudNotesResponseBodyItems) SetMeetingAssistanceFilePath(v string) *DescribeCloudNotesResponseBodyItems {
	s.MeetingAssistanceFilePath = &v
	return s
}

func (s *DescribeCloudNotesResponseBodyItems) SetRegion(v int32) *DescribeCloudNotesResponseBodyItems {
	s.Region = &v
	return s
}

func (s *DescribeCloudNotesResponseBodyItems) SetServiceInspectionFilePath(v string) *DescribeCloudNotesResponseBodyItems {
	s.ServiceInspectionFilePath = &v
	return s
}

func (s *DescribeCloudNotesResponseBodyItems) SetStartTs(v int64) *DescribeCloudNotesResponseBodyItems {
	s.StartTs = &v
	return s
}

func (s *DescribeCloudNotesResponseBodyItems) SetSummarizationFilePath(v string) *DescribeCloudNotesResponseBodyItems {
	s.SummarizationFilePath = &v
	return s
}

func (s *DescribeCloudNotesResponseBodyItems) SetTaskId(v string) *DescribeCloudNotesResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeCloudNotesResponseBodyItems) SetTextPolishFilePath(v string) *DescribeCloudNotesResponseBodyItems {
	s.TextPolishFilePath = &v
	return s
}

func (s *DescribeCloudNotesResponseBodyItems) SetTranscriptionFilePath(v string) *DescribeCloudNotesResponseBodyItems {
	s.TranscriptionFilePath = &v
	return s
}

func (s *DescribeCloudNotesResponseBodyItems) SetVendor(v int32) *DescribeCloudNotesResponseBodyItems {
	s.Vendor = &v
	return s
}

func (s *DescribeCloudNotesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
