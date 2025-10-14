// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoScanNew(v string) *UpdateDataSourceTemplateShrinkRequest
	GetAutoScanNew() *string
	SetDataSourceRecognizeEnabled(v bool) *UpdateDataSourceTemplateShrinkRequest
	GetDataSourceRecognizeEnabled() *bool
	SetDataSourceTemplateId(v string) *UpdateDataSourceTemplateShrinkRequest
	GetDataSourceTemplateId() *string
	SetDataSourceTemplateName(v string) *UpdateDataSourceTemplateShrinkRequest
	GetDataSourceTemplateName() *string
	SetLang(v string) *UpdateDataSourceTemplateShrinkRequest
	GetLang() *string
	SetLogProjectPattern(v string) *UpdateDataSourceTemplateShrinkRequest
	GetLogProjectPattern() *string
	SetLogRegionIds(v string) *UpdateDataSourceTemplateShrinkRequest
	GetLogRegionIds() *string
	SetLogStorePattern(v string) *UpdateDataSourceTemplateShrinkRequest
	GetLogStorePattern() *string
	SetLogUserIdsShrink(v string) *UpdateDataSourceTemplateShrinkRequest
	GetLogUserIdsShrink() *string
	SetRegionId(v string) *UpdateDataSourceTemplateShrinkRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateDataSourceTemplateShrinkRequest
	GetRoleFor() *int64
}

type UpdateDataSourceTemplateShrinkRequest struct {
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
	LogStorePattern  *string `json:"LogStorePattern,omitempty" xml:"LogStorePattern,omitempty"`
	LogUserIdsShrink *string `json:"LogUserIds,omitempty" xml:"LogUserIds,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s UpdateDataSourceTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceTemplateShrinkRequest) GetAutoScanNew() *string {
	return s.AutoScanNew
}

func (s *UpdateDataSourceTemplateShrinkRequest) GetDataSourceRecognizeEnabled() *bool {
	return s.DataSourceRecognizeEnabled
}

func (s *UpdateDataSourceTemplateShrinkRequest) GetDataSourceTemplateId() *string {
	return s.DataSourceTemplateId
}

func (s *UpdateDataSourceTemplateShrinkRequest) GetDataSourceTemplateName() *string {
	return s.DataSourceTemplateName
}

func (s *UpdateDataSourceTemplateShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataSourceTemplateShrinkRequest) GetLogProjectPattern() *string {
	return s.LogProjectPattern
}

func (s *UpdateDataSourceTemplateShrinkRequest) GetLogRegionIds() *string {
	return s.LogRegionIds
}

func (s *UpdateDataSourceTemplateShrinkRequest) GetLogStorePattern() *string {
	return s.LogStorePattern
}

func (s *UpdateDataSourceTemplateShrinkRequest) GetLogUserIdsShrink() *string {
	return s.LogUserIdsShrink
}

func (s *UpdateDataSourceTemplateShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataSourceTemplateShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateDataSourceTemplateShrinkRequest) SetAutoScanNew(v string) *UpdateDataSourceTemplateShrinkRequest {
	s.AutoScanNew = &v
	return s
}

func (s *UpdateDataSourceTemplateShrinkRequest) SetDataSourceRecognizeEnabled(v bool) *UpdateDataSourceTemplateShrinkRequest {
	s.DataSourceRecognizeEnabled = &v
	return s
}

func (s *UpdateDataSourceTemplateShrinkRequest) SetDataSourceTemplateId(v string) *UpdateDataSourceTemplateShrinkRequest {
	s.DataSourceTemplateId = &v
	return s
}

func (s *UpdateDataSourceTemplateShrinkRequest) SetDataSourceTemplateName(v string) *UpdateDataSourceTemplateShrinkRequest {
	s.DataSourceTemplateName = &v
	return s
}

func (s *UpdateDataSourceTemplateShrinkRequest) SetLang(v string) *UpdateDataSourceTemplateShrinkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataSourceTemplateShrinkRequest) SetLogProjectPattern(v string) *UpdateDataSourceTemplateShrinkRequest {
	s.LogProjectPattern = &v
	return s
}

func (s *UpdateDataSourceTemplateShrinkRequest) SetLogRegionIds(v string) *UpdateDataSourceTemplateShrinkRequest {
	s.LogRegionIds = &v
	return s
}

func (s *UpdateDataSourceTemplateShrinkRequest) SetLogStorePattern(v string) *UpdateDataSourceTemplateShrinkRequest {
	s.LogStorePattern = &v
	return s
}

func (s *UpdateDataSourceTemplateShrinkRequest) SetLogUserIdsShrink(v string) *UpdateDataSourceTemplateShrinkRequest {
	s.LogUserIdsShrink = &v
	return s
}

func (s *UpdateDataSourceTemplateShrinkRequest) SetRegionId(v string) *UpdateDataSourceTemplateShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataSourceTemplateShrinkRequest) SetRoleFor(v int64) *UpdateDataSourceTemplateShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateDataSourceTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
