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
	DataSourceFrom *string `json:"DataSourceFrom,omitempty" xml:"DataSourceFrom,omitempty"`
	// example:
	//
	// ds-014frtpy28m5ct2eoyo1。
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// ActiontrailLog。
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// example:
	//
	// true
	DataSourceRecognizeEnabled *bool                                      `json:"DataSourceRecognizeEnabled,omitempty" xml:"DataSourceRecognizeEnabled,omitempty"`
	DataSourceStores           []*UpdateDataSourceRequestDataSourceStores `json:"DataSourceStores,omitempty" xml:"DataSourceStores,omitempty" type:"Repeated"`
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// aliyun-cloudsiem-data-173326*******-cn-hangzhou。
	LogProjectName *string `json:"LogProjectName,omitempty" xml:"LogProjectName,omitempty"`
	// example:
	//
	// cn-hangzhou。
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// example:
	//
	// cn-rds-sqlaudit。
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// example:
	//
	// 173326*******。
	LogUserId  *int64  `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
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
	DataSourceStoreFrom *string `json:"DataSourceStoreFrom,omitempty" xml:"DataSourceStoreFrom,omitempty"`
	DataSourceStoreId   *string `json:"DataSourceStoreId,omitempty" xml:"DataSourceStoreId,omitempty"`
	// example:
	//
	// aliyun-cloudsiem-data-173326*******-cn-hangzhou。
	LogProjectName *string `json:"LogProjectName,omitempty" xml:"LogProjectName,omitempty"`
	// example:
	//
	// cn-hangzhou。
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// example:
	//
	// cn-rds-sqlaudit。
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
