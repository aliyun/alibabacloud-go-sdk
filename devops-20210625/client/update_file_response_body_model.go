// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateFileResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateFileResponseBody
	GetRequestId() *string
	SetResult(v *UpdateFileResponseBodyResult) *UpdateFileResponseBody
	GetResult() *UpdateFileResponseBodyResult
	SetSuccess(v bool) *UpdateFileResponseBody
	GetSuccess() *bool
}

type UpdateFileResponseBody struct {
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
	// C2F153F6-BB43-50C4-9F4F-40593203E19A
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFileResponseBody) GetResult() *UpdateFileResponseBodyResult {
	return s.Result
}

func (s *UpdateFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateFileResponseBody) SetErrorCode(v string) *UpdateFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateFileResponseBody) SetErrorMessage(v string) *UpdateFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateFileResponseBody) SetRequestId(v string) *UpdateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileResponseBody) SetResult(v *UpdateFileResponseBodyResult) *UpdateFileResponseBody {
	s.Result = v
	return s
}

func (s *UpdateFileResponseBody) SetSuccess(v bool) *UpdateFileResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFileResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateFileResponseBodyResult struct {
	// example:
	//
	// master
	BranchName *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	// example:
	//
	// src/main/update.txt
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
}

func (s UpdateFileResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateFileResponseBodyResult) GetBranchName() *string {
	return s.BranchName
}

func (s *UpdateFileResponseBodyResult) GetFilePath() *string {
	return s.FilePath
}

func (s *UpdateFileResponseBodyResult) SetBranchName(v string) *UpdateFileResponseBodyResult {
	s.BranchName = &v
	return s
}

func (s *UpdateFileResponseBodyResult) SetFilePath(v string) *UpdateFileResponseBodyResult {
	s.FilePath = &v
	return s
}

func (s *UpdateFileResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
