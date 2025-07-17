// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCorrectBackupFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataCorrectBackupFiles(v *GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles) *GetDataCorrectBackupFilesResponseBody
	GetDataCorrectBackupFiles() *GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles
	SetErrorCode(v string) *GetDataCorrectBackupFilesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataCorrectBackupFilesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataCorrectBackupFilesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataCorrectBackupFilesResponseBody
	GetSuccess() *bool
}

type GetDataCorrectBackupFilesResponseBody struct {
	// The download URL of the backup file for the ticket.
	DataCorrectBackupFiles *GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles `json:"DataCorrectBackupFiles,omitempty" xml:"DataCorrectBackupFiles,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4AFF4109-FEFB-44E8-96A3-923B1FA8C46E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataCorrectBackupFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectBackupFilesResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCorrectBackupFilesResponseBody) GetDataCorrectBackupFiles() *GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles {
	return s.DataCorrectBackupFiles
}

func (s *GetDataCorrectBackupFilesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataCorrectBackupFilesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataCorrectBackupFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCorrectBackupFilesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataCorrectBackupFilesResponseBody) SetDataCorrectBackupFiles(v *GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles) *GetDataCorrectBackupFilesResponseBody {
	s.DataCorrectBackupFiles = v
	return s
}

func (s *GetDataCorrectBackupFilesResponseBody) SetErrorCode(v string) *GetDataCorrectBackupFilesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataCorrectBackupFilesResponseBody) SetErrorMessage(v string) *GetDataCorrectBackupFilesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataCorrectBackupFilesResponseBody) SetRequestId(v string) *GetDataCorrectBackupFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCorrectBackupFilesResponseBody) SetSuccess(v bool) *GetDataCorrectBackupFilesResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCorrectBackupFilesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles struct {
	FileUrl []*string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty" type:"Repeated"`
}

func (s GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles) GoString() string {
	return s.String()
}

func (s *GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles) GetFileUrl() []*string {
	return s.FileUrl
}

func (s *GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles) SetFileUrl(v []*string) *GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles {
	s.FileUrl = v
	return s
}

func (s *GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles) Validate() error {
	return dara.Validate(s)
}
