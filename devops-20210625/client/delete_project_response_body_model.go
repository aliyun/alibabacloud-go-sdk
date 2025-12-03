// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteProjectResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *DeleteProjectResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *DeleteProjectResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteProjectResponseBody
	GetResult() *bool
	SetSuccess(v bool) *DeleteProjectResponseBody
	GetSuccess() *bool
}

type DeleteProjectResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteProjectResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DeleteProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProjectResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteProjectResponseBody) SetErrorCode(v string) *DeleteProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteProjectResponseBody) SetErrorMsg(v string) *DeleteProjectResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectResponseBody) SetResult(v bool) *DeleteProjectResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteProjectResponseBody) SetSuccess(v bool) *DeleteProjectResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
