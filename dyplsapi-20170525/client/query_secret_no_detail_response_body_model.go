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
	// The request status code.
	//
	// - OK indicates that the request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the status code.
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
	// The number properties.
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
	// The number verification status. Valid values:
	//
	// - **0**: Unverified.
	//
	// - **1**: Verified.
	//
	// example:
	//
	// 0
	CertifyStatus *int32 `json:"CertifyStatus,omitempty" xml:"CertifyStatus,omitempty"`
	// The city.
	//
	// example:
	//
	// Xi\\"an
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The province.
	//
	// example:
	//
	// Shaanxi
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The purchase time of the number, in the format of YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2021-12-03 15:19:27
	PurchaseTime *string `json:"PurchaseTime,omitempty" xml:"PurchaseTime,omitempty"`
	// The number status. Valid values:
	//
	// - **0**: Unbound.
	//
	// - **1**: Bound.
	//
	// - **2**: Locked.
	//
	// - **3**: Frozen.
	//
	// example:
	//
	// 0
	SecretStatus *int64 `json:"SecretStatus,omitempty" xml:"SecretStatus,omitempty"`
	// The carrier. Valid values:
	//
	// - **1**: China Mobile
	//
	// - **2**: China Unicom
	//
	// - **3**: China Telecom
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
