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
	// example:
	//
	// 200
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeVersionConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6D462855-7835-5F91-835E-A62E44EC01CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 10
	AgentlessCapacity *int64 `json:"AgentlessCapacity,omitempty" xml:"AgentlessCapacity,omitempty"`
	// example:
	//
	// 1
	AllowPartialBuy *int32 `json:"AllowPartialBuy,omitempty" xml:"AllowPartialBuy,omitempty"`
	// example:
	//
	// 1680
	AntiRansomwareCapacity *int32 `json:"AntiRansomwareCapacity,omitempty" xml:"AntiRansomwareCapacity,omitempty"`
	// example:
	//
	// 1
	AntiRansomwareService *int32 `json:"AntiRansomwareService,omitempty" xml:"AntiRansomwareService,omitempty"`
	// example:
	//
	// 0
	AppWhiteList *int32 `json:"AppWhiteList,omitempty" xml:"AppWhiteList,omitempty"`
	// example:
	//
	// 20
	AppWhiteListAuthCount *int64 `json:"AppWhiteListAuthCount,omitempty" xml:"AppWhiteListAuthCount,omitempty"`
	// example:
	//
	// 30
	AssetLevel *int32 `json:"AssetLevel,omitempty" xml:"AssetLevel,omitempty"`
	// example:
	//
	// 0
	CanTryPostPaidPackage *int32 `json:"CanTryPostPaidPackage,omitempty" xml:"CanTryPostPaidPackage,omitempty"`
	// example:
	//
	// 10
	CspmCapacity *int64 `json:"CspmCapacity,omitempty" xml:"CspmCapacity,omitempty"`
	// example:
	//
	// 1
	HighestVersion *int32 `json:"HighestVersion,omitempty" xml:"HighestVersion,omitempty"`
	// example:
	//
	// 0
	HoneypotCapacity *int64 `json:"HoneypotCapacity,omitempty" xml:"HoneypotCapacity,omitempty"`
	// example:
	//
	// 1900
	ImageScanCapacity *int64 `json:"ImageScanCapacity,omitempty" xml:"ImageScanCapacity,omitempty"`
	// example:
	//
	// 1
	InstanceBuyType *int32 `json:"InstanceBuyType,omitempty" xml:"InstanceBuyType,omitempty"`
	// example:
	//
	// 100
	IntelligentAnalysisFlow *int32 `json:"IntelligentAnalysisFlow,omitempty" xml:"IntelligentAnalysisFlow,omitempty"`
	// example:
	//
	// true
	IsNewContainerVersion *bool `json:"IsNewContainerVersion,omitempty" xml:"IsNewContainerVersion,omitempty"`
	// example:
	//
	// true
	IsNewMultiVersion *bool `json:"IsNewMultiVersion,omitempty" xml:"IsNewMultiVersion,omitempty"`
	// example:
	//
	// false
	IsOverBalance *bool `json:"IsOverBalance,omitempty" xml:"IsOverBalance,omitempty"`
	// example:
	//
	// true
	IsPostpay *bool `json:"IsPostpay,omitempty" xml:"IsPostpay,omitempty"`
	// example:
	//
	// 0
	IsTrialVersion *int32 `json:"IsTrialVersion,omitempty" xml:"IsTrialVersion,omitempty"`
	// example:
	//
	// 1603934844000
	LastTrailEndTime *int64 `json:"LastTrailEndTime,omitempty" xml:"LastTrailEndTime,omitempty"`
	// example:
	//
	// 1
	MergedVersion *int32 `json:"MergedVersion,omitempty" xml:"MergedVersion,omitempty"`
	// example:
	//
	// null
	MultiVersion *string `json:"MultiVersion,omitempty" xml:"MultiVersion,omitempty"`
	// example:
	//
	// 0
	MvAuthCount *int32 `json:"MvAuthCount,omitempty" xml:"MvAuthCount,omitempty"`
	// example:
	//
	// 0
	MvUnusedAuthCount *int32 `json:"MvUnusedAuthCount,omitempty" xml:"MvUnusedAuthCount,omitempty"`
	// example:
	//
	// 0
	NewThreatAnalysis *int32 `json:"NewThreatAnalysis,omitempty" xml:"NewThreatAnalysis,omitempty"`
	// example:
	//
	// 0
	OnboardedAssets *int32 `json:"OnboardedAssets,omitempty" xml:"OnboardedAssets,omitempty"`
	// example:
	//
	// 1657244824669
	OpenTime *int64 `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	// example:
	//
	// 1
	PostPayHostVersion *int32 `json:"PostPayHostVersion,omitempty" xml:"PostPayHostVersion,omitempty"`
	// example:
	//
	// postpay-sas-frme8vjfiw2j
	PostPayInstanceId *string `json:"PostPayInstanceId,omitempty" xml:"PostPayInstanceId,omitempty"`
	// example:
	//
	// {\\"BASIC_SERVICE\\":0,\\"VUL\\":0}
	PostPayModuleSwitch *string `json:"PostPayModuleSwitch,omitempty" xml:"PostPayModuleSwitch,omitempty"`
	// example:
	//
	// 1698915219000
	PostPayOpenTime *int64 `json:"PostPayOpenTime,omitempty" xml:"PostPayOpenTime,omitempty"`
	// example:
	//
	// 1
	PostPayStatus *int32 `json:"PostPayStatus,omitempty" xml:"PostPayStatus,omitempty"`
	// example:
	//
	// 7
	RaspCapacity *int64 `json:"RaspCapacity,omitempty" xml:"RaspCapacity,omitempty"`
	// example:
	//
	// 1625846400000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// example:
	//
	// A6FB9AC3-4431-538F-BA8A-2A13AEA208A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	SasLog *int32 `json:"SasLog,omitempty" xml:"SasLog,omitempty"`
	// example:
	//
	// 0
	SasScreen *int32 `json:"SasScreen,omitempty" xml:"SasScreen,omitempty"`
	// example:
	//
	// 10
	SdkCapacity *int64 `json:"SdkCapacity,omitempty" xml:"SdkCapacity,omitempty"`
	// example:
	//
	// 200
	SlsCapacity *int64 `json:"SlsCapacity,omitempty" xml:"SlsCapacity,omitempty"`
	// example:
	//
	// 10
	ThreatAnalysisCapacity *int64 `json:"ThreatAnalysisCapacity,omitempty" xml:"ThreatAnalysisCapacity,omitempty"`
	// example:
	//
	// 10
	ThreatAnalysisFlow *int32 `json:"ThreatAnalysisFlow,omitempty" xml:"ThreatAnalysisFlow,omitempty"`
	// example:
	//
	// 0
	UserDefinedAlarms *int32 `json:"UserDefinedAlarms,omitempty" xml:"UserDefinedAlarms,omitempty"`
	// example:
	//
	// 3
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// 10
	VmCores *int32 `json:"VmCores,omitempty" xml:"VmCores,omitempty"`
	// example:
	//
	// 8
	VulFixCapacity *int64 `json:"VulFixCapacity,omitempty" xml:"VulFixCapacity,omitempty"`
	// example:
	//
	// 0
	WebLock *int32 `json:"WebLock,omitempty" xml:"WebLock,omitempty"`
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
