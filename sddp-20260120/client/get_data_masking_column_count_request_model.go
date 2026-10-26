// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataMaskingColumnCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetDataMaskingColumnCountRequest
	GetLang() *string
	SetProductIds(v string) *GetDataMaskingColumnCountRequest
	GetProductIds() *string
	SetTemplateId(v int64) *GetDataMaskingColumnCountRequest
	GetTemplateId() *int64
}

type GetDataMaskingColumnCountRequest struct {
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
	// The list of cloud product type identifiers. Separate multiple values with commas (,). For column encryption scenarios, the values correspond to 5 (ApsaraDB RDS), 8 (PolarDB), and 27 (PolarDB-X).
	//
	// example:
	//
	// 5
	ProductIds *string `json:"ProductIds,omitempty" xml:"ProductIds,omitempty"`
	// The data classification template ID. You can obtain the ID from the Items[].Id parameter in the ListTemplates response. The ID must be consistent with the selected sensitivity level and detection model.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetDataMaskingColumnCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataMaskingColumnCountRequest) GoString() string {
	return s.String()
}

func (s *GetDataMaskingColumnCountRequest) GetLang() *string {
	return s.Lang
}

func (s *GetDataMaskingColumnCountRequest) GetProductIds() *string {
	return s.ProductIds
}

func (s *GetDataMaskingColumnCountRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetDataMaskingColumnCountRequest) SetLang(v string) *GetDataMaskingColumnCountRequest {
	s.Lang = &v
	return s
}

func (s *GetDataMaskingColumnCountRequest) SetProductIds(v string) *GetDataMaskingColumnCountRequest {
	s.ProductIds = &v
	return s
}

func (s *GetDataMaskingColumnCountRequest) SetTemplateId(v int64) *GetDataMaskingColumnCountRequest {
	s.TemplateId = &v
	return s
}

func (s *GetDataMaskingColumnCountRequest) Validate() error {
	return dara.Validate(s)
}
