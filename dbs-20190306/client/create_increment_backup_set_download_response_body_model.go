// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIncrementBackupSetDownloadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetDownloadTaskId(v string) *CreateIncrementBackupSetDownloadResponseBody
	GetBackupSetDownloadTaskId() *string
	SetErrCode(v string) *CreateIncrementBackupSetDownloadResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateIncrementBackupSetDownloadResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CreateIncrementBackupSetDownloadResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateIncrementBackupSetDownloadResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateIncrementBackupSetDownloadResponseBody
	GetSuccess() *bool
}

type CreateIncrementBackupSetDownloadResponseBody struct {
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

func (s CreateIncrementBackupSetDownloadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIncrementBackupSetDownloadResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIncrementBackupSetDownloadResponseBody) GetBackupSetDownloadTaskId() *string {
	return s.BackupSetDownloadTaskId
}

func (s *CreateIncrementBackupSetDownloadResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateIncrementBackupSetDownloadResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateIncrementBackupSetDownloadResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateIncrementBackupSetDownloadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIncrementBackupSetDownloadResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateIncrementBackupSetDownloadResponseBody) SetBackupSetDownloadTaskId(v string) *CreateIncrementBackupSetDownloadResponseBody {
	s.BackupSetDownloadTaskId = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadResponseBody) SetErrCode(v string) *CreateIncrementBackupSetDownloadResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadResponseBody) SetErrMessage(v string) *CreateIncrementBackupSetDownloadResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadResponseBody) SetHttpStatusCode(v int32) *CreateIncrementBackupSetDownloadResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadResponseBody) SetRequestId(v string) *CreateIncrementBackupSetDownloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadResponseBody) SetSuccess(v bool) *CreateIncrementBackupSetDownloadResponseBody {
	s.Success = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadResponseBody) Validate() error {
	return dara.Validate(s)
}
