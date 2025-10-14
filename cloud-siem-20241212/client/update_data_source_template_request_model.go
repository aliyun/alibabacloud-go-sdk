// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoScanNew(v string) *UpdateDataSourceTemplateRequest
	GetAutoScanNew() *string
	SetDataSourceRecognizeEnabled(v bool) *UpdateDataSourceTemplateRequest
	GetDataSourceRecognizeEnabled() *bool
	SetDataSourceTemplateId(v string) *UpdateDataSourceTemplateRequest
	GetDataSourceTemplateId() *string
	SetDataSourceTemplateName(v string) *UpdateDataSourceTemplateRequest
	GetDataSourceTemplateName() *string
	SetLang(v string) *UpdateDataSourceTemplateRequest
	GetLang() *string
	SetLogProjectPattern(v string) *UpdateDataSourceTemplateRequest
	GetLogProjectPattern() *string
	SetLogRegionIds(v string) *UpdateDataSourceTemplateRequest
	GetLogRegionIds() *string
	SetLogStorePattern(v string) *UpdateDataSourceTemplateRequest
	GetLogStorePattern() *string
	SetLogUserIds(v []*string) *UpdateDataSourceTemplateRequest
	GetLogUserIds() []*string
	SetRegionId(v string) *UpdateDataSourceTemplateRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateDataSourceTemplateRequest
	GetRoleFor() *int64
}

type UpdateDataSourceTemplateRequest struct {
	// example:
	//
	// enabled。
	AutoScanNew *string `json:"AutoScanNew,omitempty" xml:"AutoScanNew,omitempty"`
	// example:
	//
	// true。
	DataSourceRecognizeEnabled *bool `json:"DataSourceRecognizeEnabled,omitempty" xml:"DataSourceRecognizeEnabled,omitempty"`
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
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// aliyun-cloudsiem-data-173326*******。
	LogProjectPattern *string `json:"LogProjectPattern,omitempty" xml:"LogProjectPattern,omitempty"`
	// example:
	//
	// cn-hangzhou。
	LogRegionIds *string `json:"LogRegionIds,omitempty" xml:"LogRegionIds,omitempty"`
	// example:
	//
	// audit-activity。
	LogStorePattern *string   `json:"LogStorePattern,omitempty" xml:"LogStorePattern,omitempty"`
	LogUserIds      []*string `json:"LogUserIds,omitempty" xml:"LogUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s UpdateDataSourceTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceTemplateRequest) GetAutoScanNew() *string {
	return s.AutoScanNew
}

func (s *UpdateDataSourceTemplateRequest) GetDataSourceRecognizeEnabled() *bool {
	return s.DataSourceRecognizeEnabled
}

func (s *UpdateDataSourceTemplateRequest) GetDataSourceTemplateId() *string {
	return s.DataSourceTemplateId
}

func (s *UpdateDataSourceTemplateRequest) GetDataSourceTemplateName() *string {
	return s.DataSourceTemplateName
}

func (s *UpdateDataSourceTemplateRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataSourceTemplateRequest) GetLogProjectPattern() *string {
	return s.LogProjectPattern
}

func (s *UpdateDataSourceTemplateRequest) GetLogRegionIds() *string {
	return s.LogRegionIds
}

func (s *UpdateDataSourceTemplateRequest) GetLogStorePattern() *string {
	return s.LogStorePattern
}

func (s *UpdateDataSourceTemplateRequest) GetLogUserIds() []*string {
	return s.LogUserIds
}

func (s *UpdateDataSourceTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataSourceTemplateRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateDataSourceTemplateRequest) SetAutoScanNew(v string) *UpdateDataSourceTemplateRequest {
	s.AutoScanNew = &v
	return s
}

func (s *UpdateDataSourceTemplateRequest) SetDataSourceRecognizeEnabled(v bool) *UpdateDataSourceTemplateRequest {
	s.DataSourceRecognizeEnabled = &v
	return s
}

func (s *UpdateDataSourceTemplateRequest) SetDataSourceTemplateId(v string) *UpdateDataSourceTemplateRequest {
	s.DataSourceTemplateId = &v
	return s
}

func (s *UpdateDataSourceTemplateRequest) SetDataSourceTemplateName(v string) *UpdateDataSourceTemplateRequest {
	s.DataSourceTemplateName = &v
	return s
}

func (s *UpdateDataSourceTemplateRequest) SetLang(v string) *UpdateDataSourceTemplateRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataSourceTemplateRequest) SetLogProjectPattern(v string) *UpdateDataSourceTemplateRequest {
	s.LogProjectPattern = &v
	return s
}

func (s *UpdateDataSourceTemplateRequest) SetLogRegionIds(v string) *UpdateDataSourceTemplateRequest {
	s.LogRegionIds = &v
	return s
}

func (s *UpdateDataSourceTemplateRequest) SetLogStorePattern(v string) *UpdateDataSourceTemplateRequest {
	s.LogStorePattern = &v
	return s
}

func (s *UpdateDataSourceTemplateRequest) SetLogUserIds(v []*string) *UpdateDataSourceTemplateRequest {
	s.LogUserIds = v
	return s
}

func (s *UpdateDataSourceTemplateRequest) SetRegionId(v string) *UpdateDataSourceTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataSourceTemplateRequest) SetRoleFor(v int64) *UpdateDataSourceTemplateRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateDataSourceTemplateRequest) Validate() error {
	return dara.Validate(s)
}
