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
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// doc_test_3
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// example:
	//
	// 2C331582-7390-5949-8D9A-AC8239185B37
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

func (s *QueryAvatarProjectResponseBody) SetProjectName(v string) *QueryAvatarProjectResponseBody {
	s.ProjectName = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetRequestId(v string) *QueryAvatarProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetStatus(v string) *QueryAvatarProjectResponseBody {
	s.Status = &v
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
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
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
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// https://meeting.dingtalk.com/j/1COFppy0POR
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
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

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetUrl(v string) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.Url = &v
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
	AudioUrl *string `json:"audioUrl,omitempty" xml:"audioUrl,omitempty"`
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
