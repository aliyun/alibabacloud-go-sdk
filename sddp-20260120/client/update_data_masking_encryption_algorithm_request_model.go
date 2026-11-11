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
	// The name of the instance-level column encryption algorithm. Select an available algorithm from ListDataMaskingEncryptionAlgorithms.
	//
	// example:
	//
	// AES_256_GCM
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty" xml:"EncryptionAlgorithm,omitempty"`
	// The identifier of the Key Management Service (KMS) key. You can obtain it from Items[].KeyId in the response of ListKmsKeys. Provide this parameter when using a KMS key. Do not pass this parameter in local key scenarios.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// The database instance identity. You can obtain it from Items[].InstanceId in the response of ListDataMaskingInstances.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response description. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The code of the cloud service. Column encryption supports ApsaraDB RDS, PolarDB, and PolarDB-X 2.0. If ProductId is also provided, the two parameters must match.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The identifier of the cloud service type. Valid values:
	//
	// - 5: ApsaraDB RDS.
	//
	// - 8: PolarDB.
	//
	// - 27: PolarDB-X 2.0.
	//
	// If specified together with ProductCode, the values must match.
	//
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
