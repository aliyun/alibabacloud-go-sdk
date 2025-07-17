// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTeeEvidenceVO interface {
	dara.Model
	String() string
	GoString() string
	SetCipherSuite(v string) *TeeEvidenceVO
	GetCipherSuite() *string
	SetEnclaveData(v string) *TeeEvidenceVO
	GetEnclaveData() *string
	SetEncryptPublicKeyPem(v string) *TeeEvidenceVO
	GetEncryptPublicKeyPem() *string
	SetEncryptPublicKeyType(v string) *TeeEvidenceVO
	GetEncryptPublicKeyType() *string
	SetModifiedDate(v string) *TeeEvidenceVO
	GetModifiedDate() *string
	SetPublicKey(v string) *TeeEvidenceVO
	GetPublicKey() *string
	SetPublicKeyRaBase64(v string) *TeeEvidenceVO
	GetPublicKeyRaBase64() *string
	SetPublicKeyRaType(v string) *TeeEvidenceVO
	GetPublicKeyRaType() *string
	SetQuoteReport(v string) *TeeEvidenceVO
	GetQuoteReport() *string
	SetSignPublicKeyPem(v string) *TeeEvidenceVO
	GetSignPublicKeyPem() *string
	SetSignPublicKeyType(v string) *TeeEvidenceVO
	GetSignPublicKeyType() *string
	SetTrustedMrEnclave(v []*string) *TeeEvidenceVO
	GetTrustedMrEnclave() []*string
}

type TeeEvidenceVO struct {
	CipherSuite          *string   `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	EnclaveData          *string   `json:"EnclaveData,omitempty" xml:"EnclaveData,omitempty"`
	EncryptPublicKeyPem  *string   `json:"EncryptPublicKeyPem,omitempty" xml:"EncryptPublicKeyPem,omitempty"`
	EncryptPublicKeyType *string   `json:"EncryptPublicKeyType,omitempty" xml:"EncryptPublicKeyType,omitempty"`
	ModifiedDate         *string   `json:"ModifiedDate,omitempty" xml:"ModifiedDate,omitempty"`
	PublicKey            *string   `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	PublicKeyRaBase64    *string   `json:"PublicKeyRaBase64,omitempty" xml:"PublicKeyRaBase64,omitempty"`
	PublicKeyRaType      *string   `json:"PublicKeyRaType,omitempty" xml:"PublicKeyRaType,omitempty"`
	QuoteReport          *string   `json:"QuoteReport,omitempty" xml:"QuoteReport,omitempty"`
	SignPublicKeyPem     *string   `json:"SignPublicKeyPem,omitempty" xml:"SignPublicKeyPem,omitempty"`
	SignPublicKeyType    *string   `json:"SignPublicKeyType,omitempty" xml:"SignPublicKeyType,omitempty"`
	TrustedMrEnclave     []*string `json:"TrustedMrEnclave,omitempty" xml:"TrustedMrEnclave,omitempty" type:"Repeated"`
}

func (s TeeEvidenceVO) String() string {
	return dara.Prettify(s)
}

func (s TeeEvidenceVO) GoString() string {
	return s.String()
}

func (s *TeeEvidenceVO) GetCipherSuite() *string {
	return s.CipherSuite
}

func (s *TeeEvidenceVO) GetEnclaveData() *string {
	return s.EnclaveData
}

func (s *TeeEvidenceVO) GetEncryptPublicKeyPem() *string {
	return s.EncryptPublicKeyPem
}

func (s *TeeEvidenceVO) GetEncryptPublicKeyType() *string {
	return s.EncryptPublicKeyType
}

func (s *TeeEvidenceVO) GetModifiedDate() *string {
	return s.ModifiedDate
}

func (s *TeeEvidenceVO) GetPublicKey() *string {
	return s.PublicKey
}

func (s *TeeEvidenceVO) GetPublicKeyRaBase64() *string {
	return s.PublicKeyRaBase64
}

func (s *TeeEvidenceVO) GetPublicKeyRaType() *string {
	return s.PublicKeyRaType
}

func (s *TeeEvidenceVO) GetQuoteReport() *string {
	return s.QuoteReport
}

func (s *TeeEvidenceVO) GetSignPublicKeyPem() *string {
	return s.SignPublicKeyPem
}

func (s *TeeEvidenceVO) GetSignPublicKeyType() *string {
	return s.SignPublicKeyType
}

func (s *TeeEvidenceVO) GetTrustedMrEnclave() []*string {
	return s.TrustedMrEnclave
}

func (s *TeeEvidenceVO) SetCipherSuite(v string) *TeeEvidenceVO {
	s.CipherSuite = &v
	return s
}

func (s *TeeEvidenceVO) SetEnclaveData(v string) *TeeEvidenceVO {
	s.EnclaveData = &v
	return s
}

func (s *TeeEvidenceVO) SetEncryptPublicKeyPem(v string) *TeeEvidenceVO {
	s.EncryptPublicKeyPem = &v
	return s
}

func (s *TeeEvidenceVO) SetEncryptPublicKeyType(v string) *TeeEvidenceVO {
	s.EncryptPublicKeyType = &v
	return s
}

func (s *TeeEvidenceVO) SetModifiedDate(v string) *TeeEvidenceVO {
	s.ModifiedDate = &v
	return s
}

func (s *TeeEvidenceVO) SetPublicKey(v string) *TeeEvidenceVO {
	s.PublicKey = &v
	return s
}

func (s *TeeEvidenceVO) SetPublicKeyRaBase64(v string) *TeeEvidenceVO {
	s.PublicKeyRaBase64 = &v
	return s
}

func (s *TeeEvidenceVO) SetPublicKeyRaType(v string) *TeeEvidenceVO {
	s.PublicKeyRaType = &v
	return s
}

func (s *TeeEvidenceVO) SetQuoteReport(v string) *TeeEvidenceVO {
	s.QuoteReport = &v
	return s
}

func (s *TeeEvidenceVO) SetSignPublicKeyPem(v string) *TeeEvidenceVO {
	s.SignPublicKeyPem = &v
	return s
}

func (s *TeeEvidenceVO) SetSignPublicKeyType(v string) *TeeEvidenceVO {
	s.SignPublicKeyType = &v
	return s
}

func (s *TeeEvidenceVO) SetTrustedMrEnclave(v []*string) *TeeEvidenceVO {
	s.TrustedMrEnclave = v
	return s
}

func (s *TeeEvidenceVO) Validate() error {
	return dara.Validate(s)
}
