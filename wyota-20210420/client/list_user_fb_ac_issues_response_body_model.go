// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserFbAcIssuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUserFbAcIssuesResponseBody
	GetCode() *string
	SetData(v *ListUserFbAcIssuesResponseBodyData) *ListUserFbAcIssuesResponseBody
	GetData() *ListUserFbAcIssuesResponseBodyData
	SetMessage(v string) *ListUserFbAcIssuesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListUserFbAcIssuesResponseBody
	GetRequestId() *string
}

type ListUserFbAcIssuesResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListUserFbAcIssuesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserFbAcIssuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserFbAcIssuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserFbAcIssuesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUserFbAcIssuesResponseBody) GetData() *ListUserFbAcIssuesResponseBodyData {
	return s.Data
}

func (s *ListUserFbAcIssuesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUserFbAcIssuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserFbAcIssuesResponseBody) SetCode(v string) *ListUserFbAcIssuesResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBody) SetData(v *ListUserFbAcIssuesResponseBodyData) *ListUserFbAcIssuesResponseBody {
	s.Data = v
	return s
}

func (s *ListUserFbAcIssuesResponseBody) SetMessage(v string) *ListUserFbAcIssuesResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBody) SetRequestId(v string) *ListUserFbAcIssuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUserFbAcIssuesResponseBodyData struct {
	IssueDataList []*ListUserFbAcIssuesResponseBodyDataIssueDataList `json:"IssueDataList,omitempty" xml:"IssueDataList,omitempty" type:"Repeated"`
	TotalCount    *int64                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserFbAcIssuesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUserFbAcIssuesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserFbAcIssuesResponseBodyData) GetIssueDataList() []*ListUserFbAcIssuesResponseBodyDataIssueDataList {
	return s.IssueDataList
}

func (s *ListUserFbAcIssuesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUserFbAcIssuesResponseBodyData) SetIssueDataList(v []*ListUserFbAcIssuesResponseBodyDataIssueDataList) *ListUserFbAcIssuesResponseBodyData {
	s.IssueDataList = v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyData) SetTotalCount(v int64) *ListUserFbAcIssuesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListUserFbAcIssuesResponseBodyDataIssueDataList struct {
	Account       *string                                                    `json:"Account,omitempty" xml:"Account,omitempty"`
	ClientVersion *string                                                    `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ErrorMessage  *string                                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FileList      []*ListUserFbAcIssuesResponseBodyDataIssueDataListFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	GmtCreated    *string                                                    `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified   *string                                                    `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InstanceId    *string                                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IssueId       *int64                                                     `json:"IssueId,omitempty" xml:"IssueId,omitempty"`
	Label         *string                                                    `json:"Label,omitempty" xml:"Label,omitempty"`
	ReservedA     *string                                                    `json:"ReservedA,omitempty" xml:"ReservedA,omitempty"`
	ReservedB     *string                                                    `json:"ReservedB,omitempty" xml:"ReservedB,omitempty"`
	UserEmail     *string                                                    `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
}

func (s ListUserFbAcIssuesResponseBodyDataIssueDataList) String() string {
	return dara.Prettify(s)
}

func (s ListUserFbAcIssuesResponseBodyDataIssueDataList) GoString() string {
	return s.String()
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) GetAccount() *string {
	return s.Account
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) GetFileList() []*ListUserFbAcIssuesResponseBodyDataIssueDataListFileList {
	return s.FileList
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) GetIssueId() *int64 {
	return s.IssueId
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) GetLabel() *string {
	return s.Label
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) GetReservedA() *string {
	return s.ReservedA
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) GetReservedB() *string {
	return s.ReservedB
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) GetUserEmail() *string {
	return s.UserEmail
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetAccount(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.Account = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetClientVersion(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.ClientVersion = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetErrorMessage(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.ErrorMessage = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetFileList(v []*ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.FileList = v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetGmtCreated(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.GmtCreated = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetGmtModified(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.GmtModified = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetInstanceId(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.InstanceId = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetIssueId(v int64) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.IssueId = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetLabel(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.Label = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetReservedA(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.ReservedA = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetReservedB(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.ReservedB = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) SetUserEmail(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataList {
	s.UserEmail = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataList) Validate() error {
	return dara.Validate(s)
}

type ListUserFbAcIssuesResponseBodyDataIssueDataListFileList struct {
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileSize  *int32  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileType  *int32  `json:"FileType,omitempty" xml:"FileType,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) String() string {
	return dara.Prettify(s)
}

func (s ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) GoString() string {
	return s.String()
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) GetFileName() *string {
	return s.FileName
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) GetFileSize() *int32 {
	return s.FileSize
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) GetFileType() *int32 {
	return s.FileType
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) GetSessionId() *string {
	return s.SessionId
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) SetFileName(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList {
	s.FileName = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) SetFileSize(v int32) *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList {
	s.FileSize = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) SetFileType(v int32) *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList {
	s.FileType = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) SetSessionId(v string) *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList {
	s.SessionId = &v
	return s
}

func (s *ListUserFbAcIssuesResponseBodyDataIssueDataListFileList) Validate() error {
	return dara.Validate(s)
}
