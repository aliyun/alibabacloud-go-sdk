// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMergeRequestFilesReadsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListMergeRequestFilesReadsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListMergeRequestFilesReadsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListMergeRequestFilesReadsResponseBody
	GetRequestId() *string
	SetResult(v []*ListMergeRequestFilesReadsResponseBodyResult) *ListMergeRequestFilesReadsResponseBody
	GetResult() []*ListMergeRequestFilesReadsResponseBodyResult
	SetSuccess(v bool) *ListMergeRequestFilesReadsResponseBody
	GetSuccess() *bool
}

type ListMergeRequestFilesReadsResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListMergeRequestFilesReadsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListMergeRequestFilesReadsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestFilesReadsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMergeRequestFilesReadsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMergeRequestFilesReadsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListMergeRequestFilesReadsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMergeRequestFilesReadsResponseBody) GetResult() []*ListMergeRequestFilesReadsResponseBodyResult {
	return s.Result
}

func (s *ListMergeRequestFilesReadsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMergeRequestFilesReadsResponseBody) SetErrorCode(v string) *ListMergeRequestFilesReadsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBody) SetErrorMessage(v string) *ListMergeRequestFilesReadsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBody) SetRequestId(v string) *ListMergeRequestFilesReadsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBody) SetResult(v []*ListMergeRequestFilesReadsResponseBodyResult) *ListMergeRequestFilesReadsResponseBody {
	s.Result = v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBody) SetSuccess(v bool) *ListMergeRequestFilesReadsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMergeRequestFilesReadsResponseBodyResult struct {
	// example:
	//
	// false
	DeletedFile *string `json:"deletedFile,omitempty" xml:"deletedFile,omitempty"`
	// example:
	//
	// true
	NewFile *bool `json:"newFile,omitempty" xml:"newFile,omitempty"`
	// example:
	//
	// text.txt
	NewFilePath *string `json:"newFilePath,omitempty" xml:"newFilePath,omitempty"`
	// example:
	//
	// text.txt
	OldFilePath *string                                                  `json:"oldFilePath,omitempty" xml:"oldFilePath,omitempty"`
	ReadUsers   []*ListMergeRequestFilesReadsResponseBodyResultReadUsers `json:"readUsers,omitempty" xml:"readUsers,omitempty" type:"Repeated"`
	// example:
	//
	// false
	RenamedFile *string `json:"renamedFile,omitempty" xml:"renamedFile,omitempty"`
}

func (s ListMergeRequestFilesReadsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestFilesReadsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMergeRequestFilesReadsResponseBodyResult) GetDeletedFile() *string {
	return s.DeletedFile
}

func (s *ListMergeRequestFilesReadsResponseBodyResult) GetNewFile() *bool {
	return s.NewFile
}

func (s *ListMergeRequestFilesReadsResponseBodyResult) GetNewFilePath() *string {
	return s.NewFilePath
}

func (s *ListMergeRequestFilesReadsResponseBodyResult) GetOldFilePath() *string {
	return s.OldFilePath
}

func (s *ListMergeRequestFilesReadsResponseBodyResult) GetReadUsers() []*ListMergeRequestFilesReadsResponseBodyResultReadUsers {
	return s.ReadUsers
}

func (s *ListMergeRequestFilesReadsResponseBodyResult) GetRenamedFile() *string {
	return s.RenamedFile
}

func (s *ListMergeRequestFilesReadsResponseBodyResult) SetDeletedFile(v string) *ListMergeRequestFilesReadsResponseBodyResult {
	s.DeletedFile = &v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBodyResult) SetNewFile(v bool) *ListMergeRequestFilesReadsResponseBodyResult {
	s.NewFile = &v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBodyResult) SetNewFilePath(v string) *ListMergeRequestFilesReadsResponseBodyResult {
	s.NewFilePath = &v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBodyResult) SetOldFilePath(v string) *ListMergeRequestFilesReadsResponseBodyResult {
	s.OldFilePath = &v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBodyResult) SetReadUsers(v []*ListMergeRequestFilesReadsResponseBodyResultReadUsers) *ListMergeRequestFilesReadsResponseBodyResult {
	s.ReadUsers = v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBodyResult) SetRenamedFile(v string) *ListMergeRequestFilesReadsResponseBodyResult {
	s.RenamedFile = &v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBodyResult) Validate() error {
	if s.ReadUsers != nil {
		for _, item := range s.ReadUsers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMergeRequestFilesReadsResponseBodyResultReadUsers struct {
	// example:
	//
	// 204485087002425236
	AliyunPk *string `json:"aliyunPk,omitempty" xml:"aliyunPk,omitempty"`
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// test-codeup
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListMergeRequestFilesReadsResponseBodyResultReadUsers) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestFilesReadsResponseBodyResultReadUsers) GoString() string {
	return s.String()
}

func (s *ListMergeRequestFilesReadsResponseBodyResultReadUsers) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *ListMergeRequestFilesReadsResponseBodyResultReadUsers) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListMergeRequestFilesReadsResponseBodyResultReadUsers) GetEmail() *string {
	return s.Email
}

func (s *ListMergeRequestFilesReadsResponseBodyResultReadUsers) GetName() *string {
	return s.Name
}

func (s *ListMergeRequestFilesReadsResponseBodyResultReadUsers) GetState() *string {
	return s.State
}

func (s *ListMergeRequestFilesReadsResponseBodyResultReadUsers) GetUsername() *string {
	return s.Username
}

func (s *ListMergeRequestFilesReadsResponseBodyResultReadUsers) SetAliyunPk(v string) *ListMergeRequestFilesReadsResponseBodyResultReadUsers {
	s.AliyunPk = &v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBodyResultReadUsers) SetAvatarUrl(v string) *ListMergeRequestFilesReadsResponseBodyResultReadUsers {
	s.AvatarUrl = &v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBodyResultReadUsers) SetEmail(v string) *ListMergeRequestFilesReadsResponseBodyResultReadUsers {
	s.Email = &v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBodyResultReadUsers) SetName(v string) *ListMergeRequestFilesReadsResponseBodyResultReadUsers {
	s.Name = &v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBodyResultReadUsers) SetState(v string) *ListMergeRequestFilesReadsResponseBodyResultReadUsers {
	s.State = &v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBodyResultReadUsers) SetUsername(v string) *ListMergeRequestFilesReadsResponseBodyResultReadUsers {
	s.Username = &v
	return s
}

func (s *ListMergeRequestFilesReadsResponseBodyResultReadUsers) Validate() error {
	return dara.Validate(s)
}
