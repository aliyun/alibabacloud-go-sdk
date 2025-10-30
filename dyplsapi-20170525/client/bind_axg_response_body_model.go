// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BindAxgResponseBody
	GetCode() *string
	SetMessage(v string) *BindAxgResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindAxgResponseBody
	GetRequestId() *string
	SetSecretBindDTO(v *BindAxgResponseBodySecretBindDTO) *BindAxgResponseBody
	GetSecretBindDTO() *BindAxgResponseBodySecretBindDTO
}

type BindAxgResponseBody struct {
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
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information returned after the phone numbers were bound.
	SecretBindDTO *BindAxgResponseBodySecretBindDTO `json:"SecretBindDTO,omitempty" xml:"SecretBindDTO,omitempty" type:"Struct"`
}

func (s BindAxgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAxgResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxgResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindAxgResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindAxgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAxgResponseBody) GetSecretBindDTO() *BindAxgResponseBodySecretBindDTO {
	return s.SecretBindDTO
}

func (s *BindAxgResponseBody) SetCode(v string) *BindAxgResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxgResponseBody) SetMessage(v string) *BindAxgResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxgResponseBody) SetRequestId(v string) *BindAxgResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxgResponseBody) SetSecretBindDTO(v *BindAxgResponseBodySecretBindDTO) *BindAxgResponseBody {
	s.SecretBindDTO = v
	return s
}

func (s *BindAxgResponseBody) Validate() error {
	if s.SecretBindDTO != nil {
		if err := s.SecretBindDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAxgResponseBodySecretBindDTO struct {
	// The extension of the phone number.
	//
	// >  The BindAxg operation does not involve an extension. Ignore this parameter.
	//
	// example:
	//
	// 139****0000
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The private number, that is, phone number X.
	//
	// example:
	//
	// 139****0000
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// The binding ID.
	//
	// example:
	//
	// 1************3
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s BindAxgResponseBodySecretBindDTO) String() string {
	return dara.Prettify(s)
}

func (s BindAxgResponseBodySecretBindDTO) GoString() string {
	return s.String()
}

func (s *BindAxgResponseBodySecretBindDTO) GetExtension() *string {
	return s.Extension
}

func (s *BindAxgResponseBodySecretBindDTO) GetSecretNo() *string {
	return s.SecretNo
}

func (s *BindAxgResponseBodySecretBindDTO) GetSubsId() *string {
	return s.SubsId
}

func (s *BindAxgResponseBodySecretBindDTO) SetExtension(v string) *BindAxgResponseBodySecretBindDTO {
	s.Extension = &v
	return s
}

func (s *BindAxgResponseBodySecretBindDTO) SetSecretNo(v string) *BindAxgResponseBodySecretBindDTO {
	s.SecretNo = &v
	return s
}

func (s *BindAxgResponseBodySecretBindDTO) SetSubsId(v string) *BindAxgResponseBodySecretBindDTO {
	s.SubsId = &v
	return s
}

func (s *BindAxgResponseBodySecretBindDTO) Validate() error {
	return dara.Validate(s)
}
