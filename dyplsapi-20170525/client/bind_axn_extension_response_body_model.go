// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxnExtensionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BindAxnExtensionResponseBody
	GetCode() *string
	SetMessage(v string) *BindAxnExtensionResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindAxnExtensionResponseBody
	GetRequestId() *string
	SetSecretBindDTO(v *BindAxnExtensionResponseBodySecretBindDTO) *BindAxnExtensionResponseBody
	GetSecretBindDTO() *BindAxnExtensionResponseBodySecretBindDTO
}

type BindAxnExtensionResponseBody struct {
	// The request status code.
	//
	// 	- A value of OK indicates that the request is successful.
	//
	// 	- For other error codes, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
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
	// 9297B722-A016-43FB-B51A-E54050D9369D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The struct returned upon successful binding.
	SecretBindDTO *BindAxnExtensionResponseBodySecretBindDTO `json:"SecretBindDTO,omitempty" xml:"SecretBindDTO,omitempty" type:"Struct"`
}

func (s BindAxnExtensionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAxnExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindAxnExtensionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindAxnExtensionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAxnExtensionResponseBody) GetSecretBindDTO() *BindAxnExtensionResponseBodySecretBindDTO {
	return s.SecretBindDTO
}

func (s *BindAxnExtensionResponseBody) SetCode(v string) *BindAxnExtensionResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxnExtensionResponseBody) SetMessage(v string) *BindAxnExtensionResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxnExtensionResponseBody) SetRequestId(v string) *BindAxnExtensionResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxnExtensionResponseBody) SetSecretBindDTO(v *BindAxnExtensionResponseBodySecretBindDTO) *BindAxnExtensionResponseBody {
	s.SecretBindDTO = v
	return s
}

func (s *BindAxnExtensionResponseBody) Validate() error {
	if s.SecretBindDTO != nil {
		if err := s.SecretBindDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAxnExtensionResponseBodySecretBindDTO struct {
	// The extension number.
	//
	// example:
	//
	// 130
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The privacy number, which is number X.
	//
	// example:
	//
	// 139*****1111
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// The binding relationship ID.
	//
	// example:
	//
	// 1***************3
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s BindAxnExtensionResponseBodySecretBindDTO) String() string {
	return dara.Prettify(s)
}

func (s BindAxnExtensionResponseBodySecretBindDTO) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionResponseBodySecretBindDTO) GetExtension() *string {
	return s.Extension
}

func (s *BindAxnExtensionResponseBodySecretBindDTO) GetSecretNo() *string {
	return s.SecretNo
}

func (s *BindAxnExtensionResponseBodySecretBindDTO) GetSubsId() *string {
	return s.SubsId
}

func (s *BindAxnExtensionResponseBodySecretBindDTO) SetExtension(v string) *BindAxnExtensionResponseBodySecretBindDTO {
	s.Extension = &v
	return s
}

func (s *BindAxnExtensionResponseBodySecretBindDTO) SetSecretNo(v string) *BindAxnExtensionResponseBodySecretBindDTO {
	s.SecretNo = &v
	return s
}

func (s *BindAxnExtensionResponseBodySecretBindDTO) SetSubsId(v string) *BindAxnExtensionResponseBodySecretBindDTO {
	s.SubsId = &v
	return s
}

func (s *BindAxnExtensionResponseBodySecretBindDTO) Validate() error {
	return dara.Validate(s)
}
