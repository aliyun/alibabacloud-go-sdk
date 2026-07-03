// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceFrom(v string) *UpdateDataSourceRequest
	GetDataSourceFrom() *string
	SetDataSourceId(v string) *UpdateDataSourceRequest
	GetDataSourceId() *string
	SetDataSourceName(v string) *UpdateDataSourceRequest
	GetDataSourceName() *string
	SetDataSourceRecognizeEnabled(v bool) *UpdateDataSourceRequest
	GetDataSourceRecognizeEnabled() *bool
	SetDataSourceStores(v []*UpdateDataSourceRequestDataSourceStores) *UpdateDataSourceRequest
	GetDataSourceStores() []*UpdateDataSourceRequestDataSourceStores
	SetLang(v string) *UpdateDataSourceRequest
	GetLang() *string
	SetLogProjectName(v string) *UpdateDataSourceRequest
	GetLogProjectName() *string
	SetLogRegionId(v string) *UpdateDataSourceRequest
	GetLogRegionId() *string
	SetLogStoreName(v string) *UpdateDataSourceRequest
	GetLogStoreName() *string
	SetLogUserId(v int64) *UpdateDataSourceRequest
	GetLogUserId() *int64
	SetOrderField(v string) *UpdateDataSourceRequest
	GetOrderField() *string
	SetRegionId(v string) *UpdateDataSourceRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateDataSourceRequest
	GetRoleFor() *int64
}

type UpdateDataSourceRequest struct {
	// The source of the data. Valid values:
	//
	// - center
	//
	// - custom
	//
	// example:
	//
	// custom
	DataSourceFrom *string `json:"DataSourceFrom,omitempty" xml:"DataSourceFrom,omitempty"`
	// The ID of the data source.
	//
	// example:
	//
	// ds-014frtpy28m5ct2eoyo1
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// ActiontrailLog
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// Specifies whether to automatically discover new Logstores.
	//
	// example:
	//
	// true
	DataSourceRecognizeEnabled *bool `json:"DataSourceRecognizeEnabled,omitempty" xml:"DataSourceRecognizeEnabled,omitempty"`
	// The list of Simple Log Service Logstores.
	DataSourceStores []*UpdateDataSourceRequestDataSourceStores `json:"DataSourceStores,omitempty" xml:"DataSourceStores,omitempty" type:"Repeated"`
	// The language of the response messages. Valid values:
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
	// cn-rds-sqlaudit
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The user ID for data access.
	//
	// example:
	//
	// 173326*******
	LogUserId *int64 `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	// The field to use for sorting the rule list. Valid values:
	//
	// - GmtModified: Sorts the list by modification time.
	//
	// - Id: Sorts the list by rule ID. This is the default value.
	//
	// example:
	//
	// Id
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The region of the Data Management center for threat analysis. Select a region based on the location of the assets. Valid values:
	//
	// - cn-hangzhou: Assets are in the Chinese mainland.
	//
	// - ap-southeast-1: Assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member whose perspective the administrator switches to.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s UpdateDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceRequest) GetDataSourceFrom() *string {
	return s.DataSourceFrom
}

func (s *UpdateDataSourceRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *UpdateDataSourceRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *UpdateDataSourceRequest) GetDataSourceRecognizeEnabled() *bool {
	return s.DataSourceRecognizeEnabled
}

func (s *UpdateDataSourceRequest) GetDataSourceStores() []*UpdateDataSourceRequestDataSourceStores {
	return s.DataSourceStores
}

func (s *UpdateDataSourceRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataSourceRequest) GetLogProjectName() *string {
	return s.LogProjectName
}

func (s *UpdateDataSourceRequest) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *UpdateDataSourceRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *UpdateDataSourceRequest) GetLogUserId() *int64 {
	return s.LogUserId
}

func (s *UpdateDataSourceRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *UpdateDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataSourceRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateDataSourceRequest) SetDataSourceFrom(v string) *UpdateDataSourceRequest {
	s.DataSourceFrom = &v
	return s
}

func (s *UpdateDataSourceRequest) SetDataSourceId(v string) *UpdateDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *UpdateDataSourceRequest) SetDataSourceName(v string) *UpdateDataSourceRequest {
	s.DataSourceName = &v
	return s
}

func (s *UpdateDataSourceRequest) SetDataSourceRecognizeEnabled(v bool) *UpdateDataSourceRequest {
	s.DataSourceRecognizeEnabled = &v
	return s
}

func (s *UpdateDataSourceRequest) SetDataSourceStores(v []*UpdateDataSourceRequestDataSourceStores) *UpdateDataSourceRequest {
	s.DataSourceStores = v
	return s
}

func (s *UpdateDataSourceRequest) SetLang(v string) *UpdateDataSourceRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataSourceRequest) SetLogProjectName(v string) *UpdateDataSourceRequest {
	s.LogProjectName = &v
	return s
}

func (s *UpdateDataSourceRequest) SetLogRegionId(v string) *UpdateDataSourceRequest {
	s.LogRegionId = &v
	return s
}

func (s *UpdateDataSourceRequest) SetLogStoreName(v string) *UpdateDataSourceRequest {
	s.LogStoreName = &v
	return s
}

func (s *UpdateDataSourceRequest) SetLogUserId(v int64) *UpdateDataSourceRequest {
	s.LogUserId = &v
	return s
}

func (s *UpdateDataSourceRequest) SetOrderField(v string) *UpdateDataSourceRequest {
	s.OrderField = &v
	return s
}

func (s *UpdateDataSourceRequest) SetRegionId(v string) *UpdateDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataSourceRequest) SetRoleFor(v int64) *UpdateDataSourceRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateDataSourceRequest) Validate() error {
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

type UpdateDataSourceRequestDataSourceStores struct {
	// The source of the data. Valid values:
	//
	// - center
	//
	// - custom
	//
	// example:
	//
	// custom
	DataSourceStoreFrom *string `json:"DataSourceStoreFrom,omitempty" xml:"DataSourceStoreFrom,omitempty"`
	// The ID of the log storage.
	//
	// example:
	//
	// 1
	DataSourceStoreId *string `json:"DataSourceStoreId,omitempty" xml:"DataSourceStoreId,omitempty"`
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
	// cn-rds-sqlaudit
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
}

func (s UpdateDataSourceRequestDataSourceStores) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceRequestDataSourceStores) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceRequestDataSourceStores) GetDataSourceStoreFrom() *string {
	return s.DataSourceStoreFrom
}

func (s *UpdateDataSourceRequestDataSourceStores) GetDataSourceStoreId() *string {
	return s.DataSourceStoreId
}

func (s *UpdateDataSourceRequestDataSourceStores) GetLogProjectName() *string {
	return s.LogProjectName
}

func (s *UpdateDataSourceRequestDataSourceStores) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *UpdateDataSourceRequestDataSourceStores) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *UpdateDataSourceRequestDataSourceStores) SetDataSourceStoreFrom(v string) *UpdateDataSourceRequestDataSourceStores {
	s.DataSourceStoreFrom = &v
	return s
}

func (s *UpdateDataSourceRequestDataSourceStores) SetDataSourceStoreId(v string) *UpdateDataSourceRequestDataSourceStores {
	s.DataSourceStoreId = &v
	return s
}

func (s *UpdateDataSourceRequestDataSourceStores) SetLogProjectName(v string) *UpdateDataSourceRequestDataSourceStores {
	s.LogProjectName = &v
	return s
}

func (s *UpdateDataSourceRequestDataSourceStores) SetLogRegionId(v string) *UpdateDataSourceRequestDataSourceStores {
	s.LogRegionId = &v
	return s
}

func (s *UpdateDataSourceRequestDataSourceStores) SetLogStoreName(v string) *UpdateDataSourceRequestDataSourceStores {
	s.LogStoreName = &v
	return s
}

func (s *UpdateDataSourceRequestDataSourceStores) Validate() error {
	return dara.Validate(s)
}
