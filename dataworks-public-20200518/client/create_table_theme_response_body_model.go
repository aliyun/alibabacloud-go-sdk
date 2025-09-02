// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableThemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateTableThemeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateTableThemeResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateTableThemeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateTableThemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTableThemeResponseBody
	GetSuccess() *bool
	SetThemeId(v int64) *CreateTableThemeResponseBody
	GetThemeId() *int64
}

type CreateTableThemeResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
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
	// ABCd
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The table theme ID.
	//
	// example:
	//
	// 123
	ThemeId *int64 `json:"ThemeId,omitempty" xml:"ThemeId,omitempty"`
}

func (s CreateTableThemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTableThemeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTableThemeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateTableThemeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateTableThemeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateTableThemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTableThemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTableThemeResponseBody) GetThemeId() *int64 {
	return s.ThemeId
}

func (s *CreateTableThemeResponseBody) SetErrorCode(v string) *CreateTableThemeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateTableThemeResponseBody) SetErrorMessage(v string) *CreateTableThemeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateTableThemeResponseBody) SetHttpStatusCode(v int32) *CreateTableThemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTableThemeResponseBody) SetRequestId(v string) *CreateTableThemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTableThemeResponseBody) SetSuccess(v bool) *CreateTableThemeResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTableThemeResponseBody) SetThemeId(v int64) *CreateTableThemeResponseBody {
	s.ThemeId = &v
	return s
}

func (s *CreateTableThemeResponseBody) Validate() error {
	return dara.Validate(s)
}
