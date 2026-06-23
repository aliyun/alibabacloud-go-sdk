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
	// Version details of the file.
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
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The unique ID of this request. If an error occurs, you can troubleshoot the issue using this ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// - true: Succeeded.
	//
	// - false: Failed.
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
	// The change type of this file version, including CREATE, UPDATE, and DELETE.
	//
	// example:
	//
	// UPDATE
	ChangeType *string `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// Description of this file version.
	//
	// example:
	//
	// Second version submission
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// UNIX timestamp (in milliseconds) when the file version was generated.
	//
	// example:
	//
	// 1593881265000
	CommitTime *int64 `json:"CommitTime,omitempty" xml:"CommitTime,omitempty"`
	// User ID of the Alibaba Cloud user who generated this file version.
	//
	// example:
	//
	// 7384234****
	CommitUser *string `json:"CommitUser,omitempty" xml:"CommitUser,omitempty"`
	// The code of the file for this version.
	//
	// example:
	//
	// SHOW TABLES;
	FileContent *string `json:"FileContent,omitempty" xml:"FileContent,omitempty"`
	// File name used to generate this file version.
	//
	// example:
	//
	// ods_user_info_d
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Basic information of the file used to generate this file version.
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
	// Indicates whether this file version is the latest version in the current production environment.
	//
	// - true: It is the latest version.
	//
	// - false: It is not the latest version.
	//
	// example:
	//
	// true
	IsCurrentProd *bool `json:"IsCurrentProd,omitempty" xml:"IsCurrentProd,omitempty"`
	// The scan configuration at the time this file version was generated.
	//
	// example:
	//
	// {"cycleType":0,"cronExpress":"00 05 00 	- 	- ?"}
	NodeContent *string `json:"NodeContent,omitempty" xml:"NodeContent,omitempty"`
	// The ID of the scheduling task corresponding to the generation of this file version.
	//
	// example:
	//
	// 3000001
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Current status of the file version. Valid values:
	//
	// - COMMITTING (Submitting)
	//
	// - COMMITTED or CHECK_OK (Submitted)
	//
	// - PACKAGED (Preparing for publish)
	//
	// - DEPLOYING (In Publish)
	//
	// - DEPLOYED (Published)
	//
	// - CANCELLED (Publish canceled)
	//
	// example:
	//
	// COMMITTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Function module to which the file belongs. Valid values:
	//
	// - 0: NORMAL (Data Development)
	//
	// - 1: MANUAL (one-time task)
	//
	// - 2: MANUAL_BIZ (manual pipeline)
	//
	// - 3: SKIP (Dry-Run scheduling in Data Development)
	//
	// - 10: ADHOCQUERY (Ad Hoc Query)
	//
	// - 30: COMPONENT (widget Management)
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
