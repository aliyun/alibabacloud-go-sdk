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
	// The returned details.
	Data *GetDataStorageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
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
	// The cold storage capacity used by user logs.
	//
	// example:
	//
	// 100.0
	ColdStorageUsedCapacity *float64 `json:"ColdStorageUsedCapacity,omitempty" xml:"ColdStorageUsedCapacity,omitempty"`
	// The storage region of user-side logs.
	//
	// example:
	//
	// cn-shanghai
	DataStorageRegionId *string `json:"DataStorageRegionId,omitempty" xml:"DataStorageRegionId,omitempty"`
	// Indicates whether the storage region can be modified. By default, modification is not allowed. Contact the product manager to reset the region. The region can be reset only once. Valid values:
	//
	// - allow: Modification is allowed.
	//
	// - deny: Modification is not allowed.
	//
	// example:
	//
	// deny
	DataStorageRegionPermission *string `json:"DataStorageRegionPermission,omitempty" xml:"DataStorageRegionPermission,omitempty"`
	// The storage capacity purchased in the prepaid scenario.
	//
	// example:
	//
	// 100
	DataStorageTotalCapacity *int64 `json:"DataStorageTotalCapacity,omitempty" xml:"DataStorageTotalCapacity,omitempty"`
	// The storage capacity used in user log management.
	//
	// example:
	//
	// 100.0
	DataStorageUsedCapacity *float64 `json:"DataStorageUsedCapacity,omitempty" xml:"DataStorageUsedCapacity,omitempty"`
	// The storage usage details in log management.
	//
	// example:
	//
	// {\\"purchasedHotStorageCapacity\\":1000,\\"usedHotStorageCapacity\\":4.2,\\"usedHotStorageCapacityDetail\\":{\\"ap-southeast-1\\":4.2,\\"cn-shenzhen\\":0.0,\\"cn-shanghai\\":0.0}}
	DataStorageUsedCapacityDetail *string `json:"DataStorageUsedCapacityDetail,omitempty" xml:"DataStorageUsedCapacityDetail,omitempty"`
	// The name of the Simple Log Service (SLS) project that stores user logs.
	//
	// example:
	//
	// aliyun-cloudsiem-data-171835723111****-cn-shanghai
	LogProject                   *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	LogProjectState              *string `json:"LogProjectState,omitempty" xml:"LogProjectState,omitempty"`
	LogProjectStateChangeAllowed *bool   `json:"LogProjectStateChangeAllowed,omitempty" xml:"LogProjectStateChangeAllowed,omitempty"`
	LogServiceDisabled           *bool   `json:"LogServiceDisabled,omitempty" xml:"LogServiceDisabled,omitempty"`
	// The details of Logstores for normalized data.
	NormalizationLogStores []*GetDataStorageResponseBodyDataNormalizationLogStores `json:"NormalizationLogStores,omitempty" xml:"NormalizationLogStores,omitempty" type:"Repeated"`
	// The details of normalized datasets.
	NormalizationLogViews []*GetDataStorageResponseBodyDataNormalizationLogViews `json:"NormalizationLogViews,omitempty" xml:"NormalizationLogViews,omitempty" type:"Repeated"`
	// The list of record Logstores.
	RecordLogStores []*GetDataStorageResponseBodyDataRecordLogStores `json:"RecordLogStores,omitempty" xml:"RecordLogStores,omitempty" type:"Repeated"`
	// The details of raw log storage in Security Center.
	SasLogStores []*GetDataStorageResponseBodyDataSasLogStores `json:"SasLogStores,omitempty" xml:"SasLogStores,omitempty" type:"Repeated"`
	// The list of SIEM V1 legacy Logstores.
	UnusedLogStores []*GetDataStorageResponseBodyDataUnusedLogStores `json:"UnusedLogStores,omitempty" xml:"UnusedLogStores,omitempty" type:"Repeated"`
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

func (s *GetDataStorageResponseBodyData) GetLogProjectState() *string {
	return s.LogProjectState
}

func (s *GetDataStorageResponseBodyData) GetLogProjectStateChangeAllowed() *bool {
	return s.LogProjectStateChangeAllowed
}

func (s *GetDataStorageResponseBodyData) GetLogServiceDisabled() *bool {
	return s.LogServiceDisabled
}

func (s *GetDataStorageResponseBodyData) GetNormalizationLogStores() []*GetDataStorageResponseBodyDataNormalizationLogStores {
	return s.NormalizationLogStores
}

func (s *GetDataStorageResponseBodyData) GetNormalizationLogViews() []*GetDataStorageResponseBodyDataNormalizationLogViews {
	return s.NormalizationLogViews
}

func (s *GetDataStorageResponseBodyData) GetRecordLogStores() []*GetDataStorageResponseBodyDataRecordLogStores {
	return s.RecordLogStores
}

func (s *GetDataStorageResponseBodyData) GetSasLogStores() []*GetDataStorageResponseBodyDataSasLogStores {
	return s.SasLogStores
}

func (s *GetDataStorageResponseBodyData) GetUnusedLogStores() []*GetDataStorageResponseBodyDataUnusedLogStores {
	return s.UnusedLogStores
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

func (s *GetDataStorageResponseBodyData) SetLogProjectState(v string) *GetDataStorageResponseBodyData {
	s.LogProjectState = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetLogProjectStateChangeAllowed(v bool) *GetDataStorageResponseBodyData {
	s.LogProjectStateChangeAllowed = &v
	return s
}

func (s *GetDataStorageResponseBodyData) SetLogServiceDisabled(v bool) *GetDataStorageResponseBodyData {
	s.LogServiceDisabled = &v
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

func (s *GetDataStorageResponseBodyData) SetRecordLogStores(v []*GetDataStorageResponseBodyDataRecordLogStores) *GetDataStorageResponseBodyData {
	s.RecordLogStores = v
	return s
}

func (s *GetDataStorageResponseBodyData) SetSasLogStores(v []*GetDataStorageResponseBodyDataSasLogStores) *GetDataStorageResponseBodyData {
	s.SasLogStores = v
	return s
}

func (s *GetDataStorageResponseBodyData) SetUnusedLogStores(v []*GetDataStorageResponseBodyDataUnusedLogStores) *GetDataStorageResponseBodyData {
	s.UnusedLogStores = v
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
	if s.RecordLogStores != nil {
		for _, item := range s.RecordLogStores {
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
	if s.UnusedLogStores != nil {
		for _, item := range s.UnusedLogStores {
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
	// The name of the Logstore that stores normalized data.
	//
	// example:
	//
	// vulnerability-activity
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The storage duration of normalized data.
	//
	// example:
	//
	// 180
	LogStoreTtl *int32 `json:"LogStoreTtl,omitempty" xml:"LogStoreTtl,omitempty"`
	// The hot storage used capacity.
	//
	// example:
	//
	// 10.333
	UsedCapacity *float64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
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

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) GetUsedCapacity() *float64 {
	return s.UsedCapacity
}

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) SetLogStoreName(v string) *GetDataStorageResponseBodyDataNormalizationLogStores {
	s.LogStoreName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) SetLogStoreTtl(v int32) *GetDataStorageResponseBodyDataNormalizationLogStores {
	s.LogStoreTtl = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) SetUsedCapacity(v float64) *GetDataStorageResponseBodyDataNormalizationLogStores {
	s.UsedCapacity = &v
	return s
}

func (s *GetDataStorageResponseBodyDataNormalizationLogStores) Validate() error {
	return dara.Validate(s)
}

type GetDataStorageResponseBodyDataNormalizationLogViews struct {
	// The log type of the normalized log.
	//
	// example:
	//
	// API security risk log
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// The category of the normalized log.
	//
	// example:
	//
	// Security Category
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The number of times the normalized dataset is referenced in the threat analysis access center.
	//
	// example:
	//
	// 3
	DetectionRuleReferenceCount *int32 `json:"DetectionRuleReferenceCount,omitempty" xml:"DetectionRuleReferenceCount,omitempty"`
	// The list of products that reference the normalized dataset in the threat analysis access center.
	DetectionRuleReferenceProductIds []*string `json:"DetectionRuleReferenceProductIds,omitempty" xml:"DetectionRuleReferenceProductIds,omitempty" type:"Repeated"`
	// The query statement used to query the log type in the normalized dataset.
	//
	// example:
	//
	// [{\\"SCHEMA\\":\\"AZURE_ACTIVE_DIRECTORY_AUDIT_ACTIVITY\\"}]
	LogSearchConditions *string `json:"LogSearchConditions,omitempty" xml:"LogSearchConditions,omitempty"`
	// The Logstore where threat analysis stores normalized logs.
	//
	// example:
	//
	// risk-activity
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// Indicates whether the normalized dataset exists. Valid values:
	//
	// - true: The normalized dataset exists.
	//
	// - false: The normalized dataset does not exist.
	//
	// example:
	//
	// true
	LogViewExisted *bool `json:"LogViewExisted,omitempty" xml:"LogViewExisted,omitempty"`
	// The name of the normalized dataset.
	//
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

type GetDataStorageResponseBodyDataRecordLogStores struct {
	// The Logstore name.
	//
	// example:
	//
	// alert-record
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The Logstore TTL.
	//
	// example:
	//
	// 90
	LogStoreTtl *int32 `json:"LogStoreTtl,omitempty" xml:"LogStoreTtl,omitempty"`
	// The Logstore used capacity.
	//
	// example:
	//
	// 11.111
	UsedCapacity *float64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s GetDataStorageResponseBodyDataRecordLogStores) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyDataRecordLogStores) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyDataRecordLogStores) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *GetDataStorageResponseBodyDataRecordLogStores) GetLogStoreTtl() *int32 {
	return s.LogStoreTtl
}

func (s *GetDataStorageResponseBodyDataRecordLogStores) GetUsedCapacity() *float64 {
	return s.UsedCapacity
}

func (s *GetDataStorageResponseBodyDataRecordLogStores) SetLogStoreName(v string) *GetDataStorageResponseBodyDataRecordLogStores {
	s.LogStoreName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataRecordLogStores) SetLogStoreTtl(v int32) *GetDataStorageResponseBodyDataRecordLogStores {
	s.LogStoreTtl = &v
	return s
}

func (s *GetDataStorageResponseBodyDataRecordLogStores) SetUsedCapacity(v float64) *GetDataStorageResponseBodyDataRecordLogStores {
	s.UsedCapacity = &v
	return s
}

func (s *GetDataStorageResponseBodyDataRecordLogStores) Validate() error {
	return dara.Validate(s)
}

type GetDataStorageResponseBodyDataSasLogStores struct {
	// The log code.
	//
	// example:
	//
	// sas-net-block
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The group to which the log belongs. Valid values:
	//
	// - host: Host logs.
	//
	// - security: Security logs.
	//
	// example:
	//
	// host
	LogDeliveryGroup *string `json:"LogDeliveryGroup,omitempty" xml:"LogDeliveryGroup,omitempty"`
	// Indicates whether you are allowed to toggle the log delivery switch. Log delivery cannot be performed if the service is not purchased. Valid values:
	//
	// - allow: Allowed.
	//
	// - deny: Not allowed.
	//
	// example:
	//
	// deny
	LogDeliveryPermission *string `json:"LogDeliveryPermission,omitempty" xml:"LogDeliveryPermission,omitempty"`
	// The log delivery status. Valid values:
	//
	// - enable: Log delivery is enabled.
	//
	// - disable: Log delivery is disabled.
	//
	// example:
	//
	// enable
	LogDeliveryStatus *string `json:"LogDeliveryStatus,omitempty" xml:"LogDeliveryStatus,omitempty"`
	// The time when the log delivery was last modified.
	//
	// example:
	//
	// 2025-07-16T15:10:29
	LogDeliveryUpdateTime *string `json:"LogDeliveryUpdateTime,omitempty" xml:"LogDeliveryUpdateTime,omitempty"`
	// The log name.
	//
	// example:
	//
	// Process Snapshot
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// The default log query conditions for the log. When multiple logs are stored in the same Logstore, query conditions are required to perform a log query for a specific log.
	//
	// example:
	//
	// [{\\"__topic__\\":\\"sas-net-block\\"}]
	LogSearchConditions *string `json:"LogSearchConditions,omitempty" xml:"LogSearchConditions,omitempty"`
	// Indicates whether the Logstore where the log is stored exists. Valid values:
	//
	// - true: The Logstore exists.
	//
	// - false: The Logstore does not exist.
	//
	// example:
	//
	// true
	LogStoreExisted *bool `json:"LogStoreExisted,omitempty" xml:"LogStoreExisted,omitempty"`
	// The name of the Logstore where the log is stored.
	//
	// example:
	//
	// sas-security-log
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The storage duration of the Logstore where the log is stored. Logs are stored for at least 30 days.
	//
	// example:
	//
	// 180
	LogStoreTtl *int32 `json:"LogStoreTtl,omitempty" xml:"LogStoreTtl,omitempty"`
	// The hot storage used capacity.
	//
	// example:
	//
	// 10.333
	UsedCapacity *float64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
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

func (s *GetDataStorageResponseBodyDataSasLogStores) GetUsedCapacity() *float64 {
	return s.UsedCapacity
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

func (s *GetDataStorageResponseBodyDataSasLogStores) SetUsedCapacity(v float64) *GetDataStorageResponseBodyDataSasLogStores {
	s.UsedCapacity = &v
	return s
}

func (s *GetDataStorageResponseBodyDataSasLogStores) Validate() error {
	return dara.Validate(s)
}

type GetDataStorageResponseBodyDataUnusedLogStores struct {
	// The Logstore name.
	//
	// example:
	//
	// cloud-siem
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The data storage duration.
	//
	// example:
	//
	// 180
	LogStoreTtl *int32 `json:"LogStoreTtl,omitempty" xml:"LogStoreTtl,omitempty"`
	// The hot storage used capacity.
	//
	// example:
	//
	// 10.333
	UsedCapacity *float64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s GetDataStorageResponseBodyDataUnusedLogStores) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponseBodyDataUnusedLogStores) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponseBodyDataUnusedLogStores) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *GetDataStorageResponseBodyDataUnusedLogStores) GetLogStoreTtl() *int32 {
	return s.LogStoreTtl
}

func (s *GetDataStorageResponseBodyDataUnusedLogStores) GetUsedCapacity() *float64 {
	return s.UsedCapacity
}

func (s *GetDataStorageResponseBodyDataUnusedLogStores) SetLogStoreName(v string) *GetDataStorageResponseBodyDataUnusedLogStores {
	s.LogStoreName = &v
	return s
}

func (s *GetDataStorageResponseBodyDataUnusedLogStores) SetLogStoreTtl(v int32) *GetDataStorageResponseBodyDataUnusedLogStores {
	s.LogStoreTtl = &v
	return s
}

func (s *GetDataStorageResponseBodyDataUnusedLogStores) SetUsedCapacity(v float64) *GetDataStorageResponseBodyDataUnusedLogStores {
	s.UsedCapacity = &v
	return s
}

func (s *GetDataStorageResponseBodyDataUnusedLogStores) Validate() error {
	return dara.Validate(s)
}
