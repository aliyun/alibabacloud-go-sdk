// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVersionConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeVersionConfigResponseBody
	GetCode() *string
	SetData(v *DescribeVersionConfigResponseBodyData) *DescribeVersionConfigResponseBody
	GetData() *DescribeVersionConfigResponseBodyData
	SetMessage(v string) *DescribeVersionConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeVersionConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeVersionConfigResponseBody
	GetSuccess() *bool
}

type DescribeVersionConfigResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeVersionConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message. A success message is returned if the request succeeds. An error message is returned if the request fails.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6D462855-7835-5F91-835E-A62E44EC01CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values: true: The operation is successful. false: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeVersionConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVersionConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeVersionConfigResponseBody) GetData() *DescribeVersionConfigResponseBodyData {
	return s.Data
}

func (s *DescribeVersionConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeVersionConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVersionConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeVersionConfigResponseBody) SetCode(v string) *DescribeVersionConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetData(v *DescribeVersionConfigResponseBodyData) *DescribeVersionConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetMessage(v string) *DescribeVersionConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetRequestId(v string) *DescribeVersionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetSuccess(v bool) *DescribeVersionConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVersionConfigResponseBodyData struct {
	// The message body content encoded by the Base64 algorithm.
	Body *DescribeVersionConfigResponseBodyDataBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
}

func (s DescribeVersionConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeVersionConfigResponseBodyData) GetBody() *DescribeVersionConfigResponseBodyDataBody {
	return s.Body
}

func (s *DescribeVersionConfigResponseBodyData) SetBody(v *DescribeVersionConfigResponseBodyDataBody) *DescribeVersionConfigResponseBodyData {
	s.Body = v
	return s
}

func (s *DescribeVersionConfigResponseBodyData) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVersionConfigResponseBodyDataBody struct {
	// The agentless detection quota.
	//
	// >Agentless detection is not available for purchase. You can ignore this field.
	//
	// example:
	//
	// 10
	AgentlessCapacity *int64 `json:"AgentlessCapacity,omitempty" xml:"AgentlessCapacity,omitempty"`
	// Indicates whether pay-as-you-go purchase is allowed.
	//
	// - **0**: Not allowed.
	//
	// - **1**: Allowed.
	//
	// example:
	//
	// 1
	AllowPartialBuy *int32 `json:"AllowPartialBuy,omitempty" xml:"AllowPartialBuy,omitempty"`
	// The allocated anti-ransomware capacity. Unit: GB.
	//
	// example:
	//
	// 1680
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
	// 0
	AppWhiteList *int32 `json:"AppWhiteList,omitempty" xml:"AppWhiteList,omitempty"`
	// The number of application whitelist authorizations.
	//
	// > One authorization allows you to apply an application whitelist policy to one server. After the application whitelist feature is enabled, the account has 20 authorizations by default.
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
	// Indicates whether the pay-as-you-go trial plan can be activated. Valid values:
	//
	// - **0**: Not supported.
	//
	// - **1**: Supported.
	//
	// example:
	//
	// 0
	CanTryPostPaidPackage *int32 `json:"CanTryPostPaidPackage,omitempty" xml:"CanTryPostPaidPackage,omitempty"`
	// The allocated Cloud Security Posture Management (CSPM) scan quota. Unit: times/month.
	//
	// example:
	//
	// 10
	CspmCapacity *int64 `json:"CspmCapacity,omitempty" xml:"CspmCapacity,omitempty"`
	// The highest purchased edition of Security Center. Valid values:
	//
	// - **1**: Free Edition.
	//
	// - **3**: Enterprise Edition.
	//
	// - **5**: Premium Edition.
	//
	// - **6**: Anti-virus Edition.
	//
	// - **7**: Ultimate Edition.
	//
	// - **10**: Value-added services only.
	//
	// > If a single edition is purchased, this value indicates the corresponding edition. If multiple editions are purchased, this value indicates the highest edition among them.
	//
	// example:
	//
	// 1
	HighestVersion *int32 `json:"HighestVersion,omitempty" xml:"HighestVersion,omitempty"`
	// The allocated number of honeypot authorizations.
	//
	// example:
	//
	// 0
	HoneypotCapacity *int64 `json:"HoneypotCapacity,omitempty" xml:"HoneypotCapacity,omitempty"`
	// The number of purchased image scan authorizations.
	//
	// example:
	//
	// 1900
	ImageScanCapacity *int64 `json:"ImageScanCapacity,omitempty" xml:"ImageScanCapacity,omitempty"`
	// The instance purchase type. Valid values:
	//
	// - **0**: Self-purchased.
	//
	// - **1**: Allocated by multi-account management.
	//
	// example:
	//
	// 1
	InstanceBuyType *int32 `json:"InstanceBuyType,omitempty" xml:"InstanceBuyType,omitempty"`
	// The AI digital human analysis traffic.
	//
	// example:
	//
	// 100
	IntelligentAnalysisFlow *int32 `json:"IntelligentAnalysisFlow,omitempty" xml:"IntelligentAnalysisFlow,omitempty"`
	// Indicates whether the instance is the new Ultimate Edition.
	//
	// - **true**: The instance is the latest edition.
	//
	// - **false**: The instance is not the latest edition.
	//
	// example:
	//
	// true
	IsNewContainerVersion *bool `json:"IsNewContainerVersion,omitempty" xml:"IsNewContainerVersion,omitempty"`
	// Indicates whether the instance is the new multi-edition version.
	//
	// - **true**: The instance is the latest multi-edition version.
	//
	// - **false**: The instance is not the latest multi-edition version.
	//
	// example:
	//
	// true
	IsNewMultiVersion *bool `json:"IsNewMultiVersion,omitempty" xml:"IsNewMultiVersion,omitempty"`
	// Indicates whether the number of existing servers exceeds the maximum purchased authorization quota. Valid values:
	//
	// - **false**: Not exceeded.
	//
	// - **true**: Exceeded.
	//
	// 	Notice: This parameter is deprecated. You can ignore it.
	//
	// example:
	//
	// false
	IsOverBalance *bool `json:"IsOverBalance,omitempty" xml:"IsOverBalance,omitempty"`
	// Indicates whether pay-as-you-go billing is enabled. Valid values:
	//
	// - **false**: Not enabled.
	//
	// - **true**: Enabled.
	//
	// example:
	//
	// true
	IsPostpay *bool `json:"IsPostpay,omitempty" xml:"IsPostpay,omitempty"`
	// Indicates whether the current Security Center edition is a trial version. Valid values:
	//
	// - **0**: Not a trial version.
	//
	// - **1**: A trial version.
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
	// The higher protection edition when both subscription and pay-as-you-go host and container security services are enabled. Valid values:
	//
	// - **1**: Free Edition.
	//
	// - **6**: Anti-virus Edition.
	//
	// - **5**: Premium Edition.
	//
	// - **3**: Enterprise Edition.
	//
	// - **7**: Ultimate Edition.
	//
	// example:
	//
	// 1
	MergedVersion *int32 `json:"MergedVersion,omitempty" xml:"MergedVersion,omitempty"`
	// The multi-edition version numbers and authorization usage.
	//
	// example:
	//
	// null
	MultiVersion *string `json:"MultiVersion,omitempty" xml:"MultiVersion,omitempty"`
	// The total number of authorizations when multiple editions are purchased.
	//
	// example:
	//
	// 0
	MvAuthCount *int32 `json:"MvAuthCount,omitempty" xml:"MvAuthCount,omitempty"`
	// The total number of remaining authorizations when multiple editions are purchased.
	//
	// example:
	//
	// 0
	MvUnusedAuthCount *int32 `json:"MvUnusedAuthCount,omitempty" xml:"MvUnusedAuthCount,omitempty"`
	// Indicates whether the new version of Cloud Threat Detection and Response (CTDR) is enabled. The new version supports purchasing access traffic and log storage capacity for Cloud Threat Detection and Response (CTDR). Valid values:
	//
	// - **0**: No.
	//
	// - **1**: Yes.
	//
	// example:
	//
	// 0
	NewThreatAnalysis *int32 `json:"NewThreatAnalysis,omitempty" xml:"NewThreatAnalysis,omitempty"`
	// The AI digital human managed instances.
	//
	// example:
	//
	// 0
	OnboardedAssets *int32 `json:"OnboardedAssets,omitempty" xml:"OnboardedAssets,omitempty"`
	// The timestamp when the service was activated. Unit: milliseconds.
	//
	// example:
	//
	// 1657244824669
	OpenTime *int64 `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	// The highest protection edition for bound assets when the pay-as-you-go host and container security service is enabled. Valid values:
	//
	// - **1**: Free Edition.
	//
	// - **3**: Enterprise Edition.
	//
	// - **5**: Premium Edition.
	//
	// - **6**: Anti-virus Edition.
	//
	// - **7**: Ultimate Edition.
	//
	// example:
	//
	// 1
	PostPayHostVersion *int32 `json:"PostPayHostVersion,omitempty" xml:"PostPayHostVersion,omitempty"`
	// The pay-as-you-go instance ID.
	//
	// example:
	//
	// postpay-sas-frme8vjfiw2j
	PostPayInstanceId *string `json:"PostPayInstanceId,omitempty" xml:"PostPayInstanceId,omitempty"`
	// The switch status of pay-as-you-go modules in JSON string format. Valid values:
	//
	// - Key:
	//
	//   - **VUL**: Vulnerability fix module.
	//
	//   - **CSPM**: Cloud Security Posture Management module.
	//
	//   - **AGENTLESS**: Agentless detection module.
	//
	//   - **SERVERLESS**: Serverless security module.
	//
	//   - **CTDR**: Threat detection and response module.
	//
	//   - **POST_HOST**: Host and container security module.
	//
	//   - **SDK**: Malicious file detection SDK module.
	//
	//   - **RASP**: Application protection module.
	//
	// - Value: 0 indicates disabled. 1 indicates enabled.
	//
	// example:
	//
	// {\\"BASIC_SERVICE\\":0,\\"VUL\\":0}
	PostPayModuleSwitch *string `json:"PostPayModuleSwitch,omitempty" xml:"PostPayModuleSwitch,omitempty"`
	// The time when pay-as-you-go billing was activated.
	//
	// example:
	//
	// 1698915219000
	PostPayOpenTime *int64 `json:"PostPayOpenTime,omitempty" xml:"PostPayOpenTime,omitempty"`
	// The instance status of the pay-as-you-go instance. Valid values:
	//
	// - **1**: Normal.
	//
	// - **2**: Suspended due to overdue payment.
	//
	// example:
	//
	// 1
	PostPayStatus *int32 `json:"PostPayStatus,omitempty" xml:"PostPayStatus,omitempty"`
	// The number of purchased application protection quotas. Unit: count/month.
	//
	// example:
	//
	// 7
	RaspCapacity *int64 `json:"RaspCapacity,omitempty" xml:"RaspCapacity,omitempty"`
	// The UNIX timestamp that indicates when the Security Center instance expires. Unit: milliseconds.
	//
	// > If you do not perform renewal within 7 days after the instance expires, your paid edition instance is downgraded to Free Edition. You can no longer use the features of the paid edition, and your previous Security Center configuration data and historical alerting data (such as DDoS alerts) become inaccessible. In this case, you must repurchase Security Center to re-enable the paid edition. For more information, refer to the Security Center purchase documentation.
	//
	// example:
	//
	// 1625846400000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The request ID of Security Center.
	//
	// example:
	//
	// A6FB9AC3-4431-538F-BA8A-2A13AEA208A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether log analysis has been purchased. Valid values:
	//
	// - **0**: Not purchased.
	//
	// - **1**: Purchased.
	//
	// example:
	//
	// 0
	SasLog *int32 `json:"SasLog,omitempty" xml:"SasLog,omitempty"`
	// Indicates whether the security dashboard has been purchased. Valid values:
	//
	// - **0**: Not purchased.
	//
	// - **1**: Purchased.
	//
	// example:
	//
	// 0
	SasScreen *int32 `json:"SasScreen,omitempty" xml:"SasScreen,omitempty"`
	// The number of malicious file detection SDK authorizations.
	//
	// example:
	//
	// 10
	SdkCapacity *int64 `json:"SdkCapacity,omitempty" xml:"SdkCapacity,omitempty"`
	// The purchased log storage capacity. Unit: GB. Value range: 0 to 200000.
	//
	// example:
	//
	// 200
	SlsCapacity *int64 `json:"SlsCapacity,omitempty" xml:"SlsCapacity,omitempty"`
	// The purchased threat analysis capacity. Unit: GB.
	//
	// example:
	//
	// 10
	ThreatAnalysisCapacity *int64 `json:"ThreatAnalysisCapacity,omitempty" xml:"ThreatAnalysisCapacity,omitempty"`
	// The purchased threat detection and response log access traffic. Unit: GB/day.
	//
	// example:
	//
	// 10
	ThreatAnalysisFlow *int32 `json:"ThreatAnalysisFlow,omitempty" xml:"ThreatAnalysisFlow,omitempty"`
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
	// The purchased edition of Security Center. Valid values:
	//
	// - **1**: Free Edition.
	//
	// - **3**: Enterprise Edition.
	//
	// - **5**: Premium Edition.
	//
	// - **6**: Anti-virus Edition.
	//
	// - **7**: Ultimate Edition.
	//
	// - **8**: Multi-edition.
	//
	// - **10**: Value-added services only.
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
	// The number of purchased vulnerability fix quotas. Unit: times/month.
	//
	// example:
	//
	// 8
	VulFixCapacity *int64 `json:"VulFixCapacity,omitempty" xml:"VulFixCapacity,omitempty"`
	// Indicates whether the web tamper proofing service is enabled. Valid values:
	//
	// - **0**: Not enabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 0
	WebLock *int32 `json:"WebLock,omitempty" xml:"WebLock,omitempty"`
	// The number of purchased web tamper proofing authorizations. One authorization enables web tamper proofing protection for one server. Value range: 0 to N.
	//
	// > N is the number of servers that you own.
	//
	// example:
	//
	// 0
	WebLockAuthCount *int64 `json:"WebLockAuthCount,omitempty" xml:"WebLockAuthCount,omitempty"`
}

func (s DescribeVersionConfigResponseBodyDataBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionConfigResponseBodyDataBody) GoString() string {
	return s.String()
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetAgentlessCapacity() *int64 {
	return s.AgentlessCapacity
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetAllowPartialBuy() *int32 {
	return s.AllowPartialBuy
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetAntiRansomwareCapacity() *int32 {
	return s.AntiRansomwareCapacity
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetAntiRansomwareService() *int32 {
	return s.AntiRansomwareService
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetAppWhiteList() *int32 {
	return s.AppWhiteList
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetAppWhiteListAuthCount() *int64 {
	return s.AppWhiteListAuthCount
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetAssetLevel() *int32 {
	return s.AssetLevel
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetCanTryPostPaidPackage() *int32 {
	return s.CanTryPostPaidPackage
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetCspmCapacity() *int64 {
	return s.CspmCapacity
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetHighestVersion() *int32 {
	return s.HighestVersion
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetHoneypotCapacity() *int64 {
	return s.HoneypotCapacity
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetImageScanCapacity() *int64 {
	return s.ImageScanCapacity
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetInstanceBuyType() *int32 {
	return s.InstanceBuyType
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetIntelligentAnalysisFlow() *int32 {
	return s.IntelligentAnalysisFlow
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetIsNewContainerVersion() *bool {
	return s.IsNewContainerVersion
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetIsNewMultiVersion() *bool {
	return s.IsNewMultiVersion
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetIsOverBalance() *bool {
	return s.IsOverBalance
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetIsPostpay() *bool {
	return s.IsPostpay
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetIsTrialVersion() *int32 {
	return s.IsTrialVersion
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetLastTrailEndTime() *int64 {
	return s.LastTrailEndTime
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetMergedVersion() *int32 {
	return s.MergedVersion
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetMultiVersion() *string {
	return s.MultiVersion
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetMvAuthCount() *int32 {
	return s.MvAuthCount
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetMvUnusedAuthCount() *int32 {
	return s.MvUnusedAuthCount
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetNewThreatAnalysis() *int32 {
	return s.NewThreatAnalysis
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetOnboardedAssets() *int32 {
	return s.OnboardedAssets
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetOpenTime() *int64 {
	return s.OpenTime
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetPostPayHostVersion() *int32 {
	return s.PostPayHostVersion
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetPostPayInstanceId() *string {
	return s.PostPayInstanceId
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetPostPayModuleSwitch() *string {
	return s.PostPayModuleSwitch
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetPostPayOpenTime() *int64 {
	return s.PostPayOpenTime
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetPostPayStatus() *int32 {
	return s.PostPayStatus
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetRaspCapacity() *int64 {
	return s.RaspCapacity
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetReleaseTime() *int64 {
	return s.ReleaseTime
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetSasLog() *int32 {
	return s.SasLog
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetSasScreen() *int32 {
	return s.SasScreen
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetSdkCapacity() *int64 {
	return s.SdkCapacity
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetSlsCapacity() *int64 {
	return s.SlsCapacity
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetThreatAnalysisCapacity() *int64 {
	return s.ThreatAnalysisCapacity
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetThreatAnalysisFlow() *int32 {
	return s.ThreatAnalysisFlow
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetUserDefinedAlarms() *int32 {
	return s.UserDefinedAlarms
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetVmCores() *int32 {
	return s.VmCores
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetVulFixCapacity() *int64 {
	return s.VulFixCapacity
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetWebLock() *int32 {
	return s.WebLock
}

func (s *DescribeVersionConfigResponseBodyDataBody) GetWebLockAuthCount() *int64 {
	return s.WebLockAuthCount
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetAgentlessCapacity(v int64) *DescribeVersionConfigResponseBodyDataBody {
	s.AgentlessCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetAllowPartialBuy(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.AllowPartialBuy = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetAntiRansomwareCapacity(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.AntiRansomwareCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetAntiRansomwareService(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.AntiRansomwareService = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetAppWhiteList(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.AppWhiteList = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetAppWhiteListAuthCount(v int64) *DescribeVersionConfigResponseBodyDataBody {
	s.AppWhiteListAuthCount = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetAssetLevel(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.AssetLevel = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetCanTryPostPaidPackage(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.CanTryPostPaidPackage = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetCspmCapacity(v int64) *DescribeVersionConfigResponseBodyDataBody {
	s.CspmCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetHighestVersion(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.HighestVersion = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetHoneypotCapacity(v int64) *DescribeVersionConfigResponseBodyDataBody {
	s.HoneypotCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetImageScanCapacity(v int64) *DescribeVersionConfigResponseBodyDataBody {
	s.ImageScanCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetInstanceBuyType(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.InstanceBuyType = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetIntelligentAnalysisFlow(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.IntelligentAnalysisFlow = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetIsNewContainerVersion(v bool) *DescribeVersionConfigResponseBodyDataBody {
	s.IsNewContainerVersion = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetIsNewMultiVersion(v bool) *DescribeVersionConfigResponseBodyDataBody {
	s.IsNewMultiVersion = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetIsOverBalance(v bool) *DescribeVersionConfigResponseBodyDataBody {
	s.IsOverBalance = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetIsPostpay(v bool) *DescribeVersionConfigResponseBodyDataBody {
	s.IsPostpay = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetIsTrialVersion(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.IsTrialVersion = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetLastTrailEndTime(v int64) *DescribeVersionConfigResponseBodyDataBody {
	s.LastTrailEndTime = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetMergedVersion(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.MergedVersion = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetMultiVersion(v string) *DescribeVersionConfigResponseBodyDataBody {
	s.MultiVersion = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetMvAuthCount(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.MvAuthCount = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetMvUnusedAuthCount(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.MvUnusedAuthCount = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetNewThreatAnalysis(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.NewThreatAnalysis = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetOnboardedAssets(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.OnboardedAssets = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetOpenTime(v int64) *DescribeVersionConfigResponseBodyDataBody {
	s.OpenTime = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetPostPayHostVersion(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.PostPayHostVersion = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetPostPayInstanceId(v string) *DescribeVersionConfigResponseBodyDataBody {
	s.PostPayInstanceId = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetPostPayModuleSwitch(v string) *DescribeVersionConfigResponseBodyDataBody {
	s.PostPayModuleSwitch = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetPostPayOpenTime(v int64) *DescribeVersionConfigResponseBodyDataBody {
	s.PostPayOpenTime = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetPostPayStatus(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.PostPayStatus = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetRaspCapacity(v int64) *DescribeVersionConfigResponseBodyDataBody {
	s.RaspCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetReleaseTime(v int64) *DescribeVersionConfigResponseBodyDataBody {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetRequestId(v string) *DescribeVersionConfigResponseBodyDataBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetSasLog(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.SasLog = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetSasScreen(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.SasScreen = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetSdkCapacity(v int64) *DescribeVersionConfigResponseBodyDataBody {
	s.SdkCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetSlsCapacity(v int64) *DescribeVersionConfigResponseBodyDataBody {
	s.SlsCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetThreatAnalysisCapacity(v int64) *DescribeVersionConfigResponseBodyDataBody {
	s.ThreatAnalysisCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetThreatAnalysisFlow(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.ThreatAnalysisFlow = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetUserDefinedAlarms(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.UserDefinedAlarms = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetVersion(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.Version = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetVmCores(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.VmCores = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetVulFixCapacity(v int64) *DescribeVersionConfigResponseBodyDataBody {
	s.VulFixCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetWebLock(v int32) *DescribeVersionConfigResponseBodyDataBody {
	s.WebLock = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) SetWebLockAuthCount(v int64) *DescribeVersionConfigResponseBodyDataBody {
	s.WebLockAuthCount = &v
	return s
}

func (s *DescribeVersionConfigResponseBodyDataBody) Validate() error {
	return dara.Validate(s)
}
