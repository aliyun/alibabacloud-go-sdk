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
	SetUpdateTime(v int64) *CreateDataSourceShrinkRequest
	GetUpdateTime() *int64
}

type CreateDataSourceShrinkRequest struct {
	// example:
	//
	// center。
	DataSourceFrom      *string `json:"DataSourceFrom,omitempty" xml:"DataSourceFrom,omitempty"`
	DataSourceIdsShrink *string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty"`
	// example:
	//
	// AD_LOG。
	DataSourceName             *string                                          `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	DataSourceRecognizeEnabled *bool                                            `json:"DataSourceRecognizeEnabled,omitempty" xml:"DataSourceRecognizeEnabled,omitempty"`
	DataSourceRecognizer       *string                                          `json:"DataSourceRecognizer,omitempty" xml:"DataSourceRecognizer,omitempty"`
	DataSourceReferencesShrink *string                                          `json:"DataSourceReferences,omitempty" xml:"DataSourceReferences,omitempty"`
	DataSourceStores           []*CreateDataSourceShrinkRequestDataSourceStores `json:"DataSourceStores,omitempty" xml:"DataSourceStores,omitempty" type:"Repeated"`
	DataSourceTemplateId       *string                                          `json:"DataSourceTemplateId,omitempty" xml:"DataSourceTemplateId,omitempty"`
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

func (s *CreateDataSourceShrinkRequest) GetUpdateTime() *int64 {
	return s.UpdateTime
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

func (s *CreateDataSourceShrinkRequest) SetUpdateTime(v int64) *CreateDataSourceShrinkRequest {
	s.UpdateTime = &v
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

func (s CreateDataSourceShrinkRequestDataSourceStores) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceShrinkRequestDataSourceStores) GoString() string {
	return s.String()
}

func (s *CreateDataSourceShrinkRequestDataSourceStores) GetCreateTime() *int64 {
	return s.CreateTime
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

func (s *CreateDataSourceShrinkRequestDataSourceStores) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *CreateDataSourceShrinkRequestDataSourceStores) SetCreateTime(v int64) *CreateDataSourceShrinkRequestDataSourceStores {
	s.CreateTime = &v
	return s
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

func (s *CreateDataSourceShrinkRequestDataSourceStores) SetUpdateTime(v int64) *CreateDataSourceShrinkRequestDataSourceStores {
	s.UpdateTime = &v
	return s
}

func (s *CreateDataSourceShrinkRequestDataSourceStores) Validate() error {
	return dara.Validate(s)
}
