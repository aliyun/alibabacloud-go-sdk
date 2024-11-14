// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AnchorResponse struct {
	AnchorId           *string `json:"anchorId,omitempty" xml:"anchorId,omitempty"`
	AnchorMaterialName *string `json:"anchorMaterialName,omitempty" xml:"anchorMaterialName,omitempty"`
	AnchorType         *string `json:"anchorType,omitempty" xml:"anchorType,omitempty"`
	CoverHeight        *int32  `json:"coverHeight,omitempty" xml:"coverHeight,omitempty"`
	CoverRate          *string `json:"coverRate,omitempty" xml:"coverRate,omitempty"`
	CoverThumbnailUrl  *string `json:"coverThumbnailUrl,omitempty" xml:"coverThumbnailUrl,omitempty"`
	CoverUrl           *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	CoverWeight        *int32  `json:"coverWeight,omitempty" xml:"coverWeight,omitempty"`
	DigitalHumanType   *string `json:"digitalHumanType,omitempty" xml:"digitalHumanType,omitempty"`
	Gender             *string `json:"gender,omitempty" xml:"gender,omitempty"`
	ResourceTypeDesc   *string `json:"resourceTypeDesc,omitempty" xml:"resourceTypeDesc,omitempty"`
	Status             *string `json:"status,omitempty" xml:"status,omitempty"`
	SupportBgChange    *int32  `json:"supportBgChange,omitempty" xml:"supportBgChange,omitempty"`
	UseScene           *string `json:"useScene,omitempty" xml:"useScene,omitempty"`
}

func (s AnchorResponse) String() string {
	return tea.Prettify(s)
}

func (s AnchorResponse) GoString() string {
	return s.String()
}

func (s *AnchorResponse) SetAnchorId(v string) *AnchorResponse {
	s.AnchorId = &v
	return s
}

func (s *AnchorResponse) SetAnchorMaterialName(v string) *AnchorResponse {
	s.AnchorMaterialName = &v
	return s
}

func (s *AnchorResponse) SetAnchorType(v string) *AnchorResponse {
	s.AnchorType = &v
	return s
}

func (s *AnchorResponse) SetCoverHeight(v int32) *AnchorResponse {
	s.CoverHeight = &v
	return s
}

func (s *AnchorResponse) SetCoverRate(v string) *AnchorResponse {
	s.CoverRate = &v
	return s
}

func (s *AnchorResponse) SetCoverThumbnailUrl(v string) *AnchorResponse {
	s.CoverThumbnailUrl = &v
	return s
}

func (s *AnchorResponse) SetCoverUrl(v string) *AnchorResponse {
	s.CoverUrl = &v
	return s
}

func (s *AnchorResponse) SetCoverWeight(v int32) *AnchorResponse {
	s.CoverWeight = &v
	return s
}

func (s *AnchorResponse) SetDigitalHumanType(v string) *AnchorResponse {
	s.DigitalHumanType = &v
	return s
}

func (s *AnchorResponse) SetGender(v string) *AnchorResponse {
	s.Gender = &v
	return s
}

func (s *AnchorResponse) SetResourceTypeDesc(v string) *AnchorResponse {
	s.ResourceTypeDesc = &v
	return s
}

func (s *AnchorResponse) SetStatus(v string) *AnchorResponse {
	s.Status = &v
	return s
}

func (s *AnchorResponse) SetSupportBgChange(v int32) *AnchorResponse {
	s.SupportBgChange = &v
	return s
}

func (s *AnchorResponse) SetUseScene(v string) *AnchorResponse {
	s.UseScene = &v
	return s
}

type GetOssUploadTokenResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	UploadInfo *UploadInfo `json:"uploadInfo,omitempty" xml:"uploadInfo,omitempty"`
}

func (s GetOssUploadTokenResult) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadTokenResult) GoString() string {
	return s.String()
}

func (s *GetOssUploadTokenResult) SetRequestId(v string) *GetOssUploadTokenResult {
	s.RequestId = &v
	return s
}

func (s *GetOssUploadTokenResult) SetUploadInfo(v *UploadInfo) *GetOssUploadTokenResult {
	s.UploadInfo = v
	return s
}

type Illustration struct {
	IllustrationId *int64  `json:"illustrationId,omitempty" xml:"illustrationId,omitempty"`
	Oss            *string `json:"oss,omitempty" xml:"oss,omitempty"`
}

func (s Illustration) String() string {
	return tea.Prettify(s)
}

func (s Illustration) GoString() string {
	return s.String()
}

func (s *Illustration) SetIllustrationId(v int64) *Illustration {
	s.IllustrationId = &v
	return s
}

func (s *Illustration) SetOss(v string) *Illustration {
	s.Oss = &v
	return s
}

type IllustrationResult struct {
	Illustration *Illustration `json:"illustration,omitempty" xml:"illustration,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s IllustrationResult) String() string {
	return tea.Prettify(s)
}

func (s IllustrationResult) GoString() string {
	return s.String()
}

func (s *IllustrationResult) SetIllustration(v *Illustration) *IllustrationResult {
	s.Illustration = v
	return s
}

func (s *IllustrationResult) SetRequestId(v string) *IllustrationResult {
	s.RequestId = &v
	return s
}

type IllustrationTask struct {
	GmtCreate          *string  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified        *string  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	IllustrationIds    []*int64 `json:"illustrationIds,omitempty" xml:"illustrationIds,omitempty" type:"Repeated"`
	IllustrationTaskId *int64   `json:"illustrationTaskId,omitempty" xml:"illustrationTaskId,omitempty"`
	// example:
	//
	// Success
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
	TextId     *int64  `json:"textId,omitempty" xml:"textId,omitempty"`
}

func (s IllustrationTask) String() string {
	return tea.Prettify(s)
}

func (s IllustrationTask) GoString() string {
	return s.String()
}

func (s *IllustrationTask) SetGmtCreate(v string) *IllustrationTask {
	s.GmtCreate = &v
	return s
}

func (s *IllustrationTask) SetGmtModified(v string) *IllustrationTask {
	s.GmtModified = &v
	return s
}

func (s *IllustrationTask) SetIllustrationIds(v []*int64) *IllustrationTask {
	s.IllustrationIds = v
	return s
}

func (s *IllustrationTask) SetIllustrationTaskId(v int64) *IllustrationTask {
	s.IllustrationTaskId = &v
	return s
}

func (s *IllustrationTask) SetTaskStatus(v string) *IllustrationTask {
	s.TaskStatus = &v
	return s
}

func (s *IllustrationTask) SetTextId(v int64) *IllustrationTask {
	s.TextId = &v
	return s
}

type IllustrationTaskCreateCmd struct {
	// example:
	//
	// 0-不换背景，1-换背景
	BackgroundType *int32 `json:"backgroundType,omitempty" xml:"backgroundType,omitempty"`
	// example:
	//
	// 1024
	DstHeight *int32 `json:"dstHeight,omitempty" xml:"dstHeight,omitempty"`
	// example:
	//
	// 1024
	DstWidth *int32 `json:"dstWidth,omitempty" xml:"dstWidth,omitempty"`
	// example:
	//
	// 28274623764834
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	// example:
	//
	// 4
	Nums        *int32    `json:"nums,omitempty" xml:"nums,omitempty"`
	OssPaths    []*string `json:"ossPaths,omitempty" xml:"ossPaths,omitempty" type:"Repeated"`
	StickerText *string   `json:"stickerText,omitempty" xml:"stickerText,omitempty"`
}

func (s IllustrationTaskCreateCmd) String() string {
	return tea.Prettify(s)
}

func (s IllustrationTaskCreateCmd) GoString() string {
	return s.String()
}

func (s *IllustrationTaskCreateCmd) SetBackgroundType(v int32) *IllustrationTaskCreateCmd {
	s.BackgroundType = &v
	return s
}

func (s *IllustrationTaskCreateCmd) SetDstHeight(v int32) *IllustrationTaskCreateCmd {
	s.DstHeight = &v
	return s
}

func (s *IllustrationTaskCreateCmd) SetDstWidth(v int32) *IllustrationTaskCreateCmd {
	s.DstWidth = &v
	return s
}

func (s *IllustrationTaskCreateCmd) SetIdempotentId(v string) *IllustrationTaskCreateCmd {
	s.IdempotentId = &v
	return s
}

func (s *IllustrationTaskCreateCmd) SetNums(v int32) *IllustrationTaskCreateCmd {
	s.Nums = &v
	return s
}

func (s *IllustrationTaskCreateCmd) SetOssPaths(v []*string) *IllustrationTaskCreateCmd {
	s.OssPaths = v
	return s
}

func (s *IllustrationTaskCreateCmd) SetStickerText(v string) *IllustrationTaskCreateCmd {
	s.StickerText = &v
	return s
}

type IllustrationTaskResult struct {
	IllustrationTask *IllustrationTask `json:"illustrationTask,omitempty" xml:"illustrationTask,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s IllustrationTaskResult) String() string {
	return tea.Prettify(s)
}

func (s IllustrationTaskResult) GoString() string {
	return s.String()
}

func (s *IllustrationTaskResult) SetIllustrationTask(v *IllustrationTask) *IllustrationTaskResult {
	s.IllustrationTask = v
	return s
}

func (s *IllustrationTaskResult) SetRequestId(v string) *IllustrationTaskResult {
	s.RequestId = &v
	return s
}

type ReferenceTag struct {
	ReferenceContent *string `json:"referenceContent,omitempty" xml:"referenceContent,omitempty"`
	ReferenceTitle   *string `json:"referenceTitle,omitempty" xml:"referenceTitle,omitempty"`
}

func (s ReferenceTag) String() string {
	return tea.Prettify(s)
}

func (s ReferenceTag) GoString() string {
	return s.String()
}

func (s *ReferenceTag) SetReferenceContent(v string) *ReferenceTag {
	s.ReferenceContent = &v
	return s
}

func (s *ReferenceTag) SetReferenceTitle(v string) *ReferenceTag {
	s.ReferenceTitle = &v
	return s
}

type Text struct {
	AgentId   *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AgentName *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	Desc                   *string  `json:"desc,omitempty" xml:"desc,omitempty"`
	ErrMsg                 *string  `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	GmtCreate              *string  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified            *string  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	IllustrationTaskIdList []*int64 `json:"illustrationTaskIdList,omitempty" xml:"illustrationTaskIdList,omitempty" type:"Repeated"`
	PublishStatus          *string  `json:"publishStatus,omitempty" xml:"publishStatus,omitempty"`
	TextContent            *string  `json:"textContent,omitempty" xml:"textContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TextId *int64 `json:"textId,omitempty" xml:"textId,omitempty"`
	// example:
	//
	// true
	TextIllustrationTag *bool   `json:"textIllustrationTag,omitempty" xml:"textIllustrationTag,omitempty"`
	TextModeType        *string `json:"textModeType,omitempty" xml:"textModeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Generating
	TextStatus    *string `json:"textStatus,omitempty" xml:"textStatus,omitempty"`
	TextStyleType *string `json:"textStyleType,omitempty" xml:"textStyleType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	TextTaskId *int64    `json:"textTaskId,omitempty" xml:"textTaskId,omitempty"`
	TextThemes []*string `json:"textThemes,omitempty" xml:"textThemes,omitempty" type:"Repeated"`
	// example:
	//
	// xxx
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	UserNameCreate *string `json:"userNameCreate,omitempty" xml:"userNameCreate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	UserNameModified *string `json:"userNameModified,omitempty" xml:"userNameModified,omitempty"`
}

func (s Text) String() string {
	return tea.Prettify(s)
}

func (s Text) GoString() string {
	return s.String()
}

func (s *Text) SetAgentId(v string) *Text {
	s.AgentId = &v
	return s
}

func (s *Text) SetAgentName(v string) *Text {
	s.AgentName = &v
	return s
}

func (s *Text) SetDesc(v string) *Text {
	s.Desc = &v
	return s
}

func (s *Text) SetErrMsg(v string) *Text {
	s.ErrMsg = &v
	return s
}

func (s *Text) SetGmtCreate(v string) *Text {
	s.GmtCreate = &v
	return s
}

func (s *Text) SetGmtModified(v string) *Text {
	s.GmtModified = &v
	return s
}

func (s *Text) SetIllustrationTaskIdList(v []*int64) *Text {
	s.IllustrationTaskIdList = v
	return s
}

func (s *Text) SetPublishStatus(v string) *Text {
	s.PublishStatus = &v
	return s
}

func (s *Text) SetTextContent(v string) *Text {
	s.TextContent = &v
	return s
}

func (s *Text) SetTextId(v int64) *Text {
	s.TextId = &v
	return s
}

func (s *Text) SetTextIllustrationTag(v bool) *Text {
	s.TextIllustrationTag = &v
	return s
}

func (s *Text) SetTextModeType(v string) *Text {
	s.TextModeType = &v
	return s
}

func (s *Text) SetTextStatus(v string) *Text {
	s.TextStatus = &v
	return s
}

func (s *Text) SetTextStyleType(v string) *Text {
	s.TextStyleType = &v
	return s
}

func (s *Text) SetTextTaskId(v int64) *Text {
	s.TextTaskId = &v
	return s
}

func (s *Text) SetTextThemes(v []*string) *Text {
	s.TextThemes = v
	return s
}

func (s *Text) SetTitle(v string) *Text {
	s.Title = &v
	return s
}

func (s *Text) SetUserNameCreate(v string) *Text {
	s.UserNameCreate = &v
	return s
}

func (s *Text) SetUserNameModified(v string) *Text {
	s.UserNameModified = &v
	return s
}

type TextQueryResult struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Texts     []*Text `json:"texts,omitempty" xml:"texts,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s TextQueryResult) String() string {
	return tea.Prettify(s)
}

func (s TextQueryResult) GoString() string {
	return s.String()
}

func (s *TextQueryResult) SetRequestId(v string) *TextQueryResult {
	s.RequestId = &v
	return s
}

func (s *TextQueryResult) SetTexts(v []*Text) *TextQueryResult {
	s.Texts = v
	return s
}

func (s *TextQueryResult) SetTotal(v int32) *TextQueryResult {
	s.Total = &v
	return s
}

type TextResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Text *Text `json:"text,omitempty" xml:"text,omitempty"`
}

func (s TextResult) String() string {
	return tea.Prettify(s)
}

func (s TextResult) GoString() string {
	return s.String()
}

func (s *TextResult) SetRequestId(v string) *TextResult {
	s.RequestId = &v
	return s
}

func (s *TextResult) SetText(v *Text) *TextResult {
	s.Text = v
	return s
}

type TextTask struct {
	AgentId   *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AgentName *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
	// example:
	//
	// 九寨沟三日游攻略
	ContentRequirement *string `json:"contentRequirement,omitempty" xml:"contentRequirement,omitempty"`
	GmtCreate          *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified        *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Introduction       *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Nums *int32 `json:"nums,omitempty" xml:"nums,omitempty"`
	// example:
	//
	// xxx
	Point        *string       `json:"point,omitempty" xml:"point,omitempty"`
	ReferenceTag *ReferenceTag `json:"referenceTag,omitempty" xml:"referenceTag,omitempty"`
	RelatedRagId *int32        `json:"relatedRagId,omitempty" xml:"relatedRagId,omitempty"`
	// This parameter is required.
	Style   *string  `json:"style,omitempty" xml:"style,omitempty"`
	Target  *string  `json:"target,omitempty" xml:"target,omitempty"`
	TextIds []*int64 `json:"textIds,omitempty" xml:"textIds,omitempty" type:"Repeated"`
	// This parameter is required.
	TextModeType   *string `json:"textModeType,omitempty" xml:"textModeType,omitempty"`
	TextTaskId     *int64  `json:"textTaskId,omitempty" xml:"textTaskId,omitempty"`
	TextTaskStatus *string `json:"textTaskStatus,omitempty" xml:"textTaskStatus,omitempty"`
	Texts          []*Text `json:"texts,omitempty" xml:"texts,omitempty" type:"Repeated"`
	// example:
	//
	// 旅游路线
	Theme     *string `json:"theme,omitempty" xml:"theme,omitempty"`
	ThemeDesc *string `json:"themeDesc,omitempty" xml:"themeDesc,omitempty"`
}

func (s TextTask) String() string {
	return tea.Prettify(s)
}

func (s TextTask) GoString() string {
	return s.String()
}

func (s *TextTask) SetAgentId(v string) *TextTask {
	s.AgentId = &v
	return s
}

func (s *TextTask) SetAgentName(v string) *TextTask {
	s.AgentName = &v
	return s
}

func (s *TextTask) SetContentRequirement(v string) *TextTask {
	s.ContentRequirement = &v
	return s
}

func (s *TextTask) SetGmtCreate(v string) *TextTask {
	s.GmtCreate = &v
	return s
}

func (s *TextTask) SetGmtModified(v string) *TextTask {
	s.GmtModified = &v
	return s
}

func (s *TextTask) SetIntroduction(v string) *TextTask {
	s.Introduction = &v
	return s
}

func (s *TextTask) SetNums(v int32) *TextTask {
	s.Nums = &v
	return s
}

func (s *TextTask) SetPoint(v string) *TextTask {
	s.Point = &v
	return s
}

func (s *TextTask) SetReferenceTag(v *ReferenceTag) *TextTask {
	s.ReferenceTag = v
	return s
}

func (s *TextTask) SetRelatedRagId(v int32) *TextTask {
	s.RelatedRagId = &v
	return s
}

func (s *TextTask) SetStyle(v string) *TextTask {
	s.Style = &v
	return s
}

func (s *TextTask) SetTarget(v string) *TextTask {
	s.Target = &v
	return s
}

func (s *TextTask) SetTextIds(v []*int64) *TextTask {
	s.TextIds = v
	return s
}

func (s *TextTask) SetTextModeType(v string) *TextTask {
	s.TextModeType = &v
	return s
}

func (s *TextTask) SetTextTaskId(v int64) *TextTask {
	s.TextTaskId = &v
	return s
}

func (s *TextTask) SetTextTaskStatus(v string) *TextTask {
	s.TextTaskStatus = &v
	return s
}

func (s *TextTask) SetTexts(v []*Text) *TextTask {
	s.Texts = v
	return s
}

func (s *TextTask) SetTheme(v string) *TextTask {
	s.Theme = &v
	return s
}

func (s *TextTask) SetThemeDesc(v string) *TextTask {
	s.ThemeDesc = &v
	return s
}

type TextTaskCreateCmd struct {
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// 极氪007新车上市
	ContentRequirement *string `json:"contentRequirement,omitempty" xml:"contentRequirement,omitempty"`
	// example:
	//
	// 28274623764834
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	Industry     *string `json:"industry,omitempty" xml:"industry,omitempty"`
	// example:
	//
	// xxx
	Introduction *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// example:
	//
	// 超强续航
	Point        *string       `json:"point,omitempty" xml:"point,omitempty"`
	ReferenceTag *ReferenceTag `json:"referenceTag,omitempty" xml:"referenceTag,omitempty"`
	// example:
	//
	// 1
	RelatedRagIds []*int64 `json:"relatedRagIds,omitempty" xml:"relatedRagIds,omitempty" type:"Repeated"`
	// example:
	//
	// true
	StreamApi *bool `json:"streamApi,omitempty" xml:"streamApi,omitempty"`
	// This parameter is required.
	Style  *string `json:"style,omitempty" xml:"style,omitempty"`
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
	// This parameter is required.
	TextModeType *string `json:"textModeType,omitempty" xml:"textModeType,omitempty"`
	// example:
	//
	// 旅游路线
	Theme  *string   `json:"theme,omitempty" xml:"theme,omitempty"`
	Themes []*string `json:"themes,omitempty" xml:"themes,omitempty" type:"Repeated"`
}

func (s TextTaskCreateCmd) String() string {
	return tea.Prettify(s)
}

func (s TextTaskCreateCmd) GoString() string {
	return s.String()
}

func (s *TextTaskCreateCmd) SetAgentId(v string) *TextTaskCreateCmd {
	s.AgentId = &v
	return s
}

func (s *TextTaskCreateCmd) SetContentRequirement(v string) *TextTaskCreateCmd {
	s.ContentRequirement = &v
	return s
}

func (s *TextTaskCreateCmd) SetIdempotentId(v string) *TextTaskCreateCmd {
	s.IdempotentId = &v
	return s
}

func (s *TextTaskCreateCmd) SetIndustry(v string) *TextTaskCreateCmd {
	s.Industry = &v
	return s
}

func (s *TextTaskCreateCmd) SetIntroduction(v string) *TextTaskCreateCmd {
	s.Introduction = &v
	return s
}

func (s *TextTaskCreateCmd) SetNumber(v int32) *TextTaskCreateCmd {
	s.Number = &v
	return s
}

func (s *TextTaskCreateCmd) SetPoint(v string) *TextTaskCreateCmd {
	s.Point = &v
	return s
}

func (s *TextTaskCreateCmd) SetReferenceTag(v *ReferenceTag) *TextTaskCreateCmd {
	s.ReferenceTag = v
	return s
}

func (s *TextTaskCreateCmd) SetRelatedRagIds(v []*int64) *TextTaskCreateCmd {
	s.RelatedRagIds = v
	return s
}

func (s *TextTaskCreateCmd) SetStreamApi(v bool) *TextTaskCreateCmd {
	s.StreamApi = &v
	return s
}

func (s *TextTaskCreateCmd) SetStyle(v string) *TextTaskCreateCmd {
	s.Style = &v
	return s
}

func (s *TextTaskCreateCmd) SetTarget(v string) *TextTaskCreateCmd {
	s.Target = &v
	return s
}

func (s *TextTaskCreateCmd) SetTextModeType(v string) *TextTaskCreateCmd {
	s.TextModeType = &v
	return s
}

func (s *TextTaskCreateCmd) SetTheme(v string) *TextTaskCreateCmd {
	s.Theme = &v
	return s
}

func (s *TextTaskCreateCmd) SetThemes(v []*string) *TextTaskCreateCmd {
	s.Themes = v
	return s
}

type TextTaskResult struct {
	TextTask *TextTask `json:"textTask,omitempty" xml:"textTask,omitempty"`
}

func (s TextTaskResult) String() string {
	return tea.Prettify(s)
}

func (s TextTaskResult) GoString() string {
	return s.String()
}

func (s *TextTaskResult) SetTextTask(v *TextTask) *TextTaskResult {
	s.TextTask = v
	return s
}

type TextTheme struct {
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TextTheme) String() string {
	return tea.Prettify(s)
}

func (s TextTheme) GoString() string {
	return s.String()
}

func (s *TextTheme) SetDesc(v string) *TextTheme {
	s.Desc = &v
	return s
}

func (s *TextTheme) SetName(v string) *TextTheme {
	s.Name = &v
	return s
}

type TextThemeListResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	TextThemeList []*TextTheme `json:"textThemeList,omitempty" xml:"textThemeList,omitempty" type:"Repeated"`
}

func (s TextThemeListResult) String() string {
	return tea.Prettify(s)
}

func (s TextThemeListResult) GoString() string {
	return s.String()
}

func (s *TextThemeListResult) SetRequestId(v string) *TextThemeListResult {
	s.RequestId = &v
	return s
}

func (s *TextThemeListResult) SetTextThemeList(v []*TextTheme) *TextThemeListResult {
	s.TextThemeList = v
	return s
}

type UploadInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxxx
	AccessId *string `json:"accessId,omitempty" xml:"accessId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yic-pre.oss-cn-hangzhou.aliyuncs.com
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234/temp-novels/xxxx-xxx-xx.txt
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxxxxx
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxxxx
	PolicySignature *string `json:"policySignature,omitempty" xml:"policySignature,omitempty"`
	// example:
	//
	// xxxxxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s UploadInfo) String() string {
	return tea.Prettify(s)
}

func (s UploadInfo) GoString() string {
	return s.String()
}

func (s *UploadInfo) SetAccessId(v string) *UploadInfo {
	s.AccessId = &v
	return s
}

func (s *UploadInfo) SetHost(v string) *UploadInfo {
	s.Host = &v
	return s
}

func (s *UploadInfo) SetKey(v string) *UploadInfo {
	s.Key = &v
	return s
}

func (s *UploadInfo) SetPolicy(v string) *UploadInfo {
	s.Policy = &v
	return s
}

func (s *UploadInfo) SetPolicySignature(v string) *UploadInfo {
	s.PolicySignature = &v
	return s
}

func (s *UploadInfo) SetUrl(v string) *UploadInfo {
	s.Url = &v
	return s
}

type VoiceModelResponse struct {
	ResourceTypeDesc *string `json:"resourceTypeDesc,omitempty" xml:"resourceTypeDesc,omitempty"`
	TtsVersion       *int32  `json:"ttsVersion,omitempty" xml:"ttsVersion,omitempty"`
	UseScene         *string `json:"useScene,omitempty" xml:"useScene,omitempty"`
	VoiceDesc        *string `json:"voiceDesc,omitempty" xml:"voiceDesc,omitempty"`
	VoiceGender      *string `json:"voiceGender,omitempty" xml:"voiceGender,omitempty"`
	VoiceId          *int64  `json:"voiceId,omitempty" xml:"voiceId,omitempty"`
	VoiceLanguage    *string `json:"voiceLanguage,omitempty" xml:"voiceLanguage,omitempty"`
	VoiceModel       *string `json:"voiceModel,omitempty" xml:"voiceModel,omitempty"`
	VoiceName        *string `json:"voiceName,omitempty" xml:"voiceName,omitempty"`
	VoiceType        *string `json:"voiceType,omitempty" xml:"voiceType,omitempty"`
	VoiceUrl         *string `json:"voiceUrl,omitempty" xml:"voiceUrl,omitempty"`
}

func (s VoiceModelResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceModelResponse) GoString() string {
	return s.String()
}

func (s *VoiceModelResponse) SetResourceTypeDesc(v string) *VoiceModelResponse {
	s.ResourceTypeDesc = &v
	return s
}

func (s *VoiceModelResponse) SetTtsVersion(v int32) *VoiceModelResponse {
	s.TtsVersion = &v
	return s
}

func (s *VoiceModelResponse) SetUseScene(v string) *VoiceModelResponse {
	s.UseScene = &v
	return s
}

func (s *VoiceModelResponse) SetVoiceDesc(v string) *VoiceModelResponse {
	s.VoiceDesc = &v
	return s
}

func (s *VoiceModelResponse) SetVoiceGender(v string) *VoiceModelResponse {
	s.VoiceGender = &v
	return s
}

func (s *VoiceModelResponse) SetVoiceId(v int64) *VoiceModelResponse {
	s.VoiceId = &v
	return s
}

func (s *VoiceModelResponse) SetVoiceLanguage(v string) *VoiceModelResponse {
	s.VoiceLanguage = &v
	return s
}

func (s *VoiceModelResponse) SetVoiceModel(v string) *VoiceModelResponse {
	s.VoiceModel = &v
	return s
}

func (s *VoiceModelResponse) SetVoiceName(v string) *VoiceModelResponse {
	s.VoiceName = &v
	return s
}

func (s *VoiceModelResponse) SetVoiceType(v string) *VoiceModelResponse {
	s.VoiceType = &v
	return s
}

func (s *VoiceModelResponse) SetVoiceUrl(v string) *VoiceModelResponse {
	s.VoiceUrl = &v
	return s
}

type AddTextFeedbackRequest struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 1
	Quality *int32 `json:"quality,omitempty" xml:"quality,omitempty"`
	// example:
	//
	// 8478
	TextId *int64 `json:"textId,omitempty" xml:"textId,omitempty"`
}

func (s AddTextFeedbackRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTextFeedbackRequest) GoString() string {
	return s.String()
}

func (s *AddTextFeedbackRequest) SetContent(v string) *AddTextFeedbackRequest {
	s.Content = &v
	return s
}

func (s *AddTextFeedbackRequest) SetQuality(v int32) *AddTextFeedbackRequest {
	s.Quality = &v
	return s
}

func (s *AddTextFeedbackRequest) SetTextId(v int64) *AddTextFeedbackRequest {
	s.TextId = &v
	return s
}

type AddTextFeedbackResponseBody struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddTextFeedbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTextFeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *AddTextFeedbackResponseBody) SetRequestId(v string) *AddTextFeedbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTextFeedbackResponseBody) SetSuccess(v bool) *AddTextFeedbackResponseBody {
	s.Success = &v
	return s
}

type AddTextFeedbackResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTextFeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTextFeedbackResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTextFeedbackResponse) GoString() string {
	return s.String()
}

func (s *AddTextFeedbackResponse) SetHeaders(v map[string]*string) *AddTextFeedbackResponse {
	s.Headers = v
	return s
}

func (s *AddTextFeedbackResponse) SetStatusCode(v int32) *AddTextFeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTextFeedbackResponse) SetBody(v *AddTextFeedbackResponseBody) *AddTextFeedbackResponse {
	s.Body = v
	return s
}

type BatchGetProjectTaskRequest struct {
	TaskIdList []*string `json:"taskIdList,omitempty" xml:"taskIdList,omitempty" type:"Repeated"`
}

func (s BatchGetProjectTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetProjectTaskRequest) GoString() string {
	return s.String()
}

func (s *BatchGetProjectTaskRequest) SetTaskIdList(v []*string) *BatchGetProjectTaskRequest {
	s.TaskIdList = v
	return s
}

type BatchGetProjectTaskShrinkRequest struct {
	TaskIdListShrink *string `json:"taskIdList,omitempty" xml:"taskIdList,omitempty"`
}

func (s BatchGetProjectTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetProjectTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetProjectTaskShrinkRequest) SetTaskIdListShrink(v string) *BatchGetProjectTaskShrinkRequest {
	s.TaskIdListShrink = &v
	return s
}

type BatchGetProjectTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 11
	RequestId  *string                                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResultList []*BatchGetProjectTaskResponseBodyResultList `json:"resultList,omitempty" xml:"resultList,omitempty" type:"Repeated"`
}

func (s BatchGetProjectTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetProjectTaskResponseBody) SetRequestId(v string) *BatchGetProjectTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetProjectTaskResponseBody) SetResultList(v []*BatchGetProjectTaskResponseBodyResultList) *BatchGetProjectTaskResponseBody {
	s.ResultList = v
	return s
}

type BatchGetProjectTaskResponseBodyResultList struct {
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 11
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// http
	VideoDownloadUrl *string `json:"videoDownloadUrl,omitempty" xml:"videoDownloadUrl,omitempty"`
	// example:
	//
	// 1000
	VideoDuration *int32 `json:"videoDuration,omitempty" xml:"videoDuration,omitempty"`
	// example:
	//
	// http
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s BatchGetProjectTaskResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s BatchGetProjectTaskResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *BatchGetProjectTaskResponseBodyResultList) SetErrorMsg(v string) *BatchGetProjectTaskResponseBodyResultList {
	s.ErrorMsg = &v
	return s
}

func (s *BatchGetProjectTaskResponseBodyResultList) SetStatus(v string) *BatchGetProjectTaskResponseBodyResultList {
	s.Status = &v
	return s
}

func (s *BatchGetProjectTaskResponseBodyResultList) SetTaskId(v string) *BatchGetProjectTaskResponseBodyResultList {
	s.TaskId = &v
	return s
}

func (s *BatchGetProjectTaskResponseBodyResultList) SetVideoDownloadUrl(v string) *BatchGetProjectTaskResponseBodyResultList {
	s.VideoDownloadUrl = &v
	return s
}

func (s *BatchGetProjectTaskResponseBodyResultList) SetVideoDuration(v int32) *BatchGetProjectTaskResponseBodyResultList {
	s.VideoDuration = &v
	return s
}

func (s *BatchGetProjectTaskResponseBodyResultList) SetVideoUrl(v string) *BatchGetProjectTaskResponseBodyResultList {
	s.VideoUrl = &v
	return s
}

type BatchGetProjectTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetProjectTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetProjectTaskResponse) GoString() string {
	return s.String()
}

func (s *BatchGetProjectTaskResponse) SetHeaders(v map[string]*string) *BatchGetProjectTaskResponse {
	s.Headers = v
	return s
}

func (s *BatchGetProjectTaskResponse) SetStatusCode(v int32) *BatchGetProjectTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetProjectTaskResponse) SetBody(v *BatchGetProjectTaskResponseBody) *BatchGetProjectTaskResponse {
	s.Body = v
	return s
}

type CheckSessionRequest struct {
	// example:
	//
	// 11111
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 121dlsga4o7golrl1hoja
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s CheckSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckSessionRequest) GoString() string {
	return s.String()
}

func (s *CheckSessionRequest) SetProjectId(v string) *CheckSessionRequest {
	s.ProjectId = &v
	return s
}

func (s *CheckSessionRequest) SetSessionId(v string) *CheckSessionRequest {
	s.SessionId = &v
	return s
}

type CheckSessionResponseBody struct {
	// example:
	//
	// 5389BE87-571B-573C-90ED-F07C5E68760B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// FREE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CheckSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSessionResponseBody) SetRequestId(v string) *CheckSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSessionResponseBody) SetStatus(v string) *CheckSessionResponseBody {
	s.Status = &v
	return s
}

type CheckSessionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckSessionResponse) GoString() string {
	return s.String()
}

func (s *CheckSessionResponse) SetHeaders(v map[string]*string) *CheckSessionResponse {
	s.Headers = v
	return s
}

func (s *CheckSessionResponse) SetStatusCode(v int32) *CheckSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSessionResponse) SetBody(v *CheckSessionResponseBody) *CheckSessionResponse {
	s.Body = v
	return s
}

type CloseAICoachTaskSessionRequest struct {
	// example:
	//
	// 11
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 273610276967782972
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CloseAICoachTaskSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseAICoachTaskSessionRequest) GoString() string {
	return s.String()
}

func (s *CloseAICoachTaskSessionRequest) SetSessionId(v string) *CloseAICoachTaskSessionRequest {
	s.SessionId = &v
	return s
}

func (s *CloseAICoachTaskSessionRequest) SetUid(v string) *CloseAICoachTaskSessionRequest {
	s.Uid = &v
	return s
}

type CloseAICoachTaskSessionResponseBody struct {
	// example:
	//
	// 0E06E0AA-D5B6-538C-8CE9-BAB79C68B690
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CloseAICoachTaskSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseAICoachTaskSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CloseAICoachTaskSessionResponseBody) SetRequestId(v string) *CloseAICoachTaskSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseAICoachTaskSessionResponseBody) SetStatus(v string) *CloseAICoachTaskSessionResponseBody {
	s.Status = &v
	return s
}

type CloseAICoachTaskSessionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseAICoachTaskSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseAICoachTaskSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseAICoachTaskSessionResponse) GoString() string {
	return s.String()
}

func (s *CloseAICoachTaskSessionResponse) SetHeaders(v map[string]*string) *CloseAICoachTaskSessionResponse {
	s.Headers = v
	return s
}

func (s *CloseAICoachTaskSessionResponse) SetStatusCode(v int32) *CloseAICoachTaskSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseAICoachTaskSessionResponse) SetBody(v *CloseAICoachTaskSessionResponseBody) *CloseAICoachTaskSessionResponse {
	s.Body = v
	return s
}

type CountTextRequest struct {
	// API
	//
	// example:
	//
	// PLATFORM
	GenerationSource *string `json:"generationSource,omitempty" xml:"generationSource,omitempty"`
	// example:
	//
	// Garment
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
	// example:
	//
	// 1
	PublishStatus *string `json:"publishStatus,omitempty" xml:"publishStatus,omitempty"`
	// example:
	//
	// RED_BOOK
	Style *string `json:"style,omitempty" xml:"style,omitempty"`
}

func (s CountTextRequest) String() string {
	return tea.Prettify(s)
}

func (s CountTextRequest) GoString() string {
	return s.String()
}

func (s *CountTextRequest) SetGenerationSource(v string) *CountTextRequest {
	s.GenerationSource = &v
	return s
}

func (s *CountTextRequest) SetIndustry(v string) *CountTextRequest {
	s.Industry = &v
	return s
}

func (s *CountTextRequest) SetPublishStatus(v string) *CountTextRequest {
	s.PublishStatus = &v
	return s
}

func (s *CountTextRequest) SetStyle(v string) *CountTextRequest {
	s.Style = &v
	return s
}

type CountTextResponseBody struct {
	// example:
	//
	// 6C9CB64D-E2D3-5BF2-A9E6-2445F952F178
	RequestId        *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CountTextCmdList []*CountTextResponseBodyCountTextCmdList `json:"countTextCmdList,omitempty" xml:"countTextCmdList,omitempty" type:"Repeated"`
}

func (s CountTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountTextResponseBody) GoString() string {
	return s.String()
}

func (s *CountTextResponseBody) SetRequestId(v string) *CountTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountTextResponseBody) SetCountTextCmdList(v []*CountTextResponseBodyCountTextCmdList) *CountTextResponseBody {
	s.CountTextCmdList = v
	return s
}

type CountTextResponseBodyCountTextCmdList struct {
	// example:
	//
	// 4
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// example:
	//
	// RED_BOOK
	Theme *string `json:"theme,omitempty" xml:"theme,omitempty"`
}

func (s CountTextResponseBodyCountTextCmdList) String() string {
	return tea.Prettify(s)
}

func (s CountTextResponseBodyCountTextCmdList) GoString() string {
	return s.String()
}

func (s *CountTextResponseBodyCountTextCmdList) SetCount(v int64) *CountTextResponseBodyCountTextCmdList {
	s.Count = &v
	return s
}

func (s *CountTextResponseBodyCountTextCmdList) SetTheme(v string) *CountTextResponseBodyCountTextCmdList {
	s.Theme = &v
	return s
}

type CountTextResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CountTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountTextResponse) String() string {
	return tea.Prettify(s)
}

func (s CountTextResponse) GoString() string {
	return s.String()
}

func (s *CountTextResponse) SetHeaders(v map[string]*string) *CountTextResponse {
	s.Headers = v
	return s
}

func (s *CountTextResponse) SetStatusCode(v int32) *CountTextResponse {
	s.StatusCode = &v
	return s
}

func (s *CountTextResponse) SetBody(v *CountTextResponseBody) *CountTextResponse {
	s.Body = v
	return s
}

type CreateAICoachTaskSessionRequest struct {
	// example:
	//
	// 821882330423951360
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 1730530943640489
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreateAICoachTaskSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAICoachTaskSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskSessionRequest) SetTaskId(v string) *CreateAICoachTaskSessionRequest {
	s.TaskId = &v
	return s
}

func (s *CreateAICoachTaskSessionRequest) SetUid(v string) *CreateAICoachTaskSessionRequest {
	s.Uid = &v
	return s
}

type CreateAICoachTaskSessionResponseBody struct {
	// rtctoken
	//
	// example:
	//
	// 11
	ChannelToken *string `json:"channelToken,omitempty" xml:"channelToken,omitempty"`
	// example:
	//
	// 4830493A-728F-5F19-BBCC-1443292E9C49
	RequestId  *string                                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ScriptInfo *CreateAICoachTaskSessionResponseBodyScriptInfo `json:"scriptInfo,omitempty" xml:"scriptInfo,omitempty" type:"Struct"`
	// example:
	//
	// 111
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Token
	//
	// example:
	//
	// 11
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// example:
	//
	// 11
	WebSocketUrl *string `json:"webSocketUrl,omitempty" xml:"webSocketUrl,omitempty"`
}

func (s CreateAICoachTaskSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAICoachTaskSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskSessionResponseBody) SetChannelToken(v string) *CreateAICoachTaskSessionResponseBody {
	s.ChannelToken = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBody) SetRequestId(v string) *CreateAICoachTaskSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBody) SetScriptInfo(v *CreateAICoachTaskSessionResponseBodyScriptInfo) *CreateAICoachTaskSessionResponseBody {
	s.ScriptInfo = v
	return s
}

func (s *CreateAICoachTaskSessionResponseBody) SetSessionId(v string) *CreateAICoachTaskSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBody) SetToken(v string) *CreateAICoachTaskSessionResponseBody {
	s.Token = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBody) SetWebSocketUrl(v string) *CreateAICoachTaskSessionResponseBody {
	s.WebSocketUrl = &v
	return s
}

type CreateAICoachTaskSessionResponseBodyScriptInfo struct {
	Initiator *string `json:"initiator,omitempty" xml:"initiator,omitempty"`
	// example:
	//
	// 11
	MaxDuration *int64 `json:"maxDuration,omitempty" xml:"maxDuration,omitempty"`
	// example:
	//
	// test
	ScriptDesc *string `json:"scriptDesc,omitempty" xml:"scriptDesc,omitempty"`
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
}

func (s CreateAICoachTaskSessionResponseBodyScriptInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAICoachTaskSessionResponseBodyScriptInfo) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetInitiator(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.Initiator = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetMaxDuration(v int64) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.MaxDuration = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetScriptDesc(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.ScriptDesc = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetScriptName(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.ScriptName = &v
	return s
}

type CreateAICoachTaskSessionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAICoachTaskSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAICoachTaskSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAICoachTaskSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskSessionResponse) SetHeaders(v map[string]*string) *CreateAICoachTaskSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateAICoachTaskSessionResponse) SetStatusCode(v int32) *CreateAICoachTaskSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAICoachTaskSessionResponse) SetBody(v *CreateAICoachTaskSessionResponseBody) *CreateAICoachTaskSessionResponse {
	s.Body = v
	return s
}

type CreateIllustrationTaskRequest struct {
	Body *IllustrationTaskCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIllustrationTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIllustrationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateIllustrationTaskRequest) SetBody(v *IllustrationTaskCreateCmd) *CreateIllustrationTaskRequest {
	s.Body = v
	return s
}

type CreateIllustrationTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IllustrationTaskResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIllustrationTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIllustrationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateIllustrationTaskResponse) SetHeaders(v map[string]*string) *CreateIllustrationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateIllustrationTaskResponse) SetStatusCode(v int32) *CreateIllustrationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIllustrationTaskResponse) SetBody(v *IllustrationTaskResult) *CreateIllustrationTaskResponse {
	s.Body = v
	return s
}

type CreateRealisticPortraitRequest struct {
	Ages []*int32 `json:"ages,omitempty" xml:"ages,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Cloth *int32 `json:"cloth,omitempty" xml:"cloth,omitempty"`
	// example:
	//
	// 1
	Color *int32 `json:"color,omitempty" xml:"color,omitempty"`
	// example:
	//
	// 11
	Custom *string  `json:"custom,omitempty" xml:"custom,omitempty"`
	Face   []*int32 `json:"face,omitempty" xml:"face,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Figure *int32 `json:"figure,omitempty" xml:"figure,omitempty"`
	// example:
	//
	// 1
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// 1
	HairColor *int32 `json:"hairColor,omitempty" xml:"hairColor,omitempty"`
	// example:
	//
	// 1
	Hairstyle *int32 `json:"hairstyle,omitempty" xml:"hairstyle,omitempty"`
	// example:
	//
	// 500
	Height   *int32  `json:"height,omitempty" xml:"height,omitempty"`
	ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	// example:
	//
	// 4
	Numbers *int32 `json:"numbers,omitempty" xml:"numbers,omitempty"`
	// example:
	//
	// 1:1
	Ratio *string `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// example:
	//
	// 500
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s CreateRealisticPortraitRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRealisticPortraitRequest) GoString() string {
	return s.String()
}

func (s *CreateRealisticPortraitRequest) SetAges(v []*int32) *CreateRealisticPortraitRequest {
	s.Ages = v
	return s
}

func (s *CreateRealisticPortraitRequest) SetCloth(v int32) *CreateRealisticPortraitRequest {
	s.Cloth = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetColor(v int32) *CreateRealisticPortraitRequest {
	s.Color = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetCustom(v string) *CreateRealisticPortraitRequest {
	s.Custom = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetFace(v []*int32) *CreateRealisticPortraitRequest {
	s.Face = v
	return s
}

func (s *CreateRealisticPortraitRequest) SetFigure(v int32) *CreateRealisticPortraitRequest {
	s.Figure = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetGender(v int32) *CreateRealisticPortraitRequest {
	s.Gender = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetHairColor(v int32) *CreateRealisticPortraitRequest {
	s.HairColor = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetHairstyle(v int32) *CreateRealisticPortraitRequest {
	s.Hairstyle = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetHeight(v int32) *CreateRealisticPortraitRequest {
	s.Height = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetImageUrl(v string) *CreateRealisticPortraitRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetNumbers(v int32) *CreateRealisticPortraitRequest {
	s.Numbers = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetRatio(v string) *CreateRealisticPortraitRequest {
	s.Ratio = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetWidth(v int32) *CreateRealisticPortraitRequest {
	s.Width = &v
	return s
}

type CreateRealisticPortraitResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// D5798660-1531-5D12-9C20-16FEE9D22351
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 313123123
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateRealisticPortraitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRealisticPortraitResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRealisticPortraitResponseBody) SetRequestId(v string) *CreateRealisticPortraitResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRealisticPortraitResponseBody) SetTaskId(v string) *CreateRealisticPortraitResponseBody {
	s.TaskId = &v
	return s
}

type CreateRealisticPortraitResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRealisticPortraitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRealisticPortraitResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRealisticPortraitResponse) GoString() string {
	return s.String()
}

func (s *CreateRealisticPortraitResponse) SetHeaders(v map[string]*string) *CreateRealisticPortraitResponse {
	s.Headers = v
	return s
}

func (s *CreateRealisticPortraitResponse) SetStatusCode(v int32) *CreateRealisticPortraitResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRealisticPortraitResponse) SetBody(v *CreateRealisticPortraitResponseBody) *CreateRealisticPortraitResponse {
	s.Body = v
	return s
}

type CreateTextTaskRequest struct {
	Body *TextTaskCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTextTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTextTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTextTaskRequest) SetBody(v *TextTaskCreateCmd) *CreateTextTaskRequest {
	s.Body = v
	return s
}

type CreateTextTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextTaskResult    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTextTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTextTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTextTaskResponse) SetHeaders(v map[string]*string) *CreateTextTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTextTaskResponse) SetStatusCode(v int32) *CreateTextTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTextTaskResponse) SetBody(v *TextTaskResult) *CreateTextTaskResponse {
	s.Body = v
	return s
}

type FinishAICoachTaskSessionRequest struct {
	// example:
	//
	// 111
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 222
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s FinishAICoachTaskSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishAICoachTaskSessionRequest) GoString() string {
	return s.String()
}

func (s *FinishAICoachTaskSessionRequest) SetSessionId(v string) *FinishAICoachTaskSessionRequest {
	s.SessionId = &v
	return s
}

func (s *FinishAICoachTaskSessionRequest) SetUid(v string) *FinishAICoachTaskSessionRequest {
	s.Uid = &v
	return s
}

type FinishAICoachTaskSessionResponseBody struct {
	// example:
	//
	// 14878724-A835-578D-9DD5-4779ADCE9221
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s FinishAICoachTaskSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FinishAICoachTaskSessionResponseBody) GoString() string {
	return s.String()
}

func (s *FinishAICoachTaskSessionResponseBody) SetRequestId(v string) *FinishAICoachTaskSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *FinishAICoachTaskSessionResponseBody) SetStatus(v string) *FinishAICoachTaskSessionResponseBody {
	s.Status = &v
	return s
}

type FinishAICoachTaskSessionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FinishAICoachTaskSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FinishAICoachTaskSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishAICoachTaskSessionResponse) GoString() string {
	return s.String()
}

func (s *FinishAICoachTaskSessionResponse) SetHeaders(v map[string]*string) *FinishAICoachTaskSessionResponse {
	s.Headers = v
	return s
}

func (s *FinishAICoachTaskSessionResponse) SetStatusCode(v int32) *FinishAICoachTaskSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishAICoachTaskSessionResponse) SetBody(v *FinishAICoachTaskSessionResponseBody) *FinishAICoachTaskSessionResponse {
	s.Body = v
	return s
}

type GetAICoachTaskSessionHistoryRequest struct {
	// example:
	//
	// 111
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 1251317954812712
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetAICoachTaskSessionHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachTaskSessionHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionHistoryRequest) SetSessionId(v string) *GetAICoachTaskSessionHistoryRequest {
	s.SessionId = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryRequest) SetUid(v string) *GetAICoachTaskSessionHistoryRequest {
	s.Uid = &v
	return s
}

type GetAICoachTaskSessionHistoryResponseBody struct {
	ConversationList []*GetAICoachTaskSessionHistoryResponseBodyConversationList `json:"conversationList,omitempty" xml:"conversationList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 2024-11-08 09:33:21
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// D5798660-1531-5D12-9C20-16FEE9D22351
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	// example:
	//
	// 2024-08-21 05:00:01
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1579404690269235
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetAICoachTaskSessionHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachTaskSessionHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetConversationList(v []*GetAICoachTaskSessionHistoryResponseBodyConversationList) *GetAICoachTaskSessionHistoryResponseBody {
	s.ConversationList = v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetDuration(v int64) *GetAICoachTaskSessionHistoryResponseBody {
	s.Duration = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetEndTime(v string) *GetAICoachTaskSessionHistoryResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetRequestId(v string) *GetAICoachTaskSessionHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetScriptName(v string) *GetAICoachTaskSessionHistoryResponseBody {
	s.ScriptName = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetStartTime(v string) *GetAICoachTaskSessionHistoryResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetStatus(v string) *GetAICoachTaskSessionHistoryResponseBody {
	s.Status = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetUid(v string) *GetAICoachTaskSessionHistoryResponseBody {
	s.Uid = &v
	return s
}

type GetAICoachTaskSessionHistoryResponseBodyConversationList struct {
	// example:
	//
	// {}
	EvaluationResult *string `json:"evaluationResult,omitempty" xml:"evaluationResult,omitempty"`
	Message          *string `json:"message,omitempty" xml:"message,omitempty"`
	Role             *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GetAICoachTaskSessionHistoryResponseBodyConversationList) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachTaskSessionHistoryResponseBodyConversationList) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) SetEvaluationResult(v string) *GetAICoachTaskSessionHistoryResponseBodyConversationList {
	s.EvaluationResult = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) SetMessage(v string) *GetAICoachTaskSessionHistoryResponseBodyConversationList {
	s.Message = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) SetRole(v string) *GetAICoachTaskSessionHistoryResponseBodyConversationList {
	s.Role = &v
	return s
}

type GetAICoachTaskSessionHistoryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAICoachTaskSessionHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAICoachTaskSessionHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachTaskSessionHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionHistoryResponse) SetHeaders(v map[string]*string) *GetAICoachTaskSessionHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponse) SetStatusCode(v int32) *GetAICoachTaskSessionHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponse) SetBody(v *GetAICoachTaskSessionHistoryResponseBody) *GetAICoachTaskSessionHistoryResponse {
	s.Body = v
	return s
}

type GetAICoachTaskSessionReportRequest struct {
	// example:
	//
	// 1111
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 1707732338016307
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetAICoachTaskSessionReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachTaskSessionReportRequest) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionReportRequest) SetSessionId(v string) *GetAICoachTaskSessionReportRequest {
	s.SessionId = &v
	return s
}

func (s *GetAICoachTaskSessionReportRequest) SetUid(v string) *GetAICoachTaskSessionReportRequest {
	s.Uid = &v
	return s
}

type GetAICoachTaskSessionReportResponseBody struct {
	// example:
	//
	// 0
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 2024-11-08 09:33:21
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// {}
	EvaluationResult *string `json:"evaluationResult,omitempty" xml:"evaluationResult,omitempty"`
	// example:
	//
	// 0E8B1746-AE35-5C4B-A3A8-345B274AE32C
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	// example:
	//
	// 2024-10-11 09:58:01
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1276673855116835
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetAICoachTaskSessionReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachTaskSessionReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionReportResponseBody) SetDuration(v int64) *GetAICoachTaskSessionReportResponseBody {
	s.Duration = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetEndTime(v string) *GetAICoachTaskSessionReportResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetEvaluationResult(v string) *GetAICoachTaskSessionReportResponseBody {
	s.EvaluationResult = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetRequestId(v string) *GetAICoachTaskSessionReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetScriptName(v string) *GetAICoachTaskSessionReportResponseBody {
	s.ScriptName = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetStartTime(v string) *GetAICoachTaskSessionReportResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetStatus(v string) *GetAICoachTaskSessionReportResponseBody {
	s.Status = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetUid(v string) *GetAICoachTaskSessionReportResponseBody {
	s.Uid = &v
	return s
}

type GetAICoachTaskSessionReportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAICoachTaskSessionReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAICoachTaskSessionReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachTaskSessionReportResponse) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionReportResponse) SetHeaders(v map[string]*string) *GetAICoachTaskSessionReportResponse {
	s.Headers = v
	return s
}

func (s *GetAICoachTaskSessionReportResponse) SetStatusCode(v int32) *GetAICoachTaskSessionReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponse) SetBody(v *GetAICoachTaskSessionReportResponseBody) *GetAICoachTaskSessionReportResponse {
	s.Body = v
	return s
}

type GetIllustrationResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IllustrationResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIllustrationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIllustrationResponse) GoString() string {
	return s.String()
}

func (s *GetIllustrationResponse) SetHeaders(v map[string]*string) *GetIllustrationResponse {
	s.Headers = v
	return s
}

func (s *GetIllustrationResponse) SetStatusCode(v int32) *GetIllustrationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIllustrationResponse) SetBody(v *IllustrationResult) *GetIllustrationResponse {
	s.Body = v
	return s
}

type GetIllustrationTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IllustrationTaskResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIllustrationTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIllustrationTaskResponse) GoString() string {
	return s.String()
}

func (s *GetIllustrationTaskResponse) SetHeaders(v map[string]*string) *GetIllustrationTaskResponse {
	s.Headers = v
	return s
}

func (s *GetIllustrationTaskResponse) SetStatusCode(v int32) *GetIllustrationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIllustrationTaskResponse) SetBody(v *IllustrationTaskResult) *GetIllustrationTaskResponse {
	s.Body = v
	return s
}

type GetOssUploadTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 8021678.png
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ProductImage
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
}

func (s GetOssUploadTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadTokenRequest) GoString() string {
	return s.String()
}

func (s *GetOssUploadTokenRequest) SetFileName(v string) *GetOssUploadTokenRequest {
	s.FileName = &v
	return s
}

func (s *GetOssUploadTokenRequest) SetFileType(v string) *GetOssUploadTokenRequest {
	s.FileType = &v
	return s
}

type GetOssUploadTokenResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssUploadTokenResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssUploadTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadTokenResponse) GoString() string {
	return s.String()
}

func (s *GetOssUploadTokenResponse) SetHeaders(v map[string]*string) *GetOssUploadTokenResponse {
	s.Headers = v
	return s
}

func (s *GetOssUploadTokenResponse) SetStatusCode(v int32) *GetOssUploadTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssUploadTokenResponse) SetBody(v *GetOssUploadTokenResult) *GetOssUploadTokenResponse {
	s.Body = v
	return s
}

type GetProjectTaskRequest struct {
	// example:
	//
	// 20230823218109326025-1200
	IdempotentId *string `json:"IdempotentId,omitempty" xml:"IdempotentId,omitempty"`
	// example:
	//
	// 313123123
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetProjectTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectTaskRequest) GoString() string {
	return s.String()
}

func (s *GetProjectTaskRequest) SetIdempotentId(v string) *GetProjectTaskRequest {
	s.IdempotentId = &v
	return s
}

func (s *GetProjectTaskRequest) SetTaskId(v string) *GetProjectTaskRequest {
	s.TaskId = &v
	return s
}

type GetProjectTaskResponseBody struct {
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 5389BE87-571B-573C-90ED-F07C5E68760B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// www.ali.com
	VideoDownloadUrl *string `json:"videoDownloadUrl,omitempty" xml:"videoDownloadUrl,omitempty"`
	// example:
	//
	// 111
	VideoDuration *int32 `json:"videoDuration,omitempty" xml:"videoDuration,omitempty"`
	// example:
	//
	// www.ali.com
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s GetProjectTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectTaskResponseBody) SetErrorMsg(v string) *GetProjectTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetProjectTaskResponseBody) SetRequestId(v string) *GetProjectTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectTaskResponseBody) SetStatus(v string) *GetProjectTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetProjectTaskResponseBody) SetVideoDownloadUrl(v string) *GetProjectTaskResponseBody {
	s.VideoDownloadUrl = &v
	return s
}

func (s *GetProjectTaskResponseBody) SetVideoDuration(v int32) *GetProjectTaskResponseBody {
	s.VideoDuration = &v
	return s
}

func (s *GetProjectTaskResponseBody) SetVideoUrl(v string) *GetProjectTaskResponseBody {
	s.VideoUrl = &v
	return s
}

type GetProjectTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectTaskResponse) GoString() string {
	return s.String()
}

func (s *GetProjectTaskResponse) SetHeaders(v map[string]*string) *GetProjectTaskResponse {
	s.Headers = v
	return s
}

func (s *GetProjectTaskResponse) SetStatusCode(v int32) *GetProjectTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectTaskResponse) SetBody(v *GetProjectTaskResponseBody) *GetProjectTaskResponse {
	s.Body = v
	return s
}

type GetTextResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextResult        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTextResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTextResponse) GoString() string {
	return s.String()
}

func (s *GetTextResponse) SetHeaders(v map[string]*string) *GetTextResponse {
	s.Headers = v
	return s
}

func (s *GetTextResponse) SetStatusCode(v int32) *GetTextResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTextResponse) SetBody(v *TextResult) *GetTextResponse {
	s.Body = v
	return s
}

type GetTextTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextTaskResult    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTextTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTextTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTextTaskResponse) SetHeaders(v map[string]*string) *GetTextTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTextTaskResponse) SetStatusCode(v int32) *GetTextTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTextTaskResponse) SetBody(v *TextTaskResult) *GetTextTaskResponse {
	s.Body = v
	return s
}

type GetTextTemplateRequest struct {
	// example:
	//
	// Car
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
}

func (s GetTextTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTextTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTextTemplateRequest) SetIndustry(v string) *GetTextTemplateRequest {
	s.Industry = &v
	return s
}

type GetTextTemplateResponseBody struct {
	AvailableIndustry *GetTextTemplateResponseBodyAvailableIndustry `json:"availableIndustry,omitempty" xml:"availableIndustry,omitempty" type:"Struct"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetTextTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTextTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTextTemplateResponseBody) SetAvailableIndustry(v *GetTextTemplateResponseBodyAvailableIndustry) *GetTextTemplateResponseBody {
	s.AvailableIndustry = v
	return s
}

func (s *GetTextTemplateResponseBody) SetRequestId(v string) *GetTextTemplateResponseBody {
	s.RequestId = &v
	return s
}

type GetTextTemplateResponseBodyAvailableIndustry struct {
	// example:
	//
	// Car
	Name          *string                                                      `json:"name,omitempty" xml:"name,omitempty"`
	TextModeTypes []*GetTextTemplateResponseBodyAvailableIndustryTextModeTypes `json:"textModeTypes,omitempty" xml:"textModeTypes,omitempty" type:"Repeated"`
}

func (s GetTextTemplateResponseBodyAvailableIndustry) String() string {
	return tea.Prettify(s)
}

func (s GetTextTemplateResponseBodyAvailableIndustry) GoString() string {
	return s.String()
}

func (s *GetTextTemplateResponseBodyAvailableIndustry) SetName(v string) *GetTextTemplateResponseBodyAvailableIndustry {
	s.Name = &v
	return s
}

func (s *GetTextTemplateResponseBodyAvailableIndustry) SetTextModeTypes(v []*GetTextTemplateResponseBodyAvailableIndustryTextModeTypes) *GetTextTemplateResponseBodyAvailableIndustry {
	s.TextModeTypes = v
	return s
}

type GetTextTemplateResponseBodyAvailableIndustryTextModeTypes struct {
	// example:
	//
	// Rewrite
	Name       *string                                                                `json:"name,omitempty" xml:"name,omitempty"`
	TextStyles []*GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles `json:"textStyles,omitempty" xml:"textStyles,omitempty" type:"Repeated"`
}

func (s GetTextTemplateResponseBodyAvailableIndustryTextModeTypes) String() string {
	return tea.Prettify(s)
}

func (s GetTextTemplateResponseBodyAvailableIndustryTextModeTypes) GoString() string {
	return s.String()
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypes) SetName(v string) *GetTextTemplateResponseBodyAvailableIndustryTextModeTypes {
	s.Name = &v
	return s
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypes) SetTextStyles(v []*GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) *GetTextTemplateResponseBodyAvailableIndustryTextModeTypes {
	s.TextStyles = v
	return s
}

type GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles struct {
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// false
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// example:
	//
	// RED_BOOK
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 111
	TemplateKey *string `json:"templateKey,omitempty" xml:"templateKey,omitempty"`
}

func (s GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) String() string {
	return tea.Prettify(s)
}

func (s GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) GoString() string {
	return s.String()
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) SetDesc(v string) *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles {
	s.Desc = &v
	return s
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) SetDisabled(v bool) *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles {
	s.Disabled = &v
	return s
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) SetName(v string) *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles {
	s.Name = &v
	return s
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) SetTemplateKey(v string) *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles {
	s.TemplateKey = &v
	return s
}

type GetTextTemplateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTextTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTextTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTextTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTextTemplateResponse) SetHeaders(v map[string]*string) *GetTextTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTextTemplateResponse) SetStatusCode(v int32) *GetTextTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTextTemplateResponse) SetBody(v *GetTextTemplateResponseBody) *GetTextTemplateResponse {
	s.Body = v
	return s
}

type InteractTextRequest struct {
	// example:
	//
	// 1000222
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 144285195534941
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s InteractTextRequest) String() string {
	return tea.Prettify(s)
}

func (s InteractTextRequest) GoString() string {
	return s.String()
}

func (s *InteractTextRequest) SetAgentId(v string) *InteractTextRequest {
	s.AgentId = &v
	return s
}

func (s *InteractTextRequest) SetContent(v string) *InteractTextRequest {
	s.Content = &v
	return s
}

func (s *InteractTextRequest) SetSessionId(v string) *InteractTextRequest {
	s.SessionId = &v
	return s
}

type InteractTextResponseBody struct {
	// example:
	//
	// false
	End *bool `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 1
	Index         *int32    `json:"index,omitempty" xml:"index,omitempty"`
	Message       *string   `json:"message,omitempty" xml:"message,omitempty"`
	RelatedImages []*string `json:"relatedImages,omitempty" xml:"relatedImages,omitempty" type:"Repeated"`
	RelatedVideos []*string `json:"relatedVideos,omitempty" xml:"relatedVideos,omitempty" type:"Repeated"`
	// example:
	//
	// 79e954faffe2415ebd18188ba787d78e
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 2
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InteractTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InteractTextResponseBody) GoString() string {
	return s.String()
}

func (s *InteractTextResponseBody) SetEnd(v bool) *InteractTextResponseBody {
	s.End = &v
	return s
}

func (s *InteractTextResponseBody) SetIndex(v int32) *InteractTextResponseBody {
	s.Index = &v
	return s
}

func (s *InteractTextResponseBody) SetMessage(v string) *InteractTextResponseBody {
	s.Message = &v
	return s
}

func (s *InteractTextResponseBody) SetRelatedImages(v []*string) *InteractTextResponseBody {
	s.RelatedImages = v
	return s
}

func (s *InteractTextResponseBody) SetRelatedVideos(v []*string) *InteractTextResponseBody {
	s.RelatedVideos = v
	return s
}

func (s *InteractTextResponseBody) SetSessionId(v string) *InteractTextResponseBody {
	s.SessionId = &v
	return s
}

func (s *InteractTextResponseBody) SetType(v int32) *InteractTextResponseBody {
	s.Type = &v
	return s
}

type InteractTextResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InteractTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InteractTextResponse) String() string {
	return tea.Prettify(s)
}

func (s InteractTextResponse) GoString() string {
	return s.String()
}

func (s *InteractTextResponse) SetHeaders(v map[string]*string) *InteractTextResponse {
	s.Headers = v
	return s
}

func (s *InteractTextResponse) SetStatusCode(v int32) *InteractTextResponse {
	s.StatusCode = &v
	return s
}

func (s *InteractTextResponse) SetBody(v *InteractTextResponseBody) *InteractTextResponse {
	s.Body = v
	return s
}

type ListAICoachTaskPageRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// FINISHED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 111
	StudentId *string `json:"studentId,omitempty" xml:"studentId,omitempty"`
	// example:
	//
	// 313123123
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListAICoachTaskPageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAICoachTaskPageRequest) GoString() string {
	return s.String()
}

func (s *ListAICoachTaskPageRequest) SetPageNumber(v int32) *ListAICoachTaskPageRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAICoachTaskPageRequest) SetPageSize(v int32) *ListAICoachTaskPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListAICoachTaskPageRequest) SetStatus(v string) *ListAICoachTaskPageRequest {
	s.Status = &v
	return s
}

func (s *ListAICoachTaskPageRequest) SetStudentId(v string) *ListAICoachTaskPageRequest {
	s.StudentId = &v
	return s
}

func (s *ListAICoachTaskPageRequest) SetTaskId(v string) *ListAICoachTaskPageRequest {
	s.TaskId = &v
	return s
}

type ListAICoachTaskPageResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// D7F2B74F-63F2-5DD6-95E4-F408EAD6617E
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TaskList  []*ListAICoachTaskPageResponseBodyTaskList `json:"taskList,omitempty" xml:"taskList,omitempty" type:"Repeated"`
}

func (s ListAICoachTaskPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAICoachTaskPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListAICoachTaskPageResponseBody) SetRequestId(v string) *ListAICoachTaskPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAICoachTaskPageResponseBody) SetTaskList(v []*ListAICoachTaskPageResponseBodyTaskList) *ListAICoachTaskPageResponseBody {
	s.TaskList = v
	return s
}

type ListAICoachTaskPageResponseBodyTaskList struct {
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// example:
	//
	// FINISHED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 222
	StudentId *string `json:"studentId,omitempty" xml:"studentId,omitempty"`
	// example:
	//
	// 11111111111
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListAICoachTaskPageResponseBodyTaskList) String() string {
	return tea.Prettify(s)
}

func (s ListAICoachTaskPageResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *ListAICoachTaskPageResponseBodyTaskList) SetFinishTime(v string) *ListAICoachTaskPageResponseBodyTaskList {
	s.FinishTime = &v
	return s
}

func (s *ListAICoachTaskPageResponseBodyTaskList) SetStatus(v string) *ListAICoachTaskPageResponseBodyTaskList {
	s.Status = &v
	return s
}

func (s *ListAICoachTaskPageResponseBodyTaskList) SetStudentId(v string) *ListAICoachTaskPageResponseBodyTaskList {
	s.StudentId = &v
	return s
}

func (s *ListAICoachTaskPageResponseBodyTaskList) SetTaskId(v string) *ListAICoachTaskPageResponseBodyTaskList {
	s.TaskId = &v
	return s
}

type ListAICoachTaskPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAICoachTaskPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAICoachTaskPageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAICoachTaskPageResponse) GoString() string {
	return s.String()
}

func (s *ListAICoachTaskPageResponse) SetHeaders(v map[string]*string) *ListAICoachTaskPageResponse {
	s.Headers = v
	return s
}

func (s *ListAICoachTaskPageResponse) SetStatusCode(v int32) *ListAICoachTaskPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAICoachTaskPageResponse) SetBody(v *ListAICoachTaskPageResponseBody) *ListAICoachTaskPageResponse {
	s.Body = v
	return s
}

type ListAnchorRequest struct {
	// example:
	//
	// PUBLIC_MODEL
	AnchorType *string `json:"anchorType,omitempty" xml:"anchorType,omitempty"`
	// example:
	//
	// 9:16
	CoverRate *string `json:"coverRate,omitempty" xml:"coverRate,omitempty"`
	// example:
	//
	// staticTransparency
	DigitalHumanType *string `json:"digitalHumanType,omitempty" xml:"digitalHumanType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize    *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResSpecType *string `json:"resSpecType,omitempty" xml:"resSpecType,omitempty"`
	// example:
	//
	// offlineSynthesis
	UseScene *string `json:"useScene,omitempty" xml:"useScene,omitempty"`
}

func (s ListAnchorRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAnchorRequest) GoString() string {
	return s.String()
}

func (s *ListAnchorRequest) SetAnchorType(v string) *ListAnchorRequest {
	s.AnchorType = &v
	return s
}

func (s *ListAnchorRequest) SetCoverRate(v string) *ListAnchorRequest {
	s.CoverRate = &v
	return s
}

func (s *ListAnchorRequest) SetDigitalHumanType(v string) *ListAnchorRequest {
	s.DigitalHumanType = &v
	return s
}

func (s *ListAnchorRequest) SetPageNumber(v int32) *ListAnchorRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAnchorRequest) SetPageSize(v int32) *ListAnchorRequest {
	s.PageSize = &v
	return s
}

func (s *ListAnchorRequest) SetResSpecType(v string) *ListAnchorRequest {
	s.ResSpecType = &v
	return s
}

func (s *ListAnchorRequest) SetUseScene(v string) *ListAnchorRequest {
	s.UseScene = &v
	return s
}

type ListAnchorResponseBody struct {
	// code
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// Deduct.DeductTaskAlreadySuccess
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Failed to proxy flink ui request, message: An error occurred: Invalid UUID string: jobsn
	ErrorMessage *string           `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	List         []*AnchorResponse `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// 5389BE87-571B-573C-90ED-F07C5E68760B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAnchorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAnchorResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnchorResponseBody) SetCode(v string) *ListAnchorResponseBody {
	s.Code = &v
	return s
}

func (s *ListAnchorResponseBody) SetErrorCode(v string) *ListAnchorResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAnchorResponseBody) SetErrorMessage(v string) *ListAnchorResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAnchorResponseBody) SetList(v []*AnchorResponse) *ListAnchorResponseBody {
	s.List = v
	return s
}

func (s *ListAnchorResponseBody) SetRequestId(v string) *ListAnchorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnchorResponseBody) SetSuccess(v bool) *ListAnchorResponseBody {
	s.Success = &v
	return s
}

func (s *ListAnchorResponseBody) SetTotal(v int32) *ListAnchorResponseBody {
	s.Total = &v
	return s
}

type ListAnchorResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAnchorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAnchorResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAnchorResponse) GoString() string {
	return s.String()
}

func (s *ListAnchorResponse) SetHeaders(v map[string]*string) *ListAnchorResponse {
	s.Headers = v
	return s
}

func (s *ListAnchorResponse) SetStatusCode(v int32) *ListAnchorResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAnchorResponse) SetBody(v *ListAnchorResponseBody) *ListAnchorResponse {
	s.Body = v
	return s
}

type ListAvatarProjectRequest struct {
	ProjectIdList []*string `json:"projectIdList,omitempty" xml:"projectIdList,omitempty" type:"Repeated"`
}

func (s ListAvatarProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvatarProjectRequest) GoString() string {
	return s.String()
}

func (s *ListAvatarProjectRequest) SetProjectIdList(v []*string) *ListAvatarProjectRequest {
	s.ProjectIdList = v
	return s
}

type ListAvatarProjectShrinkRequest struct {
	ProjectIdListShrink *string `json:"projectIdList,omitempty" xml:"projectIdList,omitempty"`
}

func (s ListAvatarProjectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvatarProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAvatarProjectShrinkRequest) SetProjectIdListShrink(v string) *ListAvatarProjectShrinkRequest {
	s.ProjectIdListShrink = &v
	return s
}

type ListAvatarProjectResponseBody struct {
	QueryAvatarProjectResultList []*ListAvatarProjectResponseBodyQueryAvatarProjectResultList `json:"queryAvatarProjectResultList,omitempty" xml:"queryAvatarProjectResultList,omitempty" type:"Repeated"`
	// example:
	//
	// D7F2B74F-63F2-5DD6-95E4-F408EAD6617E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAvatarProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvatarProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvatarProjectResponseBody) SetQueryAvatarProjectResultList(v []*ListAvatarProjectResponseBodyQueryAvatarProjectResultList) *ListAvatarProjectResponseBody {
	s.QueryAvatarProjectResultList = v
	return s
}

func (s *ListAvatarProjectResponseBody) SetRequestId(v string) *ListAvatarProjectResponseBody {
	s.RequestId = &v
	return s
}

type ListAvatarProjectResponseBodyQueryAvatarProjectResultList struct {
	// example:
	//
	// 1000206
	AgentId  *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 12826084562688
	ProjectId   *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// example:
	//
	// DEPLOYING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListAvatarProjectResponseBodyQueryAvatarProjectResultList) String() string {
	return tea.Prettify(s)
}

func (s ListAvatarProjectResponseBodyQueryAvatarProjectResultList) GoString() string {
	return s.String()
}

func (s *ListAvatarProjectResponseBodyQueryAvatarProjectResultList) SetAgentId(v string) *ListAvatarProjectResponseBodyQueryAvatarProjectResultList {
	s.AgentId = &v
	return s
}

func (s *ListAvatarProjectResponseBodyQueryAvatarProjectResultList) SetErrorMsg(v string) *ListAvatarProjectResponseBodyQueryAvatarProjectResultList {
	s.ErrorMsg = &v
	return s
}

func (s *ListAvatarProjectResponseBodyQueryAvatarProjectResultList) SetProjectId(v string) *ListAvatarProjectResponseBodyQueryAvatarProjectResultList {
	s.ProjectId = &v
	return s
}

func (s *ListAvatarProjectResponseBodyQueryAvatarProjectResultList) SetProjectName(v string) *ListAvatarProjectResponseBodyQueryAvatarProjectResultList {
	s.ProjectName = &v
	return s
}

func (s *ListAvatarProjectResponseBodyQueryAvatarProjectResultList) SetStatus(v string) *ListAvatarProjectResponseBodyQueryAvatarProjectResultList {
	s.Status = &v
	return s
}

type ListAvatarProjectResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvatarProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvatarProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvatarProjectResponse) GoString() string {
	return s.String()
}

func (s *ListAvatarProjectResponse) SetHeaders(v map[string]*string) *ListAvatarProjectResponse {
	s.Headers = v
	return s
}

func (s *ListAvatarProjectResponse) SetStatusCode(v int32) *ListAvatarProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvatarProjectResponse) SetBody(v *ListAvatarProjectResponseBody) *ListAvatarProjectResponse {
	s.Body = v
	return s
}

type ListTextThemesRequest struct {
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
}

func (s ListTextThemesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTextThemesRequest) GoString() string {
	return s.String()
}

func (s *ListTextThemesRequest) SetIndustry(v string) *ListTextThemesRequest {
	s.Industry = &v
	return s
}

type ListTextThemesResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextThemeListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTextThemesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTextThemesResponse) GoString() string {
	return s.String()
}

func (s *ListTextThemesResponse) SetHeaders(v map[string]*string) *ListTextThemesResponse {
	s.Headers = v
	return s
}

func (s *ListTextThemesResponse) SetStatusCode(v int32) *ListTextThemesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTextThemesResponse) SetBody(v *TextThemeListResult) *ListTextThemesResponse {
	s.Body = v
	return s
}

type ListTextsRequest struct {
	// example:
	//
	// API
	GenerationSource *string `json:"generationSource,omitempty" xml:"generationSource,omitempty"`
	// example:
	//
	// Common
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
	Keyword  *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// PUBLISH
	PublishStatus *string `json:"publishStatus,omitempty" xml:"publishStatus,omitempty"`
	// example:
	//
	// WECHAT_MOMENT
	TextStyleType *string `json:"textStyleType,omitempty" xml:"textStyleType,omitempty"`
	// example:
	//
	// xxx
	TextTheme *string `json:"textTheme,omitempty" xml:"textTheme,omitempty"`
}

func (s ListTextsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTextsRequest) GoString() string {
	return s.String()
}

func (s *ListTextsRequest) SetGenerationSource(v string) *ListTextsRequest {
	s.GenerationSource = &v
	return s
}

func (s *ListTextsRequest) SetIndustry(v string) *ListTextsRequest {
	s.Industry = &v
	return s
}

func (s *ListTextsRequest) SetKeyword(v string) *ListTextsRequest {
	s.Keyword = &v
	return s
}

func (s *ListTextsRequest) SetPageNumber(v int32) *ListTextsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTextsRequest) SetPageSize(v int32) *ListTextsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTextsRequest) SetPublishStatus(v string) *ListTextsRequest {
	s.PublishStatus = &v
	return s
}

func (s *ListTextsRequest) SetTextStyleType(v string) *ListTextsRequest {
	s.TextStyleType = &v
	return s
}

func (s *ListTextsRequest) SetTextTheme(v string) *ListTextsRequest {
	s.TextTheme = &v
	return s
}

type ListTextsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextQueryResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTextsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTextsResponse) GoString() string {
	return s.String()
}

func (s *ListTextsResponse) SetHeaders(v map[string]*string) *ListTextsResponse {
	s.Headers = v
	return s
}

func (s *ListTextsResponse) SetStatusCode(v int32) *ListTextsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTextsResponse) SetBody(v *TextQueryResult) *ListTextsResponse {
	s.Body = v
	return s
}

type ListVoiceModelsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize    *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResSpecType *string `json:"resSpecType,omitempty" xml:"resSpecType,omitempty"`
	// example:
	//
	// offlineSynthesis
	UseScene *string `json:"useScene,omitempty" xml:"useScene,omitempty"`
	// example:
	//
	// PRIVATE_VOICE
	VoiceType *string `json:"voiceType,omitempty" xml:"voiceType,omitempty"`
}

func (s ListVoiceModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceModelsRequest) GoString() string {
	return s.String()
}

func (s *ListVoiceModelsRequest) SetPageNumber(v int32) *ListVoiceModelsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVoiceModelsRequest) SetPageSize(v int32) *ListVoiceModelsRequest {
	s.PageSize = &v
	return s
}

func (s *ListVoiceModelsRequest) SetResSpecType(v string) *ListVoiceModelsRequest {
	s.ResSpecType = &v
	return s
}

func (s *ListVoiceModelsRequest) SetUseScene(v string) *ListVoiceModelsRequest {
	s.UseScene = &v
	return s
}

func (s *ListVoiceModelsRequest) SetVoiceType(v string) *ListVoiceModelsRequest {
	s.VoiceType = &v
	return s
}

type ListVoiceModelsResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 040002
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Failed to proxy flink ui request, message: An error occurred: Invalid UUID string: jobsn
	ErrorMessage *string               `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	List         []*VoiceModelResponse `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 14878724-A835-578D-9DD5-4779ADCE9221
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListVoiceModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVoiceModelsResponseBody) SetCode(v string) *ListVoiceModelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListVoiceModelsResponseBody) SetErrorCode(v string) *ListVoiceModelsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListVoiceModelsResponseBody) SetErrorMessage(v string) *ListVoiceModelsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListVoiceModelsResponseBody) SetList(v []*VoiceModelResponse) *ListVoiceModelsResponseBody {
	s.List = v
	return s
}

func (s *ListVoiceModelsResponseBody) SetRequestId(v string) *ListVoiceModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVoiceModelsResponseBody) SetSuccess(v bool) *ListVoiceModelsResponseBody {
	s.Success = &v
	return s
}

func (s *ListVoiceModelsResponseBody) SetTotal(v int32) *ListVoiceModelsResponseBody {
	s.Total = &v
	return s
}

type ListVoiceModelsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVoiceModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVoiceModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVoiceModelsResponse) GoString() string {
	return s.String()
}

func (s *ListVoiceModelsResponse) SetHeaders(v map[string]*string) *ListVoiceModelsResponse {
	s.Headers = v
	return s
}

func (s *ListVoiceModelsResponse) SetStatusCode(v int32) *ListVoiceModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVoiceModelsResponse) SetBody(v *ListVoiceModelsResponseBody) *ListVoiceModelsResponse {
	s.Body = v
	return s
}

type OperateAvatarProjectRequest struct {
	// example:
	//
	// DELETE
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// example:
	//
	// 812907463682949120
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 1
	ResChannelNumber *int32 `json:"resChannelNumber,omitempty" xml:"resChannelNumber,omitempty"`
	// example:
	//
	// FREE
	ResType *string `json:"resType,omitempty" xml:"resType,omitempty"`
}

func (s OperateAvatarProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateAvatarProjectRequest) GoString() string {
	return s.String()
}

func (s *OperateAvatarProjectRequest) SetOperateType(v string) *OperateAvatarProjectRequest {
	s.OperateType = &v
	return s
}

func (s *OperateAvatarProjectRequest) SetProjectId(v string) *OperateAvatarProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *OperateAvatarProjectRequest) SetResChannelNumber(v int32) *OperateAvatarProjectRequest {
	s.ResChannelNumber = &v
	return s
}

func (s *OperateAvatarProjectRequest) SetResType(v string) *OperateAvatarProjectRequest {
	s.ResType = &v
	return s
}

type OperateAvatarProjectResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 86A90C40-D1AB-50DA-A4B1-0D545F80F2FE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OperateAvatarProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateAvatarProjectResponseBody) GoString() string {
	return s.String()
}

func (s *OperateAvatarProjectResponseBody) SetRequestId(v string) *OperateAvatarProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateAvatarProjectResponseBody) SetSuccess(v bool) *OperateAvatarProjectResponseBody {
	s.Success = &v
	return s
}

type OperateAvatarProjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateAvatarProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateAvatarProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateAvatarProjectResponse) GoString() string {
	return s.String()
}

func (s *OperateAvatarProjectResponse) SetHeaders(v map[string]*string) *OperateAvatarProjectResponse {
	s.Headers = v
	return s
}

func (s *OperateAvatarProjectResponse) SetStatusCode(v int32) *OperateAvatarProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateAvatarProjectResponse) SetBody(v *OperateAvatarProjectResponseBody) *OperateAvatarProjectResponse {
	s.Body = v
	return s
}

type QueryAvatarProjectRequest struct {
	// example:
	//
	// 11111
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s QueryAvatarProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarProjectRequest) GoString() string {
	return s.String()
}

func (s *QueryAvatarProjectRequest) SetProjectId(v string) *QueryAvatarProjectRequest {
	s.ProjectId = &v
	return s
}

type QueryAvatarProjectResponseBody struct {
	// example:
	//
	// 1000222
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string                                 `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Frames   []*QueryAvatarProjectResponseBodyFrames `json:"frames,omitempty" xml:"frames,omitempty" type:"Repeated"`
	// example:
	//
	// doc_test_3
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// example:
	//
	// 2C331582-7390-5949-8D9A-AC8239185B37
	RequestId   *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResSpecType *string `json:"resSpecType,omitempty" xml:"resSpecType,omitempty"`
	ScaleType   *string `json:"scaleType,omitempty" xml:"scaleType,omitempty"`
	// example:
	//
	// DEPLOYING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryAvatarProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarProjectResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAvatarProjectResponseBody) SetAgentId(v string) *QueryAvatarProjectResponseBody {
	s.AgentId = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetErrorMsg(v string) *QueryAvatarProjectResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetFrames(v []*QueryAvatarProjectResponseBodyFrames) *QueryAvatarProjectResponseBody {
	s.Frames = v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetProjectName(v string) *QueryAvatarProjectResponseBody {
	s.ProjectName = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetRequestId(v string) *QueryAvatarProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetResSpecType(v string) *QueryAvatarProjectResponseBody {
	s.ResSpecType = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetScaleType(v string) *QueryAvatarProjectResponseBody {
	s.ScaleType = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetStatus(v string) *QueryAvatarProjectResponseBody {
	s.Status = &v
	return s
}

type QueryAvatarProjectResponseBodyFrames struct {
	Layers      []*QueryAvatarProjectResponseBodyFramesLayers    `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	VideoScript *QueryAvatarProjectResponseBodyFramesVideoScript `json:"videoScript,omitempty" xml:"videoScript,omitempty" type:"Struct"`
}

func (s QueryAvatarProjectResponseBodyFrames) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarProjectResponseBodyFrames) GoString() string {
	return s.String()
}

func (s *QueryAvatarProjectResponseBodyFrames) SetLayers(v []*QueryAvatarProjectResponseBodyFramesLayers) *QueryAvatarProjectResponseBodyFrames {
	s.Layers = v
	return s
}

func (s *QueryAvatarProjectResponseBodyFrames) SetVideoScript(v *QueryAvatarProjectResponseBodyFramesVideoScript) *QueryAvatarProjectResponseBodyFrames {
	s.VideoScript = v
	return s
}

type QueryAvatarProjectResponseBodyFramesLayers struct {
	Height    *int32                                              `json:"height,omitempty" xml:"height,omitempty"`
	Material  *QueryAvatarProjectResponseBodyFramesLayersMaterial `json:"material,omitempty" xml:"material,omitempty" type:"Struct"`
	PositionX *int32                                              `json:"positionX,omitempty" xml:"positionX,omitempty"`
	PositionY *int32                                              `json:"positionY,omitempty" xml:"positionY,omitempty"`
	Type      *string                                             `json:"type,omitempty" xml:"type,omitempty"`
	Width     *int32                                              `json:"width,omitempty" xml:"width,omitempty"`
}

func (s QueryAvatarProjectResponseBodyFramesLayers) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarProjectResponseBodyFramesLayers) GoString() string {
	return s.String()
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) SetHeight(v int32) *QueryAvatarProjectResponseBodyFramesLayers {
	s.Height = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) SetMaterial(v *QueryAvatarProjectResponseBodyFramesLayersMaterial) *QueryAvatarProjectResponseBodyFramesLayers {
	s.Material = v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) SetPositionX(v int32) *QueryAvatarProjectResponseBodyFramesLayers {
	s.PositionX = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) SetPositionY(v int32) *QueryAvatarProjectResponseBodyFramesLayers {
	s.PositionY = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) SetType(v string) *QueryAvatarProjectResponseBodyFramesLayers {
	s.Type = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) SetWidth(v int32) *QueryAvatarProjectResponseBodyFramesLayers {
	s.Width = &v
	return s
}

type QueryAvatarProjectResponseBodyFramesLayersMaterial struct {
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	Id     *string `json:"id,omitempty" xml:"id,omitempty"`
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s QueryAvatarProjectResponseBodyFramesLayersMaterial) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarProjectResponseBodyFramesLayersMaterial) GoString() string {
	return s.String()
}

func (s *QueryAvatarProjectResponseBodyFramesLayersMaterial) SetFormat(v string) *QueryAvatarProjectResponseBodyFramesLayersMaterial {
	s.Format = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayersMaterial) SetId(v string) *QueryAvatarProjectResponseBodyFramesLayersMaterial {
	s.Id = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayersMaterial) SetUrl(v string) *QueryAvatarProjectResponseBodyFramesLayersMaterial {
	s.Url = &v
	return s
}

type QueryAvatarProjectResponseBodyFramesVideoScript struct {
	SpeedRate       *string `json:"speedRate,omitempty" xml:"speedRate,omitempty"`
	VoiceTemplateId *string `json:"voiceTemplateId,omitempty" xml:"voiceTemplateId,omitempty"`
}

func (s QueryAvatarProjectResponseBodyFramesVideoScript) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarProjectResponseBodyFramesVideoScript) GoString() string {
	return s.String()
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) SetSpeedRate(v string) *QueryAvatarProjectResponseBodyFramesVideoScript {
	s.SpeedRate = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) SetVoiceTemplateId(v string) *QueryAvatarProjectResponseBodyFramesVideoScript {
	s.VoiceTemplateId = &v
	return s
}

type QueryAvatarProjectResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAvatarProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAvatarProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarProjectResponse) GoString() string {
	return s.String()
}

func (s *QueryAvatarProjectResponse) SetHeaders(v map[string]*string) *QueryAvatarProjectResponse {
	s.Headers = v
	return s
}

func (s *QueryAvatarProjectResponse) SetStatusCode(v int32) *QueryAvatarProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAvatarProjectResponse) SetBody(v *QueryAvatarProjectResponseBody) *QueryAvatarProjectResponse {
	s.Body = v
	return s
}

type QueryAvatarResourceRequest struct {
	// example:
	//
	// 11111
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
}

func (s QueryAvatarResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarResourceRequest) GoString() string {
	return s.String()
}

func (s *QueryAvatarResourceRequest) SetIdempotentId(v string) *QueryAvatarResourceRequest {
	s.IdempotentId = &v
	return s
}

type QueryAvatarResourceResponseBody struct {
	QueryResourceInfoList []*QueryAvatarResourceResponseBodyQueryResourceInfoList `json:"queryResourceInfoList,omitempty" xml:"queryResourceInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// D5798660-1531-5D12-9C20-16FEE9D22351
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryAvatarResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarResourceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAvatarResourceResponseBody) SetQueryResourceInfoList(v []*QueryAvatarResourceResponseBodyQueryResourceInfoList) *QueryAvatarResourceResponseBody {
	s.QueryResourceInfoList = v
	return s
}

func (s *QueryAvatarResourceResponseBody) SetRequestId(v string) *QueryAvatarResourceResponseBody {
	s.RequestId = &v
	return s
}

type QueryAvatarResourceResponseBodyQueryResourceInfoList struct {
	// example:
	//
	// 21275
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// STANDARD
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1719904342237
	ValidPeriodTime *string `json:"validPeriodTime,omitempty" xml:"validPeriodTime,omitempty"`
}

func (s QueryAvatarResourceResponseBodyQueryResourceInfoList) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarResourceResponseBodyQueryResourceInfoList) GoString() string {
	return s.String()
}

func (s *QueryAvatarResourceResponseBodyQueryResourceInfoList) SetResourceId(v string) *QueryAvatarResourceResponseBodyQueryResourceInfoList {
	s.ResourceId = &v
	return s
}

func (s *QueryAvatarResourceResponseBodyQueryResourceInfoList) SetType(v string) *QueryAvatarResourceResponseBodyQueryResourceInfoList {
	s.Type = &v
	return s
}

func (s *QueryAvatarResourceResponseBodyQueryResourceInfoList) SetValidPeriodTime(v string) *QueryAvatarResourceResponseBodyQueryResourceInfoList {
	s.ValidPeriodTime = &v
	return s
}

type QueryAvatarResourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAvatarResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAvatarResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarResourceResponse) GoString() string {
	return s.String()
}

func (s *QueryAvatarResourceResponse) SetHeaders(v map[string]*string) *QueryAvatarResourceResponse {
	s.Headers = v
	return s
}

func (s *QueryAvatarResourceResponse) SetStatusCode(v int32) *QueryAvatarResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAvatarResourceResponse) SetBody(v *QueryAvatarResourceResponseBody) *QueryAvatarResourceResponse {
	s.Body = v
	return s
}

type QuerySessionInfoRequest struct {
	// example:
	//
	// 1
	PageNo *int32 `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 805800890535673856
	ProjectId  *string   `json:"projectId,omitempty" xml:"projectId,omitempty"`
	StatusList []*string `json:"statusList,omitempty" xml:"statusList,omitempty" type:"Repeated"`
}

func (s QuerySessionInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySessionInfoRequest) GoString() string {
	return s.String()
}

func (s *QuerySessionInfoRequest) SetPageNo(v int32) *QuerySessionInfoRequest {
	s.PageNo = &v
	return s
}

func (s *QuerySessionInfoRequest) SetPageSize(v int32) *QuerySessionInfoRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySessionInfoRequest) SetProjectId(v string) *QuerySessionInfoRequest {
	s.ProjectId = &v
	return s
}

func (s *QuerySessionInfoRequest) SetStatusList(v []*string) *QuerySessionInfoRequest {
	s.StatusList = v
	return s
}

type QuerySessionInfoShrinkRequest struct {
	// example:
	//
	// 1
	PageNo *int32 `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 805800890535673856
	ProjectId        *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	StatusListShrink *string `json:"statusList,omitempty" xml:"statusList,omitempty"`
}

func (s QuerySessionInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySessionInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuerySessionInfoShrinkRequest) SetPageNo(v int32) *QuerySessionInfoShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *QuerySessionInfoShrinkRequest) SetPageSize(v int32) *QuerySessionInfoShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySessionInfoShrinkRequest) SetProjectId(v string) *QuerySessionInfoShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *QuerySessionInfoShrinkRequest) SetStatusListShrink(v string) *QuerySessionInfoShrinkRequest {
	s.StatusListShrink = &v
	return s
}

type QuerySessionInfoResponseBody struct {
	QueryResourceInfoList []*QuerySessionInfoResponseBodyQueryResourceInfoList `json:"queryResourceInfoList,omitempty" xml:"queryResourceInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 4D902811-B75C-5D1B-8882-D515F8E2F977
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 26
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s QuerySessionInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySessionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySessionInfoResponseBody) SetQueryResourceInfoList(v []*QuerySessionInfoResponseBodyQueryResourceInfoList) *QuerySessionInfoResponseBody {
	s.QueryResourceInfoList = v
	return s
}

func (s *QuerySessionInfoResponseBody) SetRequestId(v string) *QuerySessionInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySessionInfoResponseBody) SetTotal(v int64) *QuerySessionInfoResponseBody {
	s.Total = &v
	return s
}

type QuerySessionInfoResponseBodyQueryResourceInfoList struct {
	// example:
	//
	// a169e9ec18404edc9972afd80866dc97
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// FREE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QuerySessionInfoResponseBodyQueryResourceInfoList) String() string {
	return tea.Prettify(s)
}

func (s QuerySessionInfoResponseBodyQueryResourceInfoList) GoString() string {
	return s.String()
}

func (s *QuerySessionInfoResponseBodyQueryResourceInfoList) SetSessionId(v string) *QuerySessionInfoResponseBodyQueryResourceInfoList {
	s.SessionId = &v
	return s
}

func (s *QuerySessionInfoResponseBodyQueryResourceInfoList) SetStatus(v string) *QuerySessionInfoResponseBodyQueryResourceInfoList {
	s.Status = &v
	return s
}

type QuerySessionInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySessionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySessionInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySessionInfoResponse) GoString() string {
	return s.String()
}

func (s *QuerySessionInfoResponse) SetHeaders(v map[string]*string) *QuerySessionInfoResponse {
	s.Headers = v
	return s
}

func (s *QuerySessionInfoResponse) SetStatusCode(v int32) *QuerySessionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySessionInfoResponse) SetBody(v *QuerySessionInfoResponseBody) *QuerySessionInfoResponse {
	s.Body = v
	return s
}

type QueryTextStreamResponseBody struct {
	// example:
	//
	// false
	End *bool `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 1
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// Id of the request
	//
	// example:
	//
	// None
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryTextStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTextStreamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTextStreamResponseBody) SetEnd(v bool) *QueryTextStreamResponseBody {
	s.End = &v
	return s
}

func (s *QueryTextStreamResponseBody) SetIndex(v int32) *QueryTextStreamResponseBody {
	s.Index = &v
	return s
}

func (s *QueryTextStreamResponseBody) SetMessage(v string) *QueryTextStreamResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTextStreamResponseBody) SetType(v int32) *QueryTextStreamResponseBody {
	s.Type = &v
	return s
}

type QueryTextStreamResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTextStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTextStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTextStreamResponse) GoString() string {
	return s.String()
}

func (s *QueryTextStreamResponse) SetHeaders(v map[string]*string) *QueryTextStreamResponse {
	s.Headers = v
	return s
}

func (s *QueryTextStreamResponse) SetStatusCode(v int32) *QueryTextStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTextStreamResponse) SetBody(v *QueryTextStreamResponseBody) *QueryTextStreamResponse {
	s.Body = v
	return s
}

type SaveAvatarProjectRequest struct {
	// example:
	//
	// 1000196
	AgentId *string                           `json:"agentId,omitempty" xml:"agentId,omitempty"`
	Frames  []*SaveAvatarProjectRequestFrames `json:"frames,omitempty" xml:"frames,omitempty" type:"Repeated"`
	// example:
	//
	// CREATE
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// example:
	//
	// 787594567117586432
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// df_cs_471437
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// example:
	//
	// STANDARD
	ResSpecType *string `json:"resSpecType,omitempty" xml:"resSpecType,omitempty"`
	// example:
	//
	// 9:16
	ScaleType *string `json:"scaleType,omitempty" xml:"scaleType,omitempty"`
}

func (s SaveAvatarProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveAvatarProjectRequest) GoString() string {
	return s.String()
}

func (s *SaveAvatarProjectRequest) SetAgentId(v string) *SaveAvatarProjectRequest {
	s.AgentId = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetFrames(v []*SaveAvatarProjectRequestFrames) *SaveAvatarProjectRequest {
	s.Frames = v
	return s
}

func (s *SaveAvatarProjectRequest) SetOperateType(v string) *SaveAvatarProjectRequest {
	s.OperateType = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetProjectId(v string) *SaveAvatarProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetProjectName(v string) *SaveAvatarProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetResSpecType(v string) *SaveAvatarProjectRequest {
	s.ResSpecType = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetScaleType(v string) *SaveAvatarProjectRequest {
	s.ScaleType = &v
	return s
}

type SaveAvatarProjectRequestFrames struct {
	Layers      []*SaveAvatarProjectRequestFramesLayers    `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	VideoScript *SaveAvatarProjectRequestFramesVideoScript `json:"videoScript,omitempty" xml:"videoScript,omitempty" type:"Struct"`
}

func (s SaveAvatarProjectRequestFrames) String() string {
	return tea.Prettify(s)
}

func (s SaveAvatarProjectRequestFrames) GoString() string {
	return s.String()
}

func (s *SaveAvatarProjectRequestFrames) SetLayers(v []*SaveAvatarProjectRequestFramesLayers) *SaveAvatarProjectRequestFrames {
	s.Layers = v
	return s
}

func (s *SaveAvatarProjectRequestFrames) SetVideoScript(v *SaveAvatarProjectRequestFramesVideoScript) *SaveAvatarProjectRequestFrames {
	s.VideoScript = v
	return s
}

type SaveAvatarProjectRequestFramesLayers struct {
	// example:
	//
	// 100
	Height   *int32                                        `json:"height,omitempty" xml:"height,omitempty"`
	Material *SaveAvatarProjectRequestFramesLayersMaterial `json:"material,omitempty" xml:"material,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PositionX *int32 `json:"positionX,omitempty" xml:"positionX,omitempty"`
	// example:
	//
	// 1
	PositionY *int32 `json:"positionY,omitempty" xml:"positionY,omitempty"`
	// example:
	//
	// ANCHOR
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 100
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s SaveAvatarProjectRequestFramesLayers) String() string {
	return tea.Prettify(s)
}

func (s SaveAvatarProjectRequestFramesLayers) GoString() string {
	return s.String()
}

func (s *SaveAvatarProjectRequestFramesLayers) SetHeight(v int32) *SaveAvatarProjectRequestFramesLayers {
	s.Height = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayers) SetMaterial(v *SaveAvatarProjectRequestFramesLayersMaterial) *SaveAvatarProjectRequestFramesLayers {
	s.Material = v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayers) SetPositionX(v int32) *SaveAvatarProjectRequestFramesLayers {
	s.PositionX = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayers) SetPositionY(v int32) *SaveAvatarProjectRequestFramesLayers {
	s.PositionY = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayers) SetType(v string) *SaveAvatarProjectRequestFramesLayers {
	s.Type = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayers) SetWidth(v int32) *SaveAvatarProjectRequestFramesLayers {
	s.Width = &v
	return s
}

type SaveAvatarProjectRequestFramesLayersMaterial struct {
	// example:
	//
	// image/png
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// example:
	//
	// 434508
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// https://alidocs.dingtalk.com/i/nodes/vy20BglGWOxjGpq0C5G4DlN0VA7depqY
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SaveAvatarProjectRequestFramesLayersMaterial) String() string {
	return tea.Prettify(s)
}

func (s SaveAvatarProjectRequestFramesLayersMaterial) GoString() string {
	return s.String()
}

func (s *SaveAvatarProjectRequestFramesLayersMaterial) SetFormat(v string) *SaveAvatarProjectRequestFramesLayersMaterial {
	s.Format = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayersMaterial) SetId(v string) *SaveAvatarProjectRequestFramesLayersMaterial {
	s.Id = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayersMaterial) SetUrl(v string) *SaveAvatarProjectRequestFramesLayersMaterial {
	s.Url = &v
	return s
}

type SaveAvatarProjectRequestFramesVideoScript struct {
	// example:
	//
	// 1.0
	SpeedRate *string `json:"speedRate,omitempty" xml:"speedRate,omitempty"`
	// example:
	//
	// 1
	VoiceTemplateId *string `json:"voiceTemplateId,omitempty" xml:"voiceTemplateId,omitempty"`
	// example:
	//
	// 50
	Volume *string `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s SaveAvatarProjectRequestFramesVideoScript) String() string {
	return tea.Prettify(s)
}

func (s SaveAvatarProjectRequestFramesVideoScript) GoString() string {
	return s.String()
}

func (s *SaveAvatarProjectRequestFramesVideoScript) SetSpeedRate(v string) *SaveAvatarProjectRequestFramesVideoScript {
	s.SpeedRate = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesVideoScript) SetVoiceTemplateId(v string) *SaveAvatarProjectRequestFramesVideoScript {
	s.VoiceTemplateId = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesVideoScript) SetVolume(v string) *SaveAvatarProjectRequestFramesVideoScript {
	s.Volume = &v
	return s
}

type SaveAvatarProjectResponseBody struct {
	// example:
	//
	// 1000222
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// 040002
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Failed to proxy flink ui request, message: An error occurred: Invalid UUID string: jobsn
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 812907463682949120
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// doc_test_3
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// example:
	//
	// 0E8B1746-AE35-5C4B-A3A8-345B274AE32C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// DRAFT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SaveAvatarProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveAvatarProjectResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAvatarProjectResponseBody) SetAgentId(v string) *SaveAvatarProjectResponseBody {
	s.AgentId = &v
	return s
}

func (s *SaveAvatarProjectResponseBody) SetErrorCode(v string) *SaveAvatarProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SaveAvatarProjectResponseBody) SetErrorMessage(v string) *SaveAvatarProjectResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SaveAvatarProjectResponseBody) SetErrorMsg(v string) *SaveAvatarProjectResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SaveAvatarProjectResponseBody) SetProjectId(v string) *SaveAvatarProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *SaveAvatarProjectResponseBody) SetProjectName(v string) *SaveAvatarProjectResponseBody {
	s.ProjectName = &v
	return s
}

func (s *SaveAvatarProjectResponseBody) SetRequestId(v string) *SaveAvatarProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveAvatarProjectResponseBody) SetStatus(v string) *SaveAvatarProjectResponseBody {
	s.Status = &v
	return s
}

type SaveAvatarProjectResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveAvatarProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveAvatarProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveAvatarProjectResponse) GoString() string {
	return s.String()
}

func (s *SaveAvatarProjectResponse) SetHeaders(v map[string]*string) *SaveAvatarProjectResponse {
	s.Headers = v
	return s
}

func (s *SaveAvatarProjectResponse) SetStatusCode(v int32) *SaveAvatarProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveAvatarProjectResponse) SetBody(v *SaveAvatarProjectResponseBody) *SaveAvatarProjectResponse {
	s.Body = v
	return s
}

type SelectImageTaskResponseBody struct {
	// example:
	//
	// Failed to proxy flink ui request, message: An error occurred: Invalid UUID string: jobsn
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 1
	Failed *string `json:"failed,omitempty" xml:"failed,omitempty"`
	// example:
	//
	// PLATFORM
	GenerationSource *string `json:"generationSource,omitempty" xml:"generationSource,omitempty"`
	// example:
	//
	// 1
	GmtCreate  *string                                  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	ImageInfos []*SelectImageTaskResponseBodyImageInfos `json:"imageInfos,omitempty" xml:"imageInfos,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 0E8B1746-AE35-5C4B-A3A8-345B274AE32C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// Successed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1
	SubtaskProcessing *string `json:"subtaskProcessing,omitempty" xml:"subtaskProcessing,omitempty"`
	// example:
	//
	// 1
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 10
	Total *string `json:"total,omitempty" xml:"total,omitempty"`
}

func (s SelectImageTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SelectImageTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SelectImageTaskResponseBody) SetErrorMessage(v string) *SelectImageTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetFailed(v string) *SelectImageTaskResponseBody {
	s.Failed = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetGenerationSource(v string) *SelectImageTaskResponseBody {
	s.GenerationSource = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetGmtCreate(v string) *SelectImageTaskResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetImageInfos(v []*SelectImageTaskResponseBodyImageInfos) *SelectImageTaskResponseBody {
	s.ImageInfos = v
	return s
}

func (s *SelectImageTaskResponseBody) SetRequestId(v string) *SelectImageTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetScene(v string) *SelectImageTaskResponseBody {
	s.Scene = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetStatus(v string) *SelectImageTaskResponseBody {
	s.Status = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetSubtaskProcessing(v string) *SelectImageTaskResponseBody {
	s.SubtaskProcessing = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetSuccess(v string) *SelectImageTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SelectImageTaskResponseBody) SetTotal(v string) *SelectImageTaskResponseBody {
	s.Total = &v
	return s
}

type SelectImageTaskResponseBodyImageInfos struct {
	// example:
	//
	// www.ali.com
	CustomImageUrl *string `json:"customImageUrl,omitempty" xml:"customImageUrl,omitempty"`
	// example:
	//
	// 1
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 500
	ImageH *string `json:"imageH,omitempty" xml:"imageH,omitempty"`
	// example:
	//
	// 500
	ImageW *string `json:"imageW,omitempty" xml:"imageW,omitempty"`
}

func (s SelectImageTaskResponseBodyImageInfos) String() string {
	return tea.Prettify(s)
}

func (s SelectImageTaskResponseBodyImageInfos) GoString() string {
	return s.String()
}

func (s *SelectImageTaskResponseBodyImageInfos) SetCustomImageUrl(v string) *SelectImageTaskResponseBodyImageInfos {
	s.CustomImageUrl = &v
	return s
}

func (s *SelectImageTaskResponseBodyImageInfos) SetGmtCreate(v string) *SelectImageTaskResponseBodyImageInfos {
	s.GmtCreate = &v
	return s
}

func (s *SelectImageTaskResponseBodyImageInfos) SetImageH(v string) *SelectImageTaskResponseBodyImageInfos {
	s.ImageH = &v
	return s
}

func (s *SelectImageTaskResponseBodyImageInfos) SetImageW(v string) *SelectImageTaskResponseBodyImageInfos {
	s.ImageW = &v
	return s
}

type SelectImageTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SelectImageTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SelectImageTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SelectImageTaskResponse) GoString() string {
	return s.String()
}

func (s *SelectImageTaskResponse) SetHeaders(v map[string]*string) *SelectImageTaskResponse {
	s.Headers = v
	return s
}

func (s *SelectImageTaskResponse) SetStatusCode(v int32) *SelectImageTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SelectImageTaskResponse) SetBody(v *SelectImageTaskResponseBody) *SelectImageTaskResponse {
	s.Body = v
	return s
}

type SelectResourceRequest struct {
	// example:
	//
	// 1111
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
}

func (s SelectResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s SelectResourceRequest) GoString() string {
	return s.String()
}

func (s *SelectResourceRequest) SetIdempotentId(v string) *SelectResourceRequest {
	s.IdempotentId = &v
	return s
}

type SelectResourceResponseBody struct {
	// example:
	//
	// 0E8B1746-AE35-5C4B-A3A8-345B274AE32C
	RequestId        *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResourceInfoList []*SelectResourceResponseBodyResourceInfoList `json:"resourceInfoList,omitempty" xml:"resourceInfoList,omitempty" type:"Repeated"`
}

func (s SelectResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SelectResourceResponseBody) GoString() string {
	return s.String()
}

func (s *SelectResourceResponseBody) SetRequestId(v string) *SelectResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SelectResourceResponseBody) SetResourceInfoList(v []*SelectResourceResponseBodyResourceInfoList) *SelectResourceResponseBody {
	s.ResourceInfoList = v
	return s
}

type SelectResourceResponseBodyResourceInfoList struct {
	// example:
	//
	// 111
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// example:
	//
	// 11
	LastExpire *int32 `json:"lastExpire,omitempty" xml:"lastExpire,omitempty"`
	// example:
	//
	// 1249
	RemainCount *int32 `json:"remainCount,omitempty" xml:"remainCount,omitempty"`
	// example:
	//
	// 2
	ResourceType *int32 `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// second
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s SelectResourceResponseBodyResourceInfoList) String() string {
	return tea.Prettify(s)
}

func (s SelectResourceResponseBodyResourceInfoList) GoString() string {
	return s.String()
}

func (s *SelectResourceResponseBodyResourceInfoList) SetExpireTime(v string) *SelectResourceResponseBodyResourceInfoList {
	s.ExpireTime = &v
	return s
}

func (s *SelectResourceResponseBodyResourceInfoList) SetLastExpire(v int32) *SelectResourceResponseBodyResourceInfoList {
	s.LastExpire = &v
	return s
}

func (s *SelectResourceResponseBodyResourceInfoList) SetRemainCount(v int32) *SelectResourceResponseBodyResourceInfoList {
	s.RemainCount = &v
	return s
}

func (s *SelectResourceResponseBodyResourceInfoList) SetResourceType(v int32) *SelectResourceResponseBodyResourceInfoList {
	s.ResourceType = &v
	return s
}

func (s *SelectResourceResponseBodyResourceInfoList) SetUnit(v string) *SelectResourceResponseBodyResourceInfoList {
	s.Unit = &v
	return s
}

type SelectResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SelectResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SelectResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s SelectResourceResponse) GoString() string {
	return s.String()
}

func (s *SelectResourceResponse) SetHeaders(v map[string]*string) *SelectResourceResponse {
	s.Headers = v
	return s
}

func (s *SelectResourceResponse) SetStatusCode(v int32) *SelectResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *SelectResourceResponse) SetBody(v *SelectResourceResponseBody) *SelectResourceResponse {
	s.Body = v
	return s
}

type SendTextMsgRequest struct {
	// example:
	//
	// 126000030
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 52775239-1575-5C07-A4AE-1835D120E4A6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// tcm9xac9dsfbfgm8hf5k94l3cqybwh9o3mn0iuyytdgd9qoejxf1crxsdvuvr8fu0zudk5px4vsa3e3fgcclplkiuo7kyy3sqgscvhejmooblaiv64ww8cvlxvin2urzyhooqj33y7gvodef0sxn22n9q58o7xlupabiknxsv46qe7kof8nuc4be8kyhi01
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SendTextMsgRequest) String() string {
	return tea.Prettify(s)
}

func (s SendTextMsgRequest) GoString() string {
	return s.String()
}

func (s *SendTextMsgRequest) SetProjectId(v string) *SendTextMsgRequest {
	s.ProjectId = &v
	return s
}

func (s *SendTextMsgRequest) SetRequestId(v string) *SendTextMsgRequest {
	s.RequestId = &v
	return s
}

func (s *SendTextMsgRequest) SetSessionId(v string) *SendTextMsgRequest {
	s.SessionId = &v
	return s
}

func (s *SendTextMsgRequest) SetText(v string) *SendTextMsgRequest {
	s.Text = &v
	return s
}

func (s *SendTextMsgRequest) SetType(v int32) *SendTextMsgRequest {
	s.Type = &v
	return s
}

type SendTextMsgResponseBody struct {
	// example:
	//
	// 827BF714-19E7-51B5-A434-C21BFEE05983
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SendTextMsgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendTextMsgResponseBody) GoString() string {
	return s.String()
}

func (s *SendTextMsgResponseBody) SetRequestId(v string) *SendTextMsgResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendTextMsgResponseBody) SetStatus(v string) *SendTextMsgResponseBody {
	s.Status = &v
	return s
}

type SendTextMsgResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendTextMsgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendTextMsgResponse) String() string {
	return tea.Prettify(s)
}

func (s SendTextMsgResponse) GoString() string {
	return s.String()
}

func (s *SendTextMsgResponse) SetHeaders(v map[string]*string) *SendTextMsgResponse {
	s.Headers = v
	return s
}

func (s *SendTextMsgResponse) SetStatusCode(v int32) *SendTextMsgResponse {
	s.StatusCode = &v
	return s
}

func (s *SendTextMsgResponse) SetBody(v *SendTextMsgResponseBody) *SendTextMsgResponse {
	s.Body = v
	return s
}

type StartAvatarSessionRequest struct {
	CustomPushUrl *string `json:"customPushUrl,omitempty" xml:"customPushUrl,omitempty"`
	// example:
	//
	// 13534711288320
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 15ED6083-B0B8-5B2A-BEDB-94A5C687C812
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartAvatarSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s StartAvatarSessionRequest) GoString() string {
	return s.String()
}

func (s *StartAvatarSessionRequest) SetCustomPushUrl(v string) *StartAvatarSessionRequest {
	s.CustomPushUrl = &v
	return s
}

func (s *StartAvatarSessionRequest) SetProjectId(v string) *StartAvatarSessionRequest {
	s.ProjectId = &v
	return s
}

func (s *StartAvatarSessionRequest) SetRequestId(v string) *StartAvatarSessionRequest {
	s.RequestId = &v
	return s
}

type StartAvatarSessionResponseBody struct {
	ChannelToken *string `json:"channelToken,omitempty" xml:"channelToken,omitempty"`
	// example:
	//
	// 0E8B1746-AE35-5C4B-A3A8-345B274AE32C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 121dlsga4o7golrl1hoja
	SessionId    *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	Token        *string `json:"token,omitempty" xml:"token,omitempty"`
	WebSocketUrl *string `json:"webSocketUrl,omitempty" xml:"webSocketUrl,omitempty"`
}

func (s StartAvatarSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartAvatarSessionResponseBody) GoString() string {
	return s.String()
}

func (s *StartAvatarSessionResponseBody) SetChannelToken(v string) *StartAvatarSessionResponseBody {
	s.ChannelToken = &v
	return s
}

func (s *StartAvatarSessionResponseBody) SetRequestId(v string) *StartAvatarSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAvatarSessionResponseBody) SetSessionId(v string) *StartAvatarSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *StartAvatarSessionResponseBody) SetToken(v string) *StartAvatarSessionResponseBody {
	s.Token = &v
	return s
}

func (s *StartAvatarSessionResponseBody) SetWebSocketUrl(v string) *StartAvatarSessionResponseBody {
	s.WebSocketUrl = &v
	return s
}

type StartAvatarSessionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAvatarSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAvatarSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s StartAvatarSessionResponse) GoString() string {
	return s.String()
}

func (s *StartAvatarSessionResponse) SetHeaders(v map[string]*string) *StartAvatarSessionResponse {
	s.Headers = v
	return s
}

func (s *StartAvatarSessionResponse) SetStatusCode(v int32) *StartAvatarSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAvatarSessionResponse) SetBody(v *StartAvatarSessionResponseBody) *StartAvatarSessionResponse {
	s.Body = v
	return s
}

type StopAvatarSessionRequest struct {
	// example:
	//
	// 124900036
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 121dlsga4o7golrl1hojazg0u9lvysk0uyczgd79be2a4hkr9ijrblmb5qohi5iaja3p5j633doqj4t2uu3sek2i49hzkao0bli4bch4tnloyx22odd7sot9dxl5xfd0hbp7fl9dehnqofkb9csebf0nuezj8bwgec8ei6dby0encu5y88ky6oqensuqnj
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s StopAvatarSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAvatarSessionRequest) GoString() string {
	return s.String()
}

func (s *StopAvatarSessionRequest) SetProjectId(v string) *StopAvatarSessionRequest {
	s.ProjectId = &v
	return s
}

func (s *StopAvatarSessionRequest) SetSessionId(v string) *StopAvatarSessionRequest {
	s.SessionId = &v
	return s
}

type StopAvatarSessionResponseBody struct {
	// example:
	//
	// 725E87CD-F2DE-5FC4-8A09-2EBDFBF26DAA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// Stopped
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s StopAvatarSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopAvatarSessionResponseBody) GoString() string {
	return s.String()
}

func (s *StopAvatarSessionResponseBody) SetRequestId(v string) *StopAvatarSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopAvatarSessionResponseBody) SetStatus(v string) *StopAvatarSessionResponseBody {
	s.Status = &v
	return s
}

type StopAvatarSessionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAvatarSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAvatarSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s StopAvatarSessionResponse) GoString() string {
	return s.String()
}

func (s *StopAvatarSessionResponse) SetHeaders(v map[string]*string) *StopAvatarSessionResponse {
	s.Headers = v
	return s
}

func (s *StopAvatarSessionResponse) SetStatusCode(v int32) *StopAvatarSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAvatarSessionResponse) SetBody(v *StopAvatarSessionResponseBody) *StopAvatarSessionResponse {
	s.Body = v
	return s
}

type StopProjectTaskRequest struct {
	// example:
	//
	// 1111111
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s StopProjectTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StopProjectTaskRequest) GoString() string {
	return s.String()
}

func (s *StopProjectTaskRequest) SetTaskId(v string) *StopProjectTaskRequest {
	s.TaskId = &v
	return s
}

type StopProjectTaskResponseBody struct {
	// example:
	//
	// 86A90C40-D1AB-50DA-A4B1-0D545F80F2FE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StopProjectTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopProjectTaskResponseBody) SetRequestId(v string) *StopProjectTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopProjectTaskResponseBody) SetSuccess(v bool) *StopProjectTaskResponseBody {
	s.Success = &v
	return s
}

type StopProjectTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopProjectTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopProjectTaskResponse) GoString() string {
	return s.String()
}

func (s *StopProjectTaskResponse) SetHeaders(v map[string]*string) *StopProjectTaskResponse {
	s.Headers = v
	return s
}

func (s *StopProjectTaskResponse) SetStatusCode(v int32) *StopProjectTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopProjectTaskResponse) SetBody(v *StopProjectTaskResponseBody) *StopProjectTaskResponse {
	s.Body = v
	return s
}

type SubmitProjectTaskRequest struct {
	// frame
	Frames []*SubmitProjectTaskRequestFrames `json:"frames,omitempty" xml:"frames,omitempty" type:"Repeated"`
	// example:
	//
	// 9:16
	ScaleType *string `json:"scaleType,omitempty" xml:"scaleType,omitempty"`
	// example:
	//
	// 1
	SubtitleTag           *int32 `json:"subtitleTag,omitempty" xml:"subtitleTag,omitempty"`
	TransparentBackground *int32 `json:"transparentBackground,omitempty" xml:"transparentBackground,omitempty"`
}

func (s SubmitProjectTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitProjectTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskRequest) SetFrames(v []*SubmitProjectTaskRequestFrames) *SubmitProjectTaskRequest {
	s.Frames = v
	return s
}

func (s *SubmitProjectTaskRequest) SetScaleType(v string) *SubmitProjectTaskRequest {
	s.ScaleType = &v
	return s
}

func (s *SubmitProjectTaskRequest) SetSubtitleTag(v int32) *SubmitProjectTaskRequest {
	s.SubtitleTag = &v
	return s
}

func (s *SubmitProjectTaskRequest) SetTransparentBackground(v int32) *SubmitProjectTaskRequest {
	s.TransparentBackground = &v
	return s
}

type SubmitProjectTaskRequestFrames struct {
	// example:
	//
	// 1
	Index       *int32                                     `json:"index,omitempty" xml:"index,omitempty"`
	Layers      []*SubmitProjectTaskRequestFramesLayers    `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	Subtitle    *SubmitProjectTaskRequestFramesSubtitle    `json:"subtitle,omitempty" xml:"subtitle,omitempty" type:"Struct"`
	VideoScript *SubmitProjectTaskRequestFramesVideoScript `json:"videoScript,omitempty" xml:"videoScript,omitempty" type:"Struct"`
}

func (s SubmitProjectTaskRequestFrames) String() string {
	return tea.Prettify(s)
}

func (s SubmitProjectTaskRequestFrames) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskRequestFrames) SetIndex(v int32) *SubmitProjectTaskRequestFrames {
	s.Index = &v
	return s
}

func (s *SubmitProjectTaskRequestFrames) SetLayers(v []*SubmitProjectTaskRequestFramesLayers) *SubmitProjectTaskRequestFrames {
	s.Layers = v
	return s
}

func (s *SubmitProjectTaskRequestFrames) SetSubtitle(v *SubmitProjectTaskRequestFramesSubtitle) *SubmitProjectTaskRequestFrames {
	s.Subtitle = v
	return s
}

func (s *SubmitProjectTaskRequestFrames) SetVideoScript(v *SubmitProjectTaskRequestFramesVideoScript) *SubmitProjectTaskRequestFrames {
	s.VideoScript = v
	return s
}

type SubmitProjectTaskRequestFramesLayers struct {
	// example:
	//
	// 222
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// 1
	Index    *int32                                        `json:"index,omitempty" xml:"index,omitempty"`
	Material *SubmitProjectTaskRequestFramesLayersMaterial `json:"material,omitempty" xml:"material,omitempty" type:"Struct"`
	// example:
	//
	// 11
	PositionX *int32 `json:"positionX,omitempty" xml:"positionX,omitempty"`
	// example:
	//
	// 22
	PositionY *int32 `json:"positionY,omitempty" xml:"positionY,omitempty"`
	// example:
	//
	// ANCHOR
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 111
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s SubmitProjectTaskRequestFramesLayers) String() string {
	return tea.Prettify(s)
}

func (s SubmitProjectTaskRequestFramesLayers) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskRequestFramesLayers) SetHeight(v int32) *SubmitProjectTaskRequestFramesLayers {
	s.Height = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayers) SetIndex(v int32) *SubmitProjectTaskRequestFramesLayers {
	s.Index = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayers) SetMaterial(v *SubmitProjectTaskRequestFramesLayersMaterial) *SubmitProjectTaskRequestFramesLayers {
	s.Material = v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayers) SetPositionX(v int32) *SubmitProjectTaskRequestFramesLayers {
	s.PositionX = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayers) SetPositionY(v int32) *SubmitProjectTaskRequestFramesLayers {
	s.PositionY = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayers) SetType(v string) *SubmitProjectTaskRequestFramesLayers {
	s.Type = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayers) SetWidth(v int32) *SubmitProjectTaskRequestFramesLayers {
	s.Width = &v
	return s
}

type SubmitProjectTaskRequestFramesLayersMaterial struct {
	// example:
	//
	// video/mp4
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// example:
	//
	// 38863
	Id    *string `json:"id,omitempty" xml:"id,omitempty"`
	Speed *string `json:"speed,omitempty" xml:"speed,omitempty"`
	// example:
	//
	// https://meeting.dingtalk.com/j/1COFppy0POR
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
	Volume *int32  `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s SubmitProjectTaskRequestFramesLayersMaterial) String() string {
	return tea.Prettify(s)
}

func (s SubmitProjectTaskRequestFramesLayersMaterial) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetFormat(v string) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.Format = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetId(v string) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.Id = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetSpeed(v string) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.Speed = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetUrl(v string) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.Url = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetVolume(v int32) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.Volume = &v
	return s
}

type SubmitProjectTaskRequestFramesSubtitle struct {
	// example:
	//
	// BottomLeft
	Alignment *string `json:"alignment,omitempty" xml:"alignment,omitempty"`
	// example:
	//
	// #ffffff
	BackgroundColor *string `json:"backgroundColor,omitempty" xml:"backgroundColor,omitempty"`
	// example:
	//
	// SimSun
	Font *string `json:"font,omitempty" xml:"font,omitempty"`
	// example:
	//
	// #ffffff
	FontColor *string `json:"fontColor,omitempty" xml:"fontColor,omitempty"`
	// example:
	//
	// 32
	FontSize *int32 `json:"fontSize,omitempty" xml:"fontSize,omitempty"`
	// example:
	//
	// 11
	MaxCharLength *int32 `json:"maxCharLength,omitempty" xml:"maxCharLength,omitempty"`
	// example:
	//
	// 2
	PositionX *int32 `json:"positionX,omitempty" xml:"positionX,omitempty"`
	// example:
	//
	// 1
	PositionY *int32 `json:"positionY,omitempty" xml:"positionY,omitempty"`
	// example:
	//
	// 22
	TextHeight *int32 `json:"textHeight,omitempty" xml:"textHeight,omitempty"`
	// example:
	//
	// 11
	TextWidth *int32 `json:"textWidth,omitempty" xml:"textWidth,omitempty"`
}

func (s SubmitProjectTaskRequestFramesSubtitle) String() string {
	return tea.Prettify(s)
}

func (s SubmitProjectTaskRequestFramesSubtitle) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetAlignment(v string) *SubmitProjectTaskRequestFramesSubtitle {
	s.Alignment = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetBackgroundColor(v string) *SubmitProjectTaskRequestFramesSubtitle {
	s.BackgroundColor = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetFont(v string) *SubmitProjectTaskRequestFramesSubtitle {
	s.Font = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetFontColor(v string) *SubmitProjectTaskRequestFramesSubtitle {
	s.FontColor = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetFontSize(v int32) *SubmitProjectTaskRequestFramesSubtitle {
	s.FontSize = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetMaxCharLength(v int32) *SubmitProjectTaskRequestFramesSubtitle {
	s.MaxCharLength = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetPositionX(v int32) *SubmitProjectTaskRequestFramesSubtitle {
	s.PositionX = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetPositionY(v int32) *SubmitProjectTaskRequestFramesSubtitle {
	s.PositionY = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetTextHeight(v int32) *SubmitProjectTaskRequestFramesSubtitle {
	s.TextHeight = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetTextWidth(v int32) *SubmitProjectTaskRequestFramesSubtitle {
	s.TextWidth = &v
	return s
}

type SubmitProjectTaskRequestFramesVideoScript struct {
	// example:
	//
	// https://meeting.dingtalk.com/j/1COFppy0POR
	AudioUrl   *string `json:"audioUrl,omitempty" xml:"audioUrl,omitempty"`
	SpeechOpen *bool   `json:"speechOpen,omitempty" xml:"speechOpen,omitempty"`
	// example:
	//
	// 2.0
	SpeedRate   *string `json:"speedRate,omitempty" xml:"speedRate,omitempty"`
	TextContent *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
	// example:
	//
	// TEXT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 11
	VoiceTemplateId *int64 `json:"voiceTemplateId,omitempty" xml:"voiceTemplateId,omitempty"`
	// example:
	//
	// 20
	Volume *int32 `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s SubmitProjectTaskRequestFramesVideoScript) String() string {
	return tea.Prettify(s)
}

func (s SubmitProjectTaskRequestFramesVideoScript) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetAudioUrl(v string) *SubmitProjectTaskRequestFramesVideoScript {
	s.AudioUrl = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetSpeechOpen(v bool) *SubmitProjectTaskRequestFramesVideoScript {
	s.SpeechOpen = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetSpeedRate(v string) *SubmitProjectTaskRequestFramesVideoScript {
	s.SpeedRate = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetTextContent(v string) *SubmitProjectTaskRequestFramesVideoScript {
	s.TextContent = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetType(v string) *SubmitProjectTaskRequestFramesVideoScript {
	s.Type = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetVoiceTemplateId(v int64) *SubmitProjectTaskRequestFramesVideoScript {
	s.VoiceTemplateId = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetVolume(v int32) *SubmitProjectTaskRequestFramesVideoScript {
	s.Volume = &v
	return s
}

type SubmitProjectTaskResponseBody struct {
	// example:
	//
	// 551FF252-6CFC-5DDA-9F84-9B07302385C2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 313123123
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitProjectTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskResponseBody) SetRequestId(v string) *SubmitProjectTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitProjectTaskResponseBody) SetTaskId(v string) *SubmitProjectTaskResponseBody {
	s.TaskId = &v
	return s
}

type SubmitProjectTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitProjectTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitProjectTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskResponse) SetHeaders(v map[string]*string) *SubmitProjectTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitProjectTaskResponse) SetStatusCode(v int32) *SubmitProjectTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitProjectTaskResponse) SetBody(v *SubmitProjectTaskResponseBody) *SubmitProjectTaskResponse {
	s.Body = v
	return s
}

type TransferPortraitStyleRequest struct {
	// example:
	//
	// 500
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// WWW
	ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	// example:
	//
	// 4
	Numbers *int32 `json:"numbers,omitempty" xml:"numbers,omitempty"`
	// example:
	//
	// 1
	RedrawAmplitude *int32 `json:"redrawAmplitude,omitempty" xml:"redrawAmplitude,omitempty"`
	// example:
	//
	// 1
	Style *int32 `json:"style,omitempty" xml:"style,omitempty"`
	// example:
	//
	// 500
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s TransferPortraitStyleRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferPortraitStyleRequest) GoString() string {
	return s.String()
}

func (s *TransferPortraitStyleRequest) SetHeight(v int32) *TransferPortraitStyleRequest {
	s.Height = &v
	return s
}

func (s *TransferPortraitStyleRequest) SetImageUrl(v string) *TransferPortraitStyleRequest {
	s.ImageUrl = &v
	return s
}

func (s *TransferPortraitStyleRequest) SetNumbers(v int32) *TransferPortraitStyleRequest {
	s.Numbers = &v
	return s
}

func (s *TransferPortraitStyleRequest) SetRedrawAmplitude(v int32) *TransferPortraitStyleRequest {
	s.RedrawAmplitude = &v
	return s
}

func (s *TransferPortraitStyleRequest) SetStyle(v int32) *TransferPortraitStyleRequest {
	s.Style = &v
	return s
}

func (s *TransferPortraitStyleRequest) SetWidth(v int32) *TransferPortraitStyleRequest {
	s.Width = &v
	return s
}

type TransferPortraitStyleResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 725E87CD-F2DE-5FC4-8A09-2EBDFBF26DAA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 313123123
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s TransferPortraitStyleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferPortraitStyleResponseBody) GoString() string {
	return s.String()
}

func (s *TransferPortraitStyleResponseBody) SetRequestId(v string) *TransferPortraitStyleResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferPortraitStyleResponseBody) SetTaskId(v string) *TransferPortraitStyleResponseBody {
	s.TaskId = &v
	return s
}

type TransferPortraitStyleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferPortraitStyleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferPortraitStyleResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferPortraitStyleResponse) GoString() string {
	return s.String()
}

func (s *TransferPortraitStyleResponse) SetHeaders(v map[string]*string) *TransferPortraitStyleResponse {
	s.Headers = v
	return s
}

func (s *TransferPortraitStyleResponse) SetStatusCode(v int32) *TransferPortraitStyleResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferPortraitStyleResponse) SetBody(v *TransferPortraitStyleResponseBody) *TransferPortraitStyleResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("intelligentcreation"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 添加文案反馈
//
// @param request - AddTextFeedbackRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTextFeedbackResponse
func (client *Client) AddTextFeedbackWithOptions(request *AddTextFeedbackRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddTextFeedbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Quality)) {
		body["quality"] = request.Quality
	}

	if !tea.BoolValue(util.IsUnset(request.TextId)) {
		body["textId"] = request.TextId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddTextFeedback"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/addTextFeedback"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddTextFeedbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加文案反馈
//
// @param request - AddTextFeedbackRequest
//
// @return AddTextFeedbackResponse
func (client *Client) AddTextFeedback(request *AddTextFeedbackRequest) (_result *AddTextFeedbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddTextFeedbackResponse{}
	_body, _err := client.AddTextFeedbackWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询项目信息
//
// @param tmpReq - BatchGetProjectTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetProjectTaskResponse
func (client *Client) BatchGetProjectTaskWithOptions(tmpReq *BatchGetProjectTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchGetProjectTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchGetProjectTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TaskIdList)) {
		request.TaskIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIdList, tea.String("taskIdList"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskIdListShrink)) {
		query["taskIdList"] = request.TaskIdListShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchGetProjectTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/project/batchGetProjectTask"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchGetProjectTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询项目信息
//
// @param request - BatchGetProjectTaskRequest
//
// @return BatchGetProjectTaskResponse
func (client *Client) BatchGetProjectTask(request *BatchGetProjectTaskRequest) (_result *BatchGetProjectTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchGetProjectTaskResponse{}
	_body, _err := client.BatchGetProjectTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查会话状态
//
// @param request - CheckSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSessionResponse
func (client *Client) CheckSessionWithOptions(request *CheckSessionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckSession"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/avatar/project/checkSession"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查会话状态
//
// @param request - CheckSessionRequest
//
// @return CheckSessionResponse
func (client *Client) CheckSession(request *CheckSessionRequest) (_result *CheckSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckSessionResponse{}
	_body, _err := client.CheckSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 学员关闭会话
//
// @param request - CloseAICoachTaskSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseAICoachTaskSessionResponse
func (client *Client) CloseAICoachTaskSessionWithOptions(request *CloseAICoachTaskSessionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloseAICoachTaskSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		body["uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseAICoachTaskSession"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/aicoach/closeSession"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseAICoachTaskSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 学员关闭会话
//
// @param request - CloseAICoachTaskSessionRequest
//
// @return CloseAICoachTaskSessionResponse
func (client *Client) CloseAICoachTaskSession(request *CloseAICoachTaskSessionRequest) (_result *CloseAICoachTaskSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloseAICoachTaskSessionResponse{}
	_body, _err := client.CloseAICoachTaskSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文本数量统计
//
// @param request - CountTextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CountTextResponse
func (client *Client) CountTextWithOptions(request *CountTextRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CountTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GenerationSource)) {
		query["generationSource"] = request.GenerationSource
	}

	if !tea.BoolValue(util.IsUnset(request.Industry)) {
		query["industry"] = request.Industry
	}

	if !tea.BoolValue(util.IsUnset(request.PublishStatus)) {
		query["publishStatus"] = request.PublishStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Style)) {
		query["style"] = request.Style
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CountText"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/countText"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CountTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文本数量统计
//
// @param request - CountTextRequest
//
// @return CountTextResponse
func (client *Client) CountText(request *CountTextRequest) (_result *CountTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CountTextResponse{}
	_body, _err := client.CountTextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 学员开启对练会话
//
// @param request - CreateAICoachTaskSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAICoachTaskSessionResponse
func (client *Client) CreateAICoachTaskSessionWithOptions(request *CreateAICoachTaskSessionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAICoachTaskSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		body["uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAICoachTaskSession"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/aicoach/startSession"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAICoachTaskSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 学员开启对练会话
//
// @param request - CreateAICoachTaskSessionRequest
//
// @return CreateAICoachTaskSessionResponse
func (client *Client) CreateAICoachTaskSession(request *CreateAICoachTaskSessionRequest) (_result *CreateAICoachTaskSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAICoachTaskSessionResponse{}
	_body, _err := client.CreateAICoachTaskSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建配图生成任务
//
// @param request - CreateIllustrationTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIllustrationTaskResponse
func (client *Client) CreateIllustrationTaskWithOptions(textId *string, request *CreateIllustrationTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateIllustrationTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIllustrationTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/texts/" + tea.StringValue(openapiutil.GetEncodeParam(textId)) + "/illustrationTasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIllustrationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建配图生成任务
//
// @param request - CreateIllustrationTaskRequest
//
// @return CreateIllustrationTaskResponse
func (client *Client) CreateIllustrationTask(textId *string, request *CreateIllustrationTaskRequest) (_result *CreateIllustrationTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIllustrationTaskResponse{}
	_body, _err := client.CreateIllustrationTaskWithOptions(textId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 写实人像创作
//
// @param request - CreateRealisticPortraitRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRealisticPortraitResponse
func (client *Client) CreateRealisticPortraitWithOptions(request *CreateRealisticPortraitRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRealisticPortraitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ages)) {
		body["ages"] = request.Ages
	}

	if !tea.BoolValue(util.IsUnset(request.Cloth)) {
		body["cloth"] = request.Cloth
	}

	if !tea.BoolValue(util.IsUnset(request.Color)) {
		body["color"] = request.Color
	}

	if !tea.BoolValue(util.IsUnset(request.Custom)) {
		body["custom"] = request.Custom
	}

	if !tea.BoolValue(util.IsUnset(request.Face)) {
		body["face"] = request.Face
	}

	if !tea.BoolValue(util.IsUnset(request.Figure)) {
		body["figure"] = request.Figure
	}

	if !tea.BoolValue(util.IsUnset(request.Gender)) {
		body["gender"] = request.Gender
	}

	if !tea.BoolValue(util.IsUnset(request.HairColor)) {
		body["hairColor"] = request.HairColor
	}

	if !tea.BoolValue(util.IsUnset(request.Hairstyle)) {
		body["hairstyle"] = request.Hairstyle
	}

	if !tea.BoolValue(util.IsUnset(request.Height)) {
		body["height"] = request.Height
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["imageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Numbers)) {
		body["numbers"] = request.Numbers
	}

	if !tea.BoolValue(util.IsUnset(request.Ratio)) {
		body["ratio"] = request.Ratio
	}

	if !tea.BoolValue(util.IsUnset(request.Width)) {
		body["width"] = request.Width
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRealisticPortrait"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/images/portrait/realistic"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRealisticPortraitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 写实人像创作
//
// @param request - CreateRealisticPortraitRequest
//
// @return CreateRealisticPortraitResponse
func (client *Client) CreateRealisticPortrait(request *CreateRealisticPortraitRequest) (_result *CreateRealisticPortraitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRealisticPortraitResponse{}
	_body, _err := client.CreateRealisticPortraitWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建文案生成任务
//
// @param request - CreateTextTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTextTaskResponse
func (client *Client) CreateTextTaskWithOptions(request *CreateTextTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTextTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTextTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/textTasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTextTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建文案生成任务
//
// @param request - CreateTextTaskRequest
//
// @return CreateTextTaskResponse
func (client *Client) CreateTextTask(request *CreateTextTaskRequest) (_result *CreateTextTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTextTaskResponse{}
	_body, _err := client.CreateTextTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 学员完成会话
//
// @param request - FinishAICoachTaskSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FinishAICoachTaskSessionResponse
func (client *Client) FinishAICoachTaskSessionWithOptions(request *FinishAICoachTaskSessionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FinishAICoachTaskSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		body["uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FinishAICoachTaskSession"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/aicoach/finishSession"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FinishAICoachTaskSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 学员完成会话
//
// @param request - FinishAICoachTaskSessionRequest
//
// @return FinishAICoachTaskSessionResponse
func (client *Client) FinishAICoachTaskSession(request *FinishAICoachTaskSessionRequest) (_result *FinishAICoachTaskSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FinishAICoachTaskSessionResponse{}
	_body, _err := client.FinishAICoachTaskSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 学员查询会话历史
//
// @param request - GetAICoachTaskSessionHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAICoachTaskSessionHistoryResponse
func (client *Client) GetAICoachTaskSessionHistoryWithOptions(request *GetAICoachTaskSessionHistoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAICoachTaskSessionHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAICoachTaskSessionHistory"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/aicoach/querySessionHistory"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAICoachTaskSessionHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 学员查询会话历史
//
// @param request - GetAICoachTaskSessionHistoryRequest
//
// @return GetAICoachTaskSessionHistoryResponse
func (client *Client) GetAICoachTaskSessionHistory(request *GetAICoachTaskSessionHistoryRequest) (_result *GetAICoachTaskSessionHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAICoachTaskSessionHistoryResponse{}
	_body, _err := client.GetAICoachTaskSessionHistoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 学员查询会话评测报告
//
// @param request - GetAICoachTaskSessionReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAICoachTaskSessionReportResponse
func (client *Client) GetAICoachTaskSessionReportWithOptions(request *GetAICoachTaskSessionReportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAICoachTaskSessionReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAICoachTaskSessionReport"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/aicoach/queryTaskSessionReport"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAICoachTaskSessionReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 学员查询会话评测报告
//
// @param request - GetAICoachTaskSessionReportRequest
//
// @return GetAICoachTaskSessionReportResponse
func (client *Client) GetAICoachTaskSessionReport(request *GetAICoachTaskSessionReportRequest) (_result *GetAICoachTaskSessionReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAICoachTaskSessionReportResponse{}
	_body, _err := client.GetAICoachTaskSessionReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询配图
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIllustrationResponse
func (client *Client) GetIllustrationWithOptions(textId *string, illustrationId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIllustrationResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetIllustration"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/texts/" + tea.StringValue(openapiutil.GetEncodeParam(textId)) + "/illustrations/" + tea.StringValue(openapiutil.GetEncodeParam(illustrationId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIllustrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询配图
//
// @return GetIllustrationResponse
func (client *Client) GetIllustration(textId *string, illustrationId *string) (_result *GetIllustrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIllustrationResponse{}
	_body, _err := client.GetIllustrationWithOptions(textId, illustrationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询配图任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIllustrationTaskResponse
func (client *Client) GetIllustrationTaskWithOptions(textId *string, illustrationTaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIllustrationTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetIllustrationTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/texts/" + tea.StringValue(openapiutil.GetEncodeParam(textId)) + "/illustrationTasks/" + tea.StringValue(openapiutil.GetEncodeParam(illustrationTaskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIllustrationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询配图任务
//
// @return GetIllustrationTaskResponse
func (client *Client) GetIllustrationTask(textId *string, illustrationTaskId *string) (_result *GetIllustrationTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIllustrationTaskResponse{}
	_body, _err := client.GetIllustrationTaskWithOptions(textId, illustrationTaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取图片上传oss token
//
// @param request - GetOssUploadTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssUploadTokenResponse
func (client *Client) GetOssUploadTokenWithOptions(request *GetOssUploadTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetOssUploadTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		query["fileType"] = request.FileType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOssUploadToken"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/uploadToken"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOssUploadTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取图片上传oss token
//
// @param request - GetOssUploadTokenRequest
//
// @return GetOssUploadTokenResponse
func (client *Client) GetOssUploadToken(request *GetOssUploadTokenRequest) (_result *GetOssUploadTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOssUploadTokenResponse{}
	_body, _err := client.GetOssUploadTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据人合成信息
//
// @param request - GetProjectTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectTaskResponse
func (client *Client) GetProjectTaskWithOptions(request *GetProjectTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdempotentId)) {
		query["IdempotentId"] = request.IdempotentId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProjectTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/project/getProjectTask"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据人合成信息
//
// @param request - GetProjectTaskRequest
//
// @return GetProjectTaskResponse
func (client *Client) GetProjectTask(request *GetProjectTaskRequest) (_result *GetProjectTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectTaskResponse{}
	_body, _err := client.GetProjectTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询文案
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextResponse
func (client *Client) GetTextWithOptions(textId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTextResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetText"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/texts/" + tea.StringValue(openapiutil.GetEncodeParam(textId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询文案
//
// @return GetTextResponse
func (client *Client) GetText(textId *string) (_result *GetTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTextResponse{}
	_body, _err := client.GetTextWithOptions(textId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询文案任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextTaskResponse
func (client *Client) GetTextTaskWithOptions(textTaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTextTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTextTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/textTasks/" + tea.StringValue(openapiutil.GetEncodeParam(textTaskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTextTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询文案任务
//
// @return GetTextTaskResponse
func (client *Client) GetTextTask(textTaskId *string) (_result *GetTextTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTextTaskResponse{}
	_body, _err := client.GetTextTaskWithOptions(textTaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询表单配置
//
// @param request - GetTextTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextTemplateResponse
func (client *Client) GetTextTemplateWithOptions(request *GetTextTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTextTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Industry)) {
		query["industry"] = request.Industry
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTextTemplate"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/texts/commands/getTextTemplate"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTextTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询表单配置
//
// @param request - GetTextTemplateRequest
//
// @return GetTextTemplateResponse
func (client *Client) GetTextTemplate(request *GetTextTemplateRequest) (_result *GetTextTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTextTemplateResponse{}
	_body, _err := client.GetTextTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 营销文案互动问答
//
// @param request - InteractTextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InteractTextResponse
func (client *Client) InteractTextWithOptions(request *InteractTextRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InteractTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InteractText"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/stream/interactText"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InteractTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 营销文案互动问答
//
// @param request - InteractTextRequest
//
// @return InteractTextResponse
func (client *Client) InteractText(request *InteractTextRequest) (_result *InteractTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InteractTextResponse{}
	_body, _err := client.InteractTextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询任务列表
//
// @param request - ListAICoachTaskPageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAICoachTaskPageResponse
func (client *Client) ListAICoachTaskPageWithOptions(request *ListAICoachTaskPageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAICoachTaskPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.StudentId)) {
		query["studentId"] = request.StudentId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAICoachTaskPage"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/aicoach/listTaskPage"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAICoachTaskPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务列表
//
// @param request - ListAICoachTaskPageRequest
//
// @return ListAICoachTaskPageResponse
func (client *Client) ListAICoachTaskPage(request *ListAICoachTaskPageRequest) (_result *ListAICoachTaskPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAICoachTaskPageResponse{}
	_body, _err := client.ListAICoachTaskPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数字人模特列表
//
// @param request - ListAnchorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAnchorResponse
func (client *Client) ListAnchorWithOptions(request *ListAnchorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAnchorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnchorType)) {
		query["anchorType"] = request.AnchorType
	}

	if !tea.BoolValue(util.IsUnset(request.CoverRate)) {
		query["coverRate"] = request.CoverRate
	}

	if !tea.BoolValue(util.IsUnset(request.DigitalHumanType)) {
		query["digitalHumanType"] = request.DigitalHumanType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResSpecType)) {
		query["resSpecType"] = request.ResSpecType
	}

	if !tea.BoolValue(util.IsUnset(request.UseScene)) {
		query["useScene"] = request.UseScene
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAnchor"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/anchorOpen/listAnchor"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAnchorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数字人模特列表
//
// @param request - ListAnchorRequest
//
// @return ListAnchorResponse
func (client *Client) ListAnchor(request *ListAnchorRequest) (_result *ListAnchorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAnchorResponse{}
	_body, _err := client.ListAnchorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询数字人项目启动结果
//
// @param tmpReq - ListAvatarProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvatarProjectResponse
func (client *Client) ListAvatarProjectWithOptions(tmpReq *ListAvatarProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAvatarProjectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListAvatarProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ProjectIdList)) {
		request.ProjectIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProjectIdList, tea.String("projectIdList"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectIdListShrink)) {
		query["projectIdList"] = request.ProjectIdListShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAvatarProject"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/avatar/project/listAvatarProject"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAvatarProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询数字人项目启动结果
//
// @param request - ListAvatarProjectRequest
//
// @return ListAvatarProjectResponse
func (client *Client) ListAvatarProject(request *ListAvatarProjectRequest) (_result *ListAvatarProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAvatarProjectResponse{}
	_body, _err := client.ListAvatarProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询文案主题列表
//
// @param request - ListTextThemesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTextThemesResponse
func (client *Client) ListTextThemesWithOptions(request *ListTextThemesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTextThemesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Industry)) {
		query["industry"] = request.Industry
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTextThemes"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/textThemes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTextThemesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询文案主题列表
//
// @param request - ListTextThemesRequest
//
// @return ListTextThemesResponse
func (client *Client) ListTextThemes(request *ListTextThemesRequest) (_result *ListTextThemesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTextThemesResponse{}
	_body, _err := client.ListTextThemesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举文案
//
// @param request - ListTextsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTextsResponse
func (client *Client) ListTextsWithOptions(request *ListTextsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTextsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GenerationSource)) {
		query["generationSource"] = request.GenerationSource
	}

	if !tea.BoolValue(util.IsUnset(request.Industry)) {
		query["industry"] = request.Industry
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PublishStatus)) {
		query["publishStatus"] = request.PublishStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TextStyleType)) {
		query["textStyleType"] = request.TextStyleType
	}

	if !tea.BoolValue(util.IsUnset(request.TextTheme)) {
		query["textTheme"] = request.TextTheme
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTexts"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/texts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTextsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举文案
//
// @param request - ListTextsRequest
//
// @return ListTextsResponse
func (client *Client) ListTexts(request *ListTextsRequest) (_result *ListTextsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTextsResponse{}
	_body, _err := client.ListTextsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取声音模版列表
//
// @param request - ListVoiceModelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVoiceModelsResponse
func (client *Client) ListVoiceModelsWithOptions(request *ListVoiceModelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListVoiceModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResSpecType)) {
		query["resSpecType"] = request.ResSpecType
	}

	if !tea.BoolValue(util.IsUnset(request.UseScene)) {
		query["useScene"] = request.UseScene
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceType)) {
		query["voiceType"] = request.VoiceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVoiceModels"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/voiceOpen/listVoiceModels"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVoiceModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取声音模版列表
//
// @param request - ListVoiceModelsRequest
//
// @return ListVoiceModelsResponse
func (client *Client) ListVoiceModels(request *ListVoiceModelsRequest) (_result *ListVoiceModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListVoiceModelsResponse{}
	_body, _err := client.ListVoiceModelsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 操作实时数字人项目
//
// @param request - OperateAvatarProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateAvatarProjectResponse
func (client *Client) OperateAvatarProjectWithOptions(request *OperateAvatarProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *OperateAvatarProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		body["operateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ResChannelNumber)) {
		body["resChannelNumber"] = request.ResChannelNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		body["resType"] = request.ResType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateAvatarProject"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/avatar/project/operateProjectAvatar"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &OperateAvatarProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 操作实时数字人项目
//
// @param request - OperateAvatarProjectRequest
//
// @return OperateAvatarProjectResponse
func (client *Client) OperateAvatarProject(request *OperateAvatarProjectRequest) (_result *OperateAvatarProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OperateAvatarProjectResponse{}
	_body, _err := client.OperateAvatarProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数字人项目信息
//
// @param request - QueryAvatarProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAvatarProjectResponse
func (client *Client) QueryAvatarProjectWithOptions(request *QueryAvatarProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryAvatarProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["projectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAvatarProject"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/avatar/project/queryAvatarProject"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAvatarProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数字人项目信息
//
// @param request - QueryAvatarProjectRequest
//
// @return QueryAvatarProjectResponse
func (client *Client) QueryAvatarProject(request *QueryAvatarProjectRequest) (_result *QueryAvatarProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryAvatarProjectResponse{}
	_body, _err := client.QueryAvatarProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查找资源
//
// @param request - QueryAvatarResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAvatarResourceResponse
func (client *Client) QueryAvatarResourceWithOptions(request *QueryAvatarResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryAvatarResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdempotentId)) {
		query["idempotentId"] = request.IdempotentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAvatarResource"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/avatar/project/queryResource"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAvatarResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查找资源
//
// @param request - QueryAvatarResourceRequest
//
// @return QueryAvatarResourceResponse
func (client *Client) QueryAvatarResource(request *QueryAvatarResourceRequest) (_result *QueryAvatarResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryAvatarResourceResponse{}
	_body, _err := client.QueryAvatarResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询会话信息
//
// @param tmpReq - QuerySessionInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySessionInfoResponse
func (client *Client) QuerySessionInfoWithOptions(tmpReq *QuerySessionInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QuerySessionInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QuerySessionInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.StatusList)) {
		request.StatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StatusList, tea.String("statusList"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["pageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.StatusListShrink)) {
		query["statusList"] = request.StatusListShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySessionInfo"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/avatar/project/querySessionInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySessionInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会话信息
//
// @param request - QuerySessionInfoRequest
//
// @return QuerySessionInfoResponse
func (client *Client) QuerySessionInfo(request *QuerySessionInfoRequest) (_result *QuerySessionInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QuerySessionInfoResponse{}
	_body, _err := client.QuerySessionInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 流式输出文案
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTextStreamResponse
func (client *Client) QueryTextStreamWithOptions(textId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryTextStreamResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTextStream"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/stream/queryTextStream/" + tea.StringValue(openapiutil.GetEncodeParam(textId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTextStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 流式输出文案
//
// @return QueryTextStreamResponse
func (client *Client) QueryTextStream(textId *string) (_result *QueryTextStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryTextStreamResponse{}
	_body, _err := client.QueryTextStreamWithOptions(textId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存实时数字人项目
//
// @param request - SaveAvatarProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveAvatarProjectResponse
func (client *Client) SaveAvatarProjectWithOptions(request *SaveAvatarProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveAvatarProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.Frames)) {
		body["frames"] = request.Frames
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		body["operateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["projectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ResSpecType)) {
		body["resSpecType"] = request.ResSpecType
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleType)) {
		body["scaleType"] = request.ScaleType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveAvatarProject"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/avatar/project/saveAvatarProject"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveAvatarProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存实时数字人项目
//
// @param request - SaveAvatarProjectRequest
//
// @return SaveAvatarProjectResponse
func (client *Client) SaveAvatarProject(request *SaveAvatarProjectRequest) (_result *SaveAvatarProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveAvatarProjectResponse{}
	_body, _err := client.SaveAvatarProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询图片任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SelectImageTaskResponse
func (client *Client) SelectImageTaskWithOptions(taskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SelectImageTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("SelectImageTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/images/portrait/select/" + tea.StringValue(openapiutil.GetEncodeParam(taskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SelectImageTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询图片任务
//
// @return SelectImageTaskResponse
func (client *Client) SelectImageTask(taskId *string) (_result *SelectImageTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SelectImageTaskResponse{}
	_body, _err := client.SelectImageTaskWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询离线数字人剩余资源
//
// @param request - SelectResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SelectResourceResponse
func (client *Client) SelectResourceWithOptions(request *SelectResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SelectResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdempotentId)) {
		query["idempotentId"] = request.IdempotentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SelectResource"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/project/commands/overview"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SelectResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询离线数字人剩余资源
//
// @param request - SelectResourceRequest
//
// @return SelectResourceResponse
func (client *Client) SelectResource(request *SelectResourceRequest) (_result *SelectResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SelectResourceResponse{}
	_body, _err := client.SelectResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送文本消息
//
// @param request - SendTextMsgRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendTextMsgResponse
func (client *Client) SendTextMsgWithOptions(request *SendTextMsgRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SendTextMsgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendTextMsg"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/avatar/project/sendTextMsg"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SendTextMsgResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送文本消息
//
// @param request - SendTextMsgRequest
//
// @return SendTextMsgResponse
func (client *Client) SendTextMsg(request *SendTextMsgRequest) (_result *SendTextMsgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendTextMsgResponse{}
	_body, _err := client.SendTextMsgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动会话
//
// @param request - StartAvatarSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAvatarSessionResponse
func (client *Client) StartAvatarSessionWithOptions(request *StartAvatarSessionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartAvatarSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomPushUrl)) {
		body["customPushUrl"] = request.CustomPushUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartAvatarSession"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/avatar/project/startAvatarSession"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartAvatarSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动会话
//
// @param request - StartAvatarSessionRequest
//
// @return StartAvatarSessionResponse
func (client *Client) StartAvatarSession(request *StartAvatarSessionRequest) (_result *StartAvatarSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartAvatarSessionResponse{}
	_body, _err := client.StartAvatarSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止会话
//
// @param request - StopAvatarSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAvatarSessionResponse
func (client *Client) StopAvatarSessionWithOptions(request *StopAvatarSessionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopAvatarSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopAvatarSession"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/avatar/project/stopAvatarSession"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopAvatarSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止会话
//
// @param request - StopAvatarSessionRequest
//
// @return StopAvatarSessionResponse
func (client *Client) StopAvatarSession(request *StopAvatarSessionRequest) (_result *StopAvatarSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopAvatarSessionResponse{}
	_body, _err := client.StopAvatarSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视频合成任务停止
//
// @param request - StopProjectTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopProjectTaskResponse
func (client *Client) StopProjectTaskWithOptions(request *StopProjectTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopProjectTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopProjectTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/project/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopProjectTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频合成任务停止
//
// @param request - StopProjectTaskRequest
//
// @return StopProjectTaskResponse
func (client *Client) StopProjectTask(request *StopProjectTaskRequest) (_result *StopProjectTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopProjectTaskResponse{}
	_body, _err := client.StopProjectTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交离线数字人合成任务
//
// @param request - SubmitProjectTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitProjectTaskResponse
func (client *Client) SubmitProjectTaskWithOptions(request *SubmitProjectTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitProjectTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Frames)) {
		body["frames"] = request.Frames
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleType)) {
		body["scaleType"] = request.ScaleType
	}

	if !tea.BoolValue(util.IsUnset(request.SubtitleTag)) {
		body["subtitleTag"] = request.SubtitleTag
	}

	if !tea.BoolValue(util.IsUnset(request.TransparentBackground)) {
		body["transparentBackground"] = request.TransparentBackground
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitProjectTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/project/submitProjectTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitProjectTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交离线数字人合成任务
//
// @param request - SubmitProjectTaskRequest
//
// @return SubmitProjectTaskResponse
func (client *Client) SubmitProjectTask(request *SubmitProjectTaskRequest) (_result *SubmitProjectTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitProjectTaskResponse{}
	_body, _err := client.SubmitProjectTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人像风格变化
//
// @param request - TransferPortraitStyleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferPortraitStyleResponse
func (client *Client) TransferPortraitStyleWithOptions(request *TransferPortraitStyleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TransferPortraitStyleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Height)) {
		body["height"] = request.Height
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["imageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Numbers)) {
		body["numbers"] = request.Numbers
	}

	if !tea.BoolValue(util.IsUnset(request.RedrawAmplitude)) {
		body["redrawAmplitude"] = request.RedrawAmplitude
	}

	if !tea.BoolValue(util.IsUnset(request.Style)) {
		body["style"] = request.Style
	}

	if !tea.BoolValue(util.IsUnset(request.Width)) {
		body["width"] = request.Width
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferPortraitStyle"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/images/portrait/transferPortraitStyle"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferPortraitStyleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人像风格变化
//
// @param request - TransferPortraitStyleRequest
//
// @return TransferPortraitStyleResponse
func (client *Client) TransferPortraitStyle(request *TransferPortraitStyleRequest) (_result *TransferPortraitStyleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TransferPortraitStyleResponse{}
	_body, _err := client.TransferPortraitStyleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
