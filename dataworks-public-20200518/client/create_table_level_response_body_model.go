// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateTableLevelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateTableLevelResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateTableLevelResponseBody
	GetHttpStatusCode() *int32
	SetLevelId(v int64) *CreateTableLevelResponseBody
	GetLevelId() *int64
	SetRequestId(v string) *CreateTableLevelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTableLevelResponseBody
	GetSuccess() *bool
}

type CreateTableLevelResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 1031203110005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The table level ID.
	//
	// example:
	//
	// 123
	LevelId *int64 `json:"LevelId,omitempty" xml:"LevelId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// The error message returned.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTableLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTableLevelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTableLevelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateTableLevelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateTableLevelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateTableLevelResponseBody) GetLevelId() *int64 {
	return s.LevelId
}

func (s *CreateTableLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTableLevelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTableLevelResponseBody) SetErrorCode(v string) *CreateTableLevelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateTableLevelResponseBody) SetErrorMessage(v string) *CreateTableLevelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateTableLevelResponseBody) SetHttpStatusCode(v int32) *CreateTableLevelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTableLevelResponseBody) SetLevelId(v int64) *CreateTableLevelResponseBody {
	s.LevelId = &v
	return s
}

func (s *CreateTableLevelResponseBody) SetRequestId(v string) *CreateTableLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTableLevelResponseBody) SetSuccess(v bool) *CreateTableLevelResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTableLevelResponseBody) Validate() error {
	return dara.Validate(s)
}
