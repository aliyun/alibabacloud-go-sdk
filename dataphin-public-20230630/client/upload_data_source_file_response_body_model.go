// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDataSourceFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadDataSourceFileResponseBody
	GetCode() *string
	SetFileId(v string) *UploadDataSourceFileResponseBody
	GetFileId() *string
	SetHttpStatusCode(v int32) *UploadDataSourceFileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UploadDataSourceFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadDataSourceFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadDataSourceFileResponseBody
	GetSuccess() *bool
}

type UploadDataSourceFileResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The file identifier (fileId). Reference this value in the datasource config when creating a datasource, such as kafka.kerberos.keytab.file=fileId.
	//
	// example:
	//
	// a1b2c3d4e5f6
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error details returned by the backend.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadDataSourceFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSourceFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDataSourceFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadDataSourceFileResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *UploadDataSourceFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UploadDataSourceFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadDataSourceFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadDataSourceFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadDataSourceFileResponseBody) SetCode(v string) *UploadDataSourceFileResponseBody {
	s.Code = &v
	return s
}

func (s *UploadDataSourceFileResponseBody) SetFileId(v string) *UploadDataSourceFileResponseBody {
	s.FileId = &v
	return s
}

func (s *UploadDataSourceFileResponseBody) SetHttpStatusCode(v int32) *UploadDataSourceFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UploadDataSourceFileResponseBody) SetMessage(v string) *UploadDataSourceFileResponseBody {
	s.Message = &v
	return s
}

func (s *UploadDataSourceFileResponseBody) SetRequestId(v string) *UploadDataSourceFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDataSourceFileResponseBody) SetSuccess(v bool) *UploadDataSourceFileResponseBody {
	s.Success = &v
	return s
}

func (s *UploadDataSourceFileResponseBody) Validate() error {
	return dara.Validate(s)
}
