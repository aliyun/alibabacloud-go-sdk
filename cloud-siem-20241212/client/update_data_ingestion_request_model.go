// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataIngestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataIngestionId(v string) *UpdateDataIngestionRequest
	GetDataIngestionId() *string
	SetDataIngestionMode(v string) *UpdateDataIngestionRequest
	GetDataIngestionMode() *string
	SetDataSourceId(v string) *UpdateDataIngestionRequest
	GetDataSourceId() *string
	SetLang(v string) *UpdateDataIngestionRequest
	GetLang() *string
	SetNormalizationRuleId(v string) *UpdateDataIngestionRequest
	GetNormalizationRuleId() *string
	SetRegionId(v string) *UpdateDataIngestionRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateDataIngestionRequest
	GetRoleFor() *int64
}

type UpdateDataIngestionRequest struct {
	// example:
	//
	// alibaba_cloud_actiontrail_event_ingestion_173326*******。
	DataIngestionId *string `json:"DataIngestionId,omitempty" xml:"DataIngestionId,omitempty"`
	// example:
	//
	// realtime。
	DataIngestionMode *string `json:"DataIngestionMode,omitempty" xml:"DataIngestionMode,omitempty"`
	// example:
	//
	// alibaba_cloud_actiontrail_event_log_173326*******。
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// alibaba_cloud_actiontrail_event_rule。
	NormalizationRuleId *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s UpdateDataIngestionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataIngestionRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataIngestionRequest) GetDataIngestionId() *string {
	return s.DataIngestionId
}

func (s *UpdateDataIngestionRequest) GetDataIngestionMode() *string {
	return s.DataIngestionMode
}

func (s *UpdateDataIngestionRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *UpdateDataIngestionRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataIngestionRequest) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *UpdateDataIngestionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataIngestionRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateDataIngestionRequest) SetDataIngestionId(v string) *UpdateDataIngestionRequest {
	s.DataIngestionId = &v
	return s
}

func (s *UpdateDataIngestionRequest) SetDataIngestionMode(v string) *UpdateDataIngestionRequest {
	s.DataIngestionMode = &v
	return s
}

func (s *UpdateDataIngestionRequest) SetDataSourceId(v string) *UpdateDataIngestionRequest {
	s.DataSourceId = &v
	return s
}

func (s *UpdateDataIngestionRequest) SetLang(v string) *UpdateDataIngestionRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataIngestionRequest) SetNormalizationRuleId(v string) *UpdateDataIngestionRequest {
	s.NormalizationRuleId = &v
	return s
}

func (s *UpdateDataIngestionRequest) SetRegionId(v string) *UpdateDataIngestionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataIngestionRequest) SetRoleFor(v int64) *UpdateDataIngestionRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateDataIngestionRequest) Validate() error {
	return dara.Validate(s)
}
