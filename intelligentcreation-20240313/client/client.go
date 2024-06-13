// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	Texts          *Text   `json:"texts,omitempty" xml:"texts,omitempty"`
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

func (s *TextTask) SetTexts(v *Text) *TextTask {
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
