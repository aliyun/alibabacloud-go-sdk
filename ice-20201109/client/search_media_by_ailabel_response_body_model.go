// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaByAILabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchMediaByAILabelResponseBody
	GetCode() *string
	SetMediaList(v []*SearchMediaByAILabelResponseBodyMediaList) *SearchMediaByAILabelResponseBody
	GetMediaList() []*SearchMediaByAILabelResponseBodyMediaList
	SetRequestId(v string) *SearchMediaByAILabelResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SearchMediaByAILabelResponseBody
	GetSuccess() *string
	SetTotal(v int64) *SearchMediaByAILabelResponseBody
	GetTotal() *int64
}

type SearchMediaByAILabelResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The media assets that contain the specified content.
	MediaList []*SearchMediaByAILabelResponseBodyMediaList `json:"MediaList,omitempty" xml:"MediaList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of audio and video files that meet the conditions.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchMediaByAILabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByAILabelResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMediaByAILabelResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchMediaByAILabelResponseBody) GetMediaList() []*SearchMediaByAILabelResponseBodyMediaList {
	return s.MediaList
}

func (s *SearchMediaByAILabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchMediaByAILabelResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SearchMediaByAILabelResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *SearchMediaByAILabelResponseBody) SetCode(v string) *SearchMediaByAILabelResponseBody {
	s.Code = &v
	return s
}

func (s *SearchMediaByAILabelResponseBody) SetMediaList(v []*SearchMediaByAILabelResponseBodyMediaList) *SearchMediaByAILabelResponseBody {
	s.MediaList = v
	return s
}

func (s *SearchMediaByAILabelResponseBody) SetRequestId(v string) *SearchMediaByAILabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMediaByAILabelResponseBody) SetSuccess(v string) *SearchMediaByAILabelResponseBody {
	s.Success = &v
	return s
}

func (s *SearchMediaByAILabelResponseBody) SetTotal(v int64) *SearchMediaByAILabelResponseBody {
	s.Total = &v
	return s
}

func (s *SearchMediaByAILabelResponseBody) Validate() error {
	if s.MediaList != nil {
		for _, item := range s.MediaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchMediaByAILabelResponseBodyMediaList struct {
	// The details of the AI job.
	AiData *SearchMediaByAILabelResponseBodyMediaListAiData `json:"AiData,omitempty" xml:"AiData,omitempty" type:"Struct"`
	// The ID of the application. Default value: app-1000000.
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The URL of the thumbnail.
	//
	// example:
	//
	// http://example.aliyundoc.com/snapshot/****.jpg?auth_key=1498476426-0-0-f00b9455c49a423ce69cf4e27333****
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// The time when the media asset was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-14T09:15:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the media asset.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The duration. Unit: seconds.
	//
	// example:
	//
	// 12.2
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// 1c6ce34007d571ed94667630a6bc****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The time when the media asset was updated. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-14T09:15:50Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The size of the source file. Unit: bytes.
	//
	// example:
	//
	// 10897890
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The array of video snapshot URLs.
	Snapshots []*string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// The status of the video.
	//
	// Valid values:
	//
	// 	- PrepareFail
	//
	// 	- UploadFail
	//
	// 	- Init
	//
	// 	- UploadSucc
	//
	// 	- Transcoding
	//
	// 	- TranscodeFail
	//
	// 	- Deleted
	//
	// 	- Normal
	//
	// 	- Uploading
	//
	// 	- Preparing
	//
	// 	- Blocked
	//
	// 	- Checking
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage address.
	//
	// example:
	//
	// out-****.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags of the media asset.
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the media asset.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SearchMediaByAILabelResponseBodyMediaList) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByAILabelResponseBodyMediaList) GoString() string {
	return s.String()
}

func (s *SearchMediaByAILabelResponseBodyMediaList) GetAiData() *SearchMediaByAILabelResponseBodyMediaListAiData {
	return s.AiData
}

func (s *SearchMediaByAILabelResponseBodyMediaList) GetAppId() *string {
	return s.AppId
}

func (s *SearchMediaByAILabelResponseBodyMediaList) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *SearchMediaByAILabelResponseBodyMediaList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SearchMediaByAILabelResponseBodyMediaList) GetDescription() *string {
	return s.Description
}

func (s *SearchMediaByAILabelResponseBodyMediaList) GetDuration() *float32 {
	return s.Duration
}

func (s *SearchMediaByAILabelResponseBodyMediaList) GetMediaId() *string {
	return s.MediaId
}

func (s *SearchMediaByAILabelResponseBodyMediaList) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *SearchMediaByAILabelResponseBodyMediaList) GetSize() *int64 {
	return s.Size
}

func (s *SearchMediaByAILabelResponseBodyMediaList) GetSnapshots() []*string {
	return s.Snapshots
}

func (s *SearchMediaByAILabelResponseBodyMediaList) GetStatus() *string {
	return s.Status
}

func (s *SearchMediaByAILabelResponseBodyMediaList) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *SearchMediaByAILabelResponseBodyMediaList) GetTags() *string {
	return s.Tags
}

func (s *SearchMediaByAILabelResponseBodyMediaList) GetTitle() *string {
	return s.Title
}

func (s *SearchMediaByAILabelResponseBodyMediaList) SetAiData(v *SearchMediaByAILabelResponseBodyMediaListAiData) *SearchMediaByAILabelResponseBodyMediaList {
	s.AiData = v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaList) SetAppId(v string) *SearchMediaByAILabelResponseBodyMediaList {
	s.AppId = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaList) SetCoverUrl(v string) *SearchMediaByAILabelResponseBodyMediaList {
	s.CoverUrl = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaList) SetCreationTime(v string) *SearchMediaByAILabelResponseBodyMediaList {
	s.CreationTime = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaList) SetDescription(v string) *SearchMediaByAILabelResponseBodyMediaList {
	s.Description = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaList) SetDuration(v float32) *SearchMediaByAILabelResponseBodyMediaList {
	s.Duration = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaList) SetMediaId(v string) *SearchMediaByAILabelResponseBodyMediaList {
	s.MediaId = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaList) SetModificationTime(v string) *SearchMediaByAILabelResponseBodyMediaList {
	s.ModificationTime = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaList) SetSize(v int64) *SearchMediaByAILabelResponseBodyMediaList {
	s.Size = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaList) SetSnapshots(v []*string) *SearchMediaByAILabelResponseBodyMediaList {
	s.Snapshots = v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaList) SetStatus(v string) *SearchMediaByAILabelResponseBodyMediaList {
	s.Status = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaList) SetStorageLocation(v string) *SearchMediaByAILabelResponseBodyMediaList {
	s.StorageLocation = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaList) SetTags(v string) *SearchMediaByAILabelResponseBodyMediaList {
	s.Tags = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaList) SetTitle(v string) *SearchMediaByAILabelResponseBodyMediaList {
	s.Title = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaList) Validate() error {
	if s.AiData != nil {
		if err := s.AiData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchMediaByAILabelResponseBodyMediaListAiData struct {
	// The tags of the AI job.
	AiLabelInfo []*SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo `json:"AiLabelInfo,omitempty" xml:"AiLabelInfo,omitempty" type:"Repeated"`
	// The information about audio files.
	AsrInfo []*SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo `json:"AsrInfo,omitempty" xml:"AsrInfo,omitempty" type:"Repeated"`
	// The information about subtitle files.
	OcrInfo []*SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo `json:"OcrInfo,omitempty" xml:"OcrInfo,omitempty" type:"Repeated"`
}

func (s SearchMediaByAILabelResponseBodyMediaListAiData) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByAILabelResponseBodyMediaListAiData) GoString() string {
	return s.String()
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiData) GetAiLabelInfo() []*SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo {
	return s.AiLabelInfo
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiData) GetAsrInfo() []*SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo {
	return s.AsrInfo
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiData) GetOcrInfo() []*SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo {
	return s.OcrInfo
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiData) SetAiLabelInfo(v []*SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) *SearchMediaByAILabelResponseBodyMediaListAiData {
	s.AiLabelInfo = v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiData) SetAsrInfo(v []*SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo) *SearchMediaByAILabelResponseBodyMediaListAiData {
	s.AsrInfo = v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiData) SetOcrInfo(v []*SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo) *SearchMediaByAILabelResponseBodyMediaListAiData {
	s.OcrInfo = v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiData) Validate() error {
	if s.AiLabelInfo != nil {
		for _, item := range s.AiLabelInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AsrInfo != nil {
		for _, item := range s.AsrInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OcrInfo != nil {
		for _, item := range s.OcrInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo struct {
	// The category.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The ID of the face.
	//
	// example:
	//
	// 5FE19530C7A422197535FE74F5DB****
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	// The ID of the entity.
	//
	// example:
	//
	// 103102503**
	LabelId *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	// The name of the entity.
	LabelName *string `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	// The type of the tag.
	LabelType *string `json:"LabelType,omitempty" xml:"LabelType,omitempty"`
	// The information about the clips.
	Occurrences []*SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences `json:"Occurrences,omitempty" xml:"Occurrences,omitempty" type:"Repeated"`
	// The source.
	//
	// example:
	//
	// vision
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) GoString() string {
	return s.String()
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) GetCategory() *string {
	return s.Category
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) GetFaceId() *string {
	return s.FaceId
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) GetLabelId() *string {
	return s.LabelId
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) GetLabelName() *string {
	return s.LabelName
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) GetLabelType() *string {
	return s.LabelType
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) GetOccurrences() []*SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences {
	return s.Occurrences
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) GetSource() *string {
	return s.Source
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) SetCategory(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo {
	s.Category = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) SetFaceId(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo {
	s.FaceId = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) SetLabelId(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo {
	s.LabelId = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) SetLabelName(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo {
	s.LabelName = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) SetLabelType(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo {
	s.LabelType = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) SetOccurrences(v []*SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo {
	s.Occurrences = v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) SetSource(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo {
	s.Source = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfo) Validate() error {
	if s.Occurrences != nil {
		for _, item := range s.Occurrences {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences struct {
	// The ID of the clip.
	//
	// example:
	//
	// 158730355E4B82257D8AA1583A58****
	ClipId *string `json:"ClipId,omitempty" xml:"ClipId,omitempty"`
	// The content of the text.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The fine-grained ID of the entity.
	//
	// example:
	//
	// 103102503**
	FinegrainId *string `json:"FinegrainId,omitempty" xml:"FinegrainId,omitempty"`
	// The fine-grained name of the entity.
	FinegrainName *string `json:"FinegrainName,omitempty" xml:"FinegrainName,omitempty"`
	// The start time of the clip.
	//
	// example:
	//
	// 1.4
	From *float64 `json:"From,omitempty" xml:"From,omitempty"`
	// The image that contains the most face information.
	//
	// example:
	//
	// https://service-****-public.oss-cn-hangzhou.aliyuncs.com/1563457****438522/service-image/f788974f-9595-43b2-a478-7c7a1afb****.jpg
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The score.
	//
	// example:
	//
	// 0.75287705
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The sequence ID of the vector table.
	//
	// example:
	//
	// 85010D1**
	TableBatchSeqId *string `json:"TableBatchSeqId,omitempty" xml:"TableBatchSeqId,omitempty"`
	// The end time of the clip.
	//
	// example:
	//
	// 2.5
	To *float64 `json:"To,omitempty" xml:"To,omitempty"`
	// The tracks.
	Tracks []*SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrencesTracks `json:"Tracks,omitempty" xml:"Tracks,omitempty" type:"Repeated"`
}

func (s SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) GoString() string {
	return s.String()
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) GetClipId() *string {
	return s.ClipId
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) GetContent() *string {
	return s.Content
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) GetFinegrainId() *string {
	return s.FinegrainId
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) GetFinegrainName() *string {
	return s.FinegrainName
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) GetFrom() *float64 {
	return s.From
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) GetImage() *string {
	return s.Image
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) GetScore() *float64 {
	return s.Score
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) GetTableBatchSeqId() *string {
	return s.TableBatchSeqId
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) GetTo() *float64 {
	return s.To
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) GetTracks() []*SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrencesTracks {
	return s.Tracks
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) SetClipId(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences {
	s.ClipId = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) SetContent(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences {
	s.Content = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) SetFinegrainId(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences {
	s.FinegrainId = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) SetFinegrainName(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences {
	s.FinegrainName = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) SetFrom(v float64) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences {
	s.From = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) SetImage(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences {
	s.Image = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) SetScore(v float64) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences {
	s.Score = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) SetTableBatchSeqId(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences {
	s.TableBatchSeqId = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) SetTo(v float64) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences {
	s.To = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) SetTracks(v []*SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrencesTracks) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences {
	s.Tracks = v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrences) Validate() error {
	if s.Tracks != nil {
		for _, item := range s.Tracks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrencesTracks struct {
	// The coordinates of the bounding box.
	//
	// example:
	//
	// 468.0;67.0;615.0;267.0
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// The size of the bounding box.
	//
	// example:
	//
	// 50
	Size *float64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The timestamp of the track.
	//
	// example:
	//
	// 1.4
	Timestamp *float64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrencesTracks) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrencesTracks) GoString() string {
	return s.String()
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrencesTracks) GetPosition() *string {
	return s.Position
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrencesTracks) GetSize() *float64 {
	return s.Size
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrencesTracks) GetTimestamp() *float64 {
	return s.Timestamp
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrencesTracks) SetPosition(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrencesTracks {
	s.Position = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrencesTracks) SetSize(v float64) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrencesTracks {
	s.Size = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrencesTracks) SetTimestamp(v float64) *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrencesTracks {
	s.Timestamp = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAiLabelInfoOccurrencesTracks) Validate() error {
	return dara.Validate(s)
}

type SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo struct {
	// The ID of the clip.
	//
	// example:
	//
	// 5FE19530C7A422197535FE74F5DB****
	ClipId *string `json:"ClipId,omitempty" xml:"ClipId,omitempty"`
	// The content of the audio.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The start time of the clip.
	//
	// example:
	//
	// 1.4
	From *float64 `json:"From,omitempty" xml:"From,omitempty"`
	// The timestamp of the clip.
	//
	// example:
	//
	// 1.4
	Timestamp *float64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The end time of the clip.
	//
	// example:
	//
	// 2.5
	To *float64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo) GoString() string {
	return s.String()
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo) GetClipId() *string {
	return s.ClipId
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo) GetContent() *string {
	return s.Content
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo) GetFrom() *float64 {
	return s.From
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo) GetTimestamp() *float64 {
	return s.Timestamp
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo) GetTo() *float64 {
	return s.To
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo) SetClipId(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo {
	s.ClipId = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo) SetContent(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo {
	s.Content = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo) SetFrom(v float64) *SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo {
	s.From = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo) SetTimestamp(v float64) *SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo {
	s.Timestamp = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo) SetTo(v float64) *SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo {
	s.To = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataAsrInfo) Validate() error {
	return dara.Validate(s)
}

type SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo struct {
	// The ID of the clip.
	//
	// example:
	//
	// 5FE19530C7A422197535FE74F5DB****
	ClipId *string `json:"ClipId,omitempty" xml:"ClipId,omitempty"`
	// The content of the text.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The start time of the clip.
	//
	// example:
	//
	// 1.4
	From *float64 `json:"From,omitempty" xml:"From,omitempty"`
	// The timestamp of the clip.
	//
	// example:
	//
	// 1.4
	Timestamp *float64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The end time of the clip.
	//
	// example:
	//
	// 2.5
	To *float64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo) GoString() string {
	return s.String()
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo) GetClipId() *string {
	return s.ClipId
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo) GetContent() *string {
	return s.Content
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo) GetFrom() *float64 {
	return s.From
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo) GetTimestamp() *float64 {
	return s.Timestamp
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo) GetTo() *float64 {
	return s.To
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo) SetClipId(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo {
	s.ClipId = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo) SetContent(v string) *SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo {
	s.Content = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo) SetFrom(v float64) *SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo {
	s.From = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo) SetTimestamp(v float64) *SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo {
	s.Timestamp = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo) SetTo(v float64) *SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo {
	s.To = &v
	return s
}

func (s *SearchMediaByAILabelResponseBodyMediaListAiDataOcrInfo) Validate() error {
	return dara.Validate(s)
}
