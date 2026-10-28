// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterCheckResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChecks(v []*ListClusterCheckResultResponseBodyChecks) *ListClusterCheckResultResponseBody
	GetChecks() []*ListClusterCheckResultResponseBodyChecks
	SetPageInfo(v *ListClusterCheckResultResponseBodyPageInfo) *ListClusterCheckResultResponseBody
	GetPageInfo() *ListClusterCheckResultResponseBodyPageInfo
	SetRequestId(v string) *ListClusterCheckResultResponseBody
	GetRequestId() *string
}

type ListClusterCheckResultResponseBody struct {
	// The check result information.
	Checks []*ListClusterCheckResultResponseBodyChecks `json:"Checks,omitempty" xml:"Checks,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListClusterCheckResultResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447DDCAE1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClusterCheckResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterCheckResultResponseBody) GetChecks() []*ListClusterCheckResultResponseBodyChecks {
	return s.Checks
}

func (s *ListClusterCheckResultResponseBody) GetPageInfo() *ListClusterCheckResultResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListClusterCheckResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterCheckResultResponseBody) SetChecks(v []*ListClusterCheckResultResponseBodyChecks) *ListClusterCheckResultResponseBody {
	s.Checks = v
	return s
}

func (s *ListClusterCheckResultResponseBody) SetPageInfo(v *ListClusterCheckResultResponseBodyPageInfo) *ListClusterCheckResultResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListClusterCheckResultResponseBody) SetRequestId(v string) *ListClusterCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterCheckResultResponseBody) Validate() error {
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

type ListClusterCheckResultResponseBodyChecks struct {
	// The subtype of the cloud service.
	//
	// example:
	//
	// 0
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// The asset type.
	//
	// example:
	//
	// 0
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The ID of the check item.
	//
	// example:
	//
	// 5
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The standards, regulations, and sections associated with the check result.
	CheckPolicies []*ListClusterCheckResultResponseBodyChecksCheckPolicies `json:"CheckPolicies,omitempty" xml:"CheckPolicies,omitempty" type:"Repeated"`
	// The name of the check item.
	//
	// example:
	//
	// OSS-PublicReadOpenManifestFileWithoutEncryption
	CheckShowName *string `json:"CheckShowName,omitempty" xml:"CheckShowName,omitempty"`
	// The source type of the Threat Detection Service check item. Valid values:
	//
	//  - **CUSTOM**: Custom.
	//
	//  - **SYSTEM**: Predefined by the Threat Detection Service platform.
	//
	// example:
	//
	// SYSTEM
	CheckType *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The asset subtype of the cloud service. Valid values:
	//
	// - If **InstanceType*	- is set to **ECS**, valid values are:
	//
	//     - **INSTANCE**
	//
	//     - **DISK**
	//
	//     - **SECURITY_GROUP**
	//
	// - If **InstanceType*	- is set to **ACR**, valid values are:
	//
	//     - **REPOSITORY_ENTERPRISE**
	//
	//     - **REPOSITORY_PERSON**
	//
	// - If **InstanceType*	- is set to **RAM**, valid values are:
	//
	//     - **ALIAS**
	//
	//     - **USER**
	//
	//     - **POLICY**
	//
	//     - **GROUP**
	//
	// - If **InstanceType*	- is set to **WAF**, valid values are:
	//
	//     - **DOMAIN**
	//
	// - If **InstanceType*	- is set to other values, valid values are:
	//
	//     - **INSTANCE**
	//
	// example:
	//
	// DISK
	InstanceSubType *string `json:"InstanceSubType,omitempty" xml:"InstanceSubType,omitempty"`
	// The asset type of the cloud service.
	//
	// example:
	//
	// Workload
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The timestamp of the latest check. Unit: milliseconds.
	//
	// example:
	//
	// 1657793398000
	LastCheckTime *int64 `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
	// The risk level of the check item. Valid values:
	//
	// - **HIGH**: High.
	//
	// - **MEDIUM**: Medium.
	//
	// - **LOW**: Low.
	//
	// example:
	//
	// HIGH
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The status of the check item. Valid values:
	//
	// - **PASS**: Passed.
	//
	// - **NOT_PASS**: Failed.
	//
	// - **CHECKING**: Checking.
	//
	// - **NOT_CHECK**: Not checked.
	//
	// - **WHITELIST**: Whitelisted.
	//
	// example:
	//
	// PASS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Specifies whether the check depends on the TRIAL permission.
	//
	// example:
	//
	// true
	TrialPermission *bool `json:"TrialPermission,omitempty" xml:"TrialPermission,omitempty"`
	// Specifies whether the check item requires enabling data delivery for ActionTrail for more than 30 days to build a behavior baseline. Valid values:
	//
	// - **1**: Yes.
	//
	// - **0**: No.
	//
	// example:
	//
	// 1
	TrialPermissionType *int32 `json:"TrialPermissionType,omitempty" xml:"TrialPermissionType,omitempty"`
	// The vendor of the asset. Valid values:
	//
	// 0: Alibaba Cloud.
	//
	// 3: Other cloud.
	//
	// 4: Other cloud.
	//
	// 5: Other cloud.
	//
	// 7: Other cloud.
	//
	// example:
	//
	// 0
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListClusterCheckResultResponseBodyChecks) String() string {
	return dara.Prettify(s)
}

func (s ListClusterCheckResultResponseBodyChecks) GoString() string {
	return s.String()
}

func (s *ListClusterCheckResultResponseBodyChecks) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *ListClusterCheckResultResponseBodyChecks) GetAssetType() *int32 {
	return s.AssetType
}

func (s *ListClusterCheckResultResponseBodyChecks) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ListClusterCheckResultResponseBodyChecks) GetCheckPolicies() []*ListClusterCheckResultResponseBodyChecksCheckPolicies {
	return s.CheckPolicies
}

func (s *ListClusterCheckResultResponseBodyChecks) GetCheckShowName() *string {
	return s.CheckShowName
}

func (s *ListClusterCheckResultResponseBodyChecks) GetCheckType() *string {
	return s.CheckType
}

func (s *ListClusterCheckResultResponseBodyChecks) GetInstanceSubType() *string {
	return s.InstanceSubType
}

func (s *ListClusterCheckResultResponseBodyChecks) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListClusterCheckResultResponseBodyChecks) GetLastCheckTime() *int64 {
	return s.LastCheckTime
}

func (s *ListClusterCheckResultResponseBodyChecks) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListClusterCheckResultResponseBodyChecks) GetStatus() *string {
	return s.Status
}

func (s *ListClusterCheckResultResponseBodyChecks) GetTrialPermission() *bool {
	return s.TrialPermission
}

func (s *ListClusterCheckResultResponseBodyChecks) GetTrialPermissionType() *int32 {
	return s.TrialPermissionType
}

func (s *ListClusterCheckResultResponseBodyChecks) GetVendor() *string {
	return s.Vendor
}

func (s *ListClusterCheckResultResponseBodyChecks) SetAssetSubType(v int32) *ListClusterCheckResultResponseBodyChecks {
	s.AssetSubType = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecks) SetAssetType(v int32) *ListClusterCheckResultResponseBodyChecks {
	s.AssetType = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecks) SetCheckId(v int64) *ListClusterCheckResultResponseBodyChecks {
	s.CheckId = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecks) SetCheckPolicies(v []*ListClusterCheckResultResponseBodyChecksCheckPolicies) *ListClusterCheckResultResponseBodyChecks {
	s.CheckPolicies = v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecks) SetCheckShowName(v string) *ListClusterCheckResultResponseBodyChecks {
	s.CheckShowName = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecks) SetCheckType(v string) *ListClusterCheckResultResponseBodyChecks {
	s.CheckType = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecks) SetInstanceSubType(v string) *ListClusterCheckResultResponseBodyChecks {
	s.InstanceSubType = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecks) SetInstanceType(v string) *ListClusterCheckResultResponseBodyChecks {
	s.InstanceType = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecks) SetLastCheckTime(v int64) *ListClusterCheckResultResponseBodyChecks {
	s.LastCheckTime = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecks) SetRiskLevel(v string) *ListClusterCheckResultResponseBodyChecks {
	s.RiskLevel = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecks) SetStatus(v string) *ListClusterCheckResultResponseBodyChecks {
	s.Status = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecks) SetTrialPermission(v bool) *ListClusterCheckResultResponseBodyChecks {
	s.TrialPermission = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecks) SetTrialPermissionType(v int32) *ListClusterCheckResultResponseBodyChecks {
	s.TrialPermissionType = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecks) SetVendor(v string) *ListClusterCheckResultResponseBodyChecks {
	s.Vendor = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecks) Validate() error {
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

type ListClusterCheckResultResponseBodyChecksCheckPolicies struct {
	// The ID of the regulation for the check item.
	//
	// example:
	//
	// 2
	RequirementId *int64 `json:"RequirementId,omitempty" xml:"RequirementId,omitempty"`
	// The display name of the regulation for the check item.
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
	// The ID of the standard for the check item.
	//
	// example:
	//
	// 1
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
	// The display name of the standard for the check item.
	//
	// example:
	//
	// Best security practices
	StandardShowName *string `json:"StandardShowName,omitempty" xml:"StandardShowName,omitempty"`
}

func (s ListClusterCheckResultResponseBodyChecksCheckPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListClusterCheckResultResponseBodyChecksCheckPolicies) GoString() string {
	return s.String()
}

func (s *ListClusterCheckResultResponseBodyChecksCheckPolicies) GetRequirementId() *int64 {
	return s.RequirementId
}

func (s *ListClusterCheckResultResponseBodyChecksCheckPolicies) GetRequirementShowName() *string {
	return s.RequirementShowName
}

func (s *ListClusterCheckResultResponseBodyChecksCheckPolicies) GetSectionId() *int64 {
	return s.SectionId
}

func (s *ListClusterCheckResultResponseBodyChecksCheckPolicies) GetSectionShowName() *string {
	return s.SectionShowName
}

func (s *ListClusterCheckResultResponseBodyChecksCheckPolicies) GetStandardId() *int64 {
	return s.StandardId
}

func (s *ListClusterCheckResultResponseBodyChecksCheckPolicies) GetStandardShowName() *string {
	return s.StandardShowName
}

func (s *ListClusterCheckResultResponseBodyChecksCheckPolicies) SetRequirementId(v int64) *ListClusterCheckResultResponseBodyChecksCheckPolicies {
	s.RequirementId = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecksCheckPolicies) SetRequirementShowName(v string) *ListClusterCheckResultResponseBodyChecksCheckPolicies {
	s.RequirementShowName = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecksCheckPolicies) SetSectionId(v int64) *ListClusterCheckResultResponseBodyChecksCheckPolicies {
	s.SectionId = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecksCheckPolicies) SetSectionShowName(v string) *ListClusterCheckResultResponseBodyChecksCheckPolicies {
	s.SectionShowName = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecksCheckPolicies) SetStandardId(v int64) *ListClusterCheckResultResponseBodyChecksCheckPolicies {
	s.StandardId = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecksCheckPolicies) SetStandardShowName(v string) *ListClusterCheckResultResponseBodyChecksCheckPolicies {
	s.StandardShowName = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyChecksCheckPolicies) Validate() error {
	return dara.Validate(s)
}

type ListClusterCheckResultResponseBodyPageInfo struct {
	// The number of data entries displayed on the current page during paging in a paged query.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the current page during paging in a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of alerting events to display on each page during paging in a paged query. Default value: **20**, which indicates that 20 alerting events are displayed on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 83
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClusterCheckResultResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListClusterCheckResultResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListClusterCheckResultResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListClusterCheckResultResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListClusterCheckResultResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClusterCheckResultResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListClusterCheckResultResponseBodyPageInfo) SetCount(v int32) *ListClusterCheckResultResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyPageInfo) SetCurrentPage(v int32) *ListClusterCheckResultResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyPageInfo) SetPageSize(v int32) *ListClusterCheckResultResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyPageInfo) SetTotalCount(v int32) *ListClusterCheckResultResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListClusterCheckResultResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
