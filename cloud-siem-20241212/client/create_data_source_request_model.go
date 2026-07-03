// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceFrom(v string) *CreateDataSourceRequest
	GetDataSourceFrom() *string
	SetDataSourceIds(v []*string) *CreateDataSourceRequest
	GetDataSourceIds() []*string
	SetDataSourceName(v string) *CreateDataSourceRequest
	GetDataSourceName() *string
	SetDataSourceRecognizeEnabled(v bool) *CreateDataSourceRequest
	GetDataSourceRecognizeEnabled() *bool
	SetDataSourceRecognizer(v string) *CreateDataSourceRequest
	GetDataSourceRecognizer() *string
	SetDataSourceReferences(v []*string) *CreateDataSourceRequest
	GetDataSourceReferences() []*string
	SetDataSourceStores(v []*CreateDataSourceRequestDataSourceStores) *CreateDataSourceRequest
	GetDataSourceStores() []*CreateDataSourceRequestDataSourceStores
	SetDataSourceTemplateId(v string) *CreateDataSourceRequest
	GetDataSourceTemplateId() *string
	SetDataSourceType(v string) *CreateDataSourceRequest
	GetDataSourceType() *string
	SetLang(v string) *CreateDataSourceRequest
	GetLang() *string
	SetLogProjectName(v string) *CreateDataSourceRequest
	GetLogProjectName() *string
	SetLogRegionId(v string) *CreateDataSourceRequest
	GetLogRegionId() *string
	SetLogStoreName(v string) *CreateDataSourceRequest
	GetLogStoreName() *string
	SetLogUserId(v int64) *CreateDataSourceRequest
	GetLogUserId() *int64
	SetOrder(v string) *CreateDataSourceRequest
	GetOrder() *string
	SetRegionId(v string) *CreateDataSourceRequest
	GetRegionId() *string
	SetRoleFor(v int64) *CreateDataSourceRequest
	GetRoleFor() *int64
}

type CreateDataSourceRequest struct {
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
	DataSourceIds []*string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty" type:"Repeated"`
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
	DataSourceReferences []*string `json:"DataSourceReferences,omitempty" xml:"DataSourceReferences,omitempty" type:"Repeated"`
	// The list of Simple Log Service projects.
	DataSourceStores []*CreateDataSourceRequestDataSourceStores `json:"DataSourceStores,omitempty" xml:"DataSourceStores,omitempty" type:"Repeated"`
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

func (s CreateDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequest) GetDataSourceFrom() *string {
	return s.DataSourceFrom
}

func (s *CreateDataSourceRequest) GetDataSourceIds() []*string {
	return s.DataSourceIds
}

func (s *CreateDataSourceRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *CreateDataSourceRequest) GetDataSourceRecognizeEnabled() *bool {
	return s.DataSourceRecognizeEnabled
}

func (s *CreateDataSourceRequest) GetDataSourceRecognizer() *string {
	return s.DataSourceRecognizer
}

func (s *CreateDataSourceRequest) GetDataSourceReferences() []*string {
	return s.DataSourceReferences
}

func (s *CreateDataSourceRequest) GetDataSourceStores() []*CreateDataSourceRequestDataSourceStores {
	return s.DataSourceStores
}

func (s *CreateDataSourceRequest) GetDataSourceTemplateId() *string {
	return s.DataSourceTemplateId
}

func (s *CreateDataSourceRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *CreateDataSourceRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateDataSourceRequest) GetLogProjectName() *string {
	return s.LogProjectName
}

func (s *CreateDataSourceRequest) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *CreateDataSourceRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *CreateDataSourceRequest) GetLogUserId() *int64 {
	return s.LogUserId
}

func (s *CreateDataSourceRequest) GetOrder() *string {
	return s.Order
}

func (s *CreateDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDataSourceRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CreateDataSourceRequest) SetDataSourceFrom(v string) *CreateDataSourceRequest {
	s.DataSourceFrom = &v
	return s
}

func (s *CreateDataSourceRequest) SetDataSourceIds(v []*string) *CreateDataSourceRequest {
	s.DataSourceIds = v
	return s
}

func (s *CreateDataSourceRequest) SetDataSourceName(v string) *CreateDataSourceRequest {
	s.DataSourceName = &v
	return s
}

func (s *CreateDataSourceRequest) SetDataSourceRecognizeEnabled(v bool) *CreateDataSourceRequest {
	s.DataSourceRecognizeEnabled = &v
	return s
}

func (s *CreateDataSourceRequest) SetDataSourceRecognizer(v string) *CreateDataSourceRequest {
	s.DataSourceRecognizer = &v
	return s
}

func (s *CreateDataSourceRequest) SetDataSourceReferences(v []*string) *CreateDataSourceRequest {
	s.DataSourceReferences = v
	return s
}

func (s *CreateDataSourceRequest) SetDataSourceStores(v []*CreateDataSourceRequestDataSourceStores) *CreateDataSourceRequest {
	s.DataSourceStores = v
	return s
}

func (s *CreateDataSourceRequest) SetDataSourceTemplateId(v string) *CreateDataSourceRequest {
	s.DataSourceTemplateId = &v
	return s
}

func (s *CreateDataSourceRequest) SetDataSourceType(v string) *CreateDataSourceRequest {
	s.DataSourceType = &v
	return s
}

func (s *CreateDataSourceRequest) SetLang(v string) *CreateDataSourceRequest {
	s.Lang = &v
	return s
}

func (s *CreateDataSourceRequest) SetLogProjectName(v string) *CreateDataSourceRequest {
	s.LogProjectName = &v
	return s
}

func (s *CreateDataSourceRequest) SetLogRegionId(v string) *CreateDataSourceRequest {
	s.LogRegionId = &v
	return s
}

func (s *CreateDataSourceRequest) SetLogStoreName(v string) *CreateDataSourceRequest {
	s.LogStoreName = &v
	return s
}

func (s *CreateDataSourceRequest) SetLogUserId(v int64) *CreateDataSourceRequest {
	s.LogUserId = &v
	return s
}

func (s *CreateDataSourceRequest) SetOrder(v string) *CreateDataSourceRequest {
	s.Order = &v
	return s
}

func (s *CreateDataSourceRequest) SetRegionId(v string) *CreateDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDataSourceRequest) SetRoleFor(v int64) *CreateDataSourceRequest {
	s.RoleFor = &v
	return s
}

func (s *CreateDataSourceRequest) Validate() error {
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

type CreateDataSourceRequestDataSourceStores struct {
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

func (s CreateDataSourceRequestDataSourceStores) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceRequestDataSourceStores) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestDataSourceStores) GetDataSourceStoreFrom() *string {
	return s.DataSourceStoreFrom
}

func (s *CreateDataSourceRequestDataSourceStores) GetDataSourceStoreId() *string {
	return s.DataSourceStoreId
}

func (s *CreateDataSourceRequestDataSourceStores) GetDataSourceStoreStatus() *string {
	return s.DataSourceStoreStatus
}

func (s *CreateDataSourceRequestDataSourceStores) GetLogProjectName() *string {
	return s.LogProjectName
}

func (s *CreateDataSourceRequestDataSourceStores) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *CreateDataSourceRequestDataSourceStores) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *CreateDataSourceRequestDataSourceStores) SetDataSourceStoreFrom(v string) *CreateDataSourceRequestDataSourceStores {
	s.DataSourceStoreFrom = &v
	return s
}

func (s *CreateDataSourceRequestDataSourceStores) SetDataSourceStoreId(v string) *CreateDataSourceRequestDataSourceStores {
	s.DataSourceStoreId = &v
	return s
}

func (s *CreateDataSourceRequestDataSourceStores) SetDataSourceStoreStatus(v string) *CreateDataSourceRequestDataSourceStores {
	s.DataSourceStoreStatus = &v
	return s
}

func (s *CreateDataSourceRequestDataSourceStores) SetLogProjectName(v string) *CreateDataSourceRequestDataSourceStores {
	s.LogProjectName = &v
	return s
}

func (s *CreateDataSourceRequestDataSourceStores) SetLogRegionId(v string) *CreateDataSourceRequestDataSourceStores {
	s.LogRegionId = &v
	return s
}

func (s *CreateDataSourceRequestDataSourceStores) SetLogStoreName(v string) *CreateDataSourceRequestDataSourceStores {
	s.LogStoreName = &v
	return s
}

func (s *CreateDataSourceRequestDataSourceStores) Validate() error {
	return dara.Validate(s)
}
