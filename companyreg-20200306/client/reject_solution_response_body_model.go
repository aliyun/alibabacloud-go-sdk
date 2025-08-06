// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectSolutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RejectSolutionResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *RejectSolutionResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *RejectSolutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RejectSolutionResponseBody
	GetSuccess() *bool
}

type RejectSolutionResponseBody struct {
	// example:
	//
	// PARTNER.CONFIG.NOT.FOUND
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 2174AA97-56FB-50FA-B243-0460B9E4CE0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RejectSolutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RejectSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *RejectSolutionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RejectSolutionResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *RejectSolutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RejectSolutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RejectSolutionResponseBody) SetErrorCode(v string) *RejectSolutionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RejectSolutionResponseBody) SetErrorMsg(v string) *RejectSolutionResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RejectSolutionResponseBody) SetRequestId(v string) *RejectSolutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RejectSolutionResponseBody) SetSuccess(v bool) *RejectSolutionResponseBody {
	s.Success = &v
	return s
}

func (s *RejectSolutionResponseBody) Validate() error {
	return dara.Validate(s)
}
