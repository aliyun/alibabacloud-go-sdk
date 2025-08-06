// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSolutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfirmUrl(v string) *SubmitSolutionResponseBody
	GetConfirmUrl() *string
	SetErrorCode(v string) *SubmitSolutionResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *SubmitSolutionResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *SubmitSolutionResponseBody
	GetRequestId() *string
	SetSolutionBizId(v string) *SubmitSolutionResponseBody
	GetSolutionBizId() *string
	SetSuccess(v bool) *SubmitSolutionResponseBody
	GetSuccess() *bool
}

type SubmitSolutionResponseBody struct {
	// example:
	//
	// https://companyreg.console.aliyun.com/#/intention-notarize?Type=119&bizid=I20220114181457000001
	ConfirmUrl *string `json:"ConfirmUrl,omitempty" xml:"ConfirmUrl,omitempty"`
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 0A3CFCF5-E0C0-5C0B-A2ED-03827F16D85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// S20211109140729000001
	SolutionBizId *string `json:"SolutionBizId,omitempty" xml:"SolutionBizId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitSolutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSolutionResponseBody) GetConfirmUrl() *string {
	return s.ConfirmUrl
}

func (s *SubmitSolutionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitSolutionResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *SubmitSolutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSolutionResponseBody) GetSolutionBizId() *string {
	return s.SolutionBizId
}

func (s *SubmitSolutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitSolutionResponseBody) SetConfirmUrl(v string) *SubmitSolutionResponseBody {
	s.ConfirmUrl = &v
	return s
}

func (s *SubmitSolutionResponseBody) SetErrorCode(v string) *SubmitSolutionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitSolutionResponseBody) SetErrorMsg(v string) *SubmitSolutionResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SubmitSolutionResponseBody) SetRequestId(v string) *SubmitSolutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSolutionResponseBody) SetSolutionBizId(v string) *SubmitSolutionResponseBody {
	s.SolutionBizId = &v
	return s
}

func (s *SubmitSolutionResponseBody) SetSuccess(v bool) *SubmitSolutionResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSolutionResponseBody) Validate() error {
	return dara.Validate(s)
}
