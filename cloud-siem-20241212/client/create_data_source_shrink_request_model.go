// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceFrom(v string) *CreateDataSourceShrinkRequest
	GetDataSourceFrom() *string
	SetDataSourceIdsShrink(v string) *CreateDataSourceShrinkRequest
	GetDataSourceIdsShrink() *string
	SetDataSourceName(v string) *CreateDataSourceShrinkRequest
	GetDataSourceName() *string
	SetDataSourceRecognizeEnabled(v bool) *CreateDataSourceShrinkRequest
	GetDataSourceRecognizeEnabled() *bool
	SetDataSourceRecognizer(v string) *CreateDataSourceShrinkRequest
	GetDataSourceRecognizer() *string
	SetDataSourceReferencesShrink(v string) *CreateDataSourceShrinkRequest
	GetDataSourceReferencesShrink() *string
	SetDataSourceStores(v []*CreateDataSourceShrinkRequestDataSourceStores) *CreateDataSourceShrinkRequest
	GetDataSourceStores() []*CreateDataSourceShrinkRequestDataSourceStores
	SetDataSourceTemplateId(v string) *CreateDataSourceShrinkRequest
	GetDataSourceTemplateId() *string
	SetDataSourceType(v string) *CreateDataSourceShrinkRequest
	GetDataSourceType() *string
	SetLang(v string) *CreateDataSourceShrinkRequest
	GetLang() *string
	SetLogProjectName(v string) *CreateDataSourceShrinkRequest
	GetLogProjectName() *string
	SetLogRegionId(v string) *CreateDataSourceShrinkRequest
	GetLogRegionId() *string
	SetLogStoreName(v string) *CreateDataSourceShrinkRequest
	GetLogStoreName() *string
	SetLogUserId(v int64) *CreateDataSourceShrinkRequest
	GetLogUserId() *int64
	SetOrder(v string) *CreateDataSourceShrinkRequest
	GetOrder() *string
	SetRegionId(v string) *CreateDataSourceShrinkRequest
	GetRegionId() *string
	SetRoleFor(v int64) *CreateDataSourceShrinkRequest
	GetRoleFor() *int64
}

type CreateDataSourceShrinkRequest struct {
	// The source of the data. Valid values:
	//
	// - center
	//
	// - custom
	//
	// example:
	//
	// center
	DataSourceFrom *string `json:"DataSourceFrom,omitempty" xml:"DataSourceFrom,omitempty"`
	// A list of data source IDs.
	DataSourceIdsShrink *string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// AD_LOG
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// Specifies whether to automatically discover new data sources.
	//
	// example:
	//
	// true
	DataSourceRecognizeEnabled *bool `json:"DataSourceRecognizeEnabled,omitempty" xml:"DataSourceRecognizeEnabled,omitempty"`
	// The data source recognizer.
	//
	// example:
	//
	// alibaba_cloud_waf_flow_log_1766185894104675
	DataSourceRecognizer *string `json:"DataSourceRecognizer,omitempty" xml:"DataSourceRecognizer,omitempty"`
	// The IDs of associated data access instances.
	DataSourceReferencesShrink *string `json:"DataSourceReferences,omitempty" xml:"DataSourceReferences,omitempty"`
	// The list of Simple Log Service projects.
	DataSourceStores []*CreateDataSourceShrinkRequestDataSourceStores `json:"DataSourceStores,omitempty" xml:"DataSourceStores,omitempty" type:"Repeated"`
	// The ID of the data source template.
	//
	// example:
	//
	// dst_alibaba_cloud_nas_audit_log_1358117679873357
	DataSourceTemplateId *string `json:"DataSourceTemplateId,omitempty" xml:"DataSourceTemplateId,omitempty"`
	// The type of the data source. Valid values:
	//
	// - preset
	//
	// - custom
	//
	// example:
	//
	// preset
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// aliyun-cloudsiem-data-173326*******-cn-hangzhou
	LogProjectName *string `json:"LogProjectName,omitempty" xml:"LogProjectName,omitempty"`
	// The ID of the log storage region.
	//
	// example:
	//
	// cn-hangzhou
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// The name of the Simple Log Service Logstore.
	//
	// example:
	//
	// mde_raw
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The user ID for data ingestion.
	//
	// example:
	//
	// 173326*******
	LogUserId *int64 `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	// The sort order. Valid values:
	//
	// - desc: descending.
	//
	// - asc: ascending.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The region of the Management Hub. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Assets are in the Chinese mainland.
	//
	// - ap-southeast-1: Assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member whose perspective the administrator assumes.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s CreateDataSourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceShrinkRequest) GetDataSourceFrom() *string {
	return s.DataSourceFrom
}

func (s *CreateDataSourceShrinkRequest) GetDataSourceIdsShrink() *string {
	return s.DataSourceIdsShrink
}

func (s *CreateDataSourceShrinkRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *CreateDataSourceShrinkRequest) GetDataSourceRecognizeEnabled() *bool {
	return s.DataSourceRecognizeEnabled
}

func (s *CreateDataSourceShrinkRequest) GetDataSourceRecognizer() *string {
	return s.DataSourceRecognizer
}

func (s *CreateDataSourceShrinkRequest) GetDataSourceReferencesShrink() *string {
	return s.DataSourceReferencesShrink
}

func (s *CreateDataSourceShrinkRequest) GetDataSourceStores() []*CreateDataSourceShrinkRequestDataSourceStores {
	return s.DataSourceStores
}

func (s *CreateDataSourceShrinkRequest) GetDataSourceTemplateId() *string {
	return s.DataSourceTemplateId
}

func (s *CreateDataSourceShrinkRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *CreateDataSourceShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateDataSourceShrinkRequest) GetLogProjectName() *string {
	return s.LogProjectName
}

func (s *CreateDataSourceShrinkRequest) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *CreateDataSourceShrinkRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *CreateDataSourceShrinkRequest) GetLogUserId() *int64 {
	return s.LogUserId
}

func (s *CreateDataSourceShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *CreateDataSourceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDataSourceShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CreateDataSourceShrinkRequest) SetDataSourceFrom(v string) *CreateDataSourceShrinkRequest {
	s.DataSourceFrom = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetDataSourceIdsShrink(v string) *CreateDataSourceShrinkRequest {
	s.DataSourceIdsShrink = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetDataSourceName(v string) *CreateDataSourceShrinkRequest {
	s.DataSourceName = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetDataSourceRecognizeEnabled(v bool) *CreateDataSourceShrinkRequest {
	s.DataSourceRecognizeEnabled = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetDataSourceRecognizer(v string) *CreateDataSourceShrinkRequest {
	s.DataSourceRecognizer = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetDataSourceReferencesShrink(v string) *CreateDataSourceShrinkRequest {
	s.DataSourceReferencesShrink = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetDataSourceStores(v []*CreateDataSourceShrinkRequestDataSourceStores) *CreateDataSourceShrinkRequest {
	s.DataSourceStores = v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetDataSourceTemplateId(v string) *CreateDataSourceShrinkRequest {
	s.DataSourceTemplateId = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetDataSourceType(v string) *CreateDataSourceShrinkRequest {
	s.DataSourceType = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetLang(v string) *CreateDataSourceShrinkRequest {
	s.Lang = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetLogProjectName(v string) *CreateDataSourceShrinkRequest {
	s.LogProjectName = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetLogRegionId(v string) *CreateDataSourceShrinkRequest {
	s.LogRegionId = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetLogStoreName(v string) *CreateDataSourceShrinkRequest {
	s.LogStoreName = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetLogUserId(v int64) *CreateDataSourceShrinkRequest {
	s.LogUserId = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetOrder(v string) *CreateDataSourceShrinkRequest {
	s.Order = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetRegionId(v string) *CreateDataSourceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetRoleFor(v int64) *CreateDataSourceShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) Validate() error {
	if s.DataSourceStores != nil {
		for _, item := range s.DataSourceStores {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDataSourceShrinkRequestDataSourceStores struct {
	// The source of the data. Valid values:
	//
	// - center
	//
	// - custom
	//
	// example:
	//
	// center
	DataSourceStoreFrom *string `json:"DataSourceStoreFrom,omitempty" xml:"DataSourceStoreFrom,omitempty"`
	// The ID of the log storage.
	//
	// example:
	//
	// 1
	DataSourceStoreId *string `json:"DataSourceStoreId,omitempty" xml:"DataSourceStoreId,omitempty"`
	// The status of the log storage. Valid values:
	//
	// - normal
	//
	// - abnormal
	//
	// example:
	//
	// normal
	DataSourceStoreStatus *string `json:"DataSourceStoreStatus,omitempty" xml:"DataSourceStoreStatus,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// aliyun-cloudsiem-data-173326*******-cn-hangzhou
	LogProjectName *string `json:"LogProjectName,omitempty" xml:"LogProjectName,omitempty"`
	// The ID of the log storage region.
	//
	// example:
	//
	// cn-hangzhou
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// The name of the Simple Log Service Logstore.
	//
	// example:
	//
	// actiontrail_management-events
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
}

func (s CreateDataSourceShrinkRequestDataSourceStores) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceShrinkRequestDataSourceStores) GoString() string {
	return s.String()
}

func (s *CreateDataSourceShrinkRequestDataSourceStores) GetDataSourceStoreFrom() *string {
	return s.DataSourceStoreFrom
}

func (s *CreateDataSourceShrinkRequestDataSourceStores) GetDataSourceStoreId() *string {
	return s.DataSourceStoreId
}

func (s *CreateDataSourceShrinkRequestDataSourceStores) GetDataSourceStoreStatus() *string {
	return s.DataSourceStoreStatus
}

func (s *CreateDataSourceShrinkRequestDataSourceStores) GetLogProjectName() *string {
	return s.LogProjectName
}

func (s *CreateDataSourceShrinkRequestDataSourceStores) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *CreateDataSourceShrinkRequestDataSourceStores) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *CreateDataSourceShrinkRequestDataSourceStores) SetDataSourceStoreFrom(v string) *CreateDataSourceShrinkRequestDataSourceStores {
	s.DataSourceStoreFrom = &v
	return s
}

func (s *CreateDataSourceShrinkRequestDataSourceStores) SetDataSourceStoreId(v string) *CreateDataSourceShrinkRequestDataSourceStores {
	s.DataSourceStoreId = &v
	return s
}

func (s *CreateDataSourceShrinkRequestDataSourceStores) SetDataSourceStoreStatus(v string) *CreateDataSourceShrinkRequestDataSourceStores {
	s.DataSourceStoreStatus = &v
	return s
}

func (s *CreateDataSourceShrinkRequestDataSourceStores) SetLogProjectName(v string) *CreateDataSourceShrinkRequestDataSourceStores {
	s.LogProjectName = &v
	return s
}

func (s *CreateDataSourceShrinkRequestDataSourceStores) SetLogRegionId(v string) *CreateDataSourceShrinkRequestDataSourceStores {
	s.LogRegionId = &v
	return s
}

func (s *CreateDataSourceShrinkRequestDataSourceStores) SetLogStoreName(v string) *CreateDataSourceShrinkRequestDataSourceStores {
	s.LogStoreName = &v
	return s
}

func (s *CreateDataSourceShrinkRequestDataSourceStores) Validate() error {
	return dara.Validate(s)
}
