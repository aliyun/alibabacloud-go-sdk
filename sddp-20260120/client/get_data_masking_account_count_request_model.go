// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataMaskingAccountCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetDataMaskingAccountCountRequest
	GetLang() *string
	SetProductIds(v string) *GetDataMaskingAccountCountRequest
	GetProductIds() *string
}

type GetDataMaskingAccountCountRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 5
	ProductIds *string `json:"ProductIds,omitempty" xml:"ProductIds,omitempty"`
}

func (s GetDataMaskingAccountCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataMaskingAccountCountRequest) GoString() string {
	return s.String()
}

func (s *GetDataMaskingAccountCountRequest) GetLang() *string {
	return s.Lang
}

func (s *GetDataMaskingAccountCountRequest) GetProductIds() *string {
	return s.ProductIds
}

func (s *GetDataMaskingAccountCountRequest) SetLang(v string) *GetDataMaskingAccountCountRequest {
	s.Lang = &v
	return s
}

func (s *GetDataMaskingAccountCountRequest) SetProductIds(v string) *GetDataMaskingAccountCountRequest {
	s.ProductIds = &v
	return s
}

func (s *GetDataMaskingAccountCountRequest) Validate() error {
	return dara.Validate(s)
}
