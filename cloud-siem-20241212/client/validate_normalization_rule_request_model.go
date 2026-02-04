// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateNormalizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ValidateNormalizationRuleRequest
	GetData() *string
	SetExtendFieldStoreMode(v string) *ValidateNormalizationRuleRequest
	GetExtendFieldStoreMode() *string
	SetLang(v string) *ValidateNormalizationRuleRequest
	GetLang() *string
	SetLogSample(v string) *ValidateNormalizationRuleRequest
	GetLogSample() *string
	SetNormalizationCategoryId(v string) *ValidateNormalizationRuleRequest
	GetNormalizationCategoryId() *string
	SetNormalizationRuleExpression(v string) *ValidateNormalizationRuleRequest
	GetNormalizationRuleExpression() *string
	SetNormalizationRuleMode(v string) *ValidateNormalizationRuleRequest
	GetNormalizationRuleMode() *string
	SetNormalizationSchemaId(v string) *ValidateNormalizationRuleRequest
	GetNormalizationSchemaId() *string
	SetProductId(v string) *ValidateNormalizationRuleRequest
	GetProductId() *string
	SetRegionId(v string) *ValidateNormalizationRuleRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ValidateNormalizationRuleRequest
	GetRoleFor() *int64
	SetVendorId(v string) *ValidateNormalizationRuleRequest
	GetVendorId() *string
}

type ValidateNormalizationRuleRequest struct {
	// example:
	//
	// 123456。
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// flat
	ExtendFieldStoreMode *string `json:"ExtendFieldStoreMode,omitempty" xml:"ExtendFieldStoreMode,omitempty"`
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// {"aaa":"bbb","xxx":"yyy"}
	LogSample *string `json:"LogSample,omitempty" xml:"LogSample,omitempty"`
	// example:
	//
	// NETWORK_CATEGORY。
	NormalizationCategoryId *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	// example:
	//
	// *
	NormalizationRuleExpression *string `json:"NormalizationRuleExpression,omitempty" xml:"NormalizationRuleExpression,omitempty"`
	// example:
	//
	// realtime
	NormalizationRuleMode *string `json:"NormalizationRuleMode,omitempty" xml:"NormalizationRuleMode,omitempty"`
	// example:
	//
	// HTTP_ACTIVITY。
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// example:
	//
	// alibaba_cloud_sas
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// alibaba_cloud
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s ValidateNormalizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateNormalizationRuleRequest) GoString() string {
	return s.String()
}

func (s *ValidateNormalizationRuleRequest) GetData() *string {
	return s.Data
}

func (s *ValidateNormalizationRuleRequest) GetExtendFieldStoreMode() *string {
	return s.ExtendFieldStoreMode
}

func (s *ValidateNormalizationRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *ValidateNormalizationRuleRequest) GetLogSample() *string {
	return s.LogSample
}

func (s *ValidateNormalizationRuleRequest) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *ValidateNormalizationRuleRequest) GetNormalizationRuleExpression() *string {
	return s.NormalizationRuleExpression
}

func (s *ValidateNormalizationRuleRequest) GetNormalizationRuleMode() *string {
	return s.NormalizationRuleMode
}

func (s *ValidateNormalizationRuleRequest) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *ValidateNormalizationRuleRequest) GetProductId() *string {
	return s.ProductId
}

func (s *ValidateNormalizationRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ValidateNormalizationRuleRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ValidateNormalizationRuleRequest) GetVendorId() *string {
	return s.VendorId
}

func (s *ValidateNormalizationRuleRequest) SetData(v string) *ValidateNormalizationRuleRequest {
	s.Data = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) SetExtendFieldStoreMode(v string) *ValidateNormalizationRuleRequest {
	s.ExtendFieldStoreMode = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) SetLang(v string) *ValidateNormalizationRuleRequest {
	s.Lang = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) SetLogSample(v string) *ValidateNormalizationRuleRequest {
	s.LogSample = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) SetNormalizationCategoryId(v string) *ValidateNormalizationRuleRequest {
	s.NormalizationCategoryId = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) SetNormalizationRuleExpression(v string) *ValidateNormalizationRuleRequest {
	s.NormalizationRuleExpression = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) SetNormalizationRuleMode(v string) *ValidateNormalizationRuleRequest {
	s.NormalizationRuleMode = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) SetNormalizationSchemaId(v string) *ValidateNormalizationRuleRequest {
	s.NormalizationSchemaId = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) SetProductId(v string) *ValidateNormalizationRuleRequest {
	s.ProductId = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) SetRegionId(v string) *ValidateNormalizationRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) SetRoleFor(v int64) *ValidateNormalizationRuleRequest {
	s.RoleFor = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) SetVendorId(v string) *ValidateNormalizationRuleRequest {
	s.VendorId = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
