// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataMaskingEncryptionAlgorithmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptionAlgorithm(v string) *UpdateDataMaskingEncryptionAlgorithmRequest
	GetEncryptionAlgorithm() *string
	SetEncryptionKeyId(v string) *UpdateDataMaskingEncryptionAlgorithmRequest
	GetEncryptionKeyId() *string
	SetInstanceId(v string) *UpdateDataMaskingEncryptionAlgorithmRequest
	GetInstanceId() *string
	SetLang(v string) *UpdateDataMaskingEncryptionAlgorithmRequest
	GetLang() *string
	SetProductCode(v string) *UpdateDataMaskingEncryptionAlgorithmRequest
	GetProductCode() *string
	SetProductId(v int64) *UpdateDataMaskingEncryptionAlgorithmRequest
	GetProductId() *int64
}

type UpdateDataMaskingEncryptionAlgorithmRequest struct {
	// example:
	//
	// AES_256_GCM
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty" xml:"EncryptionAlgorithm,omitempty"`
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s UpdateDataMaskingEncryptionAlgorithmRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataMaskingEncryptionAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataMaskingEncryptionAlgorithmRequest) GetEncryptionAlgorithm() *string {
	return s.EncryptionAlgorithm
}

func (s *UpdateDataMaskingEncryptionAlgorithmRequest) GetEncryptionKeyId() *string {
	return s.EncryptionKeyId
}

func (s *UpdateDataMaskingEncryptionAlgorithmRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateDataMaskingEncryptionAlgorithmRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataMaskingEncryptionAlgorithmRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *UpdateDataMaskingEncryptionAlgorithmRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *UpdateDataMaskingEncryptionAlgorithmRequest) SetEncryptionAlgorithm(v string) *UpdateDataMaskingEncryptionAlgorithmRequest {
	s.EncryptionAlgorithm = &v
	return s
}

func (s *UpdateDataMaskingEncryptionAlgorithmRequest) SetEncryptionKeyId(v string) *UpdateDataMaskingEncryptionAlgorithmRequest {
	s.EncryptionKeyId = &v
	return s
}

func (s *UpdateDataMaskingEncryptionAlgorithmRequest) SetInstanceId(v string) *UpdateDataMaskingEncryptionAlgorithmRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDataMaskingEncryptionAlgorithmRequest) SetLang(v string) *UpdateDataMaskingEncryptionAlgorithmRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataMaskingEncryptionAlgorithmRequest) SetProductCode(v string) *UpdateDataMaskingEncryptionAlgorithmRequest {
	s.ProductCode = &v
	return s
}

func (s *UpdateDataMaskingEncryptionAlgorithmRequest) SetProductId(v int64) *UpdateDataMaskingEncryptionAlgorithmRequest {
	s.ProductId = &v
	return s
}

func (s *UpdateDataMaskingEncryptionAlgorithmRequest) Validate() error {
	return dara.Validate(s)
}
