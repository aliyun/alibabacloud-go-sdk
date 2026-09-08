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
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductId   *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
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
