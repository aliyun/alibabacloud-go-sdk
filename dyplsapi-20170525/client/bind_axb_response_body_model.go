// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BindAxbResponseBody
	GetCode() *string
	SetMessage(v string) *BindAxbResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindAxbResponseBody
	GetRequestId() *string
	SetSecretBindDTO(v *BindAxbResponseBodySecretBindDTO) *BindAxbResponseBody
	GetSecretBindDTO() *BindAxbResponseBodySecretBindDTO
}

type BindAxbResponseBody struct {
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
	// 9297B722-A016-43FB-B51A-E54050D9369D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information returned after the phone numbers were bound.
	SecretBindDTO *BindAxbResponseBodySecretBindDTO `json:"SecretBindDTO,omitempty" xml:"SecretBindDTO,omitempty" type:"Struct"`
}

func (s BindAxbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAxbResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxbResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindAxbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindAxbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAxbResponseBody) GetSecretBindDTO() *BindAxbResponseBodySecretBindDTO {
	return s.SecretBindDTO
}

func (s *BindAxbResponseBody) SetCode(v string) *BindAxbResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxbResponseBody) SetMessage(v string) *BindAxbResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxbResponseBody) SetRequestId(v string) *BindAxbResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxbResponseBody) SetSecretBindDTO(v *BindAxbResponseBodySecretBindDTO) *BindAxbResponseBody {
	s.SecretBindDTO = v
	return s
}

func (s *BindAxbResponseBody) Validate() error {
	if s.SecretBindDTO != nil {
		if err := s.SecretBindDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAxbResponseBodySecretBindDTO struct {
	// The extension of the phone number.
	//
	// >  The BindAxb operation does not involve an extension. Ignore this parameter.
	//
	// example:
	//
	// 130
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
	// 1**************3
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s BindAxbResponseBodySecretBindDTO) String() string {
	return dara.Prettify(s)
}

func (s BindAxbResponseBodySecretBindDTO) GoString() string {
	return s.String()
}

func (s *BindAxbResponseBodySecretBindDTO) GetExtension() *string {
	return s.Extension
}

func (s *BindAxbResponseBodySecretBindDTO) GetSecretNo() *string {
	return s.SecretNo
}

func (s *BindAxbResponseBodySecretBindDTO) GetSubsId() *string {
	return s.SubsId
}

func (s *BindAxbResponseBodySecretBindDTO) SetExtension(v string) *BindAxbResponseBodySecretBindDTO {
	s.Extension = &v
	return s
}

func (s *BindAxbResponseBodySecretBindDTO) SetSecretNo(v string) *BindAxbResponseBodySecretBindDTO {
	s.SecretNo = &v
	return s
}

func (s *BindAxbResponseBodySecretBindDTO) SetSubsId(v string) *BindAxbResponseBodySecretBindDTO {
	s.SubsId = &v
	return s
}

func (s *BindAxbResponseBodySecretBindDTO) Validate() error {
	return dara.Validate(s)
}
