// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReviewMergeRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ReviewMergeRequestResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ReviewMergeRequestResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ReviewMergeRequestResponseBody
	GetRequestId() *string
	SetResult(v *ReviewMergeRequestResponseBodyResult) *ReviewMergeRequestResponseBody
	GetResult() *ReviewMergeRequestResponseBodyResult
	SetSuccess(v bool) *ReviewMergeRequestResponseBody
	GetSuccess() *bool
}

type ReviewMergeRequestResponseBody struct {
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
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *ReviewMergeRequestResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ReviewMergeRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReviewMergeRequestResponseBody) GoString() string {
	return s.String()
}

func (s *ReviewMergeRequestResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ReviewMergeRequestResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ReviewMergeRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReviewMergeRequestResponseBody) GetResult() *ReviewMergeRequestResponseBodyResult {
	return s.Result
}

func (s *ReviewMergeRequestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReviewMergeRequestResponseBody) SetErrorCode(v string) *ReviewMergeRequestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReviewMergeRequestResponseBody) SetErrorMessage(v string) *ReviewMergeRequestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ReviewMergeRequestResponseBody) SetRequestId(v string) *ReviewMergeRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReviewMergeRequestResponseBody) SetResult(v *ReviewMergeRequestResponseBodyResult) *ReviewMergeRequestResponseBody {
	s.Result = v
	return s
}

func (s *ReviewMergeRequestResponseBody) SetSuccess(v bool) *ReviewMergeRequestResponseBody {
	s.Success = &v
	return s
}

func (s *ReviewMergeRequestResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReviewMergeRequestResponseBodyResult struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ReviewMergeRequestResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ReviewMergeRequestResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ReviewMergeRequestResponseBodyResult) GetResult() *bool {
	return s.Result
}

func (s *ReviewMergeRequestResponseBodyResult) SetResult(v bool) *ReviewMergeRequestResponseBodyResult {
	s.Result = &v
	return s
}

func (s *ReviewMergeRequestResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
