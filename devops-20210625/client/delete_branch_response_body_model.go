// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBranchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteBranchResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteBranchResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteBranchResponseBody
	GetRequestId() *string
	SetResult(v *DeleteBranchResponseBodyResult) *DeleteBranchResponseBody
	GetResult() *DeleteBranchResponseBodyResult
	SetSuccess(v string) *DeleteBranchResponseBody
	GetSuccess() *string
}

type DeleteBranchResponseBody struct {
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
	// 6177543A-8D54-5736-A93B-E0195A1512CB
	RequestId *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DeleteBranchResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteBranchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBranchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBranchResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteBranchResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteBranchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBranchResponseBody) GetResult() *DeleteBranchResponseBodyResult {
	return s.Result
}

func (s *DeleteBranchResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteBranchResponseBody) SetErrorCode(v string) *DeleteBranchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteBranchResponseBody) SetErrorMessage(v string) *DeleteBranchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteBranchResponseBody) SetRequestId(v string) *DeleteBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBranchResponseBody) SetResult(v *DeleteBranchResponseBodyResult) *DeleteBranchResponseBody {
	s.Result = v
	return s
}

func (s *DeleteBranchResponseBody) SetSuccess(v string) *DeleteBranchResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBranchResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteBranchResponseBodyResult struct {
	// example:
	//
	// deleteBranch
	BranchName *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
}

func (s DeleteBranchResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteBranchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteBranchResponseBodyResult) GetBranchName() *string {
	return s.BranchName
}

func (s *DeleteBranchResponseBodyResult) SetBranchName(v string) *DeleteBranchResponseBodyResult {
	s.BranchName = &v
	return s
}

func (s *DeleteBranchResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
