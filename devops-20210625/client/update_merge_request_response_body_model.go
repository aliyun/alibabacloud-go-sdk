// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMergeRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateMergeRequestResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateMergeRequestResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateMergeRequestResponseBody
	GetRequestId() *string
	SetResult(v *UpdateMergeRequestResponseBodyResult) *UpdateMergeRequestResponseBody
	GetResult() *UpdateMergeRequestResponseBodyResult
	SetSuccess(v bool) *UpdateMergeRequestResponseBody
	GetSuccess() *bool
}

type UpdateMergeRequestResponseBody struct {
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
	// F7B85D1B-D1C2-140F-A039-341859F130B9
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateMergeRequestResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateMergeRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMergeRequestResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateMergeRequestResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateMergeRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMergeRequestResponseBody) GetResult() *UpdateMergeRequestResponseBodyResult {
	return s.Result
}

func (s *UpdateMergeRequestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMergeRequestResponseBody) SetErrorCode(v string) *UpdateMergeRequestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateMergeRequestResponseBody) SetErrorMessage(v string) *UpdateMergeRequestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateMergeRequestResponseBody) SetRequestId(v string) *UpdateMergeRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMergeRequestResponseBody) SetResult(v *UpdateMergeRequestResponseBodyResult) *UpdateMergeRequestResponseBody {
	s.Result = v
	return s
}

func (s *UpdateMergeRequestResponseBody) SetSuccess(v bool) *UpdateMergeRequestResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMergeRequestResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMergeRequestResponseBodyResult struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateMergeRequestResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateMergeRequestResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestResponseBodyResult) GetResult() *bool {
	return s.Result
}

func (s *UpdateMergeRequestResponseBodyResult) SetResult(v bool) *UpdateMergeRequestResponseBodyResult {
	s.Result = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
