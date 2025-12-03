// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCondition interface {
	dara.Model
	String() string
	GoString() string
	SetAssKeywordList(v []*string) *SearchCondition
	GetAssKeywordList() []*string
	SetAtAuthorNameList(v []*string) *SearchCondition
	GetAtAuthorNameList() []*string
	SetAuthorNameList(v []*string) *SearchCondition
	GetAuthorNameList() []*string
	SetCommentsLevel(v int32) *SearchCondition
	GetCommentsLevel() *int32
	SetContentLenLevel(v int32) *SearchCondition
	GetContentLenLevel() *int32
	SetCreateTimeEnd(v int64) *SearchCondition
	GetCreateTimeEnd() *int64
	SetCreateTimeStart(v int64) *SearchCondition
	GetCreateTimeStart() *int64
	SetDocContentSign(v string) *SearchCondition
	GetDocContentSign() *string
	SetDocIdList(v []*string) *SearchCondition
	GetDocIdList() []*string
	SetDuplicateRemoval(v bool) *SearchCondition
	GetDuplicateRemoval() *bool
	SetEmotionType(v int32) *SearchCondition
	GetEmotionType() *int32
	SetEnableKeywordHighlight(v bool) *SearchCondition
	GetEnableKeywordHighlight() *bool
	SetExcludeAtAuthorNameList(v []*string) *SearchCondition
	GetExcludeAtAuthorNameList() []*string
	SetExcludeAuthorNameList(v []*string) *SearchCondition
	GetExcludeAuthorNameList() []*string
	SetExcludeHostNameList(v []*string) *SearchCondition
	GetExcludeHostNameList() []*string
	SetExcludeKeywordList(v []*string) *SearchCondition
	GetExcludeKeywordList() []*string
	SetExcludeKeywordListInTitle(v []*string) *SearchCondition
	GetExcludeKeywordListInTitle() []*string
	SetExcludeKeywordTagIds(v []*int64) *SearchCondition
	GetExcludeKeywordTagIds() []*int64
	SetExcludeMaterialTagList(v []*string) *SearchCondition
	GetExcludeMaterialTagList() []*string
	SetExcludeMediaLibraryIdList(v []*string) *SearchCondition
	GetExcludeMediaLibraryIdList() []*string
	SetExcludeMediaNameList(v []*string) *SearchCondition
	GetExcludeMediaNameList() []*string
	SetExcludeMediaTypeList(v []*string) *SearchCondition
	GetExcludeMediaTypeList() []*string
	SetExcludeMessageTypeList(v []*string) *SearchCondition
	GetExcludeMessageTypeList() []*string
	SetFieldConditions(v []*FieldCondition) *SearchCondition
	GetFieldConditions() []*FieldCondition
	SetFilterId(v int64) *SearchCondition
	GetFilterId() *int64
	SetHasAudio(v bool) *SearchCondition
	GetHasAudio() *bool
	SetHasImage(v bool) *SearchCondition
	GetHasImage() *bool
	SetHasMultiModeContent(v bool) *SearchCondition
	GetHasMultiModeContent() *bool
	SetHasVideo(v bool) *SearchCondition
	GetHasVideo() *bool
	SetHostNameList(v []*string) *SearchCondition
	GetHostNameList() []*string
	SetInfluenceLevel(v int64) *SearchCondition
	GetInfluenceLevel() *int64
	SetKeywordTagIds(v []*int64) *SearchCondition
	GetKeywordTagIds() []*int64
	SetLikesLevel(v int32) *SearchCondition
	GetLikesLevel() *int32
	SetMaterialTagList(v []*string) *SearchCondition
	GetMaterialTagList() []*string
	SetMediaLibraryIdList(v []*string) *SearchCondition
	GetMediaLibraryIdList() []*string
	SetMediaNameList(v []*string) *SearchCondition
	GetMediaNameList() []*string
	SetMediaTypeList(v []*string) *SearchCondition
	GetMediaTypeList() []*string
	SetMessageTypeList(v []*string) *SearchCondition
	GetMessageTypeList() []*string
	SetPageNow(v int32) *SearchCondition
	GetPageNow() *int32
	SetPageSize(v int32) *SearchCondition
	GetPageSize() *int32
	SetParentDocId(v string) *SearchCondition
	GetParentDocId() *string
	SetPosKeywordList(v []*string) *SearchCondition
	GetPosKeywordList() []*string
	SetPosKeywordListInTitle(v []*string) *SearchCondition
	GetPosKeywordListInTitle() []*string
	SetProjectId(v int64) *SearchCondition
	GetProjectId() *int64
	SetPropagationLevel(v int64) *SearchCondition
	GetPropagationLevel() *int64
	SetPublishTimeEnd(v int64) *SearchCondition
	GetPublishTimeEnd() *int64
	SetPublishTimeStart(v int64) *SearchCondition
	GetPublishTimeStart() *int64
	SetReadsLevel(v int32) *SearchCondition
	GetReadsLevel() *int32
	SetRelevanceLevel(v int32) *SearchCondition
	GetRelevanceLevel() *int32
	SetRepostLevel(v int32) *SearchCondition
	GetRepostLevel() *int32
	SetSortBy(v string) *SearchCondition
	GetSortBy() *string
	SetSortByDirection(v string) *SearchCondition
	GetSortByDirection() *string
	SetTopicList(v []*string) *SearchCondition
	GetTopicList() []*string
	SetUpdateTimeEnd(v int64) *SearchCondition
	GetUpdateTimeEnd() *int64
	SetUpdateTimeStart(v int64) *SearchCondition
	GetUpdateTimeStart() *int64
}

type SearchCondition struct {
	// example:
	//
	// ["A&&B","C"]
	AssKeywordList []*string `json:"assKeywordList,omitempty" xml:"assKeywordList,omitempty" type:"Repeated"`
	// example:
	//
	// xxx
	AtAuthorNameList []*string `json:"atAuthorNameList,omitempty" xml:"atAuthorNameList,omitempty" type:"Repeated"`
	// example:
	//
	// xxx
	AuthorNameList []*string `json:"authorNameList,omitempty" xml:"authorNameList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CommentsLevel *int32 `json:"commentsLevel,omitempty" xml:"commentsLevel,omitempty"`
	// example:
	//
	// 1
	ContentLenLevel *int32 `json:"contentLenLevel,omitempty" xml:"contentLenLevel,omitempty"`
	// example:
	//
	// 1620452881429
	CreateTimeEnd *int64 `json:"createTimeEnd,omitempty" xml:"createTimeEnd,omitempty"`
	// example:
	//
	// 1610452881429
	CreateTimeStart *int64 `json:"createTimeStart,omitempty" xml:"createTimeStart,omitempty"`
	// example:
	//
	// 3478278371214
	DocContentSign *string `json:"docContentSign,omitempty" xml:"docContentSign,omitempty"`
	// example:
	//
	// 5747368272834931
	DocIdList []*string `json:"docIdList,omitempty" xml:"docIdList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	DuplicateRemoval *bool `json:"duplicateRemoval,omitempty" xml:"duplicateRemoval,omitempty"`
	// example:
	//
	// 1
	EmotionType *int32 `json:"emotionType,omitempty" xml:"emotionType,omitempty"`
	// example:
	//
	// true
	EnableKeywordHighlight *bool `json:"enableKeywordHighlight,omitempty" xml:"enableKeywordHighlight,omitempty"`
	// example:
	//
	// xxx
	ExcludeAtAuthorNameList []*string `json:"excludeAtAuthorNameList,omitempty" xml:"excludeAtAuthorNameList,omitempty" type:"Repeated"`
	// example:
	//
	// 新浪财经
	ExcludeAuthorNameList []*string `json:"excludeAuthorNameList,omitempty" xml:"excludeAuthorNameList,omitempty" type:"Repeated"`
	// example:
	//
	// finance.sina.com.cn
	ExcludeHostNameList []*string `json:"excludeHostNameList,omitempty" xml:"excludeHostNameList,omitempty" type:"Repeated"`
	// example:
	//
	// ["A&&B","C"]
	ExcludeKeywordList []*string `json:"excludeKeywordList,omitempty" xml:"excludeKeywordList,omitempty" type:"Repeated"`
	// example:
	//
	// ["A&&B","C"]
	ExcludeKeywordListInTitle []*string `json:"excludeKeywordListInTitle,omitempty" xml:"excludeKeywordListInTitle,omitempty" type:"Repeated"`
	// example:
	//
	// 123,456
	ExcludeKeywordTagIds   []*int64  `json:"excludeKeywordTagIds,omitempty" xml:"excludeKeywordTagIds,omitempty" type:"Repeated"`
	ExcludeMaterialTagList []*string `json:"excludeMaterialTagList,omitempty" xml:"excludeMaterialTagList,omitempty" type:"Repeated"`
	// example:
	//
	// 123,456
	ExcludeMediaLibraryIdList []*string `json:"excludeMediaLibraryIdList,omitempty" xml:"excludeMediaLibraryIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 新浪财经
	ExcludeMediaNameList []*string `json:"excludeMediaNameList,omitempty" xml:"excludeMediaNameList,omitempty" type:"Repeated"`
	// example:
	//
	// WEIBO-REPOST_WEIBO
	ExcludeMediaTypeList []*string `json:"excludeMediaTypeList,omitempty" xml:"excludeMediaTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// COMMENT
	ExcludeMessageTypeList []*string         `json:"excludeMessageTypeList,omitempty" xml:"excludeMessageTypeList,omitempty" type:"Repeated"`
	FieldConditions        []*FieldCondition `json:"fieldConditions,omitempty" xml:"fieldConditions,omitempty" type:"Repeated"`
	// example:
	//
	// 12345
	FilterId *int64 `json:"filterId,omitempty" xml:"filterId,omitempty"`
	// example:
	//
	// true
	HasAudio *bool `json:"hasAudio,omitempty" xml:"hasAudio,omitempty"`
	// example:
	//
	// true
	HasImage            *bool `json:"hasImage,omitempty" xml:"hasImage,omitempty"`
	HasMultiModeContent *bool `json:"hasMultiModeContent,omitempty" xml:"hasMultiModeContent,omitempty"`
	// example:
	//
	// true
	HasVideo *bool `json:"hasVideo,omitempty" xml:"hasVideo,omitempty"`
	// example:
	//
	// finance.sina.com.cn
	HostNameList []*string `json:"hostNameList,omitempty" xml:"hostNameList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	InfluenceLevel *int64 `json:"influenceLevel,omitempty" xml:"influenceLevel,omitempty"`
	// example:
	//
	// 123,456
	KeywordTagIds []*int64 `json:"keywordTagIds,omitempty" xml:"keywordTagIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	LikesLevel      *int32    `json:"likesLevel,omitempty" xml:"likesLevel,omitempty"`
	MaterialTagList []*string `json:"materialTagList,omitempty" xml:"materialTagList,omitempty" type:"Repeated"`
	// example:
	//
	// 123,456
	MediaLibraryIdList []*string `json:"mediaLibraryIdList,omitempty" xml:"mediaLibraryIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 新浪财经
	MediaNameList []*string `json:"mediaNameList,omitempty" xml:"mediaNameList,omitempty" type:"Repeated"`
	// example:
	//
	// WEIBO-REPOST_WEIBO
	MediaTypeList []*string `json:"mediaTypeList,omitempty" xml:"mediaTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// COMMENT
	MessageTypeList []*string `json:"messageTypeList,omitempty" xml:"messageTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNow *int32 `json:"pageNow,omitempty" xml:"pageNow,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 485738373837374
	ParentDocId *string `json:"parentDocId,omitempty" xml:"parentDocId,omitempty"`
	// example:
	//
	// ["A&&B","C"]
	PosKeywordList []*string `json:"posKeywordList,omitempty" xml:"posKeywordList,omitempty" type:"Repeated"`
	// example:
	//
	// ["A&&B","C"]
	PosKeywordListInTitle []*string `json:"posKeywordListInTitle,omitempty" xml:"posKeywordListInTitle,omitempty" type:"Repeated"`
	// example:
	//
	// 1234
	ProjectId *int64 `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 1
	PropagationLevel *int64 `json:"propagationLevel,omitempty" xml:"propagationLevel,omitempty"`
	// example:
	//
	// 1620452881429
	PublishTimeEnd *int64 `json:"publishTimeEnd,omitempty" xml:"publishTimeEnd,omitempty"`
	// example:
	//
	// 1610452881429
	PublishTimeStart *int64 `json:"publishTimeStart,omitempty" xml:"publishTimeStart,omitempty"`
	// example:
	//
	// 1
	ReadsLevel *int32 `json:"readsLevel,omitempty" xml:"readsLevel,omitempty"`
	// example:
	//
	// 1
	RelevanceLevel *int32 `json:"relevanceLevel,omitempty" xml:"relevanceLevel,omitempty"`
	// example:
	//
	// 1
	RepostLevel *int32 `json:"repostLevel,omitempty" xml:"repostLevel,omitempty"`
	// example:
	//
	// PUBLISH_TIME
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// example:
	//
	// +
	SortByDirection *string `json:"sortByDirection,omitempty" xml:"sortByDirection,omitempty"`
	// example:
	//
	// #xxxx#
	TopicList []*string `json:"topicList,omitempty" xml:"topicList,omitempty" type:"Repeated"`
	// example:
	//
	// 1620452881429
	UpdateTimeEnd *int64 `json:"updateTimeEnd,omitempty" xml:"updateTimeEnd,omitempty"`
	// example:
	//
	// 1610452881429
	UpdateTimeStart *int64 `json:"updateTimeStart,omitempty" xml:"updateTimeStart,omitempty"`
}

func (s SearchCondition) String() string {
	return dara.Prettify(s)
}

func (s SearchCondition) GoString() string {
	return s.String()
}

func (s *SearchCondition) GetAssKeywordList() []*string {
	return s.AssKeywordList
}

func (s *SearchCondition) GetAtAuthorNameList() []*string {
	return s.AtAuthorNameList
}

func (s *SearchCondition) GetAuthorNameList() []*string {
	return s.AuthorNameList
}

func (s *SearchCondition) GetCommentsLevel() *int32 {
	return s.CommentsLevel
}

func (s *SearchCondition) GetContentLenLevel() *int32 {
	return s.ContentLenLevel
}

func (s *SearchCondition) GetCreateTimeEnd() *int64 {
	return s.CreateTimeEnd
}

func (s *SearchCondition) GetCreateTimeStart() *int64 {
	return s.CreateTimeStart
}

func (s *SearchCondition) GetDocContentSign() *string {
	return s.DocContentSign
}

func (s *SearchCondition) GetDocIdList() []*string {
	return s.DocIdList
}

func (s *SearchCondition) GetDuplicateRemoval() *bool {
	return s.DuplicateRemoval
}

func (s *SearchCondition) GetEmotionType() *int32 {
	return s.EmotionType
}

func (s *SearchCondition) GetEnableKeywordHighlight() *bool {
	return s.EnableKeywordHighlight
}

func (s *SearchCondition) GetExcludeAtAuthorNameList() []*string {
	return s.ExcludeAtAuthorNameList
}

func (s *SearchCondition) GetExcludeAuthorNameList() []*string {
	return s.ExcludeAuthorNameList
}

func (s *SearchCondition) GetExcludeHostNameList() []*string {
	return s.ExcludeHostNameList
}

func (s *SearchCondition) GetExcludeKeywordList() []*string {
	return s.ExcludeKeywordList
}

func (s *SearchCondition) GetExcludeKeywordListInTitle() []*string {
	return s.ExcludeKeywordListInTitle
}

func (s *SearchCondition) GetExcludeKeywordTagIds() []*int64 {
	return s.ExcludeKeywordTagIds
}

func (s *SearchCondition) GetExcludeMaterialTagList() []*string {
	return s.ExcludeMaterialTagList
}

func (s *SearchCondition) GetExcludeMediaLibraryIdList() []*string {
	return s.ExcludeMediaLibraryIdList
}

func (s *SearchCondition) GetExcludeMediaNameList() []*string {
	return s.ExcludeMediaNameList
}

func (s *SearchCondition) GetExcludeMediaTypeList() []*string {
	return s.ExcludeMediaTypeList
}

func (s *SearchCondition) GetExcludeMessageTypeList() []*string {
	return s.ExcludeMessageTypeList
}

func (s *SearchCondition) GetFieldConditions() []*FieldCondition {
	return s.FieldConditions
}

func (s *SearchCondition) GetFilterId() *int64 {
	return s.FilterId
}

func (s *SearchCondition) GetHasAudio() *bool {
	return s.HasAudio
}

func (s *SearchCondition) GetHasImage() *bool {
	return s.HasImage
}

func (s *SearchCondition) GetHasMultiModeContent() *bool {
	return s.HasMultiModeContent
}

func (s *SearchCondition) GetHasVideo() *bool {
	return s.HasVideo
}

func (s *SearchCondition) GetHostNameList() []*string {
	return s.HostNameList
}

func (s *SearchCondition) GetInfluenceLevel() *int64 {
	return s.InfluenceLevel
}

func (s *SearchCondition) GetKeywordTagIds() []*int64 {
	return s.KeywordTagIds
}

func (s *SearchCondition) GetLikesLevel() *int32 {
	return s.LikesLevel
}

func (s *SearchCondition) GetMaterialTagList() []*string {
	return s.MaterialTagList
}

func (s *SearchCondition) GetMediaLibraryIdList() []*string {
	return s.MediaLibraryIdList
}

func (s *SearchCondition) GetMediaNameList() []*string {
	return s.MediaNameList
}

func (s *SearchCondition) GetMediaTypeList() []*string {
	return s.MediaTypeList
}

func (s *SearchCondition) GetMessageTypeList() []*string {
	return s.MessageTypeList
}

func (s *SearchCondition) GetPageNow() *int32 {
	return s.PageNow
}

func (s *SearchCondition) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchCondition) GetParentDocId() *string {
	return s.ParentDocId
}

func (s *SearchCondition) GetPosKeywordList() []*string {
	return s.PosKeywordList
}

func (s *SearchCondition) GetPosKeywordListInTitle() []*string {
	return s.PosKeywordListInTitle
}

func (s *SearchCondition) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *SearchCondition) GetPropagationLevel() *int64 {
	return s.PropagationLevel
}

func (s *SearchCondition) GetPublishTimeEnd() *int64 {
	return s.PublishTimeEnd
}

func (s *SearchCondition) GetPublishTimeStart() *int64 {
	return s.PublishTimeStart
}

func (s *SearchCondition) GetReadsLevel() *int32 {
	return s.ReadsLevel
}

func (s *SearchCondition) GetRelevanceLevel() *int32 {
	return s.RelevanceLevel
}

func (s *SearchCondition) GetRepostLevel() *int32 {
	return s.RepostLevel
}

func (s *SearchCondition) GetSortBy() *string {
	return s.SortBy
}

func (s *SearchCondition) GetSortByDirection() *string {
	return s.SortByDirection
}

func (s *SearchCondition) GetTopicList() []*string {
	return s.TopicList
}

func (s *SearchCondition) GetUpdateTimeEnd() *int64 {
	return s.UpdateTimeEnd
}

func (s *SearchCondition) GetUpdateTimeStart() *int64 {
	return s.UpdateTimeStart
}

func (s *SearchCondition) SetAssKeywordList(v []*string) *SearchCondition {
	s.AssKeywordList = v
	return s
}

func (s *SearchCondition) SetAtAuthorNameList(v []*string) *SearchCondition {
	s.AtAuthorNameList = v
	return s
}

func (s *SearchCondition) SetAuthorNameList(v []*string) *SearchCondition {
	s.AuthorNameList = v
	return s
}

func (s *SearchCondition) SetCommentsLevel(v int32) *SearchCondition {
	s.CommentsLevel = &v
	return s
}

func (s *SearchCondition) SetContentLenLevel(v int32) *SearchCondition {
	s.ContentLenLevel = &v
	return s
}

func (s *SearchCondition) SetCreateTimeEnd(v int64) *SearchCondition {
	s.CreateTimeEnd = &v
	return s
}

func (s *SearchCondition) SetCreateTimeStart(v int64) *SearchCondition {
	s.CreateTimeStart = &v
	return s
}

func (s *SearchCondition) SetDocContentSign(v string) *SearchCondition {
	s.DocContentSign = &v
	return s
}

func (s *SearchCondition) SetDocIdList(v []*string) *SearchCondition {
	s.DocIdList = v
	return s
}

func (s *SearchCondition) SetDuplicateRemoval(v bool) *SearchCondition {
	s.DuplicateRemoval = &v
	return s
}

func (s *SearchCondition) SetEmotionType(v int32) *SearchCondition {
	s.EmotionType = &v
	return s
}

func (s *SearchCondition) SetEnableKeywordHighlight(v bool) *SearchCondition {
	s.EnableKeywordHighlight = &v
	return s
}

func (s *SearchCondition) SetExcludeAtAuthorNameList(v []*string) *SearchCondition {
	s.ExcludeAtAuthorNameList = v
	return s
}

func (s *SearchCondition) SetExcludeAuthorNameList(v []*string) *SearchCondition {
	s.ExcludeAuthorNameList = v
	return s
}

func (s *SearchCondition) SetExcludeHostNameList(v []*string) *SearchCondition {
	s.ExcludeHostNameList = v
	return s
}

func (s *SearchCondition) SetExcludeKeywordList(v []*string) *SearchCondition {
	s.ExcludeKeywordList = v
	return s
}

func (s *SearchCondition) SetExcludeKeywordListInTitle(v []*string) *SearchCondition {
	s.ExcludeKeywordListInTitle = v
	return s
}

func (s *SearchCondition) SetExcludeKeywordTagIds(v []*int64) *SearchCondition {
	s.ExcludeKeywordTagIds = v
	return s
}

func (s *SearchCondition) SetExcludeMaterialTagList(v []*string) *SearchCondition {
	s.ExcludeMaterialTagList = v
	return s
}

func (s *SearchCondition) SetExcludeMediaLibraryIdList(v []*string) *SearchCondition {
	s.ExcludeMediaLibraryIdList = v
	return s
}

func (s *SearchCondition) SetExcludeMediaNameList(v []*string) *SearchCondition {
	s.ExcludeMediaNameList = v
	return s
}

func (s *SearchCondition) SetExcludeMediaTypeList(v []*string) *SearchCondition {
	s.ExcludeMediaTypeList = v
	return s
}

func (s *SearchCondition) SetExcludeMessageTypeList(v []*string) *SearchCondition {
	s.ExcludeMessageTypeList = v
	return s
}

func (s *SearchCondition) SetFieldConditions(v []*FieldCondition) *SearchCondition {
	s.FieldConditions = v
	return s
}

func (s *SearchCondition) SetFilterId(v int64) *SearchCondition {
	s.FilterId = &v
	return s
}

func (s *SearchCondition) SetHasAudio(v bool) *SearchCondition {
	s.HasAudio = &v
	return s
}

func (s *SearchCondition) SetHasImage(v bool) *SearchCondition {
	s.HasImage = &v
	return s
}

func (s *SearchCondition) SetHasMultiModeContent(v bool) *SearchCondition {
	s.HasMultiModeContent = &v
	return s
}

func (s *SearchCondition) SetHasVideo(v bool) *SearchCondition {
	s.HasVideo = &v
	return s
}

func (s *SearchCondition) SetHostNameList(v []*string) *SearchCondition {
	s.HostNameList = v
	return s
}

func (s *SearchCondition) SetInfluenceLevel(v int64) *SearchCondition {
	s.InfluenceLevel = &v
	return s
}

func (s *SearchCondition) SetKeywordTagIds(v []*int64) *SearchCondition {
	s.KeywordTagIds = v
	return s
}

func (s *SearchCondition) SetLikesLevel(v int32) *SearchCondition {
	s.LikesLevel = &v
	return s
}

func (s *SearchCondition) SetMaterialTagList(v []*string) *SearchCondition {
	s.MaterialTagList = v
	return s
}

func (s *SearchCondition) SetMediaLibraryIdList(v []*string) *SearchCondition {
	s.MediaLibraryIdList = v
	return s
}

func (s *SearchCondition) SetMediaNameList(v []*string) *SearchCondition {
	s.MediaNameList = v
	return s
}

func (s *SearchCondition) SetMediaTypeList(v []*string) *SearchCondition {
	s.MediaTypeList = v
	return s
}

func (s *SearchCondition) SetMessageTypeList(v []*string) *SearchCondition {
	s.MessageTypeList = v
	return s
}

func (s *SearchCondition) SetPageNow(v int32) *SearchCondition {
	s.PageNow = &v
	return s
}

func (s *SearchCondition) SetPageSize(v int32) *SearchCondition {
	s.PageSize = &v
	return s
}

func (s *SearchCondition) SetParentDocId(v string) *SearchCondition {
	s.ParentDocId = &v
	return s
}

func (s *SearchCondition) SetPosKeywordList(v []*string) *SearchCondition {
	s.PosKeywordList = v
	return s
}

func (s *SearchCondition) SetPosKeywordListInTitle(v []*string) *SearchCondition {
	s.PosKeywordListInTitle = v
	return s
}

func (s *SearchCondition) SetProjectId(v int64) *SearchCondition {
	s.ProjectId = &v
	return s
}

func (s *SearchCondition) SetPropagationLevel(v int64) *SearchCondition {
	s.PropagationLevel = &v
	return s
}

func (s *SearchCondition) SetPublishTimeEnd(v int64) *SearchCondition {
	s.PublishTimeEnd = &v
	return s
}

func (s *SearchCondition) SetPublishTimeStart(v int64) *SearchCondition {
	s.PublishTimeStart = &v
	return s
}

func (s *SearchCondition) SetReadsLevel(v int32) *SearchCondition {
	s.ReadsLevel = &v
	return s
}

func (s *SearchCondition) SetRelevanceLevel(v int32) *SearchCondition {
	s.RelevanceLevel = &v
	return s
}

func (s *SearchCondition) SetRepostLevel(v int32) *SearchCondition {
	s.RepostLevel = &v
	return s
}

func (s *SearchCondition) SetSortBy(v string) *SearchCondition {
	s.SortBy = &v
	return s
}

func (s *SearchCondition) SetSortByDirection(v string) *SearchCondition {
	s.SortByDirection = &v
	return s
}

func (s *SearchCondition) SetTopicList(v []*string) *SearchCondition {
	s.TopicList = v
	return s
}

func (s *SearchCondition) SetUpdateTimeEnd(v int64) *SearchCondition {
	s.UpdateTimeEnd = &v
	return s
}

func (s *SearchCondition) SetUpdateTimeStart(v int64) *SearchCondition {
	s.UpdateTimeStart = &v
	return s
}

func (s *SearchCondition) Validate() error {
	if s.FieldConditions != nil {
		for _, item := range s.FieldConditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
