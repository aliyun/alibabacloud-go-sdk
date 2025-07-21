// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserFbIssuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUserFbIssuesResponseBody
	GetCode() *string
	SetData(v *ListUserFbIssuesResponseBodyData) *ListUserFbIssuesResponseBody
	GetData() *ListUserFbIssuesResponseBodyData
	SetMessage(v string) *ListUserFbIssuesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListUserFbIssuesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserFbIssuesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListUserFbIssuesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListUserFbIssuesResponseBody
	GetTotalCount() *int32
}

type ListUserFbIssuesResponseBody struct {
	Code       *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *ListUserFbIssuesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserFbIssuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserFbIssuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserFbIssuesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUserFbIssuesResponseBody) GetData() *ListUserFbIssuesResponseBodyData {
	return s.Data
}

func (s *ListUserFbIssuesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUserFbIssuesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserFbIssuesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserFbIssuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserFbIssuesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserFbIssuesResponseBody) SetCode(v string) *ListUserFbIssuesResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserFbIssuesResponseBody) SetData(v *ListUserFbIssuesResponseBodyData) *ListUserFbIssuesResponseBody {
	s.Data = v
	return s
}

func (s *ListUserFbIssuesResponseBody) SetMessage(v string) *ListUserFbIssuesResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserFbIssuesResponseBody) SetPageNumber(v int32) *ListUserFbIssuesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUserFbIssuesResponseBody) SetPageSize(v int32) *ListUserFbIssuesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserFbIssuesResponseBody) SetRequestId(v string) *ListUserFbIssuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserFbIssuesResponseBody) SetTotalCount(v int32) *ListUserFbIssuesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserFbIssuesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUserFbIssuesResponseBodyData struct {
	Count             *string                                              `json:"Count,omitempty" xml:"Count,omitempty"`
	FeedbackIssueData []*ListUserFbIssuesResponseBodyDataFeedbackIssueData `json:"FeedbackIssueData,omitempty" xml:"FeedbackIssueData,omitempty" type:"Repeated"`
}

func (s ListUserFbIssuesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUserFbIssuesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserFbIssuesResponseBodyData) GetCount() *string {
	return s.Count
}

func (s *ListUserFbIssuesResponseBodyData) GetFeedbackIssueData() []*ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	return s.FeedbackIssueData
}

func (s *ListUserFbIssuesResponseBodyData) SetCount(v string) *ListUserFbIssuesResponseBodyData {
	s.Count = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyData) SetFeedbackIssueData(v []*ListUserFbIssuesResponseBodyDataFeedbackIssueData) *ListUserFbIssuesResponseBodyData {
	s.FeedbackIssueData = v
	return s
}

func (s *ListUserFbIssuesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListUserFbIssuesResponseBodyDataFeedbackIssueData struct {
	AppId       *string                                                      `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClientId    *string                                                      `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientModel *string                                                      `json:"ClientModel,omitempty" xml:"ClientModel,omitempty"`
	ClientSn    *string                                                      `json:"ClientSn,omitempty" xml:"ClientSn,omitempty"`
	CustomerId  *string                                                      `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	Description *string                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopId   *string                                                      `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	ErrorCode   *string                                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg    *string                                                      `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FbType      *int32                                                       `json:"FbType,omitempty" xml:"FbType,omitempty"`
	FileList    []*ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	GmtCreated  *string                                                      `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	IssueId     *int32                                                       `json:"IssueId,omitempty" xml:"IssueId,omitempty"`
	IssueLabel  *string                                                      `json:"IssueLabel,omitempty" xml:"IssueLabel,omitempty"`
	Status      *int32                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	Title       *string                                                      `json:"Title,omitempty" xml:"Title,omitempty"`
	UserEmail   *string                                                      `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
	UserId      *string                                                      `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserFbIssuesResponseBodyDataFeedbackIssueData) String() string {
	return dara.Prettify(s)
}

func (s ListUserFbIssuesResponseBodyDataFeedbackIssueData) GoString() string {
	return s.String()
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetAppId() *string {
	return s.AppId
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetClientId() *string {
	return s.ClientId
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetClientModel() *string {
	return s.ClientModel
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetClientSn() *string {
	return s.ClientSn
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetDescription() *string {
	return s.Description
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetFbType() *int32 {
	return s.FbType
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetFileList() []*ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList {
	return s.FileList
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetIssueId() *int32 {
	return s.IssueId
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetIssueLabel() *string {
	return s.IssueLabel
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetStatus() *int32 {
	return s.Status
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetTitle() *string {
	return s.Title
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetUserEmail() *string {
	return s.UserEmail
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) GetUserId() *string {
	return s.UserId
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetAppId(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.AppId = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetClientId(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.ClientId = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetClientModel(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.ClientModel = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetClientSn(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.ClientSn = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetCustomerId(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.CustomerId = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetDescription(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.Description = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetDesktopId(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.DesktopId = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetErrorCode(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.ErrorCode = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetErrorMsg(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.ErrorMsg = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetFbType(v int32) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.FbType = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetFileList(v []*ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.FileList = v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetGmtCreated(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.GmtCreated = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetIssueId(v int32) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.IssueId = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetIssueLabel(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.IssueLabel = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetStatus(v int32) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.Status = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetTitle(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.Title = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetUserEmail(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.UserEmail = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) SetUserId(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueData {
	s.UserId = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueData) Validate() error {
	return dara.Validate(s)
}

type ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList struct {
	FileMd5  *string `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileSize *int32  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileType *int32  `json:"FileType,omitempty" xml:"FileType,omitempty"`
	OssUrl   *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
}

func (s ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) String() string {
	return dara.Prettify(s)
}

func (s ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) GoString() string {
	return s.String()
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) GetFileMd5() *string {
	return s.FileMd5
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) GetFileName() *string {
	return s.FileName
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) GetFileSize() *int32 {
	return s.FileSize
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) GetFileType() *int32 {
	return s.FileType
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) GetOssUrl() *string {
	return s.OssUrl
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) SetFileMd5(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList {
	s.FileMd5 = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) SetFileName(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList {
	s.FileName = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) SetFileSize(v int32) *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList {
	s.FileSize = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) SetFileType(v int32) *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList {
	s.FileType = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) SetOssUrl(v string) *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList {
	s.OssUrl = &v
	return s
}

func (s *ListUserFbIssuesResponseBodyDataFeedbackIssueDataFileList) Validate() error {
	return dara.Validate(s)
}
