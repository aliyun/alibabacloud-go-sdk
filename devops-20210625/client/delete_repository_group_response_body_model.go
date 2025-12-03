// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepositoryGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteRepositoryGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteRepositoryGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteRepositoryGroupResponseBody
	GetRequestId() *string
	SetResult(v *DeleteRepositoryGroupResponseBodyResult) *DeleteRepositoryGroupResponseBody
	GetResult() *DeleteRepositoryGroupResponseBodyResult
	SetSuccess(v bool) *DeleteRepositoryGroupResponseBody
	GetSuccess() *bool
}

type DeleteRepositoryGroupResponseBody struct {
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
	// 30F2DA15-2877-5FC9-BC71-F7F394717907
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DeleteRepositoryGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteRepositoryGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteRepositoryGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteRepositoryGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRepositoryGroupResponseBody) GetResult() *DeleteRepositoryGroupResponseBodyResult {
	return s.Result
}

func (s *DeleteRepositoryGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRepositoryGroupResponseBody) SetErrorCode(v string) *DeleteRepositoryGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryGroupResponseBody) SetErrorMessage(v string) *DeleteRepositoryGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryGroupResponseBody) SetRequestId(v string) *DeleteRepositoryGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryGroupResponseBody) SetResult(v *DeleteRepositoryGroupResponseBodyResult) *DeleteRepositoryGroupResponseBody {
	s.Result = v
	return s
}

func (s *DeleteRepositoryGroupResponseBody) SetSuccess(v bool) *DeleteRepositoryGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRepositoryGroupResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteRepositoryGroupResponseBodyResult struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteRepositoryGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryGroupResponseBodyResult) GetResult() *bool {
	return s.Result
}

func (s *DeleteRepositoryGroupResponseBodyResult) SetResult(v bool) *DeleteRepositoryGroupResponseBodyResult {
	s.Result = &v
	return s
}

func (s *DeleteRepositoryGroupResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
