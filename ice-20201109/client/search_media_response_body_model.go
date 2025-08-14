// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchMediaResponseBody
	GetCode() *string
	SetMediaInfoList(v []*SearchMediaResponseBodyMediaInfoList) *SearchMediaResponseBody
	GetMediaInfoList() []*SearchMediaResponseBodyMediaInfoList
	SetRequestId(v string) *SearchMediaResponseBody
	GetRequestId() *string
	SetScrollToken(v string) *SearchMediaResponseBody
	GetScrollToken() *string
	SetSuccess(v string) *SearchMediaResponseBody
	GetSuccess() *string
	SetTotal(v int64) *SearchMediaResponseBody
	GetTotal() *int64
}

type SearchMediaResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The media assets that meet the requirements.
	MediaInfoList []*SearchMediaResponseBodyMediaInfoList `json:"MediaInfoList,omitempty" xml:"MediaInfoList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 6F61C357-ACC0-57FB-876E-D58795335E59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The pagination identifier.
	//
	// example:
	//
	// F8C4F642184DBDA5D93907A70AAE****
	ScrollToken *string `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of media assets that meet the conditions.
	//
	// example:
	//
	// 163
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchMediaResponseBody) GetMediaInfoList() []*SearchMediaResponseBodyMediaInfoList {
	return s.MediaInfoList
}

func (s *SearchMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchMediaResponseBody) GetScrollToken() *string {
	return s.ScrollToken
}

func (s *SearchMediaResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SearchMediaResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *SearchMediaResponseBody) SetCode(v string) *SearchMediaResponseBody {
	s.Code = &v
	return s
}

func (s *SearchMediaResponseBody) SetMediaInfoList(v []*SearchMediaResponseBodyMediaInfoList) *SearchMediaResponseBody {
	s.MediaInfoList = v
	return s
}

func (s *SearchMediaResponseBody) SetRequestId(v string) *SearchMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMediaResponseBody) SetScrollToken(v string) *SearchMediaResponseBody {
	s.ScrollToken = &v
	return s
}

func (s *SearchMediaResponseBody) SetSuccess(v string) *SearchMediaResponseBody {
	s.Success = &v
	return s
}

func (s *SearchMediaResponseBody) SetTotal(v int64) *SearchMediaResponseBody {
	s.Total = &v
	return s
}

func (s *SearchMediaResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaInfoList struct {
	// The details of the intelligent AI job.
	AiData *SearchMediaResponseBodyMediaInfoListAiData `json:"AiData,omitempty" xml:"AiData,omitempty" type:"Struct"`
	// The description of the AI job.
	AiRoughData *SearchMediaResponseBodyMediaInfoListAiRoughData `json:"AiRoughData,omitempty" xml:"AiRoughData,omitempty" type:"Struct"`
	// The information about the files.
	FileInfoList    []*SearchMediaResponseBodyMediaInfoListFileInfoList    `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
	IndexStatusList []*SearchMediaResponseBodyMediaInfoListIndexStatusList `json:"IndexStatusList,omitempty" xml:"IndexStatusList,omitempty" type:"Repeated"`
	// The basic information about the media asset.
	MediaBasicInfo *SearchMediaResponseBodyMediaInfoListMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// The ID of the media asset.
	//
	// example:
	//
	// 3b187b3620c8490886cfc2a9578c3ce6
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s SearchMediaResponseBodyMediaInfoList) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoList) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoList) GetAiData() *SearchMediaResponseBodyMediaInfoListAiData {
	return s.AiData
}

func (s *SearchMediaResponseBodyMediaInfoList) GetAiRoughData() *SearchMediaResponseBodyMediaInfoListAiRoughData {
	return s.AiRoughData
}

func (s *SearchMediaResponseBodyMediaInfoList) GetFileInfoList() []*SearchMediaResponseBodyMediaInfoListFileInfoList {
	return s.FileInfoList
}

func (s *SearchMediaResponseBodyMediaInfoList) GetIndexStatusList() []*SearchMediaResponseBodyMediaInfoListIndexStatusList {
	return s.IndexStatusList
}

func (s *SearchMediaResponseBodyMediaInfoList) GetMediaBasicInfo() *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	return s.MediaBasicInfo
}

func (s *SearchMediaResponseBodyMediaInfoList) GetMediaId() *string {
	return s.MediaId
}

func (s *SearchMediaResponseBodyMediaInfoList) SetAiData(v *SearchMediaResponseBodyMediaInfoListAiData) *SearchMediaResponseBodyMediaInfoList {
	s.AiData = v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoList) SetAiRoughData(v *SearchMediaResponseBodyMediaInfoListAiRoughData) *SearchMediaResponseBodyMediaInfoList {
	s.AiRoughData = v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoList) SetFileInfoList(v []*SearchMediaResponseBodyMediaInfoListFileInfoList) *SearchMediaResponseBodyMediaInfoList {
	s.FileInfoList = v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoList) SetIndexStatusList(v []*SearchMediaResponseBodyMediaInfoListIndexStatusList) *SearchMediaResponseBodyMediaInfoList {
	s.IndexStatusList = v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoList) SetMediaBasicInfo(v *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) *SearchMediaResponseBodyMediaInfoList {
	s.MediaBasicInfo = v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoList) SetMediaId(v string) *SearchMediaResponseBodyMediaInfoList {
	s.MediaId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoList) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaInfoListAiData struct {
	// The tags of the intelligent AI job.
	AiLabelInfo []*SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo `json:"AiLabelInfo,omitempty" xml:"AiLabelInfo,omitempty" type:"Repeated"`
	// The information about audio files.
	AsrInfo []*SearchMediaResponseBodyMediaInfoListAiDataAsrInfo `json:"AsrInfo,omitempty" xml:"AsrInfo,omitempty" type:"Repeated"`
	// The subtitles.
	OcrInfo []*SearchMediaResponseBodyMediaInfoListAiDataOcrInfo `json:"OcrInfo,omitempty" xml:"OcrInfo,omitempty" type:"Repeated"`
}

func (s SearchMediaResponseBodyMediaInfoListAiData) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoListAiData) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoListAiData) GetAiLabelInfo() []*SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo {
	return s.AiLabelInfo
}

func (s *SearchMediaResponseBodyMediaInfoListAiData) GetAsrInfo() []*SearchMediaResponseBodyMediaInfoListAiDataAsrInfo {
	return s.AsrInfo
}

func (s *SearchMediaResponseBodyMediaInfoListAiData) GetOcrInfo() []*SearchMediaResponseBodyMediaInfoListAiDataOcrInfo {
	return s.OcrInfo
}

func (s *SearchMediaResponseBodyMediaInfoListAiData) SetAiLabelInfo(v []*SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) *SearchMediaResponseBodyMediaInfoListAiData {
	s.AiLabelInfo = v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiData) SetAsrInfo(v []*SearchMediaResponseBodyMediaInfoListAiDataAsrInfo) *SearchMediaResponseBodyMediaInfoListAiData {
	s.AsrInfo = v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiData) SetOcrInfo(v []*SearchMediaResponseBodyMediaInfoListAiDataOcrInfo) *SearchMediaResponseBodyMediaInfoListAiData {
	s.OcrInfo = v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiData) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo struct {
	// The category.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The face ID.
	//
	// example:
	//
	// 5FE19530C7A422197535FE74F5DB2B9F
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	// The ID of the entity.
	//
	// example:
	//
	// 10310250338
	LabelId *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	// The name of the entity.
	LabelName *string `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	// The type of the tag.
	LabelType *string `json:"LabelType,omitempty" xml:"LabelType,omitempty"`
	// The clips.
	Occurrences []*SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences `json:"Occurrences,omitempty" xml:"Occurrences,omitempty" type:"Repeated"`
	// The source.
	//
	// example:
	//
	// vision
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) GetCategory() *string {
	return s.Category
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) GetFaceId() *string {
	return s.FaceId
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) GetLabelId() *string {
	return s.LabelId
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) GetLabelName() *string {
	return s.LabelName
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) GetLabelType() *string {
	return s.LabelType
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) GetOccurrences() []*SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences {
	return s.Occurrences
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) GetSource() *string {
	return s.Source
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) SetCategory(v string) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo {
	s.Category = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) SetFaceId(v string) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo {
	s.FaceId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) SetLabelId(v string) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo {
	s.LabelId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) SetLabelName(v string) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo {
	s.LabelName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) SetLabelType(v string) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo {
	s.LabelType = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) SetOccurrences(v []*SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo {
	s.Occurrences = v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) SetSource(v string) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo {
	s.Source = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfo) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences struct {
	// The text content.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The fine-grained ID of the entity.
	//
	// example:
	//
	// 10310250338
	FinegrainId *string `json:"FinegrainId,omitempty" xml:"FinegrainId,omitempty"`
	// The fine-grained name of the entity.
	FinegrainName *string `json:"FinegrainName,omitempty" xml:"FinegrainName,omitempty"`
	// The start time of the clip.
	//
	// example:
	//
	// 1.4
	From *float64 `json:"From,omitempty" xml:"From,omitempty"`
	// The optimal face image encoded in Base64.
	//
	// example:
	//
	// 99C64F6287
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
	// 85010D1
	TableBatchSeqId *string `json:"TableBatchSeqId,omitempty" xml:"TableBatchSeqId,omitempty"`
	// The end time of the clip.
	//
	// example:
	//
	// 2.5
	To *float64 `json:"To,omitempty" xml:"To,omitempty"`
	// The track sequence.
	Tracks []*SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrencesTracks `json:"Tracks,omitempty" xml:"Tracks,omitempty" type:"Repeated"`
	// The ID of the clip.
	//
	// example:
	//
	// 5FE19530C7A422197535FE74F5DB2B9F
	ClipId *string `json:"clipId,omitempty" xml:"clipId,omitempty"`
}

func (s SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) GetContent() *string {
	return s.Content
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) GetFinegrainId() *string {
	return s.FinegrainId
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) GetFinegrainName() *string {
	return s.FinegrainName
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) GetFrom() *float64 {
	return s.From
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) GetImage() *string {
	return s.Image
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) GetScore() *float64 {
	return s.Score
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) GetTableBatchSeqId() *string {
	return s.TableBatchSeqId
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) GetTo() *float64 {
	return s.To
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) GetTracks() []*SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrencesTracks {
	return s.Tracks
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) GetClipId() *string {
	return s.ClipId
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) SetContent(v string) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences {
	s.Content = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) SetFinegrainId(v string) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences {
	s.FinegrainId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) SetFinegrainName(v string) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences {
	s.FinegrainName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) SetFrom(v float64) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences {
	s.From = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) SetImage(v string) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences {
	s.Image = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) SetScore(v float64) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences {
	s.Score = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) SetTableBatchSeqId(v string) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences {
	s.TableBatchSeqId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) SetTo(v float64) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences {
	s.To = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) SetTracks(v []*SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrencesTracks) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences {
	s.Tracks = v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) SetClipId(v string) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences {
	s.ClipId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrences) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrencesTracks struct {
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
	// 50.2
	Size *float64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The timestamp of the track.
	//
	// example:
	//
	// 1.4
	Timestamp *float64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrencesTracks) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrencesTracks) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrencesTracks) GetPosition() *string {
	return s.Position
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrencesTracks) GetSize() *float64 {
	return s.Size
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrencesTracks) GetTimestamp() *float64 {
	return s.Timestamp
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrencesTracks) SetPosition(v string) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrencesTracks {
	s.Position = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrencesTracks) SetSize(v float64) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrencesTracks {
	s.Size = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrencesTracks) SetTimestamp(v float64) *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrencesTracks {
	s.Timestamp = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAiLabelInfoOccurrencesTracks) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaInfoListAiDataAsrInfo struct {
	// The ID of the clip.
	//
	// example:
	//
	// 5FE19530C7A422197535FE74F5DB2B9F
	ClipId *string `json:"ClipId,omitempty" xml:"ClipId,omitempty"`
	// The text content.
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

func (s SearchMediaResponseBodyMediaInfoListAiDataAsrInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoListAiDataAsrInfo) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAsrInfo) GetClipId() *string {
	return s.ClipId
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAsrInfo) GetContent() *string {
	return s.Content
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAsrInfo) GetFrom() *float64 {
	return s.From
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAsrInfo) GetTimestamp() *float64 {
	return s.Timestamp
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAsrInfo) GetTo() *float64 {
	return s.To
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAsrInfo) SetClipId(v string) *SearchMediaResponseBodyMediaInfoListAiDataAsrInfo {
	s.ClipId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAsrInfo) SetContent(v string) *SearchMediaResponseBodyMediaInfoListAiDataAsrInfo {
	s.Content = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAsrInfo) SetFrom(v float64) *SearchMediaResponseBodyMediaInfoListAiDataAsrInfo {
	s.From = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAsrInfo) SetTimestamp(v float64) *SearchMediaResponseBodyMediaInfoListAiDataAsrInfo {
	s.Timestamp = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAsrInfo) SetTo(v float64) *SearchMediaResponseBodyMediaInfoListAiDataAsrInfo {
	s.To = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataAsrInfo) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaInfoListAiDataOcrInfo struct {
	// The ID of the clip.
	//
	// example:
	//
	// 5FE19530C7A422197535FE74F5DB2B9F
	ClipId *string `json:"ClipId,omitempty" xml:"ClipId,omitempty"`
	// The text content.
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

func (s SearchMediaResponseBodyMediaInfoListAiDataOcrInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoListAiDataOcrInfo) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataOcrInfo) GetClipId() *string {
	return s.ClipId
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataOcrInfo) GetContent() *string {
	return s.Content
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataOcrInfo) GetFrom() *float64 {
	return s.From
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataOcrInfo) GetTimestamp() *float64 {
	return s.Timestamp
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataOcrInfo) GetTo() *float64 {
	return s.To
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataOcrInfo) SetClipId(v string) *SearchMediaResponseBodyMediaInfoListAiDataOcrInfo {
	s.ClipId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataOcrInfo) SetContent(v string) *SearchMediaResponseBodyMediaInfoListAiDataOcrInfo {
	s.Content = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataOcrInfo) SetFrom(v float64) *SearchMediaResponseBodyMediaInfoListAiDataOcrInfo {
	s.From = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataOcrInfo) SetTimestamp(v float64) *SearchMediaResponseBodyMediaInfoListAiDataOcrInfo {
	s.Timestamp = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataOcrInfo) SetTo(v float64) *SearchMediaResponseBodyMediaInfoListAiDataOcrInfo {
	s.To = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiDataOcrInfo) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaInfoListAiRoughData struct {
	// TV Series
	//
	// example:
	//
	// TV series
	AiCategory *string `json:"AiCategory,omitempty" xml:"AiCategory,omitempty"`
	// The ID of the AI job.
	//
	// example:
	//
	// cd35b0b0025f71edbfcb472190a9xxxx
	AiJobId *string `json:"AiJobId,omitempty" xml:"AiJobId,omitempty"`
	// The results of the AI job.
	//
	// example:
	//
	// http://xxxx.json
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The save type.
	//
	// example:
	//
	// TEXT
	SaveType *string `json:"SaveType,omitempty" xml:"SaveType,omitempty"`
	// The data status.
	//
	// example:
	//
	// SaveSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SearchMediaResponseBodyMediaInfoListAiRoughData) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoListAiRoughData) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoListAiRoughData) GetAiCategory() *string {
	return s.AiCategory
}

func (s *SearchMediaResponseBodyMediaInfoListAiRoughData) GetAiJobId() *string {
	return s.AiJobId
}

func (s *SearchMediaResponseBodyMediaInfoListAiRoughData) GetResult() *string {
	return s.Result
}

func (s *SearchMediaResponseBodyMediaInfoListAiRoughData) GetSaveType() *string {
	return s.SaveType
}

func (s *SearchMediaResponseBodyMediaInfoListAiRoughData) GetStatus() *string {
	return s.Status
}

func (s *SearchMediaResponseBodyMediaInfoListAiRoughData) SetAiCategory(v string) *SearchMediaResponseBodyMediaInfoListAiRoughData {
	s.AiCategory = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiRoughData) SetAiJobId(v string) *SearchMediaResponseBodyMediaInfoListAiRoughData {
	s.AiJobId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiRoughData) SetResult(v string) *SearchMediaResponseBodyMediaInfoListAiRoughData {
	s.Result = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiRoughData) SetSaveType(v string) *SearchMediaResponseBodyMediaInfoListAiRoughData {
	s.SaveType = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiRoughData) SetStatus(v string) *SearchMediaResponseBodyMediaInfoListAiRoughData {
	s.Status = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListAiRoughData) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaInfoListFileInfoList struct {
	// The basic information about the file, such as the duration and size.
	FileBasicInfo *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
}

func (s SearchMediaResponseBodyMediaInfoListFileInfoList) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoListFileInfoList) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoList) GetFileBasicInfo() *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	return s.FileBasicInfo
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoList) SetFileBasicInfo(v *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) *SearchMediaResponseBodyMediaInfoListFileInfoList {
	s.FileBasicInfo = v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoList) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo struct {
	// The bitrate of the file.
	//
	// example:
	//
	// 1912.13
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The time when the file was created.
	//
	// example:
	//
	// 2022-05-30T02:02:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The duration of the file.
	//
	// example:
	//
	// 60.00000
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// 164265080291300080527050.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The size of the file in bytes.
	//
	// example:
	//
	// 324784
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The status of the file.
	//
	// example:
	//
	// Normal
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// The type of the file.
	//
	// example:
	//
	// source_file
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The Object Storage Service (OSS) URL of the file.
	//
	// example:
	//
	// https://outin-d3f4681ddfd911ec99a600163e1403e7.oss-cn-shanghai.aliyuncs.com/sv/23d5cdd1-18180984899/23d5cdd1-18180984899.mp4
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The encapsulation format of the file.
	//
	// example:
	//
	// mov,mp4,m4a,3gp,3g2,mj2
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// The height of the file.
	//
	// example:
	//
	// 480
	Height      *string `json:"Height,omitempty" xml:"Height,omitempty"`
	ImagesInput *string `json:"ImagesInput,omitempty" xml:"ImagesInput,omitempty"`
	// The time when the file was last modified.
	//
	// example:
	//
	// 2021-12-10T12:19Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The region in which the file is stored.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The width of the file.
	//
	// example:
	//
	// 1920
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetImagesInput() *string {
	return s.ImagesInput
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetBitrate(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetCreateTime(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetDuration(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetFileName(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetFileSize(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetFileStatus(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetFileType(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetFileUrl(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetFormatName(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetHeight(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetImagesInput(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.ImagesInput = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetModifiedTime(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetRegion(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) SetWidth(v string) *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListFileInfoListFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaInfoListIndexStatusList struct {
	// example:
	//
	// Success
	IndexStatus *string `json:"IndexStatus,omitempty" xml:"IndexStatus,omitempty"`
	// example:
	//
	// mm
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
}

func (s SearchMediaResponseBodyMediaInfoListIndexStatusList) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoListIndexStatusList) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoListIndexStatusList) GetIndexStatus() *string {
	return s.IndexStatus
}

func (s *SearchMediaResponseBodyMediaInfoListIndexStatusList) GetIndexType() *string {
	return s.IndexType
}

func (s *SearchMediaResponseBodyMediaInfoListIndexStatusList) SetIndexStatus(v string) *SearchMediaResponseBodyMediaInfoListIndexStatusList {
	s.IndexStatus = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListIndexStatusList) SetIndexType(v string) *SearchMediaResponseBodyMediaInfoListIndexStatusList {
	s.IndexType = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListIndexStatusList) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaInfoListMediaBasicInfo struct {
	// The business to which the media asset belongs.
	//
	// example:
	//
	// IMS
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// The business type of the media asset.
	//
	// example:
	//
	// opening
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The ID of the category.
	//
	// example:
	//
	// 44
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The name of the category.
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The category of the media asset.
	//
	// example:
	//
	// image
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The thumbnail URL of the media asset.
	//
	// example:
	//
	// https://dtlive-bj.oss-cn-beijing.aliyuncs.com/cover/e694372e-4f5b-4821-ae09-efd064f27b63_large_cover_url.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the media asset was created.
	//
	// example:
	//
	// 2020-12-01T19:48Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the media asset was deleted.
	//
	// example:
	//
	// 2020-12-01T19:48Z
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// The description of the media asset.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The address of the media asset that is waiting to be registered.
	//
	// example:
	//
	// oss://clipres/longvideo/material/voice/prod/20220418/07d7c799f6054dc3bbef250854cf84981650248140427
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// 132bd600fc3c71ec99476732a78f6402
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The tags of the media asset.
	//
	// example:
	//
	// tags,tags2
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// The type of the media asset.
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The time when the media asset was modified.
	//
	// example:
	//
	// 2020-12-01T19:48Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The custom ID of the media asset. The ID is a string that contains 6 to 64 characters. Only letters, digits, hyphens (-), and underscores (_) are supported. Each custom ID is unique.
	//
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The snapshots of the media asset.
	//
	// example:
	//
	// [{"bucket":"example-bucket","count":"3","iceJobId":"******f48f0e4154976b2b8c45******","location":"oss-cn-beijing","snapshotRegular":"example.jpg","templateId":"******e6a6440b29eb60bd7c******"}]
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// The source of the media asset.
	//
	// example:
	//
	// oss
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The image sprite of the media asset
	//
	// example:
	//
	// [{"bucket":"example-bucket","count":"32","iceJobId":"******83ec44d58b2069def2e******","location":"oss-cn-shanghai","snapshotRegular":"example/example-{Count}.jpg","spriteRegular":"example/example-{TileCount}.jpg","templateId":"******e438b14ff39293eaec25******","tileCount":"1"}]
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
	// The state of the media asset.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The title of the media asset.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The transcoding status of the media asset.
	//
	// example:
	//
	// Init
	TranscodeStatus *string `json:"TranscodeStatus,omitempty" xml:"TranscodeStatus,omitempty"`
	// The upload source of the media asset.
	//
	// example:
	//
	// general
	UploadSource *string `json:"UploadSource,omitempty" xml:"UploadSource,omitempty"`
	// The user data.
	//
	// example:
	//
	// userData
	UserData          *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VisionDescription *string `json:"VisionDescription,omitempty" xml:"VisionDescription,omitempty"`
}

func (s SearchMediaResponseBodyMediaInfoListMediaBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetBiz() *string {
	return s.Biz
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetCateId() *int64 {
	return s.CateId
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetCateName() *string {
	return s.CateName
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetCategory() *string {
	return s.Category
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetDeletedTime() *string {
	return s.DeletedTime
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetDescription() *string {
	return s.Description
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetInputURL() *string {
	return s.InputURL
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetMediaTags() *string {
	return s.MediaTags
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetMediaType() *string {
	return s.MediaType
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetNamespace() *string {
	return s.Namespace
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetSnapshots() *string {
	return s.Snapshots
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetSource() *string {
	return s.Source
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetSpriteImages() *string {
	return s.SpriteImages
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetStatus() *string {
	return s.Status
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetTitle() *string {
	return s.Title
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetTranscodeStatus() *string {
	return s.TranscodeStatus
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetUploadSource() *string {
	return s.UploadSource
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetUserData() *string {
	return s.UserData
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) GetVisionDescription() *string {
	return s.VisionDescription
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetBiz(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.Biz = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetBusinessType(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetCateId(v int64) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.CateId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetCateName(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.CateName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetCategory(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetCoverURL(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetCreateTime(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetDeletedTime(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetDescription(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetInputURL(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetMediaId(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetMediaTags(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetMediaType(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetModifiedTime(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetNamespace(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.Namespace = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetReferenceId(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.ReferenceId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetSnapshots(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetSource(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetSpriteImages(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetStatus(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetTitle(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetTranscodeStatus(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.TranscodeStatus = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetUploadSource(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.UploadSource = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetUserData(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) SetVisionDescription(v string) *SearchMediaResponseBodyMediaInfoListMediaBasicInfo {
	s.VisionDescription = &v
	return s
}

func (s *SearchMediaResponseBodyMediaInfoListMediaBasicInfo) Validate() error {
	return dara.Validate(s)
}
