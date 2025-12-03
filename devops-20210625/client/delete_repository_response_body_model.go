// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteRepositoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteRepositoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteRepositoryResponseBody
	GetRequestId() *string
	SetResult(v *DeleteRepositoryResponseBodyResult) *DeleteRepositoryResponseBody
	GetResult() *DeleteRepositoryResponseBodyResult
	SetSuccess(v bool) *DeleteRepositoryResponseBody
	GetSuccess() *bool
}

type DeleteRepositoryResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// A7586FEB-E48D-5579-983F-74981FBFF627
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DeleteRepositoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteRepositoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRepositoryResponseBody) GetResult() *DeleteRepositoryResponseBodyResult {
	return s.Result
}

func (s *DeleteRepositoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRepositoryResponseBody) SetErrorCode(v string) *DeleteRepositoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetErrorMessage(v string) *DeleteRepositoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetRequestId(v string) *DeleteRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetResult(v *DeleteRepositoryResponseBodyResult) *DeleteRepositoryResponseBody {
	s.Result = v
	return s
}

func (s *DeleteRepositoryResponseBody) SetSuccess(v bool) *DeleteRepositoryResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRepositoryResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteRepositoryResponseBodyResult struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteRepositoryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryResponseBodyResult) GetResult() *bool {
	return s.Result
}

func (s *DeleteRepositoryResponseBodyResult) SetResult(v bool) *DeleteRepositoryResponseBodyResult {
	s.Result = &v
	return s
}

func (s *DeleteRepositoryResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
