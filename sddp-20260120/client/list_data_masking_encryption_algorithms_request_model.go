// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataMaskingEncryptionAlgorithmsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListDataMaskingEncryptionAlgorithmsRequest
	GetInstanceId() *string
	SetLang(v string) *ListDataMaskingEncryptionAlgorithmsRequest
	GetLang() *string
	SetProductCode(v string) *ListDataMaskingEncryptionAlgorithmsRequest
	GetProductCode() *string
	SetProductId(v int64) *ListDataMaskingEncryptionAlgorithmsRequest
	GetProductId() *int64
}

type ListDataMaskingEncryptionAlgorithmsRequest struct {
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

func (s ListDataMaskingEncryptionAlgorithmsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataMaskingEncryptionAlgorithmsRequest) GoString() string {
	return s.String()
}

func (s *ListDataMaskingEncryptionAlgorithmsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDataMaskingEncryptionAlgorithmsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataMaskingEncryptionAlgorithmsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListDataMaskingEncryptionAlgorithmsRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListDataMaskingEncryptionAlgorithmsRequest) SetInstanceId(v string) *ListDataMaskingEncryptionAlgorithmsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDataMaskingEncryptionAlgorithmsRequest) SetLang(v string) *ListDataMaskingEncryptionAlgorithmsRequest {
	s.Lang = &v
	return s
}

func (s *ListDataMaskingEncryptionAlgorithmsRequest) SetProductCode(v string) *ListDataMaskingEncryptionAlgorithmsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListDataMaskingEncryptionAlgorithmsRequest) SetProductId(v int64) *ListDataMaskingEncryptionAlgorithmsRequest {
	s.ProductId = &v
	return s
}

func (s *ListDataMaskingEncryptionAlgorithmsRequest) Validate() error {
	return dara.Validate(s)
}
