// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscodeSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExistVideoIds(v []*string) *GetTranscodeSummaryResponseBody
	GetNonExistVideoIds() []*string
	SetRequestId(v string) *GetTranscodeSummaryResponseBody
	GetRequestId() *string
	SetTranscodeSummaryList(v []*GetTranscodeSummaryResponseBodyTranscodeSummaryList) *GetTranscodeSummaryResponseBody
	GetTranscodeSummaryList() []*GetTranscodeSummaryResponseBodyTranscodeSummaryList
}

type GetTranscodeSummaryResponseBody struct {
	// The IDs of audio or video files that do not exist.
	NonExistVideoIds []*string `json:"NonExistVideoIds,omitempty" xml:"NonExistVideoIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The video transcoding summary list of audio or video files.
	TranscodeSummaryList []*GetTranscodeSummaryResponseBodyTranscodeSummaryList `json:"TranscodeSummaryList,omitempty" xml:"TranscodeSummaryList,omitempty" type:"Repeated"`
}

func (s GetTranscodeSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranscodeSummaryResponseBody) GetNonExistVideoIds() []*string {
	return s.NonExistVideoIds
}

func (s *GetTranscodeSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTranscodeSummaryResponseBody) GetTranscodeSummaryList() []*GetTranscodeSummaryResponseBodyTranscodeSummaryList {
	return s.TranscodeSummaryList
}

func (s *GetTranscodeSummaryResponseBody) SetNonExistVideoIds(v []*string) *GetTranscodeSummaryResponseBody {
	s.NonExistVideoIds = v
	return s
}

func (s *GetTranscodeSummaryResponseBody) SetRequestId(v string) *GetTranscodeSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTranscodeSummaryResponseBody) SetTranscodeSummaryList(v []*GetTranscodeSummaryResponseBodyTranscodeSummaryList) *GetTranscodeSummaryResponseBody {
	s.TranscodeSummaryList = v
	return s
}

func (s *GetTranscodeSummaryResponseBody) Validate() error {
	if s.TranscodeSummaryList != nil {
		for _, item := range s.TranscodeSummaryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTranscodeSummaryResponseBodyTranscodeSummaryList struct {
	// The time when the transcoding task was complete. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2019-01-23T12:40:12Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the transcoding task was created. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2019-01-23T12:35:12Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The list of transcoding job summaries.
	TranscodeJobInfoSummaryList []*GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList `json:"TranscodeJobInfoSummaryList,omitempty" xml:"TranscodeJobInfoSummaryList,omitempty" type:"Repeated"`
	// The transcoding status. Valid values:
	//
	// - **Processing**: The transcoding task is in progress.
	//
	// - **Partial**: The transcoding task is partially complete.
	//
	// - **CompleteAllSucc**: All transcoding jobs are complete and successful.
	//
	// - **CompleteAllFail**: All transcoding jobs are complete but all failed. If the source file has issues, no transcoding jobs are initiated and the entire transcoding task fails.
	//
	// - **CompletePartialSucc**: All transcoding jobs are complete but only some are successful.
	//
	// example:
	//
	// Processing
	TranscodeStatus *string `json:"TranscodeStatus,omitempty" xml:"TranscodeStatus,omitempty"`
	// The ID of the transcoding template group used for transcoding.
	//
	// example:
	//
	// 44f9e406bbb*****736a9abe876ecc0
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
	// The audio or video ID.
	//
	// example:
	//
	// e1db68cc586644b83e562bcd94****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetTranscodeSummaryResponseBodyTranscodeSummaryList) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeSummaryResponseBodyTranscodeSummaryList) GoString() string {
	return s.String()
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) GetTranscodeJobInfoSummaryList() []*GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	return s.TranscodeJobInfoSummaryList
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) GetTranscodeStatus() *string {
	return s.TranscodeStatus
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) GetTranscodeTemplateGroupId() *string {
	return s.TranscodeTemplateGroupId
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) GetVideoId() *string {
	return s.VideoId
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) SetCompleteTime(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryList {
	s.CompleteTime = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) SetCreationTime(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryList {
	s.CreationTime = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) SetTranscodeJobInfoSummaryList(v []*GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) *GetTranscodeSummaryResponseBodyTranscodeSummaryList {
	s.TranscodeJobInfoSummaryList = v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) SetTranscodeStatus(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryList {
	s.TranscodeStatus = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) SetTranscodeTemplateGroupId(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryList {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) SetVideoId(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryList {
	s.VideoId = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) Validate() error {
	if s.TranscodeJobInfoSummaryList != nil {
		for _, item := range s.TranscodeJobInfoSummaryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList struct {
	// The average bitrate of the transcoded video output. Unit: Kbps.
	//
	// example:
	//
	// 749
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The time when the transcoding job was complete. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2019-02-27T03:40:51Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the transcoding job was created. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2019-02-27T03:34:46Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The duration of the transcoded video output. Unit: seconds.
	//
	// example:
	//
	// 12
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The error code returned when the transcoding job failed.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the transcoding job failed.
	//
	// example:
	//
	// ErrorMessage
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The file size of the transcoded video output. Unit: bytes.
	//
	// example:
	//
	// 1144259
	Filesize *int64 `json:"Filesize,omitempty" xml:"Filesize,omitempty"`
	// The container format of the transcoded video output.
	//
	// example:
	//
	// mp4
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The frame rate of the transcoded video output. Unit: frames per second.
	//
	// example:
	//
	// 30
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The height of the transcoded video output. Unit: px.
	//
	// example:
	//
	// 960
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The status of the transcoding job. Valid values:
	//
	// - **Transcoding**: The transcoding job is in progress.
	//
	// - **TranscodeSuccess**: The transcoding job is successful.
	//
	// - **TranscodeFail**: The transcoding job failed.
	//
	// example:
	//
	// Transcoding
	TranscodeJobStatus *string `json:"TranscodeJobStatus,omitempty" xml:"TranscodeJobStatus,omitempty"`
	// The transcoding progress. Value range: `[0,100]`.
	//
	// example:
	//
	// 100
	TranscodeProgress *int64 `json:"TranscodeProgress,omitempty" xml:"TranscodeProgress,omitempty"`
	// The ID of the transcoding template used.
	//
	// example:
	//
	// 57496724ae2*****0968d6e08acc8f6
	TranscodeTemplateId *string `json:"TranscodeTemplateId,omitempty" xml:"TranscodeTemplateId,omitempty"`
	// The list of watermarks used for transcoding.
	WatermarkIdList []*string `json:"WatermarkIdList,omitempty" xml:"WatermarkIdList,omitempty" type:"Repeated"`
	// The width of the transcoded video output. Unit: px.
	//
	// example:
	//
	// 544
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GoString() string {
	return s.String()
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GetDuration() *string {
	return s.Duration
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GetFilesize() *int64 {
	return s.Filesize
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GetFormat() *string {
	return s.Format
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GetFps() *string {
	return s.Fps
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GetHeight() *string {
	return s.Height
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GetTranscodeJobStatus() *string {
	return s.TranscodeJobStatus
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GetTranscodeProgress() *int64 {
	return s.TranscodeProgress
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GetTranscodeTemplateId() *string {
	return s.TranscodeTemplateId
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GetWatermarkIdList() []*string {
	return s.WatermarkIdList
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GetWidth() *string {
	return s.Width
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetBitrate(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.Bitrate = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetCompleteTime(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.CompleteTime = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetCreationTime(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.CreationTime = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetDuration(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.Duration = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetErrorCode(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.ErrorCode = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetErrorMessage(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.ErrorMessage = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetFilesize(v int64) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.Filesize = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetFormat(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.Format = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetFps(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.Fps = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetHeight(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.Height = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetTranscodeJobStatus(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.TranscodeJobStatus = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetTranscodeProgress(v int64) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.TranscodeProgress = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetTranscodeTemplateId(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.TranscodeTemplateId = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetWatermarkIdList(v []*string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.WatermarkIdList = v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetWidth(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.Width = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) Validate() error {
	return dara.Validate(s)
}
