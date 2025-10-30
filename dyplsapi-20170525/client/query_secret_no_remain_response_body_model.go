// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecretNoRemainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySecretNoRemainResponseBody
	GetCode() *string
	SetMessage(v string) *QuerySecretNoRemainResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySecretNoRemainResponseBody
	GetRequestId() *string
	SetSecretRemainDTO(v *QuerySecretNoRemainResponseBodySecretRemainDTO) *QuerySecretNoRemainResponseBody
	GetSecretRemainDTO() *QuerySecretNoRemainResponseBodySecretRemainDTO
}

type QuerySecretNoRemainResponseBody struct {
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
	// 9FC30594-3841-43AD-9008-03393BCB5CD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information returned after the operation was called.
	SecretRemainDTO *QuerySecretNoRemainResponseBodySecretRemainDTO `json:"SecretRemainDTO,omitempty" xml:"SecretRemainDTO,omitempty" type:"Struct"`
}

func (s QuerySecretNoRemainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretNoRemainResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySecretNoRemainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySecretNoRemainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySecretNoRemainResponseBody) GetSecretRemainDTO() *QuerySecretNoRemainResponseBodySecretRemainDTO {
	return s.SecretRemainDTO
}

func (s *QuerySecretNoRemainResponseBody) SetCode(v string) *QuerySecretNoRemainResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySecretNoRemainResponseBody) SetMessage(v string) *QuerySecretNoRemainResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySecretNoRemainResponseBody) SetRequestId(v string) *QuerySecretNoRemainResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySecretNoRemainResponseBody) SetSecretRemainDTO(v *QuerySecretNoRemainResponseBodySecretRemainDTO) *QuerySecretNoRemainResponseBody {
	s.SecretRemainDTO = v
	return s
}

func (s *QuerySecretNoRemainResponseBody) Validate() error {
	if s.SecretRemainDTO != nil {
		if err := s.SecretRemainDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySecretNoRemainResponseBodySecretRemainDTO struct {
	// The quantity of remaining phone numbers available for online purchase.
	//
	// example:
	//
	// 0
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The home location of the phone numbers.
	//
	// example:
	//
	// hangzhou
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The information about remaining phone numbers available for online purchase.
	RemainDTOList *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList `json:"RemainDTOList,omitempty" xml:"RemainDTOList,omitempty" type:"Struct"`
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTO) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTO) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTO) GetAmount() *int64 {
	return s.Amount
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTO) GetCity() *string {
	return s.City
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTO) GetRemainDTOList() *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList {
	return s.RemainDTOList
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTO) SetAmount(v int64) *QuerySecretNoRemainResponseBodySecretRemainDTO {
	s.Amount = &v
	return s
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTO) SetCity(v string) *QuerySecretNoRemainResponseBodySecretRemainDTO {
	s.City = &v
	return s
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTO) SetRemainDTOList(v *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList) *QuerySecretNoRemainResponseBodySecretRemainDTO {
	s.RemainDTOList = v
	return s
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTO) Validate() error {
	if s.RemainDTOList != nil {
		if err := s.RemainDTOList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList struct {
	RemainDTO []*QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO `json:"remainDTO,omitempty" xml:"remainDTO,omitempty" type:"Repeated"`
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList) GetRemainDTO() []*QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO {
	return s.RemainDTO
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList) SetRemainDTO(v []*QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList {
	s.RemainDTO = v
	return s
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOList) Validate() error {
	if s.RemainDTO != nil {
		for _, item := range s.RemainDTO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO struct {
	// The quantity of remaining phone numbers available for online purchase for the city.
	//
	// example:
	//
	// 120
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The home location of the phone numbers.
	//
	// example:
	//
	// Wuhan
	City *string `json:"City,omitempty" xml:"City,omitempty"`
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) GetAmount() *int64 {
	return s.Amount
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) GetCity() *string {
	return s.City
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) SetAmount(v int64) *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO {
	s.Amount = &v
	return s
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) SetCity(v string) *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO {
	s.City = &v
	return s
}

func (s *QuerySecretNoRemainResponseBodySecretRemainDTORemainDTOListRemainDTO) Validate() error {
	return dara.Validate(s)
}
