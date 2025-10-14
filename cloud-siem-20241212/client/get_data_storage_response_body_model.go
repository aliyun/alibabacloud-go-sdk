// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDataStorageResponseBodyData) *GetDataStorageResponseBody
	GetData() *GetDataStorageResponseBodyData
	SetRequestId(v string) *GetDataStorageResponseBody
	GetRequestId() *string
}

type GetDataStorageResponseBody struct {
	Data *GetDataStorageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 81FB0DEA-52C1-55A0-8631-8E1B9A9D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBody) GetData() *GetDataStorageResponseBodyData {
	return s.Data
}

func (s *GetDataStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataStorageResponseBody) SetData(v *GetDataStorageResponseBodyData) *GetDataStorageResponseBody {
	s.Data = v
	return s
}

func (s *GetDataStorageResponseBody) SetRequestId(v string) *GetDataStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataStorageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataStorageResponseBodyData struct {
	// example:
	//
	// 100.0
	ColdStorageUsedCapacity *float64 `json:"ColdStorageUsedCapacity,omitempty" xml:"ColdStorageUsedCapacity,omitempty"`
	// example:
	//
	// cn-shanghai
	DataStorageRegionId *string `json:"DataStorageRegionId,omitempty" xml:"DataStorageRegionId,omitempty"`
	// example:
	//
	// deny
	DataStorageRegionPermission *string `json:"DataStorageRegionPermission,omitempty" xml:"DataStorageRegionPermission,omitempty"`
	// example:
	//
	// 100
	DataStorageTotalCapacity *int64 `json:"DataStorageTotalCapacity,omitempty" xml:"DataStorageTotalCapacity,omitempty"`
	// example:
	//
	// 100.0
	DataStorageUsedCapacity *float64 `json:"DataStorageUsedCapacity,omitempty" xml:"DataStorageUsedCapacity,omitempty"`
	// example:
	//
	// {\\"purchasedHotStorageCapacity\\":1000,\\"usedHotStorageCapacity\\":4.2,\\"usedHotStorageCapacityDetail\\":{\\"ap-southeast-1\\":4.2,\\"cn-shenzhen\\":0.0,\\"cn-shanghai\\":0.0}}
	DataStorageUsedCapacityDetail *string `json:"DataStorageUsedCapacityDetail,omitempty" xml:"DataStorageUsedCapacityDetail,omitempty"`
	// example:
	//
	// aliyun-cloudsiem-data-171835723111****-cn-shanghai
	LogProject             *string                                                 `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	NormalizationLogStores []*GetDataStorageResponseBodyDataNormalizationLogStores `json:"NormalizationLogStores,omitempty" xml:"NormalizationLogStores,omitempty" type:"Repeated"`
	NormalizationLogViews  []*GetDataStorageResponseBodyDataNormalizationLogViews  `json:"NormalizationLogViews,omitempty" xml:"NormalizationLogViews,omitempty" type:"Repeated"`
	SasLogStores           []*GetDataStorageResponseBodyDataSasLogStores           `json:"SasLogStores,omitempty" xml:"SasLogStores,omitempty" type:"Repeated"`
}

func (s GetDataStorageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyData) GetColdStorageUsedCapacity() *float64 {
	return s.ColdStorageUsedCapacity
}

func (s *GetDataStorageResponseBodyData) GetDataStorageRegionId() *string {
	return s.DataStorageRegionId
}

func (s *GetDataStorageResponseBodyData) GetDataStorageRegionPermission() *string {
	return s.DataStorageRegionPermission
}

func (s *GetDataStorageResponseBodyData) GetDataStorageTotalCapacity() *int64 {
	return s.DataStorageTotalCapacity
}

func (s *GetDataStorageResponseBodyData) GetDataStorageUsedCapacity() *float64 {
	return s.DataStorageUsedCapacity
}

func (s *GetDataStorageResponseBodyData) GetDataStorageUsedCapacityDetail() *string {
	return s.DataStorageUsedCapacityDetail
}

func (s *GetDataStorageResponseBodyData) GetLogProject() *string {
	return s.LogProject
}

func (s *GetDataStorageResponseBodyData) GetNormalizationLogStores() []*GetDataStorageResponseBodyDataNormalizationLogStores {
	return s.NormalizationLogStores
}

func (s *GetDataStorageResponseBodyData) GetNormalizationLogViews() []*GetDataStorageResponseBodyDataNormalizationLogViews {
	return s.NormalizationLogViews
}

func (s *GetDataStorageResponseBodyData) GetSasLogStores() []*GetDataStorageResponseBodyDataSasLogStores {
	return s.SasLogStores
}

func (s *GetDataStorageResponseBodyData) SetColdStorageUsedCapacity(v float64) *GetDataStorageResponseBodyData {
	s.ColdStorageUsedCapacity = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetDataStorageRegionId(v string) *GetDataStorageResponseBodyData {
	s.DataStorageRegionId = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetDataStorageRegionPermission(v string) *GetDataStorageResponseBodyData {
	s.DataStorageRegionPermission = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetDataStorageTotalCapacity(v int64) *GetDataStorageResponseBodyData {
	s.DataStorageTotalCapacity = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetDataStorageUsedCapacity(v float64) *GetDataStorageResponseBodyData {
	s.DataStorageUsedCapacity = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetDataStorageUsedCapacityDetail(v string) *GetDataStorageResponseBodyData {
	s.DataStorageUsedCapacityDetail = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetLogProject(v string) *GetDataStorageResponseBodyData {
	s.LogProject = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetNormalizationLogStores(v []*GetDataStorageResponseBodyDataNormalizationLogStores) *GetDataStorageResponseBodyData {
	s.NormalizationLogStores = v
	return s
}

func (s *GetDataStorageResponseBodyData) SetNormalizationLogViews(v []*GetDataStorageResponseBodyDataNormalizationLogViews) *GetDataStorageResponseBodyData {
	s.NormalizationLogViews = v
	return s
}

func (s *GetDataStorageResponseBodyData) SetSasLogStores(v []*GetDataStorageResponseBodyDataSasLogStores) *GetDataStorageResponseBodyData {
	s.SasLogStores = v
	return s
}

func (s *GetDataStorageResponseBodyData) Validate() error {
	if s.NormalizationLogStores != nil {
		for _, item := range s.NormalizationLogStores {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NormalizationLogViews != nil {
		for _, item := range s.NormalizationLogViews {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SasLogStores != nil {
		for _, item := range s.SasLogStores {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataStorageResponseBodyDataNormalizationLogStores struct {
	// example:
	//
	// vulnerability-activity
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// example:
	//
	// 180
	LogStoreTtl *int32 `json:"LogStoreTtl,omitempty" xml:"LogStoreTtl,omitempty"`
}

func (s GetDataStorageResponseBodyDataNormalizationLogStores) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyDataNormalizationLogStores) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) GetLogStoreTtl() *int32 {
	return s.LogStoreTtl
}

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) SetLogStoreName(v string) *GetDataStorageResponseBodyDataNormalizationLogStores {
	s.LogStoreName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) SetLogStoreTtl(v int32) *GetDataStorageResponseBodyDataNormalizationLogStores {
	s.LogStoreTtl = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) Validate() error {
	return dara.Validate(s)
}

type GetDataStorageResponseBodyDataNormalizationLogViews struct {
	// example:
	//
	// API security risk log
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// example:
	//
	// Security Category
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// 3
	DetectionRuleReferenceCount      *int32    `json:"DetectionRuleReferenceCount,omitempty" xml:"DetectionRuleReferenceCount,omitempty"`
	DetectionRuleReferenceProductIds []*string `json:"DetectionRuleReferenceProductIds,omitempty" xml:"DetectionRuleReferenceProductIds,omitempty" type:"Repeated"`
	// example:
	//
	// [{\\"SCHEMA\\":\\"AZURE_ACTIVE_DIRECTORY_AUDIT_ACTIVITY\\"}]
	LogSearchConditions *string `json:"LogSearchConditions,omitempty" xml:"LogSearchConditions,omitempty"`
	// example:
	//
	// risk-activity
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// example:
	//
	// true
	LogViewExisted *bool `json:"LogViewExisted,omitempty" xml:"LogViewExisted,omitempty"`
	// example:
	//
	// risk_activity
	LogViewName *string `json:"LogViewName,omitempty" xml:"LogViewName,omitempty"`
}

func (s GetDataStorageResponseBodyDataNormalizationLogViews) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyDataNormalizationLogViews) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) GetActivityName() *string {
	return s.ActivityName
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) GetCategoryName() *string {
	return s.CategoryName
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) GetDetectionRuleReferenceCount() *int32 {
	return s.DetectionRuleReferenceCount
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) GetDetectionRuleReferenceProductIds() []*string {
	return s.DetectionRuleReferenceProductIds
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) GetLogSearchConditions() *string {
	return s.LogSearchConditions
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) GetLogViewExisted() *bool {
	return s.LogViewExisted
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) GetLogViewName() *string {
	return s.LogViewName
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) SetActivityName(v string) *GetDataStorageResponseBodyDataNormalizationLogViews {
	s.ActivityName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) SetCategoryName(v string) *GetDataStorageResponseBodyDataNormalizationLogViews {
	s.CategoryName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) SetDetectionRuleReferenceCount(v int32) *GetDataStorageResponseBodyDataNormalizationLogViews {
	s.DetectionRuleReferenceCount = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) SetDetectionRuleReferenceProductIds(v []*string) *GetDataStorageResponseBodyDataNormalizationLogViews {
	s.DetectionRuleReferenceProductIds = v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) SetLogSearchConditions(v string) *GetDataStorageResponseBodyDataNormalizationLogViews {
	s.LogSearchConditions = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) SetLogStoreName(v string) *GetDataStorageResponseBodyDataNormalizationLogViews {
	s.LogStoreName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) SetLogViewExisted(v bool) *GetDataStorageResponseBodyDataNormalizationLogViews {
	s.LogViewExisted = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) SetLogViewName(v string) *GetDataStorageResponseBodyDataNormalizationLogViews {
	s.LogViewName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogViews) Validate() error {
	return dara.Validate(s)
}

type GetDataStorageResponseBodyDataSasLogStores struct {
	// example:
	//
	// sas-net-block
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// example:
	//
	// host
	LogDeliveryGroup *string `json:"LogDeliveryGroup,omitempty" xml:"LogDeliveryGroup,omitempty"`
	// example:
	//
	// deny
	LogDeliveryPermission *string `json:"LogDeliveryPermission,omitempty" xml:"LogDeliveryPermission,omitempty"`
	// example:
	//
	// enable
	LogDeliveryStatus *string `json:"LogDeliveryStatus,omitempty" xml:"LogDeliveryStatus,omitempty"`
	// example:
	//
	// 2025-07-16T15:10:29
	LogDeliveryUpdateTime *string `json:"LogDeliveryUpdateTime,omitempty" xml:"LogDeliveryUpdateTime,omitempty"`
	// example:
	//
	// Process Snapshot
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// example:
	//
	// [{\\"__topic__\\":\\"sas-net-block\\"}]
	LogSearchConditions *string `json:"LogSearchConditions,omitempty" xml:"LogSearchConditions,omitempty"`
	// example:
	//
	// true
	LogStoreExisted *bool `json:"LogStoreExisted,omitempty" xml:"LogStoreExisted,omitempty"`
	// example:
	//
	// sas-security-log
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// example:
	//
	// 180
	LogStoreTtl *int32 `json:"LogStoreTtl,omitempty" xml:"LogStoreTtl,omitempty"`
}

func (s GetDataStorageResponseBodyDataSasLogStores) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyDataSasLogStores) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogCode() *string {
	return s.LogCode
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogDeliveryGroup() *string {
	return s.LogDeliveryGroup
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogDeliveryPermission() *string {
	return s.LogDeliveryPermission
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogDeliveryStatus() *string {
	return s.LogDeliveryStatus
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogDeliveryUpdateTime() *string {
	return s.LogDeliveryUpdateTime
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogName() *string {
	return s.LogName
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogSearchConditions() *string {
	return s.LogSearchConditions
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogStoreExisted() *bool {
	return s.LogStoreExisted
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *GetDataStorageResponseBodyDataSasLogStores) GetLogStoreTtl() *int32 {
	return s.LogStoreTtl
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogCode(v string) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogCode = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogDeliveryGroup(v string) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogDeliveryGroup = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogDeliveryPermission(v string) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogDeliveryPermission = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogDeliveryStatus(v string) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogDeliveryStatus = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogDeliveryUpdateTime(v string) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogDeliveryUpdateTime = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogName(v string) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogSearchConditions(v string) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogSearchConditions = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogStoreExisted(v bool) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogStoreExisted = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogStoreName(v string) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogStoreName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) SetLogStoreTtl(v int32) *GetDataStorageResponseBodyDataSasLogStores {
	s.LogStoreTtl = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) Validate() error {
	return dara.Validate(s)
}
