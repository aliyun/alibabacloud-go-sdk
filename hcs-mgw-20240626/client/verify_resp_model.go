// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyResp interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *VerifyResp
	GetErrorCode() *string
	SetErrorMsg(v string) *VerifyResp
	GetErrorMsg() *string
	SetHttpCode(v string) *VerifyResp
	GetHttpCode() *string
}

type VerifyResp struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 200
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
}

func (s VerifyResp) String() string {
	return dara.Prettify(s)
}

func (s VerifyResp) GoString() string {
	return s.String()
}

func (s *VerifyResp) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *VerifyResp) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *VerifyResp) GetHttpCode() *string {
	return s.HttpCode
}

func (s *VerifyResp) SetErrorCode(v string) *VerifyResp {
	s.ErrorCode = &v
	return s
}

func (s *VerifyResp) SetErrorMsg(v string) *VerifyResp {
	s.ErrorMsg = &v
	return s
}

func (s *VerifyResp) SetHttpCode(v string) *VerifyResp {
	s.HttpCode = &v
	return s
}

func (s *VerifyResp) Validate() error {
	return dara.Validate(s)
}
