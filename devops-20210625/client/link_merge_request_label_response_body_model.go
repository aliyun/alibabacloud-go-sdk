// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLinkMergeRequestLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *LinkMergeRequestLabelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *LinkMergeRequestLabelResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *LinkMergeRequestLabelResponseBody
	GetRequestId() *string
	SetResult(v *LinkMergeRequestLabelResponseBodyResult) *LinkMergeRequestLabelResponseBody
	GetResult() *LinkMergeRequestLabelResponseBodyResult
	SetSuccess(v bool) *LinkMergeRequestLabelResponseBody
	GetSuccess() *bool
}

type LinkMergeRequestLabelResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *LinkMergeRequestLabelResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s LinkMergeRequestLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LinkMergeRequestLabelResponseBody) GoString() string {
	return s.String()
}

func (s *LinkMergeRequestLabelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *LinkMergeRequestLabelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *LinkMergeRequestLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LinkMergeRequestLabelResponseBody) GetResult() *LinkMergeRequestLabelResponseBodyResult {
	return s.Result
}

func (s *LinkMergeRequestLabelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *LinkMergeRequestLabelResponseBody) SetErrorCode(v string) *LinkMergeRequestLabelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *LinkMergeRequestLabelResponseBody) SetErrorMessage(v string) *LinkMergeRequestLabelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *LinkMergeRequestLabelResponseBody) SetRequestId(v string) *LinkMergeRequestLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *LinkMergeRequestLabelResponseBody) SetResult(v *LinkMergeRequestLabelResponseBodyResult) *LinkMergeRequestLabelResponseBody {
	s.Result = v
	return s
}

func (s *LinkMergeRequestLabelResponseBody) SetSuccess(v bool) *LinkMergeRequestLabelResponseBody {
	s.Success = &v
	return s
}

func (s *LinkMergeRequestLabelResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LinkMergeRequestLabelResponseBodyResult struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s LinkMergeRequestLabelResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s LinkMergeRequestLabelResponseBodyResult) GoString() string {
	return s.String()
}

func (s *LinkMergeRequestLabelResponseBodyResult) GetResult() *bool {
	return s.Result
}

func (s *LinkMergeRequestLabelResponseBodyResult) SetResult(v bool) *LinkMergeRequestLabelResponseBodyResult {
	s.Result = &v
	return s
}

func (s *LinkMergeRequestLabelResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
