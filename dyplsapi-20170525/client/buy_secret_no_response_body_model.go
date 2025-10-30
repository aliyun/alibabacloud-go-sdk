// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuySecretNoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BuySecretNoResponseBody
	GetCode() *string
	SetMessage(v string) *BuySecretNoResponseBody
	GetMessage() *string
	SetRequestId(v string) *BuySecretNoResponseBody
	GetRequestId() *string
	SetSecretBuyInfoDTO(v *BuySecretNoResponseBodySecretBuyInfoDTO) *BuySecretNoResponseBody
	GetSecretBuyInfoDTO() *BuySecretNoResponseBodySecretBuyInfoDTO
}

type BuySecretNoResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
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
	// 2D1AEB96-96D0-454E-B0DC-AE2A8DF08020
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information returned after the operation was called.
	SecretBuyInfoDTO *BuySecretNoResponseBodySecretBuyInfoDTO `json:"SecretBuyInfoDTO,omitempty" xml:"SecretBuyInfoDTO,omitempty" type:"Struct"`
}

func (s BuySecretNoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BuySecretNoResponseBody) GoString() string {
	return s.String()
}

func (s *BuySecretNoResponseBody) GetCode() *string {
	return s.Code
}

func (s *BuySecretNoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BuySecretNoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BuySecretNoResponseBody) GetSecretBuyInfoDTO() *BuySecretNoResponseBodySecretBuyInfoDTO {
	return s.SecretBuyInfoDTO
}

func (s *BuySecretNoResponseBody) SetCode(v string) *BuySecretNoResponseBody {
	s.Code = &v
	return s
}

func (s *BuySecretNoResponseBody) SetMessage(v string) *BuySecretNoResponseBody {
	s.Message = &v
	return s
}

func (s *BuySecretNoResponseBody) SetRequestId(v string) *BuySecretNoResponseBody {
	s.RequestId = &v
	return s
}

func (s *BuySecretNoResponseBody) SetSecretBuyInfoDTO(v *BuySecretNoResponseBodySecretBuyInfoDTO) *BuySecretNoResponseBody {
	s.SecretBuyInfoDTO = v
	return s
}

func (s *BuySecretNoResponseBody) Validate() error {
	if s.SecretBuyInfoDTO != nil {
		if err := s.SecretBuyInfoDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BuySecretNoResponseBodySecretBuyInfoDTO struct {
	// The private number, that is, phone number X.
	//
	// example:
	//
	// 1390000****
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s BuySecretNoResponseBodySecretBuyInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s BuySecretNoResponseBodySecretBuyInfoDTO) GoString() string {
	return s.String()
}

func (s *BuySecretNoResponseBodySecretBuyInfoDTO) GetSecretNo() *string {
	return s.SecretNo
}

func (s *BuySecretNoResponseBodySecretBuyInfoDTO) SetSecretNo(v string) *BuySecretNoResponseBodySecretBuyInfoDTO {
	s.SecretNo = &v
	return s
}

func (s *BuySecretNoResponseBodySecretBuyInfoDTO) Validate() error {
	return dara.Validate(s)
}
