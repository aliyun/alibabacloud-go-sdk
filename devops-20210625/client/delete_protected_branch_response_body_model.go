// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProtectedBranchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteProtectedBranchResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteProtectedBranchResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteProtectedBranchResponseBody
	GetRequestId() *string
	SetResult(v *DeleteProtectedBranchResponseBodyResult) *DeleteProtectedBranchResponseBody
	GetResult() *DeleteProtectedBranchResponseBodyResult
	SetSuccess(v bool) *DeleteProtectedBranchResponseBody
	GetSuccess() *bool
}

type DeleteProtectedBranchResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 1F4F342D-493A-5B2C-B133-BA78B30FF834
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DeleteProtectedBranchResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteProtectedBranchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProtectedBranchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProtectedBranchResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteProtectedBranchResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteProtectedBranchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProtectedBranchResponseBody) GetResult() *DeleteProtectedBranchResponseBodyResult {
	return s.Result
}

func (s *DeleteProtectedBranchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteProtectedBranchResponseBody) SetErrorCode(v string) *DeleteProtectedBranchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteProtectedBranchResponseBody) SetErrorMessage(v string) *DeleteProtectedBranchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteProtectedBranchResponseBody) SetRequestId(v string) *DeleteProtectedBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProtectedBranchResponseBody) SetResult(v *DeleteProtectedBranchResponseBodyResult) *DeleteProtectedBranchResponseBody {
	s.Result = v
	return s
}

func (s *DeleteProtectedBranchResponseBody) SetSuccess(v bool) *DeleteProtectedBranchResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteProtectedBranchResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteProtectedBranchResponseBodyResult struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteProtectedBranchResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteProtectedBranchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteProtectedBranchResponseBodyResult) GetResult() *bool {
	return s.Result
}

func (s *DeleteProtectedBranchResponseBodyResult) SetResult(v bool) *DeleteProtectedBranchResponseBodyResult {
	s.Result = &v
	return s
}

func (s *DeleteProtectedBranchResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
