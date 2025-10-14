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
	SetUpdateTime(v int64) *CreateDataSourceRequest
	GetUpdateTime() *int64
}

type CreateDataSourceRequest struct {
	// example:
	//
	// center。
	DataSourceFrom *string   `json:"DataSourceFrom,omitempty" xml:"DataSourceFrom,omitempty"`
	DataSourceIds  []*string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// AD_LOG。
	DataSourceName             *string                                    `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	DataSourceRecognizeEnabled *bool                                      `json:"DataSourceRecognizeEnabled,omitempty" xml:"DataSourceRecognizeEnabled,omitempty"`
	DataSourceRecognizer       *string                                    `json:"DataSourceRecognizer,omitempty" xml:"DataSourceRecognizer,omitempty"`
	DataSourceReferences       []*string                                  `json:"DataSourceReferences,omitempty" xml:"DataSourceReferences,omitempty" type:"Repeated"`
	DataSourceStores           []*CreateDataSourceRequestDataSourceStores `json:"DataSourceStores,omitempty" xml:"DataSourceStores,omitempty" type:"Repeated"`
	DataSourceTemplateId       *string                                    `json:"DataSourceTemplateId,omitempty" xml:"DataSourceTemplateId,omitempty"`
	// example:
	//
	// preset。
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
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
	// mde_raw。
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// example:
	//
	// 173326*******。
	LogUserId *int64  `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	Order     *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor    *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

func (s *CreateDataSourceRequest) GetUpdateTime() *int64 {
	return s.UpdateTime
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

func (s *CreateDataSourceRequest) SetUpdateTime(v int64) *CreateDataSourceRequest {
	s.UpdateTime = &v
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
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataSourceStoreFrom   *string `json:"DataSourceStoreFrom,omitempty" xml:"DataSourceStoreFrom,omitempty"`
	DataSourceStoreId     *string `json:"DataSourceStoreId,omitempty" xml:"DataSourceStoreId,omitempty"`
	DataSourceStoreStatus *string `json:"DataSourceStoreStatus,omitempty" xml:"DataSourceStoreStatus,omitempty"`
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
	// actiontrail_management-events。
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	UpdateTime   *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateDataSourceRequestDataSourceStores) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceRequestDataSourceStores) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestDataSourceStores) GetCreateTime() *int64 {
	return s.CreateTime
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

func (s *CreateDataSourceRequestDataSourceStores) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *CreateDataSourceRequestDataSourceStores) SetCreateTime(v int64) *CreateDataSourceRequestDataSourceStores {
	s.CreateTime = &v
	return s
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

func (s *CreateDataSourceRequestDataSourceStores) SetUpdateTime(v int64) *CreateDataSourceRequestDataSourceStores {
	s.UpdateTime = &v
	return s
}

func (s *CreateDataSourceRequestDataSourceStores) Validate() error {
	return dara.Validate(s)
}
