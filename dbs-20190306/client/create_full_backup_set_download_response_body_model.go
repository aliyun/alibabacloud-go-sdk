// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFullBackupSetDownloadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetDownloadTaskId(v string) *CreateFullBackupSetDownloadResponseBody
	GetBackupSetDownloadTaskId() *string
	SetErrCode(v string) *CreateFullBackupSetDownloadResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateFullBackupSetDownloadResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CreateFullBackupSetDownloadResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateFullBackupSetDownloadResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateFullBackupSetDownloadResponseBody
	GetSuccess() *bool
}

type CreateFullBackupSetDownloadResponseBody struct {
	// The ID of the backup set download task.
	//
	// example:
	//
	// dbstooi01e****
	BackupSetDownloadTaskId *string `json:"BackupSetDownloadTaskId,omitempty" xml:"BackupSetDownloadTaskId,omitempty"`
	// The error code.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D6E068C3-25BC-455A-85FE-45F0B22ECB1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFullBackupSetDownloadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFullBackupSetDownloadResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFullBackupSetDownloadResponseBody) GetBackupSetDownloadTaskId() *string {
	return s.BackupSetDownloadTaskId
}

func (s *CreateFullBackupSetDownloadResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateFullBackupSetDownloadResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateFullBackupSetDownloadResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateFullBackupSetDownloadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFullBackupSetDownloadResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateFullBackupSetDownloadResponseBody) SetBackupSetDownloadTaskId(v string) *CreateFullBackupSetDownloadResponseBody {
	s.BackupSetDownloadTaskId = &v
	return s
}

func (s *CreateFullBackupSetDownloadResponseBody) SetErrCode(v string) *CreateFullBackupSetDownloadResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateFullBackupSetDownloadResponseBody) SetErrMessage(v string) *CreateFullBackupSetDownloadResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateFullBackupSetDownloadResponseBody) SetHttpStatusCode(v int32) *CreateFullBackupSetDownloadResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateFullBackupSetDownloadResponseBody) SetRequestId(v string) *CreateFullBackupSetDownloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFullBackupSetDownloadResponseBody) SetSuccess(v bool) *CreateFullBackupSetDownloadResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFullBackupSetDownloadResponseBody) Validate() error {
	return dara.Validate(s)
}
