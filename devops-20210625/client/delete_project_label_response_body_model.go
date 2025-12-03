// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteProjectLabelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteProjectLabelResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteProjectLabelResponseBody
	GetRequestId() *string
	SetResult(v *DeleteProjectLabelResponseBodyResult) *DeleteProjectLabelResponseBody
	GetResult() *DeleteProjectLabelResponseBodyResult
	SetSuccess(v bool) *DeleteProjectLabelResponseBody
	GetSuccess() *bool
}

type DeleteProjectLabelResponseBody struct {
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
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DeleteProjectLabelResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteProjectLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectLabelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectLabelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteProjectLabelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteProjectLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProjectLabelResponseBody) GetResult() *DeleteProjectLabelResponseBodyResult {
	return s.Result
}

func (s *DeleteProjectLabelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteProjectLabelResponseBody) SetErrorCode(v string) *DeleteProjectLabelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteProjectLabelResponseBody) SetErrorMessage(v string) *DeleteProjectLabelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteProjectLabelResponseBody) SetRequestId(v string) *DeleteProjectLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectLabelResponseBody) SetResult(v *DeleteProjectLabelResponseBodyResult) *DeleteProjectLabelResponseBody {
	s.Result = v
	return s
}

func (s *DeleteProjectLabelResponseBody) SetSuccess(v bool) *DeleteProjectLabelResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteProjectLabelResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteProjectLabelResponseBodyResult struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteProjectLabelResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectLabelResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteProjectLabelResponseBodyResult) GetResult() *bool {
	return s.Result
}

func (s *DeleteProjectLabelResponseBodyResult) SetResult(v bool) *DeleteProjectLabelResponseBodyResult {
	s.Result = &v
	return s
}

func (s *DeleteProjectLabelResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
