// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaList(v []*SearchMediaResponseBodyMediaList) *SearchMediaResponseBody
	GetMediaList() []*SearchMediaResponseBodyMediaList
	SetRequestId(v string) *SearchMediaResponseBody
	GetRequestId() *string
	SetScrollToken(v string) *SearchMediaResponseBody
	GetScrollToken() *string
	SetTotal(v int64) *SearchMediaResponseBody
	GetTotal() *int64
}

type SearchMediaResponseBody struct {
	// The information about the media assets.
	MediaList []*SearchMediaResponseBodyMediaList `json:"MediaList,omitempty" xml:"MediaList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 3E0CEF83-FB09-4E34-BA1451814B03****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The pagination identifier.
	//
	// example:
	//
	// 24e0fba7188fae707e146esa54****
	ScrollToken *string `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
	// The total number of data records that meet the specified filter criteria.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBody) GetMediaList() []*SearchMediaResponseBodyMediaList {
	return s.MediaList
}

func (s *SearchMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchMediaResponseBody) GetScrollToken() *string {
	return s.ScrollToken
}

func (s *SearchMediaResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *SearchMediaResponseBody) SetMediaList(v []*SearchMediaResponseBodyMediaList) *SearchMediaResponseBody {
	s.MediaList = v
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

func (s *SearchMediaResponseBody) SetTotal(v int64) *SearchMediaResponseBody {
	s.Total = &v
	return s
}

func (s *SearchMediaResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaList struct {
	// Details about AI data.
	AiData *SearchMediaResponseBodyMediaListAiData `json:"AiData,omitempty" xml:"AiData,omitempty" type:"Struct"`
	// The basic information about AI data.
	AiRoughData *SearchMediaResponseBodyMediaListAiRoughData `json:"AiRoughData,omitempty" xml:"AiRoughData,omitempty" type:"Struct"`
	// [The information about the auxiliary media asset](https://help.aliyun.com/document_detail/86991.html).
	AttachedMedia *SearchMediaResponseBodyMediaListAttachedMedia `json:"AttachedMedia,omitempty" xml:"AttachedMedia,omitempty" type:"Struct"`
	// [The information about the audio](https://help.aliyun.com/document_detail/86991.html).
	Audio *SearchMediaResponseBodyMediaListAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	// The time when the media asset was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-07-19T03:45:25Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// [The information about the image](https://help.aliyun.com/document_detail/86991.html).
	Image *SearchMediaResponseBodyMediaListImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// The ID of the file.
	//
	// example:
	//
	// a82a2cd7d4e147bbed6c1ee372****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The type of the media asset. Valid values:
	//
	// 	- **video**
	//
	// 	- **audio**
	//
	// 	- **image**
	//
	// 	- **attached**
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// [The information about the video](https://help.aliyun.com/document_detail/86991.html).
	Video *SearchMediaResponseBodyMediaListVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s SearchMediaResponseBodyMediaList) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaList) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaList) GetAiData() *SearchMediaResponseBodyMediaListAiData {
	return s.AiData
}

func (s *SearchMediaResponseBodyMediaList) GetAiRoughData() *SearchMediaResponseBodyMediaListAiRoughData {
	return s.AiRoughData
}

func (s *SearchMediaResponseBodyMediaList) GetAttachedMedia() *SearchMediaResponseBodyMediaListAttachedMedia {
	return s.AttachedMedia
}

func (s *SearchMediaResponseBodyMediaList) GetAudio() *SearchMediaResponseBodyMediaListAudio {
	return s.Audio
}

func (s *SearchMediaResponseBodyMediaList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SearchMediaResponseBodyMediaList) GetImage() *SearchMediaResponseBodyMediaListImage {
	return s.Image
}

func (s *SearchMediaResponseBodyMediaList) GetMediaId() *string {
	return s.MediaId
}

func (s *SearchMediaResponseBodyMediaList) GetMediaType() *string {
	return s.MediaType
}

func (s *SearchMediaResponseBodyMediaList) GetVideo() *SearchMediaResponseBodyMediaListVideo {
	return s.Video
}

func (s *SearchMediaResponseBodyMediaList) SetAiData(v *SearchMediaResponseBodyMediaListAiData) *SearchMediaResponseBodyMediaList {
	s.AiData = v
	return s
}

func (s *SearchMediaResponseBodyMediaList) SetAiRoughData(v *SearchMediaResponseBodyMediaListAiRoughData) *SearchMediaResponseBodyMediaList {
	s.AiRoughData = v
	return s
}

func (s *SearchMediaResponseBodyMediaList) SetAttachedMedia(v *SearchMediaResponseBodyMediaListAttachedMedia) *SearchMediaResponseBodyMediaList {
	s.AttachedMedia = v
	return s
}

func (s *SearchMediaResponseBodyMediaList) SetAudio(v *SearchMediaResponseBodyMediaListAudio) *SearchMediaResponseBodyMediaList {
	s.Audio = v
	return s
}

func (s *SearchMediaResponseBodyMediaList) SetCreationTime(v string) *SearchMediaResponseBodyMediaList {
	s.CreationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaList) SetImage(v *SearchMediaResponseBodyMediaListImage) *SearchMediaResponseBodyMediaList {
	s.Image = v
	return s
}

func (s *SearchMediaResponseBodyMediaList) SetMediaId(v string) *SearchMediaResponseBodyMediaList {
	s.MediaId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaList) SetMediaType(v string) *SearchMediaResponseBodyMediaList {
	s.MediaType = &v
	return s
}

func (s *SearchMediaResponseBodyMediaList) SetVideo(v *SearchMediaResponseBodyMediaListVideo) *SearchMediaResponseBodyMediaList {
	s.Video = v
	return s
}

func (s *SearchMediaResponseBodyMediaList) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaListAiData struct {
	// The AI tags.
	AiLabelInfo []*SearchMediaResponseBodyMediaListAiDataAiLabelInfo `json:"AiLabelInfo,omitempty" xml:"AiLabelInfo,omitempty" type:"Repeated"`
	// The information about subtitles.
	OcrInfo []*SearchMediaResponseBodyMediaListAiDataOcrInfo `json:"OcrInfo,omitempty" xml:"OcrInfo,omitempty" type:"Repeated"`
}

func (s SearchMediaResponseBodyMediaListAiData) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaListAiData) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaListAiData) GetAiLabelInfo() []*SearchMediaResponseBodyMediaListAiDataAiLabelInfo {
	return s.AiLabelInfo
}

func (s *SearchMediaResponseBodyMediaListAiData) GetOcrInfo() []*SearchMediaResponseBodyMediaListAiDataOcrInfo {
	return s.OcrInfo
}

func (s *SearchMediaResponseBodyMediaListAiData) SetAiLabelInfo(v []*SearchMediaResponseBodyMediaListAiDataAiLabelInfo) *SearchMediaResponseBodyMediaListAiData {
	s.AiLabelInfo = v
	return s
}

func (s *SearchMediaResponseBodyMediaListAiData) SetOcrInfo(v []*SearchMediaResponseBodyMediaListAiDataOcrInfo) *SearchMediaResponseBodyMediaListAiData {
	s.OcrInfo = v
	return s
}

func (s *SearchMediaResponseBodyMediaListAiData) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaListAiDataAiLabelInfo struct {
	// The category.
	//
	// example:
	//
	// Transportation
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The ID of the tag.
	//
	// example:
	//
	// 10310250338
	LabelId *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	// The name of the tag.
	//
	// example:
	//
	// Vehicles
	LabelName *string `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	// The clips.
	Occurrences []*SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences `json:"Occurrences,omitempty" xml:"Occurrences,omitempty" type:"Repeated"`
}

func (s SearchMediaResponseBodyMediaListAiDataAiLabelInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaListAiDataAiLabelInfo) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaListAiDataAiLabelInfo) GetCategory() *string {
	return s.Category
}

func (s *SearchMediaResponseBodyMediaListAiDataAiLabelInfo) GetLabelId() *string {
	return s.LabelId
}

func (s *SearchMediaResponseBodyMediaListAiDataAiLabelInfo) GetLabelName() *string {
	return s.LabelName
}

func (s *SearchMediaResponseBodyMediaListAiDataAiLabelInfo) GetOccurrences() []*SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences {
	return s.Occurrences
}

func (s *SearchMediaResponseBodyMediaListAiDataAiLabelInfo) SetCategory(v string) *SearchMediaResponseBodyMediaListAiDataAiLabelInfo {
	s.Category = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAiDataAiLabelInfo) SetLabelId(v string) *SearchMediaResponseBodyMediaListAiDataAiLabelInfo {
	s.LabelId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAiDataAiLabelInfo) SetLabelName(v string) *SearchMediaResponseBodyMediaListAiDataAiLabelInfo {
	s.LabelName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAiDataAiLabelInfo) SetOccurrences(v []*SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences) *SearchMediaResponseBodyMediaListAiDataAiLabelInfo {
	s.Occurrences = v
	return s
}

func (s *SearchMediaResponseBodyMediaListAiDataAiLabelInfo) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences struct {
	// The start time of the clip.
	//
	// example:
	//
	// 1.4
	From *float64 `json:"From,omitempty" xml:"From,omitempty"`
	// The score.
	//
	// example:
	//
	// 0.75287705
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The end time of the clip.
	//
	// example:
	//
	// 2.5
	To *float64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences) GetFrom() *float64 {
	return s.From
}

func (s *SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences) GetScore() *float64 {
	return s.Score
}

func (s *SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences) GetTo() *float64 {
	return s.To
}

func (s *SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences) SetFrom(v float64) *SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences {
	s.From = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences) SetScore(v float64) *SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences {
	s.Score = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences) SetTo(v float64) *SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences {
	s.To = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaListAiDataOcrInfo struct {
	// The text content.
	//
	// example:
	//
	// I\\"m Jane.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The start time of the subtitle.
	//
	// example:
	//
	// 1.4
	From *float64 `json:"From,omitempty" xml:"From,omitempty"`
	// The end time of the subtitle.
	//
	// example:
	//
	// 2.5
	To *float64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SearchMediaResponseBodyMediaListAiDataOcrInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaListAiDataOcrInfo) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaListAiDataOcrInfo) GetContent() *string {
	return s.Content
}

func (s *SearchMediaResponseBodyMediaListAiDataOcrInfo) GetFrom() *float64 {
	return s.From
}

func (s *SearchMediaResponseBodyMediaListAiDataOcrInfo) GetTo() *float64 {
	return s.To
}

func (s *SearchMediaResponseBodyMediaListAiDataOcrInfo) SetContent(v string) *SearchMediaResponseBodyMediaListAiDataOcrInfo {
	s.Content = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAiDataOcrInfo) SetFrom(v float64) *SearchMediaResponseBodyMediaListAiDataOcrInfo {
	s.From = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAiDataOcrInfo) SetTo(v float64) *SearchMediaResponseBodyMediaListAiDataOcrInfo {
	s.To = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAiDataOcrInfo) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaListAiRoughData struct {
	// The AI category.
	//
	// example:
	//
	// TV series
	AiCategory *string `json:"AiCategory,omitempty" xml:"AiCategory,omitempty"`
	// The ID of the AI task.
	//
	// example:
	//
	// cd35b0b0025f71edbfcb472190a9xxxx
	AiJobId *string `json:"AiJobId,omitempty" xml:"AiJobId,omitempty"`
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

func (s SearchMediaResponseBodyMediaListAiRoughData) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaListAiRoughData) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaListAiRoughData) GetAiCategory() *string {
	return s.AiCategory
}

func (s *SearchMediaResponseBodyMediaListAiRoughData) GetAiJobId() *string {
	return s.AiJobId
}

func (s *SearchMediaResponseBodyMediaListAiRoughData) GetSaveType() *string {
	return s.SaveType
}

func (s *SearchMediaResponseBodyMediaListAiRoughData) GetStatus() *string {
	return s.Status
}

func (s *SearchMediaResponseBodyMediaListAiRoughData) SetAiCategory(v string) *SearchMediaResponseBodyMediaListAiRoughData {
	s.AiCategory = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAiRoughData) SetAiJobId(v string) *SearchMediaResponseBodyMediaListAiRoughData {
	s.AiJobId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAiRoughData) SetSaveType(v string) *SearchMediaResponseBodyMediaListAiRoughData {
	s.SaveType = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAiRoughData) SetStatus(v string) *SearchMediaResponseBodyMediaListAiRoughData {
	s.Status = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAiRoughData) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaListAttachedMedia struct {
	// The ID of the application.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The type of the auxiliary media asset. Valid values:
	//
	// 	- **watermark**
	//
	// 	- **subtitle**
	//
	// 	- **material**
	//
	// example:
	//
	// watermark
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The list of category IDs.
	Categories []*SearchMediaResponseBodyMediaListAttachedMediaCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The time when the auxiliary media asset was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-07-19T03:45:25Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the auxiliary media asset.
	//
	// example:
	//
	// test3
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the auxiliary media asset.
	//
	// example:
	//
	// a82a2cd7d4e147ba0ed6c1ee372****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The time when the auxiliary media asset was updated. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-07-19T03:48:25Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The status of the auxiliary media asset. Valid values:
	//
	// 	- **Uploading**
	//
	// 	- **Normal**
	//
	// 	- **UploadFail**
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The region in which the auxiliary media asset is stored.
	//
	// example:
	//
	// outin-bfefbb90a47c11*****7426.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags of the auxiliary media asset.
	//
	// example:
	//
	// test2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the auxiliary media asset.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URL of the auxiliary media asset.
	//
	// example:
	//
	// https://example.com/****.png
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s SearchMediaResponseBodyMediaListAttachedMedia) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaListAttachedMedia) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) GetAppId() *string {
	return s.AppId
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) GetBusinessType() *string {
	return s.BusinessType
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) GetCategories() []*SearchMediaResponseBodyMediaListAttachedMediaCategories {
	return s.Categories
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) GetDescription() *string {
	return s.Description
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) GetMediaId() *string {
	return s.MediaId
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) GetStatus() *string {
	return s.Status
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) GetTags() *string {
	return s.Tags
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) GetTitle() *string {
	return s.Title
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) GetURL() *string {
	return s.URL
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetAppId(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.AppId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetBusinessType(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.BusinessType = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetCategories(v []*SearchMediaResponseBodyMediaListAttachedMediaCategories) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.Categories = v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetCreationTime(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.CreationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetDescription(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.Description = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetMediaId(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.MediaId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetModificationTime(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.ModificationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetStatus(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.Status = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetStorageLocation(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.StorageLocation = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetTags(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.Tags = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetTitle(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.Title = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetURL(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.URL = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaListAttachedMediaCategories struct {
	// The category ID of the auxiliary media asset.
	//
	// example:
	//
	// 10027394
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// test1
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The level of the category.
	//
	// example:
	//
	// 1
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The ID of the parent node.
	//
	// example:
	//
	// -1
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s SearchMediaResponseBodyMediaListAttachedMediaCategories) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaListAttachedMediaCategories) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaListAttachedMediaCategories) GetCateId() *int64 {
	return s.CateId
}

func (s *SearchMediaResponseBodyMediaListAttachedMediaCategories) GetCateName() *string {
	return s.CateName
}

func (s *SearchMediaResponseBodyMediaListAttachedMediaCategories) GetLevel() *int64 {
	return s.Level
}

func (s *SearchMediaResponseBodyMediaListAttachedMediaCategories) GetParentId() *int64 {
	return s.ParentId
}

func (s *SearchMediaResponseBodyMediaListAttachedMediaCategories) SetCateId(v int64) *SearchMediaResponseBodyMediaListAttachedMediaCategories {
	s.CateId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMediaCategories) SetCateName(v string) *SearchMediaResponseBodyMediaListAttachedMediaCategories {
	s.CateName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMediaCategories) SetLevel(v int64) *SearchMediaResponseBodyMediaListAttachedMediaCategories {
	s.Level = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMediaCategories) SetParentId(v int64) *SearchMediaResponseBodyMediaListAttachedMediaCategories {
	s.ParentId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMediaCategories) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaListAudio struct {
	// The ID of the application.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the audio file.
	//
	// example:
	//
	// a82a2cd7d4e147bbed6c1ee372****
	AudioId *string `json:"AudioId,omitempty" xml:"AudioId,omitempty"`
	// The ID of the category.
	//
	// example:
	//
	// 10000123
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// ceshi
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The URL of the thumbnail.
	//
	// example:
	//
	// http://example.com/image04.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the audio stream was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-07-19T03:45:25Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the audio file.
	//
	// example:
	//
	// audio description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The download switch. The audio file can be downloaded offline only when the download switch is turned on. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	DownloadSwitch *string `json:"DownloadSwitch,omitempty" xml:"DownloadSwitch,omitempty"`
	// The duration of the audio file.
	//
	// example:
	//
	// 123
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The source of the audio file. Valid values:
	//
	// 	- **general**: The audio file is uploaded by using ApsaraVideo VOD.
	//
	// 	- **short_video**: The audio file is uploaded to ApsaraVideo VOD by using the short video SDK. For more information, see [Introduction](https://help.aliyun.com/document_detail/53407.html).
	//
	// 	- **editing**: The audio file is uploaded to ApsaraVideo VOD after online editing and production. For more information, see [ProduceEditingProjectVideo](https://help.aliyun.com/document_detail/68536.html).
	//
	// 	- **live**: The audio file is recorded and uploaded as a file to ApsaraVideo VOD.
	//
	// example:
	//
	// general
	MediaSource *string `json:"MediaSource,omitempty" xml:"MediaSource,omitempty"`
	// The time when the audio file was updated. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-07-19T03:48:25Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The preprocessing status. Only preprocessed videos can be used for live streaming in the production studio. Valid values:
	//
	// 	- **UnPreprocess**
	//
	// 	- **Preprocessing**
	//
	// 	- **PreprocessSucceed**
	//
	// 	- **PreprocessFailed**
	//
	// example:
	//
	// UnPreprocess
	PreprocessStatus *string `json:"PreprocessStatus,omitempty" xml:"PreprocessStatus,omitempty"`
	// The period of time in which the audio file remains in the restored state.
	//
	// example:
	//
	// 2023-03-30T10:14:14Z
	RestoreExpiration *string `json:"RestoreExpiration,omitempty" xml:"RestoreExpiration,omitempty"`
	// The restoration status of the audio file. Valid values:
	//
	// 	- **Processing**
	//
	// 	- **Success**
	//
	// 	- **Failed**
	//
	// example:
	//
	// Success
	RestoreStatus *string `json:"RestoreStatus,omitempty" xml:"RestoreStatus,omitempty"`
	// The size of the audio file.
	//
	// example:
	//
	// 123
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The automatic snapshots.
	Snapshots []*string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// The sprite snapshots.
	SpriteSnapshots []*string `json:"SpriteSnapshots,omitempty" xml:"SpriteSnapshots,omitempty" type:"Repeated"`
	// The status of the audio file. Valid values:
	//
	// 	- **Uploading**
	//
	// 	- **Normal**
	//
	// 	- **UploadFail**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage class of the audio file. Valid values:
	//
	// 	- **Standard**: All media resources are stored as Standard objects.
	//
	// 	- **IA**: All media resources are stored as IA objects.
	//
	// 	- **Archive**: All media resources are stored as Archive objects.
	//
	// 	- **ColdArchive**: All media resources are stored as Cold Archive objects.
	//
	// 	- **SourceIA**: Only the source file is stored as an IA object.
	//
	// 	- **SourceArchive**: Only the source file is stored as an Archive object.
	//
	// 	- **SourceColdArchive**: Only the source file is stored as a Cold Archive object.
	//
	// 	- **Changing**: The storage class is being modified.
	//
	// example:
	//
	// Standard
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The region in which the audio is stored.
	//
	// example:
	//
	// outin-aaa*****aa.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags of the audio file.
	//
	// example:
	//
	// tag1,tag2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the audio file
	//
	// example:
	//
	// audio
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The transcoding mode. Valid values:
	//
	// 	- **FastTranscode**: The audio file is immediately transcoded after it is uploaded. You cannot play the file before it is transcoded.
	//
	// 	- **NoTranscode**: The audio file can be played without being transcoded. You can immediately play the file after it is uploaded.
	//
	// 	- **AsyncTranscode**: The audio file can be immediately played and asynchronously transcoded after it is uploaded.
	//
	// example:
	//
	// FastTranscode
	TranscodeMode *string `json:"TranscodeMode,omitempty" xml:"TranscodeMode,omitempty"`
}

func (s SearchMediaResponseBodyMediaListAudio) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaListAudio) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaListAudio) GetAppId() *string {
	return s.AppId
}

func (s *SearchMediaResponseBodyMediaListAudio) GetAudioId() *string {
	return s.AudioId
}

func (s *SearchMediaResponseBodyMediaListAudio) GetCateId() *int64 {
	return s.CateId
}

func (s *SearchMediaResponseBodyMediaListAudio) GetCateName() *string {
	return s.CateName
}

func (s *SearchMediaResponseBodyMediaListAudio) GetCoverURL() *string {
	return s.CoverURL
}

func (s *SearchMediaResponseBodyMediaListAudio) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SearchMediaResponseBodyMediaListAudio) GetDescription() *string {
	return s.Description
}

func (s *SearchMediaResponseBodyMediaListAudio) GetDownloadSwitch() *string {
	return s.DownloadSwitch
}

func (s *SearchMediaResponseBodyMediaListAudio) GetDuration() *float32 {
	return s.Duration
}

func (s *SearchMediaResponseBodyMediaListAudio) GetMediaSource() *string {
	return s.MediaSource
}

func (s *SearchMediaResponseBodyMediaListAudio) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *SearchMediaResponseBodyMediaListAudio) GetPreprocessStatus() *string {
	return s.PreprocessStatus
}

func (s *SearchMediaResponseBodyMediaListAudio) GetRestoreExpiration() *string {
	return s.RestoreExpiration
}

func (s *SearchMediaResponseBodyMediaListAudio) GetRestoreStatus() *string {
	return s.RestoreStatus
}

func (s *SearchMediaResponseBodyMediaListAudio) GetSize() *int64 {
	return s.Size
}

func (s *SearchMediaResponseBodyMediaListAudio) GetSnapshots() []*string {
	return s.Snapshots
}

func (s *SearchMediaResponseBodyMediaListAudio) GetSpriteSnapshots() []*string {
	return s.SpriteSnapshots
}

func (s *SearchMediaResponseBodyMediaListAudio) GetStatus() *string {
	return s.Status
}

func (s *SearchMediaResponseBodyMediaListAudio) GetStorageClass() *string {
	return s.StorageClass
}

func (s *SearchMediaResponseBodyMediaListAudio) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *SearchMediaResponseBodyMediaListAudio) GetTags() *string {
	return s.Tags
}

func (s *SearchMediaResponseBodyMediaListAudio) GetTitle() *string {
	return s.Title
}

func (s *SearchMediaResponseBodyMediaListAudio) GetTranscodeMode() *string {
	return s.TranscodeMode
}

func (s *SearchMediaResponseBodyMediaListAudio) SetAppId(v string) *SearchMediaResponseBodyMediaListAudio {
	s.AppId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetAudioId(v string) *SearchMediaResponseBodyMediaListAudio {
	s.AudioId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetCateId(v int64) *SearchMediaResponseBodyMediaListAudio {
	s.CateId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetCateName(v string) *SearchMediaResponseBodyMediaListAudio {
	s.CateName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetCoverURL(v string) *SearchMediaResponseBodyMediaListAudio {
	s.CoverURL = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetCreationTime(v string) *SearchMediaResponseBodyMediaListAudio {
	s.CreationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetDescription(v string) *SearchMediaResponseBodyMediaListAudio {
	s.Description = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetDownloadSwitch(v string) *SearchMediaResponseBodyMediaListAudio {
	s.DownloadSwitch = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetDuration(v float32) *SearchMediaResponseBodyMediaListAudio {
	s.Duration = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetMediaSource(v string) *SearchMediaResponseBodyMediaListAudio {
	s.MediaSource = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetModificationTime(v string) *SearchMediaResponseBodyMediaListAudio {
	s.ModificationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetPreprocessStatus(v string) *SearchMediaResponseBodyMediaListAudio {
	s.PreprocessStatus = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetRestoreExpiration(v string) *SearchMediaResponseBodyMediaListAudio {
	s.RestoreExpiration = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetRestoreStatus(v string) *SearchMediaResponseBodyMediaListAudio {
	s.RestoreStatus = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetSize(v int64) *SearchMediaResponseBodyMediaListAudio {
	s.Size = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetSnapshots(v []*string) *SearchMediaResponseBodyMediaListAudio {
	s.Snapshots = v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetSpriteSnapshots(v []*string) *SearchMediaResponseBodyMediaListAudio {
	s.SpriteSnapshots = v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetStatus(v string) *SearchMediaResponseBodyMediaListAudio {
	s.Status = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetStorageClass(v string) *SearchMediaResponseBodyMediaListAudio {
	s.StorageClass = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetStorageLocation(v string) *SearchMediaResponseBodyMediaListAudio {
	s.StorageLocation = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetTags(v string) *SearchMediaResponseBodyMediaListAudio {
	s.Tags = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetTitle(v string) *SearchMediaResponseBodyMediaListAudio {
	s.Title = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetTranscodeMode(v string) *SearchMediaResponseBodyMediaListAudio {
	s.TranscodeMode = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaListImage struct {
	// The ID of the application.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the category.
	//
	// example:
	//
	// 1000123
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// beauty
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The time when the image was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-07-19T03:45:25Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the image file.
	//
	// example:
	//
	// image test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the image file.
	//
	// example:
	//
	// 11130843741se99wqmoes****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The time when the image file was updated. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-07-19T03:48:25Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The status of the image file.
	//
	// 	- **Uploading**
	//
	// 	- **Normal**
	//
	// 	- **UploadFail**
	//
	// example:
	//
	// Uploading
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The region in which the image is stored.
	//
	// example:
	//
	// outin-bfefbb90a47c******163e1c7426.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags of the image file.
	//
	// example:
	//
	// tag1
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the image file.
	//
	// example:
	//
	// image1
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URL of the image file.
	//
	// example:
	//
	// https://example.com/****.png
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s SearchMediaResponseBodyMediaListImage) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaListImage) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaListImage) GetAppId() *string {
	return s.AppId
}

func (s *SearchMediaResponseBodyMediaListImage) GetCateId() *int64 {
	return s.CateId
}

func (s *SearchMediaResponseBodyMediaListImage) GetCateName() *string {
	return s.CateName
}

func (s *SearchMediaResponseBodyMediaListImage) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SearchMediaResponseBodyMediaListImage) GetDescription() *string {
	return s.Description
}

func (s *SearchMediaResponseBodyMediaListImage) GetImageId() *string {
	return s.ImageId
}

func (s *SearchMediaResponseBodyMediaListImage) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *SearchMediaResponseBodyMediaListImage) GetStatus() *string {
	return s.Status
}

func (s *SearchMediaResponseBodyMediaListImage) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *SearchMediaResponseBodyMediaListImage) GetTags() *string {
	return s.Tags
}

func (s *SearchMediaResponseBodyMediaListImage) GetTitle() *string {
	return s.Title
}

func (s *SearchMediaResponseBodyMediaListImage) GetURL() *string {
	return s.URL
}

func (s *SearchMediaResponseBodyMediaListImage) SetAppId(v string) *SearchMediaResponseBodyMediaListImage {
	s.AppId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetCateId(v int64) *SearchMediaResponseBodyMediaListImage {
	s.CateId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetCateName(v string) *SearchMediaResponseBodyMediaListImage {
	s.CateName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetCreationTime(v string) *SearchMediaResponseBodyMediaListImage {
	s.CreationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetDescription(v string) *SearchMediaResponseBodyMediaListImage {
	s.Description = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetImageId(v string) *SearchMediaResponseBodyMediaListImage {
	s.ImageId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetModificationTime(v string) *SearchMediaResponseBodyMediaListImage {
	s.ModificationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetStatus(v string) *SearchMediaResponseBodyMediaListImage {
	s.Status = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetStorageLocation(v string) *SearchMediaResponseBodyMediaListImage {
	s.StorageLocation = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetTags(v string) *SearchMediaResponseBodyMediaListImage {
	s.Tags = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetTitle(v string) *SearchMediaResponseBodyMediaListImage {
	s.Title = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetURL(v string) *SearchMediaResponseBodyMediaListImage {
	s.URL = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) Validate() error {
	return dara.Validate(s)
}

type SearchMediaResponseBodyMediaListVideo struct {
	// The ID of the application.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the category.
	//
	// example:
	//
	// 10000123
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// video1
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The URL of the thumbnail.
	//
	// example:
	//
	// https://example.aliyundoc.com/image01.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the video file was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-07-19T03:45:25Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the video file.
	//
	// example:
	//
	// Video test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The download switch. The video file can be downloaded offline only when the download switch is turned on. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	DownloadSwitch *string `json:"DownloadSwitch,omitempty" xml:"DownloadSwitch,omitempty"`
	// The duration of the video file. Unit: seconds.
	//
	// example:
	//
	// 123
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The source of the video file. Valid values:
	//
	// 	- **general**: The video file is uploaded by using ApsaraVideo VOD.
	//
	// 	- **short_video**: The video file is uploaded by using the short video SDK.
	//
	// 	- **editing**: The video file is produced after online editing.
	//
	// 	- **live**: The video stream is recorded and uploaded as a file.
	//
	// example:
	//
	// general
	MediaSource *string `json:"MediaSource,omitempty" xml:"MediaSource,omitempty"`
	// The time when the video file was updated. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-07-19T03:48:25Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The preprocessing status. Valid values:
	//
	// 	- **UnPreprocess**
	//
	// 	- **Preprocessing**
	//
	// 	- **PreprocessSucceed**
	//
	// 	- **PreprocessFailed**
	//
	// example:
	//
	// Preprocessing
	PreprocessStatus *string `json:"PreprocessStatus,omitempty" xml:"PreprocessStatus,omitempty"`
	// The period of time in which the video file remains in the restored state.
	//
	// example:
	//
	// 2023-03-30T10:14:14Z
	RestoreExpiration *string `json:"RestoreExpiration,omitempty" xml:"RestoreExpiration,omitempty"`
	// The restoration status of the video file. Valid values:
	//
	// 	- **Processing**
	//
	// 	- **Success**
	//
	// 	- **Failed**
	//
	// example:
	//
	// Success
	RestoreStatus *string `json:"RestoreStatus,omitempty" xml:"RestoreStatus,omitempty"`
	// The size of the video file.
	//
	// example:
	//
	// 123
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The automatic snapshots.
	Snapshots []*string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// The sprite snapshots.
	SpriteSnapshots []*string `json:"SpriteSnapshots,omitempty" xml:"SpriteSnapshots,omitempty" type:"Repeated"`
	// The status of the file. Valid values:
	//
	// 	- **Uploading**
	//
	// 	- **UploadFail**
	//
	// 	- **UploadSucc**
	//
	// 	- **Transcoding**
	//
	// 	- **TranscodeFail**
	//
	// 	- **Blocked**
	//
	// 	- **Normal**
	//
	// example:
	//
	// UploadSucc
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage class of the video file. Valid values:
	//
	// 	- **Standard**: All media resources are stored as Standard objects.
	//
	// 	- **IA**: All media resources are stored as IA objects.
	//
	// 	- **Archive**: All media resources are stored as Archive objects.
	//
	// 	- **ColdArchive**: All media resources are stored as Cold Archive objects.
	//
	// 	- **SourceIA**: Only the source file is stored as an IA object.
	//
	// 	- **SourceArchive**: Only the source file is stored as an Archive object.
	//
	// 	- **SourceColdArchive**: Only the source file is stored as a Cold Archive object.
	//
	// 	- **Changing**: The storage class of the video file is being changed.
	//
	// 	- **SourceChanging**: The storage class of the source file is being changed.
	//
	// example:
	//
	// Standard
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The region in which the video is stored.
	//
	// example:
	//
	// outin-bfefbb90a47c******163e1c7426.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags of the video file.
	//
	// example:
	//
	// tag1
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the video.
	//
	// example:
	//
	// ceshi
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The transcoding mode. Valid values:
	//
	// 	- **FastTranscode**: The video file is immediately transcoded after it is uploaded. You cannot play the file before it is transcoded.
	//
	// 	- **NoTranscode**: The video file can be played without being transcoded. You can immediately play the file after it is uploaded.
	//
	// 	- **AsyncTranscode**: The video file can be immediately played and asynchronously transcoded after it is uploaded.
	//
	// example:
	//
	// FastTranscode
	TranscodeMode *string `json:"TranscodeMode,omitempty" xml:"TranscodeMode,omitempty"`
	// The ID of the video file.
	//
	// example:
	//
	// a82a2asdasqadaf3faa0ed6c1ee372****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s SearchMediaResponseBodyMediaListVideo) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponseBodyMediaListVideo) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaListVideo) GetAppId() *string {
	return s.AppId
}

func (s *SearchMediaResponseBodyMediaListVideo) GetCateId() *int64 {
	return s.CateId
}

func (s *SearchMediaResponseBodyMediaListVideo) GetCateName() *string {
	return s.CateName
}

func (s *SearchMediaResponseBodyMediaListVideo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *SearchMediaResponseBodyMediaListVideo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SearchMediaResponseBodyMediaListVideo) GetDescription() *string {
	return s.Description
}

func (s *SearchMediaResponseBodyMediaListVideo) GetDownloadSwitch() *string {
	return s.DownloadSwitch
}

func (s *SearchMediaResponseBodyMediaListVideo) GetDuration() *float32 {
	return s.Duration
}

func (s *SearchMediaResponseBodyMediaListVideo) GetMediaSource() *string {
	return s.MediaSource
}

func (s *SearchMediaResponseBodyMediaListVideo) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *SearchMediaResponseBodyMediaListVideo) GetPreprocessStatus() *string {
	return s.PreprocessStatus
}

func (s *SearchMediaResponseBodyMediaListVideo) GetRestoreExpiration() *string {
	return s.RestoreExpiration
}

func (s *SearchMediaResponseBodyMediaListVideo) GetRestoreStatus() *string {
	return s.RestoreStatus
}

func (s *SearchMediaResponseBodyMediaListVideo) GetSize() *int64 {
	return s.Size
}

func (s *SearchMediaResponseBodyMediaListVideo) GetSnapshots() []*string {
	return s.Snapshots
}

func (s *SearchMediaResponseBodyMediaListVideo) GetSpriteSnapshots() []*string {
	return s.SpriteSnapshots
}

func (s *SearchMediaResponseBodyMediaListVideo) GetStatus() *string {
	return s.Status
}

func (s *SearchMediaResponseBodyMediaListVideo) GetStorageClass() *string {
	return s.StorageClass
}

func (s *SearchMediaResponseBodyMediaListVideo) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *SearchMediaResponseBodyMediaListVideo) GetTags() *string {
	return s.Tags
}

func (s *SearchMediaResponseBodyMediaListVideo) GetTitle() *string {
	return s.Title
}

func (s *SearchMediaResponseBodyMediaListVideo) GetTranscodeMode() *string {
	return s.TranscodeMode
}

func (s *SearchMediaResponseBodyMediaListVideo) GetVideoId() *string {
	return s.VideoId
}

func (s *SearchMediaResponseBodyMediaListVideo) SetAppId(v string) *SearchMediaResponseBodyMediaListVideo {
	s.AppId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetCateId(v int64) *SearchMediaResponseBodyMediaListVideo {
	s.CateId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetCateName(v string) *SearchMediaResponseBodyMediaListVideo {
	s.CateName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetCoverURL(v string) *SearchMediaResponseBodyMediaListVideo {
	s.CoverURL = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetCreationTime(v string) *SearchMediaResponseBodyMediaListVideo {
	s.CreationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetDescription(v string) *SearchMediaResponseBodyMediaListVideo {
	s.Description = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetDownloadSwitch(v string) *SearchMediaResponseBodyMediaListVideo {
	s.DownloadSwitch = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetDuration(v float32) *SearchMediaResponseBodyMediaListVideo {
	s.Duration = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetMediaSource(v string) *SearchMediaResponseBodyMediaListVideo {
	s.MediaSource = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetModificationTime(v string) *SearchMediaResponseBodyMediaListVideo {
	s.ModificationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetPreprocessStatus(v string) *SearchMediaResponseBodyMediaListVideo {
	s.PreprocessStatus = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetRestoreExpiration(v string) *SearchMediaResponseBodyMediaListVideo {
	s.RestoreExpiration = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetRestoreStatus(v string) *SearchMediaResponseBodyMediaListVideo {
	s.RestoreStatus = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetSize(v int64) *SearchMediaResponseBodyMediaListVideo {
	s.Size = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetSnapshots(v []*string) *SearchMediaResponseBodyMediaListVideo {
	s.Snapshots = v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetSpriteSnapshots(v []*string) *SearchMediaResponseBodyMediaListVideo {
	s.SpriteSnapshots = v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetStatus(v string) *SearchMediaResponseBodyMediaListVideo {
	s.Status = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetStorageClass(v string) *SearchMediaResponseBodyMediaListVideo {
	s.StorageClass = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetStorageLocation(v string) *SearchMediaResponseBodyMediaListVideo {
	s.StorageLocation = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetTags(v string) *SearchMediaResponseBodyMediaListVideo {
	s.Tags = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetTitle(v string) *SearchMediaResponseBodyMediaListVideo {
	s.Title = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetTranscodeMode(v string) *SearchMediaResponseBodyMediaListVideo {
	s.TranscodeMode = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetVideoId(v string) *SearchMediaResponseBodyMediaListVideo {
	s.VideoId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) Validate() error {
	return dara.Validate(s)
}
