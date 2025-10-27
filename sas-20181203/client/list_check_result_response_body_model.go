// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChecks(v []*ListCheckResultResponseBodyChecks) *ListCheckResultResponseBody
	GetChecks() []*ListCheckResultResponseBodyChecks
	SetPageInfo(v *ListCheckResultResponseBodyPageInfo) *ListCheckResultResponseBody
	GetPageInfo() *ListCheckResultResponseBodyPageInfo
	SetRequestId(v string) *ListCheckResultResponseBody
	GetRequestId() *string
}

type ListCheckResultResponseBody struct {
	// The check items.
	Checks []*ListCheckResultResponseBodyChecks `json:"Checks,omitempty" xml:"Checks,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListCheckResultResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F9B6DD67-B289-5406-B35C-B0F4A217S23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCheckResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListCheckResultResponseBody) GetChecks() []*ListCheckResultResponseBodyChecks {
	return s.Checks
}

func (s *ListCheckResultResponseBody) GetPageInfo() *ListCheckResultResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListCheckResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCheckResultResponseBody) SetChecks(v []*ListCheckResultResponseBodyChecks) *ListCheckResultResponseBody {
	s.Checks = v
	return s
}

func (s *ListCheckResultResponseBody) SetPageInfo(v *ListCheckResultResponseBodyPageInfo) *ListCheckResultResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListCheckResultResponseBody) SetRequestId(v string) *ListCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCheckResultResponseBody) Validate() error {
	if s.Checks != nil {
		for _, item := range s.Checks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCheckResultResponseBodyChecks struct {
	// The subtype of the cloud service.
	//
	// example:
	//
	// 0
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- **0**: an ECS instance
	//
	// 	- **1**: a SLB instance
	//
	// 	- **2**: a NAT gateway
	//
	// 	- **3**: an ApsaraDB RDS instance
	//
	// 	- **4**: an ApsaraDB for MongoDB instance
	//
	// 	- **5**: an ApsaraDB for Redis instance
	//
	// 	- **6**: a container image
	//
	// 	- **7**: a container
	//
	// example:
	//
	// 0
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The service provider of the asset. Valid values:
	//
	// 	- **0**: Alibaba Cloud
	//
	// 	- **3**: Huawei Cloud
	//
	// 	- **4**: Microsoft Azure
	//
	// 	- **5**: AWS
	//
	// 	- **7**: Tencent Cloud
	//
	// example:
	//
	// 3
	AssetVendor *int32 `json:"AssetVendor,omitempty" xml:"AssetVendor,omitempty"`
	// The ID of the check item.
	//
	// example:
	//
	// 5
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The check policies.
	CheckPolicies []*ListCheckResultResponseBodyChecksCheckPolicies `json:"CheckPolicies,omitempty" xml:"CheckPolicies,omitempty" type:"Repeated"`
	// The type of the check item. Valid values:
	//
	// 	- **0**: paid
	//
	// 	- **1**: free
	//
	// example:
	//
	// 1
	CheckSaleType *int32 `json:"CheckSaleType,omitempty" xml:"CheckSaleType,omitempty"`
	// The name of the check item.
	//
	// example:
	//
	// OSS-PublicReadOpenManifestFileWithoutEncryption
	CheckShowName *string `json:"CheckShowName,omitempty" xml:"CheckShowName,omitempty"`
	// The source type of the situation awareness check item:
	//
	// - **CUSTOM**: User-defined
	//
	// - **SYSTEM**: Predefined by the situation awareness platform
	//
	// example:
	//
	// SYSTEM
	CheckType *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The asset subtype of the cloud service. Valid values:
	//
	// 	- If the **InstanceType*	- parameter is set to **ECS**, this parameter supports the following valid values:
	//
	//     	- **INSTANCE**
	//
	//     	- **DISK**
	//
	//     	- **SECURITY_GROUP**
	//
	// 	- If the **InstanceType*	- parameter is set to **ACR**, this parameter supports the following valid values:
	//
	//     	- **REPOSITORY_ENTERPRISE**
	//
	//     	- **REPOSITORY_PERSON**
	//
	// 	- If the **InstanceType*	- parameter is set to **RAM**, this parameter supports the following valid values:
	//
	//     	- **ALIAS**
	//
	//     	- **USER**
	//
	//     	- **POLICY**
	//
	//     	- **GROUP**
	//
	// 	- If the **InstanceType*	- parameter is set to **WAF**, this parameter supports the following valid values:
	//
	//     	- **DOMAIN**
	//
	// 	- If the **InstanceType*	- parameter is set to other values, this parameter supports the following valid values:
	//
	//     	- **INSTANCE**
	//
	// example:
	//
	// DISK
	InstanceSubType *string `json:"InstanceSubType,omitempty" xml:"InstanceSubType,omitempty"`
	// The asset type of the cloud service. Valid values:
	//
	// 	- **ECS**: ECS
	//
	// 	- **SLB**: SLB
	//
	// 	- **RDS**: ApsaraDB RDS
	//
	// 	- **MONGODB**: MongoDB
	//
	// 	- **KVSTORE**: Redis
	//
	// 	- **ACR**: Container Registry
	//
	// 	- **CSK**: ACK
	//
	// 	- **VPC**: VPC
	//
	// 	- **ACTIONTRAIL**: ActionTrail
	//
	// 	- **CDN**: CDN
	//
	// 	- **CAS**: Certificate Management Service (formerly SSL Certificates Service)
	//
	// 	- **RDC**: Apsara Devops
	//
	// 	- **RAM**: RAM
	//
	// 	- **DDOS**: Anti-DDoS
	//
	// 	- **WAF**: WAF
	//
	// 	- **OSS**: OSS
	//
	// 	- **POLARDB**: PolarDB
	//
	// 	- **POSTGRESQL**: ApsaraDB RDS for PostgreSQL
	//
	// 	- **MSE**: MSE
	//
	// 	- **NAS**: NAS
	//
	// 	- **SDDP**: SDDP
	//
	// 	- **EIP**: EIP
	//
	// example:
	//
	// ECS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The timestamp when the last check was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1657793398000
	LastCheckTime *int64 `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
	// Indicates whether fixing is supported. Valid values:
	//
	// 	- **SUPPORT_REPAIR**
	//
	// 	- **NOT_SUPPORT_REPAIR**
	//
	// example:
	//
	// NOT_SUPPORT_REPAIR
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The risk level of the check item. Valid values:
	//
	// 	- **HIGH**
	//
	// 	- **MEDIUM**
	//
	// 	- **LOW**
	//
	// example:
	//
	// HIGH
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The status of the check item. Valid values:
	//
	// 	- **PASS**: passed
	//
	// 	- **NOT_PASS**: failed
	//
	// 	- **CHECKING**: being checked
	//
	// 	- **NOT_CHECK**: not checked
	//
	// 	- **WHITELIST**: added to the whitelist
	//
	// example:
	//
	// PASS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The message returned if the status of the check item is abnormal.
	//
	// example:
	//
	// TIMEOUT
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The ID of the check task.
	//
	// example:
	//
	// 64
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Indicates whether the TRIAL permission is required.
	//
	// example:
	//
	// true
	TrialPermission *bool `json:"TrialPermission,omitempty" xml:"TrialPermission,omitempty"`
	// Check whether the data delivery period for ActionTrail is enabled for more than 30 days to establish a baseline of behaviour.
	//
	// 	- **0**: REQUIRED
	//
	// 	- **1**: NOT REQUIRED
	//
	// example:
	//
	// 1
	TrialPermissionType *int32 `json:"TrialPermissionType,omitempty" xml:"TrialPermissionType,omitempty"`
	// The cloud service provider.
	//
	// example:
	//
	// ALIYUN
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The name of the cloud service provider.
	//
	// example:
	//
	// Aliyun
	VendorShowName *string `json:"VendorShowName,omitempty" xml:"VendorShowName,omitempty"`
}

func (s ListCheckResultResponseBodyChecks) String() string {
	return dara.Prettify(s)
}

func (s ListCheckResultResponseBodyChecks) GoString() string {
	return s.String()
}

func (s *ListCheckResultResponseBodyChecks) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *ListCheckResultResponseBodyChecks) GetAssetType() *int32 {
	return s.AssetType
}

func (s *ListCheckResultResponseBodyChecks) GetAssetVendor() *int32 {
	return s.AssetVendor
}

func (s *ListCheckResultResponseBodyChecks) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ListCheckResultResponseBodyChecks) GetCheckPolicies() []*ListCheckResultResponseBodyChecksCheckPolicies {
	return s.CheckPolicies
}

func (s *ListCheckResultResponseBodyChecks) GetCheckSaleType() *int32 {
	return s.CheckSaleType
}

func (s *ListCheckResultResponseBodyChecks) GetCheckShowName() *string {
	return s.CheckShowName
}

func (s *ListCheckResultResponseBodyChecks) GetCheckType() *string {
	return s.CheckType
}

func (s *ListCheckResultResponseBodyChecks) GetInstanceSubType() *string {
	return s.InstanceSubType
}

func (s *ListCheckResultResponseBodyChecks) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListCheckResultResponseBodyChecks) GetLastCheckTime() *int64 {
	return s.LastCheckTime
}

func (s *ListCheckResultResponseBodyChecks) GetOperationType() *string {
	return s.OperationType
}

func (s *ListCheckResultResponseBodyChecks) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListCheckResultResponseBodyChecks) GetStatus() *string {
	return s.Status
}

func (s *ListCheckResultResponseBodyChecks) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *ListCheckResultResponseBodyChecks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListCheckResultResponseBodyChecks) GetTrialPermission() *bool {
	return s.TrialPermission
}

func (s *ListCheckResultResponseBodyChecks) GetTrialPermissionType() *int32 {
	return s.TrialPermissionType
}

func (s *ListCheckResultResponseBodyChecks) GetVendor() *string {
	return s.Vendor
}

func (s *ListCheckResultResponseBodyChecks) GetVendorShowName() *string {
	return s.VendorShowName
}

func (s *ListCheckResultResponseBodyChecks) SetAssetSubType(v int32) *ListCheckResultResponseBodyChecks {
	s.AssetSubType = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetAssetType(v int32) *ListCheckResultResponseBodyChecks {
	s.AssetType = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetAssetVendor(v int32) *ListCheckResultResponseBodyChecks {
	s.AssetVendor = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetCheckId(v int64) *ListCheckResultResponseBodyChecks {
	s.CheckId = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetCheckPolicies(v []*ListCheckResultResponseBodyChecksCheckPolicies) *ListCheckResultResponseBodyChecks {
	s.CheckPolicies = v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetCheckSaleType(v int32) *ListCheckResultResponseBodyChecks {
	s.CheckSaleType = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetCheckShowName(v string) *ListCheckResultResponseBodyChecks {
	s.CheckShowName = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetCheckType(v string) *ListCheckResultResponseBodyChecks {
	s.CheckType = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetInstanceSubType(v string) *ListCheckResultResponseBodyChecks {
	s.InstanceSubType = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetInstanceType(v string) *ListCheckResultResponseBodyChecks {
	s.InstanceType = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetLastCheckTime(v int64) *ListCheckResultResponseBodyChecks {
	s.LastCheckTime = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetOperationType(v string) *ListCheckResultResponseBodyChecks {
	s.OperationType = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetRiskLevel(v string) *ListCheckResultResponseBodyChecks {
	s.RiskLevel = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetStatus(v string) *ListCheckResultResponseBodyChecks {
	s.Status = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetStatusMessage(v string) *ListCheckResultResponseBodyChecks {
	s.StatusMessage = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetTaskId(v string) *ListCheckResultResponseBodyChecks {
	s.TaskId = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetTrialPermission(v bool) *ListCheckResultResponseBodyChecks {
	s.TrialPermission = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetTrialPermissionType(v int32) *ListCheckResultResponseBodyChecks {
	s.TrialPermissionType = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetVendor(v string) *ListCheckResultResponseBodyChecks {
	s.Vendor = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) SetVendorShowName(v string) *ListCheckResultResponseBodyChecks {
	s.VendorShowName = &v
	return s
}

func (s *ListCheckResultResponseBodyChecks) Validate() error {
	if s.CheckPolicies != nil {
		for _, item := range s.CheckPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCheckResultResponseBodyChecksCheckPolicies struct {
	// The ID of the requirement item for the check item.
	//
	// example:
	//
	// 2
	RequirementId *int64 `json:"RequirementId,omitempty" xml:"RequirementId,omitempty"`
	// The display name of the requirement item for the check item.
	//
	// example:
	//
	// Alibaba cloud OSS best security practices
	RequirementShowName *string `json:"RequirementShowName,omitempty" xml:"RequirementShowName,omitempty"`
	// The ID of the section for the check item.
	//
	// example:
	//
	// 3
	SectionId *int64 `json:"SectionId,omitempty" xml:"SectionId,omitempty"`
	// The display name of the section for the check item.
	//
	// example:
	//
	// Log Audit
	SectionShowName *string `json:"SectionShowName,omitempty" xml:"SectionShowName,omitempty"`
	// The standard ID of the check item.
	//
	// example:
	//
	// 1
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
	// The standard display name of the check item.
	//
	// example:
	//
	// Best security practices
	StandardShowName *string `json:"StandardShowName,omitempty" xml:"StandardShowName,omitempty"`
}

func (s ListCheckResultResponseBodyChecksCheckPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListCheckResultResponseBodyChecksCheckPolicies) GoString() string {
	return s.String()
}

func (s *ListCheckResultResponseBodyChecksCheckPolicies) GetRequirementId() *int64 {
	return s.RequirementId
}

func (s *ListCheckResultResponseBodyChecksCheckPolicies) GetRequirementShowName() *string {
	return s.RequirementShowName
}

func (s *ListCheckResultResponseBodyChecksCheckPolicies) GetSectionId() *int64 {
	return s.SectionId
}

func (s *ListCheckResultResponseBodyChecksCheckPolicies) GetSectionShowName() *string {
	return s.SectionShowName
}

func (s *ListCheckResultResponseBodyChecksCheckPolicies) GetStandardId() *int64 {
	return s.StandardId
}

func (s *ListCheckResultResponseBodyChecksCheckPolicies) GetStandardShowName() *string {
	return s.StandardShowName
}

func (s *ListCheckResultResponseBodyChecksCheckPolicies) SetRequirementId(v int64) *ListCheckResultResponseBodyChecksCheckPolicies {
	s.RequirementId = &v
	return s
}

func (s *ListCheckResultResponseBodyChecksCheckPolicies) SetRequirementShowName(v string) *ListCheckResultResponseBodyChecksCheckPolicies {
	s.RequirementShowName = &v
	return s
}

func (s *ListCheckResultResponseBodyChecksCheckPolicies) SetSectionId(v int64) *ListCheckResultResponseBodyChecksCheckPolicies {
	s.SectionId = &v
	return s
}

func (s *ListCheckResultResponseBodyChecksCheckPolicies) SetSectionShowName(v string) *ListCheckResultResponseBodyChecksCheckPolicies {
	s.SectionShowName = &v
	return s
}

func (s *ListCheckResultResponseBodyChecksCheckPolicies) SetStandardId(v int64) *ListCheckResultResponseBodyChecksCheckPolicies {
	s.StandardId = &v
	return s
}

func (s *ListCheckResultResponseBodyChecksCheckPolicies) SetStandardShowName(v string) *ListCheckResultResponseBodyChecksCheckPolicies {
	s.StandardShowName = &v
	return s
}

func (s *ListCheckResultResponseBodyChecksCheckPolicies) Validate() error {
	return dara.Validate(s)
}

type ListCheckResultResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCheckResultResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCheckResultResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListCheckResultResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListCheckResultResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCheckResultResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckResultResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCheckResultResponseBodyPageInfo) SetCount(v int32) *ListCheckResultResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListCheckResultResponseBodyPageInfo) SetCurrentPage(v int32) *ListCheckResultResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCheckResultResponseBodyPageInfo) SetPageSize(v int32) *ListCheckResultResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCheckResultResponseBodyPageInfo) SetTotalCount(v int32) *ListCheckResultResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCheckResultResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
