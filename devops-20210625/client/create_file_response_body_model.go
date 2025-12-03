// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateFileResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateFileResponseBody
	GetRequestId() *string
	SetResult(v *CreateFileResponseBodyResult) *CreateFileResponseBody
	GetResult() *CreateFileResponseBodyResult
	SetSuccess(v bool) *CreateFileResponseBody
	GetSuccess() *bool
}

type CreateFileResponseBody struct {
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
	// F8053E32-9623-511F-8B46-F0E5FD206524
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFileResponseBody) GetResult() *CreateFileResponseBodyResult {
	return s.Result
}

func (s *CreateFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateFileResponseBody) SetErrorCode(v string) *CreateFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFileResponseBody) SetErrorMessage(v string) *CreateFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateFileResponseBody) SetRequestId(v string) *CreateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileResponseBody) SetResult(v *CreateFileResponseBodyResult) *CreateFileResponseBody {
	s.Result = v
	return s
}

func (s *CreateFileResponseBody) SetSuccess(v bool) *CreateFileResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFileResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateFileResponseBodyResult struct {
	// example:
	//
	// master
	BranchName *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	// example:
	//
	// /src/main/test.java
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
}

func (s CreateFileResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateFileResponseBodyResult) GetBranchName() *string {
	return s.BranchName
}

func (s *CreateFileResponseBodyResult) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateFileResponseBodyResult) SetBranchName(v string) *CreateFileResponseBodyResult {
	s.BranchName = &v
	return s
}

func (s *CreateFileResponseBodyResult) SetFilePath(v string) *CreateFileResponseBodyResult {
	s.FilePath = &v
	return s
}

func (s *CreateFileResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
