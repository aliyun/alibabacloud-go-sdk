// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BindAxnResponseBody
	GetCode() *string
	SetMessage(v string) *BindAxnResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindAxnResponseBody
	GetRequestId() *string
	SetSecretBindDTO(v *BindAxnResponseBodySecretBindDTO) *BindAxnResponseBody
	GetSecretBindDTO() *BindAxnResponseBodySecretBindDTO
}

type BindAxnResponseBody struct {
	// The status code of the request.
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
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The struct returned after the number is successfully bound.
	SecretBindDTO *BindAxnResponseBodySecretBindDTO `json:"SecretBindDTO,omitempty" xml:"SecretBindDTO,omitempty" type:"Struct"`
}

func (s BindAxnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAxnResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxnResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindAxnResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindAxnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAxnResponseBody) GetSecretBindDTO() *BindAxnResponseBodySecretBindDTO {
	return s.SecretBindDTO
}

func (s *BindAxnResponseBody) SetCode(v string) *BindAxnResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxnResponseBody) SetMessage(v string) *BindAxnResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxnResponseBody) SetRequestId(v string) *BindAxnResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxnResponseBody) SetSecretBindDTO(v *BindAxnResponseBodySecretBindDTO) *BindAxnResponseBody {
	s.SecretBindDTO = v
	return s
}

func (s *BindAxnResponseBody) Validate() error {
	if s.SecretBindDTO != nil {
		if err := s.SecretBindDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAxnResponseBodySecretBindDTO struct {
	// The extension number.
	//
	// > The BindAxn operation does not involve extension numbers. Ignore this parameter.
	//
	// example:
	//
	// 130
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The privacy number, which is Number X.
	//
	// example:
	//
	// 139****1111
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// The ID of the binding relationship.
	//
	// example:
	//
	// 1***************3
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s BindAxnResponseBodySecretBindDTO) String() string {
	return dara.Prettify(s)
}

func (s BindAxnResponseBodySecretBindDTO) GoString() string {
	return s.String()
}

func (s *BindAxnResponseBodySecretBindDTO) GetExtension() *string {
	return s.Extension
}

func (s *BindAxnResponseBodySecretBindDTO) GetSecretNo() *string {
	return s.SecretNo
}

func (s *BindAxnResponseBodySecretBindDTO) GetSubsId() *string {
	return s.SubsId
}

func (s *BindAxnResponseBodySecretBindDTO) SetExtension(v string) *BindAxnResponseBodySecretBindDTO {
	s.Extension = &v
	return s
}

func (s *BindAxnResponseBodySecretBindDTO) SetSecretNo(v string) *BindAxnResponseBodySecretBindDTO {
	s.SecretNo = &v
	return s
}

func (s *BindAxnResponseBodySecretBindDTO) SetSubsId(v string) *BindAxnResponseBodySecretBindDTO {
	s.SubsId = &v
	return s
}

func (s *BindAxnResponseBodySecretBindDTO) Validate() error {
	return dara.Validate(s)
}
