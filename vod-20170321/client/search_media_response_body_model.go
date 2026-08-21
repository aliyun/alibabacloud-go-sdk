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
	// The list of media asset information.
	MediaList []*SearchMediaResponseBodyMediaList `json:"MediaList,omitempty" xml:"MediaList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 3E0CEF83-FB09-4E34-BA1451814B03****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// 24e0fba7188fae707e146esa54****
	ScrollToken *string `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
	// The total number of media assets that match the search conditions.
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

type SearchMediaResponseBodyMediaList struct {
	AiData      *SearchMediaResponseBodyMediaListAiData      `json:"AiData,omitempty" xml:"AiData,omitempty" type:"Struct"`
	AiRoughData *SearchMediaResponseBodyMediaListAiRoughData `json:"AiRoughData,omitempty" xml:"AiRoughData,omitempty" type:"Struct"`
	// [Auxiliary media asset information](https://help.aliyun.com/document_detail/86991.html).
	AttachedMedia *SearchMediaResponseBodyMediaListAttachedMedia `json:"AttachedMedia,omitempty" xml:"AttachedMedia,omitempty" type:"Struct"`
	// [Audio information](https://help.aliyun.com/document_detail/86991.html).
	Audio *SearchMediaResponseBodyMediaListAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	// The time when the media asset was created. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2018-07-19T03:45:25Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// [Image information](https://help.aliyun.com/document_detail/86991.html).
	Image *SearchMediaResponseBodyMediaListImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// The media ID.
	//
	// example:
	//
	// a82a2cd7d4e147bbed6c1ee372****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The media type. Valid values:
	//
	// - **video**: video.
	//
	// - **audio**: audio.
	//
	// - **image**: image.
	//
	// - **attached**: auxiliary media asset.
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// [Video information](https://help.aliyun.com/document_detail/86991.html).
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
	if s.AiData != nil {
		if err := s.AiData.Validate(); err != nil {
			return err
		}
	}
	if s.AiRoughData != nil {
		if err := s.AiRoughData.Validate(); err != nil {
			return err
		}
	}
	if s.AttachedMedia != nil {
		if err := s.AttachedMedia.Validate(); err != nil {
			return err
		}
	}
	if s.Audio != nil {
		if err := s.Audio.Validate(); err != nil {
			return err
		}
	}
	if s.Image != nil {
		if err := s.Image.Validate(); err != nil {
			return err
		}
	}
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchMediaResponseBodyMediaListAiData struct {
	AiLabelInfo []*SearchMediaResponseBodyMediaListAiDataAiLabelInfo `json:"AiLabelInfo,omitempty" xml:"AiLabelInfo,omitempty" type:"Repeated"`
	OcrInfo     []*SearchMediaResponseBodyMediaListAiDataOcrInfo     `json:"OcrInfo,omitempty" xml:"OcrInfo,omitempty" type:"Repeated"`
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
	if s.AiLabelInfo != nil {
		for _, item := range s.AiLabelInfo {
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

type SearchMediaResponseBodyMediaListAiDataAiLabelInfo struct {
	Category    *string                                                         `json:"Category,omitempty" xml:"Category,omitempty"`
	LabelId     *string                                                         `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	LabelName   *string                                                         `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
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

type SearchMediaResponseBodyMediaListAiDataAiLabelInfoOccurrences struct {
	From  *float64 `json:"From,omitempty" xml:"From,omitempty"`
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	To    *float64 `json:"To,omitempty" xml:"To,omitempty"`
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
	Content *string  `json:"Content,omitempty" xml:"Content,omitempty"`
	From    *float64 `json:"From,omitempty" xml:"From,omitempty"`
	To      *float64 `json:"To,omitempty" xml:"To,omitempty"`
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
	AiCategory *string `json:"AiCategory,omitempty" xml:"AiCategory,omitempty"`
	AiJobId    *string `json:"AiJobId,omitempty" xml:"AiJobId,omitempty"`
	SaveType   *string `json:"SaveType,omitempty" xml:"SaveType,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// The application ID.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The business type. Valid values:
	//
	// - **watermark**: watermark.
	//
	// - **subtitle**: subtitle.
	//
	// - **material**: material.
	//
	// example:
	//
	// watermark
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The list of category IDs.
	Categories []*SearchMediaResponseBodyMediaListAttachedMediaCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The creation time. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2018-07-19T03:45:25Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description.
	//
	// example:
	//
	// Alibaba Cloud VOD-assisted media asset description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The auxiliary media asset ID.
	//
	// example:
	//
	// a82a2cd7d4e147ba0ed6c1ee372****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The last modification time. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2018-07-19T03:48:25Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The status. Valid values:
	//
	// - **Uploading*	- (uploading): the initial state. The auxiliary media asset is being uploaded.
	//
	// - **Normal*	- (normal): the auxiliary media asset is uploaded.
	//
	// - **UploadFail*	- (failed): the auxiliary media asset failed to be uploaded.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage region.
	//
	// example:
	//
	// outin-bfefbb90a47c11*****7426.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags.
	//
	// example:
	//
	// tag1
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title.
	//
	// example:
	//
	// Alibaba Cloud VOD-assisted media asset Title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The auxiliary media asset URL.
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
	if s.Categories != nil {
		for _, item := range s.Categories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchMediaResponseBodyMediaListAttachedMediaCategories struct {
	// The category ID.
	//
	// example:
	//
	// 10027394
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The category name.
	//
	// example:
	//
	// cate1
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The category level.
	//
	// example:
	//
	// 1
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The parent node ID.
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
	// The application ID.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The audio ID.
	//
	// example:
	//
	// a82a2cd7d4e147bbed6c1ee372****
	AudioId *string `json:"AudioId,omitempty" xml:"AudioId,omitempty"`
	// The category ID.
	//
	// example:
	//
	// 10000123
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The category name.
	//
	// example:
	//
	// cate1
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The thumbnail URL.
	//
	// example:
	//
	// http://example.com/image04.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The creation time. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2018-07-19T03:45:25Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description.
	//
	// example:
	//
	// Alibaba Cloud VOD Audio Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The download switch. Offline download is allowed only when the switch is enabled. Valid values:
	//
	// - **on*	- (enabled): the initial state. Offline download is allowed.
	//
	// - **off*	- (disabled): offline download is disabled.
	//
	// example:
	//
	// on
	DownloadSwitch *string `json:"DownloadSwitch,omitempty" xml:"DownloadSwitch,omitempty"`
	// The duration.
	//
	// example:
	//
	// 123
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The source. Valid values:
	//
	// - **general*	- (ApsaraVideo VOD upload): standard upload.
	//
	// - **short_video*	- (the short video SDK): files uploaded to ApsaraVideo VOD by using the short video SDK. For more information, see [Short video SDK](https://help.aliyun.com/document_detail/53407.html).
	//
	// - **editing*	- (online editing): files uploaded to ApsaraVideo VOD by using online editing. For more information, see [Produce videos](https://help.aliyun.com/document_detail/68536.html).
	//
	// - **live*	- (live recording): files uploaded to ApsaraVideo VOD through live recording.
	//
	// example:
	//
	// general
	MediaSource *string `json:"MediaSource,omitempty" xml:"MediaSource,omitempty"`
	// The last modification time. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2018-07-19T03:48:25Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The preprocessing status. Only preprocessed videos can be used for live streaming. Valid values:
	//
	// - **UnPreprocess**: not preprocessed.
	//
	// - **Preprocessing**: preprocessing.
	//
	// - **PreprocessSucceed**: preprocessing complete.
	//
	// - **PreprocessFailed**: preprocessing failed.
	//
	// example:
	//
	// UnPreprocess
	PreprocessStatus *string `json:"PreprocessStatus,omitempty" xml:"PreprocessStatus,omitempty"`
	// The custom ID. Only lowercase letters, uppercase letters, digits, hyphens, and underscores are supported. The value must be 6 to 64 characters in length and is unique at the user level.
	//
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The expiration time of the media asset restoration.
	//
	// example:
	//
	// 2023-03-30T10:14:14Z
	RestoreExpiration *string `json:"RestoreExpiration,omitempty" xml:"RestoreExpiration,omitempty"`
	// The media asset restoration status. Valid values:
	//
	// - **Processing**: restoring.
	//
	// - **Success**: restoration successful.
	//
	// - **Failed**: restoration failed.
	//
	// example:
	//
	// Success
	RestoreStatus *string `json:"RestoreStatus,omitempty" xml:"RestoreStatus,omitempty"`
	// The size.
	//
	// example:
	//
	// 123
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The list of automatic snapshots.
	Snapshots []*string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// The list of sprites.
	SpriteSnapshots []*string `json:"SpriteSnapshots,omitempty" xml:"SpriteSnapshots,omitempty" type:"Repeated"`
	// The status. Valid values:
	//
	// - **Uploading**: uploading.
	//
	// - **Normal**: normal.
	//
	// - **UploadFail**: upload failed.
	//
	// - **Deleted**: deleted.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage class of the media asset. Valid values:
	//
	// - **Standard**: standard.
	//
	// - **IA**: Infrequent Access (media asset).
	//
	// - **Archive**: Archive (media asset).
	//
	// - **ColdArchive**: Cold Archive (media asset).
	//
	// - **SourceIA**: Infrequent Access (source file).
	//
	// - **SourceArchive**: Archive (source file).
	//
	// - **SourceColdArchive**: Cold Archive (source file).
	//
	// - **Changing**: the storage class is being changed.
	//
	// example:
	//
	// Standard
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The storage region.
	//
	// example:
	//
	// outin-aaa*****aa.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags.
	//
	// example:
	//
	// tag1,tag2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title.
	//
	// example:
	//
	// Alibaba Cloud VOD Audio Title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The transcoding mode. Valid values:
	//
	// - **FastTranscode*	- (standard transcoding, default): transcoding starts after the upload is complete, and the audio can be played only after transcoding is complete.
	//
	// - **NoTranscode*	- (distribution without transcoding): the audio can be played immediately after the upload is complete without transcoding.
	//
	// - **AsyncTranscode*	- (distribution and transcoding upon upload): the audio can be played immediately after the upload is complete, and transcoding is performed asynchronously.
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

func (s *SearchMediaResponseBodyMediaListAudio) GetReferenceId() *string {
	return s.ReferenceId
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

func (s *SearchMediaResponseBodyMediaListAudio) SetReferenceId(v string) *SearchMediaResponseBodyMediaListAudio {
	s.ReferenceId = &v
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
	// The application ID.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The category ID.
	//
	// example:
	//
	// 1000123
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The category name.
	//
	// example:
	//
	// cate1
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The creation time. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2018-07-19T03:45:25Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description.
	//
	// example:
	//
	// Alibaba Cloud VOD Image Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image ID.
	//
	// example:
	//
	// 11130843741se99wqmoes****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The last modification time. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2018-07-19T03:48:25Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The image status. Valid values:
	//
	// - **Uploading*	- (uploading): the initial state. The image is being uploaded.
	//
	// - **Normal*	- (normal): the image is uploaded.
	//
	// - **UploadFail*	- (failed): the image failed to be uploaded.
	//
	// example:
	//
	// Uploading
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage region.
	//
	// example:
	//
	// outin-bfefbb90a47c******163e1c7426.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags.
	//
	// example:
	//
	// tag1
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title.
	//
	// example:
	//
	// Alibaba Cloud VOD Image Title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The image URL.
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
	// The application ID.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The category ID.
	//
	// example:
	//
	// 10000123
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The category name.
	//
	// example:
	//
	// video1
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The thumbnail URL.
	//
	// example:
	//
	// https://example.aliyundoc.com/image01.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the video information was created. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2018-07-19T03:45:25Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The video description.
	//
	// example:
	//
	// Alibaba Cloud VOD video description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The download switch. Offline download is allowed only when the switch is enabled. Valid values:
	//
	// - **on*	- (enabled): the initial state. Offline download is allowed.
	//
	// - **off*	- (disabled): offline download is disabled.
	//
	// example:
	//
	// on
	DownloadSwitch *string `json:"DownloadSwitch,omitempty" xml:"DownloadSwitch,omitempty"`
	// The video duration. Unit: seconds.
	//
	// example:
	//
	// 123
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The source. Valid values:
	//
	// - **general**: ApsaraVideo VOD upload.
	//
	// - **short_video**: the short video SDK.
	//
	// - **editing**: online editing.
	//
	// - **live**: live recording.
	//
	// example:
	//
	// general
	MediaSource *string `json:"MediaSource,omitempty" xml:"MediaSource,omitempty"`
	// The time when the video information was last modified. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2018-07-19T03:48:25Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The preprocessing status. Valid values:
	//
	// - **UnPreprocess**: not preprocessed.
	//
	// - **Preprocessing**: preprocessing.
	//
	// - **PreprocessSucceed**: preprocessing complete.
	//
	// - **PreprocessFailed**: preprocessing failed.
	//
	// example:
	//
	// Preprocessing
	PreprocessStatus *string `json:"PreprocessStatus,omitempty" xml:"PreprocessStatus,omitempty"`
	// The custom ID. Only lowercase letters, uppercase letters, digits, hyphens, and underscores are supported. The value must be 6 to 64 characters in length and is unique at the user level.
	//
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The expiration time of the media asset restoration.
	//
	// example:
	//
	// 2023-03-30T10:14:14Z
	RestoreExpiration *string `json:"RestoreExpiration,omitempty" xml:"RestoreExpiration,omitempty"`
	// The media asset restoration status. Valid values:
	//
	// - **Processing**: restoring.
	//
	// - **Success**: restoration successful.
	//
	// - **Failed**: restoration failed.
	//
	// example:
	//
	// Success
	RestoreStatus *string `json:"RestoreStatus,omitempty" xml:"RestoreStatus,omitempty"`
	// The video size.
	//
	// example:
	//
	// 123
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The list of automatic snapshots.
	Snapshots []*string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// The list of sprites.
	SpriteSnapshots []*string `json:"SpriteSnapshots,omitempty" xml:"SpriteSnapshots,omitempty" type:"Repeated"`
	// The status. Valid values:
	//
	// - **Uploading**: uploading.
	//
	// - **UploadFail**: upload failed.
	//
	// - **UploadSucc**: upload complete.
	//
	// - **Transcoding**: transcoding.
	//
	// - **TranscodeFail**: transcoding failed.
	//
	// - **Blocked**: blocked.
	//
	// - **Normal**: normal.
	//
	// example:
	//
	// UploadSucc
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage class of the media asset. Valid values:
	//
	// - **Standard**: standard.
	//
	// - **IA**: Infrequent Access (media asset).
	//
	// - **Archive**: Archive (media asset).
	//
	// - **ColdArchive**: Cold Archive (media asset).
	//
	// - **SourceIA**: Infrequent Access (source file).
	//
	// - **SourceArchive**: Archive (source file).
	//
	// - **SourceColdArchive**: Cold Archive (source file).
	//
	// - **Changing**: the storage class of the media asset is being changed.
	//
	// - **SourceChanging**: the storage class of the source file is being changed.
	//
	// example:
	//
	// Standard
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The storage region.
	//
	// example:
	//
	// outin-bfefbb90a47c******163e1c7426.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The video tags.
	//
	// example:
	//
	// tag1
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The video title.
	//
	// example:
	//
	// Alibaba Cloud VOD Video Title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The transcoding mode. Valid values:
	//
	// - **FastTranscode*	- (standard transcoding): the default mode. Transcoding starts after the upload is complete, and the video can be played only after transcoding is complete.
	//
	// - **NoTranscode*	- (distribution without transcoding): the video can be played immediately after the upload is complete without transcoding.
	//
	// - **AsyncTranscode*	- (distribution and transcoding upon upload): the video can be played immediately after the upload is complete, and transcoding is performed asynchronously.
	//
	// example:
	//
	// FastTranscode
	TranscodeMode *string `json:"TranscodeMode,omitempty" xml:"TranscodeMode,omitempty"`
	// The video ID.
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

func (s *SearchMediaResponseBodyMediaListVideo) GetReferenceId() *string {
	return s.ReferenceId
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

func (s *SearchMediaResponseBodyMediaListVideo) SetReferenceId(v string) *SearchMediaResponseBodyMediaListVideo {
	s.ReferenceId = &v
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
