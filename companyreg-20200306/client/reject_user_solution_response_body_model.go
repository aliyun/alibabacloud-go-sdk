// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectUserSolutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RejectUserSolutionResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *RejectUserSolutionResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *RejectUserSolutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RejectUserSolutionResponseBody
	GetSuccess() *bool
}

type RejectUserSolutionResponseBody struct {
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 717711FB-F887-597B-8121-B77437E89B97
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RejectUserSolutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RejectUserSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *RejectUserSolutionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RejectUserSolutionResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *RejectUserSolutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RejectUserSolutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RejectUserSolutionResponseBody) SetErrorCode(v string) *RejectUserSolutionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RejectUserSolutionResponseBody) SetErrorMsg(v string) *RejectUserSolutionResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RejectUserSolutionResponseBody) SetRequestId(v string) *RejectUserSolutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RejectUserSolutionResponseBody) SetSuccess(v bool) *RejectUserSolutionResponseBody {
	s.Success = &v
	return s
}

func (s *RejectUserSolutionResponseBody) Validate() error {
	return dara.Validate(s)
}
