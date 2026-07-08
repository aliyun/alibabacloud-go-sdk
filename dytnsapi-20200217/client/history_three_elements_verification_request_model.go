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
	// The authorization code. You can obtain it from the following sources:
	//
	// - On the [Tag Plaza](https://dytns.console.aliyun.com/analysis/square) page in the Phone Number Intelligence console, select the **three-element ID verification*	- tag and submit an application. You will receive an authorization code after the application is approved.
	//
	// - On the [My Applications](https://dytns.console.aliyun.com/analysis/apply) page in the Phone Number Intelligence console, find the authorization ID for your approved **three-element ID verification*	- service.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The carrier to query. By default, the system queries the number\\"s carrier of record. Specify this parameter to route the query to a specific carrier.
	//
	// Valid values:
	//
	// - `CMCC`: China Mobile
	//
	// - `CUCC`: China Unicom
	//
	// - `CTCC`: China Telecom
	//
	// > Due to number portability, a ported number\\"s historical carrier may be unknown. Use this parameter to explicitly query a specific carrier. If omitted, the query defaults to the number\\"s current carrier of record.
	//
	// >
	//
	// > **Important*	- Specifying China Broadcasting Network is not supported and results in an HTTP 400 error.
	//
	// example:
	//
	// CMCC
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The ID number to verify.
	//
	// - If `Mask` is set to `NORMAL`, the value of this parameter is in plaintext.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	CertCode *string `json:"CertCode,omitempty" xml:"CertCode,omitempty"`
	// The phone number to query.
	//
	// - If `Mask` is set to `NORMAL`, this parameter must be an 11-digit mobile phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method. Valid value:
	//
	// - **NORMAL**: The phone number is not encrypted.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The name to verify.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The historical point in time to verify, in `yyyyMMddHHmmss` format. If the specific time of day is unknown, set the `HHmmss` portion to `000000`. For example, `20230615000000` verifies ownership as of June 15, 2023.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
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
