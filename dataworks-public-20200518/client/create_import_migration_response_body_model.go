// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImportMigrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateImportMigrationResponseBody
	GetData() *int64
	SetErrorCode(v string) *CreateImportMigrationResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateImportMigrationResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateImportMigrationResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateImportMigrationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateImportMigrationResponseBody
	GetSuccess() *bool
}

type CreateImportMigrationResponseBody struct {
	// The import task ID. The ID is used as an input parameter if you want the system to run the import task or you want to obtain the running progress of the import task.
	//
	// example:
	//
	// 123456
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 110001123456
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// test error message
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// ADFASDFASDFA-ADFASDF-ASDFADSDF-AFFADS
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateImportMigrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImportMigrationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImportMigrationResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateImportMigrationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateImportMigrationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateImportMigrationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateImportMigrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImportMigrationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateImportMigrationResponseBody) SetData(v int64) *CreateImportMigrationResponseBody {
	s.Data = &v
	return s
}

func (s *CreateImportMigrationResponseBody) SetErrorCode(v string) *CreateImportMigrationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateImportMigrationResponseBody) SetErrorMessage(v string) *CreateImportMigrationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateImportMigrationResponseBody) SetHttpStatusCode(v int32) *CreateImportMigrationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateImportMigrationResponseBody) SetRequestId(v string) *CreateImportMigrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImportMigrationResponseBody) SetSuccess(v bool) *CreateImportMigrationResponseBody {
	s.Success = &v
	return s
}

func (s *CreateImportMigrationResponseBody) Validate() error {
	return dara.Validate(s)
}
