// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataIngestionTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataIngestionTemplates(v []*ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) *ListDataIngestionTemplatesResponseBody
	GetDataIngestionTemplates() []*ListDataIngestionTemplatesResponseBodyDataIngestionTemplates
	SetPageNumber(v string) *ListDataIngestionTemplatesResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *ListDataIngestionTemplatesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *ListDataIngestionTemplatesResponseBody
	GetRequestId() *string
}

type ListDataIngestionTemplatesResponseBody struct {
	DataIngestionTemplates []*ListDataIngestionTemplatesResponseBodyDataIngestionTemplates `json:"DataIngestionTemplates,omitempty" xml:"DataIngestionTemplates,omitempty" type:"Repeated"`
	// example:
	//
	// 1。
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10。
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataIngestionTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataIngestionTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataIngestionTemplatesResponseBody) GetDataIngestionTemplates() []*ListDataIngestionTemplatesResponseBodyDataIngestionTemplates {
	return s.DataIngestionTemplates
}

func (s *ListDataIngestionTemplatesResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDataIngestionTemplatesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDataIngestionTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataIngestionTemplatesResponseBody) SetDataIngestionTemplates(v []*ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) *ListDataIngestionTemplatesResponseBody {
	s.DataIngestionTemplates = v
	return s
}

func (s *ListDataIngestionTemplatesResponseBody) SetPageNumber(v string) *ListDataIngestionTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDataIngestionTemplatesResponseBody) SetPageSize(v string) *ListDataIngestionTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataIngestionTemplatesResponseBody) SetRequestId(v string) *ListDataIngestionTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataIngestionTemplatesResponseBody) Validate() error {
	if s.DataIngestionTemplates != nil {
		for _, item := range s.DataIngestionTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataIngestionTemplatesResponseBodyDataIngestionTemplates struct {
	// example:
	//
	// 3。
	CapacityCount *string `json:"CapacityCount,omitempty" xml:"CapacityCount,omitempty"`
	// example:
	//
	// 1733269771123。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// realtime。
	DataIngestionMode *string `json:"DataIngestionMode,omitempty" xml:"DataIngestionMode,omitempty"`
	// example:
	//
	// enabled。
	DataIngestionStatus *string `json:"DataIngestionStatus,omitempty" xml:"DataIngestionStatus,omitempty"`
	// example:
	//
	// alibaba_cloud_sas_account_snapshot_log。
	DataIngestionTemplateId *string `json:"DataIngestionTemplateId,omitempty" xml:"DataIngestionTemplateId,omitempty"`
	// example:
	//
	// alibaba_cloud_sas_account_snapshot_log。
	DataIngestionTemplateName *string `json:"DataIngestionTemplateName,omitempty" xml:"DataIngestionTemplateName,omitempty"`
	// example:
	//
	// running。
	DataIngestionTemplateStatus *string `json:"DataIngestionTemplateStatus,omitempty" xml:"DataIngestionTemplateStatus,omitempty"`
	// example:
	//
	// alibaba_cloud_sas_account_snapshot_log_173326*******。
	DataSourceTemplateId *string `json:"DataSourceTemplateId,omitempty" xml:"DataSourceTemplateId,omitempty"`
	// example:
	//
	// alibaba_cloud_actiontrail_event_rule。
	NormalizationRuleId *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
	// example:
	//
	// normalization_rule_ke1RN。
	NormalizationRuleName *string `json:"NormalizationRuleName,omitempty" xml:"NormalizationRuleName,omitempty"`
	// example:
	//
	// 173326*******。
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) GoString() string {
	return s.String()
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) GetCapacityCount() *string {
	return s.CapacityCount
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) GetDataIngestionMode() *string {
	return s.DataIngestionMode
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) GetDataIngestionStatus() *string {
	return s.DataIngestionStatus
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) GetDataIngestionTemplateId() *string {
	return s.DataIngestionTemplateId
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) GetDataIngestionTemplateName() *string {
	return s.DataIngestionTemplateName
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) GetDataIngestionTemplateStatus() *string {
	return s.DataIngestionTemplateStatus
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) GetDataSourceTemplateId() *string {
	return s.DataSourceTemplateId
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) GetNormalizationRuleName() *string {
	return s.NormalizationRuleName
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) SetCapacityCount(v string) *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates {
	s.CapacityCount = &v
	return s
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) SetCreateTime(v int64) *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates {
	s.CreateTime = &v
	return s
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) SetDataIngestionMode(v string) *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates {
	s.DataIngestionMode = &v
	return s
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) SetDataIngestionStatus(v string) *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates {
	s.DataIngestionStatus = &v
	return s
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) SetDataIngestionTemplateId(v string) *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates {
	s.DataIngestionTemplateId = &v
	return s
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) SetDataIngestionTemplateName(v string) *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates {
	s.DataIngestionTemplateName = &v
	return s
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) SetDataIngestionTemplateStatus(v string) *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates {
	s.DataIngestionTemplateStatus = &v
	return s
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) SetDataSourceTemplateId(v string) *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates {
	s.DataSourceTemplateId = &v
	return s
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) SetNormalizationRuleId(v string) *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates {
	s.NormalizationRuleId = &v
	return s
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) SetNormalizationRuleName(v string) *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates {
	s.NormalizationRuleName = &v
	return s
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) SetUpdateTime(v int64) *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates {
	s.UpdateTime = &v
	return s
}

func (s *ListDataIngestionTemplatesResponseBodyDataIngestionTemplates) Validate() error {
	return dara.Validate(s)
}
