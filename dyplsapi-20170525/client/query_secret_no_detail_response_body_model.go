// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecretNoDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySecretNoDetailResponseBody
	GetCode() *string
	SetMessage(v string) *QuerySecretNoDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySecretNoDetailResponseBody
	GetRequestId() *string
	SetSecretNoInfoDTO(v *QuerySecretNoDetailResponseBodySecretNoInfoDTO) *QuerySecretNoDetailResponseBody
	GetSecretNoInfoDTO() *QuerySecretNoDetailResponseBodySecretNoInfoDTO
}

type QuerySecretNoDetailResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 066E6E47-04CB-4774-A976-4F73CB76D4A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The attributes of the phone number.
	SecretNoInfoDTO *QuerySecretNoDetailResponseBodySecretNoInfoDTO `json:"SecretNoInfoDTO,omitempty" xml:"SecretNoInfoDTO,omitempty" type:"Struct"`
}

func (s QuerySecretNoDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretNoDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySecretNoDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySecretNoDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySecretNoDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySecretNoDetailResponseBody) GetSecretNoInfoDTO() *QuerySecretNoDetailResponseBodySecretNoInfoDTO {
	return s.SecretNoInfoDTO
}

func (s *QuerySecretNoDetailResponseBody) SetCode(v string) *QuerySecretNoDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySecretNoDetailResponseBody) SetMessage(v string) *QuerySecretNoDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySecretNoDetailResponseBody) SetRequestId(v string) *QuerySecretNoDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySecretNoDetailResponseBody) SetSecretNoInfoDTO(v *QuerySecretNoDetailResponseBodySecretNoInfoDTO) *QuerySecretNoDetailResponseBody {
	s.SecretNoInfoDTO = v
	return s
}

func (s *QuerySecretNoDetailResponseBody) Validate() error {
	if s.SecretNoInfoDTO != nil {
		if err := s.SecretNoInfoDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySecretNoDetailResponseBodySecretNoInfoDTO struct {
	// The verification status of the phone number. Valid values:
	//
	// 	- **0**: The phone number is not verified.
	//
	// 	- **1**: The phone number is verified.
	//
	// example:
	//
	// 0
	CertifyStatus *int32 `json:"CertifyStatus,omitempty" xml:"CertifyStatus,omitempty"`
	// The city.
	//
	// example:
	//
	// chengdu
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The province.
	//
	// example:
	//
	// sichuan
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The time when the phone number was purchased.
	//
	// example:
	//
	// 2021-12-03 15:19:27
	PurchaseTime *string `json:"PurchaseTime,omitempty" xml:"PurchaseTime,omitempty"`
	// The status of the phone number. Valid values:
	//
	// 	- **0**: The phone number is not bound to other phone numbers.
	//
	// 	- **1**: The phone number is bound to other phone numbers.
	//
	// 	- **2**: The phone number is locked.
	//
	// 	- **3**: The phone number is frozen.
	//
	// example:
	//
	// 0
	SecretStatus *int64 `json:"SecretStatus,omitempty" xml:"SecretStatus,omitempty"`
	// The carrier to which the phone number belongs. Valid values:
	//
	// 	- **1**: China Mobile
	//
	// 	- **2**: China Unicom
	//
	// 	- **3**: China Telecom
	//
	// example:
	//
	// 1
	Vendor *int64 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s QuerySecretNoDetailResponseBodySecretNoInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretNoDetailResponseBodySecretNoInfoDTO) GoString() string {
	return s.String()
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) GetCertifyStatus() *int32 {
	return s.CertifyStatus
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) GetCity() *string {
	return s.City
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) GetProvince() *string {
	return s.Province
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) GetPurchaseTime() *string {
	return s.PurchaseTime
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) GetSecretStatus() *int64 {
	return s.SecretStatus
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) GetVendor() *int64 {
	return s.Vendor
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) SetCertifyStatus(v int32) *QuerySecretNoDetailResponseBodySecretNoInfoDTO {
	s.CertifyStatus = &v
	return s
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) SetCity(v string) *QuerySecretNoDetailResponseBodySecretNoInfoDTO {
	s.City = &v
	return s
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) SetProvince(v string) *QuerySecretNoDetailResponseBodySecretNoInfoDTO {
	s.Province = &v
	return s
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) SetPurchaseTime(v string) *QuerySecretNoDetailResponseBodySecretNoInfoDTO {
	s.PurchaseTime = &v
	return s
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) SetSecretStatus(v int64) *QuerySecretNoDetailResponseBodySecretNoInfoDTO {
	s.SecretStatus = &v
	return s
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) SetVendor(v int64) *QuerySecretNoDetailResponseBodySecretNoInfoDTO {
	s.Vendor = &v
	return s
}

func (s *QuerySecretNoDetailResponseBodySecretNoInfoDTO) Validate() error {
	return dara.Validate(s)
}
