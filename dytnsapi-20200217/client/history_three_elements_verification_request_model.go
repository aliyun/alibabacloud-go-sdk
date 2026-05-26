// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHistoryThreeElementsVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *HistoryThreeElementsVerificationRequest
	GetAuthCode() *string
	SetCarrier(v string) *HistoryThreeElementsVerificationRequest
	GetCarrier() *string
	SetCertCode(v string) *HistoryThreeElementsVerificationRequest
	GetCertCode() *string
	SetInputNumber(v string) *HistoryThreeElementsVerificationRequest
	GetInputNumber() *string
	SetMask(v string) *HistoryThreeElementsVerificationRequest
	GetMask() *string
	SetName(v string) *HistoryThreeElementsVerificationRequest
	GetName() *string
	SetVerificationTime(v string) *HistoryThreeElementsVerificationRequest
	GetVerificationTime() *string
}

type HistoryThreeElementsVerificationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// example:
	//
	// CMCC
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	CertCode *string `json:"CertCode,omitempty" xml:"CertCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	VerificationTime *string `json:"VerificationTime,omitempty" xml:"VerificationTime,omitempty"`
}

func (s HistoryThreeElementsVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s HistoryThreeElementsVerificationRequest) GoString() string {
	return s.String()
}

func (s *HistoryThreeElementsVerificationRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *HistoryThreeElementsVerificationRequest) GetCarrier() *string {
	return s.Carrier
}

func (s *HistoryThreeElementsVerificationRequest) GetCertCode() *string {
	return s.CertCode
}

func (s *HistoryThreeElementsVerificationRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *HistoryThreeElementsVerificationRequest) GetMask() *string {
	return s.Mask
}

func (s *HistoryThreeElementsVerificationRequest) GetName() *string {
	return s.Name
}

func (s *HistoryThreeElementsVerificationRequest) GetVerificationTime() *string {
	return s.VerificationTime
}

func (s *HistoryThreeElementsVerificationRequest) SetAuthCode(v string) *HistoryThreeElementsVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *HistoryThreeElementsVerificationRequest) SetCarrier(v string) *HistoryThreeElementsVerificationRequest {
	s.Carrier = &v
	return s
}

func (s *HistoryThreeElementsVerificationRequest) SetCertCode(v string) *HistoryThreeElementsVerificationRequest {
	s.CertCode = &v
	return s
}

func (s *HistoryThreeElementsVerificationRequest) SetInputNumber(v string) *HistoryThreeElementsVerificationRequest {
	s.InputNumber = &v
	return s
}

func (s *HistoryThreeElementsVerificationRequest) SetMask(v string) *HistoryThreeElementsVerificationRequest {
	s.Mask = &v
	return s
}

func (s *HistoryThreeElementsVerificationRequest) SetName(v string) *HistoryThreeElementsVerificationRequest {
	s.Name = &v
	return s
}

func (s *HistoryThreeElementsVerificationRequest) SetVerificationTime(v string) *HistoryThreeElementsVerificationRequest {
	s.VerificationTime = &v
	return s
}

func (s *HistoryThreeElementsVerificationRequest) Validate() error {
	return dara.Validate(s)
}
