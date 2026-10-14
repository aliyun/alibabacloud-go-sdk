// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKmsKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListKmsKeysRequest
	GetInstanceId() *string
	SetLang(v string) *ListKmsKeysRequest
	GetLang() *string
	SetProductCode(v string) *ListKmsKeysRequest
	GetProductCode() *string
	SetProductId(v int64) *ListKmsKeysRequest
	GetProductId() *int64
}

type ListKmsKeysRequest struct {
	// The database instance identity. You can obtain this value from Items[].InstanceId in the response of the ListDataMaskingInstances operation.
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
	// The cloud product code. Column encryption supports RDS, POLARDB, and PolarDBX2. If you also specify ProductId, the two values must correspond to each other.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The cloud product type ID. Valid values:
	//
	// - 5: RDS.
	//
	// - 8: POLARDB.
	//
	// - 27: PolarDBX2.
	//
	// If you specify both ProductId and ProductCode, the values must correspond to each other.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s ListKmsKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKmsKeysRequest) GoString() string {
	return s.String()
}

func (s *ListKmsKeysRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListKmsKeysRequest) GetLang() *string {
	return s.Lang
}

func (s *ListKmsKeysRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListKmsKeysRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListKmsKeysRequest) SetInstanceId(v string) *ListKmsKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *ListKmsKeysRequest) SetLang(v string) *ListKmsKeysRequest {
	s.Lang = &v
	return s
}

func (s *ListKmsKeysRequest) SetProductCode(v string) *ListKmsKeysRequest {
	s.ProductCode = &v
	return s
}

func (s *ListKmsKeysRequest) SetProductId(v int64) *ListKmsKeysRequest {
	s.ProductId = &v
	return s
}

func (s *ListKmsKeysRequest) Validate() error {
	return dara.Validate(s)
}
