// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateEditingProjectRequest struct {
	// 云剪辑工程标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 云剪辑工程描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 云剪辑工程时间线，Json格式
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// 云剪辑工程封面
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 前端编辑器工程数据结构
	FEExtend *string `json:"FEExtend,omitempty" xml:"FEExtend,omitempty"`
}

func (s CreateEditingProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateEditingProjectRequest) SetTitle(v string) *CreateEditingProjectRequest {
	s.Title = &v
	return s
}

func (s *CreateEditingProjectRequest) SetDescription(v string) *CreateEditingProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateEditingProjectRequest) SetTimeline(v string) *CreateEditingProjectRequest {
	s.Timeline = &v
	return s
}

func (s *CreateEditingProjectRequest) SetCoverURL(v string) *CreateEditingProjectRequest {
	s.CoverURL = &v
	return s
}

func (s *CreateEditingProjectRequest) SetFEExtend(v string) *CreateEditingProjectRequest {
	s.FEExtend = &v
	return s
}

type CreateEditingProjectResponseBody struct {
	// Id of the request
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Project   *CreateEditingProjectResponseBodyProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
}

func (s CreateEditingProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEditingProjectResponseBody) SetRequestId(v string) *CreateEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEditingProjectResponseBody) SetProject(v *CreateEditingProjectResponseBodyProject) *CreateEditingProjectResponseBody {
	s.Project = v
	return s
}

type CreateEditingProjectResponseBodyProject struct {
	ProjectId    *string  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Title        *string  `json:"Title,omitempty" xml:"Title,omitempty"`
	Description  *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Timeline     *string  `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	CoverURL     *string  `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	Status       *int64   `json:"Status,omitempty" xml:"Status,omitempty"`
	CreateTime   *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifiedTime *string  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Duration     *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s CreateEditingProjectResponseBodyProject) String() string {
	return tea.Prettify(s)
}

func (s CreateEditingProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *CreateEditingProjectResponseBodyProject) SetProjectId(v string) *CreateEditingProjectResponseBodyProject {
	s.ProjectId = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetTitle(v string) *CreateEditingProjectResponseBodyProject {
	s.Title = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetDescription(v string) *CreateEditingProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetTimeline(v string) *CreateEditingProjectResponseBodyProject {
	s.Timeline = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetCoverURL(v string) *CreateEditingProjectResponseBodyProject {
	s.CoverURL = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetStatus(v int64) *CreateEditingProjectResponseBodyProject {
	s.Status = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetCreateTime(v string) *CreateEditingProjectResponseBodyProject {
	s.CreateTime = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetModifiedTime(v string) *CreateEditingProjectResponseBodyProject {
	s.ModifiedTime = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetDuration(v float32) *CreateEditingProjectResponseBodyProject {
	s.Duration = &v
	return s
}

type CreateEditingProjectResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEditingProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateEditingProjectResponse) SetHeaders(v map[string]*string) *CreateEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateEditingProjectResponse) SetBody(v *CreateEditingProjectResponseBody) *CreateEditingProjectResponse {
	s.Body = v
	return s
}

type DeleteEditingProjectsRequest struct {
	// 云剪辑工程ID。支持多个云剪辑工程，以逗号分隔。
	ProjectIds *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
}

func (s DeleteEditingProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEditingProjectsRequest) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectsRequest) SetProjectIds(v string) *DeleteEditingProjectsRequest {
	s.ProjectIds = &v
	return s
}

type DeleteEditingProjectsResponseBody struct {
	// Id of the request
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IgnoredList *string `json:"IgnoredList,omitempty" xml:"IgnoredList,omitempty"`
}

func (s DeleteEditingProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEditingProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectsResponseBody) SetRequestId(v string) *DeleteEditingProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEditingProjectsResponseBody) SetIgnoredList(v string) *DeleteEditingProjectsResponseBody {
	s.IgnoredList = &v
	return s
}

type DeleteEditingProjectsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEditingProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEditingProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEditingProjectsResponse) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectsResponse) SetHeaders(v map[string]*string) *DeleteEditingProjectsResponse {
	s.Headers = v
	return s
}

func (s *DeleteEditingProjectsResponse) SetBody(v *DeleteEditingProjectsResponseBody) *DeleteEditingProjectsResponse {
	s.Body = v
	return s
}

type DeleteMediaInfosRequest struct {
	// ICE 媒资ID
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// 待注册的媒资在相应系统中的地址
	InputURLs *string `json:"InputURLs,omitempty" xml:"InputURLs,omitempty"`
}

func (s DeleteMediaInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMediaInfosRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediaInfosRequest) SetMediaIds(v string) *DeleteMediaInfosRequest {
	s.MediaIds = &v
	return s
}

func (s *DeleteMediaInfosRequest) SetInputURLs(v string) *DeleteMediaInfosRequest {
	s.InputURLs = &v
	return s
}

type DeleteMediaInfosResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 出现获取错误的ID或inputUr
	IgnoredList *string `json:"IgnoredList,omitempty" xml:"IgnoredList,omitempty"`
}

func (s DeleteMediaInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMediaInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediaInfosResponseBody) SetRequestId(v string) *DeleteMediaInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediaInfosResponseBody) SetIgnoredList(v string) *DeleteMediaInfosResponseBody {
	s.IgnoredList = &v
	return s
}

type DeleteMediaInfosResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMediaInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMediaInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMediaInfosResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediaInfosResponse) SetHeaders(v map[string]*string) *DeleteMediaInfosResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediaInfosResponse) SetBody(v *DeleteMediaInfosResponseBody) *DeleteMediaInfosResponse {
	s.Body = v
	return s
}

type DescribeIceProductStatusResponseBody struct {
	// Id of the request
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ICEServiceAvaliable *bool   `json:"ICEServiceAvaliable,omitempty" xml:"ICEServiceAvaliable,omitempty"`
}

func (s DescribeIceProductStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIceProductStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIceProductStatusResponseBody) SetRequestId(v string) *DescribeIceProductStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIceProductStatusResponseBody) SetICEServiceAvaliable(v bool) *DescribeIceProductStatusResponseBody {
	s.ICEServiceAvaliable = &v
	return s
}

type DescribeIceProductStatusResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIceProductStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIceProductStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIceProductStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeIceProductStatusResponse) SetHeaders(v map[string]*string) *DescribeIceProductStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeIceProductStatusResponse) SetBody(v *DescribeIceProductStatusResponseBody) *DescribeIceProductStatusResponse {
	s.Body = v
	return s
}

type DescribeRelatedAuthorizationStatusResponseBody struct {
	// Id of the request
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OSSAuthorized *bool   `json:"OSSAuthorized,omitempty" xml:"OSSAuthorized,omitempty"`
	MTSAuthorized *bool   `json:"MTSAuthorized,omitempty" xml:"MTSAuthorized,omitempty"`
	MNSAuthorized *bool   `json:"MNSAuthorized,omitempty" xml:"MNSAuthorized,omitempty"`
}

func (s DescribeRelatedAuthorizationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelatedAuthorizationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRelatedAuthorizationStatusResponseBody) SetRequestId(v string) *DescribeRelatedAuthorizationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRelatedAuthorizationStatusResponseBody) SetOSSAuthorized(v bool) *DescribeRelatedAuthorizationStatusResponseBody {
	s.OSSAuthorized = &v
	return s
}

func (s *DescribeRelatedAuthorizationStatusResponseBody) SetMTSAuthorized(v bool) *DescribeRelatedAuthorizationStatusResponseBody {
	s.MTSAuthorized = &v
	return s
}

func (s *DescribeRelatedAuthorizationStatusResponseBody) SetMNSAuthorized(v bool) *DescribeRelatedAuthorizationStatusResponseBody {
	s.MNSAuthorized = &v
	return s
}

type DescribeRelatedAuthorizationStatusResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRelatedAuthorizationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRelatedAuthorizationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelatedAuthorizationStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeRelatedAuthorizationStatusResponse) SetHeaders(v map[string]*string) *DescribeRelatedAuthorizationStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeRelatedAuthorizationStatusResponse) SetBody(v *DescribeRelatedAuthorizationStatusResponseBody) *DescribeRelatedAuthorizationStatusResponse {
	s.Body = v
	return s
}

type GetEditingProjectRequest struct {
	// 云剪辑工程ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 是否返回编辑器数据结构
	FEExtendFlag *int64 `json:"FEExtendFlag,omitempty" xml:"FEExtendFlag,omitempty"`
}

func (s GetEditingProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *GetEditingProjectRequest) SetProjectId(v string) *GetEditingProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *GetEditingProjectRequest) SetFEExtendFlag(v int64) *GetEditingProjectRequest {
	s.FEExtendFlag = &v
	return s
}

type GetEditingProjectResponseBody struct {
	// Id of the request
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Project   *GetEditingProjectResponseBodyProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
}

func (s GetEditingProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetEditingProjectResponseBody) SetRequestId(v string) *GetEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEditingProjectResponseBody) SetProject(v *GetEditingProjectResponseBodyProject) *GetEditingProjectResponseBody {
	s.Project = v
	return s
}

type GetEditingProjectResponseBodyProject struct {
	// 云剪辑工程ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 云剪辑工程标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 云剪辑工程时间线
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// 云剪辑工程描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 云剪辑工程封面
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 云剪辑工程创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 云剪辑工程最新修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// 云剪辑工程总时长
	Duration *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetEditingProjectResponseBodyProject) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *GetEditingProjectResponseBodyProject) SetProjectId(v string) *GetEditingProjectResponseBodyProject {
	s.ProjectId = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetTitle(v string) *GetEditingProjectResponseBodyProject {
	s.Title = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetTimeline(v string) *GetEditingProjectResponseBodyProject {
	s.Timeline = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetDescription(v string) *GetEditingProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetCoverURL(v string) *GetEditingProjectResponseBodyProject {
	s.CoverURL = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetCreateTime(v string) *GetEditingProjectResponseBodyProject {
	s.CreateTime = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetModifiedTime(v string) *GetEditingProjectResponseBodyProject {
	s.ModifiedTime = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetDuration(v int64) *GetEditingProjectResponseBodyProject {
	s.Duration = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetStatus(v string) *GetEditingProjectResponseBodyProject {
	s.Status = &v
	return s
}

type GetEditingProjectResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEditingProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *GetEditingProjectResponse) SetHeaders(v map[string]*string) *GetEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *GetEditingProjectResponse) SetBody(v *GetEditingProjectResponseBody) *GetEditingProjectResponse {
	s.Body = v
	return s
}

type GetMediaInfoRequest struct {
	MediaId  *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
}

func (s GetMediaInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMediaInfoRequest) SetMediaId(v string) *GetMediaInfoRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaInfoRequest) SetInputURL(v string) *GetMediaInfoRequest {
	s.InputURL = &v
	return s
}

type GetMediaInfoResponseBody struct {
	// RequestId
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MediaInfo *GetMediaInfoResponseBodyMediaInfo `json:"MediaInfo,omitempty" xml:"MediaInfo,omitempty" type:"Struct"`
}

func (s GetMediaInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBody) SetRequestId(v string) *GetMediaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaInfoResponseBody) SetMediaInfo(v *GetMediaInfoResponseBodyMediaInfo) *GetMediaInfoResponseBody {
	s.MediaInfo = v
	return s
}

type GetMediaInfoResponseBodyMediaInfo struct {
	// 媒资ID
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// BasicInfo
	MediaBasicInfo *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// 其他元数据
	DynamicMetaDataList []*GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList `json:"DynamicMetaDataList,omitempty" xml:"DynamicMetaDataList,omitempty" type:"Repeated"`
	// AIMetadata
	AiRoughDataList []*GetMediaInfoResponseBodyMediaInfoAiRoughDataList `json:"AiRoughDataList,omitempty" xml:"AiRoughDataList,omitempty" type:"Repeated"`
	// FileInfos
	FileInfoList []*GetMediaInfoResponseBodyMediaInfoFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
}

func (s GetMediaInfoResponseBodyMediaInfo) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfo) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfo) SetMediaId(v string) *GetMediaInfoResponseBodyMediaInfo {
	s.MediaId = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfo) SetMediaBasicInfo(v *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) *GetMediaInfoResponseBodyMediaInfo {
	s.MediaBasicInfo = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfo) SetDynamicMetaDataList(v []*GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList) *GetMediaInfoResponseBodyMediaInfo {
	s.DynamicMetaDataList = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfo) SetAiRoughDataList(v []*GetMediaInfoResponseBodyMediaInfoAiRoughDataList) *GetMediaInfoResponseBodyMediaInfo {
	s.AiRoughDataList = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfo) SetFileInfoList(v []*GetMediaInfoResponseBodyMediaInfoFileInfoList) *GetMediaInfoResponseBodyMediaInfo {
	s.FileInfoList = v
	return s
}

type GetMediaInfoResponseBodyMediaInfoMediaBasicInfo struct {
	// MediaId
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// 待注册的媒资在相应系统中的地址
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// 媒资媒体类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 内容描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 标签
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// 封面地址
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 用户数据
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 资源状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 媒资创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 媒资修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// 媒资删除时间
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetMediaId(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetInputURL(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetMediaType(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetBusinessType(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetSource(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetTitle(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetDescription(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCategory(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetMediaTags(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCoverURL(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetUserData(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetStatus(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCreateTime(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetModifiedTime(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetDeletedTime(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

type GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList struct {
	// 开始时间
	In *float32 `json:"In,omitempty" xml:"In,omitempty"`
	// 结束时间
	Out *float32 `json:"Out,omitempty" xml:"Out,omitempty"`
	// 类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 元数据json string
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList) SetIn(v float32) *GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList {
	s.In = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList) SetOut(v float32) *GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList {
	s.Out = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList) SetType(v string) *GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList {
	s.Type = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList) SetData(v string) *GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList {
	s.Data = &v
	return s
}

type GetMediaInfoResponseBodyMediaInfoAiRoughDataList struct {
	// AI类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// AI原始结果
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoAiRoughDataList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoAiRoughDataList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataList) SetType(v string) *GetMediaInfoResponseBodyMediaInfoAiRoughDataList {
	s.Type = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataList) SetResult(v string) *GetMediaInfoResponseBodyMediaInfoAiRoughDataList {
	s.Result = &v
	return s
}

type GetMediaInfoResponseBodyMediaInfoFileInfoList struct {
	// 文件基础信息，包含时长，大小等
	FileBasicInfo *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
	// 音频流信息，一个媒资可能有多条音频流
	AudioStreamInfoList []*GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList `json:"AudioStreamInfoList,omitempty" xml:"AudioStreamInfoList,omitempty" type:"Repeated"`
	// 视频流信息，一个媒资可能有多条视频流
	VideoStreamInfoList []*GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList `json:"VideoStreamInfoList,omitempty" xml:"VideoStreamInfoList,omitempty" type:"Repeated"`
	// 字幕流信息，一个媒资可能有多条字幕流
	SubtitleStreamInfoList []*GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList `json:"SubtitleStreamInfoList,omitempty" xml:"SubtitleStreamInfoList,omitempty" type:"Repeated"`
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) SetFileBasicInfo(v *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) *GetMediaInfoResponseBodyMediaInfoFileInfoList {
	s.FileBasicInfo = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) SetAudioStreamInfoList(v []*GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) *GetMediaInfoResponseBodyMediaInfoFileInfoList {
	s.AudioStreamInfoList = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) SetVideoStreamInfoList(v []*GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) *GetMediaInfoResponseBodyMediaInfoFileInfoList {
	s.VideoStreamInfoList = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) SetSubtitleStreamInfoList(v []*GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) *GetMediaInfoResponseBodyMediaInfoFileInfoList {
	s.SubtitleStreamInfoList = v
	return s
}

type GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo struct {
	// 文件名
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件状态
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// 文件类型
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// 文件大小（字节）
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// 文件oss地址
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// 文件存储区域
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 封装格式
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// 时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 码率
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// 宽
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// 高
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileStatus(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileType(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileSize(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileUrl(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetRegion(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFormatName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetDuration(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetBitrate(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetWidth(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetHeight(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

type GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList struct {
	// 音频流序号
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// 编码格式简述名
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// 编码格式长述名
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// 编码时基
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// 编码格式标记文本
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// 编码格式标记
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// 编码预置
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// 采样格式
	SampleFmt *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	// 采样率
	SampleRate *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// 声道数
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// 声道输出样式
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// 时基
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 码率
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// 音频帧率
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// 总帧数
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// 语言
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetIndex(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecLongName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTimeBase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTagString(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTag(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetProfile(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Profile = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetSampleFmt(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.SampleFmt = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetSampleRate(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.SampleRate = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetChannels(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Channels = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetChannelLayout(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.ChannelLayout = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetTimebase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetStartTime(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetDuration(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetBitrate(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetFps(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Fps = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetNumFrames(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetLang(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Lang = &v
	return s
}

type GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList struct {
	// 视频流序号
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// 编码格式简述名
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// 编码格式长述名
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// 编码预置
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// 编码时基
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// 编码格式标记文本
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// 编码格式标记
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// 宽
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// 高
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// 是否有B帧
	HasBFrames *string `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	// 编码信号分辨率比
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	// 编码显示分辨率比
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// 像素格式
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// 编码等级
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// 视频帧率
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// 平均帧率
	AvgFPS *string `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	// 时基
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 码率
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// 总帧数
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// 语言
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// 旋转
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// 总帧数
	NbFrames *string `json:"Nb_frames,omitempty" xml:"Nb_frames,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetIndex(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecLongName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetProfile(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Profile = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTimeBase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTagString(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTag(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetWidth(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Width = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetHeight(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Height = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetHasBFrames(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.HasBFrames = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetSar(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Sar = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetDar(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Dar = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetPixFmt(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.PixFmt = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetLevel(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Level = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetFps(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Fps = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetAvgFPS(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.AvgFPS = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetTimebase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetStartTime(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetDuration(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetBitrate(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetNumFrames(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetLang(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetRotate(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Rotate = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetNbFrames(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.NbFrames = &v
	return s
}

type GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList struct {
	// 音频流序号
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// 编码格式简述名
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// 编码格式长述名
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// 编码时基
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// 编码格式标记文本
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// 编码格式标记
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// 时基
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 语言
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetIndex(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecLongName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTimeBase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTagString(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTag(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetTimebase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetStartTime(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetDuration(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetLang(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Lang = &v
	return s
}

type GetMediaInfoResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMediaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponse) SetHeaders(v map[string]*string) *GetMediaInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMediaInfoResponse) SetBody(v *GetMediaInfoResponseBody) *GetMediaInfoResponse {
	s.Body = v
	return s
}

type GetMediaProducingJobRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetMediaProducingJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaProducingJobRequest) GoString() string {
	return s.String()
}

func (s *GetMediaProducingJobRequest) SetJobId(v string) *GetMediaProducingJobRequest {
	s.JobId = &v
	return s
}

type GetMediaProducingJobResponseBody struct {
	// Id of the request
	RequestId         *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MediaProducingJob *GetMediaProducingJobResponseBodyMediaProducingJob `json:"MediaProducingJob,omitempty" xml:"MediaProducingJob,omitempty" type:"Struct"`
}

func (s GetMediaProducingJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaProducingJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaProducingJobResponseBody) SetRequestId(v string) *GetMediaProducingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaProducingJobResponseBody) SetMediaProducingJob(v *GetMediaProducingJobResponseBodyMediaProducingJob) *GetMediaProducingJobResponseBody {
	s.MediaProducingJob = v
	return s
}

type GetMediaProducingJobResponseBodyMediaProducingJob struct {
	JobId        *string  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ProjectId    *string  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	MediaId      *string  `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	MediaURL     *string  `json:"MediaURL,omitempty" xml:"MediaURL,omitempty"`
	Timeline     *string  `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	TemplateId   *string  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	ClipsParam   *string  `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	Duration     *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	CreateTime   *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CompleteTime *string  `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	Status       *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Code         *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string  `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetMediaProducingJobResponseBodyMediaProducingJob) String() string {
	return tea.Prettify(s)
}

func (s GetMediaProducingJobResponseBodyMediaProducingJob) GoString() string {
	return s.String()
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetJobId(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.JobId = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetProjectId(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.ProjectId = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetMediaId(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.MediaId = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetMediaURL(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.MediaURL = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetTimeline(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.Timeline = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetTemplateId(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.TemplateId = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetClipsParam(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.ClipsParam = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetDuration(v float32) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.Duration = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetCreateTime(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.CreateTime = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetCompleteTime(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.CompleteTime = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetStatus(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.Status = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetCode(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.Code = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetMessage(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.Message = &v
	return s
}

type GetMediaProducingJobResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMediaProducingJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaProducingJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaProducingJobResponse) GoString() string {
	return s.String()
}

func (s *GetMediaProducingJobResponse) SetHeaders(v map[string]*string) *GetMediaProducingJobResponse {
	s.Headers = v
	return s
}

func (s *GetMediaProducingJobResponse) SetBody(v *GetMediaProducingJobResponseBody) *GetMediaProducingJobResponse {
	s.Body = v
	return s
}

type ListAllPublicMediaTagsRequest struct {
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
}

func (s ListAllPublicMediaTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAllPublicMediaTagsRequest) GoString() string {
	return s.String()
}

func (s *ListAllPublicMediaTagsRequest) SetBusinessType(v string) *ListAllPublicMediaTagsRequest {
	s.BusinessType = &v
	return s
}

type ListAllPublicMediaTagsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 公共素材库标签列表
	MediaTagList []*ListAllPublicMediaTagsResponseBodyMediaTagList `json:"MediaTagList,omitempty" xml:"MediaTagList,omitempty" type:"Repeated"`
}

func (s ListAllPublicMediaTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAllPublicMediaTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllPublicMediaTagsResponseBody) SetRequestId(v string) *ListAllPublicMediaTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllPublicMediaTagsResponseBody) SetMediaTagList(v []*ListAllPublicMediaTagsResponseBodyMediaTagList) *ListAllPublicMediaTagsResponseBody {
	s.MediaTagList = v
	return s
}

type ListAllPublicMediaTagsResponseBodyMediaTagList struct {
	// 素材标签id
	MediaTagId *string `json:"MediaTagId,omitempty" xml:"MediaTagId,omitempty"`
}

func (s ListAllPublicMediaTagsResponseBodyMediaTagList) String() string {
	return tea.Prettify(s)
}

func (s ListAllPublicMediaTagsResponseBodyMediaTagList) GoString() string {
	return s.String()
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagList) SetMediaTagId(v string) *ListAllPublicMediaTagsResponseBodyMediaTagList {
	s.MediaTagId = &v
	return s
}

type ListAllPublicMediaTagsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAllPublicMediaTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAllPublicMediaTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAllPublicMediaTagsResponse) GoString() string {
	return s.String()
}

func (s *ListAllPublicMediaTagsResponse) SetHeaders(v map[string]*string) *ListAllPublicMediaTagsResponse {
	s.Headers = v
	return s
}

func (s *ListAllPublicMediaTagsResponse) SetBody(v *ListAllPublicMediaTagsResponseBody) *ListAllPublicMediaTagsResponse {
	s.Body = v
	return s
}

type ListMediaBasicInfosRequest struct {
	// 创建时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 媒资媒体类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 资源状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 页号
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页大小
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 排序
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 页数
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// 分页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 返回值中是否包含文件基础信息
	IncludeFileBasicInfo *bool `json:"IncludeFileBasicInfo,omitempty" xml:"IncludeFileBasicInfo,omitempty"`
	// 针对媒资标题进行关键词搜索
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
}

func (s ListMediaBasicInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMediaBasicInfosRequest) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosRequest) SetStartTime(v string) *ListMediaBasicInfosRequest {
	s.StartTime = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetEndTime(v string) *ListMediaBasicInfosRequest {
	s.EndTime = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetMediaType(v string) *ListMediaBasicInfosRequest {
	s.MediaType = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetBusinessType(v string) *ListMediaBasicInfosRequest {
	s.BusinessType = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetSource(v string) *ListMediaBasicInfosRequest {
	s.Source = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetCategory(v string) *ListMediaBasicInfosRequest {
	s.Category = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetStatus(v string) *ListMediaBasicInfosRequest {
	s.Status = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetNextToken(v string) *ListMediaBasicInfosRequest {
	s.NextToken = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetMaxResults(v int32) *ListMediaBasicInfosRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetSortBy(v string) *ListMediaBasicInfosRequest {
	s.SortBy = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetPageNo(v int32) *ListMediaBasicInfosRequest {
	s.PageNo = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetPageSize(v int32) *ListMediaBasicInfosRequest {
	s.PageSize = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetIncludeFileBasicInfo(v bool) *ListMediaBasicInfosRequest {
	s.IncludeFileBasicInfo = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetKeyword(v string) *ListMediaBasicInfosRequest {
	s.Keyword = &v
	return s
}

type ListMediaBasicInfosResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合要求的媒资总数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 符合要求的媒资集合
	MediaInfos []*ListMediaBasicInfosResponseBodyMediaInfos `json:"MediaInfos,omitempty" xml:"MediaInfos,omitempty" type:"Repeated"`
	NextToken  *string                                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32                                       `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListMediaBasicInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMediaBasicInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponseBody) SetRequestId(v string) *ListMediaBasicInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaBasicInfosResponseBody) SetTotalCount(v int64) *ListMediaBasicInfosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMediaBasicInfosResponseBody) SetMediaInfos(v []*ListMediaBasicInfosResponseBodyMediaInfos) *ListMediaBasicInfosResponseBody {
	s.MediaInfos = v
	return s
}

func (s *ListMediaBasicInfosResponseBody) SetNextToken(v string) *ListMediaBasicInfosResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMediaBasicInfosResponseBody) SetMaxResults(v int32) *ListMediaBasicInfosResponseBody {
	s.MaxResults = &v
	return s
}

type ListMediaBasicInfosResponseBodyMediaInfos struct {
	// 媒资ID
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// BasicInfo
	MediaBasicInfo *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// FileInfos
	FileInfoList []*ListMediaBasicInfosResponseBodyMediaInfosFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
}

func (s ListMediaBasicInfosResponseBodyMediaInfos) String() string {
	return tea.Prettify(s)
}

func (s ListMediaBasicInfosResponseBodyMediaInfos) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponseBodyMediaInfos) SetMediaId(v string) *ListMediaBasicInfosResponseBodyMediaInfos {
	s.MediaId = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfos) SetMediaBasicInfo(v *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) *ListMediaBasicInfosResponseBodyMediaInfos {
	s.MediaBasicInfo = v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfos) SetFileInfoList(v []*ListMediaBasicInfosResponseBodyMediaInfosFileInfoList) *ListMediaBasicInfosResponseBodyMediaInfos {
	s.FileInfoList = v
	return s
}

type ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo struct {
	// MediaId
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// 待注册的媒资在相应系统中的地址
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// 媒资媒体类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 内容描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 标签
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// 封面地址
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 用户数据
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 截图
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// 资源状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 转码状态
	TranscodeStatus *string `json:"TranscodeStatus,omitempty" xml:"TranscodeStatus,omitempty"`
	// 媒资创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 媒资修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// 媒资删除时间
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
}

func (s ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaId(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetInputURL(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaType(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetBusinessType(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetSource(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetTitle(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetDescription(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCategory(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaTags(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCoverURL(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetUserData(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetSnapshots(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetStatus(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetTranscodeStatus(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.TranscodeStatus = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCreateTime(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetModifiedTime(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetDeletedTime(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

type ListMediaBasicInfosResponseBodyMediaInfosFileInfoList struct {
	// 文件基础信息，包含时长，大小等
	FileBasicInfo *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
}

func (s ListMediaBasicInfosResponseBodyMediaInfosFileInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListMediaBasicInfosResponseBodyMediaInfosFileInfoList) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoList) SetFileBasicInfo(v *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoList {
	s.FileBasicInfo = v
	return s
}

type ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo struct {
	// 文件名
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件状态
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// 文件类型
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// 文件大小（字节）
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// 文件oss地址
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// 文件存储区域
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 封装格式
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// 时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 码率
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// 宽
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// 高
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileName(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileStatus(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileType(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileSize(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileUrl(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetRegion(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFormatName(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetDuration(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetBitrate(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetWidth(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetHeight(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

type ListMediaBasicInfosResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMediaBasicInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMediaBasicInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMediaBasicInfosResponse) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponse) SetHeaders(v map[string]*string) *ListMediaBasicInfosResponse {
	s.Headers = v
	return s
}

func (s *ListMediaBasicInfosResponse) SetBody(v *ListMediaBasicInfosResponseBody) *ListMediaBasicInfosResponse {
	s.Body = v
	return s
}

type ListMediaProducingJobsRequest struct {
	// 查询以下状态的合成任务，支持多值，以英文逗号分隔
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMediaProducingJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMediaProducingJobsRequest) GoString() string {
	return s.String()
}

func (s *ListMediaProducingJobsRequest) SetStatus(v string) *ListMediaProducingJobsRequest {
	s.Status = &v
	return s
}

type ListMediaProducingJobsResponseBody struct {
	// Id of the request
	RequestId             *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MediaProducingJobList []*ListMediaProducingJobsResponseBodyMediaProducingJobList `json:"MediaProducingJobList,omitempty" xml:"MediaProducingJobList,omitempty" type:"Repeated"`
}

func (s ListMediaProducingJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMediaProducingJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaProducingJobsResponseBody) SetRequestId(v string) *ListMediaProducingJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaProducingJobsResponseBody) SetMediaProducingJobList(v []*ListMediaProducingJobsResponseBodyMediaProducingJobList) *ListMediaProducingJobsResponseBody {
	s.MediaProducingJobList = v
	return s
}

type ListMediaProducingJobsResponseBodyMediaProducingJobList struct {
	JobId        *string  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ProjectId    *string  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	MediaId      *string  `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	MediaURL     *string  `json:"MediaURL,omitempty" xml:"MediaURL,omitempty"`
	Timeline     *string  `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	TemplateId   *string  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	ClipsParam   *string  `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	Duration     *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	CreateTime   *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CompleteTime *string  `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	ModifiedTime *string  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Status       *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Code         *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string  `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListMediaProducingJobsResponseBodyMediaProducingJobList) String() string {
	return tea.Prettify(s)
}

func (s ListMediaProducingJobsResponseBodyMediaProducingJobList) GoString() string {
	return s.String()
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetJobId(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.JobId = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetProjectId(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.ProjectId = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetMediaId(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.MediaId = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetMediaURL(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.MediaURL = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetTimeline(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.Timeline = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetTemplateId(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.TemplateId = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetClipsParam(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.ClipsParam = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetDuration(v float32) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.Duration = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetCreateTime(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.CreateTime = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetCompleteTime(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.CompleteTime = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetModifiedTime(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.ModifiedTime = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetStatus(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.Status = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetCode(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.Code = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetMessage(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.Message = &v
	return s
}

type ListMediaProducingJobsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMediaProducingJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMediaProducingJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMediaProducingJobsResponse) GoString() string {
	return s.String()
}

func (s *ListMediaProducingJobsResponse) SetHeaders(v map[string]*string) *ListMediaProducingJobsResponse {
	s.Headers = v
	return s
}

func (s *ListMediaProducingJobsResponse) SetBody(v *ListMediaProducingJobsResponseBody) *ListMediaProducingJobsResponse {
	s.Body = v
	return s
}

type ListPublicMediaBasicInfosRequest struct {
	// 标签
	MediaTagId *string `json:"MediaTagId,omitempty" xml:"MediaTagId,omitempty"`
	// 下一次读取的位置
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页大小
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 页数
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// 分页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 返回值中是否包含文件基础信息
	IncludeFileBasicInfo *bool `json:"IncludeFileBasicInfo,omitempty" xml:"IncludeFileBasicInfo,omitempty"`
}

func (s ListPublicMediaBasicInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublicMediaBasicInfosRequest) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosRequest) SetMediaTagId(v string) *ListPublicMediaBasicInfosRequest {
	s.MediaTagId = &v
	return s
}

func (s *ListPublicMediaBasicInfosRequest) SetNextToken(v string) *ListPublicMediaBasicInfosRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicMediaBasicInfosRequest) SetMaxResults(v int32) *ListPublicMediaBasicInfosRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPublicMediaBasicInfosRequest) SetPageNo(v int32) *ListPublicMediaBasicInfosRequest {
	s.PageNo = &v
	return s
}

func (s *ListPublicMediaBasicInfosRequest) SetPageSize(v int32) *ListPublicMediaBasicInfosRequest {
	s.PageSize = &v
	return s
}

func (s *ListPublicMediaBasicInfosRequest) SetIncludeFileBasicInfo(v bool) *ListPublicMediaBasicInfosRequest {
	s.IncludeFileBasicInfo = &v
	return s
}

type ListPublicMediaBasicInfosResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合要求的媒资总数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 符合要求的媒资集合
	MediaInfos []*ListPublicMediaBasicInfosResponseBodyMediaInfos `json:"MediaInfos,omitempty" xml:"MediaInfos,omitempty" type:"Repeated"`
	NextToken  *string                                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32                                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListPublicMediaBasicInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponseBody) SetRequestId(v string) *ListPublicMediaBasicInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBody) SetTotalCount(v int64) *ListPublicMediaBasicInfosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBody) SetMediaInfos(v []*ListPublicMediaBasicInfosResponseBodyMediaInfos) *ListPublicMediaBasicInfosResponseBody {
	s.MediaInfos = v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBody) SetNextToken(v string) *ListPublicMediaBasicInfosResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBody) SetMaxResults(v int32) *ListPublicMediaBasicInfosResponseBody {
	s.MaxResults = &v
	return s
}

type ListPublicMediaBasicInfosResponseBodyMediaInfos struct {
	// 媒资ID
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// BasicInfo
	MediaBasicInfo *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// FileInfos
	FileInfoList []*ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfos) String() string {
	return tea.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfos) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfos) SetMediaId(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfos {
	s.MediaId = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfos) SetMediaBasicInfo(v *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) *ListPublicMediaBasicInfosResponseBodyMediaInfos {
	s.MediaBasicInfo = v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfos) SetFileInfoList(v []*ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList) *ListPublicMediaBasicInfosResponseBodyMediaInfos {
	s.FileInfoList = v
	return s
}

type ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo struct {
	// MediaId
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// 待注册的媒资在相应系统中的地址
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// 媒资媒体类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 内容描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 标签
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// 封面地址
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 用户数据
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 截图
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// 资源状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 转码状态
	TranscodeStatus *string `json:"TranscodeStatus,omitempty" xml:"TranscodeStatus,omitempty"`
	// 媒资创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 媒资修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// 媒资删除时间
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaId(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetInputURL(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaType(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetBusinessType(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetSource(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetTitle(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetDescription(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCategory(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaTags(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCoverURL(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetUserData(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetSnapshots(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetStatus(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetTranscodeStatus(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.TranscodeStatus = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCreateTime(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetModifiedTime(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetDeletedTime(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

type ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList struct {
	// 文件基础信息，包含时长，大小等
	FileBasicInfo *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList) SetFileBasicInfo(v *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList {
	s.FileBasicInfo = v
	return s
}

type ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo struct {
	// 文件名
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件状态
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// 文件类型
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// 文件大小（字节）
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// 文件oss地址
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// 文件存储区域
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 封装格式
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// 时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 码率
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// 宽
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// 高
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileName(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileStatus(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileType(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileSize(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileUrl(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetRegion(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFormatName(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetDuration(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetBitrate(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetWidth(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetHeight(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

type ListPublicMediaBasicInfosResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPublicMediaBasicInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPublicMediaBasicInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponse) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponse) SetHeaders(v map[string]*string) *ListPublicMediaBasicInfosResponse {
	s.Headers = v
	return s
}

func (s *ListPublicMediaBasicInfosResponse) SetBody(v *ListPublicMediaBasicInfosResponseBody) *ListPublicMediaBasicInfosResponse {
	s.Body = v
	return s
}

type RegisterMediaInfoRequest struct {
	// 媒资媒体url
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// 媒资媒体类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 标签,如果有多个标签用逗号隔开
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// 封面图，仅视频媒资有效
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 用户自定义元数据
	DynamicMetaDataList *string `json:"DynamicMetaDataList,omitempty" xml:"DynamicMetaDataList,omitempty"`
	// 用户数据，最大1024字节
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 是否覆盖已有媒资
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// 客户端token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s RegisterMediaInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterMediaInfoRequest) GoString() string {
	return s.String()
}

func (s *RegisterMediaInfoRequest) SetInputURL(v string) *RegisterMediaInfoRequest {
	s.InputURL = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetMediaType(v string) *RegisterMediaInfoRequest {
	s.MediaType = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetBusinessType(v string) *RegisterMediaInfoRequest {
	s.BusinessType = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetTitle(v string) *RegisterMediaInfoRequest {
	s.Title = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetDescription(v string) *RegisterMediaInfoRequest {
	s.Description = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetCategory(v string) *RegisterMediaInfoRequest {
	s.Category = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetMediaTags(v string) *RegisterMediaInfoRequest {
	s.MediaTags = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetCoverURL(v string) *RegisterMediaInfoRequest {
	s.CoverURL = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetDynamicMetaDataList(v string) *RegisterMediaInfoRequest {
	s.DynamicMetaDataList = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetUserData(v string) *RegisterMediaInfoRequest {
	s.UserData = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetOverwrite(v bool) *RegisterMediaInfoRequest {
	s.Overwrite = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetClientToken(v string) *RegisterMediaInfoRequest {
	s.ClientToken = &v
	return s
}

type RegisterMediaInfoResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// ICE媒资ID
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s RegisterMediaInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterMediaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterMediaInfoResponseBody) SetRequestId(v string) *RegisterMediaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterMediaInfoResponseBody) SetMediaId(v string) *RegisterMediaInfoResponseBody {
	s.MediaId = &v
	return s
}

type RegisterMediaInfoResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterMediaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterMediaInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterMediaInfoResponse) GoString() string {
	return s.String()
}

func (s *RegisterMediaInfoResponse) SetHeaders(v map[string]*string) *RegisterMediaInfoResponse {
	s.Headers = v
	return s
}

func (s *RegisterMediaInfoResponse) SetBody(v *RegisterMediaInfoResponseBody) *RegisterMediaInfoResponse {
	s.Body = v
	return s
}

type SearchEditingProjectRequest struct {
	// CreateTime（创建时间）的开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// CreationTime（创建时间）的结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 云剪辑工程状态。多个用逗号分隔
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 结果排序方式
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 分页参数
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页参数
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 分页参数
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// 分页参数
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s SearchEditingProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectRequest) SetStartTime(v string) *SearchEditingProjectRequest {
	s.StartTime = &v
	return s
}

func (s *SearchEditingProjectRequest) SetEndTime(v string) *SearchEditingProjectRequest {
	s.EndTime = &v
	return s
}

func (s *SearchEditingProjectRequest) SetStatus(v string) *SearchEditingProjectRequest {
	s.Status = &v
	return s
}

func (s *SearchEditingProjectRequest) SetSortBy(v string) *SearchEditingProjectRequest {
	s.SortBy = &v
	return s
}

func (s *SearchEditingProjectRequest) SetNextToken(v string) *SearchEditingProjectRequest {
	s.NextToken = &v
	return s
}

func (s *SearchEditingProjectRequest) SetMaxResults(v int64) *SearchEditingProjectRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchEditingProjectRequest) SetPageNo(v int64) *SearchEditingProjectRequest {
	s.PageNo = &v
	return s
}

func (s *SearchEditingProjectRequest) SetPageSize(v int64) *SearchEditingProjectRequest {
	s.PageSize = &v
	return s
}

type SearchEditingProjectResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 云剪辑工程列表
	ProjectList []*SearchEditingProjectResponseBodyProjectList `json:"ProjectList,omitempty" xml:"ProjectList,omitempty" type:"Repeated"`
	// 云剪辑工程总数
	MaxResults *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s SearchEditingProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectResponseBody) SetRequestId(v string) *SearchEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchEditingProjectResponseBody) SetProjectList(v []*SearchEditingProjectResponseBodyProjectList) *SearchEditingProjectResponseBody {
	s.ProjectList = v
	return s
}

func (s *SearchEditingProjectResponseBody) SetMaxResults(v int64) *SearchEditingProjectResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchEditingProjectResponseBody) SetTotalCount(v int64) *SearchEditingProjectResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchEditingProjectResponseBody) SetNextToken(v string) *SearchEditingProjectResponseBody {
	s.NextToken = &v
	return s
}

type SearchEditingProjectResponseBodyProjectList struct {
	// 云剪辑工程ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 云剪辑工程标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 云剪辑工程时间线
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// 云剪辑工程描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 云剪辑工程封面
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 云剪辑工程创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 云剪辑工程最新修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// 云剪辑工程总时长
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 云剪辑工程状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SearchEditingProjectResponseBodyProjectList) String() string {
	return tea.Prettify(s)
}

func (s SearchEditingProjectResponseBodyProjectList) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectResponseBodyProjectList) SetProjectId(v string) *SearchEditingProjectResponseBodyProjectList {
	s.ProjectId = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetTitle(v string) *SearchEditingProjectResponseBodyProjectList {
	s.Title = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetTimeline(v string) *SearchEditingProjectResponseBodyProjectList {
	s.Timeline = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetDescription(v string) *SearchEditingProjectResponseBodyProjectList {
	s.Description = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetCoverURL(v string) *SearchEditingProjectResponseBodyProjectList {
	s.CoverURL = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetCreateTime(v string) *SearchEditingProjectResponseBodyProjectList {
	s.CreateTime = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetModifiedTime(v string) *SearchEditingProjectResponseBodyProjectList {
	s.ModifiedTime = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetDuration(v int64) *SearchEditingProjectResponseBodyProjectList {
	s.Duration = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetStatus(v string) *SearchEditingProjectResponseBodyProjectList {
	s.Status = &v
	return s
}

type SearchEditingProjectResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchEditingProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectResponse) SetHeaders(v map[string]*string) *SearchEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *SearchEditingProjectResponse) SetBody(v *SearchEditingProjectResponseBody) *SearchEditingProjectResponse {
	s.Body = v
	return s
}

type SubmitMediaProducingJobRequest struct {
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Timeline          *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	TemplateId        *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	ClipsParam        *string `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	ProjectMetadata   *string `json:"ProjectMetadata,omitempty" xml:"ProjectMetadata,omitempty"`
	OutputMediaTarget *string `json:"OutputMediaTarget,omitempty" xml:"OutputMediaTarget,omitempty"`
	OutputMediaConfig *string `json:"OutputMediaConfig,omitempty" xml:"OutputMediaConfig,omitempty"`
	UserData          *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s SubmitMediaProducingJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitMediaProducingJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitMediaProducingJobRequest) SetProjectId(v string) *SubmitMediaProducingJobRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetTimeline(v string) *SubmitMediaProducingJobRequest {
	s.Timeline = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetTemplateId(v string) *SubmitMediaProducingJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetClipsParam(v string) *SubmitMediaProducingJobRequest {
	s.ClipsParam = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetProjectMetadata(v string) *SubmitMediaProducingJobRequest {
	s.ProjectMetadata = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetOutputMediaTarget(v string) *SubmitMediaProducingJobRequest {
	s.OutputMediaTarget = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetOutputMediaConfig(v string) *SubmitMediaProducingJobRequest {
	s.OutputMediaConfig = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetUserData(v string) *SubmitMediaProducingJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetClientToken(v string) *SubmitMediaProducingJobRequest {
	s.ClientToken = &v
	return s
}

type SubmitMediaProducingJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 剪辑工程Id
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 合成作业Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitMediaProducingJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitMediaProducingJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitMediaProducingJobResponseBody) SetRequestId(v string) *SubmitMediaProducingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitMediaProducingJobResponseBody) SetProjectId(v string) *SubmitMediaProducingJobResponseBody {
	s.ProjectId = &v
	return s
}

func (s *SubmitMediaProducingJobResponseBody) SetJobId(v string) *SubmitMediaProducingJobResponseBody {
	s.JobId = &v
	return s
}

type SubmitMediaProducingJobResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitMediaProducingJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitMediaProducingJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitMediaProducingJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitMediaProducingJobResponse) SetHeaders(v map[string]*string) *SubmitMediaProducingJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitMediaProducingJobResponse) SetBody(v *SubmitMediaProducingJobResponseBody) *SubmitMediaProducingJobResponse {
	s.Body = v
	return s
}

type UpdateEditingProjectRequest struct {
	// 云剪辑工程标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 云剪辑工程描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 云剪辑工程时间线，Json格式
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// 云剪辑工程封面
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 云剪辑工程ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 前端编辑器工程数据结构
	FEExtend *string `json:"FEExtend,omitempty" xml:"FEExtend,omitempty"`
}

func (s UpdateEditingProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateEditingProjectRequest) SetTitle(v string) *UpdateEditingProjectRequest {
	s.Title = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetDescription(v string) *UpdateEditingProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetTimeline(v string) *UpdateEditingProjectRequest {
	s.Timeline = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetCoverURL(v string) *UpdateEditingProjectRequest {
	s.CoverURL = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetProjectId(v string) *UpdateEditingProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetFEExtend(v string) *UpdateEditingProjectRequest {
	s.FEExtend = &v
	return s
}

type UpdateEditingProjectResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEditingProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEditingProjectResponseBody) SetRequestId(v string) *UpdateEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

type UpdateEditingProjectResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEditingProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateEditingProjectResponse) SetHeaders(v map[string]*string) *UpdateEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateEditingProjectResponse) SetBody(v *UpdateEditingProjectResponseBody) *UpdateEditingProjectResponse {
	s.Body = v
	return s
}

type UpdateMediaInfoRequest struct {
	// 媒资Id
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// 媒资媒体类型
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 标签,如果有多个标签用逗号隔开
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// 封面图，仅视频媒资有效
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 截图
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// 用户自定义元数据
	DynamicMetaDataList *string `json:"DynamicMetaDataList,omitempty" xml:"DynamicMetaDataList,omitempty"`
	// 用户数据，最大1024字节
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 是否以append的形式更新Tags字段
	AppendTags *bool `json:"AppendTags,omitempty" xml:"AppendTags,omitempty"`
	// 是否以append的形式更新Snapshots字段
	AppendSnapshots *bool `json:"AppendSnapshots,omitempty" xml:"AppendSnapshots,omitempty"`
	// 是否以append的形式更新DynamicMetaDataList字段
	AppendDynamicMeta *bool `json:"AppendDynamicMeta,omitempty" xml:"AppendDynamicMeta,omitempty"`
}

func (s UpdateMediaInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMediaInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaInfoRequest) SetMediaId(v string) *UpdateMediaInfoRequest {
	s.MediaId = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetInputURL(v string) *UpdateMediaInfoRequest {
	s.InputURL = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetBusinessType(v string) *UpdateMediaInfoRequest {
	s.BusinessType = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetTitle(v string) *UpdateMediaInfoRequest {
	s.Title = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetDescription(v string) *UpdateMediaInfoRequest {
	s.Description = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetCategory(v string) *UpdateMediaInfoRequest {
	s.Category = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetMediaTags(v string) *UpdateMediaInfoRequest {
	s.MediaTags = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetCoverURL(v string) *UpdateMediaInfoRequest {
	s.CoverURL = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetSnapshots(v string) *UpdateMediaInfoRequest {
	s.Snapshots = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetDynamicMetaDataList(v string) *UpdateMediaInfoRequest {
	s.DynamicMetaDataList = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetUserData(v string) *UpdateMediaInfoRequest {
	s.UserData = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetAppendTags(v bool) *UpdateMediaInfoRequest {
	s.AppendTags = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetAppendSnapshots(v bool) *UpdateMediaInfoRequest {
	s.AppendSnapshots = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetAppendDynamicMeta(v bool) *UpdateMediaInfoRequest {
	s.AppendDynamicMeta = &v
	return s
}

type UpdateMediaInfoResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// ICE媒资ID
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s UpdateMediaInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMediaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaInfoResponseBody) SetRequestId(v string) *UpdateMediaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaInfoResponseBody) SetMediaId(v string) *UpdateMediaInfoResponseBody {
	s.MediaId = &v
	return s
}

type UpdateMediaInfoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateMediaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMediaInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMediaInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaInfoResponse) SetHeaders(v map[string]*string) *UpdateMediaInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaInfoResponse) SetBody(v *UpdateMediaInfoResponseBody) *UpdateMediaInfoResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("ice.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("ice.aliyuncs.com"),
		"ap-south-1":                  tea.String("ice.aliyuncs.com"),
		"ap-southeast-1":              tea.String("ice.aliyuncs.com"),
		"ap-southeast-2":              tea.String("ice.aliyuncs.com"),
		"ap-southeast-3":              tea.String("ice.aliyuncs.com"),
		"ap-southeast-5":              tea.String("ice.aliyuncs.com"),
		"cn-beijing":                  tea.String("ice.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("ice.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("ice.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("ice.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("ice.aliyuncs.com"),
		"cn-chengdu":                  tea.String("ice.aliyuncs.com"),
		"cn-edge-1":                   tea.String("ice.aliyuncs.com"),
		"cn-fujian":                   tea.String("ice.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("ice.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("ice.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("ice.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("ice.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("ice.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("ice.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("ice.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("ice.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("ice.aliyuncs.com"),
		"cn-hongkong":                 tea.String("ice.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("ice.aliyuncs.com"),
		"cn-huhehaote":                tea.String("ice.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("ice.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("ice.aliyuncs.com"),
		"cn-qingdao":                  tea.String("ice.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("ice.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("ice.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("ice.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("ice.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("ice.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("ice.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("ice.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("ice.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("ice.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("ice.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("ice.aliyuncs.com"),
		"cn-wuhan":                    tea.String("ice.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("ice.aliyuncs.com"),
		"cn-yushanfang":               tea.String("ice.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("ice.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("ice.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("ice.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("ice.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("ice.aliyuncs.com"),
		"eu-central-1":                tea.String("ice.aliyuncs.com"),
		"eu-west-1":                   tea.String("ice.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("ice.aliyuncs.com"),
		"me-east-1":                   tea.String("ice.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("ice.aliyuncs.com"),
		"us-east-1":                   tea.String("ice.aliyuncs.com"),
		"us-west-1":                   tea.String("ice.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ice"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateEditingProjectWithOptions(request *CreateEditingProjectRequest, runtime *util.RuntimeOptions) (_result *CreateEditingProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateEditingProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateEditingProject"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEditingProject(request *CreateEditingProjectRequest) (_result *CreateEditingProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEditingProjectResponse{}
	_body, _err := client.CreateEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEditingProjectsWithOptions(request *DeleteEditingProjectsRequest, runtime *util.RuntimeOptions) (_result *DeleteEditingProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteEditingProjectsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteEditingProjects"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEditingProjects(request *DeleteEditingProjectsRequest) (_result *DeleteEditingProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEditingProjectsResponse{}
	_body, _err := client.DeleteEditingProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMediaInfosWithOptions(request *DeleteMediaInfosRequest, runtime *util.RuntimeOptions) (_result *DeleteMediaInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMediaInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMediaInfos"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMediaInfos(request *DeleteMediaInfosRequest) (_result *DeleteMediaInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMediaInfosResponse{}
	_body, _err := client.DeleteMediaInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIceProductStatusWithOptions(runtime *util.RuntimeOptions) (_result *DescribeIceProductStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeIceProductStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeIceProductStatus"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIceProductStatus() (_result *DescribeIceProductStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIceProductStatusResponse{}
	_body, _err := client.DescribeIceProductStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRelatedAuthorizationStatusWithOptions(runtime *util.RuntimeOptions) (_result *DescribeRelatedAuthorizationStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeRelatedAuthorizationStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRelatedAuthorizationStatus"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRelatedAuthorizationStatus() (_result *DescribeRelatedAuthorizationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRelatedAuthorizationStatusResponse{}
	_body, _err := client.DescribeRelatedAuthorizationStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEditingProjectWithOptions(request *GetEditingProjectRequest, runtime *util.RuntimeOptions) (_result *GetEditingProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetEditingProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetEditingProject"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEditingProject(request *GetEditingProjectRequest) (_result *GetEditingProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEditingProjectResponse{}
	_body, _err := client.GetEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaInfoWithOptions(request *GetMediaInfoRequest, runtime *util.RuntimeOptions) (_result *GetMediaInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMediaInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMediaInfo"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaInfo(request *GetMediaInfoRequest) (_result *GetMediaInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMediaInfoResponse{}
	_body, _err := client.GetMediaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaProducingJobWithOptions(request *GetMediaProducingJobRequest, runtime *util.RuntimeOptions) (_result *GetMediaProducingJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetMediaProducingJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMediaProducingJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaProducingJob(request *GetMediaProducingJobRequest) (_result *GetMediaProducingJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMediaProducingJobResponse{}
	_body, _err := client.GetMediaProducingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAllPublicMediaTagsWithOptions(request *ListAllPublicMediaTagsRequest, runtime *util.RuntimeOptions) (_result *ListAllPublicMediaTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAllPublicMediaTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAllPublicMediaTags"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAllPublicMediaTags(request *ListAllPublicMediaTagsRequest) (_result *ListAllPublicMediaTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAllPublicMediaTagsResponse{}
	_body, _err := client.ListAllPublicMediaTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMediaBasicInfosWithOptions(request *ListMediaBasicInfosRequest, runtime *util.RuntimeOptions) (_result *ListMediaBasicInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMediaBasicInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMediaBasicInfos"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMediaBasicInfos(request *ListMediaBasicInfosRequest) (_result *ListMediaBasicInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMediaBasicInfosResponse{}
	_body, _err := client.ListMediaBasicInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMediaProducingJobsWithOptions(request *ListMediaProducingJobsRequest, runtime *util.RuntimeOptions) (_result *ListMediaProducingJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMediaProducingJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMediaProducingJobs"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMediaProducingJobs(request *ListMediaProducingJobsRequest) (_result *ListMediaProducingJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMediaProducingJobsResponse{}
	_body, _err := client.ListMediaProducingJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPublicMediaBasicInfosWithOptions(request *ListPublicMediaBasicInfosRequest, runtime *util.RuntimeOptions) (_result *ListPublicMediaBasicInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPublicMediaBasicInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPublicMediaBasicInfos"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPublicMediaBasicInfos(request *ListPublicMediaBasicInfosRequest) (_result *ListPublicMediaBasicInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublicMediaBasicInfosResponse{}
	_body, _err := client.ListPublicMediaBasicInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterMediaInfoWithOptions(request *RegisterMediaInfoRequest, runtime *util.RuntimeOptions) (_result *RegisterMediaInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegisterMediaInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterMediaInfo"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterMediaInfo(request *RegisterMediaInfoRequest) (_result *RegisterMediaInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterMediaInfoResponse{}
	_body, _err := client.RegisterMediaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchEditingProjectWithOptions(request *SearchEditingProjectRequest, runtime *util.RuntimeOptions) (_result *SearchEditingProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchEditingProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchEditingProject"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchEditingProject(request *SearchEditingProjectRequest) (_result *SearchEditingProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchEditingProjectResponse{}
	_body, _err := client.SearchEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitMediaProducingJobWithOptions(request *SubmitMediaProducingJobRequest, runtime *util.RuntimeOptions) (_result *SubmitMediaProducingJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitMediaProducingJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitMediaProducingJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitMediaProducingJob(request *SubmitMediaProducingJobRequest) (_result *SubmitMediaProducingJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitMediaProducingJobResponse{}
	_body, _err := client.SubmitMediaProducingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEditingProjectWithOptions(request *UpdateEditingProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateEditingProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateEditingProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateEditingProject"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEditingProject(request *UpdateEditingProjectRequest) (_result *UpdateEditingProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEditingProjectResponse{}
	_body, _err := client.UpdateEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMediaInfoWithOptions(request *UpdateMediaInfoRequest, runtime *util.RuntimeOptions) (_result *UpdateMediaInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateMediaInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateMediaInfo"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMediaInfo(request *UpdateMediaInfoRequest) (_result *UpdateMediaInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMediaInfoResponse{}
	_body, _err := client.UpdateMediaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
