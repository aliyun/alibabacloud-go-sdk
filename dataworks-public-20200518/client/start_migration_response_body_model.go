// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMigrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *StartMigrationResponseBody
	GetData() *bool
	SetErrorCode(v string) *StartMigrationResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StartMigrationResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *StartMigrationResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *StartMigrationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartMigrationResponseBody
	GetSuccess() *bool
}

type StartMigrationResponseBody struct {
	// Indicates whether the migration task is started. Valid values:
	//
	// 	- true: The migration task is started.
	//
	// 	- false: The migration task fails to be started.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 110001123445
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// test error msg
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
	// FSDFSDF-WERWER-XVCX-DSFSDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartMigrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartMigrationResponseBody) GoString() string {
	return s.String()
}

func (s *StartMigrationResponseBody) GetData() *bool {
	return s.Data
}

func (s *StartMigrationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StartMigrationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StartMigrationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartMigrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartMigrationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartMigrationResponseBody) SetData(v bool) *StartMigrationResponseBody {
	s.Data = &v
	return s
}

func (s *StartMigrationResponseBody) SetErrorCode(v string) *StartMigrationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartMigrationResponseBody) SetErrorMessage(v string) *StartMigrationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StartMigrationResponseBody) SetHttpStatusCode(v int32) *StartMigrationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartMigrationResponseBody) SetRequestId(v string) *StartMigrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartMigrationResponseBody) SetSuccess(v bool) *StartMigrationResponseBody {
	s.Success = &v
	return s
}

func (s *StartMigrationResponseBody) Validate() error {
	return dara.Validate(s)
}
