// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMdsCubeResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListMdsCubeResourcesResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMdsCubeResourcesResponseBody
	GetResultCode() *string
	SetResultContent(v *ListMdsCubeResourcesResponseBodyResultContent) *ListMdsCubeResourcesResponseBody
	GetResultContent() *ListMdsCubeResourcesResponseBodyResultContent
	SetResultMessage(v string) *ListMdsCubeResourcesResponseBody
	GetResultMessage() *string
}

type ListMdsCubeResourcesResponseBody struct {
	RequestId     *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                        `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *ListMdsCubeResourcesResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                        `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMdsCubeResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMdsCubeResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMdsCubeResourcesResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMdsCubeResourcesResponseBody) GetResultContent() *ListMdsCubeResourcesResponseBodyResultContent {
	return s.ResultContent
}

func (s *ListMdsCubeResourcesResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMdsCubeResourcesResponseBody) SetRequestId(v string) *ListMdsCubeResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBody) SetResultCode(v string) *ListMdsCubeResourcesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBody) SetResultContent(v *ListMdsCubeResourcesResponseBodyResultContent) *ListMdsCubeResourcesResponseBody {
	s.ResultContent = v
	return s
}

func (s *ListMdsCubeResourcesResponseBody) SetResultMessage(v string) *ListMdsCubeResourcesResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMdsCubeResourcesResponseBodyResultContent struct {
	Data      *ListMdsCubeResourcesResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMdsCubeResourcesResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeResourcesResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *ListMdsCubeResourcesResponseBodyResultContent) GetData() *ListMdsCubeResourcesResponseBodyResultContentData {
	return s.Data
}

func (s *ListMdsCubeResourcesResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMdsCubeResourcesResponseBodyResultContent) SetData(v *ListMdsCubeResourcesResponseBodyResultContentData) *ListMdsCubeResourcesResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContent) SetRequestId(v string) *ListMdsCubeResourcesResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMdsCubeResourcesResponseBodyResultContentData struct {
	Content   *ListMdsCubeResourcesResponseBodyResultContentDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	ErrorCode *string                                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg *string                                                   `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMdsCubeResourcesResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeResourcesResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *ListMdsCubeResourcesResponseBodyResultContentData) GetContent() *ListMdsCubeResourcesResponseBodyResultContentDataContent {
	return s.Content
}

func (s *ListMdsCubeResourcesResponseBodyResultContentData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMdsCubeResourcesResponseBodyResultContentData) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMdsCubeResourcesResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListMdsCubeResourcesResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *ListMdsCubeResourcesResponseBodyResultContentData) SetContent(v *ListMdsCubeResourcesResponseBodyResultContentDataContent) *ListMdsCubeResourcesResponseBodyResultContentData {
	s.Content = v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentData) SetErrorCode(v string) *ListMdsCubeResourcesResponseBodyResultContentData {
	s.ErrorCode = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentData) SetRequestId(v string) *ListMdsCubeResourcesResponseBodyResultContentData {
	s.RequestId = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentData) SetResultMsg(v string) *ListMdsCubeResourcesResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentData) SetSuccess(v bool) *ListMdsCubeResourcesResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentData) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMdsCubeResourcesResponseBodyResultContentDataContent struct {
	CurrentMaxAndroidVersion *string                                                         `json:"CurrentMaxAndroidVersion,omitempty" xml:"CurrentMaxAndroidVersion,omitempty"`
	CurrentMaxIosVersion     *string                                                         `json:"CurrentMaxIosVersion,omitempty" xml:"CurrentMaxIosVersion,omitempty"`
	FirstPage                *bool                                                           `json:"FirstPage,omitempty" xml:"FirstPage,omitempty"`
	FirstResult              *int64                                                          `json:"FirstResult,omitempty" xml:"FirstResult,omitempty"`
	LastPage                 *bool                                                           `json:"LastPage,omitempty" xml:"LastPage,omitempty"`
	List                     []*ListMdsCubeResourcesResponseBodyResultContentDataContentList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	NextPage                 *int64                                                          `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageNo                   *int32                                                          `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize                 *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage                  *int64                                                          `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	TotalCount               *int32                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMdsCubeResourcesResponseBodyResultContentDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeResourcesResponseBodyResultContentDataContent) GoString() string {
	return s.String()
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) GetCurrentMaxAndroidVersion() *string {
	return s.CurrentMaxAndroidVersion
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) GetCurrentMaxIosVersion() *string {
	return s.CurrentMaxIosVersion
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) GetFirstPage() *bool {
	return s.FirstPage
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) GetFirstResult() *int64 {
	return s.FirstResult
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) GetLastPage() *bool {
	return s.LastPage
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) GetList() []*ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	return s.List
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) GetNextPage() *int64 {
	return s.NextPage
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) GetPrePage() *int64 {
	return s.PrePage
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) SetCurrentMaxAndroidVersion(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContent {
	s.CurrentMaxAndroidVersion = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) SetCurrentMaxIosVersion(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContent {
	s.CurrentMaxIosVersion = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) SetFirstPage(v bool) *ListMdsCubeResourcesResponseBodyResultContentDataContent {
	s.FirstPage = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) SetFirstResult(v int64) *ListMdsCubeResourcesResponseBodyResultContentDataContent {
	s.FirstResult = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) SetLastPage(v bool) *ListMdsCubeResourcesResponseBodyResultContentDataContent {
	s.LastPage = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) SetList(v []*ListMdsCubeResourcesResponseBodyResultContentDataContentList) *ListMdsCubeResourcesResponseBodyResultContentDataContent {
	s.List = v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) SetNextPage(v int64) *ListMdsCubeResourcesResponseBodyResultContentDataContent {
	s.NextPage = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) SetPageNo(v int32) *ListMdsCubeResourcesResponseBodyResultContentDataContent {
	s.PageNo = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) SetPageSize(v int32) *ListMdsCubeResourcesResponseBodyResultContentDataContent {
	s.PageSize = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) SetPrePage(v int64) *ListMdsCubeResourcesResponseBodyResultContentDataContent {
	s.PrePage = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) SetTotalCount(v int32) *ListMdsCubeResourcesResponseBodyResultContentDataContent {
	s.TotalCount = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContent) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMdsCubeResourcesResponseBodyResultContentDataContentList struct {
	AndroidMaxVersion       *string `json:"AndroidMaxVersion,omitempty" xml:"AndroidMaxVersion,omitempty"`
	AndroidMinVersion       *string `json:"AndroidMinVersion,omitempty" xml:"AndroidMinVersion,omitempty"`
	AppCode                 *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	BinFileMd5              *string `json:"BinFileMd5,omitempty" xml:"BinFileMd5,omitempty"`
	BinPrivateFileUrl       *string `json:"BinPrivateFileUrl,omitempty" xml:"BinPrivateFileUrl,omitempty"`
	BinPublicFileUrl        *string `json:"BinPublicFileUrl,omitempty" xml:"BinPublicFileUrl,omitempty"`
	ExtendInfo              *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	GmtCreate               *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified             *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IosMaxVersion           *string `json:"IosMaxVersion,omitempty" xml:"IosMaxVersion,omitempty"`
	IosMinVersion           *string `json:"IosMinVersion,omitempty" xml:"IosMinVersion,omitempty"`
	JsonPrivateFileUrl      *string `json:"JsonPrivateFileUrl,omitempty" xml:"JsonPrivateFileUrl,omitempty"`
	JsonPublicFileUrl       *string `json:"JsonPublicFileUrl,omitempty" xml:"JsonPublicFileUrl,omitempty"`
	MinCubeSdkVersion       *string `json:"MinCubeSdkVersion,omitempty" xml:"MinCubeSdkVersion,omitempty"`
	MockDataDownloadUrl     *string `json:"MockDataDownloadUrl,omitempty" xml:"MockDataDownloadUrl,omitempty"`
	Operator                *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Platform                *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	PreviewPictureUrl       *string `json:"PreviewPictureUrl,omitempty" xml:"PreviewPictureUrl,omitempty"`
	ResourceStatus          *int64  `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	Status                  *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateId              *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateResourceDesc    *string `json:"TemplateResourceDesc,omitempty" xml:"TemplateResourceDesc,omitempty"`
	TemplateResourceVersion *string `json:"TemplateResourceVersion,omitempty" xml:"TemplateResourceVersion,omitempty"`
}

func (s ListMdsCubeResourcesResponseBodyResultContentDataContentList) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeResourcesResponseBodyResultContentDataContentList) GoString() string {
	return s.String()
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetAndroidMaxVersion() *string {
	return s.AndroidMaxVersion
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetAndroidMinVersion() *string {
	return s.AndroidMinVersion
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetAppCode() *string {
	return s.AppCode
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetBinFileMd5() *string {
	return s.BinFileMd5
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetBinPrivateFileUrl() *string {
	return s.BinPrivateFileUrl
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetBinPublicFileUrl() *string {
	return s.BinPublicFileUrl
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetId() *int64 {
	return s.Id
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetIosMaxVersion() *string {
	return s.IosMaxVersion
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetIosMinVersion() *string {
	return s.IosMinVersion
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetJsonPrivateFileUrl() *string {
	return s.JsonPrivateFileUrl
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetJsonPublicFileUrl() *string {
	return s.JsonPublicFileUrl
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetMinCubeSdkVersion() *string {
	return s.MinCubeSdkVersion
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetMockDataDownloadUrl() *string {
	return s.MockDataDownloadUrl
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetOperator() *string {
	return s.Operator
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetPlatform() *string {
	return s.Platform
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetPreviewPictureUrl() *string {
	return s.PreviewPictureUrl
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetResourceStatus() *int64 {
	return s.ResourceStatus
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetStatus() *int32 {
	return s.Status
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetTemplateResourceDesc() *string {
	return s.TemplateResourceDesc
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) GetTemplateResourceVersion() *string {
	return s.TemplateResourceVersion
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetAndroidMaxVersion(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.AndroidMaxVersion = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetAndroidMinVersion(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.AndroidMinVersion = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetAppCode(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.AppCode = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetBinFileMd5(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.BinFileMd5 = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetBinPrivateFileUrl(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.BinPrivateFileUrl = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetBinPublicFileUrl(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.BinPublicFileUrl = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetExtendInfo(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.ExtendInfo = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetGmtCreate(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.GmtCreate = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetGmtModified(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.GmtModified = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetId(v int64) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.Id = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetIosMaxVersion(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.IosMaxVersion = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetIosMinVersion(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.IosMinVersion = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetJsonPrivateFileUrl(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.JsonPrivateFileUrl = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetJsonPublicFileUrl(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.JsonPublicFileUrl = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetMinCubeSdkVersion(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.MinCubeSdkVersion = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetMockDataDownloadUrl(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.MockDataDownloadUrl = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetOperator(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.Operator = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetPlatform(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.Platform = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetPreviewPictureUrl(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.PreviewPictureUrl = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetResourceStatus(v int64) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.ResourceStatus = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetStatus(v int32) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.Status = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetTemplateId(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.TemplateId = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetTemplateResourceDesc(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.TemplateResourceDesc = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) SetTemplateResourceVersion(v string) *ListMdsCubeResourcesResponseBodyResultContentDataContentList {
	s.TemplateResourceVersion = &v
	return s
}

func (s *ListMdsCubeResourcesResponseBodyResultContentDataContentList) Validate() error {
	return dara.Validate(s)
}
