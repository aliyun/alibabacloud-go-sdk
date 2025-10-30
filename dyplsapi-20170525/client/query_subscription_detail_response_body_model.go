// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySubscriptionDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySubscriptionDetailResponseBody
	GetCode() *string
	SetMessage(v string) *QuerySubscriptionDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySubscriptionDetailResponseBody
	GetRequestId() *string
	SetSecretBindDetailDTO(v *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) *QuerySubscriptionDetailResponseBody
	GetSecretBindDetailDTO() *QuerySubscriptionDetailResponseBodySecretBindDetailDTO
}

type QuerySubscriptionDetailResponseBody struct {
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
	// The information returned after the QuerySubscriptionDetail operation was called.
	SecretBindDetailDTO *QuerySubscriptionDetailResponseBodySecretBindDetailDTO `json:"SecretBindDetailDTO,omitempty" xml:"SecretBindDetailDTO,omitempty" type:"Struct"`
}

func (s QuerySubscriptionDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySubscriptionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySubscriptionDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySubscriptionDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySubscriptionDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySubscriptionDetailResponseBody) GetSecretBindDetailDTO() *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	return s.SecretBindDetailDTO
}

func (s *QuerySubscriptionDetailResponseBody) SetCode(v string) *QuerySubscriptionDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBody) SetMessage(v string) *QuerySubscriptionDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBody) SetRequestId(v string) *QuerySubscriptionDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBody) SetSecretBindDetailDTO(v *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) *QuerySubscriptionDetailResponseBody {
	s.SecretBindDetailDTO = v
	return s
}

func (s *QuerySubscriptionDetailResponseBody) Validate() error {
	if s.SecretBindDetailDTO != nil {
		if err := s.SecretBindDetailDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySubscriptionDetailResponseBodySecretBindDetailDTO struct {
	// The ID of the ASR model.
	//
	// example:
	//
	// 123456
	ASRModelId *string `json:"ASRModelId,omitempty" xml:"ASRModelId,omitempty"`
	// Indicates whether automatic speech recognition (ASR) is enabled. Valid values:
	//
	// 	- **false**: ASR is disabled.
	//
	// 	- **true**: ASR is enabled.
	//
	// example:
	//
	// true
	ASRStatus *bool `json:"ASRStatus,omitempty" xml:"ASRStatus,omitempty"`
	// The status of one-way call restrictions. No value is returned for this parameter if one-way call restrictions are not set. Valid values:
	//
	// 	- **CONTROL_AX_DISABLE**: Phone number A cannot be used to call phone number X.
	//
	// 	- **CONTROL_BX_DISABLE**: Phone number B cannot be used to call phone number X.
	//
	// example:
	//
	// CONTROL_BX_DISABLE
	CallRestrict *string `json:"CallRestrict,omitempty" xml:"CallRestrict,omitempty"`
	// The expiration time of the binding.
	//
	// example:
	//
	// 2019-09-05 12:00:00
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The extension in the AXG extension binding.
	//
	// example:
	//
	// 130
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The creation time of the binding.
	//
	// example:
	//
	// 2019-03-05 12:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of number group G in the binding.
	//
	// example:
	//
	// 2000000130001
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Indicates whether all calls made by the bound phone numbers are recorded. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	NeedRecord *bool `json:"NeedRecord,omitempty" xml:"NeedRecord,omitempty"`
	// Phone number A in the binding.
	//
	// example:
	//
	// 13900001111
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// Phone number B in the binding.
	//
	// example:
	//
	// 13900002222
	PhoneNoB *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	// The private number in the binding, that is, phone number X.
	//
	// example:
	//
	// 13900001234
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// The binding status. Valid values:
	//
	// 	- **0**: The binding expired.
	//
	// 	- **1**: The binding is in effect.
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The binding ID.
	//
	// example:
	//
	// 100000076879****
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s QuerySubscriptionDetailResponseBodySecretBindDetailDTO) String() string {
	return dara.Prettify(s)
}

func (s QuerySubscriptionDetailResponseBodySecretBindDetailDTO) GoString() string {
	return s.String()
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) GetASRModelId() *string {
	return s.ASRModelId
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) GetASRStatus() *bool {
	return s.ASRStatus
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) GetCallRestrict() *string {
	return s.CallRestrict
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) GetExtension() *string {
	return s.Extension
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) GetGroupId() *int64 {
	return s.GroupId
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) GetNeedRecord() *bool {
	return s.NeedRecord
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) GetPhoneNoA() *string {
	return s.PhoneNoA
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) GetPhoneNoB() *string {
	return s.PhoneNoB
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) GetPhoneNoX() *string {
	return s.PhoneNoX
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) GetStatus() *int64 {
	return s.Status
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) GetSubsId() *string {
	return s.SubsId
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetASRModelId(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.ASRModelId = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetASRStatus(v bool) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.ASRStatus = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetCallRestrict(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.CallRestrict = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetExpireDate(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.ExpireDate = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetExtension(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.Extension = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetGmtCreate(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.GmtCreate = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetGroupId(v int64) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.GroupId = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetNeedRecord(v bool) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.NeedRecord = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetPhoneNoA(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.PhoneNoA = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetPhoneNoB(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.PhoneNoB = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetPhoneNoX(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.PhoneNoX = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetStatus(v int64) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.Status = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) SetSubsId(v string) *QuerySubscriptionDetailResponseBodySecretBindDetailDTO {
	s.SubsId = &v
	return s
}

func (s *QuerySubscriptionDetailResponseBodySecretBindDetailDTO) Validate() error {
	return dara.Validate(s)
}
