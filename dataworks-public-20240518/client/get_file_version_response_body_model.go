// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetFileVersionResponseBodyData) *GetFileVersionResponseBody
	GetData() *GetFileVersionResponseBodyData
	SetErrorCode(v string) *GetFileVersionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetFileVersionResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetFileVersionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetFileVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFileVersionResponseBody
	GetSuccess() *bool
}

type GetFileVersionResponseBody struct {
	// The version details of the file.
	Data *GetFileVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The unique ID of the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFileVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileVersionResponseBody) GetData() *GetFileVersionResponseBodyData {
	return s.Data
}

func (s *GetFileVersionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetFileVersionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetFileVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetFileVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFileVersionResponseBody) SetData(v *GetFileVersionResponseBodyData) *GetFileVersionResponseBody {
	s.Data = v
	return s
}

func (s *GetFileVersionResponseBody) SetErrorCode(v string) *GetFileVersionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetFileVersionResponseBody) SetErrorMessage(v string) *GetFileVersionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetFileVersionResponseBody) SetHttpStatusCode(v int32) *GetFileVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetFileVersionResponseBody) SetRequestId(v string) *GetFileVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileVersionResponseBody) SetSuccess(v bool) *GetFileVersionResponseBody {
	s.Success = &v
	return s
}

func (s *GetFileVersionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileVersionResponseBodyData struct {
	// The change type of this file version. Valid values: CREATE, UPDATE, and DELETE.
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
	// The timestamp when the file version was generated, in milliseconds.
	//
	// example:
	//
	// 1593881265000
	CommitTime *int64 `json:"CommitTime,omitempty" xml:"CommitTime,omitempty"`
	// The Alibaba Cloud user ID that generated this file version.
	//
	// example:
	//
	// 7384234****
	CommitUser *string `json:"CommitUser,omitempty" xml:"CommitUser,omitempty"`
	// The file code that generated this file version.
	//
	// example:
	//
	// SHOW TABLES;
	FileContent *string `json:"FileContent,omitempty" xml:"FileContent,omitempty"`
	// The name of the file that generated this file version.
	//
	// example:
	//
	// ods_user_info_d
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The basic file information when this file version was generated.
	//
	// example:
	//
	// {"fileName":"ods_user_info_d","fileType":10}
	FilePropertyContent *string `json:"FilePropertyContent,omitempty" xml:"FilePropertyContent,omitempty"`
	// The version of the file.
	//
	// example:
	//
	// 2
	FileVersion *int32 `json:"FileVersion,omitempty" xml:"FileVersion,omitempty"`
	// Indicates whether this file version is the latest version in the production environment. Valid values:
	//
	// - true: The version is the latest version.
	//
	// - false: The version is not the latest version.
	//
	// example:
	//
	// true
	IsCurrentProd *bool `json:"IsCurrentProd,omitempty" xml:"IsCurrentProd,omitempty"`
	// The scheduling configuration when this file version was generated.
	//
	// example:
	//
	// {"cycleType":0,"cronExpress":"00 05 00 	- 	- ?"}
	NodeContent *string `json:"NodeContent,omitempty" xml:"NodeContent,omitempty"`
	// The ID of the scheduling node associated with the file version when it was generated.
	//
	// example:
	//
	// 3000001
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The current status of the file version. Valid values:
	//
	// - COMMITTING: The version is being committed.
	//
	// - COMMITTED or CHECK_OK: The version has been committed.
	//
	// - PACKAGED: The version is ready for deployment.
	//
	// - DEPLOYING: The version is being deployed.
	//
	// - DEPLOYED: The version has been deployed.
	//
	// - CANCELLED: The deployment has been canceled.
	//
	// example:
	//
	// COMMITTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The functional module to which the file belongs. Valid values:
	//
	// - 0: NORMAL (DataStudio)
	//
	// - 1: MANUAL (manual node)
	//
	// - 2: MANUAL_BIZ (manual workflow)
	//
	// - 3: SKIP (dry-run scheduling in DataStudio)
	//
	// - 10: ADHOCQUERY (ad hoc query)
	//
	// - 30: COMPONENT (component management)
	//
	// example:
	//
	// 0
	UseType *string `json:"UseType,omitempty" xml:"UseType,omitempty"`
}

func (s GetFileVersionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFileVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileVersionResponseBodyData) GetChangeType() *string {
	return s.ChangeType
}

func (s *GetFileVersionResponseBodyData) GetComment() *string {
	return s.Comment
}

func (s *GetFileVersionResponseBodyData) GetCommitTime() *int64 {
	return s.CommitTime
}

func (s *GetFileVersionResponseBodyData) GetCommitUser() *string {
	return s.CommitUser
}

func (s *GetFileVersionResponseBodyData) GetFileContent() *string {
	return s.FileContent
}

func (s *GetFileVersionResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetFileVersionResponseBodyData) GetFilePropertyContent() *string {
	return s.FilePropertyContent
}

func (s *GetFileVersionResponseBodyData) GetFileVersion() *int32 {
	return s.FileVersion
}

func (s *GetFileVersionResponseBodyData) GetIsCurrentProd() *bool {
	return s.IsCurrentProd
}

func (s *GetFileVersionResponseBodyData) GetNodeContent() *string {
	return s.NodeContent
}

func (s *GetFileVersionResponseBodyData) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetFileVersionResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetFileVersionResponseBodyData) GetUseType() *string {
	return s.UseType
}

func (s *GetFileVersionResponseBodyData) SetChangeType(v string) *GetFileVersionResponseBodyData {
	s.ChangeType = &v
	return s
}

func (s *GetFileVersionResponseBodyData) SetComment(v string) *GetFileVersionResponseBodyData {
	s.Comment = &v
	return s
}

func (s *GetFileVersionResponseBodyData) SetCommitTime(v int64) *GetFileVersionResponseBodyData {
	s.CommitTime = &v
	return s
}

func (s *GetFileVersionResponseBodyData) SetCommitUser(v string) *GetFileVersionResponseBodyData {
	s.CommitUser = &v
	return s
}

func (s *GetFileVersionResponseBodyData) SetFileContent(v string) *GetFileVersionResponseBodyData {
	s.FileContent = &v
	return s
}

func (s *GetFileVersionResponseBodyData) SetFileName(v string) *GetFileVersionResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetFileVersionResponseBodyData) SetFilePropertyContent(v string) *GetFileVersionResponseBodyData {
	s.FilePropertyContent = &v
	return s
}

func (s *GetFileVersionResponseBodyData) SetFileVersion(v int32) *GetFileVersionResponseBodyData {
	s.FileVersion = &v
	return s
}

func (s *GetFileVersionResponseBodyData) SetIsCurrentProd(v bool) *GetFileVersionResponseBodyData {
	s.IsCurrentProd = &v
	return s
}

func (s *GetFileVersionResponseBodyData) SetNodeContent(v string) *GetFileVersionResponseBodyData {
	s.NodeContent = &v
	return s
}

func (s *GetFileVersionResponseBodyData) SetNodeId(v int64) *GetFileVersionResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *GetFileVersionResponseBodyData) SetStatus(v string) *GetFileVersionResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetFileVersionResponseBodyData) SetUseType(v string) *GetFileVersionResponseBodyData {
	s.UseType = &v
	return s
}

func (s *GetFileVersionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
