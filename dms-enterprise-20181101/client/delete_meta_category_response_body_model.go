// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetaCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteMetaCategoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteMetaCategoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteMetaCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMetaCategoryResponseBody
	GetSuccess() *bool
}

type DeleteMetaCategoryResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7FAD400F-7A5C-4193-8F9A-39D86C4F0231
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMetaCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetaCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetaCategoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteMetaCategoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteMetaCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMetaCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMetaCategoryResponseBody) SetErrorCode(v string) *DeleteMetaCategoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMetaCategoryResponseBody) SetErrorMessage(v string) *DeleteMetaCategoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteMetaCategoryResponseBody) SetRequestId(v string) *DeleteMetaCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetaCategoryResponseBody) SetSuccess(v bool) *DeleteMetaCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMetaCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
