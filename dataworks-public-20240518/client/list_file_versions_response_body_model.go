// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListFileVersionsResponseBodyData) *ListFileVersionsResponseBody
	GetData() *ListFileVersionsResponseBodyData
	SetErrorCode(v string) *ListFileVersionsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListFileVersionsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListFileVersionsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListFileVersionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListFileVersionsResponseBody
	GetSuccess() *bool
}

type ListFileVersionsResponseBody struct {
	// The list of file versions.
	Data *ListFileVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call succeeded. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFileVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFileVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileVersionsResponseBody) GetData() *ListFileVersionsResponseBodyData {
	return s.Data
}

func (s *ListFileVersionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListFileVersionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListFileVersionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListFileVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFileVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFileVersionsResponseBody) SetData(v *ListFileVersionsResponseBodyData) *ListFileVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListFileVersionsResponseBody) SetErrorCode(v string) *ListFileVersionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListFileVersionsResponseBody) SetErrorMessage(v string) *ListFileVersionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListFileVersionsResponseBody) SetHttpStatusCode(v int32) *ListFileVersionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListFileVersionsResponseBody) SetRequestId(v string) *ListFileVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileVersionsResponseBody) SetSuccess(v bool) *ListFileVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListFileVersionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFileVersionsResponseBodyData struct {
	// The list of file version details.
	FileVersions []*ListFileVersionsResponseBodyDataFileVersions `json:"FileVersions,omitempty" xml:"FileVersions,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on the current page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 13
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFileVersionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFileVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFileVersionsResponseBodyData) GetFileVersions() []*ListFileVersionsResponseBodyDataFileVersions {
	return s.FileVersions
}

func (s *ListFileVersionsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFileVersionsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFileVersionsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFileVersionsResponseBodyData) SetFileVersions(v []*ListFileVersionsResponseBodyDataFileVersions) *ListFileVersionsResponseBodyData {
	s.FileVersions = v
	return s
}

func (s *ListFileVersionsResponseBodyData) SetPageNumber(v int32) *ListFileVersionsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListFileVersionsResponseBodyData) SetPageSize(v int32) *ListFileVersionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListFileVersionsResponseBodyData) SetTotalCount(v int32) *ListFileVersionsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListFileVersionsResponseBodyData) Validate() error {
	if s.FileVersions != nil {
		for _, item := range s.FileVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFileVersionsResponseBodyDataFileVersions struct {
	// The change type for this file version. Valid values: CREATE, UPDATE, and DELETE.
	//
	// example:
	//
	// UPDATE
	ChangeType *string `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// The description of this file version.
	//
	// example:
	//
	// Second version submission
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The timestamp (in milliseconds) when the file version was created.
	//
	// example:
	//
	// 1593881265000
	CommitTime *int64 `json:"CommitTime,omitempty" xml:"CommitTime,omitempty"`
	// The Alibaba Cloud account ID of the user who created this file version.
	//
	// example:
	//
	// 73842342****
	CommitUser *string `json:"CommitUser,omitempty" xml:"CommitUser,omitempty"`
	// The file code for this version.
	//
	// example:
	//
	// SHOW TABLES;
	FileContent *string `json:"FileContent,omitempty" xml:"FileContent,omitempty"`
	// The file name for this file version.
	//
	// example:
	//
	// ods_user_info_d
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The text information for this file version.
	//
	// example:
	//
	// {"fileName":"ods_user_info_d","fileType":10}
	FilePropertyContent *string `json:"FilePropertyContent,omitempty" xml:"FilePropertyContent,omitempty"`
	// The file version.
	//
	// example:
	//
	// 2
	FileVersion *int32 `json:"FileVersion,omitempty" xml:"FileVersion,omitempty"`
	// Indicates whether this file version is the latest version in the production environment.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsCurrentProd *bool `json:"IsCurrentProd,omitempty" xml:"IsCurrentProd,omitempty"`
	// The scheduling configuration for this file version.
	//
	// example:
	//
	// {"cycleType":0,"cronExpress":"00 05 00 	- 	- ?"}
	NodeContent *string `json:"NodeContent,omitempty" xml:"NodeContent,omitempty"`
	// The scheduling task ID associated with this file version.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The current status of the file version. Valid values: COMMITTING (committing), COMMITTED or CHECK_OK (committed), PACKAGED (ready for deployment), DEPLOYING (deploying), DEPLOYED (deployed), and CANCELLED (deployment canceled).
	//
	// example:
	//
	// COMMITTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The functional module to which the file belongs. Valid values: NORMAL (Data Studio), MANUAL (manual task), MANUAL_BIZ (manual workflow), SKIP (dry-run scheduling in Data Studio), ADHOCQUERY (ad hoc query), and COMPONENT (component management).
	//
	// example:
	//
	// NORMAL
	UseType *string `json:"UseType,omitempty" xml:"UseType,omitempty"`
}

func (s ListFileVersionsResponseBodyDataFileVersions) String() string {
	return dara.Prettify(s)
}

func (s ListFileVersionsResponseBodyDataFileVersions) GoString() string {
	return s.String()
}

func (s *ListFileVersionsResponseBodyDataFileVersions) GetChangeType() *string {
	return s.ChangeType
}

func (s *ListFileVersionsResponseBodyDataFileVersions) GetComment() *string {
	return s.Comment
}

func (s *ListFileVersionsResponseBodyDataFileVersions) GetCommitTime() *int64 {
	return s.CommitTime
}

func (s *ListFileVersionsResponseBodyDataFileVersions) GetCommitUser() *string {
	return s.CommitUser
}

func (s *ListFileVersionsResponseBodyDataFileVersions) GetFileContent() *string {
	return s.FileContent
}

func (s *ListFileVersionsResponseBodyDataFileVersions) GetFileName() *string {
	return s.FileName
}

func (s *ListFileVersionsResponseBodyDataFileVersions) GetFilePropertyContent() *string {
	return s.FilePropertyContent
}

func (s *ListFileVersionsResponseBodyDataFileVersions) GetFileVersion() *int32 {
	return s.FileVersion
}

func (s *ListFileVersionsResponseBodyDataFileVersions) GetIsCurrentProd() *bool {
	return s.IsCurrentProd
}

func (s *ListFileVersionsResponseBodyDataFileVersions) GetNodeContent() *string {
	return s.NodeContent
}

func (s *ListFileVersionsResponseBodyDataFileVersions) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListFileVersionsResponseBodyDataFileVersions) GetStatus() *string {
	return s.Status
}

func (s *ListFileVersionsResponseBodyDataFileVersions) GetUseType() *string {
	return s.UseType
}

func (s *ListFileVersionsResponseBodyDataFileVersions) SetChangeType(v string) *ListFileVersionsResponseBodyDataFileVersions {
	s.ChangeType = &v
	return s
}

func (s *ListFileVersionsResponseBodyDataFileVersions) SetComment(v string) *ListFileVersionsResponseBodyDataFileVersions {
	s.Comment = &v
	return s
}

func (s *ListFileVersionsResponseBodyDataFileVersions) SetCommitTime(v int64) *ListFileVersionsResponseBodyDataFileVersions {
	s.CommitTime = &v
	return s
}

func (s *ListFileVersionsResponseBodyDataFileVersions) SetCommitUser(v string) *ListFileVersionsResponseBodyDataFileVersions {
	s.CommitUser = &v
	return s
}

func (s *ListFileVersionsResponseBodyDataFileVersions) SetFileContent(v string) *ListFileVersionsResponseBodyDataFileVersions {
	s.FileContent = &v
	return s
}

func (s *ListFileVersionsResponseBodyDataFileVersions) SetFileName(v string) *ListFileVersionsResponseBodyDataFileVersions {
	s.FileName = &v
	return s
}

func (s *ListFileVersionsResponseBodyDataFileVersions) SetFilePropertyContent(v string) *ListFileVersionsResponseBodyDataFileVersions {
	s.FilePropertyContent = &v
	return s
}

func (s *ListFileVersionsResponseBodyDataFileVersions) SetFileVersion(v int32) *ListFileVersionsResponseBodyDataFileVersions {
	s.FileVersion = &v
	return s
}

func (s *ListFileVersionsResponseBodyDataFileVersions) SetIsCurrentProd(v bool) *ListFileVersionsResponseBodyDataFileVersions {
	s.IsCurrentProd = &v
	return s
}

func (s *ListFileVersionsResponseBodyDataFileVersions) SetNodeContent(v string) *ListFileVersionsResponseBodyDataFileVersions {
	s.NodeContent = &v
	return s
}

func (s *ListFileVersionsResponseBodyDataFileVersions) SetNodeId(v int64) *ListFileVersionsResponseBodyDataFileVersions {
	s.NodeId = &v
	return s
}

func (s *ListFileVersionsResponseBodyDataFileVersions) SetStatus(v string) *ListFileVersionsResponseBodyDataFileVersions {
	s.Status = &v
	return s
}

func (s *ListFileVersionsResponseBodyDataFileVersions) SetUseType(v string) *ListFileVersionsResponseBodyDataFileVersions {
	s.UseType = &v
	return s
}

func (s *ListFileVersionsResponseBodyDataFileVersions) Validate() error {
	return dara.Validate(s)
}
