// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddDocumentInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// pdf
	DocumentType *string `json:"documentType,omitempty" xml:"documentType,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s AddDocumentInfo) String() string {
	return tea.Prettify(s)
}

func (s AddDocumentInfo) GoString() string {
	return s.String()
}

func (s *AddDocumentInfo) SetDocumentType(v string) *AddDocumentInfo {
	s.DocumentType = &v
	return s
}

func (s *AddDocumentInfo) SetName(v string) *AddDocumentInfo {
	s.Name = &v
	return s
}

func (s *AddDocumentInfo) SetUrl(v string) *AddDocumentInfo {
	s.Url = &v
	return s
}

type AddDocumentResult struct {
	// example:
	//
	// example.pdf
	DocName      *string       `json:"docName,omitempty" xml:"docName,omitempty"`
	DocumentInfo *DocumentInfo `json:"documentInfo,omitempty" xml:"documentInfo,omitempty"`
	// example:
	//
	// true
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddDocumentResult) String() string {
	return tea.Prettify(s)
}

func (s AddDocumentResult) GoString() string {
	return s.String()
}

func (s *AddDocumentResult) SetDocName(v string) *AddDocumentResult {
	s.DocName = &v
	return s
}

func (s *AddDocumentResult) SetDocumentInfo(v *DocumentInfo) *AddDocumentResult {
	s.DocumentInfo = v
	return s
}

func (s *AddDocumentResult) SetErrorMessage(v string) *AddDocumentResult {
	s.ErrorMessage = &v
	return s
}

func (s *AddDocumentResult) SetSuccess(v bool) *AddDocumentResult {
	s.Success = &v
	return s
}

type AnchorResponse struct {
	AnchorCategory     *string `json:"anchorCategory,omitempty" xml:"anchorCategory,omitempty"`
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

func (s *AnchorResponse) SetAnchorCategory(v string) *AnchorResponse {
	s.AnchorCategory = &v
	return s
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

type BatchAddDocumentResult struct {
	// This parameter is required.
	AddDocumentResults []*AddDocumentResult `json:"addDocumentResults,omitempty" xml:"addDocumentResults,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchAddDocumentResult) String() string {
	return tea.Prettify(s)
}

func (s BatchAddDocumentResult) GoString() string {
	return s.String()
}

func (s *BatchAddDocumentResult) SetAddDocumentResults(v []*AddDocumentResult) *BatchAddDocumentResult {
	s.AddDocumentResults = v
	return s
}

func (s *BatchAddDocumentResult) SetRequestId(v string) *BatchAddDocumentResult {
	s.RequestId = &v
	return s
}

type DocumentInfo struct {
	DocumentType  *string `json:"documentType,omitempty" xml:"documentType,omitempty"`
	Id            *string `json:"id,omitempty" xml:"id,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	ProcessStatus *string `json:"processStatus,omitempty" xml:"processStatus,omitempty"`
}

func (s DocumentInfo) String() string {
	return tea.Prettify(s)
}

func (s DocumentInfo) GoString() string {
	return s.String()
}

func (s *DocumentInfo) SetDocumentType(v string) *DocumentInfo {
	s.DocumentType = &v
	return s
}

func (s *DocumentInfo) SetId(v string) *DocumentInfo {
	s.Id = &v
	return s
}

func (s *DocumentInfo) SetName(v string) *DocumentInfo {
	s.Name = &v
	return s
}

func (s *DocumentInfo) SetProcessStatus(v string) *DocumentInfo {
	s.ProcessStatus = &v
	return s
}

type DocumentResult struct {
	DocumentInfo *DocumentInfo `json:"documentInfo,omitempty" xml:"documentInfo,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DocumentResult) String() string {
	return tea.Prettify(s)
}

func (s DocumentResult) GoString() string {
	return s.String()
}

func (s *DocumentResult) SetDocumentInfo(v *DocumentInfo) *DocumentResult {
	s.DocumentInfo = v
	return s
}

func (s *DocumentResult) SetRequestId(v string) *DocumentResult {
	s.RequestId = &v
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
	IdempotentId *string   `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	ImageUrls    []*string `json:"imageUrls,omitempty" xml:"imageUrls,omitempty" type:"Repeated"`
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

func (s *IllustrationTaskCreateCmd) SetImageUrls(v []*string) *IllustrationTaskCreateCmd {
	s.ImageUrls = v
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

type KnowledgeBaseInfo struct {
	Description       *string `json:"description,omitempty" xml:"description,omitempty"`
	Id                *string `json:"id,omitempty" xml:"id,omitempty"`
	Industry          *string `json:"industry,omitempty" xml:"industry,omitempty"`
	KnowledgeBaseType *string `json:"knowledgeBaseType,omitempty" xml:"knowledgeBaseType,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s KnowledgeBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s KnowledgeBaseInfo) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseInfo) SetDescription(v string) *KnowledgeBaseInfo {
	s.Description = &v
	return s
}

func (s *KnowledgeBaseInfo) SetId(v string) *KnowledgeBaseInfo {
	s.Id = &v
	return s
}

func (s *KnowledgeBaseInfo) SetIndustry(v string) *KnowledgeBaseInfo {
	s.Industry = &v
	return s
}

func (s *KnowledgeBaseInfo) SetKnowledgeBaseType(v string) *KnowledgeBaseInfo {
	s.KnowledgeBaseType = &v
	return s
}

func (s *KnowledgeBaseInfo) SetName(v string) *KnowledgeBaseInfo {
	s.Name = &v
	return s
}

type KnowledgeBaseListResult struct {
	KnowledgeBases []*KnowledgeBaseInfo `json:"knowledgeBases,omitempty" xml:"knowledgeBases,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s KnowledgeBaseListResult) String() string {
	return tea.Prettify(s)
}

func (s KnowledgeBaseListResult) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseListResult) SetKnowledgeBases(v []*KnowledgeBaseInfo) *KnowledgeBaseListResult {
	s.KnowledgeBases = v
	return s
}

func (s *KnowledgeBaseListResult) SetRequestId(v string) *KnowledgeBaseListResult {
	s.RequestId = &v
	return s
}

func (s *KnowledgeBaseListResult) SetTotal(v int32) *KnowledgeBaseListResult {
	s.Total = &v
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
	Point         *string       `json:"point,omitempty" xml:"point,omitempty"`
	ReferenceTag  *ReferenceTag `json:"referenceTag,omitempty" xml:"referenceTag,omitempty"`
	RelatedRagIds []*int64      `json:"relatedRagIds,omitempty" xml:"relatedRagIds,omitempty" type:"Repeated"`
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

func (s *TextTask) SetRelatedRagIds(v []*int64) *TextTask {
	s.RelatedRagIds = v
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

type BatchAddDocumentRequest struct {
	AddDocumentInfos []*AddDocumentInfo `json:"addDocumentInfos,omitempty" xml:"addDocumentInfos,omitempty" type:"Repeated"`
}

func (s BatchAddDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddDocumentRequest) GoString() string {
	return s.String()
}

func (s *BatchAddDocumentRequest) SetAddDocumentInfos(v []*AddDocumentInfo) *BatchAddDocumentRequest {
	s.AddDocumentInfos = v
	return s
}

type BatchAddDocumentResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddDocumentResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAddDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAddDocumentResponse) GoString() string {
	return s.String()
}

func (s *BatchAddDocumentResponse) SetHeaders(v map[string]*string) *BatchAddDocumentResponse {
	s.Headers = v
	return s
}

func (s *BatchAddDocumentResponse) SetStatusCode(v int32) *BatchAddDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddDocumentResponse) SetBody(v *BatchAddDocumentResult) *BatchAddDocumentResponse {
	s.Body = v
	return s
}

type BatchCreateAICoachTaskRequest struct {
	// example:
	//
	// 7915125A-0D96-5A25-A54B-D3B739A86AFC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1
	ScriptRecordId *string                                     `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
	StudentIds     []*string                                   `json:"studentIds,omitempty" xml:"studentIds,omitempty" type:"Repeated"`
	StudentList    []*BatchCreateAICoachTaskRequestStudentList `json:"studentList,omitempty" xml:"studentList,omitempty" type:"Repeated"`
}

func (s BatchCreateAICoachTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateAICoachTaskRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateAICoachTaskRequest) SetRequestId(v string) *BatchCreateAICoachTaskRequest {
	s.RequestId = &v
	return s
}

func (s *BatchCreateAICoachTaskRequest) SetScriptRecordId(v string) *BatchCreateAICoachTaskRequest {
	s.ScriptRecordId = &v
	return s
}

func (s *BatchCreateAICoachTaskRequest) SetStudentIds(v []*string) *BatchCreateAICoachTaskRequest {
	s.StudentIds = v
	return s
}

func (s *BatchCreateAICoachTaskRequest) SetStudentList(v []*BatchCreateAICoachTaskRequestStudentList) *BatchCreateAICoachTaskRequest {
	s.StudentList = v
	return s
}

type BatchCreateAICoachTaskRequestStudentList struct {
	StudentAudioUrl *string `json:"studentAudioUrl,omitempty" xml:"studentAudioUrl,omitempty"`
	StudentId       *string `json:"studentId,omitempty" xml:"studentId,omitempty"`
}

func (s BatchCreateAICoachTaskRequestStudentList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateAICoachTaskRequestStudentList) GoString() string {
	return s.String()
}

func (s *BatchCreateAICoachTaskRequestStudentList) SetStudentAudioUrl(v string) *BatchCreateAICoachTaskRequestStudentList {
	s.StudentAudioUrl = &v
	return s
}

func (s *BatchCreateAICoachTaskRequestStudentList) SetStudentId(v string) *BatchCreateAICoachTaskRequestStudentList {
	s.StudentId = &v
	return s
}

type BatchCreateAICoachTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 10923AA3-F7A1-5EA0-ACCA-D704269EAA78
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TaskIds   []*string `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
}

func (s BatchCreateAICoachTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateAICoachTaskResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateAICoachTaskResponseBody) SetRequestId(v string) *BatchCreateAICoachTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateAICoachTaskResponseBody) SetTaskIds(v []*string) *BatchCreateAICoachTaskResponseBody {
	s.TaskIds = v
	return s
}

type BatchCreateAICoachTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateAICoachTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateAICoachTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateAICoachTaskResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateAICoachTaskResponse) SetHeaders(v map[string]*string) *BatchCreateAICoachTaskResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateAICoachTaskResponse) SetStatusCode(v int32) *BatchCreateAICoachTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateAICoachTaskResponse) SetBody(v *BatchCreateAICoachTaskResponseBody) *BatchCreateAICoachTaskResponse {
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

type BatchGetTrainTaskRequest struct {
	// example:
	//
	// 1524004782431111
	AliyunMainId *string   `json:"aliyunMainId,omitempty" xml:"aliyunMainId,omitempty"`
	TaskIdList   []*string `json:"taskIdList,omitempty" xml:"taskIdList,omitempty" type:"Repeated"`
}

func (s BatchGetTrainTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetTrainTaskRequest) GoString() string {
	return s.String()
}

func (s *BatchGetTrainTaskRequest) SetAliyunMainId(v string) *BatchGetTrainTaskRequest {
	s.AliyunMainId = &v
	return s
}

func (s *BatchGetTrainTaskRequest) SetTaskIdList(v []*string) *BatchGetTrainTaskRequest {
	s.TaskIdList = v
	return s
}

type BatchGetTrainTaskShrinkRequest struct {
	// example:
	//
	// 1524004782431111
	AliyunMainId     *string `json:"aliyunMainId,omitempty" xml:"aliyunMainId,omitempty"`
	TaskIdListShrink *string `json:"taskIdList,omitempty" xml:"taskIdList,omitempty"`
}

func (s BatchGetTrainTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetTrainTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetTrainTaskShrinkRequest) SetAliyunMainId(v string) *BatchGetTrainTaskShrinkRequest {
	s.AliyunMainId = &v
	return s
}

func (s *BatchGetTrainTaskShrinkRequest) SetTaskIdListShrink(v string) *BatchGetTrainTaskShrinkRequest {
	s.TaskIdListShrink = &v
	return s
}

type BatchGetTrainTaskResponseBody struct {
	// example:
	//
	// 2226A26A-26E5-5AB9-A14A-54D612FCF96A
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	VoiceList []*BatchGetTrainTaskResponseBodyVoiceList `json:"voiceList,omitempty" xml:"voiceList,omitempty" type:"Repeated"`
}

func (s BatchGetTrainTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetTrainTaskResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetTrainTaskResponseBody) SetRequestId(v string) *BatchGetTrainTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetTrainTaskResponseBody) SetVoiceList(v []*BatchGetTrainTaskResponseBodyVoiceList) *BatchGetTrainTaskResponseBody {
	s.VoiceList = v
	return s
}

type BatchGetTrainTaskResponseBodyVoiceList struct {
	// example:
	//
	// 1524004782438111
	AliyunSubId      *string `json:"aliyunSubId,omitempty" xml:"aliyunSubId,omitempty"`
	AuditFailMessage *string `json:"auditFailMessage,omitempty" xml:"auditFailMessage,omitempty"`
	// example:
	//
	// auditFail
	AuditStatus *string `json:"auditStatus,omitempty" xml:"auditStatus,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// M
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// BASIC_MODEL
	ResSpecType *string `json:"resSpecType,omitempty" xml:"resSpecType,omitempty"`
	// example:
	//
	// 837091359375048704
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// CopyVoice
	TaskType         *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	TrainFailMessage *string `json:"trainFailMessage,omitempty" xml:"trainFailMessage,omitempty"`
	// example:
	//
	// trainFail
	TrainStatus *string `json:"trainStatus,omitempty" xml:"trainStatus,omitempty"`
	// example:
	//
	// realTimeInteractivity
	UseScene      *string                                              `json:"useScene,omitempty" xml:"useScene,omitempty"`
	VoiceMaterial *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial `json:"voiceMaterial,omitempty" xml:"voiceMaterial,omitempty" type:"Struct"`
}

func (s BatchGetTrainTaskResponseBodyVoiceList) String() string {
	return tea.Prettify(s)
}

func (s BatchGetTrainTaskResponseBodyVoiceList) GoString() string {
	return s.String()
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetAliyunSubId(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.AliyunSubId = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetAuditFailMessage(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.AuditFailMessage = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetAuditStatus(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.AuditStatus = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetCreateTime(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.CreateTime = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetGender(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.Gender = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetName(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.Name = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetResSpecType(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.ResSpecType = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetTaskId(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.TaskId = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetTaskType(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.TaskType = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetTrainFailMessage(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.TrainFailMessage = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetTrainStatus(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.TrainStatus = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetUseScene(v string) *BatchGetTrainTaskResponseBodyVoiceList {
	s.UseScene = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceList) SetVoiceMaterial(v *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial) *BatchGetTrainTaskResponseBodyVoiceList {
	s.VoiceMaterial = v
	return s
}

type BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial struct {
	// example:
	//
	// 1
	VoiceId *int64 `json:"voiceId,omitempty" xml:"voiceId,omitempty"`
	// example:
	//
	// zh
	VoiceLanguage *string `json:"voiceLanguage,omitempty" xml:"voiceLanguage,omitempty"`
	// example:
	//
	// http://www.voice.com
	VoiceUrl *string `json:"voiceUrl,omitempty" xml:"voiceUrl,omitempty"`
}

func (s BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial) String() string {
	return tea.Prettify(s)
}

func (s BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial) GoString() string {
	return s.String()
}

func (s *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial) SetVoiceId(v int64) *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial {
	s.VoiceId = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial) SetVoiceLanguage(v string) *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial {
	s.VoiceLanguage = &v
	return s
}

func (s *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial) SetVoiceUrl(v string) *BatchGetTrainTaskResponseBodyVoiceListVoiceMaterial {
	s.VoiceUrl = &v
	return s
}

type BatchGetTrainTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetTrainTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetTrainTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetTrainTaskResponse) GoString() string {
	return s.String()
}

func (s *BatchGetTrainTaskResponse) SetHeaders(v map[string]*string) *BatchGetTrainTaskResponse {
	s.Headers = v
	return s
}

func (s *BatchGetTrainTaskResponse) SetStatusCode(v int32) *BatchGetTrainTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetTrainTaskResponse) SetBody(v *BatchGetTrainTaskResponseBody) *BatchGetTrainTaskResponse {
	s.Body = v
	return s
}

type BatchGetVideoClipTaskRequest struct {
	TaskIdList []*string `json:"taskIdList,omitempty" xml:"taskIdList,omitempty" type:"Repeated"`
}

func (s BatchGetVideoClipTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetVideoClipTaskRequest) GoString() string {
	return s.String()
}

func (s *BatchGetVideoClipTaskRequest) SetTaskIdList(v []*string) *BatchGetVideoClipTaskRequest {
	s.TaskIdList = v
	return s
}

type BatchGetVideoClipTaskShrinkRequest struct {
	TaskIdListShrink *string `json:"taskIdList,omitempty" xml:"taskIdList,omitempty"`
}

func (s BatchGetVideoClipTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetVideoClipTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetVideoClipTaskShrinkRequest) SetTaskIdListShrink(v string) *BatchGetVideoClipTaskShrinkRequest {
	s.TaskIdListShrink = &v
	return s
}

type BatchGetVideoClipTaskResponseBody struct {
	// example:
	//
	// 86A90C40-D1AB-50DA-A4B1-0D545F80F2FE
	RequestId *string                                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TaskList  []*BatchGetVideoClipTaskResponseBodyTaskList `json:"taskList,omitempty" xml:"taskList,omitempty" type:"Repeated"`
}

func (s BatchGetVideoClipTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetVideoClipTaskResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetVideoClipTaskResponseBody) SetRequestId(v string) *BatchGetVideoClipTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBody) SetTaskList(v []*BatchGetVideoClipTaskResponseBodyTaskList) *BatchGetVideoClipTaskResponseBody {
	s.TaskList = v
	return s
}

type BatchGetVideoClipTaskResponseBodyTaskList struct {
	// example:
	//
	// FINISHED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 864413342857035776
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 43335
	TotalDuration *float64 `json:"totalDuration,omitempty" xml:"totalDuration,omitempty"`
	// example:
	//
	// 11
	TotalToken *int64                                                `json:"totalToken,omitempty" xml:"totalToken,omitempty"`
	VideoList  []*BatchGetVideoClipTaskResponseBodyTaskListVideoList `json:"videoList,omitempty" xml:"videoList,omitempty" type:"Repeated"`
}

func (s BatchGetVideoClipTaskResponseBodyTaskList) String() string {
	return tea.Prettify(s)
}

func (s BatchGetVideoClipTaskResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *BatchGetVideoClipTaskResponseBodyTaskList) SetStatus(v string) *BatchGetVideoClipTaskResponseBodyTaskList {
	s.Status = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskList) SetTaskId(v string) *BatchGetVideoClipTaskResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskList) SetTotalDuration(v float64) *BatchGetVideoClipTaskResponseBodyTaskList {
	s.TotalDuration = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskList) SetTotalToken(v int64) *BatchGetVideoClipTaskResponseBodyTaskList {
	s.TotalToken = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskList) SetVideoList(v []*BatchGetVideoClipTaskResponseBodyTaskListVideoList) *BatchGetVideoClipTaskResponseBodyTaskList {
	s.VideoList = v
	return s
}

type BatchGetVideoClipTaskResponseBodyTaskListVideoList struct {
	// example:
	//
	// 0
	BeginTime   *int32  `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 11110
	EndTime  *int32  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// https://e-ai.oss-cn-guangzhou.aliyuncs.com/video/jlkasdl.mp4
	VideoDownloadUrl *string `json:"videoDownloadUrl,omitempty" xml:"videoDownloadUrl,omitempty"`
	VideoName        *string `json:"videoName,omitempty" xml:"videoName,omitempty"`
	// example:
	//
	// https://e-ai.oss-cn-guangzhou.aliyuncs.com/video/jlkasdl.mp4
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s BatchGetVideoClipTaskResponseBodyTaskListVideoList) String() string {
	return tea.Prettify(s)
}

func (s BatchGetVideoClipTaskResponseBodyTaskListVideoList) GoString() string {
	return s.String()
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) SetBeginTime(v int32) *BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	s.BeginTime = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) SetDescription(v string) *BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	s.Description = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) SetEndTime(v int32) *BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	s.EndTime = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) SetErrorMsg(v string) *BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	s.ErrorMsg = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) SetTitle(v string) *BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	s.Title = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) SetVideoDownloadUrl(v string) *BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	s.VideoDownloadUrl = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) SetVideoName(v string) *BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	s.VideoName = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) SetVideoUrl(v string) *BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	s.VideoUrl = &v
	return s
}

type BatchGetVideoClipTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetVideoClipTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetVideoClipTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetVideoClipTaskResponse) GoString() string {
	return s.String()
}

func (s *BatchGetVideoClipTaskResponse) SetHeaders(v map[string]*string) *BatchGetVideoClipTaskResponse {
	s.Headers = v
	return s
}

func (s *BatchGetVideoClipTaskResponse) SetStatusCode(v int32) *BatchGetVideoClipTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetVideoClipTaskResponse) SetBody(v *BatchGetVideoClipTaskResponseBody) *BatchGetVideoClipTaskResponse {
	s.Body = v
	return s
}

type BatchQueryIndividuationTextRequest struct {
	TextIdList []*string `json:"textIdList,omitempty" xml:"textIdList,omitempty" type:"Repeated"`
}

func (s BatchQueryIndividuationTextRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryIndividuationTextRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryIndividuationTextRequest) SetTextIdList(v []*string) *BatchQueryIndividuationTextRequest {
	s.TextIdList = v
	return s
}

type BatchQueryIndividuationTextShrinkRequest struct {
	TextIdListShrink *string `json:"textIdList,omitempty" xml:"textIdList,omitempty"`
}

func (s BatchQueryIndividuationTextShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryIndividuationTextShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryIndividuationTextShrinkRequest) SetTextIdListShrink(v string) *BatchQueryIndividuationTextShrinkRequest {
	s.TextIdListShrink = &v
	return s
}

type BatchQueryIndividuationTextResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 14878724-A835-578D-9DD5-4779ADCE9221
	RequestId *string                                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TextList  []*BatchQueryIndividuationTextResponseBodyTextList `json:"textList,omitempty" xml:"textList,omitempty" type:"Repeated"`
}

func (s BatchQueryIndividuationTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryIndividuationTextResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryIndividuationTextResponseBody) SetRequestId(v string) *BatchQueryIndividuationTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBody) SetTextList(v []*BatchQueryIndividuationTextResponseBodyTextList) *BatchQueryIndividuationTextResponseBody {
	s.TextList = v
	return s
}

type BatchQueryIndividuationTextResponseBodyTextList struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 1
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 2849286
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// example:
	//
	// 812884915104530432
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 837074737851613184
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 110825
	TextId *string `json:"textId,omitempty" xml:"textId,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// 11
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchQueryIndividuationTextResponseBodyTextList) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryIndividuationTextResponseBodyTextList) GoString() string {
	return s.String()
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetContent(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.Content = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetCreateTime(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.CreateTime = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetErrorMsg(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.ErrorMsg = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetItemId(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.ItemId = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetProjectId(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.ProjectId = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetStatus(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.Status = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetTaskId(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.TaskId = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetTextId(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.TextId = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetUpdateTime(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.UpdateTime = &v
	return s
}

func (s *BatchQueryIndividuationTextResponseBodyTextList) SetUserId(v string) *BatchQueryIndividuationTextResponseBodyTextList {
	s.UserId = &v
	return s
}

type BatchQueryIndividuationTextResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryIndividuationTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryIndividuationTextResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryIndividuationTextResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryIndividuationTextResponse) SetHeaders(v map[string]*string) *BatchQueryIndividuationTextResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryIndividuationTextResponse) SetStatusCode(v int32) *BatchQueryIndividuationTextResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryIndividuationTextResponse) SetBody(v *BatchQueryIndividuationTextResponseBody) *BatchQueryIndividuationTextResponse {
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

type CreateAICoachTaskRequest struct {
	// example:
	//
	// 541E7123-2E8A-5BA2-AC38-665650C84129
	RequestId       *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ScriptRecordId  *string `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
	StudentAudioUrl *string `json:"studentAudioUrl,omitempty" xml:"studentAudioUrl,omitempty"`
	StudentId       *string `json:"studentId,omitempty" xml:"studentId,omitempty"`
}

func (s CreateAICoachTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAICoachTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskRequest) SetRequestId(v string) *CreateAICoachTaskRequest {
	s.RequestId = &v
	return s
}

func (s *CreateAICoachTaskRequest) SetScriptRecordId(v string) *CreateAICoachTaskRequest {
	s.ScriptRecordId = &v
	return s
}

func (s *CreateAICoachTaskRequest) SetStudentAudioUrl(v string) *CreateAICoachTaskRequest {
	s.StudentAudioUrl = &v
	return s
}

func (s *CreateAICoachTaskRequest) SetStudentId(v string) *CreateAICoachTaskRequest {
	s.StudentId = &v
	return s
}

type CreateAICoachTaskResponseBody struct {
	// example:
	//
	// Deduct.DeductTaskAlreadySuccess
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0E8B1746-AE35-5C4B-A3A8-345B274AE32C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 821882330423951360
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateAICoachTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAICoachTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskResponseBody) SetErrorCode(v string) *CreateAICoachTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAICoachTaskResponseBody) SetErrorMessage(v string) *CreateAICoachTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateAICoachTaskResponseBody) SetRequestId(v string) *CreateAICoachTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAICoachTaskResponseBody) SetSuccess(v bool) *CreateAICoachTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAICoachTaskResponseBody) SetTaskId(v string) *CreateAICoachTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateAICoachTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAICoachTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAICoachTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAICoachTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskResponse) SetHeaders(v map[string]*string) *CreateAICoachTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAICoachTaskResponse) SetStatusCode(v int32) *CreateAICoachTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAICoachTaskResponse) SetBody(v *CreateAICoachTaskResponseBody) *CreateAICoachTaskResponse {
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
	SessionId     *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	SessionStatus *int64  `json:"sessionStatus,omitempty" xml:"sessionStatus,omitempty"`
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

func (s *CreateAICoachTaskSessionResponseBody) SetSessionStatus(v int64) *CreateAICoachTaskSessionResponseBody {
	s.SessionStatus = &v
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
	AgentIconUrl     *string   `json:"agentIconUrl,omitempty" xml:"agentIconUrl,omitempty"`
	CharacterName    *string   `json:"characterName,omitempty" xml:"characterName,omitempty"`
	DialogueTextFlag *bool     `json:"dialogueTextFlag,omitempty" xml:"dialogueTextFlag,omitempty"`
	DialogueTipFlag  *bool     `json:"dialogueTipFlag,omitempty" xml:"dialogueTipFlag,omitempty"`
	Initiator        *string   `json:"initiator,omitempty" xml:"initiator,omitempty"`
	InputTypeList    []*string `json:"inputTypeList,omitempty" xml:"inputTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 11
	MaxDuration *int64 `json:"maxDuration,omitempty" xml:"maxDuration,omitempty"`
	// example:
	//
	// test
	ScriptDesc            *string `json:"scriptDesc,omitempty" xml:"scriptDesc,omitempty"`
	ScriptName            *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	ScriptRecordId        *string `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
	ScriptType            *int64  `json:"scriptType,omitempty" xml:"scriptType,omitempty"`
	SparringTipContent    *string `json:"sparringTipContent,omitempty" xml:"sparringTipContent,omitempty"`
	SparringTipTitle      *string `json:"sparringTipTitle,omitempty" xml:"sparringTipTitle,omitempty"`
	StudentThinkTimeFlag  *bool   `json:"studentThinkTimeFlag,omitempty" xml:"studentThinkTimeFlag,omitempty"`
	StudentThinkTimeLimit *int64  `json:"studentThinkTimeLimit,omitempty" xml:"studentThinkTimeLimit,omitempty"`
}

func (s CreateAICoachTaskSessionResponseBodyScriptInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAICoachTaskSessionResponseBodyScriptInfo) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetAgentIconUrl(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.AgentIconUrl = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetCharacterName(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.CharacterName = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetDialogueTextFlag(v bool) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.DialogueTextFlag = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetDialogueTipFlag(v bool) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.DialogueTipFlag = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetInitiator(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.Initiator = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetInputTypeList(v []*string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.InputTypeList = v
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

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetScriptRecordId(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.ScriptRecordId = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetScriptType(v int64) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.ScriptType = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetSparringTipContent(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.SparringTipContent = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetSparringTipTitle(v string) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.SparringTipTitle = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetStudentThinkTimeFlag(v bool) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.StudentThinkTimeFlag = &v
	return s
}

func (s *CreateAICoachTaskSessionResponseBodyScriptInfo) SetStudentThinkTimeLimit(v int64) *CreateAICoachTaskSessionResponseBodyScriptInfo {
	s.StudentThinkTimeLimit = &v
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

type CreateAnchorRequest struct {
	AnchorCategory     *string `json:"anchorCategory,omitempty" xml:"anchorCategory,omitempty"`
	AnchorMaterialName *string `json:"anchorMaterialName,omitempty" xml:"anchorMaterialName,omitempty"`
	// example:
	//
	// https://yic-pre.oss-cn-hangzhou.aliyuncs.com/common/image/anchor/1733474220549-1733474198960image.png?Expires=3311144948&OSSAccessKeyId=LTAI5tPHLyFPhh4UoRias4Zg&Signature=qldDufvRDj9IUTmOtb9r2451RIU%3D
	CoverUrl         *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	DigitalHumanType *string `json:"digitalHumanType,omitempty" xml:"digitalHumanType,omitempty"`
	// example:
	//
	// F
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// offlineSynthesis
	UseScene    *string `json:"useScene,omitempty" xml:"useScene,omitempty"`
	VideoOssKey *string `json:"videoOssKey,omitempty" xml:"videoOssKey,omitempty"`
}

func (s CreateAnchorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAnchorRequest) GoString() string {
	return s.String()
}

func (s *CreateAnchorRequest) SetAnchorCategory(v string) *CreateAnchorRequest {
	s.AnchorCategory = &v
	return s
}

func (s *CreateAnchorRequest) SetAnchorMaterialName(v string) *CreateAnchorRequest {
	s.AnchorMaterialName = &v
	return s
}

func (s *CreateAnchorRequest) SetCoverUrl(v string) *CreateAnchorRequest {
	s.CoverUrl = &v
	return s
}

func (s *CreateAnchorRequest) SetDigitalHumanType(v string) *CreateAnchorRequest {
	s.DigitalHumanType = &v
	return s
}

func (s *CreateAnchorRequest) SetGender(v string) *CreateAnchorRequest {
	s.Gender = &v
	return s
}

func (s *CreateAnchorRequest) SetUseScene(v string) *CreateAnchorRequest {
	s.UseScene = &v
	return s
}

func (s *CreateAnchorRequest) SetVideoOssKey(v string) *CreateAnchorRequest {
	s.VideoOssKey = &v
	return s
}

type CreateAnchorResponseBody struct {
	// 123456789
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// PARAM_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Failed to proxy flink ui request, message: An error occurred: Invalid UUID string: jobsn
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 10923AA3-F7A1-5EA0-ACCA-D704269EAA78
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateAnchorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAnchorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAnchorResponseBody) SetData(v string) *CreateAnchorResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAnchorResponseBody) SetErrorCode(v string) *CreateAnchorResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAnchorResponseBody) SetErrorMessage(v string) *CreateAnchorResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateAnchorResponseBody) SetRequestId(v string) *CreateAnchorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAnchorResponseBody) SetSuccess(v bool) *CreateAnchorResponseBody {
	s.Success = &v
	return s
}

type CreateAnchorResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAnchorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAnchorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAnchorResponse) GoString() string {
	return s.String()
}

func (s *CreateAnchorResponse) SetHeaders(v map[string]*string) *CreateAnchorResponse {
	s.Headers = v
	return s
}

func (s *CreateAnchorResponse) SetStatusCode(v int32) *CreateAnchorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAnchorResponse) SetBody(v *CreateAnchorResponseBody) *CreateAnchorResponse {
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

type CreateIndividuationProjectRequest struct {
	ProjectInfo *string `json:"projectInfo,omitempty" xml:"projectInfo,omitempty"`
	// example:
	//
	// avatar-1
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	Purpose     *string `json:"purpose,omitempty" xml:"purpose,omitempty"`
	// example:
	//
	// ail003
	SceneId *string `json:"sceneId,omitempty" xml:"sceneId,omitempty"`
}

func (s CreateIndividuationProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIndividuationProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateIndividuationProjectRequest) SetProjectInfo(v string) *CreateIndividuationProjectRequest {
	s.ProjectInfo = &v
	return s
}

func (s *CreateIndividuationProjectRequest) SetProjectName(v string) *CreateIndividuationProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateIndividuationProjectRequest) SetPurpose(v string) *CreateIndividuationProjectRequest {
	s.Purpose = &v
	return s
}

func (s *CreateIndividuationProjectRequest) SetSceneId(v string) *CreateIndividuationProjectRequest {
	s.SceneId = &v
	return s
}

type CreateIndividuationProjectResponseBody struct {
	// example:
	//
	// 812907463682949120
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4D902811-B75C-5D1B-8882-D515F8E2F977
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateIndividuationProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIndividuationProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIndividuationProjectResponseBody) SetProjectId(v string) *CreateIndividuationProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateIndividuationProjectResponseBody) SetRequestId(v string) *CreateIndividuationProjectResponseBody {
	s.RequestId = &v
	return s
}

type CreateIndividuationProjectResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIndividuationProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIndividuationProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIndividuationProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateIndividuationProjectResponse) SetHeaders(v map[string]*string) *CreateIndividuationProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateIndividuationProjectResponse) SetStatusCode(v int32) *CreateIndividuationProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIndividuationProjectResponse) SetBody(v *CreateIndividuationProjectResponseBody) *CreateIndividuationProjectResponse {
	s.Body = v
	return s
}

type CreateIndividuationTextTaskRequest struct {
	CrowdPack [][]*string `json:"crowdPack,omitempty" xml:"crowdPack,omitempty" type:"Repeated"`
	// example:
	//
	// 840015278620459008
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	TaskName  *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s CreateIndividuationTextTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIndividuationTextTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateIndividuationTextTaskRequest) SetCrowdPack(v [][]*string) *CreateIndividuationTextTaskRequest {
	s.CrowdPack = v
	return s
}

func (s *CreateIndividuationTextTaskRequest) SetProjectId(v string) *CreateIndividuationTextTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateIndividuationTextTaskRequest) SetTaskName(v string) *CreateIndividuationTextTaskRequest {
	s.TaskName = &v
	return s
}

type CreateIndividuationTextTaskResponseBody struct {
	// example:
	//
	// 56AC346B-AF40-5E4F-AFFE-FD8BA5E6FB3A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 837091359375048704
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateIndividuationTextTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIndividuationTextTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIndividuationTextTaskResponseBody) SetRequestId(v string) *CreateIndividuationTextTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIndividuationTextTaskResponseBody) SetTaskId(v string) *CreateIndividuationTextTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateIndividuationTextTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIndividuationTextTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIndividuationTextTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIndividuationTextTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateIndividuationTextTaskResponse) SetHeaders(v map[string]*string) *CreateIndividuationTextTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateIndividuationTextTaskResponse) SetStatusCode(v int32) *CreateIndividuationTextTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIndividuationTextTaskResponse) SetBody(v *CreateIndividuationTextTaskResponseBody) *CreateIndividuationTextTaskResponse {
	s.Body = v
	return s
}

type CreateProductImageRequest struct {
	BackgroundDescription *string `json:"backgroundDescription,omitempty" xml:"backgroundDescription,omitempty"`
	// example:
	//
	// 75
	BackgroundPriority *int32 `json:"backgroundPriority,omitempty" xml:"backgroundPriority,omitempty"`
	// example:
	//
	// http://xxx/background.png
	BackgroundUrl *string `json:"backgroundUrl,omitempty" xml:"backgroundUrl,omitempty"`
	HighlightText *string `json:"highlightText,omitempty" xml:"highlightText,omitempty"`
	// example:
	//
	// 1
	ImageCount *int32 `json:"imageCount,omitempty" xml:"imageCount,omitempty"`
	// example:
	//
	// http://xxx/image.png
	ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	SubTitle *string `json:"subTitle,omitempty" xml:"subTitle,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateProductImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProductImageRequest) GoString() string {
	return s.String()
}

func (s *CreateProductImageRequest) SetBackgroundDescription(v string) *CreateProductImageRequest {
	s.BackgroundDescription = &v
	return s
}

func (s *CreateProductImageRequest) SetBackgroundPriority(v int32) *CreateProductImageRequest {
	s.BackgroundPriority = &v
	return s
}

func (s *CreateProductImageRequest) SetBackgroundUrl(v string) *CreateProductImageRequest {
	s.BackgroundUrl = &v
	return s
}

func (s *CreateProductImageRequest) SetHighlightText(v string) *CreateProductImageRequest {
	s.HighlightText = &v
	return s
}

func (s *CreateProductImageRequest) SetImageCount(v int32) *CreateProductImageRequest {
	s.ImageCount = &v
	return s
}

func (s *CreateProductImageRequest) SetImageUrl(v string) *CreateProductImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateProductImageRequest) SetSubTitle(v string) *CreateProductImageRequest {
	s.SubTitle = &v
	return s
}

func (s *CreateProductImageRequest) SetTitle(v string) *CreateProductImageRequest {
	s.Title = &v
	return s
}

type CreateProductImageResponseBody struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 837091359375048704
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateProductImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProductImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductImageResponseBody) SetRequestId(v string) *CreateProductImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProductImageResponseBody) SetTaskId(v string) *CreateProductImageResponseBody {
	s.TaskId = &v
	return s
}

type CreateProductImageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProductImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProductImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProductImageResponse) GoString() string {
	return s.String()
}

func (s *CreateProductImageResponse) SetHeaders(v map[string]*string) *CreateProductImageResponse {
	s.Headers = v
	return s
}

func (s *CreateProductImageResponse) SetStatusCode(v int32) *CreateProductImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProductImageResponse) SetBody(v *CreateProductImageResponseBody) *CreateProductImageResponse {
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

type CreateTrainTaskRequest struct {
	// example:
	//
	// 13168123111
	AliyunMainId *string `json:"aliyunMainId,omitempty" xml:"aliyunMainId,omitempty"`
	// example:
	//
	// BASIC_MODEL
	ResSpecType *string `json:"resSpecType,omitempty" xml:"resSpecType,omitempty"`
	// example:
	//
	// CopyAnchorAndVoice
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// example:
	//
	// realTimeInteractivity
	UseScene *string `json:"useScene,omitempty" xml:"useScene,omitempty"`
	// example:
	//
	// M
	VoiceGender *string `json:"voiceGender,omitempty" xml:"voiceGender,omitempty"`
	// example:
	//
	// zh
	VoiceLanguage *string `json:"voiceLanguage,omitempty" xml:"voiceLanguage,omitempty"`
	VoiceName     *string `json:"voiceName,omitempty" xml:"voiceName,omitempty"`
	// example:
	//
	// https://yic-pre/video/test-0513.mp3
	VoicePath *string `json:"voicePath,omitempty" xml:"voicePath,omitempty"`
}

func (s CreateTrainTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTrainTaskRequest) SetAliyunMainId(v string) *CreateTrainTaskRequest {
	s.AliyunMainId = &v
	return s
}

func (s *CreateTrainTaskRequest) SetResSpecType(v string) *CreateTrainTaskRequest {
	s.ResSpecType = &v
	return s
}

func (s *CreateTrainTaskRequest) SetTaskType(v string) *CreateTrainTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateTrainTaskRequest) SetUseScene(v string) *CreateTrainTaskRequest {
	s.UseScene = &v
	return s
}

func (s *CreateTrainTaskRequest) SetVoiceGender(v string) *CreateTrainTaskRequest {
	s.VoiceGender = &v
	return s
}

func (s *CreateTrainTaskRequest) SetVoiceLanguage(v string) *CreateTrainTaskRequest {
	s.VoiceLanguage = &v
	return s
}

func (s *CreateTrainTaskRequest) SetVoiceName(v string) *CreateTrainTaskRequest {
	s.VoiceName = &v
	return s
}

func (s *CreateTrainTaskRequest) SetVoicePath(v string) *CreateTrainTaskRequest {
	s.VoicePath = &v
	return s
}

type CreateTrainTaskResponseBody struct {
	// example:
	//
	// 84657DE0-B68C-508B-AFE7-8ED921854E3C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 837091359375048704
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateTrainTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrainTaskResponseBody) SetRequestId(v string) *CreateTrainTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrainTaskResponseBody) SetTaskId(v string) *CreateTrainTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateTrainTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrainTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrainTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTrainTaskResponse) SetHeaders(v map[string]*string) *CreateTrainTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTrainTaskResponse) SetStatusCode(v int32) *CreateTrainTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrainTaskResponse) SetBody(v *CreateTrainTaskResponseBody) *CreateTrainTaskResponse {
	s.Body = v
	return s
}

type CreateVideoClipTaskRequest struct {
	// example:
	//
	// 1314445556
	AliyunMainId *string   `json:"aliyunMainId,omitempty" xml:"aliyunMainId,omitempty"`
	Description  *string   `json:"description,omitempty" xml:"description,omitempty"`
	OssKeys      []*string `json:"ossKeys,omitempty" xml:"ossKeys,omitempty" type:"Repeated"`
	Requirement  *string   `json:"requirement,omitempty" xml:"requirement,omitempty"`
}

func (s CreateVideoClipTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoClipTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoClipTaskRequest) SetAliyunMainId(v string) *CreateVideoClipTaskRequest {
	s.AliyunMainId = &v
	return s
}

func (s *CreateVideoClipTaskRequest) SetDescription(v string) *CreateVideoClipTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateVideoClipTaskRequest) SetOssKeys(v []*string) *CreateVideoClipTaskRequest {
	s.OssKeys = v
	return s
}

func (s *CreateVideoClipTaskRequest) SetRequirement(v string) *CreateVideoClipTaskRequest {
	s.Requirement = &v
	return s
}

type CreateVideoClipTaskResponseBody struct {
	// example:
	//
	// 86A90C40-D1AB-50DA-A4B1-0D545F80F2FE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 837091359375048704
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateVideoClipTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoClipTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoClipTaskResponseBody) SetRequestId(v string) *CreateVideoClipTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoClipTaskResponseBody) SetTaskId(v string) *CreateVideoClipTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateVideoClipTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVideoClipTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVideoClipTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoClipTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoClipTaskResponse) SetHeaders(v map[string]*string) *CreateVideoClipTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoClipTaskResponse) SetStatusCode(v int32) *CreateVideoClipTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVideoClipTaskResponse) SetBody(v *CreateVideoClipTaskResponseBody) *CreateVideoClipTaskResponse {
	s.Body = v
	return s
}

type DeleteIndividuationProjectRequest struct {
	// example:
	//
	// 840015278620459008
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s DeleteIndividuationProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndividuationProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteIndividuationProjectRequest) SetProjectId(v string) *DeleteIndividuationProjectRequest {
	s.ProjectId = &v
	return s
}

type DeleteIndividuationProjectResponseBody struct {
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 86A90C40-D1AB-50DA-A4B1-0D545F80F2FE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DeleteIndividuationProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndividuationProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndividuationProjectResponseBody) SetDesc(v string) *DeleteIndividuationProjectResponseBody {
	s.Desc = &v
	return s
}

func (s *DeleteIndividuationProjectResponseBody) SetRequestId(v string) *DeleteIndividuationProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndividuationProjectResponseBody) SetStatus(v string) *DeleteIndividuationProjectResponseBody {
	s.Status = &v
	return s
}

type DeleteIndividuationProjectResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIndividuationProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIndividuationProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndividuationProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndividuationProjectResponse) SetHeaders(v map[string]*string) *DeleteIndividuationProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndividuationProjectResponse) SetStatusCode(v int32) *DeleteIndividuationProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIndividuationProjectResponse) SetBody(v *DeleteIndividuationProjectResponseBody) *DeleteIndividuationProjectResponse {
	s.Body = v
	return s
}

type DeleteIndividuationTextRequest struct {
	TextIdList []*string `json:"textIdList,omitempty" xml:"textIdList,omitempty" type:"Repeated"`
}

func (s DeleteIndividuationTextRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndividuationTextRequest) GoString() string {
	return s.String()
}

func (s *DeleteIndividuationTextRequest) SetTextIdList(v []*string) *DeleteIndividuationTextRequest {
	s.TextIdList = v
	return s
}

type DeleteIndividuationTextResponseBody struct {
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4830493A-728F-5F19-BBCC-1443292E9C49
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DeleteIndividuationTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndividuationTextResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndividuationTextResponseBody) SetDesc(v string) *DeleteIndividuationTextResponseBody {
	s.Desc = &v
	return s
}

func (s *DeleteIndividuationTextResponseBody) SetRequestId(v string) *DeleteIndividuationTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndividuationTextResponseBody) SetStatus(v string) *DeleteIndividuationTextResponseBody {
	s.Status = &v
	return s
}

type DeleteIndividuationTextResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIndividuationTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIndividuationTextResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndividuationTextResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndividuationTextResponse) SetHeaders(v map[string]*string) *DeleteIndividuationTextResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndividuationTextResponse) SetStatusCode(v int32) *DeleteIndividuationTextResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIndividuationTextResponse) SetBody(v *DeleteIndividuationTextResponseBody) *DeleteIndividuationTextResponse {
	s.Body = v
	return s
}

type DescribeDocumentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocumentResult    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocumentResponse) GoString() string {
	return s.String()
}

func (s *DescribeDocumentResponse) SetHeaders(v map[string]*string) *DescribeDocumentResponse {
	s.Headers = v
	return s
}

func (s *DescribeDocumentResponse) SetStatusCode(v int32) *DescribeDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDocumentResponse) SetBody(v *DocumentResult) *DescribeDocumentResponse {
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

type GetAICoachAssessmentPointRequest struct {
	// example:
	//
	// 1
	PointId *string `json:"pointId,omitempty" xml:"pointId,omitempty"`
}

func (s GetAICoachAssessmentPointRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachAssessmentPointRequest) GoString() string {
	return s.String()
}

func (s *GetAICoachAssessmentPointRequest) SetPointId(v string) *GetAICoachAssessmentPointRequest {
	s.PointId = &v
	return s
}

type GetAICoachAssessmentPointResponseBody struct {
	AnswerList []*GetAICoachAssessmentPointResponseBodyAnswerList `json:"answerList,omitempty" xml:"answerList,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Citations *int32 `json:"citations,omitempty" xml:"citations,omitempty"`
	// example:
	//
	// 1
	DocumentId *string `json:"documentId,omitempty" xml:"documentId,omitempty"`
	// example:
	//
	// demo
	DocumentName *string `json:"documentName,omitempty" xml:"documentName,omitempty"`
	// example:
	//
	// 2025-02-24 12:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2025-02-24 12:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	KbId *string `json:"kbId,omitempty" xml:"kbId,omitempty"`
	// example:
	//
	// Cloudcode
	KbType        *string   `json:"kbType,omitempty" xml:"kbType,omitempty"`
	KnowledgeList []*string `json:"knowledgeList,omitempty" xml:"knowledgeList,omitempty" type:"Repeated"`
	// example:
	//
	// demo
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	PointId *string `json:"pointId,omitempty" xml:"pointId,omitempty"`
	// example:
	//
	// demo
	QuestionDescription *string `json:"questionDescription,omitempty" xml:"questionDescription,omitempty"`
	// example:
	//
	// demo
	QuestionSample *string `json:"questionSample,omitempty" xml:"questionSample,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4830493A-728F-5F19-BBCC-1443292E9C49
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// DRAFT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetAICoachAssessmentPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachAssessmentPointResponseBody) GoString() string {
	return s.String()
}

func (s *GetAICoachAssessmentPointResponseBody) SetAnswerList(v []*GetAICoachAssessmentPointResponseBodyAnswerList) *GetAICoachAssessmentPointResponseBody {
	s.AnswerList = v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetCitations(v int32) *GetAICoachAssessmentPointResponseBody {
	s.Citations = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetDocumentId(v string) *GetAICoachAssessmentPointResponseBody {
	s.DocumentId = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetDocumentName(v string) *GetAICoachAssessmentPointResponseBody {
	s.DocumentName = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetGmtCreate(v string) *GetAICoachAssessmentPointResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetGmtModified(v string) *GetAICoachAssessmentPointResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetKbId(v string) *GetAICoachAssessmentPointResponseBody {
	s.KbId = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetKbType(v string) *GetAICoachAssessmentPointResponseBody {
	s.KbType = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetKnowledgeList(v []*string) *GetAICoachAssessmentPointResponseBody {
	s.KnowledgeList = v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetName(v string) *GetAICoachAssessmentPointResponseBody {
	s.Name = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetPointId(v string) *GetAICoachAssessmentPointResponseBody {
	s.PointId = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetQuestionDescription(v string) *GetAICoachAssessmentPointResponseBody {
	s.QuestionDescription = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetQuestionSample(v string) *GetAICoachAssessmentPointResponseBody {
	s.QuestionSample = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetRequestId(v string) *GetAICoachAssessmentPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetStatus(v string) *GetAICoachAssessmentPointResponseBody {
	s.Status = &v
	return s
}

type GetAICoachAssessmentPointResponseBodyAnswerList struct {
	AnswerValues []*GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues `json:"answerValues,omitempty" xml:"answerValues,omitempty" type:"Repeated"`
	// example:
	//
	// true
	EnabledKeyword *bool     `json:"enabledKeyword,omitempty" xml:"enabledKeyword,omitempty"`
	NameList       []*string `json:"nameList,omitempty" xml:"nameList,omitempty" type:"Repeated"`
	// example:
	//
	// and
	Operators  *string                                                      `json:"operators,omitempty" xml:"operators,omitempty"`
	Parameters []*GetAICoachAssessmentPointResponseBodyAnswerListParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	// example:
	//
	// custom
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetAICoachAssessmentPointResponseBodyAnswerList) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachAssessmentPointResponseBodyAnswerList) GoString() string {
	return s.String()
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) SetAnswerValues(v []*GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) *GetAICoachAssessmentPointResponseBodyAnswerList {
	s.AnswerValues = v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) SetEnabledKeyword(v bool) *GetAICoachAssessmentPointResponseBodyAnswerList {
	s.EnabledKeyword = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) SetNameList(v []*string) *GetAICoachAssessmentPointResponseBodyAnswerList {
	s.NameList = v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) SetOperators(v string) *GetAICoachAssessmentPointResponseBodyAnswerList {
	s.Operators = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) SetParameters(v []*GetAICoachAssessmentPointResponseBodyAnswerListParameters) *GetAICoachAssessmentPointResponseBodyAnswerList {
	s.Parameters = v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) SetType(v string) *GetAICoachAssessmentPointResponseBodyAnswerList {
	s.Type = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) SetWeight(v int32) *GetAICoachAssessmentPointResponseBodyAnswerList {
	s.Weight = &v
	return s
}

type GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues struct {
	// example:
	//
	// demo
	AnswerName *string `json:"answerName,omitempty" xml:"answerName,omitempty"`
	// example:
	//
	// 50
	AnswerWeight  *int32                                                                      `json:"answerWeight,omitempty" xml:"answerWeight,omitempty"`
	KeywordValues []*GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues `json:"keywordValues,omitempty" xml:"keywordValues,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	KeywordWeight *int32                                                                     `json:"keywordWeight,omitempty" xml:"keywordWeight,omitempty"`
	ScoringRules  []*GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules `json:"scoringRules,omitempty" xml:"scoringRules,omitempty" type:"Repeated"`
}

func (s GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) GoString() string {
	return s.String()
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) SetAnswerName(v string) *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues {
	s.AnswerName = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) SetAnswerWeight(v int32) *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues {
	s.AnswerWeight = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) SetKeywordValues(v []*GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues) *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues {
	s.KeywordValues = v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) SetKeywordWeight(v int32) *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues {
	s.KeywordWeight = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) SetScoringRules(v []*GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules) *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues {
	s.ScoringRules = v
	return s
}

type GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues struct {
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 50
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues) GoString() string {
	return s.String()
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues) SetName(v string) *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues {
	s.Name = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues) SetWeight(v int32) *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues {
	s.Weight = &v
	return s
}

type GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules struct {
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules) GoString() string {
	return s.String()
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules) SetName(v string) *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules {
	s.Name = &v
	return s
}

type GetAICoachAssessmentPointResponseBodyAnswerListParameters struct {
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 441323200602114284
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetAICoachAssessmentPointResponseBodyAnswerListParameters) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachAssessmentPointResponseBodyAnswerListParameters) GoString() string {
	return s.String()
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListParameters) SetName(v string) *GetAICoachAssessmentPointResponseBodyAnswerListParameters {
	s.Name = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListParameters) SetValue(v string) *GetAICoachAssessmentPointResponseBodyAnswerListParameters {
	s.Value = &v
	return s
}

type GetAICoachAssessmentPointResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAICoachAssessmentPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAICoachAssessmentPointResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachAssessmentPointResponse) GoString() string {
	return s.String()
}

func (s *GetAICoachAssessmentPointResponse) SetHeaders(v map[string]*string) *GetAICoachAssessmentPointResponse {
	s.Headers = v
	return s
}

func (s *GetAICoachAssessmentPointResponse) SetStatusCode(v int32) *GetAICoachAssessmentPointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAICoachAssessmentPointResponse) SetBody(v *GetAICoachAssessmentPointResponseBody) *GetAICoachAssessmentPointResponse {
	s.Body = v
	return s
}

type GetAICoachCheatDetectionRequest struct {
	// example:
	//
	// 79e954faffe2415ebd18188ba787d78e
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s GetAICoachCheatDetectionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachCheatDetectionRequest) GoString() string {
	return s.String()
}

func (s *GetAICoachCheatDetectionRequest) SetSessionId(v string) *GetAICoachCheatDetectionRequest {
	s.SessionId = &v
	return s
}

type GetAICoachCheatDetectionResponseBody struct {
	// example:
	//
	// 1
	CheatId *string `json:"cheatId,omitempty" xml:"cheatId,omitempty"`
	// example:
	//
	// success
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 2025-02-24 12:00:00
	GmtCreate  *string                                         `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	ImageCheat *GetAICoachCheatDetectionResponseBodyImageCheat `json:"imageCheat,omitempty" xml:"imageCheat,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 0E8B1746-AE35-5C4B-A3A8-345B274AE32C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true
	//
	// example:
	//
	// True
	Success    *bool                                           `json:"success,omitempty" xml:"success,omitempty"`
	VoiceCheat *GetAICoachCheatDetectionResponseBodyVoiceCheat `json:"voiceCheat,omitempty" xml:"voiceCheat,omitempty" type:"Struct"`
}

func (s GetAICoachCheatDetectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachCheatDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAICoachCheatDetectionResponseBody) SetCheatId(v string) *GetAICoachCheatDetectionResponseBody {
	s.CheatId = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) SetErrorCode(v string) *GetAICoachCheatDetectionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) SetErrorMessage(v string) *GetAICoachCheatDetectionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) SetGmtCreate(v string) *GetAICoachCheatDetectionResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) SetImageCheat(v *GetAICoachCheatDetectionResponseBodyImageCheat) *GetAICoachCheatDetectionResponseBody {
	s.ImageCheat = v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) SetRequestId(v string) *GetAICoachCheatDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) SetStatus(v int32) *GetAICoachCheatDetectionResponseBody {
	s.Status = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) SetSuccess(v bool) *GetAICoachCheatDetectionResponseBody {
	s.Success = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBody) SetVoiceCheat(v *GetAICoachCheatDetectionResponseBodyVoiceCheat) *GetAICoachCheatDetectionResponseBody {
	s.VoiceCheat = v
	return s
}

type GetAICoachCheatDetectionResponseBodyImageCheat struct {
	// example:
	//
	// demo
	Desc *string                                               `json:"desc,omitempty" xml:"desc,omitempty"`
	List []*GetAICoachCheatDetectionResponseBodyImageCheatList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetAICoachCheatDetectionResponseBodyImageCheat) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachCheatDetectionResponseBodyImageCheat) GoString() string {
	return s.String()
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheat) SetDesc(v string) *GetAICoachCheatDetectionResponseBodyImageCheat {
	s.Desc = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheat) SetList(v []*GetAICoachCheatDetectionResponseBodyImageCheatList) *GetAICoachCheatDetectionResponseBodyImageCheat {
	s.List = v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheat) SetStatus(v int32) *GetAICoachCheatDetectionResponseBodyImageCheat {
	s.Status = &v
	return s
}

type GetAICoachCheatDetectionResponseBodyImageCheatList struct {
	// example:
	//
	// 2025-03-22 10:05:07
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// https://demo.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetAICoachCheatDetectionResponseBodyImageCheatList) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachCheatDetectionResponseBodyImageCheatList) GoString() string {
	return s.String()
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheatList) SetTime(v string) *GetAICoachCheatDetectionResponseBodyImageCheatList {
	s.Time = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyImageCheatList) SetUrl(v string) *GetAICoachCheatDetectionResponseBodyImageCheatList {
	s.Url = &v
	return s
}

type GetAICoachCheatDetectionResponseBodyVoiceCheat struct {
	ComparisonList []*GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList `json:"comparisonList,omitempty" xml:"comparisonList,omitempty" type:"Repeated"`
	// example:
	//
	// demo
	Desc         *string                                                       `json:"desc,omitempty" xml:"desc,omitempty"`
	OriginalList []*GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList `json:"originalList,omitempty" xml:"originalList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetAICoachCheatDetectionResponseBodyVoiceCheat) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachCheatDetectionResponseBodyVoiceCheat) GoString() string {
	return s.String()
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheat) SetComparisonList(v []*GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList) *GetAICoachCheatDetectionResponseBodyVoiceCheat {
	s.ComparisonList = v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheat) SetDesc(v string) *GetAICoachCheatDetectionResponseBodyVoiceCheat {
	s.Desc = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheat) SetOriginalList(v []*GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList) *GetAICoachCheatDetectionResponseBodyVoiceCheat {
	s.OriginalList = v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheat) SetStatus(v int32) *GetAICoachCheatDetectionResponseBodyVoiceCheat {
	s.Status = &v
	return s
}

type GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList struct {
	// example:
	//
	// 2024-12-11 10:07:23
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// https://demo.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList) GoString() string {
	return s.String()
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList) SetTime(v string) *GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList {
	s.Time = &v
	return s
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList) SetUrl(v string) *GetAICoachCheatDetectionResponseBodyVoiceCheatComparisonList {
	s.Url = &v
	return s
}

type GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList struct {
	// example:
	//
	// https://demo.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList) GoString() string {
	return s.String()
}

func (s *GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList) SetUrl(v string) *GetAICoachCheatDetectionResponseBodyVoiceCheatOriginalList {
	s.Url = &v
	return s
}

type GetAICoachCheatDetectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAICoachCheatDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAICoachCheatDetectionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachCheatDetectionResponse) GoString() string {
	return s.String()
}

func (s *GetAICoachCheatDetectionResponse) SetHeaders(v map[string]*string) *GetAICoachCheatDetectionResponse {
	s.Headers = v
	return s
}

func (s *GetAICoachCheatDetectionResponse) SetStatusCode(v int32) *GetAICoachCheatDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAICoachCheatDetectionResponse) SetBody(v *GetAICoachCheatDetectionResponseBody) *GetAICoachCheatDetectionResponse {
	s.Body = v
	return s
}

type GetAICoachScriptRequest struct {
	// example:
	//
	// 1
	ScriptRecordId *string `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
}

func (s GetAICoachScriptRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachScriptRequest) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptRequest) SetScriptRecordId(v string) *GetAICoachScriptRequest {
	s.ScriptRecordId = &v
	return s
}

type GetAICoachScriptResponseBody struct {
	AppendQuestionFlag *bool `json:"appendQuestionFlag,omitempty" xml:"appendQuestionFlag,omitempty"`
	// example:
	//
	// point
	AssessmentScope  *string                                       `json:"assessmentScope,omitempty" xml:"assessmentScope,omitempty"`
	CheckCheatConfig *GetAICoachScriptResponseBodyCheckCheatConfig `json:"checkCheatConfig,omitempty" xml:"checkCheatConfig,omitempty" type:"Struct"`
	ClosingRemarks   *string                                       `json:"closingRemarks,omitempty" xml:"closingRemarks,omitempty"`
	CompleteStrategy *GetAICoachScriptResponseBodyCompleteStrategy `json:"completeStrategy,omitempty" xml:"completeStrategy,omitempty" type:"Struct"`
	// example:
	//
	// https://demo.com
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// example:
	//
	// 500
	DialogueInputTextLimit *int32 `json:"dialogueInputTextLimit,omitempty" xml:"dialogueInputTextLimit,omitempty"`
	// example:
	//
	// true
	DialogueTextFlag *bool `json:"dialogueTextFlag,omitempty" xml:"dialogueTextFlag,omitempty"`
	// example:
	//
	// true
	DialogueTipFlag *bool `json:"dialogueTipFlag,omitempty" xml:"dialogueTipFlag,omitempty"`
	// example:
	//
	// 30
	DialogueVoiceLimit *int32 `json:"dialogueVoiceLimit,omitempty" xml:"dialogueVoiceLimit,omitempty"`
	// example:
	//
	// true
	EvaluateReportFlag *bool                                             `json:"evaluateReportFlag,omitempty" xml:"evaluateReportFlag,omitempty"`
	Expressiveness     map[string]*int32                                 `json:"expressiveness,omitempty" xml:"expressiveness,omitempty"`
	ExpressivenessList []*GetAICoachScriptResponseBodyExpressivenessList `json:"expressivenessList,omitempty" xml:"expressivenessList,omitempty" type:"Repeated"`
	GifDynamicUrl      *string                                           `json:"gifDynamicUrl,omitempty" xml:"gifDynamicUrl,omitempty"`
	GifStaticUrl       *string                                           `json:"gifStaticUrl,omitempty" xml:"gifStaticUrl,omitempty"`
	// example:
	//
	// 2025-02-24 12:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2025-02-24 12:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// coach
	Initiator             *string   `json:"initiator,omitempty" xml:"initiator,omitempty"`
	InteractionInputTypes []*string `json:"interactionInputTypes,omitempty" xml:"interactionInputTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	InteractionType *int32 `json:"interactionType,omitempty" xml:"interactionType,omitempty"`
	// example:
	//
	// demo
	Introduce *string `json:"introduce,omitempty" xml:"introduce,omitempty"`
	// example:
	//
	// demo
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OpeningRemarks *string `json:"openingRemarks,omitempty" xml:"openingRemarks,omitempty"`
	// example:
	//
	// true
	OrderAckFlag           *bool                                                 `json:"orderAckFlag,omitempty" xml:"orderAckFlag,omitempty"`
	PointDeductionRuleList []*GetAICoachScriptResponseBodyPointDeductionRuleList `json:"pointDeductionRuleList,omitempty" xml:"pointDeductionRuleList,omitempty" type:"Repeated"`
	Points                 []*GetAICoachScriptResponseBodyPoints                 `json:"points,omitempty" xml:"points,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	RequestId          *string                                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SampleDialogueList []*GetAICoachScriptResponseBodySampleDialogueList `json:"sampleDialogueList,omitempty" xml:"sampleDialogueList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	ScriptRecordId     *string `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
	SparringTipContent *string `json:"sparringTipContent,omitempty" xml:"sparringTipContent,omitempty"`
	SparringTipTitle   *string `json:"sparringTipTitle,omitempty" xml:"sparringTipTitle,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	StudentThinkTimeFlag *bool `json:"studentThinkTimeFlag,omitempty" xml:"studentThinkTimeFlag,omitempty"`
	// example:
	//
	// 100
	StudentThinkTimeLimit *int32 `json:"studentThinkTimeLimit,omitempty" xml:"studentThinkTimeLimit,omitempty"`
	// example:
	//
	// 1
	Type    *int32                               `json:"type,omitempty" xml:"type,omitempty"`
	Weights *GetAICoachScriptResponseBodyWeights `json:"weights,omitempty" xml:"weights,omitempty" type:"Struct"`
}

func (s GetAICoachScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachScriptResponseBody) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBody) SetAppendQuestionFlag(v bool) *GetAICoachScriptResponseBody {
	s.AppendQuestionFlag = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetAssessmentScope(v string) *GetAICoachScriptResponseBody {
	s.AssessmentScope = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetCheckCheatConfig(v *GetAICoachScriptResponseBodyCheckCheatConfig) *GetAICoachScriptResponseBody {
	s.CheckCheatConfig = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetClosingRemarks(v string) *GetAICoachScriptResponseBody {
	s.ClosingRemarks = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetCompleteStrategy(v *GetAICoachScriptResponseBodyCompleteStrategy) *GetAICoachScriptResponseBody {
	s.CompleteStrategy = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetCoverUrl(v string) *GetAICoachScriptResponseBody {
	s.CoverUrl = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetDialogueInputTextLimit(v int32) *GetAICoachScriptResponseBody {
	s.DialogueInputTextLimit = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetDialogueTextFlag(v bool) *GetAICoachScriptResponseBody {
	s.DialogueTextFlag = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetDialogueTipFlag(v bool) *GetAICoachScriptResponseBody {
	s.DialogueTipFlag = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetDialogueVoiceLimit(v int32) *GetAICoachScriptResponseBody {
	s.DialogueVoiceLimit = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetEvaluateReportFlag(v bool) *GetAICoachScriptResponseBody {
	s.EvaluateReportFlag = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetExpressiveness(v map[string]*int32) *GetAICoachScriptResponseBody {
	s.Expressiveness = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetExpressivenessList(v []*GetAICoachScriptResponseBodyExpressivenessList) *GetAICoachScriptResponseBody {
	s.ExpressivenessList = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetGifDynamicUrl(v string) *GetAICoachScriptResponseBody {
	s.GifDynamicUrl = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetGifStaticUrl(v string) *GetAICoachScriptResponseBody {
	s.GifStaticUrl = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetGmtCreate(v string) *GetAICoachScriptResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetGmtModified(v string) *GetAICoachScriptResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetInitiator(v string) *GetAICoachScriptResponseBody {
	s.Initiator = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetInteractionInputTypes(v []*string) *GetAICoachScriptResponseBody {
	s.InteractionInputTypes = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetInteractionType(v int32) *GetAICoachScriptResponseBody {
	s.InteractionType = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetIntroduce(v string) *GetAICoachScriptResponseBody {
	s.Introduce = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetName(v string) *GetAICoachScriptResponseBody {
	s.Name = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetOpeningRemarks(v string) *GetAICoachScriptResponseBody {
	s.OpeningRemarks = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetOrderAckFlag(v bool) *GetAICoachScriptResponseBody {
	s.OrderAckFlag = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetPointDeductionRuleList(v []*GetAICoachScriptResponseBodyPointDeductionRuleList) *GetAICoachScriptResponseBody {
	s.PointDeductionRuleList = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetPoints(v []*GetAICoachScriptResponseBodyPoints) *GetAICoachScriptResponseBody {
	s.Points = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetRequestId(v string) *GetAICoachScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetSampleDialogueList(v []*GetAICoachScriptResponseBodySampleDialogueList) *GetAICoachScriptResponseBody {
	s.SampleDialogueList = v
	return s
}

func (s *GetAICoachScriptResponseBody) SetScriptRecordId(v string) *GetAICoachScriptResponseBody {
	s.ScriptRecordId = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetSparringTipContent(v string) *GetAICoachScriptResponseBody {
	s.SparringTipContent = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetSparringTipTitle(v string) *GetAICoachScriptResponseBody {
	s.SparringTipTitle = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetStatus(v int32) *GetAICoachScriptResponseBody {
	s.Status = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetStudentThinkTimeFlag(v bool) *GetAICoachScriptResponseBody {
	s.StudentThinkTimeFlag = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetStudentThinkTimeLimit(v int32) *GetAICoachScriptResponseBody {
	s.StudentThinkTimeLimit = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetType(v int32) *GetAICoachScriptResponseBody {
	s.Type = &v
	return s
}

func (s *GetAICoachScriptResponseBody) SetWeights(v *GetAICoachScriptResponseBodyWeights) *GetAICoachScriptResponseBody {
	s.Weights = v
	return s
}

type GetAICoachScriptResponseBodyCheckCheatConfig struct {
	CheckImage *bool `json:"checkImage,omitempty" xml:"checkImage,omitempty"`
	CheckVoice *bool `json:"checkVoice,omitempty" xml:"checkVoice,omitempty"`
}

func (s GetAICoachScriptResponseBodyCheckCheatConfig) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachScriptResponseBodyCheckCheatConfig) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyCheckCheatConfig) SetCheckImage(v bool) *GetAICoachScriptResponseBodyCheckCheatConfig {
	s.CheckImage = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCheckCheatConfig) SetCheckVoice(v bool) *GetAICoachScriptResponseBodyCheckCheatConfig {
	s.CheckVoice = &v
	return s
}

type GetAICoachScriptResponseBodyCompleteStrategy struct {
	// example:
	//
	// 5
	AbnormalQuitSessionExpired *int32 `json:"abnormalQuitSessionExpired,omitempty" xml:"abnormalQuitSessionExpired,omitempty"`
	// example:
	//
	// true
	AbnormalQuitSessionExpiredFlag *bool `json:"abnormalQuitSessionExpiredFlag,omitempty" xml:"abnormalQuitSessionExpiredFlag,omitempty"`
	// example:
	//
	// true
	ClickCompleteAutoEnd *bool `json:"clickCompleteAutoEnd,omitempty" xml:"clickCompleteAutoEnd,omitempty"`
	// example:
	//
	// 15
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// true
	DurationFlag *bool `json:"durationFlag,omitempty" xml:"durationFlag,omitempty"`
	// example:
	//
	// true
	FullCoverageAutoEnd *bool `json:"fullCoverageAutoEnd,omitempty" xml:"fullCoverageAutoEnd,omitempty"`
}

func (s GetAICoachScriptResponseBodyCompleteStrategy) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachScriptResponseBodyCompleteStrategy) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) SetAbnormalQuitSessionExpired(v int32) *GetAICoachScriptResponseBodyCompleteStrategy {
	s.AbnormalQuitSessionExpired = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) SetAbnormalQuitSessionExpiredFlag(v bool) *GetAICoachScriptResponseBodyCompleteStrategy {
	s.AbnormalQuitSessionExpiredFlag = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) SetClickCompleteAutoEnd(v bool) *GetAICoachScriptResponseBodyCompleteStrategy {
	s.ClickCompleteAutoEnd = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) SetDuration(v int32) *GetAICoachScriptResponseBodyCompleteStrategy {
	s.Duration = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) SetDurationFlag(v bool) *GetAICoachScriptResponseBodyCompleteStrategy {
	s.DurationFlag = &v
	return s
}

func (s *GetAICoachScriptResponseBodyCompleteStrategy) SetFullCoverageAutoEnd(v bool) *GetAICoachScriptResponseBodyCompleteStrategy {
	s.FullCoverageAutoEnd = &v
	return s
}

type GetAICoachScriptResponseBodyExpressivenessList struct {
	Desc             *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Enabled          *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	ExpressivenessId *string `json:"expressivenessId,omitempty" xml:"expressivenessId,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	Rule             *string `json:"rule,omitempty" xml:"rule,omitempty"`
	Type             *string `json:"type,omitempty" xml:"type,omitempty"`
	Weight           *int32  `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetAICoachScriptResponseBodyExpressivenessList) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachScriptResponseBodyExpressivenessList) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) SetDesc(v string) *GetAICoachScriptResponseBodyExpressivenessList {
	s.Desc = &v
	return s
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) SetEnabled(v bool) *GetAICoachScriptResponseBodyExpressivenessList {
	s.Enabled = &v
	return s
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) SetExpressivenessId(v string) *GetAICoachScriptResponseBodyExpressivenessList {
	s.ExpressivenessId = &v
	return s
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) SetName(v string) *GetAICoachScriptResponseBodyExpressivenessList {
	s.Name = &v
	return s
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) SetRule(v string) *GetAICoachScriptResponseBodyExpressivenessList {
	s.Rule = &v
	return s
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) SetType(v string) *GetAICoachScriptResponseBodyExpressivenessList {
	s.Type = &v
	return s
}

func (s *GetAICoachScriptResponseBodyExpressivenessList) SetWeight(v int32) *GetAICoachScriptResponseBodyExpressivenessList {
	s.Weight = &v
	return s
}

type GetAICoachScriptResponseBodyPointDeductionRuleList struct {
	// example:
	//
	// demo
	Description     *string   `json:"description,omitempty" xml:"description,omitempty"`
	PunishmentTypes []*string `json:"punishmentTypes,omitempty" xml:"punishmentTypes,omitempty" type:"Repeated"`
	RuleValue       *string   `json:"ruleValue,omitempty" xml:"ruleValue,omitempty"`
	// example:
	//
	// 90
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetAICoachScriptResponseBodyPointDeductionRuleList) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachScriptResponseBodyPointDeductionRuleList) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyPointDeductionRuleList) SetDescription(v string) *GetAICoachScriptResponseBodyPointDeductionRuleList {
	s.Description = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointDeductionRuleList) SetPunishmentTypes(v []*string) *GetAICoachScriptResponseBodyPointDeductionRuleList {
	s.PunishmentTypes = v
	return s
}

func (s *GetAICoachScriptResponseBodyPointDeductionRuleList) SetRuleValue(v string) *GetAICoachScriptResponseBodyPointDeductionRuleList {
	s.RuleValue = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointDeductionRuleList) SetWeight(v int32) *GetAICoachScriptResponseBodyPointDeductionRuleList {
	s.Weight = &v
	return s
}

type GetAICoachScriptResponseBodyPoints struct {
	AnswerList    []*GetAICoachScriptResponseBodyPointsAnswerList `json:"answerList,omitempty" xml:"answerList,omitempty" type:"Repeated"`
	KnowledgeList []*string                                       `json:"knowledgeList,omitempty" xml:"knowledgeList,omitempty" type:"Repeated"`
	// example:
	//
	// demo
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	PointId *string `json:"pointId,omitempty" xml:"pointId,omitempty"`
	// example:
	//
	// test
	QuestionDescription *string `json:"questionDescription,omitempty" xml:"questionDescription,omitempty"`
	// example:
	//
	// 1
	SortNo *int32 `json:"sortNo,omitempty" xml:"sortNo,omitempty"`
	// example:
	//
	// 50
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetAICoachScriptResponseBodyPoints) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachScriptResponseBodyPoints) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyPoints) SetAnswerList(v []*GetAICoachScriptResponseBodyPointsAnswerList) *GetAICoachScriptResponseBodyPoints {
	s.AnswerList = v
	return s
}

func (s *GetAICoachScriptResponseBodyPoints) SetKnowledgeList(v []*string) *GetAICoachScriptResponseBodyPoints {
	s.KnowledgeList = v
	return s
}

func (s *GetAICoachScriptResponseBodyPoints) SetName(v string) *GetAICoachScriptResponseBodyPoints {
	s.Name = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPoints) SetPointId(v string) *GetAICoachScriptResponseBodyPoints {
	s.PointId = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPoints) SetQuestionDescription(v string) *GetAICoachScriptResponseBodyPoints {
	s.QuestionDescription = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPoints) SetSortNo(v int32) *GetAICoachScriptResponseBodyPoints {
	s.SortNo = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPoints) SetWeight(v int32) *GetAICoachScriptResponseBodyPoints {
	s.Weight = &v
	return s
}

type GetAICoachScriptResponseBodyPointsAnswerList struct {
	AnswerValues   []*GetAICoachScriptResponseBodyPointsAnswerListAnswerValues `json:"answerValues,omitempty" xml:"answerValues,omitempty" type:"Repeated"`
	EnabledKeyword *bool                                                       `json:"enabledKeyword,omitempty" xml:"enabledKeyword,omitempty"`
	Name           *string                                                     `json:"name,omitempty" xml:"name,omitempty"`
	NameList       []*string                                                   `json:"nameList,omitempty" xml:"nameList,omitempty" type:"Repeated"`
	Operators      *string                                                     `json:"operators,omitempty" xml:"operators,omitempty"`
	Parameters     []*GetAICoachScriptResponseBodyPointsAnswerListParameters   `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	// example:
	//
	// normalKnowledge
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetAICoachScriptResponseBodyPointsAnswerList) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachScriptResponseBodyPointsAnswerList) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) SetAnswerValues(v []*GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) *GetAICoachScriptResponseBodyPointsAnswerList {
	s.AnswerValues = v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) SetEnabledKeyword(v bool) *GetAICoachScriptResponseBodyPointsAnswerList {
	s.EnabledKeyword = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) SetName(v string) *GetAICoachScriptResponseBodyPointsAnswerList {
	s.Name = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) SetNameList(v []*string) *GetAICoachScriptResponseBodyPointsAnswerList {
	s.NameList = v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) SetOperators(v string) *GetAICoachScriptResponseBodyPointsAnswerList {
	s.Operators = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) SetParameters(v []*GetAICoachScriptResponseBodyPointsAnswerListParameters) *GetAICoachScriptResponseBodyPointsAnswerList {
	s.Parameters = v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) SetType(v string) *GetAICoachScriptResponseBodyPointsAnswerList {
	s.Type = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerList) SetWeight(v int32) *GetAICoachScriptResponseBodyPointsAnswerList {
	s.Weight = &v
	return s
}

type GetAICoachScriptResponseBodyPointsAnswerListAnswerValues struct {
	AnswerName    *string                                                                  `json:"answerName,omitempty" xml:"answerName,omitempty"`
	AnswerWeight  *int32                                                                   `json:"answerWeight,omitempty" xml:"answerWeight,omitempty"`
	KeywordValues []*GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues `json:"keywordValues,omitempty" xml:"keywordValues,omitempty" type:"Repeated"`
	KeywordWeight *int32                                                                   `json:"keywordWeight,omitempty" xml:"keywordWeight,omitempty"`
	ScoringRules  []*GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules  `json:"scoringRules,omitempty" xml:"scoringRules,omitempty" type:"Repeated"`
}

func (s GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) SetAnswerName(v string) *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues {
	s.AnswerName = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) SetAnswerWeight(v int32) *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues {
	s.AnswerWeight = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) SetKeywordValues(v []*GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues) *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues {
	s.KeywordValues = v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) SetKeywordWeight(v int32) *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues {
	s.KeywordWeight = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues) SetScoringRules(v []*GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules) *GetAICoachScriptResponseBodyPointsAnswerListAnswerValues {
	s.ScoringRules = v
	return s
}

type GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Weight *int32  `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues) SetName(v string) *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues {
	s.Name = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues) SetWeight(v int32) *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesKeywordValues {
	s.Weight = &v
	return s
}

type GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules) SetName(v string) *GetAICoachScriptResponseBodyPointsAnswerListAnswerValuesScoringRules {
	s.Name = &v
	return s
}

type GetAICoachScriptResponseBodyPointsAnswerListParameters struct {
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetAICoachScriptResponseBodyPointsAnswerListParameters) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachScriptResponseBodyPointsAnswerListParameters) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListParameters) SetName(v string) *GetAICoachScriptResponseBodyPointsAnswerListParameters {
	s.Name = &v
	return s
}

func (s *GetAICoachScriptResponseBodyPointsAnswerListParameters) SetValue(v string) *GetAICoachScriptResponseBodyPointsAnswerListParameters {
	s.Value = &v
	return s
}

type GetAICoachScriptResponseBodySampleDialogueList struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// coach
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GetAICoachScriptResponseBodySampleDialogueList) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachScriptResponseBodySampleDialogueList) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodySampleDialogueList) SetMessage(v string) *GetAICoachScriptResponseBodySampleDialogueList {
	s.Message = &v
	return s
}

func (s *GetAICoachScriptResponseBodySampleDialogueList) SetRole(v string) *GetAICoachScriptResponseBodySampleDialogueList {
	s.Role = &v
	return s
}

type GetAICoachScriptResponseBodyWeights struct {
	// example:
	//
	// 10
	AbilityEvaluation *int32 `json:"abilityEvaluation,omitempty" xml:"abilityEvaluation,omitempty"`
	// example:
	//
	// false
	AbilityEvaluationEnabled *bool `json:"abilityEvaluationEnabled,omitempty" xml:"abilityEvaluationEnabled,omitempty"`
	// example:
	//
	// 10
	AssessmentPoint        *int32 `json:"assessmentPoint,omitempty" xml:"assessmentPoint,omitempty"`
	AssessmentPointEnabled *bool  `json:"assessmentPointEnabled,omitempty" xml:"assessmentPointEnabled,omitempty"`
	// example:
	//
	// 10
	Expressiveness *int32 `json:"expressiveness,omitempty" xml:"expressiveness,omitempty"`
	// example:
	//
	// true
	ExpressivenessEnabled *bool `json:"expressivenessEnabled,omitempty" xml:"expressivenessEnabled,omitempty"`
	// example:
	//
	// 10
	PointDeductionRule *int32 `json:"pointDeductionRule,omitempty" xml:"pointDeductionRule,omitempty"`
	// example:
	//
	// true
	PointDeductionRuleEnabled *bool `json:"pointDeductionRuleEnabled,omitempty" xml:"pointDeductionRuleEnabled,omitempty"`
	// example:
	//
	// 10
	Standard *int32 `json:"standard,omitempty" xml:"standard,omitempty"`
	// example:
	//
	// true
	StandardEnabled *bool `json:"standardEnabled,omitempty" xml:"standardEnabled,omitempty"`
}

func (s GetAICoachScriptResponseBodyWeights) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachScriptResponseBodyWeights) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponseBodyWeights) SetAbilityEvaluation(v int32) *GetAICoachScriptResponseBodyWeights {
	s.AbilityEvaluation = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetAbilityEvaluationEnabled(v bool) *GetAICoachScriptResponseBodyWeights {
	s.AbilityEvaluationEnabled = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetAssessmentPoint(v int32) *GetAICoachScriptResponseBodyWeights {
	s.AssessmentPoint = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetAssessmentPointEnabled(v bool) *GetAICoachScriptResponseBodyWeights {
	s.AssessmentPointEnabled = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetExpressiveness(v int32) *GetAICoachScriptResponseBodyWeights {
	s.Expressiveness = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetExpressivenessEnabled(v bool) *GetAICoachScriptResponseBodyWeights {
	s.ExpressivenessEnabled = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetPointDeductionRule(v int32) *GetAICoachScriptResponseBodyWeights {
	s.PointDeductionRule = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetPointDeductionRuleEnabled(v bool) *GetAICoachScriptResponseBodyWeights {
	s.PointDeductionRuleEnabled = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetStandard(v int32) *GetAICoachScriptResponseBodyWeights {
	s.Standard = &v
	return s
}

func (s *GetAICoachScriptResponseBodyWeights) SetStandardEnabled(v bool) *GetAICoachScriptResponseBodyWeights {
	s.StandardEnabled = &v
	return s
}

type GetAICoachScriptResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAICoachScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAICoachScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachScriptResponse) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptResponse) SetHeaders(v map[string]*string) *GetAICoachScriptResponse {
	s.Headers = v
	return s
}

func (s *GetAICoachScriptResponse) SetStatusCode(v int32) *GetAICoachScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAICoachScriptResponse) SetBody(v *GetAICoachScriptResponseBody) *GetAICoachScriptResponse {
	s.Body = v
	return s
}

type GetAICoachTaskSessionHistoryRequest struct {
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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

func (s *GetAICoachTaskSessionHistoryRequest) SetPageNumber(v int32) *GetAICoachTaskSessionHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryRequest) SetPageSize(v int32) *GetAICoachTaskSessionHistoryRequest {
	s.PageSize = &v
	return s
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
	EndTime       *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	PauseDuration *int64  `json:"pauseDuration,omitempty" xml:"pauseDuration,omitempty"`
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
	Total     *int32  `json:"total,omitempty" xml:"total,omitempty"`
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

func (s *GetAICoachTaskSessionHistoryResponseBody) SetPauseDuration(v int64) *GetAICoachTaskSessionHistoryResponseBody {
	s.PauseDuration = &v
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

func (s *GetAICoachTaskSessionHistoryResponseBody) SetTotal(v int32) *GetAICoachTaskSessionHistoryResponseBody {
	s.Total = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetUid(v string) *GetAICoachTaskSessionHistoryResponseBody {
	s.Uid = &v
	return s
}

type GetAICoachTaskSessionHistoryResponseBodyConversationList struct {
	AudioUrl           *string `json:"audioUrl,omitempty" xml:"audioUrl,omitempty"`
	EvaluationFeedback *string `json:"evaluationFeedback,omitempty" xml:"evaluationFeedback,omitempty"`
	// example:
	//
	// {}
	EvaluationResult *string `json:"evaluationResult,omitempty" xml:"evaluationResult,omitempty"`
	Message          *string `json:"message,omitempty" xml:"message,omitempty"`
	RecordId         *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
	Role             *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GetAICoachTaskSessionHistoryResponseBodyConversationList) String() string {
	return tea.Prettify(s)
}

func (s GetAICoachTaskSessionHistoryResponseBodyConversationList) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) SetAudioUrl(v string) *GetAICoachTaskSessionHistoryResponseBodyConversationList {
	s.AudioUrl = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) SetEvaluationFeedback(v string) *GetAICoachTaskSessionHistoryResponseBodyConversationList {
	s.EvaluationFeedback = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) SetEvaluationResult(v string) *GetAICoachTaskSessionHistoryResponseBodyConversationList {
	s.EvaluationResult = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) SetMessage(v string) *GetAICoachTaskSessionHistoryResponseBodyConversationList {
	s.Message = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) SetRecordId(v string) *GetAICoachTaskSessionHistoryResponseBodyConversationList {
	s.RecordId = &v
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
	EndTime          *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	EvaluationRating *string `json:"evaluationRating,omitempty" xml:"evaluationRating,omitempty"`
	// example:
	//
	// {}
	EvaluationResult *string `json:"evaluationResult,omitempty" xml:"evaluationResult,omitempty"`
	Feedback         *bool   `json:"feedback,omitempty" xml:"feedback,omitempty"`
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

func (s *GetAICoachTaskSessionReportResponseBody) SetEvaluationRating(v string) *GetAICoachTaskSessionReportResponseBody {
	s.EvaluationRating = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetEvaluationResult(v string) *GetAICoachTaskSessionReportResponseBody {
	s.EvaluationResult = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetFeedback(v bool) *GetAICoachTaskSessionReportResponseBody {
	s.Feedback = &v
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
	FileType   *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	UploadType *int32  `json:"uploadType,omitempty" xml:"uploadType,omitempty"`
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

func (s *GetOssUploadTokenRequest) SetUploadType(v int32) *GetOssUploadTokenRequest {
	s.UploadType = &v
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

type ListAICoachScriptPageRequest struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	Type   *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAICoachScriptPageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAICoachScriptPageRequest) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageRequest) SetName(v string) *ListAICoachScriptPageRequest {
	s.Name = &v
	return s
}

func (s *ListAICoachScriptPageRequest) SetPageNumber(v int32) *ListAICoachScriptPageRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAICoachScriptPageRequest) SetPageSize(v int32) *ListAICoachScriptPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListAICoachScriptPageRequest) SetStatus(v int32) *ListAICoachScriptPageRequest {
	s.Status = &v
	return s
}

func (s *ListAICoachScriptPageRequest) SetType(v int32) *ListAICoachScriptPageRequest {
	s.Type = &v
	return s
}

type ListAICoachScriptPageResponseBody struct {
	// example:
	//
	// PARAM_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	ErrorMessage *string                                  `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	List         []*ListAICoachScriptPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 86A90C40-D1AB-50DA-A4B1-0D545F80F2FE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAICoachScriptPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAICoachScriptPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBody) SetErrorCode(v string) *ListAICoachScriptPageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAICoachScriptPageResponseBody) SetErrorMessage(v string) *ListAICoachScriptPageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAICoachScriptPageResponseBody) SetList(v []*ListAICoachScriptPageResponseBodyList) *ListAICoachScriptPageResponseBody {
	s.List = v
	return s
}

func (s *ListAICoachScriptPageResponseBody) SetRequestId(v string) *ListAICoachScriptPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAICoachScriptPageResponseBody) SetSuccess(v bool) *ListAICoachScriptPageResponseBody {
	s.Success = &v
	return s
}

func (s *ListAICoachScriptPageResponseBody) SetTotal(v int32) *ListAICoachScriptPageResponseBody {
	s.Total = &v
	return s
}

type ListAICoachScriptPageResponseBodyList struct {
	AppendQuestionFlag *string                                                `json:"appendQuestionFlag,omitempty" xml:"appendQuestionFlag,omitempty"`
	AssessmentScope    *string                                                `json:"assessmentScope,omitempty" xml:"assessmentScope,omitempty"`
	ClosingRemarks     *string                                                `json:"closingRemarks,omitempty" xml:"closingRemarks,omitempty"`
	CompleteStrategy   *ListAICoachScriptPageResponseBodyListCompleteStrategy `json:"completeStrategy,omitempty" xml:"completeStrategy,omitempty" type:"Struct"`
	// example:
	//
	// https://oss-ata.alibaba.com/front/live/banner1.png
	CoverUrl           *string            `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	DialogueTextFlag   *bool              `json:"dialogueTextFlag,omitempty" xml:"dialogueTextFlag,omitempty"`
	DialogueTipFlag    *bool              `json:"dialogueTipFlag,omitempty" xml:"dialogueTipFlag,omitempty"`
	EvaluateReportFlag *bool              `json:"evaluateReportFlag,omitempty" xml:"evaluateReportFlag,omitempty"`
	Expressiveness     map[string]*string `json:"expressiveness,omitempty" xml:"expressiveness,omitempty"`
	GifDynamicUrl      *string            `json:"gifDynamicUrl,omitempty" xml:"gifDynamicUrl,omitempty"`
	GifStaticUrl       *string            `json:"gifStaticUrl,omitempty" xml:"gifStaticUrl,omitempty"`
	// example:
	//
	// 2024-12-25 14:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-12-25 14:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// student
	Initiator *string `json:"initiator,omitempty" xml:"initiator,omitempty"`
	// example:
	//
	// 4
	InteractionType *string `json:"interactionType,omitempty" xml:"interactionType,omitempty"`
	Introduce       *string `json:"introduce,omitempty" xml:"introduce,omitempty"`
	// example:
	//
	// prod-ydsf
	Name               *string                                                    `json:"name,omitempty" xml:"name,omitempty"`
	OpeningRemarks     *string                                                    `json:"openingRemarks,omitempty" xml:"openingRemarks,omitempty"`
	OrderAckFlag       *bool                                                      `json:"orderAckFlag,omitempty" xml:"orderAckFlag,omitempty"`
	SampleDialogueList []*ListAICoachScriptPageResponseBodyListSampleDialogueList `json:"sampleDialogueList,omitempty" xml:"sampleDialogueList,omitempty" type:"Repeated"`
	ScoreConfig        *ListAICoachScriptPageResponseBodyListScoreConfig          `json:"scoreConfig,omitempty" xml:"scoreConfig,omitempty" type:"Struct"`
	// example:
	//
	// 1
	ScriptRecordId     *string `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
	SparringTipContent *string `json:"sparringTipContent,omitempty" xml:"sparringTipContent,omitempty"`
	SparringTipTitle   *string `json:"sparringTipTitle,omitempty" xml:"sparringTipTitle,omitempty"`
	// example:
	//
	// 1
	Status               *int32                                        `json:"status,omitempty" xml:"status,omitempty"`
	StudentThinkTimeFlag *bool                                         `json:"studentThinkTimeFlag,omitempty" xml:"studentThinkTimeFlag,omitempty"`
	Type                 *int32                                        `json:"type,omitempty" xml:"type,omitempty"`
	Weights              *ListAICoachScriptPageResponseBodyListWeights `json:"weights,omitempty" xml:"weights,omitempty" type:"Struct"`
}

func (s ListAICoachScriptPageResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyList) SetAppendQuestionFlag(v string) *ListAICoachScriptPageResponseBodyList {
	s.AppendQuestionFlag = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetAssessmentScope(v string) *ListAICoachScriptPageResponseBodyList {
	s.AssessmentScope = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetClosingRemarks(v string) *ListAICoachScriptPageResponseBodyList {
	s.ClosingRemarks = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetCompleteStrategy(v *ListAICoachScriptPageResponseBodyListCompleteStrategy) *ListAICoachScriptPageResponseBodyList {
	s.CompleteStrategy = v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetCoverUrl(v string) *ListAICoachScriptPageResponseBodyList {
	s.CoverUrl = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetDialogueTextFlag(v bool) *ListAICoachScriptPageResponseBodyList {
	s.DialogueTextFlag = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetDialogueTipFlag(v bool) *ListAICoachScriptPageResponseBodyList {
	s.DialogueTipFlag = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetEvaluateReportFlag(v bool) *ListAICoachScriptPageResponseBodyList {
	s.EvaluateReportFlag = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetExpressiveness(v map[string]*string) *ListAICoachScriptPageResponseBodyList {
	s.Expressiveness = v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetGifDynamicUrl(v string) *ListAICoachScriptPageResponseBodyList {
	s.GifDynamicUrl = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetGifStaticUrl(v string) *ListAICoachScriptPageResponseBodyList {
	s.GifStaticUrl = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetGmtCreate(v string) *ListAICoachScriptPageResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetGmtModified(v string) *ListAICoachScriptPageResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetInitiator(v string) *ListAICoachScriptPageResponseBodyList {
	s.Initiator = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetInteractionType(v string) *ListAICoachScriptPageResponseBodyList {
	s.InteractionType = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetIntroduce(v string) *ListAICoachScriptPageResponseBodyList {
	s.Introduce = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetName(v string) *ListAICoachScriptPageResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetOpeningRemarks(v string) *ListAICoachScriptPageResponseBodyList {
	s.OpeningRemarks = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetOrderAckFlag(v bool) *ListAICoachScriptPageResponseBodyList {
	s.OrderAckFlag = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetSampleDialogueList(v []*ListAICoachScriptPageResponseBodyListSampleDialogueList) *ListAICoachScriptPageResponseBodyList {
	s.SampleDialogueList = v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetScoreConfig(v *ListAICoachScriptPageResponseBodyListScoreConfig) *ListAICoachScriptPageResponseBodyList {
	s.ScoreConfig = v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetScriptRecordId(v string) *ListAICoachScriptPageResponseBodyList {
	s.ScriptRecordId = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetSparringTipContent(v string) *ListAICoachScriptPageResponseBodyList {
	s.SparringTipContent = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetSparringTipTitle(v string) *ListAICoachScriptPageResponseBodyList {
	s.SparringTipTitle = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetStatus(v int32) *ListAICoachScriptPageResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetStudentThinkTimeFlag(v bool) *ListAICoachScriptPageResponseBodyList {
	s.StudentThinkTimeFlag = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetType(v int32) *ListAICoachScriptPageResponseBodyList {
	s.Type = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyList) SetWeights(v *ListAICoachScriptPageResponseBodyListWeights) *ListAICoachScriptPageResponseBodyList {
	s.Weights = v
	return s
}

type ListAICoachScriptPageResponseBodyListCompleteStrategy struct {
	// example:
	//
	// true
	ClickCompleteAutoEnd *bool `json:"clickCompleteAutoEnd,omitempty" xml:"clickCompleteAutoEnd,omitempty"`
	// example:
	//
	// 75
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// false
	FullCoverageAutoEnd *bool `json:"fullCoverageAutoEnd,omitempty" xml:"fullCoverageAutoEnd,omitempty"`
}

func (s ListAICoachScriptPageResponseBodyListCompleteStrategy) String() string {
	return tea.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyListCompleteStrategy) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyListCompleteStrategy) SetClickCompleteAutoEnd(v bool) *ListAICoachScriptPageResponseBodyListCompleteStrategy {
	s.ClickCompleteAutoEnd = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCompleteStrategy) SetDuration(v int32) *ListAICoachScriptPageResponseBodyListCompleteStrategy {
	s.Duration = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListCompleteStrategy) SetFullCoverageAutoEnd(v bool) *ListAICoachScriptPageResponseBodyListCompleteStrategy {
	s.FullCoverageAutoEnd = &v
	return s
}

type ListAICoachScriptPageResponseBodyListSampleDialogueList struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// student
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ListAICoachScriptPageResponseBodyListSampleDialogueList) String() string {
	return tea.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyListSampleDialogueList) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyListSampleDialogueList) SetMessage(v string) *ListAICoachScriptPageResponseBodyListSampleDialogueList {
	s.Message = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListSampleDialogueList) SetRole(v string) *ListAICoachScriptPageResponseBodyListSampleDialogueList {
	s.Role = &v
	return s
}

type ListAICoachScriptPageResponseBodyListScoreConfig struct {
	Enabled   *bool  `json:"enabled,omitempty" xml:"enabled,omitempty"`
	PassScore *int32 `json:"passScore,omitempty" xml:"passScore,omitempty"`
}

func (s ListAICoachScriptPageResponseBodyListScoreConfig) String() string {
	return tea.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyListScoreConfig) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfig) SetEnabled(v bool) *ListAICoachScriptPageResponseBodyListScoreConfig {
	s.Enabled = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListScoreConfig) SetPassScore(v int32) *ListAICoachScriptPageResponseBodyListScoreConfig {
	s.PassScore = &v
	return s
}

type ListAICoachScriptPageResponseBodyListWeights struct {
	// example:
	//
	// 50
	AssessmentPoint        *int32 `json:"assessmentPoint,omitempty" xml:"assessmentPoint,omitempty"`
	AssessmentPointEnabled *bool  `json:"assessmentPointEnabled,omitempty" xml:"assessmentPointEnabled,omitempty"`
	// example:
	//
	// 30
	Expressiveness            *int32 `json:"expressiveness,omitempty" xml:"expressiveness,omitempty"`
	ExpressivenessEnabled     *bool  `json:"expressivenessEnabled,omitempty" xml:"expressivenessEnabled,omitempty"`
	PointDeductionRule        *int32 `json:"pointDeductionRule,omitempty" xml:"pointDeductionRule,omitempty"`
	PointDeductionRuleEnabled *bool  `json:"pointDeductionRuleEnabled,omitempty" xml:"pointDeductionRuleEnabled,omitempty"`
	// example:
	//
	// 20
	Standard *int32 `json:"standard,omitempty" xml:"standard,omitempty"`
	// example:
	//
	// true
	StandardEnabled *bool `json:"standardEnabled,omitempty" xml:"standardEnabled,omitempty"`
}

func (s ListAICoachScriptPageResponseBodyListWeights) String() string {
	return tea.Prettify(s)
}

func (s ListAICoachScriptPageResponseBodyListWeights) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetAssessmentPoint(v int32) *ListAICoachScriptPageResponseBodyListWeights {
	s.AssessmentPoint = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetAssessmentPointEnabled(v bool) *ListAICoachScriptPageResponseBodyListWeights {
	s.AssessmentPointEnabled = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetExpressiveness(v int32) *ListAICoachScriptPageResponseBodyListWeights {
	s.Expressiveness = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetExpressivenessEnabled(v bool) *ListAICoachScriptPageResponseBodyListWeights {
	s.ExpressivenessEnabled = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetPointDeductionRule(v int32) *ListAICoachScriptPageResponseBodyListWeights {
	s.PointDeductionRule = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetPointDeductionRuleEnabled(v bool) *ListAICoachScriptPageResponseBodyListWeights {
	s.PointDeductionRuleEnabled = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetStandard(v int32) *ListAICoachScriptPageResponseBodyListWeights {
	s.Standard = &v
	return s
}

func (s *ListAICoachScriptPageResponseBodyListWeights) SetStandardEnabled(v bool) *ListAICoachScriptPageResponseBodyListWeights {
	s.StandardEnabled = &v
	return s
}

type ListAICoachScriptPageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAICoachScriptPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAICoachScriptPageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAICoachScriptPageResponse) GoString() string {
	return s.String()
}

func (s *ListAICoachScriptPageResponse) SetHeaders(v map[string]*string) *ListAICoachScriptPageResponse {
	s.Headers = v
	return s
}

func (s *ListAICoachScriptPageResponse) SetStatusCode(v int32) *ListAICoachScriptPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAICoachScriptPageResponse) SetBody(v *ListAICoachScriptPageResponseBody) *ListAICoachScriptPageResponse {
	s.Body = v
	return s
}

type ListAICoachTaskPageRequest struct {
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize  *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
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

func (s *ListAICoachTaskPageRequest) SetEndTime(v string) *ListAICoachTaskPageRequest {
	s.EndTime = &v
	return s
}

func (s *ListAICoachTaskPageRequest) SetPageNumber(v int32) *ListAICoachTaskPageRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAICoachTaskPageRequest) SetPageSize(v int32) *ListAICoachTaskPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListAICoachTaskPageRequest) SetStartTime(v string) *ListAICoachTaskPageRequest {
	s.StartTime = &v
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
	GmtCreate  *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
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

func (s *ListAICoachTaskPageResponseBodyTaskList) SetGmtCreate(v string) *ListAICoachTaskPageResponseBodyTaskList {
	s.GmtCreate = &v
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

type ListAgentsRequest struct {
	// example:
	//
	// 840016700254633984
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// text
	AgentScene *string `json:"agentScene,omitempty" xml:"agentScene,omitempty"`
	// example:
	//
	// SYSTEM
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
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
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListAgentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentsRequest) SetAgentId(v string) *ListAgentsRequest {
	s.AgentId = &v
	return s
}

func (s *ListAgentsRequest) SetAgentScene(v string) *ListAgentsRequest {
	s.AgentScene = &v
	return s
}

func (s *ListAgentsRequest) SetOwner(v string) *ListAgentsRequest {
	s.Owner = &v
	return s
}

func (s *ListAgentsRequest) SetPageNumber(v int32) *ListAgentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAgentsRequest) SetPageSize(v int32) *ListAgentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentsRequest) SetStatus(v int32) *ListAgentsRequest {
	s.Status = &v
	return s
}

type ListAgentsResponseBody struct {
	List []*ListAgentsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 86A90C40-D1AB-50DA-A4B1-0D545F80F2FE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAgentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAgentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBody) SetList(v []*ListAgentsResponseBodyList) *ListAgentsResponseBody {
	s.List = v
	return s
}

func (s *ListAgentsResponseBody) SetRequestId(v string) *ListAgentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentsResponseBody) SetSuccess(v bool) *ListAgentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgentsResponseBody) SetTotal(v int32) *ListAgentsResponseBody {
	s.Total = &v
	return s
}

type ListAgentsResponseBodyList struct {
	// example:
	//
	// 840016700254633984
	AgentId   *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AgentName *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
	// example:
	//
	// text
	AgentScene            *string `json:"agentScene,omitempty" xml:"agentScene,omitempty"`
	CharactersDescription *string `json:"charactersDescription,omitempty" xml:"charactersDescription,omitempty"`
	// example:
	//
	// 1
	EnableInteraction *int32 `json:"enableInteraction,omitempty" xml:"enableInteraction,omitempty"`
	// example:
	//
	// Car
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
	// example:
	//
	// true
	OnlineSearch *bool `json:"onlineSearch,omitempty" xml:"onlineSearch,omitempty"`
	// example:
	//
	// SYSTEM
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// http
	//
	// ;//www.abc.com/111.mp4
	ReferenceUrl *string `json:"referenceUrl,omitempty" xml:"referenceUrl,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// RED_BOOK
	TextStyle *string `json:"textStyle,omitempty" xml:"textStyle,omitempty"`
	// example:
	//
	// Seller
	Viewer *string `json:"viewer,omitempty" xml:"viewer,omitempty"`
}

func (s ListAgentsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListAgentsResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBodyList) SetAgentId(v string) *ListAgentsResponseBodyList {
	s.AgentId = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetAgentName(v string) *ListAgentsResponseBodyList {
	s.AgentName = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetAgentScene(v string) *ListAgentsResponseBodyList {
	s.AgentScene = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetCharactersDescription(v string) *ListAgentsResponseBodyList {
	s.CharactersDescription = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetEnableInteraction(v int32) *ListAgentsResponseBodyList {
	s.EnableInteraction = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetIndustry(v string) *ListAgentsResponseBodyList {
	s.Industry = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetOnlineSearch(v bool) *ListAgentsResponseBodyList {
	s.OnlineSearch = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetOwner(v string) *ListAgentsResponseBodyList {
	s.Owner = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetReferenceUrl(v string) *ListAgentsResponseBodyList {
	s.ReferenceUrl = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetStatus(v int32) *ListAgentsResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetTextStyle(v string) *ListAgentsResponseBodyList {
	s.TextStyle = &v
	return s
}

func (s *ListAgentsResponseBodyList) SetViewer(v string) *ListAgentsResponseBodyList {
	s.Viewer = &v
	return s
}

type ListAgentsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAgentsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentsResponse) SetHeaders(v map[string]*string) *ListAgentsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentsResponse) SetStatusCode(v int32) *ListAgentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentsResponse) SetBody(v *ListAgentsResponseBody) *ListAgentsResponse {
	s.Body = v
	return s
}

type ListAnchorRequest struct {
	AnchorCategory *string `json:"anchorCategory,omitempty" xml:"anchorCategory,omitempty"`
	AnchorId       *string `json:"anchorId,omitempty" xml:"anchorId,omitempty"`
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

func (s *ListAnchorRequest) SetAnchorCategory(v string) *ListAnchorRequest {
	s.AnchorCategory = &v
	return s
}

func (s *ListAnchorRequest) SetAnchorId(v string) *ListAnchorRequest {
	s.AnchorId = &v
	return s
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

type ListKnowledgeBaseRequest struct {
	// example:
	//
	// "186432649"
	KnowledgeBaseId *string `json:"knowledgeBaseId,omitempty" xml:"knowledgeBaseId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListKnowledgeBaseRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseRequest) SetKnowledgeBaseId(v string) *ListKnowledgeBaseRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeBaseRequest) SetPageNumber(v int32) *ListKnowledgeBaseRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKnowledgeBaseRequest) SetPageSize(v int32) *ListKnowledgeBaseRequest {
	s.PageSize = &v
	return s
}

type ListKnowledgeBaseResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KnowledgeBaseListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgeBaseResponse) String() string {
	return tea.Prettify(s)
}

func (s ListKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseResponse) SetHeaders(v map[string]*string) *ListKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgeBaseResponse) SetStatusCode(v int32) *ListKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgeBaseResponse) SetBody(v *KnowledgeBaseListResult) *ListKnowledgeBaseResponse {
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
	UseScene      *string `json:"useScene,omitempty" xml:"useScene,omitempty"`
	VoiceLanguage *string `json:"voiceLanguage,omitempty" xml:"voiceLanguage,omitempty"`
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

func (s *ListVoiceModelsRequest) SetVoiceLanguage(v string) *ListVoiceModelsRequest {
	s.VoiceLanguage = &v
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
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResSpecType    *string `json:"resSpecType,omitempty" xml:"resSpecType,omitempty"`
	ScaleType      *string `json:"scaleType,omitempty" xml:"scaleType,omitempty"`
	ScriptModelTag *string `json:"scriptModelTag,omitempty" xml:"scriptModelTag,omitempty"`
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

func (s *QueryAvatarProjectResponseBody) SetScriptModelTag(v string) *QueryAvatarProjectResponseBody {
	s.ScriptModelTag = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetStatus(v string) *QueryAvatarProjectResponseBody {
	s.Status = &v
	return s
}

type QueryAvatarProjectResponseBodyFrames struct {
	Index       *int32                                           `json:"index,omitempty" xml:"index,omitempty"`
	Layers      []*QueryAvatarProjectResponseBodyFramesLayers    `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	VideoScript *QueryAvatarProjectResponseBodyFramesVideoScript `json:"videoScript,omitempty" xml:"videoScript,omitempty" type:"Struct"`
}

func (s QueryAvatarProjectResponseBodyFrames) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarProjectResponseBodyFrames) GoString() string {
	return s.String()
}

func (s *QueryAvatarProjectResponseBodyFrames) SetIndex(v int32) *QueryAvatarProjectResponseBodyFrames {
	s.Index = &v
	return s
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
	Index     *int32                                              `json:"index,omitempty" xml:"index,omitempty"`
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

func (s *QueryAvatarProjectResponseBodyFramesLayers) SetIndex(v int32) *QueryAvatarProjectResponseBodyFramesLayers {
	s.Index = &v
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
	Emotion         *string `json:"emotion,omitempty" xml:"emotion,omitempty"`
	PitchRate       *string `json:"pitchRate,omitempty" xml:"pitchRate,omitempty"`
	SpeedRate       *string `json:"speedRate,omitempty" xml:"speedRate,omitempty"`
	TextContent     *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
	VoiceLanguage   *string `json:"voiceLanguage,omitempty" xml:"voiceLanguage,omitempty"`
	VoiceTemplateId *string `json:"voiceTemplateId,omitempty" xml:"voiceTemplateId,omitempty"`
	Volume          *int32  `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s QueryAvatarProjectResponseBodyFramesVideoScript) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarProjectResponseBodyFramesVideoScript) GoString() string {
	return s.String()
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) SetEmotion(v string) *QueryAvatarProjectResponseBodyFramesVideoScript {
	s.Emotion = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) SetPitchRate(v string) *QueryAvatarProjectResponseBodyFramesVideoScript {
	s.PitchRate = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) SetSpeedRate(v string) *QueryAvatarProjectResponseBodyFramesVideoScript {
	s.SpeedRate = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) SetTextContent(v string) *QueryAvatarProjectResponseBodyFramesVideoScript {
	s.TextContent = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) SetVoiceLanguage(v string) *QueryAvatarProjectResponseBodyFramesVideoScript {
	s.VoiceLanguage = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) SetVoiceTemplateId(v string) *QueryAvatarProjectResponseBodyFramesVideoScript {
	s.VoiceTemplateId = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) SetVolume(v int32) *QueryAvatarProjectResponseBodyFramesVideoScript {
	s.Volume = &v
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

type QueryImageToVideoTaskRequest struct {
	// example:
	//
	// 868125994191405056
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QueryImageToVideoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryImageToVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryImageToVideoTaskRequest) SetTaskId(v string) *QueryImageToVideoTaskRequest {
	s.TaskId = &v
	return s
}

type QueryImageToVideoTaskResponseBody struct {
	// example:
	//
	// None
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// https://xxx/xxx.mp4
	OriginUrl *string `json:"originUrl,omitempty" xml:"originUrl,omitempty"`
	// example:
	//
	// CC2967CA-0114-57E0-A0CF-7DEEEDAB953D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 868125994191405056
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QueryImageToVideoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryImageToVideoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryImageToVideoTaskResponseBody) SetMessage(v string) *QueryImageToVideoTaskResponseBody {
	s.Message = &v
	return s
}

func (s *QueryImageToVideoTaskResponseBody) SetOriginUrl(v string) *QueryImageToVideoTaskResponseBody {
	s.OriginUrl = &v
	return s
}

func (s *QueryImageToVideoTaskResponseBody) SetRequestId(v string) *QueryImageToVideoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryImageToVideoTaskResponseBody) SetStatus(v int32) *QueryImageToVideoTaskResponseBody {
	s.Status = &v
	return s
}

func (s *QueryImageToVideoTaskResponseBody) SetSuccess(v bool) *QueryImageToVideoTaskResponseBody {
	s.Success = &v
	return s
}

func (s *QueryImageToVideoTaskResponseBody) SetTaskId(v string) *QueryImageToVideoTaskResponseBody {
	s.TaskId = &v
	return s
}

type QueryImageToVideoTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryImageToVideoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryImageToVideoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryImageToVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryImageToVideoTaskResponse) SetHeaders(v map[string]*string) *QueryImageToVideoTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryImageToVideoTaskResponse) SetStatusCode(v int32) *QueryImageToVideoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryImageToVideoTaskResponse) SetBody(v *QueryImageToVideoTaskResponseBody) *QueryImageToVideoTaskResponse {
	s.Body = v
	return s
}

type QueryIndividuationTextTaskRequest struct {
	// example:
	//
	// 829682927337963520
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QueryIndividuationTextTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryIndividuationTextTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryIndividuationTextTaskRequest) SetTaskId(v string) *QueryIndividuationTextTaskRequest {
	s.TaskId = &v
	return s
}

type QueryIndividuationTextTaskResponseBody struct {
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 56AC346B-AF40-5E4F-AFFE-FD8BA5E6FB3A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0
	Status   *int32                                            `json:"status,omitempty" xml:"status,omitempty"`
	TextList []*QueryIndividuationTextTaskResponseBodyTextList `json:"textList,omitempty" xml:"textList,omitempty" type:"Repeated"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s QueryIndividuationTextTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryIndividuationTextTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIndividuationTextTaskResponseBody) SetCreateTime(v string) *QueryIndividuationTextTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *QueryIndividuationTextTaskResponseBody) SetRequestId(v string) *QueryIndividuationTextTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryIndividuationTextTaskResponseBody) SetStatus(v int32) *QueryIndividuationTextTaskResponseBody {
	s.Status = &v
	return s
}

func (s *QueryIndividuationTextTaskResponseBody) SetTextList(v []*QueryIndividuationTextTaskResponseBodyTextList) *QueryIndividuationTextTaskResponseBody {
	s.TextList = v
	return s
}

func (s *QueryIndividuationTextTaskResponseBody) SetUpdateTime(v string) *QueryIndividuationTextTaskResponseBody {
	s.UpdateTime = &v
	return s
}

type QueryIndividuationTextTaskResponseBodyTextList struct {
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2761
	TextId *string `json:"textId,omitempty" xml:"textId,omitempty"`
	// example:
	//
	// 11
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryIndividuationTextTaskResponseBodyTextList) String() string {
	return tea.Prettify(s)
}

func (s QueryIndividuationTextTaskResponseBodyTextList) GoString() string {
	return s.String()
}

func (s *QueryIndividuationTextTaskResponseBodyTextList) SetStatus(v int32) *QueryIndividuationTextTaskResponseBodyTextList {
	s.Status = &v
	return s
}

func (s *QueryIndividuationTextTaskResponseBodyTextList) SetTextId(v string) *QueryIndividuationTextTaskResponseBodyTextList {
	s.TextId = &v
	return s
}

func (s *QueryIndividuationTextTaskResponseBodyTextList) SetUserId(v string) *QueryIndividuationTextTaskResponseBodyTextList {
	s.UserId = &v
	return s
}

type QueryIndividuationTextTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryIndividuationTextTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryIndividuationTextTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryIndividuationTextTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryIndividuationTextTaskResponse) SetHeaders(v map[string]*string) *QueryIndividuationTextTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryIndividuationTextTaskResponse) SetStatusCode(v int32) *QueryIndividuationTextTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryIndividuationTextTaskResponse) SetBody(v *QueryIndividuationTextTaskResponseBody) *QueryIndividuationTextTaskResponse {
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
	AgentId   *string                           `json:"agentId,omitempty" xml:"agentId,omitempty"`
	BitRate   *string                           `json:"bitRate,omitempty" xml:"bitRate,omitempty"`
	FrameRate *string                           `json:"frameRate,omitempty" xml:"frameRate,omitempty"`
	Frames    []*SaveAvatarProjectRequestFrames `json:"frames,omitempty" xml:"frames,omitempty" type:"Repeated"`
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
	Resolution  *string `json:"resolution,omitempty" xml:"resolution,omitempty"`
	// example:
	//
	// 9:16
	ScaleType           *string `json:"scaleType,omitempty" xml:"scaleType,omitempty"`
	ScriptModelTag      *string `json:"scriptModelTag,omitempty" xml:"scriptModelTag,omitempty"`
	SynchronizedDisplay *string `json:"synchronizedDisplay,omitempty" xml:"synchronizedDisplay,omitempty"`
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

func (s *SaveAvatarProjectRequest) SetBitRate(v string) *SaveAvatarProjectRequest {
	s.BitRate = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetFrameRate(v string) *SaveAvatarProjectRequest {
	s.FrameRate = &v
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

func (s *SaveAvatarProjectRequest) SetResolution(v string) *SaveAvatarProjectRequest {
	s.Resolution = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetScaleType(v string) *SaveAvatarProjectRequest {
	s.ScaleType = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetScriptModelTag(v string) *SaveAvatarProjectRequest {
	s.ScriptModelTag = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetSynchronizedDisplay(v string) *SaveAvatarProjectRequest {
	s.SynchronizedDisplay = &v
	return s
}

type SaveAvatarProjectRequestFrames struct {
	Index       *int32                                     `json:"index,omitempty" xml:"index,omitempty"`
	Layers      []*SaveAvatarProjectRequestFramesLayers    `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	VideoScript *SaveAvatarProjectRequestFramesVideoScript `json:"videoScript,omitempty" xml:"videoScript,omitempty" type:"Struct"`
}

func (s SaveAvatarProjectRequestFrames) String() string {
	return tea.Prettify(s)
}

func (s SaveAvatarProjectRequestFrames) GoString() string {
	return s.String()
}

func (s *SaveAvatarProjectRequestFrames) SetIndex(v int32) *SaveAvatarProjectRequestFrames {
	s.Index = &v
	return s
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
	Index    *int32                                        `json:"index,omitempty" xml:"index,omitempty"`
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

func (s *SaveAvatarProjectRequestFramesLayers) SetIndex(v int32) *SaveAvatarProjectRequestFramesLayers {
	s.Index = &v
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
	Emotion   *string `json:"emotion,omitempty" xml:"emotion,omitempty"`
	PitchRate *string `json:"pitchRate,omitempty" xml:"pitchRate,omitempty"`
	// example:
	//
	// 1.0
	SpeedRate     *string `json:"speedRate,omitempty" xml:"speedRate,omitempty"`
	TextContent   *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
	VoiceLanguage *string `json:"voiceLanguage,omitempty" xml:"voiceLanguage,omitempty"`
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

func (s *SaveAvatarProjectRequestFramesVideoScript) SetEmotion(v string) *SaveAvatarProjectRequestFramesVideoScript {
	s.Emotion = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesVideoScript) SetPitchRate(v string) *SaveAvatarProjectRequestFramesVideoScript {
	s.PitchRate = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesVideoScript) SetSpeedRate(v string) *SaveAvatarProjectRequestFramesVideoScript {
	s.SpeedRate = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesVideoScript) SetTextContent(v string) *SaveAvatarProjectRequestFramesVideoScript {
	s.TextContent = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesVideoScript) SetVoiceLanguage(v string) *SaveAvatarProjectRequestFramesVideoScript {
	s.VoiceLanguage = &v
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
	AliyunUid *string `json:"aliyunUid,omitempty" xml:"aliyunUid,omitempty"`
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

func (s *SelectResourceResponseBody) SetAliyunUid(v string) *SelectResourceResponseBody {
	s.AliyunUid = &v
	return s
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

type SendSdkMessageRequest struct {
	// example:
	//
	// {}
	Data   *string `json:"data,omitempty" xml:"data,omitempty"`
	Header *string `json:"header,omitempty" xml:"header,omitempty"`
	// example:
	//
	// avatar
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// example:
	//
	// getProject
	OperationName *string `json:"operationName,omitempty" xml:"operationName,omitempty"`
	// example:
	//
	// 123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SendSdkMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendSdkMessageRequest) GoString() string {
	return s.String()
}

func (s *SendSdkMessageRequest) SetData(v string) *SendSdkMessageRequest {
	s.Data = &v
	return s
}

func (s *SendSdkMessageRequest) SetHeader(v string) *SendSdkMessageRequest {
	s.Header = &v
	return s
}

func (s *SendSdkMessageRequest) SetModuleName(v string) *SendSdkMessageRequest {
	s.ModuleName = &v
	return s
}

func (s *SendSdkMessageRequest) SetOperationName(v string) *SendSdkMessageRequest {
	s.OperationName = &v
	return s
}

func (s *SendSdkMessageRequest) SetUserId(v string) *SendSdkMessageRequest {
	s.UserId = &v
	return s
}

type SendSdkMessageResponseBody struct {
	// example:
	//
	// {}
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// system-01
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 86A90C40-D1AB-50DA-A4B1-0D545F80F2FE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendSdkMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendSdkMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendSdkMessageResponseBody) SetData(v string) *SendSdkMessageResponseBody {
	s.Data = &v
	return s
}

func (s *SendSdkMessageResponseBody) SetErrorCode(v string) *SendSdkMessageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SendSdkMessageResponseBody) SetErrorMessage(v string) *SendSdkMessageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SendSdkMessageResponseBody) SetRequestId(v string) *SendSdkMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendSdkMessageResponseBody) SetSuccess(v bool) *SendSdkMessageResponseBody {
	s.Success = &v
	return s
}

type SendSdkMessageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendSdkMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendSdkMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendSdkMessageResponse) GoString() string {
	return s.String()
}

func (s *SendSdkMessageResponse) SetHeaders(v map[string]*string) *SendSdkMessageResponse {
	s.Headers = v
	return s
}

func (s *SendSdkMessageResponse) SetStatusCode(v int32) *SendSdkMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendSdkMessageResponse) SetBody(v *SendSdkMessageResponseBody) *SendSdkMessageResponse {
	s.Body = v
	return s
}

type SendSdkStreamMessageRequest struct {
	// example:
	//
	// {"test":""}
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// header
	//
	// example:
	//
	// {}
	Header *string `json:"header,omitempty" xml:"header,omitempty"`
	// example:
	//
	// avatar
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// example:
	//
	// GetProject
	OperationName *string `json:"operationName,omitempty" xml:"operationName,omitempty"`
	// example:
	//
	// 123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SendSdkStreamMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendSdkStreamMessageRequest) GoString() string {
	return s.String()
}

func (s *SendSdkStreamMessageRequest) SetData(v string) *SendSdkStreamMessageRequest {
	s.Data = &v
	return s
}

func (s *SendSdkStreamMessageRequest) SetHeader(v string) *SendSdkStreamMessageRequest {
	s.Header = &v
	return s
}

func (s *SendSdkStreamMessageRequest) SetModuleName(v string) *SendSdkStreamMessageRequest {
	s.ModuleName = &v
	return s
}

func (s *SendSdkStreamMessageRequest) SetOperationName(v string) *SendSdkStreamMessageRequest {
	s.OperationName = &v
	return s
}

func (s *SendSdkStreamMessageRequest) SetUserId(v string) *SendSdkStreamMessageRequest {
	s.UserId = &v
	return s
}

type SendSdkStreamMessageResponseBody struct {
	// example:
	//
	// {"id":"123"}
	CommonStreamMessage *string `json:"commonStreamMessage,omitempty" xml:"commonStreamMessage,omitempty"`
}

func (s SendSdkStreamMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendSdkStreamMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendSdkStreamMessageResponseBody) SetCommonStreamMessage(v string) *SendSdkStreamMessageResponseBody {
	s.CommonStreamMessage = &v
	return s
}

type SendSdkStreamMessageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendSdkStreamMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendSdkStreamMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendSdkStreamMessageResponse) GoString() string {
	return s.String()
}

func (s *SendSdkStreamMessageResponse) SetHeaders(v map[string]*string) *SendSdkStreamMessageResponse {
	s.Headers = v
	return s
}

func (s *SendSdkStreamMessageResponse) SetStatusCode(v int32) *SendSdkStreamMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendSdkStreamMessageResponse) SetBody(v *SendSdkStreamMessageResponseBody) *SendSdkStreamMessageResponse {
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
	ChannelToken  *string `json:"channelToken,omitempty" xml:"channelToken,omitempty"`
	CustomPushUrl *string `json:"customPushUrl,omitempty" xml:"customPushUrl,omitempty"`
	CustomUserId  *string `json:"customUserId,omitempty" xml:"customUserId,omitempty"`
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

func (s *StartAvatarSessionRequest) SetChannelToken(v string) *StartAvatarSessionRequest {
	s.ChannelToken = &v
	return s
}

func (s *StartAvatarSessionRequest) SetCustomPushUrl(v string) *StartAvatarSessionRequest {
	s.CustomPushUrl = &v
	return s
}

func (s *StartAvatarSessionRequest) SetCustomUserId(v string) *StartAvatarSessionRequest {
	s.CustomUserId = &v
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

type SubmitImageToVideoTaskRequest struct {
	// example:
	//
	// http://xxx/image.png
	ImageUrl  *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	PosPrompt *string `json:"posPrompt,omitempty" xml:"posPrompt,omitempty"`
}

func (s SubmitImageToVideoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitImageToVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitImageToVideoTaskRequest) SetImageUrl(v string) *SubmitImageToVideoTaskRequest {
	s.ImageUrl = &v
	return s
}

func (s *SubmitImageToVideoTaskRequest) SetPosPrompt(v string) *SubmitImageToVideoTaskRequest {
	s.PosPrompt = &v
	return s
}

type SubmitImageToVideoTaskResponseBody struct {
	// example:
	//
	// job added successfully
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 53AED51A-74CE-57CE-B1BF-2703F314EEC8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 868125994191405056
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitImageToVideoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitImageToVideoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitImageToVideoTaskResponseBody) SetMessage(v string) *SubmitImageToVideoTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitImageToVideoTaskResponseBody) SetRequestId(v string) *SubmitImageToVideoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitImageToVideoTaskResponseBody) SetSuccess(v bool) *SubmitImageToVideoTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitImageToVideoTaskResponseBody) SetTaskId(v string) *SubmitImageToVideoTaskResponseBody {
	s.TaskId = &v
	return s
}

type SubmitImageToVideoTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitImageToVideoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitImageToVideoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitImageToVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitImageToVideoTaskResponse) SetHeaders(v map[string]*string) *SubmitImageToVideoTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitImageToVideoTaskResponse) SetStatusCode(v int32) *SubmitImageToVideoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitImageToVideoTaskResponse) SetBody(v *SubmitImageToVideoTaskResponseBody) *SubmitImageToVideoTaskResponse {
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
	AnchorStyleLevel *string `json:"anchorStyleLevel,omitempty" xml:"anchorStyleLevel,omitempty"`
	// example:
	//
	// video/mp4
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// example:
	//
	// 38863
	Id    *string                                           `json:"id,omitempty" xml:"id,omitempty"`
	Mask  *SubmitProjectTaskRequestFramesLayersMaterialMask `json:"mask,omitempty" xml:"mask,omitempty" type:"Struct"`
	Speed *string                                           `json:"speed,omitempty" xml:"speed,omitempty"`
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

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetAnchorStyleLevel(v string) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.AnchorStyleLevel = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetFormat(v string) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.Format = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetId(v string) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.Id = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetMask(v *SubmitProjectTaskRequestFramesLayersMaterialMask) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.Mask = v
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

type SubmitProjectTaskRequestFramesLayersMaterialMask struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SubmitProjectTaskRequestFramesLayersMaterialMask) String() string {
	return tea.Prettify(s)
}

func (s SubmitProjectTaskRequestFramesLayersMaterialMask) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskRequestFramesLayersMaterialMask) SetUrl(v string) *SubmitProjectTaskRequestFramesLayersMaterialMask {
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
	AudioUrl   *string `json:"audioUrl,omitempty" xml:"audioUrl,omitempty"`
	Emotion    *string `json:"emotion,omitempty" xml:"emotion,omitempty"`
	PitchRate  *string `json:"pitchRate,omitempty" xml:"pitchRate,omitempty"`
	SpeechOpen *bool   `json:"speechOpen,omitempty" xml:"speechOpen,omitempty"`
	// example:
	//
	// 2.0
	SpeedRate   *string `json:"speedRate,omitempty" xml:"speedRate,omitempty"`
	TextContent *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
	// example:
	//
	// TEXT
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
	VoiceLanguage *string `json:"voiceLanguage,omitempty" xml:"voiceLanguage,omitempty"`
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

func (s *SubmitProjectTaskRequestFramesVideoScript) SetEmotion(v string) *SubmitProjectTaskRequestFramesVideoScript {
	s.Emotion = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetPitchRate(v string) *SubmitProjectTaskRequestFramesVideoScript {
	s.PitchRate = &v
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

func (s *SubmitProjectTaskRequestFramesVideoScript) SetVoiceLanguage(v string) *SubmitProjectTaskRequestFramesVideoScript {
	s.VoiceLanguage = &v
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
// 批量添加知识文档
//
// @param request - BatchAddDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddDocumentResponse
func (client *Client) BatchAddDocumentWithOptions(knowledgeBaseId *string, request *BatchAddDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchAddDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddDocumentInfos)) {
		body["addDocumentInfos"] = request.AddDocumentInfos
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchAddDocument"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/knowledge-base/" + tea.StringValue(openapiutil.GetEncodeParam(knowledgeBaseId)) + "/documents"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchAddDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量添加知识文档
//
// @param request - BatchAddDocumentRequest
//
// @return BatchAddDocumentResponse
func (client *Client) BatchAddDocument(knowledgeBaseId *string, request *BatchAddDocumentRequest) (_result *BatchAddDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchAddDocumentResponse{}
	_body, _err := client.BatchAddDocumentWithOptions(knowledgeBaseId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量发布剧本任务
//
// @param request - BatchCreateAICoachTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateAICoachTaskResponse
func (client *Client) BatchCreateAICoachTaskWithOptions(request *BatchCreateAICoachTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchCreateAICoachTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ScriptRecordId)) {
		body["scriptRecordId"] = request.ScriptRecordId
	}

	if !tea.BoolValue(util.IsUnset(request.StudentIds)) {
		body["studentIds"] = request.StudentIds
	}

	if !tea.BoolValue(util.IsUnset(request.StudentList)) {
		body["studentList"] = request.StudentList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchCreateAICoachTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/aicoach/batchCreateTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchCreateAICoachTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量发布剧本任务
//
// @param request - BatchCreateAICoachTaskRequest
//
// @return BatchCreateAICoachTaskResponse
func (client *Client) BatchCreateAICoachTask(request *BatchCreateAICoachTaskRequest) (_result *BatchCreateAICoachTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchCreateAICoachTaskResponse{}
	_body, _err := client.BatchCreateAICoachTaskWithOptions(request, headers, runtime)
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
// 批量查询声音复刻任务
//
// @param tmpReq - BatchGetTrainTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetTrainTaskResponse
func (client *Client) BatchGetTrainTaskWithOptions(tmpReq *BatchGetTrainTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchGetTrainTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchGetTrainTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TaskIdList)) {
		request.TaskIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIdList, tea.String("taskIdList"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunMainId)) {
		query["aliyunMainId"] = request.AliyunMainId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIdListShrink)) {
		query["taskIdList"] = request.TaskIdListShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchGetTrainTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/train/task/batchGetTrainTaskInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchGetTrainTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询声音复刻任务
//
// @param request - BatchGetTrainTaskRequest
//
// @return BatchGetTrainTaskResponse
func (client *Client) BatchGetTrainTask(request *BatchGetTrainTaskRequest) (_result *BatchGetTrainTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchGetTrainTaskResponse{}
	_body, _err := client.BatchGetTrainTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询视频切片任务信息
//
// @param tmpReq - BatchGetVideoClipTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetVideoClipTaskResponse
func (client *Client) BatchGetVideoClipTaskWithOptions(tmpReq *BatchGetVideoClipTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchGetVideoClipTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchGetVideoClipTaskShrinkRequest{}
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
		Action:      tea.String("BatchGetVideoClipTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/video/clip/batchGetVideoClipTask"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchGetVideoClipTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询视频切片任务信息
//
// @param request - BatchGetVideoClipTaskRequest
//
// @return BatchGetVideoClipTaskResponse
func (client *Client) BatchGetVideoClipTask(request *BatchGetVideoClipTaskRequest) (_result *BatchGetVideoClipTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchGetVideoClipTaskResponse{}
	_body, _err := client.BatchGetVideoClipTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询文案
//
// @param tmpReq - BatchQueryIndividuationTextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryIndividuationTextResponse
func (client *Client) BatchQueryIndividuationTextWithOptions(tmpReq *BatchQueryIndividuationTextRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchQueryIndividuationTextResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchQueryIndividuationTextShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TextIdList)) {
		request.TextIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TextIdList, tea.String("textIdList"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TextIdListShrink)) {
		query["textIdList"] = request.TextIdListShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchQueryIndividuationText"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/individuationText/batchQueryText"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchQueryIndividuationTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询文案
//
// @param request - BatchQueryIndividuationTextRequest
//
// @return BatchQueryIndividuationTextResponse
func (client *Client) BatchQueryIndividuationText(request *BatchQueryIndividuationTextRequest) (_result *BatchQueryIndividuationTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchQueryIndividuationTextResponse{}
	_body, _err := client.BatchQueryIndividuationTextWithOptions(request, headers, runtime)
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
// 查询剧本列表
//
// @param request - CreateAICoachTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAICoachTaskResponse
func (client *Client) CreateAICoachTaskWithOptions(request *CreateAICoachTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAICoachTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ScriptRecordId)) {
		body["scriptRecordId"] = request.ScriptRecordId
	}

	if !tea.BoolValue(util.IsUnset(request.StudentAudioUrl)) {
		body["studentAudioUrl"] = request.StudentAudioUrl
	}

	if !tea.BoolValue(util.IsUnset(request.StudentId)) {
		body["studentId"] = request.StudentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAICoachTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/aicoach/createTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAICoachTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询剧本列表
//
// @param request - CreateAICoachTaskRequest
//
// @return CreateAICoachTaskResponse
func (client *Client) CreateAICoachTask(request *CreateAICoachTaskRequest) (_result *CreateAICoachTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAICoachTaskResponse{}
	_body, _err := client.CreateAICoachTaskWithOptions(request, headers, runtime)
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
// 创建照片数字人
//
// @param request - CreateAnchorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAnchorResponse
func (client *Client) CreateAnchorWithOptions(request *CreateAnchorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAnchorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnchorCategory)) {
		body["anchorCategory"] = request.AnchorCategory
	}

	if !tea.BoolValue(util.IsUnset(request.AnchorMaterialName)) {
		body["anchorMaterialName"] = request.AnchorMaterialName
	}

	if !tea.BoolValue(util.IsUnset(request.CoverUrl)) {
		body["coverUrl"] = request.CoverUrl
	}

	if !tea.BoolValue(util.IsUnset(request.DigitalHumanType)) {
		body["digitalHumanType"] = request.DigitalHumanType
	}

	if !tea.BoolValue(util.IsUnset(request.Gender)) {
		body["gender"] = request.Gender
	}

	if !tea.BoolValue(util.IsUnset(request.UseScene)) {
		body["useScene"] = request.UseScene
	}

	if !tea.BoolValue(util.IsUnset(request.VideoOssKey)) {
		body["videoOssKey"] = request.VideoOssKey
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAnchor"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/anchorOpen/createAnchor"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAnchorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建照片数字人
//
// @param request - CreateAnchorRequest
//
// @return CreateAnchorResponse
func (client *Client) CreateAnchor(request *CreateAnchorRequest) (_result *CreateAnchorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAnchorResponse{}
	_body, _err := client.CreateAnchorWithOptions(request, headers, runtime)
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
// 创建个性化文案项目
//
// @param request - CreateIndividuationProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIndividuationProjectResponse
func (client *Client) CreateIndividuationProjectWithOptions(request *CreateIndividuationProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateIndividuationProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectInfo)) {
		body["projectInfo"] = request.ProjectInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["projectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Purpose)) {
		body["purpose"] = request.Purpose
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		body["sceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIndividuationProject"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/individuationText/createProject"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIndividuationProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建个性化文案项目
//
// @param request - CreateIndividuationProjectRequest
//
// @return CreateIndividuationProjectResponse
func (client *Client) CreateIndividuationProject(request *CreateIndividuationProjectRequest) (_result *CreateIndividuationProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIndividuationProjectResponse{}
	_body, _err := client.CreateIndividuationProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建个性化文案任务
//
// @param request - CreateIndividuationTextTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIndividuationTextTaskResponse
func (client *Client) CreateIndividuationTextTaskWithOptions(request *CreateIndividuationTextTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateIndividuationTextTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CrowdPack)) {
		body["crowdPack"] = request.CrowdPack
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["taskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIndividuationTextTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/individuationText/createTextTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIndividuationTextTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建个性化文案任务
//
// @param request - CreateIndividuationTextTaskRequest
//
// @return CreateIndividuationTextTaskResponse
func (client *Client) CreateIndividuationTextTask(request *CreateIndividuationTextTaskRequest) (_result *CreateIndividuationTextTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIndividuationTextTaskResponse{}
	_body, _err := client.CreateIndividuationTextTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建产品图
//
// @param request - CreateProductImageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProductImageResponse
func (client *Client) CreateProductImageWithOptions(request *CreateProductImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProductImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackgroundDescription)) {
		body["backgroundDescription"] = request.BackgroundDescription
	}

	if !tea.BoolValue(util.IsUnset(request.BackgroundPriority)) {
		body["backgroundPriority"] = request.BackgroundPriority
	}

	if !tea.BoolValue(util.IsUnset(request.BackgroundUrl)) {
		body["backgroundUrl"] = request.BackgroundUrl
	}

	if !tea.BoolValue(util.IsUnset(request.HighlightText)) {
		body["highlightText"] = request.HighlightText
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCount)) {
		body["imageCount"] = request.ImageCount
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["imageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SubTitle)) {
		body["subTitle"] = request.SubTitle
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProductImage"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/images/products"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProductImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建产品图
//
// @param request - CreateProductImageRequest
//
// @return CreateProductImageResponse
func (client *Client) CreateProductImage(request *CreateProductImageRequest) (_result *CreateProductImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProductImageResponse{}
	_body, _err := client.CreateProductImageWithOptions(request, headers, runtime)
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
// 提交声音复刻任务
//
// @param request - CreateTrainTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrainTaskResponse
func (client *Client) CreateTrainTaskWithOptions(request *CreateTrainTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTrainTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunMainId)) {
		body["aliyunMainId"] = request.AliyunMainId
	}

	if !tea.BoolValue(util.IsUnset(request.ResSpecType)) {
		body["resSpecType"] = request.ResSpecType
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["taskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.UseScene)) {
		body["useScene"] = request.UseScene
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceGender)) {
		body["voiceGender"] = request.VoiceGender
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceLanguage)) {
		body["voiceLanguage"] = request.VoiceLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceName)) {
		body["voiceName"] = request.VoiceName
	}

	if !tea.BoolValue(util.IsUnset(request.VoicePath)) {
		body["voicePath"] = request.VoicePath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrainTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/train/task/createTrainTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTrainTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交声音复刻任务
//
// @param request - CreateTrainTaskRequest
//
// @return CreateTrainTaskResponse
func (client *Client) CreateTrainTask(request *CreateTrainTaskRequest) (_result *CreateTrainTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTrainTaskResponse{}
	_body, _err := client.CreateTrainTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交视频切片任务
//
// @param request - CreateVideoClipTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVideoClipTaskResponse
func (client *Client) CreateVideoClipTaskWithOptions(request *CreateVideoClipTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateVideoClipTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunMainId)) {
		body["aliyunMainId"] = request.AliyunMainId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OssKeys)) {
		body["ossKeys"] = request.OssKeys
	}

	if !tea.BoolValue(util.IsUnset(request.Requirement)) {
		body["requirement"] = request.Requirement
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVideoClipTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/video/clip/createVideoClipTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVideoClipTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交视频切片任务
//
// @param request - CreateVideoClipTaskRequest
//
// @return CreateVideoClipTaskResponse
func (client *Client) CreateVideoClipTask(request *CreateVideoClipTaskRequest) (_result *CreateVideoClipTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateVideoClipTaskResponse{}
	_body, _err := client.CreateVideoClipTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除个性化文案项目
//
// @param request - DeleteIndividuationProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIndividuationProjectResponse
func (client *Client) DeleteIndividuationProjectWithOptions(request *DeleteIndividuationProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteIndividuationProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIndividuationProject"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/individuationText/deleteProject"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIndividuationProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除个性化文案项目
//
// @param request - DeleteIndividuationProjectRequest
//
// @return DeleteIndividuationProjectResponse
func (client *Client) DeleteIndividuationProject(request *DeleteIndividuationProjectRequest) (_result *DeleteIndividuationProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIndividuationProjectResponse{}
	_body, _err := client.DeleteIndividuationProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除个性化文案
//
// @param request - DeleteIndividuationTextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIndividuationTextResponse
func (client *Client) DeleteIndividuationTextWithOptions(request *DeleteIndividuationTextRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteIndividuationTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TextIdList)) {
		body["textIdList"] = request.TextIdList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIndividuationText"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/individuationText/deleteText"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIndividuationTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除个性化文案
//
// @param request - DeleteIndividuationTextRequest
//
// @return DeleteIndividuationTextResponse
func (client *Client) DeleteIndividuationText(request *DeleteIndividuationTextRequest) (_result *DeleteIndividuationTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIndividuationTextResponse{}
	_body, _err := client.DeleteIndividuationTextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询文档信息与状态
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDocumentResponse
func (client *Client) DescribeDocumentWithOptions(knowledgeBaseId *string, documentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeDocumentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDocument"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/knowledge-base/" + tea.StringValue(openapiutil.GetEncodeParam(knowledgeBaseId)) + "/documents/" + tea.StringValue(openapiutil.GetEncodeParam(documentId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询文档信息与状态
//
// @return DescribeDocumentResponse
func (client *Client) DescribeDocument(knowledgeBaseId *string, documentId *string) (_result *DescribeDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeDocumentResponse{}
	_body, _err := client.DescribeDocumentWithOptions(knowledgeBaseId, documentId, headers, runtime)
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
// 获取考核点详情
//
// @param request - GetAICoachAssessmentPointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAICoachAssessmentPointResponse
func (client *Client) GetAICoachAssessmentPointWithOptions(request *GetAICoachAssessmentPointRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAICoachAssessmentPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PointId)) {
		query["pointId"] = request.PointId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAICoachAssessmentPoint"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/aicoach/getAssessmentPoint"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAICoachAssessmentPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取考核点详情
//
// @param request - GetAICoachAssessmentPointRequest
//
// @return GetAICoachAssessmentPointResponse
func (client *Client) GetAICoachAssessmentPoint(request *GetAICoachAssessmentPointRequest) (_result *GetAICoachAssessmentPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAICoachAssessmentPointResponse{}
	_body, _err := client.GetAICoachAssessmentPointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询作弊检测详情
//
// @param request - GetAICoachCheatDetectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAICoachCheatDetectionResponse
func (client *Client) GetAICoachCheatDetectionWithOptions(request *GetAICoachCheatDetectionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAICoachCheatDetectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["sessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAICoachCheatDetection"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/aicoach/getCheatDetection"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAICoachCheatDetectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询作弊检测详情
//
// @param request - GetAICoachCheatDetectionRequest
//
// @return GetAICoachCheatDetectionResponse
func (client *Client) GetAICoachCheatDetection(request *GetAICoachCheatDetectionRequest) (_result *GetAICoachCheatDetectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAICoachCheatDetectionResponse{}
	_body, _err := client.GetAICoachCheatDetectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询剧本详情
//
// @param request - GetAICoachScriptRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAICoachScriptResponse
func (client *Client) GetAICoachScriptWithOptions(request *GetAICoachScriptRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAICoachScriptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ScriptRecordId)) {
		query["scriptRecordId"] = request.ScriptRecordId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAICoachScript"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/aicoach/getScript"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAICoachScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询剧本详情
//
// @param request - GetAICoachScriptRequest
//
// @return GetAICoachScriptResponse
func (client *Client) GetAICoachScript(request *GetAICoachScriptRequest) (_result *GetAICoachScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAICoachScriptResponse{}
	_body, _err := client.GetAICoachScriptWithOptions(request, headers, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

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

	if !tea.BoolValue(util.IsUnset(request.UploadType)) {
		query["uploadType"] = request.UploadType
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
// 查询剧本列表
//
// @param request - ListAICoachScriptPageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAICoachScriptPageResponse
func (client *Client) ListAICoachScriptPageWithOptions(request *ListAICoachScriptPageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAICoachScriptPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAICoachScriptPage"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/aicoach/pageScript"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAICoachScriptPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询剧本列表
//
// @param request - ListAICoachScriptPageRequest
//
// @return ListAICoachScriptPageResponse
func (client *Client) ListAICoachScriptPage(request *ListAICoachScriptPageRequest) (_result *ListAICoachScriptPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAICoachScriptPageResponse{}
	_body, _err := client.ListAICoachScriptPageWithOptions(request, headers, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
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
// 分页查询智能体
//
// @param request - ListAgentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentsResponse
func (client *Client) ListAgentsWithOptions(request *ListAgentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAgentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		query["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.AgentScene)) {
		query["agentScene"] = request.AgentScene
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		query["owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAgents"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/agent/listAgents"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAgentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询智能体
//
// @param request - ListAgentsRequest
//
// @return ListAgentsResponse
func (client *Client) ListAgents(request *ListAgentsRequest) (_result *ListAgentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentsResponse{}
	_body, _err := client.ListAgentsWithOptions(request, headers, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AnchorCategory)) {
		query["anchorCategory"] = request.AnchorCategory
	}

	if !tea.BoolValue(util.IsUnset(request.AnchorId)) {
		query["anchorId"] = request.AnchorId
	}

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
// 查询知识库
//
// @param request - ListKnowledgeBaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKnowledgeBaseResponse
func (client *Client) ListKnowledgeBaseWithOptions(request *ListKnowledgeBaseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListKnowledgeBaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KnowledgeBaseId)) {
		query["knowledgeBaseId"] = request.KnowledgeBaseId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListKnowledgeBase"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/knowledge-base"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListKnowledgeBaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询知识库
//
// @param request - ListKnowledgeBaseRequest
//
// @return ListKnowledgeBaseResponse
func (client *Client) ListKnowledgeBase(request *ListKnowledgeBaseRequest) (_result *ListKnowledgeBaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListKnowledgeBaseResponse{}
	_body, _err := client.ListKnowledgeBaseWithOptions(request, headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.VoiceLanguage)) {
		query["voiceLanguage"] = request.VoiceLanguage
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
// 查询图片转视频任务
//
// @param request - QueryImageToVideoTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryImageToVideoTaskResponse
func (client *Client) QueryImageToVideoTaskWithOptions(request *QueryImageToVideoTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryImageToVideoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryImageToVideoTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/video/imageToVideo/task"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryImageToVideoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询图片转视频任务
//
// @param request - QueryImageToVideoTaskRequest
//
// @return QueryImageToVideoTaskResponse
func (client *Client) QueryImageToVideoTask(request *QueryImageToVideoTaskRequest) (_result *QueryImageToVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryImageToVideoTaskResponse{}
	_body, _err := client.QueryImageToVideoTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询个性化文案任务
//
// @param request - QueryIndividuationTextTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryIndividuationTextTaskResponse
func (client *Client) QueryIndividuationTextTaskWithOptions(request *QueryIndividuationTextTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryIndividuationTextTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryIndividuationTextTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/individuationText/queryTextTask"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryIndividuationTextTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询个性化文案任务
//
// @param request - QueryIndividuationTextTaskRequest
//
// @return QueryIndividuationTextTaskResponse
func (client *Client) QueryIndividuationTextTask(request *QueryIndividuationTextTaskRequest) (_result *QueryIndividuationTextTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryIndividuationTextTaskResponse{}
	_body, _err := client.QueryIndividuationTextTaskWithOptions(request, headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.BitRate)) {
		body["bitRate"] = request.BitRate
	}

	if !tea.BoolValue(util.IsUnset(request.FrameRate)) {
		body["frameRate"] = request.FrameRate
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

	if !tea.BoolValue(util.IsUnset(request.Resolution)) {
		body["resolution"] = request.Resolution
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleType)) {
		body["scaleType"] = request.ScaleType
	}

	if !tea.BoolValue(util.IsUnset(request.ScriptModelTag)) {
		body["scriptModelTag"] = request.ScriptModelTag
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizedDisplay)) {
		body["synchronizedDisplay"] = request.SynchronizedDisplay
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
// 发送sdk消息
//
// @param request - SendSdkMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendSdkMessageResponse
func (client *Client) SendSdkMessageWithOptions(request *SendSdkMessageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SendSdkMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Header)) {
		body["header"] = request.Header
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		body["moduleName"] = request.ModuleName
	}

	if !tea.BoolValue(util.IsUnset(request.OperationName)) {
		body["operationName"] = request.OperationName
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendSdkMessage"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/sdk/sendMessage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SendSdkMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送sdk消息
//
// @param request - SendSdkMessageRequest
//
// @return SendSdkMessageResponse
func (client *Client) SendSdkMessage(request *SendSdkMessageRequest) (_result *SendSdkMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendSdkMessageResponse{}
	_body, _err := client.SendSdkMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送sdk流式消息
//
// @param request - SendSdkStreamMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendSdkStreamMessageResponse
func (client *Client) SendSdkStreamMessageWithOptions(request *SendSdkStreamMessageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SendSdkStreamMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Header)) {
		body["header"] = request.Header
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		body["moduleName"] = request.ModuleName
	}

	if !tea.BoolValue(util.IsUnset(request.OperationName)) {
		body["operationName"] = request.OperationName
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendSdkStreamMessage"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/sdk/stream/sendMessage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SendSdkStreamMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送sdk流式消息
//
// @param request - SendSdkStreamMessageRequest
//
// @return SendSdkStreamMessageResponse
func (client *Client) SendSdkStreamMessage(request *SendSdkStreamMessageRequest) (_result *SendSdkStreamMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendSdkStreamMessageResponse{}
	_body, _err := client.SendSdkStreamMessageWithOptions(request, headers, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.ChannelToken)) {
		body["channelToken"] = request.ChannelToken
	}

	if !tea.BoolValue(util.IsUnset(request.CustomPushUrl)) {
		body["customPushUrl"] = request.CustomPushUrl
	}

	if !tea.BoolValue(util.IsUnset(request.CustomUserId)) {
		body["customUserId"] = request.CustomUserId
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
// 提交图片转视频任务
//
// @param request - SubmitImageToVideoTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitImageToVideoTaskResponse
func (client *Client) SubmitImageToVideoTaskWithOptions(request *SubmitImageToVideoTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitImageToVideoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["imageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PosPrompt)) {
		body["posPrompt"] = request.PosPrompt
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitImageToVideoTask"),
		Version:     tea.String("2024-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/video/imageToVideo/task"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitImageToVideoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交图片转视频任务
//
// @param request - SubmitImageToVideoTaskRequest
//
// @return SubmitImageToVideoTaskResponse
func (client *Client) SubmitImageToVideoTask(request *SubmitImageToVideoTaskRequest) (_result *SubmitImageToVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitImageToVideoTaskResponse{}
	_body, _err := client.SubmitImageToVideoTaskWithOptions(request, headers, runtime)
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
