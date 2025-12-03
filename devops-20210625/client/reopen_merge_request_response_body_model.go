// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReopenMergeRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ReopenMergeRequestResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ReopenMergeRequestResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ReopenMergeRequestResponseBody
	GetRequestId() *string
	SetResult(v *ReopenMergeRequestResponseBodyResult) *ReopenMergeRequestResponseBody
	GetResult() *ReopenMergeRequestResponseBodyResult
	SetSuccess(v bool) *ReopenMergeRequestResponseBody
	GetSuccess() *bool
}

type ReopenMergeRequestResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 313A1BF6-63B7-52D4-A098-952221A65254
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *ReopenMergeRequestResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ReopenMergeRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReopenMergeRequestResponseBody) GoString() string {
	return s.String()
}

func (s *ReopenMergeRequestResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ReopenMergeRequestResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ReopenMergeRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReopenMergeRequestResponseBody) GetResult() *ReopenMergeRequestResponseBodyResult {
	return s.Result
}

func (s *ReopenMergeRequestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReopenMergeRequestResponseBody) SetErrorCode(v string) *ReopenMergeRequestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReopenMergeRequestResponseBody) SetErrorMessage(v string) *ReopenMergeRequestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ReopenMergeRequestResponseBody) SetRequestId(v string) *ReopenMergeRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReopenMergeRequestResponseBody) SetResult(v *ReopenMergeRequestResponseBodyResult) *ReopenMergeRequestResponseBody {
	s.Result = v
	return s
}

func (s *ReopenMergeRequestResponseBody) SetSuccess(v bool) *ReopenMergeRequestResponseBody {
	s.Success = &v
	return s
}

func (s *ReopenMergeRequestResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReopenMergeRequestResponseBodyResult struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ReopenMergeRequestResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ReopenMergeRequestResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ReopenMergeRequestResponseBodyResult) GetResult() *bool {
	return s.Result
}

func (s *ReopenMergeRequestResponseBodyResult) SetResult(v bool) *ReopenMergeRequestResponseBodyResult {
	s.Result = &v
	return s
}

func (s *ReopenMergeRequestResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
