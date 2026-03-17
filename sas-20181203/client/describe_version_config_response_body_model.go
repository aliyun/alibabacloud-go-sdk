// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVersionConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentlessCapacity(v int64) *DescribeVersionConfigResponseBody
	GetAgentlessCapacity() *int64
	SetAllowPartialBuy(v int32) *DescribeVersionConfigResponseBody
	GetAllowPartialBuy() *int32
	SetAntiRansomwareCapacity(v int32) *DescribeVersionConfigResponseBody
	GetAntiRansomwareCapacity() *int32
	SetAntiRansomwareService(v int32) *DescribeVersionConfigResponseBody
	GetAntiRansomwareService() *int32
	SetAppWhiteList(v int32) *DescribeVersionConfigResponseBody
	GetAppWhiteList() *int32
	SetAppWhiteListAuthCount(v int64) *DescribeVersionConfigResponseBody
	GetAppWhiteListAuthCount() *int64
	SetAssetLevel(v int32) *DescribeVersionConfigResponseBody
	GetAssetLevel() *int32
	SetCanTryPostPaidPackage(v int32) *DescribeVersionConfigResponseBody
	GetCanTryPostPaidPackage() *int32
	SetCspmCapacity(v int64) *DescribeVersionConfigResponseBody
	GetCspmCapacity() *int64
	SetHighestVersion(v int32) *DescribeVersionConfigResponseBody
	GetHighestVersion() *int32
	SetHoneypotCapacity(v int64) *DescribeVersionConfigResponseBody
	GetHoneypotCapacity() *int64
	SetImageScanCapacity(v int64) *DescribeVersionConfigResponseBody
	GetImageScanCapacity() *int64
	SetInstanceBuyType(v int32) *DescribeVersionConfigResponseBody
	GetInstanceBuyType() *int32
	SetInstanceId(v string) *DescribeVersionConfigResponseBody
	GetInstanceId() *string
	SetIntelligentAnalysisFlow(v int32) *DescribeVersionConfigResponseBody
	GetIntelligentAnalysisFlow() *int32
	SetIsNewContainerVersion(v bool) *DescribeVersionConfigResponseBody
	GetIsNewContainerVersion() *bool
	SetIsNewMultiVersion(v bool) *DescribeVersionConfigResponseBody
	GetIsNewMultiVersion() *bool
	SetIsOverBalance(v bool) *DescribeVersionConfigResponseBody
	GetIsOverBalance() *bool
	SetIsPostpay(v bool) *DescribeVersionConfigResponseBody
	GetIsPostpay() *bool
	SetIsTrialVersion(v int32) *DescribeVersionConfigResponseBody
	GetIsTrialVersion() *int32
	SetLastTrailEndTime(v int64) *DescribeVersionConfigResponseBody
	GetLastTrailEndTime() *int64
	SetMVAuthCount(v int32) *DescribeVersionConfigResponseBody
	GetMVAuthCount() *int32
	SetMVUnusedAuthCount(v int32) *DescribeVersionConfigResponseBody
	GetMVUnusedAuthCount() *int32
	SetMergedVersion(v int32) *DescribeVersionConfigResponseBody
	GetMergedVersion() *int32
	SetMultiVersion(v string) *DescribeVersionConfigResponseBody
	GetMultiVersion() *string
	SetNewThreatAnalysis(v int32) *DescribeVersionConfigResponseBody
	GetNewThreatAnalysis() *int32
	SetOnboardedAssets(v int32) *DescribeVersionConfigResponseBody
	GetOnboardedAssets() *int32
	SetOpenTime(v int64) *DescribeVersionConfigResponseBody
	GetOpenTime() *int64
	SetPostPayHostVersion(v int32) *DescribeVersionConfigResponseBody
	GetPostPayHostVersion() *int32
	SetPostPayInstanceId(v string) *DescribeVersionConfigResponseBody
	GetPostPayInstanceId() *string
	SetPostPayModuleSwitch(v string) *DescribeVersionConfigResponseBody
	GetPostPayModuleSwitch() *string
	SetPostPayOpenTime(v int64) *DescribeVersionConfigResponseBody
	GetPostPayOpenTime() *int64
	SetPostPayStatus(v int32) *DescribeVersionConfigResponseBody
	GetPostPayStatus() *int32
	SetRaspCapacity(v int64) *DescribeVersionConfigResponseBody
	GetRaspCapacity() *int64
	SetReleaseTime(v int64) *DescribeVersionConfigResponseBody
	GetReleaseTime() *int64
	SetRequestId(v string) *DescribeVersionConfigResponseBody
	GetRequestId() *string
	SetSasLog(v int32) *DescribeVersionConfigResponseBody
	GetSasLog() *int32
	SetSasScreen(v int32) *DescribeVersionConfigResponseBody
	GetSasScreen() *int32
	SetSdkCapacity(v int64) *DescribeVersionConfigResponseBody
	GetSdkCapacity() *int64
	SetSlsCapacity(v int64) *DescribeVersionConfigResponseBody
	GetSlsCapacity() *int64
	SetThreatAnalysisCapacity(v int64) *DescribeVersionConfigResponseBody
	GetThreatAnalysisCapacity() *int64
	SetThreatAnalysisFlow(v int32) *DescribeVersionConfigResponseBody
	GetThreatAnalysisFlow() *int32
	SetUserDefinedAlarms(v int32) *DescribeVersionConfigResponseBody
	GetUserDefinedAlarms() *int32
	SetVersion(v int32) *DescribeVersionConfigResponseBody
	GetVersion() *int32
	SetVmCores(v int32) *DescribeVersionConfigResponseBody
	GetVmCores() *int32
	SetVulFixCapacity(v int64) *DescribeVersionConfigResponseBody
	GetVulFixCapacity() *int64
	SetWebLock(v int32) *DescribeVersionConfigResponseBody
	GetWebLock() *int32
	SetWebLockAuthCount(v int64) *DescribeVersionConfigResponseBody
	GetWebLockAuthCount() *int64
}

type DescribeVersionConfigResponseBody struct {
	// Number of agentless detections.
	//
	// >Agentless detection is not yet available for sale, so there\\"s no need to pay attention to this field at the moment.
	//
	// example:
	//
	// 10
	AgentlessCapacity *int64 `json:"AgentlessCapacity,omitempty" xml:"AgentlessCapacity,omitempty"`
	// Whether to allow pay-as-you-go purchases.
	//
	// - **0**: Not allowed
	//
	// - **1**: Allowed
	//
	// example:
	//
	// 1
	AllowPartialBuy *int32 `json:"AllowPartialBuy,omitempty" xml:"AllowPartialBuy,omitempty"`
	// Ransomware protection backup capacity, in GB.
	//
	// example:
	//
	// 160
	AntiRansomwareCapacity *int32 `json:"AntiRansomwareCapacity,omitempty" xml:"AntiRansomwareCapacity,omitempty"`
	// Ransomware Guardian Service. Values:
	//
	//  - **0**: Not activated
	//
	//  - **1**: Activated
	//
	// example:
	//
	// 1
	AntiRansomwareService *int32 `json:"AntiRansomwareService,omitempty" xml:"AntiRansomwareService,omitempty"`
	// Whether to enable the application whitelist. Values:
	//
	// - **0**: Not enabled
	//
	// - **2**: Enabled
	//
	// example:
	//
	// 2
	AppWhiteList *int32 `json:"AppWhiteList,omitempty" xml:"AppWhiteList,omitempty"`
	// Number of application whitelist authorizations.
	//
	// > One authorization allows the application of a whitelist policy to one server. After enabling the application whitelist function, the account will have 20 authorizations by default.
	//
	// example:
	//
	// 20
	AppWhiteListAuthCount *int64 `json:"AppWhiteListAuthCount,omitempty" xml:"AppWhiteListAuthCount,omitempty"`
	// Number of purchased server licenses.
	//
	// example:
	//
	// 30
	AssetLevel *int32 `json:"AssetLevel,omitempty" xml:"AssetLevel,omitempty"`
	// Whether it supports the activation of a post-paid trial package. Values:
	//
	// - **0**: Not supported
	//
	//  - **1**: Supported
	//
	// example:
	//
	// 1
	CanTryPostPaidPackage *int32 `json:"CanTryPostPaidPackage,omitempty" xml:"CanTryPostPaidPackage,omitempty"`
	// Purchased cloud platform configuration check scan count. Unit: times/month.
	//
	// example:
	//
	// 10
	CspmCapacity *int64 `json:"CspmCapacity,omitempty" xml:"CspmCapacity,omitempty"`
	// Purchase the highest version of the Security Center. Values:
	//
	//  - **1**: Free Edition
	//
	// - **3**: Enterprise Edition
	//
	// - **5**: Advanced Edition
	//
	// - **6**: Anti-Virus Edition
	//
	//  - **7**: Flagship Edition
	//
	// - **10**: Purchase Additional Services Only
	//
	// > When purchasing a single version, it indicates the corresponding version. When purchasing multiple versions, this value represents the highest version among the purchased multi-versions of Cloud Security Center.
	//
	// example:
	//
	// 1
	HighestVersion *int32 `json:"HighestVersion,omitempty" xml:"HighestVersion,omitempty"`
	// Number of purchased honeypot licenses.
	//
	// example:
	//
	// 20
	HoneypotCapacity *int64 `json:"HoneypotCapacity,omitempty" xml:"HoneypotCapacity,omitempty"`
	// Number of purchased image scanning authorizations.
	//
	// example:
	//
	// 8954
	ImageScanCapacity *int64 `json:"ImageScanCapacity,omitempty" xml:"ImageScanCapacity,omitempty"`
	// Instance purchase type. Values:
	//
	// - **0**: Self-purchased
	//
	//  - **1**: Allocated from multiple accounts
	//
	// example:
	//
	// 0
	InstanceBuyType *int32 `json:"InstanceBuyType,omitempty" xml:"InstanceBuyType,omitempty"`
	// ID of the purchased Cloud Security Center instance.
	//
	// example:
	//
	// sas-vg6hafdsafs****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// AI digital human analyzes traffic
	//
	// example:
	//
	// 100
	IntelligentAnalysisFlow *int32 `json:"IntelligentAnalysisFlow,omitempty" xml:"IntelligentAnalysisFlow,omitempty"`
	// Whether it is the new flagship version.
	//
	// - **true**: It is the latest version
	//
	// - **false**: It is not the latest version
	//
	// example:
	//
	// true
	IsNewContainerVersion *bool `json:"IsNewContainerVersion,omitempty" xml:"IsNewContainerVersion,omitempty"`
	// Whether it is the latest multi-version.
	//
	// - **true**: It is the latest multi-version
	//
	// - **false**: It is not the latest multi-version
	//
	// example:
	//
	// true
	IsNewMultiVersion *bool `json:"IsNewMultiVersion,omitempty" xml:"IsNewMultiVersion,omitempty"`
	// Whether the number of existing servers exceeds the maximum authorized purchase quantity. Values:
	//
	// - **false**: Not exceeded
	//
	// - **true**: Exceeded
	//
	// 	Notice: This parameter is deprecated, and you do not need to pay attention to it.
	//
	// example:
	//
	// false
	IsOverBalance *bool `json:"IsOverBalance,omitempty" xml:"IsOverBalance,omitempty"`
	// Whether to enable pay-as-you-go. Values:
	//
	// - **false**: Not enabled
	//
	// - **true**: Enabled
	//
	// example:
	//
	// true
	IsPostpay *bool `json:"IsPostpay,omitempty" xml:"IsPostpay,omitempty"`
	// Indicates whether the current Cloud Security Center version is a trial version. Values:
	//
	// - **0**: Not a trial version
	//
	// - **1**: Trial version
	//
	// example:
	//
	// 0
	IsTrialVersion *int32 `json:"IsTrialVersion,omitempty" xml:"IsTrialVersion,omitempty"`
	// The timestamp of the last trial expiration for Cloud Security Center, in milliseconds.
	//
	// example:
	//
	// 1603934844000
	LastTrailEndTime *int64 `json:"LastTrailEndTime,omitempty" xml:"LastTrailEndTime,omitempty"`
	// Total number of licenses when purchasing multiple versions.
	//
	// example:
	//
	// 5000
	MVAuthCount *int32 `json:"MVAuthCount,omitempty" xml:"MVAuthCount,omitempty"`
	// Total remaining licenses when purchasing multiple versions.
	//
	// example:
	//
	// 40
	MVUnusedAuthCount *int32 `json:"MVUnusedAuthCount,omitempty" xml:"MVUnusedAuthCount,omitempty"`
	// When both the annual/monthly and pay-as-you-go services for Cloud Security Center\\"s host and container security are activated, the higher protection version of the two is selected. Values:
	//
	// - **1**: Free Edition
	//
	//  - **6**: Anti-Virus Edition
	//
	// - **5**: Advanced Edition
	//
	// - **3**: Enterprise Edition
	//
	// - **7**: Ultimate Edition
	//
	// example:
	//
	// 7
	MergedVersion *int32 `json:"MergedVersion,omitempty" xml:"MergedVersion,omitempty"`
	// Usage of multiple version numbers and license counts
	//
	// example:
	//
	// null
	MultiVersion *string `json:"MultiVersion,omitempty" xml:"MultiVersion,omitempty"`
	// Whether to enable the new version of Threat Analysis and Response service. The new version of Threat Analysis and Response service refers to the one that supports purchasing access traffic and log storage capacity. Values:
	//
	// - **0**: No
	//
	// - **1**: Yes
	//
	// example:
	//
	// 1
	NewThreatAnalysis *int32 `json:"NewThreatAnalysis,omitempty" xml:"NewThreatAnalysis,omitempty"`
	// AI Digital Human Management Instance
	//
	// example:
	//
	// 10
	OnboardedAssets *int32 `json:"OnboardedAssets,omitempty" xml:"OnboardedAssets,omitempty"`
	// Service activation timestamp, unit: milliseconds.
	//
	// example:
	//
	// 1657244824669
	OpenTime *int64 `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	// When activating the pay-as-you-go service for host and container security, it represents the highest protection version of the already bound assets. Values:
	//
	// - **1**: Free Edition
	//
	//  - **3**: Enterprise Edition
	//
	//  - **5**: Advanced Edition
	//
	//  - **6**: Anti-Virus Edition
	//
	// - **7**: Flagship Edition
	//
	// example:
	//
	// 7
	PostPayHostVersion *int32 `json:"PostPayHostVersion,omitempty" xml:"PostPayHostVersion,omitempty"`
	// Pay-As-You-Go instance ID.
	//
	// example:
	//
	// postpay-sas-**
	PostPayInstanceId *string `json:"PostPayInstanceId,omitempty" xml:"PostPayInstanceId,omitempty"`
	// Pay-as-you-go module switch status, in the format of JsonString, with values as follows:
	//
	//  - Key:
	//
	//    	- **VUL**: Vulnerability Repair Module
	//
	//    	- **CSPM**: Cloud Security Posture Management Module
	//
	//    	- **AGENTLESS**: Agentless Detection Module
	//
	//    	- **SERVERLESS**: Serverless Security Module
	//
	//    	- **CTDR**: Threat Analysis and Response Module
	//
	//    	- **POST_HOST**: Host and Container Security Module
	//
	//    	- **SDK**: Malicious File Detection SDK Module
	//
	//    	- **RASP**: Application Protection Module
	//
	//  - Value: 0 indicates off, 1 indicates on
	//
	// example:
	//
	// {"VUL":1}
	PostPayModuleSwitch *string `json:"PostPayModuleSwitch,omitempty" xml:"PostPayModuleSwitch,omitempty"`
	// Pay-as-you-go activation time
	//
	// example:
	//
	// 1698915219000
	PostPayOpenTime *int64 `json:"PostPayOpenTime,omitempty" xml:"PostPayOpenTime,omitempty"`
	// Pay-As-You-Go instance status. Values:
	//
	// - **1**: Normal
	//
	// - **2**: Stopped due to unpaid bills
	//
	// example:
	//
	// 1
	PostPayStatus *int32 `json:"PostPayStatus,omitempty" xml:"PostPayStatus,omitempty"`
	// Number of purchased application protections. Unit: per month.
	//
	// example:
	//
	// 10
	RaspCapacity *int64 `json:"RaspCapacity,omitempty" xml:"RaspCapacity,omitempty"`
	// The timestamp of when the Cloud Security Center instance will expire, in milliseconds.
	//
	// > If you do not renew the service within 7 days after it expires, your paid instance will be downgraded to a free version, and you will no longer be able to use the features of the paid version. Your previous Cloud Security Center configuration data and historical alert data (e.g., DDoS alerts) will become inaccessible. At this point, you can only re-enable the paid version of Cloud Security Center by repurchasing it. For more information, see [Purchasing Cloud Security Center](https://help.aliyun.com/document_detail/42308.html).
	//
	// example:
	//
	// 1625846400000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The unique identifier generated by Alibaba Cloud for this request.
	//
	// example:
	//
	// C2DC96D2-DD2E-49D9-A28E-85590475DF55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether log analysis has been purchased. Values:
	//
	// - **0**: Not purchased
	//
	// - **1**: Purchased
	//
	// example:
	//
	// 1
	SasLog *int32 `json:"SasLog,omitempty" xml:"SasLog,omitempty"`
	// Whether the security dashboard has been purchased. Values:
	//
	// - **0**: Not purchased
	//
	// - **1**: Purchased
	//
	// example:
	//
	// 0
	SasScreen *int32 `json:"SasScreen,omitempty" xml:"SasScreen,omitempty"`
	// Number of SDK authorizations for malicious file detection
	//
	// example:
	//
	// 10
	SdkCapacity *int64 `json:"SdkCapacity,omitempty" xml:"SdkCapacity,omitempty"`
	// Purchased log storage capacity in GB. Range: 0 to 200000.
	//
	// example:
	//
	// 10240
	SlsCapacity *int64 `json:"SlsCapacity,omitempty" xml:"SlsCapacity,omitempty"`
	// Purchased threat analysis capacity. Unit: GB.
	//
	// example:
	//
	// 25
	ThreatAnalysisCapacity *int64 `json:"ThreatAnalysisCapacity,omitempty" xml:"ThreatAnalysisCapacity,omitempty"`
	// Purchased threat analysis and response log access traffic. Unit is GB/day.
	//
	// example:
	//
	// 10
	ThreatAnalysisFlow *int32 `json:"ThreatAnalysisFlow,omitempty" xml:"ThreatAnalysisFlow,omitempty"`
	// Whether to enable the custom alarm function. Values:
	//
	//  - **0**: Not enabled
	//
	// - **2**: Enabled
	//
	// example:
	//
	// 0
	UserDefinedAlarms *int32 `json:"UserDefinedAlarms,omitempty" xml:"UserDefinedAlarms,omitempty"`
	// Purchased Cloud Security Center version. Values:
	//
	// - **1**: Free Edition
	//
	//  - **3**: Enterprise Edition
	//
	// - **5**: Advanced Edition
	//
	// - **6**: Anti-Virus Edition
	//
	// - **7**: Flagship Edition
	//
	//  - **8**: Multi-Edition
	//
	//   - **10**: Value-Added Services Only
	//
	// example:
	//
	// 3
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// Number of authorized cores purchased.
	//
	// example:
	//
	// 10
	VmCores *int32 `json:"VmCores,omitempty" xml:"VmCores,omitempty"`
	// Number of purchased vulnerability fixes. Unit: times/month.
	//
	// example:
	//
	// 10
	VulFixCapacity *int64 `json:"VulFixCapacity,omitempty" xml:"VulFixCapacity,omitempty"`
	// Indicates whether the web tamper-proof service is enabled. Values:
	//
	// - **0**: Not enabled
	//
	// - **1**: Enabled
	//
	// example:
	//
	// 0
	WebLock *int32 `json:"WebLock,omitempty" xml:"WebLock,omitempty"`
	// The number of purchased web tamper-proof licenses. One license can enable web tamper protection for one server. Value range: 0~N.
	//
	//  >N is the number of servers you have.
	//
	// example:
	//
	// 0
	WebLockAuthCount *int64 `json:"WebLockAuthCount,omitempty" xml:"WebLockAuthCount,omitempty"`
}

func (s DescribeVersionConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVersionConfigResponseBody) GetAgentlessCapacity() *int64 {
	return s.AgentlessCapacity
}

func (s *DescribeVersionConfigResponseBody) GetAllowPartialBuy() *int32 {
	return s.AllowPartialBuy
}

func (s *DescribeVersionConfigResponseBody) GetAntiRansomwareCapacity() *int32 {
	return s.AntiRansomwareCapacity
}

func (s *DescribeVersionConfigResponseBody) GetAntiRansomwareService() *int32 {
	return s.AntiRansomwareService
}

func (s *DescribeVersionConfigResponseBody) GetAppWhiteList() *int32 {
	return s.AppWhiteList
}

func (s *DescribeVersionConfigResponseBody) GetAppWhiteListAuthCount() *int64 {
	return s.AppWhiteListAuthCount
}

func (s *DescribeVersionConfigResponseBody) GetAssetLevel() *int32 {
	return s.AssetLevel
}

func (s *DescribeVersionConfigResponseBody) GetCanTryPostPaidPackage() *int32 {
	return s.CanTryPostPaidPackage
}

func (s *DescribeVersionConfigResponseBody) GetCspmCapacity() *int64 {
	return s.CspmCapacity
}

func (s *DescribeVersionConfigResponseBody) GetHighestVersion() *int32 {
	return s.HighestVersion
}

func (s *DescribeVersionConfigResponseBody) GetHoneypotCapacity() *int64 {
	return s.HoneypotCapacity
}

func (s *DescribeVersionConfigResponseBody) GetImageScanCapacity() *int64 {
	return s.ImageScanCapacity
}

func (s *DescribeVersionConfigResponseBody) GetInstanceBuyType() *int32 {
	return s.InstanceBuyType
}

func (s *DescribeVersionConfigResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVersionConfigResponseBody) GetIntelligentAnalysisFlow() *int32 {
	return s.IntelligentAnalysisFlow
}

func (s *DescribeVersionConfigResponseBody) GetIsNewContainerVersion() *bool {
	return s.IsNewContainerVersion
}

func (s *DescribeVersionConfigResponseBody) GetIsNewMultiVersion() *bool {
	return s.IsNewMultiVersion
}

func (s *DescribeVersionConfigResponseBody) GetIsOverBalance() *bool {
	return s.IsOverBalance
}

func (s *DescribeVersionConfigResponseBody) GetIsPostpay() *bool {
	return s.IsPostpay
}

func (s *DescribeVersionConfigResponseBody) GetIsTrialVersion() *int32 {
	return s.IsTrialVersion
}

func (s *DescribeVersionConfigResponseBody) GetLastTrailEndTime() *int64 {
	return s.LastTrailEndTime
}

func (s *DescribeVersionConfigResponseBody) GetMVAuthCount() *int32 {
	return s.MVAuthCount
}

func (s *DescribeVersionConfigResponseBody) GetMVUnusedAuthCount() *int32 {
	return s.MVUnusedAuthCount
}

func (s *DescribeVersionConfigResponseBody) GetMergedVersion() *int32 {
	return s.MergedVersion
}

func (s *DescribeVersionConfigResponseBody) GetMultiVersion() *string {
	return s.MultiVersion
}

func (s *DescribeVersionConfigResponseBody) GetNewThreatAnalysis() *int32 {
	return s.NewThreatAnalysis
}

func (s *DescribeVersionConfigResponseBody) GetOnboardedAssets() *int32 {
	return s.OnboardedAssets
}

func (s *DescribeVersionConfigResponseBody) GetOpenTime() *int64 {
	return s.OpenTime
}

func (s *DescribeVersionConfigResponseBody) GetPostPayHostVersion() *int32 {
	return s.PostPayHostVersion
}

func (s *DescribeVersionConfigResponseBody) GetPostPayInstanceId() *string {
	return s.PostPayInstanceId
}

func (s *DescribeVersionConfigResponseBody) GetPostPayModuleSwitch() *string {
	return s.PostPayModuleSwitch
}

func (s *DescribeVersionConfigResponseBody) GetPostPayOpenTime() *int64 {
	return s.PostPayOpenTime
}

func (s *DescribeVersionConfigResponseBody) GetPostPayStatus() *int32 {
	return s.PostPayStatus
}

func (s *DescribeVersionConfigResponseBody) GetRaspCapacity() *int64 {
	return s.RaspCapacity
}

func (s *DescribeVersionConfigResponseBody) GetReleaseTime() *int64 {
	return s.ReleaseTime
}

func (s *DescribeVersionConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVersionConfigResponseBody) GetSasLog() *int32 {
	return s.SasLog
}

func (s *DescribeVersionConfigResponseBody) GetSasScreen() *int32 {
	return s.SasScreen
}

func (s *DescribeVersionConfigResponseBody) GetSdkCapacity() *int64 {
	return s.SdkCapacity
}

func (s *DescribeVersionConfigResponseBody) GetSlsCapacity() *int64 {
	return s.SlsCapacity
}

func (s *DescribeVersionConfigResponseBody) GetThreatAnalysisCapacity() *int64 {
	return s.ThreatAnalysisCapacity
}

func (s *DescribeVersionConfigResponseBody) GetThreatAnalysisFlow() *int32 {
	return s.ThreatAnalysisFlow
}

func (s *DescribeVersionConfigResponseBody) GetUserDefinedAlarms() *int32 {
	return s.UserDefinedAlarms
}

func (s *DescribeVersionConfigResponseBody) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeVersionConfigResponseBody) GetVmCores() *int32 {
	return s.VmCores
}

func (s *DescribeVersionConfigResponseBody) GetVulFixCapacity() *int64 {
	return s.VulFixCapacity
}

func (s *DescribeVersionConfigResponseBody) GetWebLock() *int32 {
	return s.WebLock
}

func (s *DescribeVersionConfigResponseBody) GetWebLockAuthCount() *int64 {
	return s.WebLockAuthCount
}

func (s *DescribeVersionConfigResponseBody) SetAgentlessCapacity(v int64) *DescribeVersionConfigResponseBody {
	s.AgentlessCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetAllowPartialBuy(v int32) *DescribeVersionConfigResponseBody {
	s.AllowPartialBuy = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetAntiRansomwareCapacity(v int32) *DescribeVersionConfigResponseBody {
	s.AntiRansomwareCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetAntiRansomwareService(v int32) *DescribeVersionConfigResponseBody {
	s.AntiRansomwareService = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetAppWhiteList(v int32) *DescribeVersionConfigResponseBody {
	s.AppWhiteList = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetAppWhiteListAuthCount(v int64) *DescribeVersionConfigResponseBody {
	s.AppWhiteListAuthCount = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetAssetLevel(v int32) *DescribeVersionConfigResponseBody {
	s.AssetLevel = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetCanTryPostPaidPackage(v int32) *DescribeVersionConfigResponseBody {
	s.CanTryPostPaidPackage = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetCspmCapacity(v int64) *DescribeVersionConfigResponseBody {
	s.CspmCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetHighestVersion(v int32) *DescribeVersionConfigResponseBody {
	s.HighestVersion = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetHoneypotCapacity(v int64) *DescribeVersionConfigResponseBody {
	s.HoneypotCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetImageScanCapacity(v int64) *DescribeVersionConfigResponseBody {
	s.ImageScanCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetInstanceBuyType(v int32) *DescribeVersionConfigResponseBody {
	s.InstanceBuyType = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetInstanceId(v string) *DescribeVersionConfigResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetIntelligentAnalysisFlow(v int32) *DescribeVersionConfigResponseBody {
	s.IntelligentAnalysisFlow = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetIsNewContainerVersion(v bool) *DescribeVersionConfigResponseBody {
	s.IsNewContainerVersion = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetIsNewMultiVersion(v bool) *DescribeVersionConfigResponseBody {
	s.IsNewMultiVersion = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetIsOverBalance(v bool) *DescribeVersionConfigResponseBody {
	s.IsOverBalance = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetIsPostpay(v bool) *DescribeVersionConfigResponseBody {
	s.IsPostpay = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetIsTrialVersion(v int32) *DescribeVersionConfigResponseBody {
	s.IsTrialVersion = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetLastTrailEndTime(v int64) *DescribeVersionConfigResponseBody {
	s.LastTrailEndTime = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetMVAuthCount(v int32) *DescribeVersionConfigResponseBody {
	s.MVAuthCount = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetMVUnusedAuthCount(v int32) *DescribeVersionConfigResponseBody {
	s.MVUnusedAuthCount = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetMergedVersion(v int32) *DescribeVersionConfigResponseBody {
	s.MergedVersion = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetMultiVersion(v string) *DescribeVersionConfigResponseBody {
	s.MultiVersion = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetNewThreatAnalysis(v int32) *DescribeVersionConfigResponseBody {
	s.NewThreatAnalysis = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetOnboardedAssets(v int32) *DescribeVersionConfigResponseBody {
	s.OnboardedAssets = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetOpenTime(v int64) *DescribeVersionConfigResponseBody {
	s.OpenTime = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetPostPayHostVersion(v int32) *DescribeVersionConfigResponseBody {
	s.PostPayHostVersion = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetPostPayInstanceId(v string) *DescribeVersionConfigResponseBody {
	s.PostPayInstanceId = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetPostPayModuleSwitch(v string) *DescribeVersionConfigResponseBody {
	s.PostPayModuleSwitch = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetPostPayOpenTime(v int64) *DescribeVersionConfigResponseBody {
	s.PostPayOpenTime = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetPostPayStatus(v int32) *DescribeVersionConfigResponseBody {
	s.PostPayStatus = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetRaspCapacity(v int64) *DescribeVersionConfigResponseBody {
	s.RaspCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetReleaseTime(v int64) *DescribeVersionConfigResponseBody {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetRequestId(v string) *DescribeVersionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetSasLog(v int32) *DescribeVersionConfigResponseBody {
	s.SasLog = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetSasScreen(v int32) *DescribeVersionConfigResponseBody {
	s.SasScreen = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetSdkCapacity(v int64) *DescribeVersionConfigResponseBody {
	s.SdkCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetSlsCapacity(v int64) *DescribeVersionConfigResponseBody {
	s.SlsCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetThreatAnalysisCapacity(v int64) *DescribeVersionConfigResponseBody {
	s.ThreatAnalysisCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetThreatAnalysisFlow(v int32) *DescribeVersionConfigResponseBody {
	s.ThreatAnalysisFlow = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetUserDefinedAlarms(v int32) *DescribeVersionConfigResponseBody {
	s.UserDefinedAlarms = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetVersion(v int32) *DescribeVersionConfigResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetVmCores(v int32) *DescribeVersionConfigResponseBody {
	s.VmCores = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetVulFixCapacity(v int64) *DescribeVersionConfigResponseBody {
	s.VulFixCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetWebLock(v int32) *DescribeVersionConfigResponseBody {
	s.WebLock = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetWebLockAuthCount(v int64) *DescribeVersionConfigResponseBody {
	s.WebLockAuthCount = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
