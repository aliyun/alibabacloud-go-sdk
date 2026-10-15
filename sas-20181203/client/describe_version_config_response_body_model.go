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
	SetBuySasEdr(v string) *DescribeVersionConfigResponseBody
	GetBuySasEdr() *string
	SetCanTryPostPaidPackage(v int32) *DescribeVersionConfigResponseBody
	GetCanTryPostPaidPackage() *int32
	SetCspmCapacity(v int64) *DescribeVersionConfigResponseBody
	GetCspmCapacity() *int64
	SetCspmInstanceCapacity(v int32) *DescribeVersionConfigResponseBody
	GetCspmInstanceCapacity() *int32
	SetHighestVersion(v int32) *DescribeVersionConfigResponseBody
	GetHighestVersion() *int32
	SetHoneypotCapacity(v int64) *DescribeVersionConfigResponseBody
	GetHoneypotCapacity() *int64
	SetHybridPaidGrayStatus(v string) *DescribeVersionConfigResponseBody
	GetHybridPaidGrayStatus() *string
	SetHybridPaidModuleSwitchMap(v int32) *DescribeVersionConfigResponseBody
	GetHybridPaidModuleSwitchMap() *int32
	SetHybridPaidStatus(v int32) *DescribeVersionConfigResponseBody
	GetHybridPaidStatus() *int32
	SetHybridSwitch(v int32) *DescribeVersionConfigResponseBody
	GetHybridSwitch() *int32
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
	SetNewPostPaidCspm(v int32) *DescribeVersionConfigResponseBody
	GetNewPostPaidCspm() *int32
	SetNewThreatAnalysis(v int32) *DescribeVersionConfigResponseBody
	GetNewThreatAnalysis() *int32
	SetOnboardedAssets(v int32) *DescribeVersionConfigResponseBody
	GetOnboardedAssets() *int32
	SetOpenTime(v int64) *DescribeVersionConfigResponseBody
	GetOpenTime() *int64
	SetPostPaidFirstTrialTime(v int64) *DescribeVersionConfigResponseBody
	GetPostPaidFirstTrialTime() *int64
	SetPostPaidLastTrialTime(v int64) *DescribeVersionConfigResponseBody
	GetPostPaidLastTrialTime() *int64
	SetPostPaidTrialExpiredTime(v int64) *DescribeVersionConfigResponseBody
	GetPostPaidTrialExpiredTime() *int64
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
	SetSasEdrClientAuthCount(v string) *DescribeVersionConfigResponseBody
	GetSasEdrClientAuthCount() *string
	SetSasEdrPostPaidInstanceId(v string) *DescribeVersionConfigResponseBody
	GetSasEdrPostPaidInstanceId() *string
	SetSasEdrPrePaidInstanceId(v string) *DescribeVersionConfigResponseBody
	GetSasEdrPrePaidInstanceId() *string
	SetSasEdrPrePaidInstanceStatus(v string) *DescribeVersionConfigResponseBody
	GetSasEdrPrePaidInstanceStatus() *string
	SetSasEdrVersion(v string) *DescribeVersionConfigResponseBody
	GetSasEdrVersion() *string
	SetSasLog(v int32) *DescribeVersionConfigResponseBody
	GetSasLog() *int32
	SetSasScreen(v int32) *DescribeVersionConfigResponseBody
	GetSasScreen() *int32
	SetSdkAiPostPaidGray(v int32) *DescribeVersionConfigResponseBody
	GetSdkAiPostPaidGray() *int32
	SetSdkCapacity(v int64) *DescribeVersionConfigResponseBody
	GetSdkCapacity() *int64
	SetSlsCapacity(v int64) *DescribeVersionConfigResponseBody
	GetSlsCapacity() *int64
	SetThreatAnalysisCapacity(v int64) *DescribeVersionConfigResponseBody
	GetThreatAnalysisCapacity() *int64
	SetThreatAnalysisFlow(v int32) *DescribeVersionConfigResponseBody
	GetThreatAnalysisFlow() *int32
	SetTrialModuleList(v []*DescribeVersionConfigResponseBodyTrialModuleList) *DescribeVersionConfigResponseBody
	GetTrialModuleList() []*DescribeVersionConfigResponseBodyTrialModuleList
	SetTrialVersion(v int32) *DescribeVersionConfigResponseBody
	GetTrialVersion() *int32
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
	// The agentless detection quota.
	//
	// > Agentless detection is not yet available for sale. You can ignore this field.
	//
	// example:
	//
	// 10
	AgentlessCapacity *int64 `json:"AgentlessCapacity,omitempty" xml:"AgentlessCapacity,omitempty"`
	// Indicates whether pay-as-you-go purchase is allowed. Valid values:
	//
	// - **0**: Not allowed.
	//
	// - **1**: Allowed.
	//
	// example:
	//
	// 1
	AllowPartialBuy *int32 `json:"AllowPartialBuy,omitempty" xml:"AllowPartialBuy,omitempty"`
	// The anti-ransomware backup capacity. Unit: GB.
	//
	// example:
	//
	// 160
	AntiRansomwareCapacity *int32 `json:"AntiRansomwareCapacity,omitempty" xml:"AntiRansomwareCapacity,omitempty"`
	// Indicates whether the anti-ransomware managed service is enabled. Valid values:
	//
	// - **0**: Not enabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	AntiRansomwareService *int32 `json:"AntiRansomwareService,omitempty" xml:"AntiRansomwareService,omitempty"`
	// Indicates whether the application whitelist feature is enabled. Valid values:
	//
	// - **0**: Not enabled.
	//
	// - **2**: Enabled.
	//
	// example:
	//
	// 2
	AppWhiteList *int32 `json:"AppWhiteList,omitempty" xml:"AppWhiteList,omitempty"`
	// The number of application whitelist authorizations.
	//
	// > One authorization allows you to apply a whitelist policy to one server. After you enable the application whitelist feature, your account is granted 20 authorizations by default.
	//
	// example:
	//
	// 20
	AppWhiteListAuthCount *int64 `json:"AppWhiteListAuthCount,omitempty" xml:"AppWhiteListAuthCount,omitempty"`
	// The number of purchased server authorizations.
	//
	// example:
	//
	// 30
	AssetLevel *int32 `json:"AssetLevel,omitempty" xml:"AssetLevel,omitempty"`
	// Indicates whether Endpoint Detection and Response (EDR) is purchased.
	//
	// example:
	//
	// true
	BuySasEdr *string `json:"BuySasEdr,omitempty" xml:"BuySasEdr,omitempty"`
	// Indicates whether the pay-as-you-go trial plan is supported. Valid values:
	//
	// - **0**: Not supported.
	//
	// - **1**: Supported.
	//
	// example:
	//
	// 1
	CanTryPostPaidPackage *int32 `json:"CanTryPostPaidPackage,omitempty" xml:"CanTryPostPaidPackage,omitempty"`
	// The number of purchased cloud platform configuration check scans. Unit: scans per month.
	//
	// example:
	//
	// 10
	CspmCapacity *int64 `json:"CspmCapacity,omitempty" xml:"CspmCapacity,omitempty"`
	// The AI digital human analysis traffic.
	//
	// example:
	//
	// 100
	CspmInstanceCapacity *int32 `json:"CspmInstanceCapacity,omitempty" xml:"CspmInstanceCapacity,omitempty"`
	// The highest purchased edition of Security Center. Valid values:
	//
	// - **1**: Free Edition
	//
	// - **3**: Enterprise Edition
	//
	// - **5**: Advanced Edition
	//
	// - **6**: Anti-virus Edition
	//
	// - **7**: Ultimate Edition
	//
	// - **10**: Value-added services only
	//
	// > If a single edition is purchased, this value indicates the corresponding edition. If multiple editions are purchased, this value indicates the highest edition among the purchased editions of Security Center.
	//
	// example:
	//
	// 1
	HighestVersion *int32 `json:"HighestVersion,omitempty" xml:"HighestVersion,omitempty"`
	// The number of purchased honeypot authorizations.
	//
	// example:
	//
	// 20
	HoneypotCapacity *int64 `json:"HoneypotCapacity,omitempty" xml:"HoneypotCapacity,omitempty"`
	// The canary release module for elastic billing.
	//
	// example:
	//
	// {"CSPM_INSTANCE":1}
	HybridPaidGrayStatus *string `json:"HybridPaidGrayStatus,omitempty" xml:"HybridPaidGrayStatus,omitempty"`
	// The AI digital human analysis traffic.
	//
	// example:
	//
	// 100
	HybridPaidModuleSwitchMap *int32 `json:"HybridPaidModuleSwitchMap,omitempty" xml:"HybridPaidModuleSwitchMap,omitempty"`
	// The elastic billing switch status.
	//
	// example:
	//
	// 1
	HybridPaidStatus *int32 `json:"HybridPaidStatus,omitempty" xml:"HybridPaidStatus,omitempty"`
	// The AI digital human analysis traffic.
	//
	// example:
	//
	// 100
	HybridSwitch *int32 `json:"HybridSwitch,omitempty" xml:"HybridSwitch,omitempty"`
	// The number of purchased image scan authorizations.
	//
	// example:
	//
	// 8954
	ImageScanCapacity *int64 `json:"ImageScanCapacity,omitempty" xml:"ImageScanCapacity,omitempty"`
	// The instance purchase type. Valid values:
	//
	// - **0**: Self-purchased.
	//
	// - **1**: Multi-account allocation.
	//
	// example:
	//
	// 0
	InstanceBuyType *int32 `json:"InstanceBuyType,omitempty" xml:"InstanceBuyType,omitempty"`
	// The instance ID of the purchased Security Center.
	//
	// example:
	//
	// sas-vg6hafdsafs****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The AI digital human analysis traffic.
	//
	// example:
	//
	// 100
	IntelligentAnalysisFlow *int32 `json:"IntelligentAnalysisFlow,omitempty" xml:"IntelligentAnalysisFlow,omitempty"`
	// Indicates whether the instance is the latest Ultimate Edition. Valid values:
	//
	// - **true**: The instance is the latest version.
	//
	// - **false**: The instance is not the latest version.
	//
	// example:
	//
	// true
	IsNewContainerVersion *bool `json:"IsNewContainerVersion,omitempty" xml:"IsNewContainerVersion,omitempty"`
	// Indicates whether the instance is the latest multi-edition. Valid values:
	//
	// - **true**: The instance is the latest multi-edition.
	//
	// - **false**: The instance is not the latest multi-edition.
	//
	// example:
	//
	// true
	IsNewMultiVersion *bool `json:"IsNewMultiVersion,omitempty" xml:"IsNewMultiVersion,omitempty"`
	// Indicates whether the number of existing servers exceeds the maximum number of purchased authorizations. Valid values:
	//
	// - **false**: Not exceeded.
	//
	// - **true**: Exceeded.
	//
	// 	Notice: **[Deprecated]*	- This parameter is deprecated. You can ignore it.</notice>
	//
	// example:
	//
	// false
	IsOverBalance *bool `json:"IsOverBalance,omitempty" xml:"IsOverBalance,omitempty"`
	// Indicates whether the pay-as-you-go billing method is enabled. Valid values:
	//
	// - **false**: Not enabled.
	//
	// - **true**: Enabled.
	//
	// example:
	//
	// true
	IsPostpay *bool `json:"IsPostpay,omitempty" xml:"IsPostpay,omitempty"`
	// Indicates whether the current Security Center edition is a trial edition. Valid values:
	//
	// - **0**: No, it is not a trial edition.
	//
	// - **1**: Yes, it is a trial edition.
	//
	// example:
	//
	// 0
	IsTrialVersion *int32 `json:"IsTrialVersion,omitempty" xml:"IsTrialVersion,omitempty"`
	// The end timestamp of the last Security Center trial. Unit: milliseconds.
	//
	// example:
	//
	// 1603934844000
	LastTrailEndTime *int64 `json:"LastTrailEndTime,omitempty" xml:"LastTrailEndTime,omitempty"`
	// The total number of authorizations for the multi-edition purchase.
	//
	// example:
	//
	// 5000
	MVAuthCount *int32 `json:"MVAuthCount,omitempty" xml:"MVAuthCount,omitempty"`
	// The total number of remaining authorizations for the multi-edition purchase.
	//
	// example:
	//
	// 40
	MVUnusedAuthCount *int32 `json:"MVUnusedAuthCount,omitempty" xml:"MVUnusedAuthCount,omitempty"`
	// The higher protection edition between the subscription and pay-as-you-go services when both are enabled for Security Center host and container security. Valid values:
	//
	// - **1**: Free Edition
	//
	// - **6**: Anti-virus Edition
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
	// The multi-edition version numbers and authorization usage.
	//
	// example:
	//
	// null
	MultiVersion *string `json:"MultiVersion,omitempty" xml:"MultiVersion,omitempty"`
	// The AI digital human analysis traffic.
	//
	// example:
	//
	// 100
	NewPostPaidCspm *int32 `json:"NewPostPaidCspm,omitempty" xml:"NewPostPaidCspm,omitempty"`
	// Indicates whether the new version of Cloud Threat Detection and Response (CTDR) is enabled. The new version of CTDR supports the purchase of inbound traffic and log storage capacity. Valid values:
	//
	// - **0**: No.
	//
	// - **1**: Yes.
	//
	// example:
	//
	// 1
	NewThreatAnalysis *int32 `json:"NewThreatAnalysis,omitempty" xml:"NewThreatAnalysis,omitempty"`
	// The AI digital human managed instances.
	//
	// example:
	//
	// 10
	OnboardedAssets *int32 `json:"OnboardedAssets,omitempty" xml:"OnboardedAssets,omitempty"`
	// The timestamp when the service is enabled. Unit: milliseconds.
	//
	// example:
	//
	// 1657244824669
	OpenTime *int64 `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	// The first trial time for the pay-as-you-go billing method.
	//
	// example:
	//
	// 1657244824669
	PostPaidFirstTrialTime *int64 `json:"PostPaidFirstTrialTime,omitempty" xml:"PostPaidFirstTrialTime,omitempty"`
	// The latest trial time for the pay-as-you-go billing method.
	//
	// example:
	//
	// 1657244824669
	PostPaidLastTrialTime *int64 `json:"PostPaidLastTrialTime,omitempty" xml:"PostPaidLastTrialTime,omitempty"`
	// The trial expiration time for the pay-as-you-go billing method.
	//
	// example:
	//
	// 1657244824669
	PostPaidTrialExpiredTime *int64 `json:"PostPaidTrialExpiredTime,omitempty" xml:"PostPaidTrialExpiredTime,omitempty"`
	// The highest protection edition for bound assets when the pay-as-you-go service for host and container security is enabled. Valid values:
	//
	// - **1**: Free Edition
	//
	// - **3**: Enterprise Edition
	//
	// - **5**: Advanced Edition
	//
	// - **6**: Anti-virus Edition
	//
	// - **7**: Ultimate Edition
	//
	// example:
	//
	// 7
	PostPayHostVersion *int32 `json:"PostPayHostVersion,omitempty" xml:"PostPayHostVersion,omitempty"`
	// The pay-as-you-go instance ID.
	//
	// example:
	//
	// postpay-sas-**
	PostPayInstanceId *string `json:"PostPayInstanceId,omitempty" xml:"PostPayInstanceId,omitempty"`
	// The switch status of pay-as-you-go modules. The value is a JSON string. Valid values:
	//
	// - Key:
	//
	//   - **VUL**: Vulnerability fixing module
	//
	//   - **CSPM**: Cloud Security Posture Management (CSPM) module
	//
	//   - **AGENTLESS**: Agentless detection module
	//
	//   - **SERVERLESS**: Serverless security module
	//
	//   - **CTDR**: Cloud Threat Detection and Response (CTDR) module
	//
	//   - **POST_HOST**: Host and container security module
	//
	//   - **SDK**: Malicious file detection SDK module
	//
	//   - **RASP**: Application protection module
	//
	// - Value: 0 indicates disabled, and 1 indicates enabled.
	//
	// example:
	//
	// {"VUL":1}
	PostPayModuleSwitch *string `json:"PostPayModuleSwitch,omitempty" xml:"PostPayModuleSwitch,omitempty"`
	// The time when the pay-as-you-go billing method is enabled. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1698915219000
	PostPayOpenTime *int64 `json:"PostPayOpenTime,omitempty" xml:"PostPayOpenTime,omitempty"`
	// The pay-as-you-go instance status. Valid values:
	//
	// - **1**: Normal.
	//
	// - **2**: Stopped due to overdue payment.
	//
	// example:
	//
	// 1
	PostPayStatus *int32 `json:"PostPayStatus,omitempty" xml:"PostPayStatus,omitempty"`
	// The number of purchased application protection authorizations. Unit: authorizations per month.
	//
	// example:
	//
	// 10
	RaspCapacity *int64 `json:"RaspCapacity,omitempty" xml:"RaspCapacity,omitempty"`
	// The expiration UNIX timestamp of the Security Center instance. Unit: milliseconds.
	//
	// > If you do not complete the renewal within 7 days after the instance expires, your paid instance is downgraded to the Free Edition. You cannot continue to use the features of the paid edition, and you cannot view your previous Security Center configuration data and historical alerting data, such as DDoS alerting. In this case, you must repurchase the paid edition of Security Center to enable the service. For more information, see [Purchase Security Center](https://help.aliyun.com/document_detail/42308.html).
	//
	// example:
	//
	// 1625846400000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The unique identifier generated by Alibaba Cloud for the request.
	//
	// example:
	//
	// C2DC96D2-DD2E-49D9-A28E-85590475DF55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of machines purchased for EDR.
	//
	// example:
	//
	// 10
	SasEdrClientAuthCount *string `json:"SasEdrClientAuthCount,omitempty" xml:"SasEdrClientAuthCount,omitempty"`
	// The pay-as-you-go EDR instance ID.
	//
	// example:
	//
	// sas-edr-postpaid-fadaf
	SasEdrPostPaidInstanceId *string `json:"SasEdrPostPaidInstanceId,omitempty" xml:"SasEdrPostPaidInstanceId,omitempty"`
	// The subscription EDR instance.
	//
	// example:
	//
	// sas-edr-sfkhakhk
	SasEdrPrePaidInstanceId *string `json:"SasEdrPrePaidInstanceId,omitempty" xml:"SasEdrPrePaidInstanceId,omitempty"`
	// The upfront EDR instance status.
	//
	// example:
	//
	// RELEASED
	SasEdrPrePaidInstanceStatus *string `json:"SasEdrPrePaidInstanceStatus,omitempty" xml:"SasEdrPrePaidInstanceStatus,omitempty"`
	// The purchased EDR edition.
	//
	// example:
	//
	// 1
	SasEdrVersion *string `json:"SasEdrVersion,omitempty" xml:"SasEdrVersion,omitempty"`
	// Indicates whether the log analysis feature is purchased. Valid values:
	//
	// - **0**: Not purchased.
	//
	// - **1**: Purchased.
	//
	// example:
	//
	// 1
	SasLog *int32 `json:"SasLog,omitempty" xml:"SasLog,omitempty"`
	// Indicates whether the security dashboard is purchased. Valid values:
	//
	// - **0**: Not purchased.
	//
	// - **1**: Purchased.
	//
	// example:
	//
	// 0
	SasScreen *int32 `json:"SasScreen,omitempty" xml:"SasScreen,omitempty"`
	// The canary release status for the pay-as-you-go SDK.
	//
	// example:
	//
	// 1
	SdkAiPostPaidGray *int32 `json:"SdkAiPostPaidGray,omitempty" xml:"SdkAiPostPaidGray,omitempty"`
	// The number of malicious file detection SDK authorizations.
	//
	// example:
	//
	// 10
	SdkCapacity *int64 `json:"SdkCapacity,omitempty" xml:"SdkCapacity,omitempty"`
	// The purchased log storage capacity. Unit: GB. Valid values: 0 to 200000.
	//
	// example:
	//
	// 10240
	SlsCapacity *int64 `json:"SlsCapacity,omitempty" xml:"SlsCapacity,omitempty"`
	// The purchased threat analysis capacity. Unit: GB.
	//
	// example:
	//
	// 25
	ThreatAnalysisCapacity *int64 `json:"ThreatAnalysisCapacity,omitempty" xml:"ThreatAnalysisCapacity,omitempty"`
	// The purchased log inbound traffic for Cloud Threat Detection and Response (CTDR). Unit: GB per day.
	//
	// example:
	//
	// 10
	ThreatAnalysisFlow *int32 `json:"ThreatAnalysisFlow,omitempty" xml:"ThreatAnalysisFlow,omitempty"`
	// The list of child trial modules.
	TrialModuleList []*DescribeVersionConfigResponseBodyTrialModuleList `json:"TrialModuleList,omitempty" xml:"TrialModuleList,omitempty" type:"Repeated"`
	// The trial edition.
	//
	// example:
	//
	// 1
	TrialVersion *int32 `json:"TrialVersion,omitempty" xml:"TrialVersion,omitempty"`
	// Indicates whether the custom alert feature is enabled. Valid values:
	//
	// - **0**: Not enabled.
	//
	// - **2**: Enabled.
	//
	// example:
	//
	// 0
	UserDefinedAlarms *int32 `json:"UserDefinedAlarms,omitempty" xml:"UserDefinedAlarms,omitempty"`
	// The purchased Security Center edition. Valid values:
	//
	// - **1**: Free Edition
	//
	// - **3**: Enterprise Edition
	//
	// - **5**: Advanced Edition
	//
	// - **6**: Anti-virus Edition
	//
	// - **7**: Ultimate Edition
	//
	// - **8**: Multi-edition
	//
	// - **10**: Value-added services only
	//
	// example:
	//
	// 3
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// The number of purchased authorized cores.
	//
	// example:
	//
	// 10
	VmCores *int32 `json:"VmCores,omitempty" xml:"VmCores,omitempty"`
	// The number of purchased vulnerability fixing quotas. Unit: fixes per month.
	//
	// example:
	//
	// 10
	VulFixCapacity *int64 `json:"VulFixCapacity,omitempty" xml:"VulFixCapacity,omitempty"`
	// Indicates whether the web tamper-proofing service is enabled. Valid values:
	//
	// - **0**: Not enabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 0
	WebLock *int32 `json:"WebLock,omitempty" xml:"WebLock,omitempty"`
	// The number of purchased web tamper-proofing authorizations. One authorization enables web tamper-proofing for one server. Valid values: 0 to N.
	//
	// > N is the number of servers that you own.
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

func (s *DescribeVersionConfigResponseBody) GetBuySasEdr() *string {
	return s.BuySasEdr
}

func (s *DescribeVersionConfigResponseBody) GetCanTryPostPaidPackage() *int32 {
	return s.CanTryPostPaidPackage
}

func (s *DescribeVersionConfigResponseBody) GetCspmCapacity() *int64 {
	return s.CspmCapacity
}

func (s *DescribeVersionConfigResponseBody) GetCspmInstanceCapacity() *int32 {
	return s.CspmInstanceCapacity
}

func (s *DescribeVersionConfigResponseBody) GetHighestVersion() *int32 {
	return s.HighestVersion
}

func (s *DescribeVersionConfigResponseBody) GetHoneypotCapacity() *int64 {
	return s.HoneypotCapacity
}

func (s *DescribeVersionConfigResponseBody) GetHybridPaidGrayStatus() *string {
	return s.HybridPaidGrayStatus
}

func (s *DescribeVersionConfigResponseBody) GetHybridPaidModuleSwitchMap() *int32 {
	return s.HybridPaidModuleSwitchMap
}

func (s *DescribeVersionConfigResponseBody) GetHybridPaidStatus() *int32 {
	return s.HybridPaidStatus
}

func (s *DescribeVersionConfigResponseBody) GetHybridSwitch() *int32 {
	return s.HybridSwitch
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

func (s *DescribeVersionConfigResponseBody) GetNewPostPaidCspm() *int32 {
	return s.NewPostPaidCspm
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

func (s *DescribeVersionConfigResponseBody) GetPostPaidFirstTrialTime() *int64 {
	return s.PostPaidFirstTrialTime
}

func (s *DescribeVersionConfigResponseBody) GetPostPaidLastTrialTime() *int64 {
	return s.PostPaidLastTrialTime
}

func (s *DescribeVersionConfigResponseBody) GetPostPaidTrialExpiredTime() *int64 {
	return s.PostPaidTrialExpiredTime
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

func (s *DescribeVersionConfigResponseBody) GetSasEdrClientAuthCount() *string {
	return s.SasEdrClientAuthCount
}

func (s *DescribeVersionConfigResponseBody) GetSasEdrPostPaidInstanceId() *string {
	return s.SasEdrPostPaidInstanceId
}

func (s *DescribeVersionConfigResponseBody) GetSasEdrPrePaidInstanceId() *string {
	return s.SasEdrPrePaidInstanceId
}

func (s *DescribeVersionConfigResponseBody) GetSasEdrPrePaidInstanceStatus() *string {
	return s.SasEdrPrePaidInstanceStatus
}

func (s *DescribeVersionConfigResponseBody) GetSasEdrVersion() *string {
	return s.SasEdrVersion
}

func (s *DescribeVersionConfigResponseBody) GetSasLog() *int32 {
	return s.SasLog
}

func (s *DescribeVersionConfigResponseBody) GetSasScreen() *int32 {
	return s.SasScreen
}

func (s *DescribeVersionConfigResponseBody) GetSdkAiPostPaidGray() *int32 {
	return s.SdkAiPostPaidGray
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

func (s *DescribeVersionConfigResponseBody) GetTrialModuleList() []*DescribeVersionConfigResponseBodyTrialModuleList {
	return s.TrialModuleList
}

func (s *DescribeVersionConfigResponseBody) GetTrialVersion() *int32 {
	return s.TrialVersion
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

func (s *DescribeVersionConfigResponseBody) SetBuySasEdr(v string) *DescribeVersionConfigResponseBody {
	s.BuySasEdr = &v
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

func (s *DescribeVersionConfigResponseBody) SetCspmInstanceCapacity(v int32) *DescribeVersionConfigResponseBody {
	s.CspmInstanceCapacity = &v
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

func (s *DescribeVersionConfigResponseBody) SetHybridPaidGrayStatus(v string) *DescribeVersionConfigResponseBody {
	s.HybridPaidGrayStatus = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetHybridPaidModuleSwitchMap(v int32) *DescribeVersionConfigResponseBody {
	s.HybridPaidModuleSwitchMap = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetHybridPaidStatus(v int32) *DescribeVersionConfigResponseBody {
	s.HybridPaidStatus = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetHybridSwitch(v int32) *DescribeVersionConfigResponseBody {
	s.HybridSwitch = &v
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

func (s *DescribeVersionConfigResponseBody) SetNewPostPaidCspm(v int32) *DescribeVersionConfigResponseBody {
	s.NewPostPaidCspm = &v
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

func (s *DescribeVersionConfigResponseBody) SetPostPaidFirstTrialTime(v int64) *DescribeVersionConfigResponseBody {
	s.PostPaidFirstTrialTime = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetPostPaidLastTrialTime(v int64) *DescribeVersionConfigResponseBody {
	s.PostPaidLastTrialTime = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetPostPaidTrialExpiredTime(v int64) *DescribeVersionConfigResponseBody {
	s.PostPaidTrialExpiredTime = &v
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

func (s *DescribeVersionConfigResponseBody) SetSasEdrClientAuthCount(v string) *DescribeVersionConfigResponseBody {
	s.SasEdrClientAuthCount = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetSasEdrPostPaidInstanceId(v string) *DescribeVersionConfigResponseBody {
	s.SasEdrPostPaidInstanceId = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetSasEdrPrePaidInstanceId(v string) *DescribeVersionConfigResponseBody {
	s.SasEdrPrePaidInstanceId = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetSasEdrPrePaidInstanceStatus(v string) *DescribeVersionConfigResponseBody {
	s.SasEdrPrePaidInstanceStatus = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetSasEdrVersion(v string) *DescribeVersionConfigResponseBody {
	s.SasEdrVersion = &v
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

func (s *DescribeVersionConfigResponseBody) SetSdkAiPostPaidGray(v int32) *DescribeVersionConfigResponseBody {
	s.SdkAiPostPaidGray = &v
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

func (s *DescribeVersionConfigResponseBody) SetTrialModuleList(v []*DescribeVersionConfigResponseBodyTrialModuleList) *DescribeVersionConfigResponseBody {
	s.TrialModuleList = v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetTrialVersion(v int32) *DescribeVersionConfigResponseBody {
	s.TrialVersion = &v
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
	if s.TrialModuleList != nil {
		for _, item := range s.TrialModuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVersionConfigResponseBodyTrialModuleList struct {
	// The name of the child trial module.
	//
	// example:
	//
	// EDR
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeVersionConfigResponseBodyTrialModuleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionConfigResponseBodyTrialModuleList) GoString() string {
	return s.String()
}

func (s *DescribeVersionConfigResponseBodyTrialModuleList) GetName() *string {
	return s.Name
}

func (s *DescribeVersionConfigResponseBodyTrialModuleList) SetName(v string) *DescribeVersionConfigResponseBodyTrialModuleList {
	s.Name = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyTrialModuleList) Validate() error {
	return dara.Validate(s)
}
