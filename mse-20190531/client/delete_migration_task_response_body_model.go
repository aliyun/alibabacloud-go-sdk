// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMigrationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteMigrationTaskResponseBody
	GetData() *bool
	SetErrorCode(v string) *DeleteMigrationTaskResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *DeleteMigrationTaskResponseBody
	GetHttpCode() *string
	SetMessage(v string) *DeleteMigrationTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMigrationTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMigrationTaskResponseBody
	GetSuccess() *bool
}

type DeleteMigrationTaskResponseBody struct {
	// The deletion result.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8BD1E58D-0755-42AC-A599-E6B55112****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMigrationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMigrationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMigrationTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteMigrationTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteMigrationTaskResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *DeleteMigrationTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMigrationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMigrationTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMigrationTaskResponseBody) SetData(v bool) *DeleteMigrationTaskResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteMigrationTaskResponseBody) SetErrorCode(v string) *DeleteMigrationTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMigrationTaskResponseBody) SetHttpCode(v string) *DeleteMigrationTaskResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteMigrationTaskResponseBody) SetMessage(v string) *DeleteMigrationTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMigrationTaskResponseBody) SetRequestId(v string) *DeleteMigrationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMigrationTaskResponseBody) SetSuccess(v bool) *DeleteMigrationTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMigrationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
