// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ConsoleBody struct {
	AppCode       *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	InterfaceName *string `json:"interfaceName,omitempty" xml:"interfaceName,omitempty"`
	ParamJson     *string `json:"paramJson,omitempty" xml:"paramJson,omitempty"`
	RequestId     *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TeamHashId    *string `json:"teamHashId,omitempty" xml:"teamHashId,omitempty"`
}

func (s ConsoleBody) String() string {
	return tea.Prettify(s)
}

func (s ConsoleBody) GoString() string {
	return s.String()
}

func (s *ConsoleBody) SetAppCode(v string) *ConsoleBody {
	s.AppCode = &v
	return s
}

func (s *ConsoleBody) SetInterfaceName(v string) *ConsoleBody {
	s.InterfaceName = &v
	return s
}

func (s *ConsoleBody) SetParamJson(v string) *ConsoleBody {
	s.ParamJson = &v
	return s
}

func (s *ConsoleBody) SetRequestId(v string) *ConsoleBody {
	s.RequestId = &v
	return s
}

func (s *ConsoleBody) SetTeamHashId(v string) *ConsoleBody {
	s.TeamHashId = &v
	return s
}

type FieldCondition struct {
	// example:
	//
	// xxx
	FieldName      *string  `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	NestFieldPath  *string  `json:"nestFieldPath,omitempty" xml:"nestFieldPath,omitempty"`
	NestFieldValue []*int64 `json:"nestFieldValue,omitempty" xml:"nestFieldValue,omitempty" type:"Repeated"`
	// example:
	//
	// =
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// example:
	//
	// yyy
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FieldCondition) String() string {
	return tea.Prettify(s)
}

func (s FieldCondition) GoString() string {
	return s.String()
}

func (s *FieldCondition) SetFieldName(v string) *FieldCondition {
	s.FieldName = &v
	return s
}

func (s *FieldCondition) SetNestFieldPath(v string) *FieldCondition {
	s.NestFieldPath = &v
	return s
}

func (s *FieldCondition) SetNestFieldValue(v []*int64) *FieldCondition {
	s.NestFieldValue = v
	return s
}

func (s *FieldCondition) SetOperateType(v string) *FieldCondition {
	s.OperateType = &v
	return s
}

func (s *FieldCondition) SetValue(v string) *FieldCondition {
	s.Value = &v
	return s
}

type ProductInstance struct {
	// This parameter is required.
	//
	// example:
	//
	// EUWYEEQ
	AppCode *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	// example:
	//
	// xxx
	BuyerName *string `json:"buyerName,omitempty" xml:"buyerName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ID2343231321
	BuyerUid *string `json:"buyerUid,omitempty" xml:"buyerUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// example:
	//
	// {}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// 1640292843231
	End *int64 `json:"end,omitempty" xml:"end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ID3928389103844
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// D23938474923u42
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// C394884
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// C847573
	ProductSpecCode *string `json:"productSpecCode,omitempty" xml:"productSpecCode,omitempty"`
	// example:
	//
	// 1640292843231
	Start *int64 `json:"start,omitempty" xml:"start,omitempty"`
	// example:
	//
	// XXX
	TenantName *string `json:"tenantName,omitempty" xml:"tenantName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UID284747383
	TenantUid *string `json:"tenantUid,omitempty" xml:"tenantUid,omitempty"`
}

func (s ProductInstance) String() string {
	return tea.Prettify(s)
}

func (s ProductInstance) GoString() string {
	return s.String()
}

func (s *ProductInstance) SetAppCode(v string) *ProductInstance {
	s.AppCode = &v
	return s
}

func (s *ProductInstance) SetBuyerName(v string) *ProductInstance {
	s.BuyerName = &v
	return s
}

func (s *ProductInstance) SetBuyerUid(v string) *ProductInstance {
	s.BuyerUid = &v
	return s
}

func (s *ProductInstance) SetChannel(v string) *ProductInstance {
	s.Channel = &v
	return s
}

func (s *ProductInstance) SetConfig(v string) *ProductInstance {
	s.Config = &v
	return s
}

func (s *ProductInstance) SetEnd(v int64) *ProductInstance {
	s.End = &v
	return s
}

func (s *ProductInstance) SetInstanceId(v string) *ProductInstance {
	s.InstanceId = &v
	return s
}

func (s *ProductInstance) SetOrderNo(v string) *ProductInstance {
	s.OrderNo = &v
	return s
}

func (s *ProductInstance) SetProductCode(v string) *ProductInstance {
	s.ProductCode = &v
	return s
}

func (s *ProductInstance) SetProductSpecCode(v string) *ProductInstance {
	s.ProductSpecCode = &v
	return s
}

func (s *ProductInstance) SetStart(v int64) *ProductInstance {
	s.Start = &v
	return s
}

func (s *ProductInstance) SetTenantName(v string) *ProductInstance {
	s.TenantName = &v
	return s
}

func (s *ProductInstance) SetTenantUid(v string) *ProductInstance {
	s.TenantUid = &v
	return s
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
	return tea.Prettify(s)
}

func (s SearchCondition) GoString() string {
	return s.String()
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

type YuqingFinanceEvent struct {
	ComprehensiveRisk       *float64  `json:"comprehensiveRisk,omitempty" xml:"comprehensiveRisk,omitempty"`
	EntityArea              *string   `json:"entityArea,omitempty" xml:"entityArea,omitempty"`
	EntityCrn               *string   `json:"entityCrn,omitempty" xml:"entityCrn,omitempty"`
	EntityEmotionScore      *float64  `json:"entityEmotionScore,omitempty" xml:"entityEmotionScore,omitempty"`
	EntityId                *int64    `json:"entityId,omitempty" xml:"entityId,omitempty"`
	EntityName              *string   `json:"entityName,omitempty" xml:"entityName,omitempty"`
	EntityRelevanceScore    *float64  `json:"entityRelevanceScore,omitempty" xml:"entityRelevanceScore,omitempty"`
	EntityShowName          *string   `json:"entityShowName,omitempty" xml:"entityShowName,omitempty"`
	EntitySummary           *string   `json:"entitySummary,omitempty" xml:"entitySummary,omitempty"`
	EntityType              *string   `json:"entityType,omitempty" xml:"entityType,omitempty"`
	EventId                 *string   `json:"eventId,omitempty" xml:"eventId,omitempty"`
	EventLevel3Code         *int64    `json:"eventLevel3Code,omitempty" xml:"eventLevel3Code,omitempty"`
	EventLevel3Name         *string   `json:"eventLevel3Name,omitempty" xml:"eventLevel3Name,omitempty"`
	EventTags               *string   `json:"eventTags,omitempty" xml:"eventTags,omitempty"`
	EventTime               *int64    `json:"eventTime,omitempty" xml:"eventTime,omitempty"`
	SecurityAbbreviation    *string   `json:"securityAbbreviation,omitempty" xml:"securityAbbreviation,omitempty"`
	SecurityCategoryCodes   []*string `json:"securityCategoryCodes,omitempty" xml:"securityCategoryCodes,omitempty" type:"Repeated"`
	SecurityCodes           []*string `json:"securityCodes,omitempty" xml:"securityCodes,omitempty" type:"Repeated"`
	SecurityMarketsCodes    []*string `json:"securityMarketsCodes,omitempty" xml:"securityMarketsCodes,omitempty" type:"Repeated"`
	SpamScore               *float64  `json:"spamScore,omitempty" xml:"spamScore,omitempty"`
	UserSubscribeEntityTags []*string `json:"userSubscribeEntityTags,omitempty" xml:"userSubscribeEntityTags,omitempty" type:"Repeated"`
	UserSubscribeEventTags  []*int64  `json:"userSubscribeEventTags,omitempty" xml:"userSubscribeEventTags,omitempty" type:"Repeated"`
}

func (s YuqingFinanceEvent) String() string {
	return tea.Prettify(s)
}

func (s YuqingFinanceEvent) GoString() string {
	return s.String()
}

func (s *YuqingFinanceEvent) SetComprehensiveRisk(v float64) *YuqingFinanceEvent {
	s.ComprehensiveRisk = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntityArea(v string) *YuqingFinanceEvent {
	s.EntityArea = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntityCrn(v string) *YuqingFinanceEvent {
	s.EntityCrn = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntityEmotionScore(v float64) *YuqingFinanceEvent {
	s.EntityEmotionScore = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntityId(v int64) *YuqingFinanceEvent {
	s.EntityId = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntityName(v string) *YuqingFinanceEvent {
	s.EntityName = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntityRelevanceScore(v float64) *YuqingFinanceEvent {
	s.EntityRelevanceScore = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntityShowName(v string) *YuqingFinanceEvent {
	s.EntityShowName = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntitySummary(v string) *YuqingFinanceEvent {
	s.EntitySummary = &v
	return s
}

func (s *YuqingFinanceEvent) SetEntityType(v string) *YuqingFinanceEvent {
	s.EntityType = &v
	return s
}

func (s *YuqingFinanceEvent) SetEventId(v string) *YuqingFinanceEvent {
	s.EventId = &v
	return s
}

func (s *YuqingFinanceEvent) SetEventLevel3Code(v int64) *YuqingFinanceEvent {
	s.EventLevel3Code = &v
	return s
}

func (s *YuqingFinanceEvent) SetEventLevel3Name(v string) *YuqingFinanceEvent {
	s.EventLevel3Name = &v
	return s
}

func (s *YuqingFinanceEvent) SetEventTags(v string) *YuqingFinanceEvent {
	s.EventTags = &v
	return s
}

func (s *YuqingFinanceEvent) SetEventTime(v int64) *YuqingFinanceEvent {
	s.EventTime = &v
	return s
}

func (s *YuqingFinanceEvent) SetSecurityAbbreviation(v string) *YuqingFinanceEvent {
	s.SecurityAbbreviation = &v
	return s
}

func (s *YuqingFinanceEvent) SetSecurityCategoryCodes(v []*string) *YuqingFinanceEvent {
	s.SecurityCategoryCodes = v
	return s
}

func (s *YuqingFinanceEvent) SetSecurityCodes(v []*string) *YuqingFinanceEvent {
	s.SecurityCodes = v
	return s
}

func (s *YuqingFinanceEvent) SetSecurityMarketsCodes(v []*string) *YuqingFinanceEvent {
	s.SecurityMarketsCodes = v
	return s
}

func (s *YuqingFinanceEvent) SetSpamScore(v float64) *YuqingFinanceEvent {
	s.SpamScore = &v
	return s
}

func (s *YuqingFinanceEvent) SetUserSubscribeEntityTags(v []*string) *YuqingFinanceEvent {
	s.UserSubscribeEntityTags = v
	return s
}

func (s *YuqingFinanceEvent) SetUserSubscribeEventTags(v []*int64) *YuqingFinanceEvent {
	s.UserSubscribeEventTags = v
	return s
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
	return tea.Prettify(s)
}

func (s YuqingMessage) GoString() string {
	return s.String()
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

type CloseProductRequest struct {
	ProductInstance *ProductInstance `json:"productInstance,omitempty" xml:"productInstance,omitempty"`
	RequestId       *string          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CloseProductRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseProductRequest) GoString() string {
	return s.String()
}

func (s *CloseProductRequest) SetProductInstance(v *ProductInstance) *CloseProductRequest {
	s.ProductInstance = v
	return s
}

func (s *CloseProductRequest) SetRequestId(v string) *CloseProductRequest {
	s.RequestId = &v
	return s
}

type CloseProductResponseBody struct {
	Data      *int64  `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CloseProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseProductResponseBody) GoString() string {
	return s.String()
}

func (s *CloseProductResponseBody) SetData(v int64) *CloseProductResponseBody {
	s.Data = &v
	return s
}

func (s *CloseProductResponseBody) SetRequestId(v string) *CloseProductResponseBody {
	s.RequestId = &v
	return s
}

type CloseProductResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseProductResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseProductResponse) GoString() string {
	return s.String()
}

func (s *CloseProductResponse) SetHeaders(v map[string]*string) *CloseProductResponse {
	s.Headers = v
	return s
}

func (s *CloseProductResponse) SetStatusCode(v int32) *CloseProductResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseProductResponse) SetBody(v *CloseProductResponseBody) *CloseProductResponse {
	s.Body = v
	return s
}

type ConsoleApiProxyRequest struct {
	Body *ConsoleBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConsoleApiProxyRequest) String() string {
	return tea.Prettify(s)
}

func (s ConsoleApiProxyRequest) GoString() string {
	return s.String()
}

func (s *ConsoleApiProxyRequest) SetBody(v *ConsoleBody) *ConsoleApiProxyRequest {
	s.Body = v
	return s
}

type ConsoleApiProxyResponseBody struct {
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResultJson *string `json:"resultJson,omitempty" xml:"resultJson,omitempty"`
}

func (s ConsoleApiProxyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConsoleApiProxyResponseBody) GoString() string {
	return s.String()
}

func (s *ConsoleApiProxyResponseBody) SetRequestId(v string) *ConsoleApiProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConsoleApiProxyResponseBody) SetResultJson(v string) *ConsoleApiProxyResponseBody {
	s.ResultJson = &v
	return s
}

type ConsoleApiProxyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConsoleApiProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConsoleApiProxyResponse) String() string {
	return tea.Prettify(s)
}

func (s ConsoleApiProxyResponse) GoString() string {
	return s.String()
}

func (s *ConsoleApiProxyResponse) SetHeaders(v map[string]*string) *ConsoleApiProxyResponse {
	s.Headers = v
	return s
}

func (s *ConsoleApiProxyResponse) SetStatusCode(v int32) *ConsoleApiProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *ConsoleApiProxyResponse) SetBody(v *ConsoleApiProxyResponseBody) *ConsoleApiProxyResponse {
	s.Body = v
	return s
}

type ConsoleProxyRequest struct {
	AppCode       *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	InterfaceName *string `json:"interfaceName,omitempty" xml:"interfaceName,omitempty"`
	ParamJson     *string `json:"paramJson,omitempty" xml:"paramJson,omitempty"`
	RequestId     *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TeamHashId    *string `json:"teamHashId,omitempty" xml:"teamHashId,omitempty"`
}

func (s ConsoleProxyRequest) String() string {
	return tea.Prettify(s)
}

func (s ConsoleProxyRequest) GoString() string {
	return s.String()
}

func (s *ConsoleProxyRequest) SetAppCode(v string) *ConsoleProxyRequest {
	s.AppCode = &v
	return s
}

func (s *ConsoleProxyRequest) SetInterfaceName(v string) *ConsoleProxyRequest {
	s.InterfaceName = &v
	return s
}

func (s *ConsoleProxyRequest) SetParamJson(v string) *ConsoleProxyRequest {
	s.ParamJson = &v
	return s
}

func (s *ConsoleProxyRequest) SetRequestId(v string) *ConsoleProxyRequest {
	s.RequestId = &v
	return s
}

func (s *ConsoleProxyRequest) SetTeamHashId(v string) *ConsoleProxyRequest {
	s.TeamHashId = &v
	return s
}

type ConsoleProxyResponseBody struct {
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResultJson *string `json:"resultJson,omitempty" xml:"resultJson,omitempty"`
}

func (s ConsoleProxyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConsoleProxyResponseBody) GoString() string {
	return s.String()
}

func (s *ConsoleProxyResponseBody) SetRequestId(v string) *ConsoleProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConsoleProxyResponseBody) SetResultJson(v string) *ConsoleProxyResponseBody {
	s.ResultJson = &v
	return s
}

type ConsoleProxyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConsoleProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConsoleProxyResponse) String() string {
	return tea.Prettify(s)
}

func (s ConsoleProxyResponse) GoString() string {
	return s.String()
}

func (s *ConsoleProxyResponse) SetHeaders(v map[string]*string) *ConsoleProxyResponse {
	s.Headers = v
	return s
}

func (s *ConsoleProxyResponse) SetStatusCode(v int32) *ConsoleProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *ConsoleProxyResponse) SetBody(v *ConsoleProxyResponseBody) *ConsoleProxyResponse {
	s.Body = v
	return s
}

type GetAnalysisTaskResultRequest struct {
	AnalysisId *int64  `json:"analysisId,omitempty" xml:"analysisId,omitempty"`
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TeamHashId *string `json:"teamHashId,omitempty" xml:"teamHashId,omitempty"`
}

func (s GetAnalysisTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAnalysisTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetAnalysisTaskResultRequest) SetAnalysisId(v int64) *GetAnalysisTaskResultRequest {
	s.AnalysisId = &v
	return s
}

func (s *GetAnalysisTaskResultRequest) SetRequestId(v string) *GetAnalysisTaskResultRequest {
	s.RequestId = &v
	return s
}

func (s *GetAnalysisTaskResultRequest) SetTeamHashId(v string) *GetAnalysisTaskResultRequest {
	s.TeamHashId = &v
	return s
}

type GetAnalysisTaskResultResponseBody struct {
	AnalysisId *int64  `json:"analysisId,omitempty" xml:"analysisId,omitempty"`
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResultJson *string `json:"resultJson,omitempty" xml:"resultJson,omitempty"`
}

func (s GetAnalysisTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAnalysisTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAnalysisTaskResultResponseBody) SetAnalysisId(v int64) *GetAnalysisTaskResultResponseBody {
	s.AnalysisId = &v
	return s
}

func (s *GetAnalysisTaskResultResponseBody) SetRequestId(v string) *GetAnalysisTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAnalysisTaskResultResponseBody) SetResultJson(v string) *GetAnalysisTaskResultResponseBody {
	s.ResultJson = &v
	return s
}

type GetAnalysisTaskResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAnalysisTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAnalysisTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAnalysisTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetAnalysisTaskResultResponse) SetHeaders(v map[string]*string) *GetAnalysisTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetAnalysisTaskResultResponse) SetStatusCode(v int32) *GetAnalysisTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAnalysisTaskResultResponse) SetBody(v *GetAnalysisTaskResultResponseBody) *GetAnalysisTaskResultResponse {
	s.Body = v
	return s
}

type OpenProductRequest struct {
	ClientToken     *string          `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ProductInstance *ProductInstance `json:"productInstance,omitempty" xml:"productInstance,omitempty"`
	RequestId       *string          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s OpenProductRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenProductRequest) GoString() string {
	return s.String()
}

func (s *OpenProductRequest) SetClientToken(v string) *OpenProductRequest {
	s.ClientToken = &v
	return s
}

func (s *OpenProductRequest) SetProductInstance(v *ProductInstance) *OpenProductRequest {
	s.ProductInstance = v
	return s
}

func (s *OpenProductRequest) SetRequestId(v string) *OpenProductRequest {
	s.RequestId = &v
	return s
}

type OpenProductResponseBody struct {
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s OpenProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenProductResponseBody) GoString() string {
	return s.String()
}

func (s *OpenProductResponseBody) SetId(v int64) *OpenProductResponseBody {
	s.Id = &v
	return s
}

func (s *OpenProductResponseBody) SetRequestId(v string) *OpenProductResponseBody {
	s.RequestId = &v
	return s
}

type OpenProductResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenProductResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenProductResponse) GoString() string {
	return s.String()
}

func (s *OpenProductResponse) SetHeaders(v map[string]*string) *OpenProductResponse {
	s.Headers = v
	return s
}

func (s *OpenProductResponse) SetStatusCode(v int32) *OpenProductResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenProductResponse) SetBody(v *OpenProductResponseBody) *OpenProductResponse {
	s.Body = v
	return s
}

type QueryProductInstanceListRequest struct {
	AppCode   *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	FromTime  *int64  `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TenantUid *string `json:"tenantUid,omitempty" xml:"tenantUid,omitempty"`
	ToTime    *int64  `json:"toTime,omitempty" xml:"toTime,omitempty"`
}

func (s QueryProductInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProductInstanceListRequest) GoString() string {
	return s.String()
}

func (s *QueryProductInstanceListRequest) SetAppCode(v string) *QueryProductInstanceListRequest {
	s.AppCode = &v
	return s
}

func (s *QueryProductInstanceListRequest) SetFromTime(v int64) *QueryProductInstanceListRequest {
	s.FromTime = &v
	return s
}

func (s *QueryProductInstanceListRequest) SetRequestId(v string) *QueryProductInstanceListRequest {
	s.RequestId = &v
	return s
}

func (s *QueryProductInstanceListRequest) SetTenantUid(v string) *QueryProductInstanceListRequest {
	s.TenantUid = &v
	return s
}

func (s *QueryProductInstanceListRequest) SetToTime(v int64) *QueryProductInstanceListRequest {
	s.ToTime = &v
	return s
}

type QueryProductInstanceListResponseBody struct {
	InstanceList []*ProductInstance `json:"instanceList,omitempty" xml:"instanceList,omitempty" type:"Repeated"`
	RequestId    *string            `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryProductInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProductInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProductInstanceListResponseBody) SetInstanceList(v []*ProductInstance) *QueryProductInstanceListResponseBody {
	s.InstanceList = v
	return s
}

func (s *QueryProductInstanceListResponseBody) SetRequestId(v string) *QueryProductInstanceListResponseBody {
	s.RequestId = &v
	return s
}

type QueryProductInstanceListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProductInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProductInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProductInstanceListResponse) GoString() string {
	return s.String()
}

func (s *QueryProductInstanceListResponse) SetHeaders(v map[string]*string) *QueryProductInstanceListResponse {
	s.Headers = v
	return s
}

func (s *QueryProductInstanceListResponse) SetStatusCode(v int32) *QueryProductInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProductInstanceListResponse) SetBody(v *QueryProductInstanceListResponseBody) *QueryProductInstanceListResponse {
	s.Body = v
	return s
}

type QueryYuqingMessageRequest struct {
	// example:
	//
	// 5645a6c9-7d21-4926-a410-db9a1af85faa
	RequestId       *string          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SearchCondition *SearchCondition `json:"searchCondition,omitempty" xml:"searchCondition,omitempty"`
	// example:
	//
	// xxxx43434dsdsd
	TeamHashId *string `json:"teamHashId,omitempty" xml:"teamHashId,omitempty"`
}

func (s QueryYuqingMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryYuqingMessageRequest) GoString() string {
	return s.String()
}

func (s *QueryYuqingMessageRequest) SetRequestId(v string) *QueryYuqingMessageRequest {
	s.RequestId = &v
	return s
}

func (s *QueryYuqingMessageRequest) SetSearchCondition(v *SearchCondition) *QueryYuqingMessageRequest {
	s.SearchCondition = v
	return s
}

func (s *QueryYuqingMessageRequest) SetTeamHashId(v string) *QueryYuqingMessageRequest {
	s.TeamHashId = &v
	return s
}

type QueryYuqingMessageResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// adacae47-6fc0-45c6-897c-26201aefbdfd
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 20
	TotalCount     *int64           `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	YuqingMessages []*YuqingMessage `json:"yuqingMessages,omitempty" xml:"yuqingMessages,omitempty" type:"Repeated"`
}

func (s QueryYuqingMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryYuqingMessageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYuqingMessageResponseBody) SetRequestId(v string) *QueryYuqingMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryYuqingMessageResponseBody) SetTotalCount(v int64) *QueryYuqingMessageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryYuqingMessageResponseBody) SetYuqingMessages(v []*YuqingMessage) *QueryYuqingMessageResponseBody {
	s.YuqingMessages = v
	return s
}

type QueryYuqingMessageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYuqingMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYuqingMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryYuqingMessageResponse) GoString() string {
	return s.String()
}

func (s *QueryYuqingMessageResponse) SetHeaders(v map[string]*string) *QueryYuqingMessageResponse {
	s.Headers = v
	return s
}

func (s *QueryYuqingMessageResponse) SetStatusCode(v int32) *QueryYuqingMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYuqingMessageResponse) SetBody(v *QueryYuqingMessageResponseBody) *QueryYuqingMessageResponse {
	s.Body = v
	return s
}

type SubmitAnalysisTaskRequest struct {
	AnalyseType     *string          `json:"analyseType,omitempty" xml:"analyseType,omitempty"`
	RequestId       *string          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SearchCondition *SearchCondition `json:"searchCondition,omitempty" xml:"searchCondition,omitempty"`
	TeamHashId      *string          `json:"teamHashId,omitempty" xml:"teamHashId,omitempty"`
}

func (s SubmitAnalysisTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisTaskRequest) SetAnalyseType(v string) *SubmitAnalysisTaskRequest {
	s.AnalyseType = &v
	return s
}

func (s *SubmitAnalysisTaskRequest) SetRequestId(v string) *SubmitAnalysisTaskRequest {
	s.RequestId = &v
	return s
}

func (s *SubmitAnalysisTaskRequest) SetSearchCondition(v *SearchCondition) *SubmitAnalysisTaskRequest {
	s.SearchCondition = v
	return s
}

func (s *SubmitAnalysisTaskRequest) SetTeamHashId(v string) *SubmitAnalysisTaskRequest {
	s.TeamHashId = &v
	return s
}

type SubmitAnalysisTaskResponseBody struct {
	AnalysisId *int64  `json:"analysisId,omitempty" xml:"analysisId,omitempty"`
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResultJson *string `json:"resultJson,omitempty" xml:"resultJson,omitempty"`
}

func (s SubmitAnalysisTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisTaskResponseBody) SetAnalysisId(v int64) *SubmitAnalysisTaskResponseBody {
	s.AnalysisId = &v
	return s
}

func (s *SubmitAnalysisTaskResponseBody) SetRequestId(v string) *SubmitAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAnalysisTaskResponseBody) SetResultJson(v string) *SubmitAnalysisTaskResponseBody {
	s.ResultJson = &v
	return s
}

type SubmitAnalysisTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAnalysisTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisTaskResponse) SetHeaders(v map[string]*string) *SubmitAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitAnalysisTaskResponse) SetStatusCode(v int32) *SubmitAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAnalysisTaskResponse) SetBody(v *SubmitAnalysisTaskResponseBody) *SubmitAnalysisTaskResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("yuqing"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关闭舆情产品
//
// @param request - CloseProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseProductResponse
func (client *Client) CloseProductWithOptions(request *CloseProductRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloseProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["requestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductInstance)) {
		body["productInstance"] = request.ProductInstance
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseProduct"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/aliyun/closeProduct.json"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭舆情产品
//
// @param request - CloseProductRequest
//
// @return CloseProductResponse
func (client *Client) CloseProduct(request *CloseProductRequest) (_result *CloseProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloseProductResponse{}
	_body, _err := client.CloseProductWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 控制台统一代理API
//
// @param request - ConsoleApiProxyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConsoleApiProxyResponse
func (client *Client) ConsoleApiProxyWithOptions(request *ConsoleApiProxyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ConsoleApiProxyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConsoleApiProxy"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/aliyun/consoleApiProxy.json"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ConsoleApiProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 控制台统一代理API
//
// @param request - ConsoleApiProxyRequest
//
// @return ConsoleApiProxyResponse
func (client *Client) ConsoleApiProxy(request *ConsoleApiProxyRequest) (_result *ConsoleApiProxyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ConsoleApiProxyResponse{}
	_body, _err := client.ConsoleApiProxyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ConsoleProxy is deprecated
//
// Summary:
//
// 控制台统一代理API
//
// @param request - ConsoleProxyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConsoleProxyResponse
// Deprecated
func (client *Client) ConsoleProxyWithOptions(request *ConsoleProxyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ConsoleProxyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["requestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCode)) {
		body["appCode"] = request.AppCode
	}

	if !tea.BoolValue(util.IsUnset(request.InterfaceName)) {
		body["interfaceName"] = request.InterfaceName
	}

	if !tea.BoolValue(util.IsUnset(request.ParamJson)) {
		body["paramJson"] = request.ParamJson
	}

	if !tea.BoolValue(util.IsUnset(request.TeamHashId)) {
		body["teamHashId"] = request.TeamHashId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConsoleProxy"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/aliyun/consoleProxy.json"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConsoleProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ConsoleProxy is deprecated
//
// Summary:
//
// 控制台统一代理API
//
// @param request - ConsoleProxyRequest
//
// @return ConsoleProxyResponse
// Deprecated
func (client *Client) ConsoleProxy(request *ConsoleProxyRequest) (_result *ConsoleProxyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ConsoleProxyResponse{}
	_body, _err := client.ConsoleProxyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 读取分析组件计算任务结果
//
// @param request - GetAnalysisTaskResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAnalysisTaskResultResponse
func (client *Client) GetAnalysisTaskResultWithOptions(request *GetAnalysisTaskResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAnalysisTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnalysisId)) {
		query["analysisId"] = request.AnalysisId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.TeamHashId)) {
		query["teamHashId"] = request.TeamHashId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAnalysisTaskResult"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/aliyun/getAnalysisComponentResult.json"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAnalysisTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 读取分析组件计算任务结果
//
// @param request - GetAnalysisTaskResultRequest
//
// @return GetAnalysisTaskResultResponse
func (client *Client) GetAnalysisTaskResult(request *GetAnalysisTaskResultRequest) (_result *GetAnalysisTaskResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAnalysisTaskResultResponse{}
	_body, _err := client.GetAnalysisTaskResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开通舆情产品
//
// @param request - OpenProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenProductResponse
func (client *Client) OpenProductWithOptions(request *OpenProductRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *OpenProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["requestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductInstance)) {
		body["productInstance"] = request.ProductInstance
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenProduct"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/aliyun/openProduct.json"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开通舆情产品
//
// @param request - OpenProductRequest
//
// @return OpenProductResponse
func (client *Client) OpenProduct(request *OpenProductRequest) (_result *OpenProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OpenProductResponse{}
	_body, _err := client.OpenProductWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询产品开通实例列表
//
// @param request - QueryProductInstanceListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProductInstanceListResponse
func (client *Client) QueryProductInstanceListWithOptions(request *QueryProductInstanceListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryProductInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCode)) {
		query["appCode"] = request.AppCode
	}

	if !tea.BoolValue(util.IsUnset(request.FromTime)) {
		query["fromTime"] = request.FromTime
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantUid)) {
		query["tenantUid"] = request.TenantUid
	}

	if !tea.BoolValue(util.IsUnset(request.ToTime)) {
		query["toTime"] = request.ToTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryProductInstanceList"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/aliyun/queryProductInstanceList.json"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryProductInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询产品开通实例列表
//
// @param request - QueryProductInstanceListRequest
//
// @return QueryProductInstanceListResponse
func (client *Client) QueryProductInstanceList(request *QueryProductInstanceListRequest) (_result *QueryProductInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryProductInstanceListResponse{}
	_body, _err := client.QueryProductInstanceListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询舆情文章列表
//
// @param request - QueryYuqingMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYuqingMessageResponse
func (client *Client) QueryYuqingMessageWithOptions(request *QueryYuqingMessageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryYuqingMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["requestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SearchCondition)) {
		body["searchCondition"] = request.SearchCondition
	}

	if !tea.BoolValue(util.IsUnset(request.TeamHashId)) {
		body["teamHashId"] = request.TeamHashId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryYuqingMessage"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/aliyun/queryYuqingMessage.json"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryYuqingMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询舆情文章列表
//
// @param request - QueryYuqingMessageRequest
//
// @return QueryYuqingMessageResponse
func (client *Client) QueryYuqingMessage(request *QueryYuqingMessageRequest) (_result *QueryYuqingMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryYuqingMessageResponse{}
	_body, _err := client.QueryYuqingMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交分析组件计算任务
//
// @param request - SubmitAnalysisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAnalysisTaskResponse
func (client *Client) SubmitAnalysisTaskWithOptions(request *SubmitAnalysisTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitAnalysisTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["requestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnalyseType)) {
		body["analyseType"] = request.AnalyseType
	}

	if !tea.BoolValue(util.IsUnset(request.SearchCondition)) {
		body["searchCondition"] = request.SearchCondition
	}

	if !tea.BoolValue(util.IsUnset(request.TeamHashId)) {
		body["teamHashId"] = request.TeamHashId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitAnalysisTask"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/aliyun/submitAnalysisComponent.json"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitAnalysisTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交分析组件计算任务
//
// @param request - SubmitAnalysisTaskRequest
//
// @return SubmitAnalysisTaskResponse
func (client *Client) SubmitAnalysisTask(request *SubmitAnalysisTaskRequest) (_result *SubmitAnalysisTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitAnalysisTaskResponse{}
	_body, _err := client.SubmitAnalysisTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
