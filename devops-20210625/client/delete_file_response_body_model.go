// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteFileResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteFileResponseBody
	GetRequestId() *string
	SetResult(v *DeleteFileResponseBodyResult) *DeleteFileResponseBody
	GetResult() *DeleteFileResponseBodyResult
	SetSuccess(v bool) *DeleteFileResponseBody
	GetSuccess() *bool
}

type DeleteFileResponseBody struct {
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
	// 7EFAD5FB-2296-5D52-BC60-FCC992A40767
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DeleteFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFileResponseBody) GetResult() *DeleteFileResponseBodyResult {
	return s.Result
}

func (s *DeleteFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFileResponseBody) SetErrorCode(v string) *DeleteFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFileResponseBody) SetErrorMessage(v string) *DeleteFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFileResponseBody) SetRequestId(v string) *DeleteFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFileResponseBody) SetResult(v *DeleteFileResponseBodyResult) *DeleteFileResponseBody {
	s.Result = v
	return s
}

func (s *DeleteFileResponseBody) SetSuccess(v bool) *DeleteFileResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFileResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteFileResponseBodyResult struct {
	// example:
	//
	// master
	BranchName *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	// example:
	//
	// src/main/delete.java
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
}

func (s DeleteFileResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBodyResult) GetBranchName() *string {
	return s.BranchName
}

func (s *DeleteFileResponseBodyResult) GetFilePath() *string {
	return s.FilePath
}

func (s *DeleteFileResponseBodyResult) SetBranchName(v string) *DeleteFileResponseBodyResult {
	s.BranchName = &v
	return s
}

func (s *DeleteFileResponseBodyResult) SetFilePath(v string) *DeleteFileResponseBodyResult {
	s.FilePath = &v
	return s
}

func (s *DeleteFileResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
