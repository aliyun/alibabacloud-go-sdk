// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMergeRequestPersonnelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateMergeRequestPersonnelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateMergeRequestPersonnelResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateMergeRequestPersonnelResponseBody
	GetRequestId() *string
	SetResult(v *UpdateMergeRequestPersonnelResponseBodyResult) *UpdateMergeRequestPersonnelResponseBody
	GetResult() *UpdateMergeRequestPersonnelResponseBodyResult
	SetSuccess(v bool) *UpdateMergeRequestPersonnelResponseBody
	GetSuccess() *bool
}

type UpdateMergeRequestPersonnelResponseBody struct {
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
	// 4D6AF7CC-B43B-5454-86AB-023D25E44868
	RequestId *string                                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateMergeRequestPersonnelResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateMergeRequestPersonnelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMergeRequestPersonnelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestPersonnelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateMergeRequestPersonnelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateMergeRequestPersonnelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMergeRequestPersonnelResponseBody) GetResult() *UpdateMergeRequestPersonnelResponseBodyResult {
	return s.Result
}

func (s *UpdateMergeRequestPersonnelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMergeRequestPersonnelResponseBody) SetErrorCode(v string) *UpdateMergeRequestPersonnelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateMergeRequestPersonnelResponseBody) SetErrorMessage(v string) *UpdateMergeRequestPersonnelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateMergeRequestPersonnelResponseBody) SetRequestId(v string) *UpdateMergeRequestPersonnelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMergeRequestPersonnelResponseBody) SetResult(v *UpdateMergeRequestPersonnelResponseBodyResult) *UpdateMergeRequestPersonnelResponseBody {
	s.Result = v
	return s
}

func (s *UpdateMergeRequestPersonnelResponseBody) SetSuccess(v bool) *UpdateMergeRequestPersonnelResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMergeRequestPersonnelResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMergeRequestPersonnelResponseBodyResult struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateMergeRequestPersonnelResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateMergeRequestPersonnelResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestPersonnelResponseBodyResult) GetResult() *bool {
	return s.Result
}

func (s *UpdateMergeRequestPersonnelResponseBodyResult) SetResult(v bool) *UpdateMergeRequestPersonnelResponseBodyResult {
	s.Result = &v
	return s
}

func (s *UpdateMergeRequestPersonnelResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
