// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMiniEngineVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListMiniEngineVersionsRequest
	GetInstanceId() *string
	SetLang(v string) *ListMiniEngineVersionsRequest
	GetLang() *string
	SetProductCode(v string) *ListMiniEngineVersionsRequest
	GetProductCode() *string
	SetProductId(v int64) *ListMiniEngineVersionsRequest
	GetProductId() *int64
}

type ListMiniEngineVersionsRequest struct {
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

func (s ListMiniEngineVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMiniEngineVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListMiniEngineVersionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListMiniEngineVersionsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListMiniEngineVersionsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListMiniEngineVersionsRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListMiniEngineVersionsRequest) SetInstanceId(v string) *ListMiniEngineVersionsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListMiniEngineVersionsRequest) SetLang(v string) *ListMiniEngineVersionsRequest {
	s.Lang = &v
	return s
}

func (s *ListMiniEngineVersionsRequest) SetProductCode(v string) *ListMiniEngineVersionsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListMiniEngineVersionsRequest) SetProductId(v int64) *ListMiniEngineVersionsRequest {
	s.ProductId = &v
	return s
}

func (s *ListMiniEngineVersionsRequest) Validate() error {
	return dara.Validate(s)
}
