// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyAddressResp interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *VerifyAddressResp
	GetErrorCode() *string
	SetErrorMessage(v string) *VerifyAddressResp
	GetErrorMessage() *string
	SetStatus(v string) *VerifyAddressResp
	GetStatus() *string
	SetVerifyTime(v string) *VerifyAddressResp
	GetVerifyTime() *string
}

type VerifyAddressResp struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// avaliable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	VerifyTime *string `json:"VerifyTime,omitempty" xml:"VerifyTime,omitempty"`
}

func (s VerifyAddressResp) String() string {
	return dara.Prettify(s)
}

func (s VerifyAddressResp) GoString() string {
	return s.String()
}

func (s *VerifyAddressResp) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *VerifyAddressResp) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *VerifyAddressResp) GetStatus() *string {
	return s.Status
}

func (s *VerifyAddressResp) GetVerifyTime() *string {
	return s.VerifyTime
}

func (s *VerifyAddressResp) SetErrorCode(v string) *VerifyAddressResp {
	s.ErrorCode = &v
	return s
}

func (s *VerifyAddressResp) SetErrorMessage(v string) *VerifyAddressResp {
	s.ErrorMessage = &v
	return s
}

func (s *VerifyAddressResp) SetStatus(v string) *VerifyAddressResp {
	s.Status = &v
	return s
}

func (s *VerifyAddressResp) SetVerifyTime(v string) *VerifyAddressResp {
	s.VerifyTime = &v
	return s
}

func (s *VerifyAddressResp) Validate() error {
	return dara.Validate(s)
}
