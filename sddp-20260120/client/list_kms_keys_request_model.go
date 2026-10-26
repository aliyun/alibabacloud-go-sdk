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
	// The instance identity of the database. You can obtain the value from Items[].InstanceId in the response of the ListDataMaskingInstances operation.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response. Valid values: zh: Chinese. en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The code of the cloud service. Column encryption supports RDS, POLARDB, and PolarDBX2. If ProductId is also specified, the two parameters must match.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the cloud service type. Valid values: 5: RDS. 8: POLARDB. 27: PolarDBX2. If both this parameter and ProductCode are specified, they must match.
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
