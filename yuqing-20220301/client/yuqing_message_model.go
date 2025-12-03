// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iYuqingMessage interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *YuqingMessage
	GetAppName() *string
	SetAppScore(v int64) *YuqingMessage
	GetAppScore() *int64
	SetAppStoreName(v string) *YuqingMessage
	GetAppStoreName() *string
	SetAtAuthorNames(v []*string) *YuqingMessage
	GetAtAuthorNames() []*string
	SetAudioCount(v int32) *YuqingMessage
	GetAudioCount() *int32
	SetAuthorAvatarUrl(v string) *YuqingMessage
	GetAuthorAvatarUrl() *string
	SetAuthorFollowersCount(v int64) *YuqingMessage
	GetAuthorFollowersCount() *int64
	SetAuthorFriendsCount(v int64) *YuqingMessage
	GetAuthorFriendsCount() *int64
	SetAuthorGender(v string) *YuqingMessage
	GetAuthorGender() *string
	SetAuthorId(v string) *YuqingMessage
	GetAuthorId() *string
	SetAuthorLikesCount(v int64) *YuqingMessage
	GetAuthorLikesCount() *int64
	SetAuthorName(v string) *YuqingMessage
	GetAuthorName() *string
	SetAuthorProfileUrl(v string) *YuqingMessage
	GetAuthorProfileUrl() *string
	SetAuthorStatusesCount(v int64) *YuqingMessage
	GetAuthorStatusesCount() *int64
	SetAuthorVerified(v bool) *YuqingMessage
	GetAuthorVerified() *bool
	SetAuthorVerifyType(v int32) *YuqingMessage
	GetAuthorVerifyType() *int32
	SetContentAudioText(v string) *YuqingMessage
	GetContentAudioText() *string
	SetContentAudioUrls(v string) *YuqingMessage
	GetContentAudioUrls() *string
	SetContentImageText(v string) *YuqingMessage
	GetContentImageText() *string
	SetContentImageUrls(v string) *YuqingMessage
	GetContentImageUrls() *string
	SetContentLang(v string) *YuqingMessage
	GetContentLang() *string
	SetContentLen(v int64) *YuqingMessage
	GetContentLen() *int64
	SetContentVideoText(v string) *YuqingMessage
	GetContentVideoText() *string
	SetContentVideoUrls(v string) *YuqingMessage
	GetContentVideoUrls() *string
	SetCreateTime(v int64) *YuqingMessage
	GetCreateTime() *int64
	SetDocAnswersCount(v int64) *YuqingMessage
	GetDocAnswersCount() *int64
	SetDocAreas(v []*string) *YuqingMessage
	GetDocAreas() []*string
	SetDocCoinCount(v int64) *YuqingMessage
	GetDocCoinCount() *int64
	SetDocCommentsCount(v int64) *YuqingMessage
	GetDocCommentsCount() *int64
	SetDocContent(v string) *YuqingMessage
	GetDocContent() *string
	SetDocContentBrief(v string) *YuqingMessage
	GetDocContentBrief() *string
	SetDocContentSign(v string) *YuqingMessage
	GetDocContentSign() *string
	SetDocId(v string) *YuqingMessage
	GetDocId() *string
	SetDocLikesCount(v int64) *YuqingMessage
	GetDocLikesCount() *int64
	SetDocPlayCount(v int64) *YuqingMessage
	GetDocPlayCount() *int64
	SetDocReadingCount(v int64) *YuqingMessage
	GetDocReadingCount() *int64
	SetDocReadsCount(v int64) *YuqingMessage
	GetDocReadsCount() *int64
	SetDocRepostsCount(v int64) *YuqingMessage
	GetDocRepostsCount() *int64
	SetDocReprintName(v string) *YuqingMessage
	GetDocReprintName() *string
	SetDocSelfContentSign(v string) *YuqingMessage
	GetDocSelfContentSign() *string
	SetDocTitle(v string) *YuqingMessage
	GetDocTitle() *string
	SetDocUrl(v string) *YuqingMessage
	GetDocUrl() *string
	SetEmotionScore(v float64) *YuqingMessage
	GetEmotionScore() *float64
	SetEmotionType(v int32) *YuqingMessage
	GetEmotionType() *int32
	SetExtInfo(v map[string]*string) *YuqingMessage
	GetExtInfo() map[string]*string
	SetFinEventCount(v int32) *YuqingMessage
	GetFinEventCount() *int32
	SetFinanceEventList(v []*YuqingFinanceEvent) *YuqingMessage
	GetFinanceEventList() []*YuqingFinanceEvent
	SetHighlightKeywords(v []*string) *YuqingMessage
	GetHighlightKeywords() []*string
	SetImageCount(v int32) *YuqingMessage
	GetImageCount() *int32
	SetInfluenceScore(v float64) *YuqingMessage
	GetInfluenceScore() *float64
	SetMediaHosts(v []*string) *YuqingMessage
	GetMediaHosts() []*string
	SetMediaInfluenceLevel(v int32) *YuqingMessage
	GetMediaInfluenceLevel() *int32
	SetMediaName(v string) *YuqingMessage
	GetMediaName() *string
	SetMediaPropagationLevel(v int32) *YuqingMessage
	GetMediaPropagationLevel() *int32
	SetMediaType(v string) *YuqingMessage
	GetMediaType() *string
	SetMessageType(v string) *YuqingMessage
	GetMessageType() *string
	SetParentDocId(v string) *YuqingMessage
	GetParentDocId() *string
	SetPropagationScore(v float64) *YuqingMessage
	GetPropagationScore() *float64
	SetPublishTime(v int64) *YuqingMessage
	GetPublishTime() *int64
	SetRelevanceScore(v float64) *YuqingMessage
	GetRelevanceScore() *float64
	SetReportMaterialTags(v []*string) *YuqingMessage
	GetReportMaterialTags() []*string
	SetRepostList(v []*string) *YuqingMessage
	GetRepostList() []*string
	SetSimilarNumber(v int32) *YuqingMessage
	GetSimilarNumber() *int32
	SetTopics(v []*string) *YuqingMessage
	GetTopics() []*string
	SetVideoCount(v int32) *YuqingMessage
	GetVideoCount() *int32
	SetWeiboCommentId(v string) *YuqingMessage
	GetWeiboCommentId() *string
	SetWeiboMid(v string) *YuqingMessage
	GetWeiboMid() *string
}

type YuqingMessage struct {
	// example:
	//
	// Alipay
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// example:
	//
	// 3
	AppScore *int64 `json:"appScore,omitempty" xml:"appScore,omitempty"`
	// example:
	//
	// HUAWEI_APPSTORE
	AppStoreName  *string   `json:"appStoreName,omitempty" xml:"appStoreName,omitempty"`
	AtAuthorNames []*string `json:"atAuthorNames,omitempty" xml:"atAuthorNames,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	AudioCount *int32 `json:"audioCount,omitempty" xml:"audioCount,omitempty"`
	// example:
	//
	// https://xxx.png
	AuthorAvatarUrl *string `json:"authorAvatarUrl,omitempty" xml:"authorAvatarUrl,omitempty"`
	// example:
	//
	// 12
	AuthorFollowersCount *int64 `json:"authorFollowersCount,omitempty" xml:"authorFollowersCount,omitempty"`
	// example:
	//
	// 12
	AuthorFriendsCount *int64 `json:"authorFriendsCount,omitempty" xml:"authorFriendsCount,omitempty"`
	// example:
	//
	// f
	AuthorGender *string `json:"authorGender,omitempty" xml:"authorGender,omitempty"`
	// example:
	//
	// xxx
	AuthorId *string `json:"authorId,omitempty" xml:"authorId,omitempty"`
	// example:
	//
	// 100
	AuthorLikesCount *int64 `json:"authorLikesCount,omitempty" xml:"authorLikesCount,omitempty"`
	// example:
	//
	// xxx
	AuthorName *string `json:"authorName,omitempty" xml:"authorName,omitempty"`
	// example:
	//
	// http://xxx
	AuthorProfileUrl *string `json:"authorProfileUrl,omitempty" xml:"authorProfileUrl,omitempty"`
	// example:
	//
	// 12
	AuthorStatusesCount *int64 `json:"authorStatusesCount,omitempty" xml:"authorStatusesCount,omitempty"`
	// example:
	//
	// true
	AuthorVerified *bool `json:"authorVerified,omitempty" xml:"authorVerified,omitempty"`
	// example:
	//
	// 1
	AuthorVerifyType *int32 `json:"authorVerifyType,omitempty" xml:"authorVerifyType,omitempty"`
	// example:
	//
	// ["xxx","yyy"]
	ContentAudioText *string `json:"contentAudioText,omitempty" xml:"contentAudioText,omitempty"`
	// example:
	//
	// ["http://xx.mp3","http://yy.mp3"]
	ContentAudioUrls *string `json:"contentAudioUrls,omitempty" xml:"contentAudioUrls,omitempty"`
	// example:
	//
	// ["xxx","yyy"]
	ContentImageText *string `json:"contentImageText,omitempty" xml:"contentImageText,omitempty"`
	// example:
	//
	// ["http://xx.png","http://xx.jpeg"]
	ContentImageUrls *string `json:"contentImageUrls,omitempty" xml:"contentImageUrls,omitempty"`
	// example:
	//
	// zh
	ContentLang *string `json:"contentLang,omitempty" xml:"contentLang,omitempty"`
	// example:
	//
	// 100
	ContentLen *int64 `json:"contentLen,omitempty" xml:"contentLen,omitempty"`
	// example:
	//
	// ["xxx","yyy"]
	ContentVideoText *string `json:"contentVideoText,omitempty" xml:"contentVideoText,omitempty"`
	// example:
	//
	// ["http://xx.mpeg","http://yy.mp4"]
	ContentVideoUrls *string `json:"contentVideoUrls,omitempty" xml:"contentVideoUrls,omitempty"`
	// example:
	//
	// 165202930291
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 100
	DocAnswersCount *int64    `json:"docAnswersCount,omitempty" xml:"docAnswersCount,omitempty"`
	DocAreas        []*string `json:"docAreas,omitempty" xml:"docAreas,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	DocCoinCount *int64 `json:"docCoinCount,omitempty" xml:"docCoinCount,omitempty"`
	// example:
	//
	// 100
	DocCommentsCount *int64 `json:"docCommentsCount,omitempty" xml:"docCommentsCount,omitempty"`
	// example:
	//
	// xxxx
	DocContent *string `json:"docContent,omitempty" xml:"docContent,omitempty"`
	// example:
	//
	// xxxx
	DocContentBrief *string `json:"docContentBrief,omitempty" xml:"docContentBrief,omitempty"`
	// example:
	//
	// 81728391712912
	DocContentSign *string `json:"docContentSign,omitempty" xml:"docContentSign,omitempty"`
	// example:
	//
	// 48573837837232
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// 100
	DocLikesCount *int64 `json:"docLikesCount,omitempty" xml:"docLikesCount,omitempty"`
	// example:
	//
	// 100
	DocPlayCount *int64 `json:"docPlayCount,omitempty" xml:"docPlayCount,omitempty"`
	// example:
	//
	// 100
	DocReadingCount *int64 `json:"docReadingCount,omitempty" xml:"docReadingCount,omitempty"`
	// example:
	//
	// 100
	DocReadsCount *int64 `json:"docReadsCount,omitempty" xml:"docReadsCount,omitempty"`
	// example:
	//
	// 100
	DocRepostsCount *int64 `json:"docRepostsCount,omitempty" xml:"docRepostsCount,omitempty"`
	// example:
	//
	// 成都日报
	DocReprintName *string `json:"docReprintName,omitempty" xml:"docReprintName,omitempty"`
	// example:
	//
	// 81728391712912
	DocSelfContentSign *string `json:"docSelfContentSign,omitempty" xml:"docSelfContentSign,omitempty"`
	// example:
	//
	// xxx
	DocTitle *string `json:"docTitle,omitempty" xml:"docTitle,omitempty"`
	// example:
	//
	// http://xxx
	DocUrl *string `json:"docUrl,omitempty" xml:"docUrl,omitempty"`
	// example:
	//
	// 4.3
	EmotionScore *float64 `json:"emotionScore,omitempty" xml:"emotionScore,omitempty"`
	// example:
	//
	// 1
	EmotionType *int32             `json:"emotionType,omitempty" xml:"emotionType,omitempty"`
	ExtInfo     map[string]*string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// example:
	//
	// 1
	FinEventCount     *int32                `json:"finEventCount,omitempty" xml:"finEventCount,omitempty"`
	FinanceEventList  []*YuqingFinanceEvent `json:"financeEventList,omitempty" xml:"financeEventList,omitempty" type:"Repeated"`
	HighlightKeywords []*string             `json:"highlightKeywords,omitempty" xml:"highlightKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	ImageCount *int32 `json:"imageCount,omitempty" xml:"imageCount,omitempty"`
	// example:
	//
	// 4.1
	InfluenceScore *float64  `json:"influenceScore,omitempty" xml:"influenceScore,omitempty"`
	MediaHosts     []*string `json:"mediaHosts,omitempty" xml:"mediaHosts,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	MediaInfluenceLevel *int32 `json:"mediaInfluenceLevel,omitempty" xml:"mediaInfluenceLevel,omitempty"`
	// example:
	//
	// 新浪财经
	MediaName *string `json:"mediaName,omitempty" xml:"mediaName,omitempty"`
	// example:
	//
	// 1
	MediaPropagationLevel *int32 `json:"mediaPropagationLevel,omitempty" xml:"mediaPropagationLevel,omitempty"`
	// example:
	//
	// WEIBO-REPOST_WEIBO
	MediaType *string `json:"mediaType,omitempty" xml:"mediaType,omitempty"`
	// example:
	//
	// COMMENT
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// example:
	//
	// 484747382721
	ParentDocId *string `json:"parentDocId,omitempty" xml:"parentDocId,omitempty"`
	// example:
	//
	// 3.9
	PropagationScore *float64 `json:"propagationScore,omitempty" xml:"propagationScore,omitempty"`
	// example:
	//
	// 165202930291
	PublishTime *int64 `json:"publishTime,omitempty" xml:"publishTime,omitempty"`
	// example:
	//
	// 5.2
	RelevanceScore     *float64  `json:"relevanceScore,omitempty" xml:"relevanceScore,omitempty"`
	ReportMaterialTags []*string `json:"reportMaterialTags,omitempty" xml:"reportMaterialTags,omitempty" type:"Repeated"`
	RepostList         []*string `json:"repostList,omitempty" xml:"repostList,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	SimilarNumber *int32    `json:"similarNumber,omitempty" xml:"similarNumber,omitempty"`
	Topics        []*string `json:"topics,omitempty" xml:"topics,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	VideoCount *int32 `json:"videoCount,omitempty" xml:"videoCount,omitempty"`
	// example:
	//
	// 4837383832323
	WeiboCommentId *string `json:"weiboCommentId,omitempty" xml:"weiboCommentId,omitempty"`
	// example:
	//
	// 465758363823
	WeiboMid *string `json:"weiboMid,omitempty" xml:"weiboMid,omitempty"`
}

func (s YuqingMessage) String() string {
	return dara.Prettify(s)
}

func (s YuqingMessage) GoString() string {
	return s.String()
}

func (s *YuqingMessage) GetAppName() *string {
	return s.AppName
}

func (s *YuqingMessage) GetAppScore() *int64 {
	return s.AppScore
}

func (s *YuqingMessage) GetAppStoreName() *string {
	return s.AppStoreName
}

func (s *YuqingMessage) GetAtAuthorNames() []*string {
	return s.AtAuthorNames
}

func (s *YuqingMessage) GetAudioCount() *int32 {
	return s.AudioCount
}

func (s *YuqingMessage) GetAuthorAvatarUrl() *string {
	return s.AuthorAvatarUrl
}

func (s *YuqingMessage) GetAuthorFollowersCount() *int64 {
	return s.AuthorFollowersCount
}

func (s *YuqingMessage) GetAuthorFriendsCount() *int64 {
	return s.AuthorFriendsCount
}

func (s *YuqingMessage) GetAuthorGender() *string {
	return s.AuthorGender
}

func (s *YuqingMessage) GetAuthorId() *string {
	return s.AuthorId
}

func (s *YuqingMessage) GetAuthorLikesCount() *int64 {
	return s.AuthorLikesCount
}

func (s *YuqingMessage) GetAuthorName() *string {
	return s.AuthorName
}

func (s *YuqingMessage) GetAuthorProfileUrl() *string {
	return s.AuthorProfileUrl
}

func (s *YuqingMessage) GetAuthorStatusesCount() *int64 {
	return s.AuthorStatusesCount
}

func (s *YuqingMessage) GetAuthorVerified() *bool {
	return s.AuthorVerified
}

func (s *YuqingMessage) GetAuthorVerifyType() *int32 {
	return s.AuthorVerifyType
}

func (s *YuqingMessage) GetContentAudioText() *string {
	return s.ContentAudioText
}

func (s *YuqingMessage) GetContentAudioUrls() *string {
	return s.ContentAudioUrls
}

func (s *YuqingMessage) GetContentImageText() *string {
	return s.ContentImageText
}

func (s *YuqingMessage) GetContentImageUrls() *string {
	return s.ContentImageUrls
}

func (s *YuqingMessage) GetContentLang() *string {
	return s.ContentLang
}

func (s *YuqingMessage) GetContentLen() *int64 {
	return s.ContentLen
}

func (s *YuqingMessage) GetContentVideoText() *string {
	return s.ContentVideoText
}

func (s *YuqingMessage) GetContentVideoUrls() *string {
	return s.ContentVideoUrls
}

func (s *YuqingMessage) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *YuqingMessage) GetDocAnswersCount() *int64 {
	return s.DocAnswersCount
}

func (s *YuqingMessage) GetDocAreas() []*string {
	return s.DocAreas
}

func (s *YuqingMessage) GetDocCoinCount() *int64 {
	return s.DocCoinCount
}

func (s *YuqingMessage) GetDocCommentsCount() *int64 {
	return s.DocCommentsCount
}

func (s *YuqingMessage) GetDocContent() *string {
	return s.DocContent
}

func (s *YuqingMessage) GetDocContentBrief() *string {
	return s.DocContentBrief
}

func (s *YuqingMessage) GetDocContentSign() *string {
	return s.DocContentSign
}

func (s *YuqingMessage) GetDocId() *string {
	return s.DocId
}

func (s *YuqingMessage) GetDocLikesCount() *int64 {
	return s.DocLikesCount
}

func (s *YuqingMessage) GetDocPlayCount() *int64 {
	return s.DocPlayCount
}

func (s *YuqingMessage) GetDocReadingCount() *int64 {
	return s.DocReadingCount
}

func (s *YuqingMessage) GetDocReadsCount() *int64 {
	return s.DocReadsCount
}

func (s *YuqingMessage) GetDocRepostsCount() *int64 {
	return s.DocRepostsCount
}

func (s *YuqingMessage) GetDocReprintName() *string {
	return s.DocReprintName
}

func (s *YuqingMessage) GetDocSelfContentSign() *string {
	return s.DocSelfContentSign
}

func (s *YuqingMessage) GetDocTitle() *string {
	return s.DocTitle
}

func (s *YuqingMessage) GetDocUrl() *string {
	return s.DocUrl
}

func (s *YuqingMessage) GetEmotionScore() *float64 {
	return s.EmotionScore
}

func (s *YuqingMessage) GetEmotionType() *int32 {
	return s.EmotionType
}

func (s *YuqingMessage) GetExtInfo() map[string]*string {
	return s.ExtInfo
}

func (s *YuqingMessage) GetFinEventCount() *int32 {
	return s.FinEventCount
}

func (s *YuqingMessage) GetFinanceEventList() []*YuqingFinanceEvent {
	return s.FinanceEventList
}

func (s *YuqingMessage) GetHighlightKeywords() []*string {
	return s.HighlightKeywords
}

func (s *YuqingMessage) GetImageCount() *int32 {
	return s.ImageCount
}

func (s *YuqingMessage) GetInfluenceScore() *float64 {
	return s.InfluenceScore
}

func (s *YuqingMessage) GetMediaHosts() []*string {
	return s.MediaHosts
}

func (s *YuqingMessage) GetMediaInfluenceLevel() *int32 {
	return s.MediaInfluenceLevel
}

func (s *YuqingMessage) GetMediaName() *string {
	return s.MediaName
}

func (s *YuqingMessage) GetMediaPropagationLevel() *int32 {
	return s.MediaPropagationLevel
}

func (s *YuqingMessage) GetMediaType() *string {
	return s.MediaType
}

func (s *YuqingMessage) GetMessageType() *string {
	return s.MessageType
}

func (s *YuqingMessage) GetParentDocId() *string {
	return s.ParentDocId
}

func (s *YuqingMessage) GetPropagationScore() *float64 {
	return s.PropagationScore
}

func (s *YuqingMessage) GetPublishTime() *int64 {
	return s.PublishTime
}

func (s *YuqingMessage) GetRelevanceScore() *float64 {
	return s.RelevanceScore
}

func (s *YuqingMessage) GetReportMaterialTags() []*string {
	return s.ReportMaterialTags
}

func (s *YuqingMessage) GetRepostList() []*string {
	return s.RepostList
}

func (s *YuqingMessage) GetSimilarNumber() *int32 {
	return s.SimilarNumber
}

func (s *YuqingMessage) GetTopics() []*string {
	return s.Topics
}

func (s *YuqingMessage) GetVideoCount() *int32 {
	return s.VideoCount
}

func (s *YuqingMessage) GetWeiboCommentId() *string {
	return s.WeiboCommentId
}

func (s *YuqingMessage) GetWeiboMid() *string {
	return s.WeiboMid
}

func (s *YuqingMessage) SetAppName(v string) *YuqingMessage {
	s.AppName = &v
	return s
}

func (s *YuqingMessage) SetAppScore(v int64) *YuqingMessage {
	s.AppScore = &v
	return s
}

func (s *YuqingMessage) SetAppStoreName(v string) *YuqingMessage {
	s.AppStoreName = &v
	return s
}

func (s *YuqingMessage) SetAtAuthorNames(v []*string) *YuqingMessage {
	s.AtAuthorNames = v
	return s
}

func (s *YuqingMessage) SetAudioCount(v int32) *YuqingMessage {
	s.AudioCount = &v
	return s
}

func (s *YuqingMessage) SetAuthorAvatarUrl(v string) *YuqingMessage {
	s.AuthorAvatarUrl = &v
	return s
}

func (s *YuqingMessage) SetAuthorFollowersCount(v int64) *YuqingMessage {
	s.AuthorFollowersCount = &v
	return s
}

func (s *YuqingMessage) SetAuthorFriendsCount(v int64) *YuqingMessage {
	s.AuthorFriendsCount = &v
	return s
}

func (s *YuqingMessage) SetAuthorGender(v string) *YuqingMessage {
	s.AuthorGender = &v
	return s
}

func (s *YuqingMessage) SetAuthorId(v string) *YuqingMessage {
	s.AuthorId = &v
	return s
}

func (s *YuqingMessage) SetAuthorLikesCount(v int64) *YuqingMessage {
	s.AuthorLikesCount = &v
	return s
}

func (s *YuqingMessage) SetAuthorName(v string) *YuqingMessage {
	s.AuthorName = &v
	return s
}

func (s *YuqingMessage) SetAuthorProfileUrl(v string) *YuqingMessage {
	s.AuthorProfileUrl = &v
	return s
}

func (s *YuqingMessage) SetAuthorStatusesCount(v int64) *YuqingMessage {
	s.AuthorStatusesCount = &v
	return s
}

func (s *YuqingMessage) SetAuthorVerified(v bool) *YuqingMessage {
	s.AuthorVerified = &v
	return s
}

func (s *YuqingMessage) SetAuthorVerifyType(v int32) *YuqingMessage {
	s.AuthorVerifyType = &v
	return s
}

func (s *YuqingMessage) SetContentAudioText(v string) *YuqingMessage {
	s.ContentAudioText = &v
	return s
}

func (s *YuqingMessage) SetContentAudioUrls(v string) *YuqingMessage {
	s.ContentAudioUrls = &v
	return s
}

func (s *YuqingMessage) SetContentImageText(v string) *YuqingMessage {
	s.ContentImageText = &v
	return s
}

func (s *YuqingMessage) SetContentImageUrls(v string) *YuqingMessage {
	s.ContentImageUrls = &v
	return s
}

func (s *YuqingMessage) SetContentLang(v string) *YuqingMessage {
	s.ContentLang = &v
	return s
}

func (s *YuqingMessage) SetContentLen(v int64) *YuqingMessage {
	s.ContentLen = &v
	return s
}

func (s *YuqingMessage) SetContentVideoText(v string) *YuqingMessage {
	s.ContentVideoText = &v
	return s
}

func (s *YuqingMessage) SetContentVideoUrls(v string) *YuqingMessage {
	s.ContentVideoUrls = &v
	return s
}

func (s *YuqingMessage) SetCreateTime(v int64) *YuqingMessage {
	s.CreateTime = &v
	return s
}

func (s *YuqingMessage) SetDocAnswersCount(v int64) *YuqingMessage {
	s.DocAnswersCount = &v
	return s
}

func (s *YuqingMessage) SetDocAreas(v []*string) *YuqingMessage {
	s.DocAreas = v
	return s
}

func (s *YuqingMessage) SetDocCoinCount(v int64) *YuqingMessage {
	s.DocCoinCount = &v
	return s
}

func (s *YuqingMessage) SetDocCommentsCount(v int64) *YuqingMessage {
	s.DocCommentsCount = &v
	return s
}

func (s *YuqingMessage) SetDocContent(v string) *YuqingMessage {
	s.DocContent = &v
	return s
}

func (s *YuqingMessage) SetDocContentBrief(v string) *YuqingMessage {
	s.DocContentBrief = &v
	return s
}

func (s *YuqingMessage) SetDocContentSign(v string) *YuqingMessage {
	s.DocContentSign = &v
	return s
}

func (s *YuqingMessage) SetDocId(v string) *YuqingMessage {
	s.DocId = &v
	return s
}

func (s *YuqingMessage) SetDocLikesCount(v int64) *YuqingMessage {
	s.DocLikesCount = &v
	return s
}

func (s *YuqingMessage) SetDocPlayCount(v int64) *YuqingMessage {
	s.DocPlayCount = &v
	return s
}

func (s *YuqingMessage) SetDocReadingCount(v int64) *YuqingMessage {
	s.DocReadingCount = &v
	return s
}

func (s *YuqingMessage) SetDocReadsCount(v int64) *YuqingMessage {
	s.DocReadsCount = &v
	return s
}

func (s *YuqingMessage) SetDocRepostsCount(v int64) *YuqingMessage {
	s.DocRepostsCount = &v
	return s
}

func (s *YuqingMessage) SetDocReprintName(v string) *YuqingMessage {
	s.DocReprintName = &v
	return s
}

func (s *YuqingMessage) SetDocSelfContentSign(v string) *YuqingMessage {
	s.DocSelfContentSign = &v
	return s
}

func (s *YuqingMessage) SetDocTitle(v string) *YuqingMessage {
	s.DocTitle = &v
	return s
}

func (s *YuqingMessage) SetDocUrl(v string) *YuqingMessage {
	s.DocUrl = &v
	return s
}

func (s *YuqingMessage) SetEmotionScore(v float64) *YuqingMessage {
	s.EmotionScore = &v
	return s
}

func (s *YuqingMessage) SetEmotionType(v int32) *YuqingMessage {
	s.EmotionType = &v
	return s
}

func (s *YuqingMessage) SetExtInfo(v map[string]*string) *YuqingMessage {
	s.ExtInfo = v
	return s
}

func (s *YuqingMessage) SetFinEventCount(v int32) *YuqingMessage {
	s.FinEventCount = &v
	return s
}

func (s *YuqingMessage) SetFinanceEventList(v []*YuqingFinanceEvent) *YuqingMessage {
	s.FinanceEventList = v
	return s
}

func (s *YuqingMessage) SetHighlightKeywords(v []*string) *YuqingMessage {
	s.HighlightKeywords = v
	return s
}

func (s *YuqingMessage) SetImageCount(v int32) *YuqingMessage {
	s.ImageCount = &v
	return s
}

func (s *YuqingMessage) SetInfluenceScore(v float64) *YuqingMessage {
	s.InfluenceScore = &v
	return s
}

func (s *YuqingMessage) SetMediaHosts(v []*string) *YuqingMessage {
	s.MediaHosts = v
	return s
}

func (s *YuqingMessage) SetMediaInfluenceLevel(v int32) *YuqingMessage {
	s.MediaInfluenceLevel = &v
	return s
}

func (s *YuqingMessage) SetMediaName(v string) *YuqingMessage {
	s.MediaName = &v
	return s
}

func (s *YuqingMessage) SetMediaPropagationLevel(v int32) *YuqingMessage {
	s.MediaPropagationLevel = &v
	return s
}

func (s *YuqingMessage) SetMediaType(v string) *YuqingMessage {
	s.MediaType = &v
	return s
}

func (s *YuqingMessage) SetMessageType(v string) *YuqingMessage {
	s.MessageType = &v
	return s
}

func (s *YuqingMessage) SetParentDocId(v string) *YuqingMessage {
	s.ParentDocId = &v
	return s
}

func (s *YuqingMessage) SetPropagationScore(v float64) *YuqingMessage {
	s.PropagationScore = &v
	return s
}

func (s *YuqingMessage) SetPublishTime(v int64) *YuqingMessage {
	s.PublishTime = &v
	return s
}

func (s *YuqingMessage) SetRelevanceScore(v float64) *YuqingMessage {
	s.RelevanceScore = &v
	return s
}

func (s *YuqingMessage) SetReportMaterialTags(v []*string) *YuqingMessage {
	s.ReportMaterialTags = v
	return s
}

func (s *YuqingMessage) SetRepostList(v []*string) *YuqingMessage {
	s.RepostList = v
	return s
}

func (s *YuqingMessage) SetSimilarNumber(v int32) *YuqingMessage {
	s.SimilarNumber = &v
	return s
}

func (s *YuqingMessage) SetTopics(v []*string) *YuqingMessage {
	s.Topics = v
	return s
}

func (s *YuqingMessage) SetVideoCount(v int32) *YuqingMessage {
	s.VideoCount = &v
	return s
}

func (s *YuqingMessage) SetWeiboCommentId(v string) *YuqingMessage {
	s.WeiboCommentId = &v
	return s
}

func (s *YuqingMessage) SetWeiboMid(v string) *YuqingMessage {
	s.WeiboMid = &v
	return s
}

func (s *YuqingMessage) Validate() error {
	if s.FinanceEventList != nil {
		for _, item := range s.FinanceEventList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
