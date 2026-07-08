// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVirtualThreeElementsVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *VirtualThreeElementsVerificationRequest
	GetAuthCode() *string
	SetCertCode(v string) *VirtualThreeElementsVerificationRequest
	GetCertCode() *string
	SetCertName(v string) *VirtualThreeElementsVerificationRequest
	GetCertName() *string
	SetInputNumber(v string) *VirtualThreeElementsVerificationRequest
	GetInputNumber() *string
	SetMask(v string) *VirtualThreeElementsVerificationRequest
	GetMask() *string
}

type VirtualThreeElementsVerificationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	CertCode *string `json:"CertCode,omitempty" xml:"CertCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
}

func (s VirtualThreeElementsVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s VirtualThreeElementsVerificationRequest) GoString() string {
	return s.String()
}

func (s *VirtualThreeElementsVerificationRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *VirtualThreeElementsVerificationRequest) GetCertCode() *string {
	return s.CertCode
}

func (s *VirtualThreeElementsVerificationRequest) GetCertName() *string {
	return s.CertName
}

func (s *VirtualThreeElementsVerificationRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *VirtualThreeElementsVerificationRequest) GetMask() *string {
	return s.Mask
}

func (s *VirtualThreeElementsVerificationRequest) SetAuthCode(v string) *VirtualThreeElementsVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *VirtualThreeElementsVerificationRequest) SetCertCode(v string) *VirtualThreeElementsVerificationRequest {
	s.CertCode = &v
	return s
}

func (s *VirtualThreeElementsVerificationRequest) SetCertName(v string) *VirtualThreeElementsVerificationRequest {
	s.CertName = &v
	return s
}

func (s *VirtualThreeElementsVerificationRequest) SetInputNumber(v string) *VirtualThreeElementsVerificationRequest {
	s.InputNumber = &v
	return s
}

func (s *VirtualThreeElementsVerificationRequest) SetMask(v string) *VirtualThreeElementsVerificationRequest {
	s.Mask = &v
	return s
}

func (s *VirtualThreeElementsVerificationRequest) Validate() error {
	return dara.Validate(s)
}
