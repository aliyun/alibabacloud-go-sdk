// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipDataCorrectRowCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SkipDataCorrectRowCheckResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SkipDataCorrectRowCheckResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SkipDataCorrectRowCheckResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SkipDataCorrectRowCheckResponseBody
	GetSuccess() *bool
}

type SkipDataCorrectRowCheckResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
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
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SkipDataCorrectRowCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SkipDataCorrectRowCheckResponseBody) GoString() string {
	return s.String()
}

func (s *SkipDataCorrectRowCheckResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SkipDataCorrectRowCheckResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SkipDataCorrectRowCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SkipDataCorrectRowCheckResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SkipDataCorrectRowCheckResponseBody) SetErrorCode(v string) *SkipDataCorrectRowCheckResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SkipDataCorrectRowCheckResponseBody) SetErrorMessage(v string) *SkipDataCorrectRowCheckResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SkipDataCorrectRowCheckResponseBody) SetRequestId(v string) *SkipDataCorrectRowCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *SkipDataCorrectRowCheckResponseBody) SetSuccess(v bool) *SkipDataCorrectRowCheckResponseBody {
	s.Success = &v
	return s
}

func (s *SkipDataCorrectRowCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
