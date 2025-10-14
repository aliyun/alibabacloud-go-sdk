// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceTemplates(v []*ListDataSourceTemplatesResponseBodyDataSourceTemplates) *ListDataSourceTemplatesResponseBody
	GetDataSourceTemplates() []*ListDataSourceTemplatesResponseBodyDataSourceTemplates
	SetPageNumber(v string) *ListDataSourceTemplatesResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *ListDataSourceTemplatesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *ListDataSourceTemplatesResponseBody
	GetRequestId() *string
}

type ListDataSourceTemplatesResponseBody struct {
	DataSourceTemplates []*ListDataSourceTemplatesResponseBodyDataSourceTemplates `json:"DataSourceTemplates,omitempty" xml:"DataSourceTemplates,omitempty" type:"Repeated"`
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

func (s ListDataSourceTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceTemplatesResponseBody) GetDataSourceTemplates() []*ListDataSourceTemplatesResponseBodyDataSourceTemplates {
	return s.DataSourceTemplates
}

func (s *ListDataSourceTemplatesResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDataSourceTemplatesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDataSourceTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourceTemplatesResponseBody) SetDataSourceTemplates(v []*ListDataSourceTemplatesResponseBodyDataSourceTemplates) *ListDataSourceTemplatesResponseBody {
	s.DataSourceTemplates = v
	return s
}

func (s *ListDataSourceTemplatesResponseBody) SetPageNumber(v string) *ListDataSourceTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourceTemplatesResponseBody) SetPageSize(v string) *ListDataSourceTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataSourceTemplatesResponseBody) SetRequestId(v string) *ListDataSourceTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceTemplatesResponseBody) Validate() error {
	if s.DataSourceTemplates != nil {
		for _, item := range s.DataSourceTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSourceTemplatesResponseBodyDataSourceTemplates struct {
	// example:
	//
	// enabled。
	AutoScanNew *string `json:"AutoScanNew,omitempty" xml:"AutoScanNew,omitempty"`
	// example:
	//
	// 1733269771123。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// custom。
	DataSourceFrom *string `json:"DataSourceFrom,omitempty" xml:"DataSourceFrom,omitempty"`
	// example:
	//
	// true。
	DataSourceRecognizeEnabled *bool `json:"DataSourceRecognizeEnabled,omitempty" xml:"DataSourceRecognizeEnabled,omitempty"`
	// example:
	//
	// alibaba_cloud_actiontrail_event_ingestion。
	DataSourceRecognizer *string `json:"DataSourceRecognizer,omitempty" xml:"DataSourceRecognizer,omitempty"`
	// example:
	//
	// alibaba_cloud_actiontrail_event_ingestion。
	DataSourceTemplateId *string `json:"DataSourceTemplateId,omitempty" xml:"DataSourceTemplateId,omitempty"`
	// example:
	//
	// alibaba_cloud_actiontrail_event_ingestion。
	DataSourceTemplateName *string `json:"DataSourceTemplateName,omitempty" xml:"DataSourceTemplateName,omitempty"`
	// example:
	//
	// aliyun-cloudsiem-data-173326*******。
	LogProjectPattern *string   `json:"LogProjectPattern,omitempty" xml:"LogProjectPattern,omitempty"`
	LogRegionIds      []*string `json:"LogRegionIds,omitempty" xml:"LogRegionIds,omitempty" type:"Repeated"`
	// example:
	//
	// audit-activity。
	LogStorePattern *string   `json:"LogStorePattern,omitempty" xml:"LogStorePattern,omitempty"`
	LogUserIds      []*string `json:"LogUserIds,omitempty" xml:"LogUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-01-08 10:28:31
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDataSourceTemplatesResponseBodyDataSourceTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTemplatesResponseBodyDataSourceTemplates) GoString() string {
	return s.String()
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) GetAutoScanNew() *string {
	return s.AutoScanNew
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) GetDataSourceFrom() *string {
	return s.DataSourceFrom
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) GetDataSourceRecognizeEnabled() *bool {
	return s.DataSourceRecognizeEnabled
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) GetDataSourceRecognizer() *string {
	return s.DataSourceRecognizer
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) GetDataSourceTemplateId() *string {
	return s.DataSourceTemplateId
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) GetDataSourceTemplateName() *string {
	return s.DataSourceTemplateName
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) GetLogProjectPattern() *string {
	return s.LogProjectPattern
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) GetLogRegionIds() []*string {
	return s.LogRegionIds
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) GetLogStorePattern() *string {
	return s.LogStorePattern
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) GetLogUserIds() []*string {
	return s.LogUserIds
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) SetAutoScanNew(v string) *ListDataSourceTemplatesResponseBodyDataSourceTemplates {
	s.AutoScanNew = &v
	return s
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) SetCreateTime(v int64) *ListDataSourceTemplatesResponseBodyDataSourceTemplates {
	s.CreateTime = &v
	return s
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) SetDataSourceFrom(v string) *ListDataSourceTemplatesResponseBodyDataSourceTemplates {
	s.DataSourceFrom = &v
	return s
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) SetDataSourceRecognizeEnabled(v bool) *ListDataSourceTemplatesResponseBodyDataSourceTemplates {
	s.DataSourceRecognizeEnabled = &v
	return s
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) SetDataSourceRecognizer(v string) *ListDataSourceTemplatesResponseBodyDataSourceTemplates {
	s.DataSourceRecognizer = &v
	return s
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) SetDataSourceTemplateId(v string) *ListDataSourceTemplatesResponseBodyDataSourceTemplates {
	s.DataSourceTemplateId = &v
	return s
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) SetDataSourceTemplateName(v string) *ListDataSourceTemplatesResponseBodyDataSourceTemplates {
	s.DataSourceTemplateName = &v
	return s
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) SetLogProjectPattern(v string) *ListDataSourceTemplatesResponseBodyDataSourceTemplates {
	s.LogProjectPattern = &v
	return s
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) SetLogRegionIds(v []*string) *ListDataSourceTemplatesResponseBodyDataSourceTemplates {
	s.LogRegionIds = v
	return s
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) SetLogStorePattern(v string) *ListDataSourceTemplatesResponseBodyDataSourceTemplates {
	s.LogStorePattern = &v
	return s
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) SetLogUserIds(v []*string) *ListDataSourceTemplatesResponseBodyDataSourceTemplates {
	s.LogUserIds = v
	return s
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) SetUpdateTime(v int64) *ListDataSourceTemplatesResponseBodyDataSourceTemplates {
	s.UpdateTime = &v
	return s
}

func (s *ListDataSourceTemplatesResponseBodyDataSourceTemplates) Validate() error {
	return dara.Validate(s)
}
