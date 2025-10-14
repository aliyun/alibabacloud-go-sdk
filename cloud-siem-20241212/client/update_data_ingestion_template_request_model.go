// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataIngestionTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataIngestionStatus(v string) *UpdateDataIngestionTemplateRequest
	GetDataIngestionStatus() *string
	SetDataIngestionTemplateId(v string) *UpdateDataIngestionTemplateRequest
	GetDataIngestionTemplateId() *string
	SetDataIngestionTemplateName(v string) *UpdateDataIngestionTemplateRequest
	GetDataIngestionTemplateName() *string
	SetLang(v string) *UpdateDataIngestionTemplateRequest
	GetLang() *string
	SetNormalizationRuleId(v string) *UpdateDataIngestionTemplateRequest
	GetNormalizationRuleId() *string
	SetRegionId(v string) *UpdateDataIngestionTemplateRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateDataIngestionTemplateRequest
	GetRoleFor() *int64
}

type UpdateDataIngestionTemplateRequest struct {
	// example:
	//
	// enabled。
	DataIngestionStatus *string `json:"DataIngestionStatus,omitempty" xml:"DataIngestionStatus,omitempty"`
	// example:
	//
	// alibaba_cloud_actiontrail_event_ingestion_173326*******。
	DataIngestionTemplateId *string `json:"DataIngestionTemplateId,omitempty" xml:"DataIngestionTemplateId,omitempty"`
	// example:
	//
	// alibaba_cloud_actiontrail_event_ingestion_173326*******。
	DataIngestionTemplateName *string `json:"DataIngestionTemplateName,omitempty" xml:"DataIngestionTemplateName,omitempty"`
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

func (s UpdateDataIngestionTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataIngestionTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataIngestionTemplateRequest) GetDataIngestionStatus() *string {
	return s.DataIngestionStatus
}

func (s *UpdateDataIngestionTemplateRequest) GetDataIngestionTemplateId() *string {
	return s.DataIngestionTemplateId
}

func (s *UpdateDataIngestionTemplateRequest) GetDataIngestionTemplateName() *string {
	return s.DataIngestionTemplateName
}

func (s *UpdateDataIngestionTemplateRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataIngestionTemplateRequest) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *UpdateDataIngestionTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataIngestionTemplateRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateDataIngestionTemplateRequest) SetDataIngestionStatus(v string) *UpdateDataIngestionTemplateRequest {
	s.DataIngestionStatus = &v
	return s
}

func (s *UpdateDataIngestionTemplateRequest) SetDataIngestionTemplateId(v string) *UpdateDataIngestionTemplateRequest {
	s.DataIngestionTemplateId = &v
	return s
}

func (s *UpdateDataIngestionTemplateRequest) SetDataIngestionTemplateName(v string) *UpdateDataIngestionTemplateRequest {
	s.DataIngestionTemplateName = &v
	return s
}

func (s *UpdateDataIngestionTemplateRequest) SetLang(v string) *UpdateDataIngestionTemplateRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataIngestionTemplateRequest) SetNormalizationRuleId(v string) *UpdateDataIngestionTemplateRequest {
	s.NormalizationRuleId = &v
	return s
}

func (s *UpdateDataIngestionTemplateRequest) SetRegionId(v string) *UpdateDataIngestionTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataIngestionTemplateRequest) SetRoleFor(v int64) *UpdateDataIngestionTemplateRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateDataIngestionTemplateRequest) Validate() error {
	return dara.Validate(s)
}
