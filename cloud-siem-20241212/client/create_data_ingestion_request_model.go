// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataIngestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapacityCount(v int32) *CreateDataIngestionRequest
	GetCapacityCount() *int32
	SetDataIngestionMode(v string) *CreateDataIngestionRequest
	GetDataIngestionMode() *string
	SetDataIngestionStateCode(v string) *CreateDataIngestionRequest
	GetDataIngestionStateCode() *string
	SetDataIngestionType(v string) *CreateDataIngestionRequest
	GetDataIngestionType() *string
	SetDataSourceEditable(v bool) *CreateDataIngestionRequest
	GetDataSourceEditable() *bool
	SetDataSourceId(v string) *CreateDataIngestionRequest
	GetDataSourceId() *string
	SetLang(v string) *CreateDataIngestionRequest
	GetLang() *string
	SetNormalizationRuleEditable(v bool) *CreateDataIngestionRequest
	GetNormalizationRuleEditable() *bool
	SetNormalizationRuleId(v string) *CreateDataIngestionRequest
	GetNormalizationRuleId() *string
	SetProductId(v string) *CreateDataIngestionRequest
	GetProductId() *string
	SetRegionId(v string) *CreateDataIngestionRequest
	GetRegionId() *string
	SetRoleFor(v int64) *CreateDataIngestionRequest
	GetRoleFor() *int64
	SetScanDataSourceId(v string) *CreateDataIngestionRequest
	GetScanDataSourceId() *string
	SetStreamJobId(v string) *CreateDataIngestionRequest
	GetStreamJobId() *string
	SetUpdateTime(v int64) *CreateDataIngestionRequest
	GetUpdateTime() *int64
}

type CreateDataIngestionRequest struct {
	CapacityCount *int32 `json:"CapacityCount,omitempty" xml:"CapacityCount,omitempty"`
	// example:
	//
	// realtime。
	DataIngestionMode      *string `json:"DataIngestionMode,omitempty" xml:"DataIngestionMode,omitempty"`
	DataIngestionStateCode *string `json:"DataIngestionStateCode,omitempty" xml:"DataIngestionStateCode,omitempty"`
	DataIngestionType      *string `json:"DataIngestionType,omitempty" xml:"DataIngestionType,omitempty"`
	DataSourceEditable     *bool   `json:"DataSourceEditable,omitempty" xml:"DataSourceEditable,omitempty"`
	// example:
	//
	// ds-3g6lyf4eonngyohaq7tr。
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// zh。
	Lang                      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NormalizationRuleEditable *bool   `json:"NormalizationRuleEditable,omitempty" xml:"NormalizationRuleEditable,omitempty"`
	// example:
	//
	// nr-hdmady54piigkjfv17yp。
	NormalizationRuleId *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
	ProductId           *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor          *int64  `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	ScanDataSourceId *string `json:"ScanDataSourceId,omitempty" xml:"ScanDataSourceId,omitempty"`
	StreamJobId      *string `json:"StreamJobId,omitempty" xml:"StreamJobId,omitempty"`
	UpdateTime       *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateDataIngestionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataIngestionRequest) GoString() string {
	return s.String()
}

func (s *CreateDataIngestionRequest) GetCapacityCount() *int32 {
	return s.CapacityCount
}

func (s *CreateDataIngestionRequest) GetDataIngestionMode() *string {
	return s.DataIngestionMode
}

func (s *CreateDataIngestionRequest) GetDataIngestionStateCode() *string {
	return s.DataIngestionStateCode
}

func (s *CreateDataIngestionRequest) GetDataIngestionType() *string {
	return s.DataIngestionType
}

func (s *CreateDataIngestionRequest) GetDataSourceEditable() *bool {
	return s.DataSourceEditable
}

func (s *CreateDataIngestionRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateDataIngestionRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateDataIngestionRequest) GetNormalizationRuleEditable() *bool {
	return s.NormalizationRuleEditable
}

func (s *CreateDataIngestionRequest) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *CreateDataIngestionRequest) GetProductId() *string {
	return s.ProductId
}

func (s *CreateDataIngestionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDataIngestionRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CreateDataIngestionRequest) GetScanDataSourceId() *string {
	return s.ScanDataSourceId
}

func (s *CreateDataIngestionRequest) GetStreamJobId() *string {
	return s.StreamJobId
}

func (s *CreateDataIngestionRequest) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *CreateDataIngestionRequest) SetCapacityCount(v int32) *CreateDataIngestionRequest {
	s.CapacityCount = &v
	return s
}

func (s *CreateDataIngestionRequest) SetDataIngestionMode(v string) *CreateDataIngestionRequest {
	s.DataIngestionMode = &v
	return s
}

func (s *CreateDataIngestionRequest) SetDataIngestionStateCode(v string) *CreateDataIngestionRequest {
	s.DataIngestionStateCode = &v
	return s
}

func (s *CreateDataIngestionRequest) SetDataIngestionType(v string) *CreateDataIngestionRequest {
	s.DataIngestionType = &v
	return s
}

func (s *CreateDataIngestionRequest) SetDataSourceEditable(v bool) *CreateDataIngestionRequest {
	s.DataSourceEditable = &v
	return s
}

func (s *CreateDataIngestionRequest) SetDataSourceId(v string) *CreateDataIngestionRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateDataIngestionRequest) SetLang(v string) *CreateDataIngestionRequest {
	s.Lang = &v
	return s
}

func (s *CreateDataIngestionRequest) SetNormalizationRuleEditable(v bool) *CreateDataIngestionRequest {
	s.NormalizationRuleEditable = &v
	return s
}

func (s *CreateDataIngestionRequest) SetNormalizationRuleId(v string) *CreateDataIngestionRequest {
	s.NormalizationRuleId = &v
	return s
}

func (s *CreateDataIngestionRequest) SetProductId(v string) *CreateDataIngestionRequest {
	s.ProductId = &v
	return s
}

func (s *CreateDataIngestionRequest) SetRegionId(v string) *CreateDataIngestionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDataIngestionRequest) SetRoleFor(v int64) *CreateDataIngestionRequest {
	s.RoleFor = &v
	return s
}

func (s *CreateDataIngestionRequest) SetScanDataSourceId(v string) *CreateDataIngestionRequest {
	s.ScanDataSourceId = &v
	return s
}

func (s *CreateDataIngestionRequest) SetStreamJobId(v string) *CreateDataIngestionRequest {
	s.StreamJobId = &v
	return s
}

func (s *CreateDataIngestionRequest) SetUpdateTime(v int64) *CreateDataIngestionRequest {
	s.UpdateTime = &v
	return s
}

func (s *CreateDataIngestionRequest) Validate() error {
	return dara.Validate(s)
}
