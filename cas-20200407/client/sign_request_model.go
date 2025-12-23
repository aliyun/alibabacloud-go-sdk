// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *SignRequest
	GetCertIdentifier() *string
	SetCustomIdentifier(v string) *SignRequest
	GetCustomIdentifier() *string
	SetMessage(v string) *SignRequest
	GetMessage() *string
	SetMessageType(v string) *SignRequest
	GetMessageType() *string
	SetSigningAlgorithm(v string) *SignRequest
	GetSigningAlgorithm() *string
}

type SignRequest struct {
	// The unique identifier of the certificate. You can call the [ListCert](https://help.aliyun.com/document_detail/455806.html) operation to obtain the identifier.
	//
	// 	- If the certificate is an SSL certificate, the value of this parameter must be in the {Certificate ID}-cn-hangzhou format.
	//
	// 	- If the certificate is a private certificate, the value of this parameter must be the value of the Identifier field for the private certificate.
	//
	// example:
	//
	// ccaf0c629c2be1e2abb63bb76b
	CertIdentifier   *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
	// The data to sign. The value must be encoded in Base64.\\
	//
	// For example, if the hexadecimal data that you want to sign is [0x31, 0x32, 0x33, 0x34], set the parameter to the Base64-encoded value MTIzNA==. If you set MessageType to RAW, the size of the data must be less than 4 KB. If the size of the data is greater than 4 KB, you can set MessageType to DIGEST and set Message to the digest of the data. The digest is a hash value. You can compute the digest of the data on an on-premises machine. The certificate application repository uses the digest that you compute in your own certificate application system. The message digest algorithm that you use must match the specified signature algorithm. The following items describe the details:
	//
	// 	- If the signature algorithm is SHA256withRSA, SHA256withRSA/PSS, or SHA256withECDSA, the message digest algorithm must be SHA-256.
	//
	// 	- If the signature algorithm is SM3withSM2, the message digest algorithm must be SM3.
	//
	// This parameter is required.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The value type of the Message parameter. Valid values:
	//
	// 	- RAW: the raw data. This is the default value.
	//
	// 	- DIGEST: the message digest (hash value) of the raw data.
	//
	// This parameter is required.
	//
	// example:
	//
	// RAW
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The signature algorithm. Valid values:
	//
	// 	- SHA256withRSA
	//
	// 	- SHA256withRSA/PSS
	//
	// 	- SHA256withECDSA
	//
	// 	- SM3withSM2
	//
	// This parameter is required.
	//
	// example:
	//
	// SHA256withRSA
	SigningAlgorithm *string `json:"SigningAlgorithm,omitempty" xml:"SigningAlgorithm,omitempty"`
}

func (s SignRequest) String() string {
	return dara.Prettify(s)
}

func (s SignRequest) GoString() string {
	return s.String()
}

func (s *SignRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *SignRequest) GetCustomIdentifier() *string {
	return s.CustomIdentifier
}

func (s *SignRequest) GetMessage() *string {
	return s.Message
}

func (s *SignRequest) GetMessageType() *string {
	return s.MessageType
}

func (s *SignRequest) GetSigningAlgorithm() *string {
	return s.SigningAlgorithm
}

func (s *SignRequest) SetCertIdentifier(v string) *SignRequest {
	s.CertIdentifier = &v
	return s
}

func (s *SignRequest) SetCustomIdentifier(v string) *SignRequest {
	s.CustomIdentifier = &v
	return s
}

func (s *SignRequest) SetMessage(v string) *SignRequest {
	s.Message = &v
	return s
}

func (s *SignRequest) SetMessageType(v string) *SignRequest {
	s.MessageType = &v
	return s
}

func (s *SignRequest) SetSigningAlgorithm(v string) *SignRequest {
	s.SigningAlgorithm = &v
	return s
}

func (s *SignRequest) Validate() error {
	return dara.Validate(s)
}
