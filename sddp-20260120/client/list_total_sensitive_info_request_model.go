// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTotalSensitiveInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCountType(v int32) *ListTotalSensitiveInfoRequest
	GetCountType() *int32
	SetLang(v string) *ListTotalSensitiveInfoRequest
	GetLang() *string
	SetProductCode(v string) *ListTotalSensitiveInfoRequest
	GetProductCode() *string
	SetProductCodeList(v string) *ListTotalSensitiveInfoRequest
	GetProductCodeList() *string
	SetTemplateId(v int64) *ListTotalSensitiveInfoRequest
	GetTemplateId() *int64
}

type ListTotalSensitiveInfoRequest struct {
	// The statistics scenario. Valid values:
	//
	// - 41: overview.
	//
	// - 42: data catalog statistics by product.
	//
	// - 43: region and template summary by product.
	//
	// example:
	//
	// 43
	CountType *int32 `json:"CountType,omitempty" xml:"CountType,omitempty"`
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
	// The code of the cloud service. Column encryption supports RDS, POLARDB, and PolarDBX2. If you also specify ProductId, the two values must correspond to each other.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Multiple product codes separated by commas (,). This parameter is used when CountType is set to 43 and ProductCode is not specified. The results are merged by region and template.
	//
	// example:
	//
	// RDS,POLARDB,PolarDBX2
	ProductCodeList *string `json:"ProductCodeList,omitempty" xml:"ProductCodeList,omitempty"`
	// The ID of the data classification template. You can obtain this value from Items[].Id in the response of the ListTemplates operation. The value must be consistent with the selected sensitivity level and detection model.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListTotalSensitiveInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTotalSensitiveInfoRequest) GoString() string {
	return s.String()
}

func (s *ListTotalSensitiveInfoRequest) GetCountType() *int32 {
	return s.CountType
}

func (s *ListTotalSensitiveInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *ListTotalSensitiveInfoRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListTotalSensitiveInfoRequest) GetProductCodeList() *string {
	return s.ProductCodeList
}

func (s *ListTotalSensitiveInfoRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListTotalSensitiveInfoRequest) SetCountType(v int32) *ListTotalSensitiveInfoRequest {
	s.CountType = &v
	return s
}

func (s *ListTotalSensitiveInfoRequest) SetLang(v string) *ListTotalSensitiveInfoRequest {
	s.Lang = &v
	return s
}

func (s *ListTotalSensitiveInfoRequest) SetProductCode(v string) *ListTotalSensitiveInfoRequest {
	s.ProductCode = &v
	return s
}

func (s *ListTotalSensitiveInfoRequest) SetProductCodeList(v string) *ListTotalSensitiveInfoRequest {
	s.ProductCodeList = &v
	return s
}

func (s *ListTotalSensitiveInfoRequest) SetTemplateId(v int64) *ListTotalSensitiveInfoRequest {
	s.TemplateId = &v
	return s
}

func (s *ListTotalSensitiveInfoRequest) Validate() error {
	return dara.Validate(s)
}
