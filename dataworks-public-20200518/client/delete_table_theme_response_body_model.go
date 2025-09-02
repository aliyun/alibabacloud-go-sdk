// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTableThemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteResult(v bool) *DeleteTableThemeResponseBody
	GetDeleteResult() *bool
	SetErrorCode(v string) *DeleteTableThemeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteTableThemeResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteTableThemeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteTableThemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTableThemeResponseBody
	GetSuccess() *bool
}

type DeleteTableThemeResponseBody struct {
	// Indicates whether the theme was deleted.
	//
	// example:
	//
	// true
	DeleteResult *bool `json:"DeleteResult,omitempty" xml:"DeleteResult,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// abcdef
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTableThemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTableThemeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTableThemeResponseBody) GetDeleteResult() *bool {
	return s.DeleteResult
}

func (s *DeleteTableThemeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteTableThemeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteTableThemeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteTableThemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTableThemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTableThemeResponseBody) SetDeleteResult(v bool) *DeleteTableThemeResponseBody {
	s.DeleteResult = &v
	return s
}

func (s *DeleteTableThemeResponseBody) SetErrorCode(v string) *DeleteTableThemeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTableThemeResponseBody) SetErrorMessage(v string) *DeleteTableThemeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteTableThemeResponseBody) SetHttpStatusCode(v int32) *DeleteTableThemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteTableThemeResponseBody) SetRequestId(v string) *DeleteTableThemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTableThemeResponseBody) SetSuccess(v bool) *DeleteTableThemeResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTableThemeResponseBody) Validate() error {
	return dara.Validate(s)
}
