// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseMergeRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CloseMergeRequestResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CloseMergeRequestResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CloseMergeRequestResponseBody
	GetRequestId() *string
	SetResult(v *CloseMergeRequestResponseBodyResult) *CloseMergeRequestResponseBody
	GetResult() *CloseMergeRequestResponseBodyResult
	SetSuccess(v bool) *CloseMergeRequestResponseBody
	GetSuccess() *bool
}

type CloseMergeRequestResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// InvalidParam.NoPermission
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// HC93CE1A-8D7A-13A9-8306-7465DE2E5C0F
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CloseMergeRequestResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CloseMergeRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseMergeRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CloseMergeRequestResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CloseMergeRequestResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CloseMergeRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseMergeRequestResponseBody) GetResult() *CloseMergeRequestResponseBodyResult {
	return s.Result
}

func (s *CloseMergeRequestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CloseMergeRequestResponseBody) SetErrorCode(v string) *CloseMergeRequestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CloseMergeRequestResponseBody) SetErrorMessage(v string) *CloseMergeRequestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CloseMergeRequestResponseBody) SetRequestId(v string) *CloseMergeRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseMergeRequestResponseBody) SetResult(v *CloseMergeRequestResponseBodyResult) *CloseMergeRequestResponseBody {
	s.Result = v
	return s
}

func (s *CloseMergeRequestResponseBody) SetSuccess(v bool) *CloseMergeRequestResponseBody {
	s.Success = &v
	return s
}

func (s *CloseMergeRequestResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloseMergeRequestResponseBodyResult struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CloseMergeRequestResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CloseMergeRequestResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CloseMergeRequestResponseBodyResult) GetResult() *bool {
	return s.Result
}

func (s *CloseMergeRequestResponseBodyResult) SetResult(v bool) *CloseMergeRequestResponseBodyResult {
	s.Result = &v
	return s
}

func (s *CloseMergeRequestResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
