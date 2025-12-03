// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePushReviewOnOffResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdatePushReviewOnOffResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdatePushReviewOnOffResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdatePushReviewOnOffResponseBody
	GetRequestId() *string
	SetResult(v *UpdatePushReviewOnOffResponseBodyResult) *UpdatePushReviewOnOffResponseBody
	GetResult() *UpdatePushReviewOnOffResponseBodyResult
	SetSuccess(v bool) *UpdatePushReviewOnOffResponseBody
	GetSuccess() *bool
}

type UpdatePushReviewOnOffResponseBody struct {
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
	// 6177543A-8D54-5736-A93B-E0195A1512CB
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdatePushReviewOnOffResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdatePushReviewOnOffResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePushReviewOnOffResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePushReviewOnOffResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdatePushReviewOnOffResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdatePushReviewOnOffResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePushReviewOnOffResponseBody) GetResult() *UpdatePushReviewOnOffResponseBodyResult {
	return s.Result
}

func (s *UpdatePushReviewOnOffResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePushReviewOnOffResponseBody) SetErrorCode(v string) *UpdatePushReviewOnOffResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdatePushReviewOnOffResponseBody) SetErrorMessage(v string) *UpdatePushReviewOnOffResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePushReviewOnOffResponseBody) SetRequestId(v string) *UpdatePushReviewOnOffResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePushReviewOnOffResponseBody) SetResult(v *UpdatePushReviewOnOffResponseBodyResult) *UpdatePushReviewOnOffResponseBody {
	s.Result = v
	return s
}

func (s *UpdatePushReviewOnOffResponseBody) SetSuccess(v bool) *UpdatePushReviewOnOffResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePushReviewOnOffResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePushReviewOnOffResponseBodyResult struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdatePushReviewOnOffResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdatePushReviewOnOffResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdatePushReviewOnOffResponseBodyResult) GetResult() *bool {
	return s.Result
}

func (s *UpdatePushReviewOnOffResponseBodyResult) SetResult(v bool) *UpdatePushReviewOnOffResponseBodyResult {
	s.Result = &v
	return s
}

func (s *UpdatePushReviewOnOffResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
