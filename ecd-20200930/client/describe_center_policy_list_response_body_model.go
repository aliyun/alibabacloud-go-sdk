// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenterPolicyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescribePolicyGroups(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroups) *DescribeCenterPolicyListResponseBody
	GetDescribePolicyGroups() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroups
	SetRequestId(v string) *DescribeCenterPolicyListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCenterPolicyListResponseBody
	GetTotalCount() *int32
}

type DescribeCenterPolicyListResponseBody struct {
	// 雲端電腦原則的詳細資訊。
	DescribePolicyGroups []*DescribeCenterPolicyListResponseBodyDescribePolicyGroups `json:"DescribePolicyGroups,omitempty" xml:"DescribePolicyGroups,omitempty" type:"Repeated"`
	// 請求 ID。
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 查詢結果總數。
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenterPolicyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBody) GetDescribePolicyGroups() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	return s.DescribePolicyGroups
}

func (s *DescribeCenterPolicyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCenterPolicyListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCenterPolicyListResponseBody) SetDescribePolicyGroups(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroups) *DescribeCenterPolicyListResponseBody {
	s.DescribePolicyGroups = v
	return s
}

func (s *DescribeCenterPolicyListResponseBody) SetRequestId(v string) *DescribeCenterPolicyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBody) SetTotalCount(v int32) *DescribeCenterPolicyListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBody) Validate() error {
	if s.DescribePolicyGroups != nil {
		for _, item := range s.DescribePolicyGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroups struct {
	// 學術代理功能開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	AcademicProxy *string `json:"AcademicProxy,omitempty" xml:"AcademicProxy,omitempty"`
	// 使用者連線雲端電腦後是否具有管理員權限。
	//
	// > 該功能處於定向邀測中，暫不對外開放。
	//
	// example:
	//
	// deny
	AdminAccess *string `json:"AdminAccess,omitempty" xml:"AdminAccess,omitempty"`
	// 全螢幕模式下管理員鍵盤控制開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	AdminKeyboardOnFullScreen *string `json:"AdminKeyboardOnFullScreen,omitempty" xml:"AdminKeyboardOnFullScreen,omitempty"`
	// Windows 系統內管理員鍵盤控制開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	AdminKeyboardOnWindows *string `json:"AdminKeyboardOnWindows,omitempty" xml:"AdminKeyboardOnWindows,omitempty"`
	// 管理員監控開關。取值範圍：
	//
	// - 1：開啟。
	//
	// - 0：關閉。
	//
	// example:
	//
	// 0
	AdminMonitor *int32 `json:"AdminMonitor,omitempty" xml:"AdminMonitor,omitempty"`
	// 是否開啟防截圖功能。
	//
	// example:
	//
	// off
	AppContentProtection *string `json:"AppContentProtection,omitempty" xml:"AppContentProtection,omitempty"`
	// 防截圖行動裝置截圖自訂 OSS Bucket 清單。
	AppScreenshotBuckets []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets `json:"AppScreenshotBuckets,omitempty" xml:"AppScreenshotBuckets,omitempty" type:"Repeated"`
	// 用戶端存取 IP 白名單。
	AuthorizeAccessPolicyRules []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules `json:"AuthorizeAccessPolicyRules,omitempty" xml:"AuthorizeAccessPolicyRules,omitempty" type:"Repeated"`
	// 安全性群組管控規則清單。
	AuthorizeSecurityPolicyRules []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules `json:"AuthorizeSecurityPolicyRules,omitempty" xml:"AuthorizeSecurityPolicyRules,omitempty" type:"Repeated"`
	// 用戶端自動重新連線開關。
	//
	// example:
	//
	// off
	AutoReconnect *string `json:"AutoReconnect,omitempty" xml:"AutoReconnect,omitempty"`
	// 自動儲存行動裝置截圖。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	AutoSaveAppScreenshot *string `json:"AutoSaveAppScreenshot,omitempty" xml:"AutoSaveAppScreenshot,omitempty"`
	// 是否開啟本機攝影機重新導向。
	//
	// example:
	//
	// on
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// 用戶端控制選單顯示開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// on
	ClientControlMenu *string `json:"ClientControlMenu,omitempty" xml:"ClientControlMenu,omitempty"`
	// 用戶端建立自訂快照功能開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	ClientCreateSnapshot *string `json:"ClientCreateSnapshot,omitempty" xml:"ClientCreateSnapshot,omitempty"`
	// 用戶端休眠選單開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	ClientHibernate *string `json:"ClientHibernate,omitempty" xml:"ClientHibernate,omitempty"`
	// 用戶端重新啟動選單開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	ClientRestart *string `json:"ClientRestart,omitempty" xml:"ClientRestart,omitempty"`
	// 用戶端關機選單開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	ClientShutdown *string `json:"ClientShutdown,omitempty" xml:"ClientShutdown,omitempty"`
	// 登入方式管控規則。
	ClientTypes []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes `json:"ClientTypes,omitempty" xml:"ClientTypes,omitempty" type:"Repeated"`
	// 剪貼簿權限。
	//
	// example:
	//
	// off
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// 設計及 3D 應用場景選擇是否開啟色彩增強。
	//
	// example:
	//
	// off
	ColorEnhancement *string `json:"ColorEnhancement,omitempty" xml:"ColorEnhancement,omitempty"`
	// 本機驅動程式剪貼簿功能開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	CpdDriveClipboard *string `json:"CpdDriveClipboard,omitempty" xml:"CpdDriveClipboard,omitempty"`
	// CPU 降頻時長，取值範圍： 30\\~120。單位：秒。
	//
	// example:
	//
	// 30
	CpuDownGradeDuration *int32 `json:"CpuDownGradeDuration,omitempty" xml:"CpuDownGradeDuration,omitempty"`
	// CPU 負載保護開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	CpuOverload *string `json:"CpuOverload,omitempty" xml:"CpuOverload,omitempty"`
	// 處理程序名稱。
	CpuProcessors []*string `json:"CpuProcessors,omitempty" xml:"CpuProcessors,omitempty" type:"Repeated"`
	// CPU 保護模式開關。
	//
	// example:
	//
	// on
	CpuProtectedMode *string `json:"CpuProtectedMode,omitempty" xml:"CpuProtectedMode,omitempty"`
	// CPU 整體使用率百分比，取值範圍： 70\\~90。
	//
	// example:
	//
	// 70
	CpuRateLimit *int32 `json:"CpuRateLimit,omitempty" xml:"CpuRateLimit,omitempty"`
	// CPU 整體取樣時長，取值範圍： 10\\~60。單位：秒。
	//
	// example:
	//
	// 10
	CpuSampleDuration *int32 `json:"CpuSampleDuration,omitempty" xml:"CpuSampleDuration,omitempty"`
	// CPU 單核心使用率百分比，取值範圍： 70\\~100。
	//
	// example:
	//
	// 70
	CpuSingleRateLimit *int32 `json:"CpuSingleRateLimit,omitempty" xml:"CpuSingleRateLimit,omitempty"`
	// 已關聯該原則的雲端電腦數量。
	//
	// example:
	//
	// 1
	DesktopCount *int32 `json:"DesktopCount,omitempty" xml:"DesktopCount,omitempty"`
	// 已關聯該原則的雲端電腦集區數量。
	//
	// example:
	//
	// 1
	DesktopGroupCount *int32 `json:"DesktopGroupCount,omitempty" xml:"DesktopGroupCount,omitempty"`
	// 周邊裝置連線提示控制。
	//
	// example:
	//
	// off
	DeviceConnectHint *string `json:"DeviceConnectHint,omitempty" xml:"DeviceConnectHint,omitempty"`
	// 裝置重新導向規則。
	DeviceRedirects []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects `json:"DeviceRedirects,omitempty" xml:"DeviceRedirects,omitempty" type:"Repeated"`
	// 自訂周邊裝置規則。
	DeviceRules []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules `json:"DeviceRules,omitempty" xml:"DeviceRules,omitempty" type:"Repeated"`
	// 工作階段斷線保留。
	//
	// > 僅適用於雲端應用程式原則。
	//
	// example:
	//
	// persistent
	DisconnectKeepSession *string `json:"DisconnectKeepSession,omitempty" xml:"DisconnectKeepSession,omitempty"`
	// 工作階段斷線保留時長。單位：秒。
	//
	// > 僅適用於雲端應用程式原則。
	//
	// example:
	//
	// 120
	DisconnectKeepSessionTime *int32 `json:"DisconnectKeepSessionTime,omitempty" xml:"DisconnectKeepSessionTime,omitempty"`
	// 磁碟負載保護開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	DiskOverload *string `json:"DiskOverload,omitempty" xml:"DiskOverload,omitempty"`
	// 顯示模式。
	//
	// example:
	//
	// adminCustom
	DisplayMode *string `json:"DisplayMode,omitempty" xml:"DisplayMode,omitempty"`
	// 網域解析生效欄位。
	//
	// example:
	//
	// example.com
	DomainRegisterValue *string `json:"DomainRegisterValue,omitempty" xml:"DomainRegisterValue,omitempty"`
	// 網域解析原則詳情。
	DomainResolveRule []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule `json:"DomainResolveRule,omitempty" xml:"DomainResolveRule,omitempty" type:"Repeated"`
	// 網域解析原則開關。
	//
	// example:
	//
	// on
	DomainResolveRuleType *string `json:"DomainResolveRuleType,omitempty" xml:"DomainResolveRuleType,omitempty"`
	// 使用者請求管理員協助。
	//
	// example:
	//
	// off
	EndUserApplyAdminCoordinate *string `json:"EndUserApplyAdminCoordinate,omitempty" xml:"EndUserApplyAdminCoordinate,omitempty"`
	// 關聯的終端使用者數量。
	//
	// example:
	//
	// 3
	EndUserCount *string `json:"EndUserCount,omitempty" xml:"EndUserCount,omitempty"`
	// 同一個辦公網路下使用者共享雲端電腦。
	//
	// example:
	//
	// off
	EndUserGroupCoordinate *string `json:"EndUserGroupCoordinate,omitempty" xml:"EndUserGroupCoordinate,omitempty"`
	// 外接儲存裝置使用開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	ExternalDrive *string `json:"ExternalDrive,omitempty" xml:"ExternalDrive,omitempty"`
	// 檔案遷移。
	//
	// example:
	//
	// off
	FileMigrate *string `json:"FileMigrate,omitempty" xml:"FileMigrate,omitempty"`
	// 檔案傳輸功能的服務位址。
	//
	// example:
	//
	// filetransfer.example.com
	FileTransferAddress *string `json:"FileTransferAddress,omitempty" xml:"FileTransferAddress,omitempty"`
	// 單次傳入雲端電腦的檔案大小限制，需與傳入單位參數配合使用。
	//
	// example:
	//
	// 100
	FileTransferInSize *string `json:"FileTransferInSize,omitempty" xml:"FileTransferInSize,omitempty"`
	// 單次傳入雲端電腦的檔案大小限制單位。
	//
	// example:
	//
	// MB
	FileTransferInUnit *string `json:"FileTransferInUnit,omitempty" xml:"FileTransferInUnit,omitempty"`
	// 單次從雲端電腦傳出的檔案大小限制，需與傳出單位參數配合使用。
	//
	// example:
	//
	// 100
	FileTransferOutSize *string `json:"FileTransferOutSize,omitempty" xml:"FileTransferOutSize,omitempty"`
	// 單次從雲端電腦傳出的檔案大小限制單位。
	//
	// example:
	//
	// MB
	FileTransferOutUnit *string `json:"FileTransferOutUnit,omitempty" xml:"FileTransferOutUnit,omitempty"`
	// 檔案傳輸大小限制開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	FileTransferSizeLimit *string `json:"FileTransferSizeLimit,omitempty" xml:"FileTransferSizeLimit,omitempty"`
	// 檔案傳輸速度檔位。
	//
	// example:
	//
	// default
	FileTransferSpeed *string `json:"FileTransferSpeed,omitempty" xml:"FileTransferSpeed,omitempty"`
	// 用戶端設定的檔案傳輸速度生效位置。
	//
	// example:
	//
	// client
	FileTransferSpeedLocation *string `json:"FileTransferSpeedLocation,omitempty" xml:"FileTransferSpeedLocation,omitempty"`
	// 檔案同步開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	FolderSync *string `json:"FolderSync,omitempty" xml:"FolderSync,omitempty"`
	// 是否為圖形型規格雲端電腦開啟畫質策略。對雲端電腦效能和使用體驗要求較高時，例如專業設計場景下，建議開啟該策略。
	//
	// example:
	//
	// off
	GpuAcceleration *string `json:"GpuAcceleration,omitempty" xml:"GpuAcceleration,omitempty"`
	// 懸浮球設定訊息提示開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	HoverConfigMsg *string `json:"HoverConfigMsg,omitempty" xml:"HoverConfigMsg,omitempty"`
	// 懸浮球休眠按鈕開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	HoverHibernate *string `json:"HoverHibernate,omitempty" xml:"HoverHibernate,omitempty"`
	// 懸浮球重新啟動按鈕開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	HoverRestart *string `json:"HoverRestart,omitempty" xml:"HoverRestart,omitempty"`
	// 懸浮球關機按鈕開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	HoverShutdown *string `json:"HoverShutdown,omitempty" xml:"HoverShutdown,omitempty"`
	// Web用戶端存取策略。
	//
	// example:
	//
	// off
	Html5Access *string `json:"Html5Access,omitempty" xml:"Html5Access,omitempty"`
	// Web用戶端檔案傳輸策略。
	//
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// 網路通訊協定。
	//
	// example:
	//
	// tcp
	InternetCommunicationProtocol *string `json:"InternetCommunicationProtocol,omitempty" xml:"InternetCommunicationProtocol,omitempty"`
	// 網路印表機功能開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	InternetPrinter *string `json:"InternetPrinter,omitempty" xml:"InternetPrinter,omitempty"`
	// 懸浮球鍵盤控制開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	KeyboardControl *string `json:"KeyboardControl,omitempty" xml:"KeyboardControl,omitempty"`
	// KMS 伺服器開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	KmsServer *string `json:"KmsServer,omitempty" xml:"KmsServer,omitempty"`
	// KMS 伺服器設定，格式為位址:連接埠。
	//
	// example:
	//
	// kms.example.com:1688
	KmsServerConfig *string `json:"KmsServerConfig,omitempty" xml:"KmsServerConfig,omitempty"`
	// 本機磁碟映射權限。
	//
	// example:
	//
	// readwrite
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// 雲端電腦由於客觀原因造成中斷連線的最大連線重試時間。取值範圍：30\\~7200，單位：秒。
	//
	// example:
	//
	// 120
	MaxReconnectTime *int32 `json:"MaxReconnectTime,omitempty" xml:"MaxReconnectTime,omitempty"`
	// 開啟 MCP 服務。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	McpService *string `json:"McpService,omitempty" xml:"McpService,omitempty"`
	// 單個處理程序記憶體降頻時長，取值範圍： 30\\~120。單位：秒。
	//
	// example:
	//
	// 30
	MemoryDownGradeDuration *int32 `json:"MemoryDownGradeDuration,omitempty" xml:"MemoryDownGradeDuration,omitempty"`
	// 記憶體負載保護開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	MemoryOverload *string `json:"MemoryOverload,omitempty" xml:"MemoryOverload,omitempty"`
	// 處理程序名稱。
	MemoryProcessors []*string `json:"MemoryProcessors,omitempty" xml:"MemoryProcessors,omitempty" type:"Repeated"`
	// 記憶體保護模式開關。
	//
	// example:
	//
	// off
	MemoryProtectedMode *string `json:"MemoryProtectedMode,omitempty" xml:"MemoryProtectedMode,omitempty"`
	// 記憶體整體使用率百分比，取值範圍： 70\\~90。
	//
	// example:
	//
	// 70
	MemoryRateLimit *int32 `json:"MemoryRateLimit,omitempty" xml:"MemoryRateLimit,omitempty"`
	// 記憶體整體取樣時長，取值範圍： 30\\~60。單位：秒。
	//
	// example:
	//
	// 30
	MemorySampleDuration *int32 `json:"MemorySampleDuration,omitempty" xml:"MemorySampleDuration,omitempty"`
	// 單個處理程序記憶體使用率百分比，取值範圍： 30\\~60。
	//
	// example:
	//
	// 30
	MemorySingleRateLimit *int32 `json:"MemorySingleRateLimit,omitempty" xml:"MemorySingleRateLimit,omitempty"`
	// 透過行動端（Android用戶端<props="china">和iOS用戶端）連線雲端電腦時，是否在雲端電腦懸浮球內提供重新啟動按鈕。
	//
	// > 僅適用於V7.4或更高版本的行動端。
	//
	// example:
	//
	// off
	MobileRestart *string `json:"MobileRestart,omitempty" xml:"MobileRestart,omitempty"`
	// 行動端Windows 安全性控制開關。
	//
	// example:
	//
	// off
	MobileSafeMenu *string `json:"MobileSafeMenu,omitempty" xml:"MobileSafeMenu,omitempty"`
	// 透過行動端（Android用戶端<props="china">和iOS用戶端）連線雲端電腦時，是否在雲端電腦懸浮球內提供關機按鈕。
	//
	// > 僅適用於V7.4或更高版本的行動端。
	//
	// example:
	//
	// off
	MobileShutdown *string `json:"MobileShutdown,omitempty" xml:"MobileShutdown,omitempty"`
	// 行動端無影管家開關。
	//
	// example:
	//
	// off
	MobileWuyingKeeper *string `json:"MobileWuyingKeeper,omitempty" xml:"MobileWuyingKeeper,omitempty"`
	// 行動端小穎的開關。
	//
	// example:
	//
	// off
	MobileWyAssistant *string `json:"MobileWyAssistant,omitempty" xml:"MobileWyAssistant,omitempty"`
	// 模型庫功能開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	ModelLibrary *string `json:"ModelLibrary,omitempty" xml:"ModelLibrary,omitempty"`
	// 多螢幕顯示功能開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	MultiScreen *string `json:"MultiScreen,omitempty" xml:"MultiScreen,omitempty"`
	// 策略名稱。
	//
	// example:
	//
	// testPolicyGroupName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 網路重新導向。
	//
	// > 該功能處於定向邀測中，暫不對外開放。
	//
	// example:
	//
	// off
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// 網路重新導向策略詳細資訊。
	//
	// > 該功能處於定向邀測中，暫不對外開放。
	//
	// >
	NetRedirectRule []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule `json:"NetRedirectRule,omitempty" xml:"NetRedirectRule,omitempty" type:"Repeated"`
	// 網路印表機功能開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	NetworkPrinter *string `json:"NetworkPrinter,omitempty" xml:"NetworkPrinter,omitempty"`
	// 無操作中斷連線。
	//
	// > 僅適用於雲端應用策略。
	//
	// example:
	//
	// off
	NoOperationDisconnect *string `json:"NoOperationDisconnect,omitempty" xml:"NoOperationDisconnect,omitempty"`
	// 無操作中斷連線時長。單位：秒。
	//
	// > 僅適用於雲端應用策略。
	//
	// example:
	//
	// 120
	NoOperationDisconnectTime *int32 `json:"NoOperationDisconnectTime,omitempty" xml:"NoOperationDisconnectTime,omitempty"`
	// 關聯的組織數量。
	//
	// example:
	//
	// 2
	OrganizationCount *string `json:"OrganizationCount,omitempty" xml:"OrganizationCount,omitempty"`
	// 策略ID。
	//
	// example:
	//
	// pg-gx2x1dhsmthe9****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// 策略類型。
	//
	// example:
	//
	// SYSTEM
	PolicyGroupType *string `json:"PolicyGroupType,omitempty" xml:"PolicyGroupType,omitempty"`
	// 雲端電腦策略狀態。
	//
	// example:
	//
	// AVAILABLE
	PolicyStatus *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	// 連接埠代理功能開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	PortProxy *string `json:"PortProxy,omitempty" xml:"PortProxy,omitempty"`
	// 印表機彈出視窗提示。取值範圍：
	//
	// - default：預設值。
	//
	// - off：關閉。
	//
	// - custom：自訂。
	//
	// example:
	//
	// off
	PrinterAlert *string `json:"PrinterAlert,omitempty" xml:"PrinterAlert,omitempty"`
	// 印表機彈出視窗提示內容。
	//
	// example:
	//
	// Print Content.
	PrinterAlertContent *string `json:"PrinterAlertContent,omitempty" xml:"PrinterAlertContent,omitempty"`
	// 印表機彈出視窗提示標題。
	//
	// example:
	//
	// Print Title
	PrinterAlertTitle *string `json:"PrinterAlertTitle,omitempty" xml:"PrinterAlertTitle,omitempty"`
	// 印表機重新導向策略。
	//
	// example:
	//
	// off
	PrinterRedirection *string `json:"PrinterRedirection,omitempty" xml:"PrinterRedirection,omitempty"`
	// 設計及3D應用場景選擇是否開啟畫質增強。
	//
	// example:
	//
	// off
	QualityEnhancement *string `json:"QualityEnhancement,omitempty" xml:"QualityEnhancement,omitempty"`
	// 是否開啟自訂螢幕錄影。
	//
	// example:
	//
	// off
	RecordContent *string `json:"RecordContent,omitempty" xml:"RecordContent,omitempty"`
	// 自訂螢幕錄影的檔案過期時間，預設30天。
	//
	// example:
	//
	// 30
	RecordContentExpires *int64 `json:"RecordContentExpires,omitempty" xml:"RecordContentExpires,omitempty"`
	// 螢幕錄影稽核中偵測到事件發生後螢幕錄影的時長。單位：分鐘。取值範圍：10\\~60。
	//
	// example:
	//
	// 10
	RecordEventDuration *int32 `json:"RecordEventDuration,omitempty" xml:"RecordEventDuration,omitempty"`
	// 螢幕錄影稽核中的檔案監控絕對路徑。
	RecordEventFilePaths []*string `json:"RecordEventFilePaths,omitempty" xml:"RecordEventFilePaths,omitempty" type:"Repeated"`
	// 螢幕錄影稽核中的登錄檔監控絕對路徑。
	RecordEventRegisters []*string `json:"RecordEventRegisters,omitempty" xml:"RecordEventRegisters,omitempty" type:"Repeated"`
	// 是否開啟螢幕錄影。
	//
	// example:
	//
	// off
	Recording *string `json:"Recording,omitempty" xml:"Recording,omitempty"`
	// 錄製雲端電腦聲音選項。
	//
	// example:
	//
	// on
	RecordingAudio *string `json:"RecordingAudio,omitempty" xml:"RecordingAudio,omitempty"`
	// 螢幕錄影檔案閱覽時長，單位為分鐘。錄製檔案會根據您設定的時長自動分割並上傳至儲存空間，檔案達到300MB時優先滾動更新。
	//
	// example:
	//
	// 10
	RecordingDuration *int32 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// 螢幕錄影結束時間，取值：HH:MM:SS。只有Recording取值是period時，該傳回值有實際意義。
	//
	// example:
	//
	// 08:59:00
	RecordingEndTime *string `json:"RecordingEndTime,omitempty" xml:"RecordingEndTime,omitempty"`
	// 螢幕錄影檔案的保留時間。取值範圍為1\\~180天。
	//
	// example:
	//
	// 15
	RecordingExpires *int64 `json:"RecordingExpires,omitempty" xml:"RecordingExpires,omitempty"`
	// 螢幕錄影幀率。單位：FPS（幀/秒）。
	//
	// example:
	//
	// 5
	RecordingFps *int64 `json:"RecordingFps,omitempty" xml:"RecordingFps,omitempty"`
	// 螢幕錄影開始時間，取值：HH:MM:SS。只有Recording取值是period時，該傳回值有實際意義。
	//
	// example:
	//
	// 08:00:00
	RecordingStartTime *string `json:"RecordingStartTime,omitempty" xml:"RecordingStartTime,omitempty"`
	// 提示終端客戶已開啟螢幕錄影。
	//
	// example:
	//
	// off
	RecordingUserNotify *string `json:"RecordingUserNotify,omitempty" xml:"RecordingUserNotify,omitempty"`
	// 提示終端客戶開啟螢幕錄影的資訊。
	//
	// example:
	//
	// 螢幕錄影已開啟
	RecordingUserNotifyMessage *string `json:"RecordingUserNotifyMessage,omitempty" xml:"RecordingUserNotifyMessage,omitempty"`
	// 遠端協助鍵盤與滑鼠控制權限。
	//
	// example:
	//
	// fullControl
	RemoteCoordinate *string `json:"RemoteCoordinate,omitempty" xml:"RemoteCoordinate,omitempty"`
	// 重置雲端電腦。
	//
	// example:
	//
	// off
	ResetDesktop *string `json:"ResetDesktop,omitempty" xml:"ResetDesktop,omitempty"`
	// 螢幕解析度的 DPI 值。
	//
	// example:
	//
	// 96
	ResolutionDpi *int32 `json:"ResolutionDpi,omitempty" xml:"ResolutionDpi,omitempty"`
	// 解析度的高度。單位：畫素。
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// 解析度類型。
	//
	// example:
	//
	// adaptive
	ResolutionModel *string `json:"ResolutionModel,omitempty" xml:"ResolutionModel,omitempty"`
	// 解析度的寬度。單位：畫素。
	//
	// example:
	//
	// 1920
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// 已關聯該策略的資源群組數量。
	//
	// example:
	//
	// 1
	ResourceGroupCount *int32 `json:"ResourceGroupCount,omitempty" xml:"ResourceGroupCount,omitempty"`
	// 安全中心快速鍵開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	SafeMenu *string `json:"SafeMenu,omitempty" xml:"SafeMenu,omitempty"`
	// 策略的生效範圍。
	//
	// example:
	//
	// GLOBAL
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// 生效的 CIDR Block。
	ScopeValue []*string `json:"ScopeValue,omitempty" xml:"ScopeValue,omitempty" type:"Repeated"`
	// 螢幕顯示模式。
	//
	// example:
	//
	// auto
	ScreenDisplayMode *string `json:"ScreenDisplayMode,omitempty" xml:"ScreenDisplayMode,omitempty"`
	// 日常辦公場景選擇，是否開啟平滑度增強。
	//
	// example:
	//
	// off
	SmoothEnhancement *string `json:"SmoothEnhancement,omitempty" xml:"SmoothEnhancement,omitempty"`
	// 是否在雲端電腦懸浮球內提供狀態監控功能入口。
	//
	// example:
	//
	// on
	StatusMonitor *string `json:"StatusMonitor,omitempty" xml:"StatusMonitor,omitempty"`
	// 串流模式適配場景。
	//
	// example:
	//
	// smooth
	StreamingMode *string `json:"StreamingMode,omitempty" xml:"StreamingMode,omitempty"`
	// 目標幀率，取值範圍：10\\~60。
	//
	// example:
	//
	// 30
	TargetFps *int32 `json:"TargetFps,omitempty" xml:"TargetFps,omitempty"`
	// 應用程式工作列。
	//
	// > 僅適用於雲端應用程式策略。
	//
	// example:
	//
	// off
	Taskbar *string `json:"Taskbar,omitempty" xml:"Taskbar,omitempty"`
	// 三螢幕開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	ThreeScreen *string `json:"ThreeScreen,omitempty" xml:"ThreeScreen,omitempty"`
	// USB 重新導向。
	//
	// example:
	//
	// off
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	// USB 重新導向規則。
	UsbSupplyRedirectRule []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule `json:"UsbSupplyRedirectRule,omitempty" xml:"UsbSupplyRedirectRule,omitempty" type:"Repeated"`
	// 懸浮球使用時長顯示開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	UseTime *string `json:"UseTime,omitempty" xml:"UseTime,omitempty"`
	// 影片編碼平均位元率。單位：Kbps。取值範圍：1000\\~50000。
	//
	// example:
	//
	// 1000
	VideoEncAvgKbps *int32 `json:"VideoEncAvgKbps,omitempty" xml:"VideoEncAvgKbps,omitempty"`
	// 影片編碼最大 QP，畫面最低品質，取值範圍：0\\~51。
	//
	// example:
	//
	// 20
	VideoEncMaxQP *int32 `json:"VideoEncMaxQP,omitempty" xml:"VideoEncMaxQP,omitempty"`
	// 影片編碼最小 QP，最高品質。取值範圍：0\\~51。
	//
	// example:
	//
	// 20
	VideoEncMinQP *int32 `json:"VideoEncMinQP,omitempty" xml:"VideoEncMinQP,omitempty"`
	// 影片編碼峰值位元率。單位：Kbps。取值範圍：1000~50000。
	//
	// example:
	//
	// 2000
	VideoEncPeakKbps *int32 `json:"VideoEncPeakKbps,omitempty" xml:"VideoEncPeakKbps,omitempty"`
	// 影片編碼策略。
	//
	// example:
	//
	// qualityFirst
	VideoEncPolicy *string `json:"VideoEncPolicy,omitempty" xml:"VideoEncPolicy,omitempty"`
	// 多媒體重新導向。
	//
	// example:
	//
	// off
	VideoRedirect *string `json:"VideoRedirect,omitempty" xml:"VideoRedirect,omitempty"`
	// 影像顯示品質策略。
	//
	// example:
	//
	// medium
	VisualQuality *string `json:"VisualQuality,omitempty" xml:"VisualQuality,omitempty"`
	// 浮水印。
	//
	// example:
	//
	// on
	Watermark *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	// 隱藏式浮水印防拍照功能。
	//
	// example:
	//
	// off
	WatermarkAntiCam *string `json:"WatermarkAntiCam,omitempty" xml:"WatermarkAntiCam,omitempty"`
	// 浮水印字型顏色。取值範圍：0\\~16777215。
	//
	// example:
	//
	// 0
	WatermarkColor *int32 `json:"WatermarkColor,omitempty" xml:"WatermarkColor,omitempty"`
	// 若參數 `WatermarkType` 傳入 `custom`，則必須同時透過參數 `WatermarkCustomText` 傳入自訂文字的內容。
	//
	// example:
	//
	// custom-watermark
	WatermarkCustomText *string `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	// 浮水印傾斜度。取值範圍：-10\\~-30。
	//
	// example:
	//
	// -10
	WatermarkDegree *float64 `json:"WatermarkDegree,omitempty" xml:"WatermarkDegree,omitempty"`
	// 浮水印字型大小。取值範圍：10\\~20。
	//
	// example:
	//
	// 10
	WatermarkFontSize *int32 `json:"WatermarkFontSize,omitempty" xml:"WatermarkFontSize,omitempty"`
	// 浮水印字型樣式。
	//
	// example:
	//
	// plain
	WatermarkFontStyle *string `json:"WatermarkFontStyle,omitempty" xml:"WatermarkFontStyle,omitempty"`
	// 隱藏式浮水印增強功能。
	//
	// example:
	//
	// medium
	WatermarkPower *string `json:"WatermarkPower,omitempty" xml:"WatermarkPower,omitempty"`
	// 浮水印行數。
	//
	// example:
	//
	// 3
	WatermarkRowAmount *int32 `json:"WatermarkRowAmount,omitempty" xml:"WatermarkRowAmount,omitempty"`
	// 隱藏式浮水印安全優先規則。
	//
	// example:
	//
	// on
	WatermarkSecurity *string `json:"WatermarkSecurity,omitempty" xml:"WatermarkSecurity,omitempty"`
	// 浮水印陰影效果開關。取值範圍：
	//
	// - on：開啟。
	//
	// - off：關閉。
	//
	// example:
	//
	// off
	WatermarkShadow *string `json:"WatermarkShadow,omitempty" xml:"WatermarkShadow,omitempty"`
	// 浮水印透明度，值越大，透明度越小。取值範圍：10\\~100。
	//
	// example:
	//
	// 10
	WatermarkTransparencyValue *int32 `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	// 浮水印類型。
	//
	// example:
	//
	// EndUserId
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	// 無影管家開關。
	//
	// example:
	//
	// off
	WuyingKeeper *string `json:"WuyingKeeper,omitempty" xml:"WuyingKeeper,omitempty"`
	// 是否在雲端電腦懸浮球內提供無影 AI 助手功能入口。
	//
	// example:
	//
	// on
	WyAssistant *string `json:"WyAssistant,omitempty" xml:"WyAssistant,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAcademicProxy() *string {
	return s.AcademicProxy
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAdminAccess() *string {
	return s.AdminAccess
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAdminKeyboardOnFullScreen() *string {
	return s.AdminKeyboardOnFullScreen
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAdminKeyboardOnWindows() *string {
	return s.AdminKeyboardOnWindows
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAdminMonitor() *int32 {
	return s.AdminMonitor
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAppContentProtection() *string {
	return s.AppContentProtection
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAppScreenshotBuckets() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets {
	return s.AppScreenshotBuckets
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAuthorizeAccessPolicyRules() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules {
	return s.AuthorizeAccessPolicyRules
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAuthorizeSecurityPolicyRules() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	return s.AuthorizeSecurityPolicyRules
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAutoReconnect() *string {
	return s.AutoReconnect
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetAutoSaveAppScreenshot() *string {
	return s.AutoSaveAppScreenshot
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCameraRedirect() *string {
	return s.CameraRedirect
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetClientControlMenu() *string {
	return s.ClientControlMenu
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetClientCreateSnapshot() *string {
	return s.ClientCreateSnapshot
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetClientHibernate() *string {
	return s.ClientHibernate
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetClientRestart() *string {
	return s.ClientRestart
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetClientShutdown() *string {
	return s.ClientShutdown
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetClientTypes() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes {
	return s.ClientTypes
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetClipboard() *string {
	return s.Clipboard
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetColorEnhancement() *string {
	return s.ColorEnhancement
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCpdDriveClipboard() *string {
	return s.CpdDriveClipboard
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCpuDownGradeDuration() *int32 {
	return s.CpuDownGradeDuration
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCpuOverload() *string {
	return s.CpuOverload
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCpuProcessors() []*string {
	return s.CpuProcessors
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCpuProtectedMode() *string {
	return s.CpuProtectedMode
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCpuRateLimit() *int32 {
	return s.CpuRateLimit
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCpuSampleDuration() *int32 {
	return s.CpuSampleDuration
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetCpuSingleRateLimit() *int32 {
	return s.CpuSingleRateLimit
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDesktopCount() *int32 {
	return s.DesktopCount
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDesktopGroupCount() *int32 {
	return s.DesktopGroupCount
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDeviceConnectHint() *string {
	return s.DeviceConnectHint
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDeviceRedirects() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects {
	return s.DeviceRedirects
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDeviceRules() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules {
	return s.DeviceRules
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDisconnectKeepSession() *string {
	return s.DisconnectKeepSession
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDisconnectKeepSessionTime() *int32 {
	return s.DisconnectKeepSessionTime
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDiskOverload() *string {
	return s.DiskOverload
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDisplayMode() *string {
	return s.DisplayMode
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDomainRegisterValue() *string {
	return s.DomainRegisterValue
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDomainResolveRule() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule {
	return s.DomainResolveRule
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetDomainResolveRuleType() *string {
	return s.DomainResolveRuleType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetEndUserApplyAdminCoordinate() *string {
	return s.EndUserApplyAdminCoordinate
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetEndUserCount() *string {
	return s.EndUserCount
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetEndUserGroupCoordinate() *string {
	return s.EndUserGroupCoordinate
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetExternalDrive() *string {
	return s.ExternalDrive
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileMigrate() *string {
	return s.FileMigrate
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferAddress() *string {
	return s.FileTransferAddress
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferInSize() *string {
	return s.FileTransferInSize
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferInUnit() *string {
	return s.FileTransferInUnit
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferOutSize() *string {
	return s.FileTransferOutSize
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferOutUnit() *string {
	return s.FileTransferOutUnit
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferSizeLimit() *string {
	return s.FileTransferSizeLimit
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferSpeed() *string {
	return s.FileTransferSpeed
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFileTransferSpeedLocation() *string {
	return s.FileTransferSpeedLocation
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetFolderSync() *string {
	return s.FolderSync
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetGpuAcceleration() *string {
	return s.GpuAcceleration
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetHoverConfigMsg() *string {
	return s.HoverConfigMsg
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetHoverHibernate() *string {
	return s.HoverHibernate
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetHoverRestart() *string {
	return s.HoverRestart
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetHoverShutdown() *string {
	return s.HoverShutdown
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetHtml5Access() *string {
	return s.Html5Access
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetHtml5FileTransfer() *string {
	return s.Html5FileTransfer
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetInternetCommunicationProtocol() *string {
	return s.InternetCommunicationProtocol
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetInternetPrinter() *string {
	return s.InternetPrinter
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetKeyboardControl() *string {
	return s.KeyboardControl
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetKmsServer() *string {
	return s.KmsServer
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetKmsServerConfig() *string {
	return s.KmsServerConfig
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetLocalDrive() *string {
	return s.LocalDrive
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMaxReconnectTime() *int32 {
	return s.MaxReconnectTime
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMcpService() *string {
	return s.McpService
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMemoryDownGradeDuration() *int32 {
	return s.MemoryDownGradeDuration
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMemoryOverload() *string {
	return s.MemoryOverload
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMemoryProcessors() []*string {
	return s.MemoryProcessors
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMemoryProtectedMode() *string {
	return s.MemoryProtectedMode
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMemoryRateLimit() *int32 {
	return s.MemoryRateLimit
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMemorySampleDuration() *int32 {
	return s.MemorySampleDuration
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMemorySingleRateLimit() *int32 {
	return s.MemorySingleRateLimit
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMobileRestart() *string {
	return s.MobileRestart
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMobileSafeMenu() *string {
	return s.MobileSafeMenu
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMobileShutdown() *string {
	return s.MobileShutdown
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMobileWuyingKeeper() *string {
	return s.MobileWuyingKeeper
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMobileWyAssistant() *string {
	return s.MobileWyAssistant
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetModelLibrary() *string {
	return s.ModelLibrary
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetMultiScreen() *string {
	return s.MultiScreen
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetName() *string {
	return s.Name
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetNetRedirect() *string {
	return s.NetRedirect
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetNetRedirectRule() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule {
	return s.NetRedirectRule
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetNetworkPrinter() *string {
	return s.NetworkPrinter
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetNoOperationDisconnect() *string {
	return s.NoOperationDisconnect
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetNoOperationDisconnectTime() *int32 {
	return s.NoOperationDisconnectTime
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetOrganizationCount() *string {
	return s.OrganizationCount
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetPolicyGroupType() *string {
	return s.PolicyGroupType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetPolicyStatus() *string {
	return s.PolicyStatus
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetPortProxy() *string {
	return s.PortProxy
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetPrinterAlert() *string {
	return s.PrinterAlert
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetPrinterAlertContent() *string {
	return s.PrinterAlertContent
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetPrinterAlertTitle() *string {
	return s.PrinterAlertTitle
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetPrinterRedirection() *string {
	return s.PrinterRedirection
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetQualityEnhancement() *string {
	return s.QualityEnhancement
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordContent() *string {
	return s.RecordContent
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordContentExpires() *int64 {
	return s.RecordContentExpires
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordEventDuration() *int32 {
	return s.RecordEventDuration
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordEventFilePaths() []*string {
	return s.RecordEventFilePaths
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordEventRegisters() []*string {
	return s.RecordEventRegisters
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecording() *string {
	return s.Recording
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordingAudio() *string {
	return s.RecordingAudio
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordingDuration() *int32 {
	return s.RecordingDuration
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordingEndTime() *string {
	return s.RecordingEndTime
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordingExpires() *int64 {
	return s.RecordingExpires
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordingFps() *int64 {
	return s.RecordingFps
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordingStartTime() *string {
	return s.RecordingStartTime
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordingUserNotify() *string {
	return s.RecordingUserNotify
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRecordingUserNotifyMessage() *string {
	return s.RecordingUserNotifyMessage
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetRemoteCoordinate() *string {
	return s.RemoteCoordinate
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetResetDesktop() *string {
	return s.ResetDesktop
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetResolutionDpi() *int32 {
	return s.ResolutionDpi
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetResolutionModel() *string {
	return s.ResolutionModel
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetResourceGroupCount() *int32 {
	return s.ResourceGroupCount
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetSafeMenu() *string {
	return s.SafeMenu
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetScope() *string {
	return s.Scope
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetScopeValue() []*string {
	return s.ScopeValue
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetScreenDisplayMode() *string {
	return s.ScreenDisplayMode
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetSmoothEnhancement() *string {
	return s.SmoothEnhancement
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetStatusMonitor() *string {
	return s.StatusMonitor
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetStreamingMode() *string {
	return s.StreamingMode
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetTargetFps() *int32 {
	return s.TargetFps
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetTaskbar() *string {
	return s.Taskbar
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetThreeScreen() *string {
	return s.ThreeScreen
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetUsbRedirect() *string {
	return s.UsbRedirect
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetUsbSupplyRedirectRule() []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	return s.UsbSupplyRedirectRule
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetUseTime() *string {
	return s.UseTime
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetVideoEncAvgKbps() *int32 {
	return s.VideoEncAvgKbps
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetVideoEncMaxQP() *int32 {
	return s.VideoEncMaxQP
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetVideoEncMinQP() *int32 {
	return s.VideoEncMinQP
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetVideoEncPeakKbps() *int32 {
	return s.VideoEncPeakKbps
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetVideoEncPolicy() *string {
	return s.VideoEncPolicy
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetVideoRedirect() *string {
	return s.VideoRedirect
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetVisualQuality() *string {
	return s.VisualQuality
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermark() *string {
	return s.Watermark
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkAntiCam() *string {
	return s.WatermarkAntiCam
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkColor() *int32 {
	return s.WatermarkColor
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkCustomText() *string {
	return s.WatermarkCustomText
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkDegree() *float64 {
	return s.WatermarkDegree
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkFontSize() *int32 {
	return s.WatermarkFontSize
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkFontStyle() *string {
	return s.WatermarkFontStyle
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkPower() *string {
	return s.WatermarkPower
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkRowAmount() *int32 {
	return s.WatermarkRowAmount
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkSecurity() *string {
	return s.WatermarkSecurity
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkShadow() *string {
	return s.WatermarkShadow
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkTransparencyValue() *int32 {
	return s.WatermarkTransparencyValue
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWuyingKeeper() *string {
	return s.WuyingKeeper
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) GetWyAssistant() *string {
	return s.WyAssistant
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAcademicProxy(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AcademicProxy = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAdminAccess(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AdminAccess = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAdminKeyboardOnFullScreen(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AdminKeyboardOnFullScreen = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAdminKeyboardOnWindows(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AdminKeyboardOnWindows = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAdminMonitor(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AdminMonitor = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAppContentProtection(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AppContentProtection = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAppScreenshotBuckets(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AppScreenshotBuckets = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAuthorizeAccessPolicyRules(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AuthorizeAccessPolicyRules = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAuthorizeSecurityPolicyRules(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AuthorizeSecurityPolicyRules = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAutoReconnect(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AutoReconnect = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetAutoSaveAppScreenshot(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.AutoSaveAppScreenshot = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCameraRedirect(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CameraRedirect = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetClientControlMenu(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ClientControlMenu = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetClientCreateSnapshot(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ClientCreateSnapshot = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetClientHibernate(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ClientHibernate = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetClientRestart(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ClientRestart = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetClientShutdown(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ClientShutdown = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetClientTypes(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ClientTypes = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetClipboard(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.Clipboard = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetColorEnhancement(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ColorEnhancement = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCpdDriveClipboard(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CpdDriveClipboard = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCpuDownGradeDuration(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CpuDownGradeDuration = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCpuOverload(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CpuOverload = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCpuProcessors(v []*string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CpuProcessors = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCpuProtectedMode(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CpuProtectedMode = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCpuRateLimit(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CpuRateLimit = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCpuSampleDuration(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CpuSampleDuration = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetCpuSingleRateLimit(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.CpuSingleRateLimit = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDesktopCount(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DesktopCount = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDesktopGroupCount(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DesktopGroupCount = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDeviceConnectHint(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DeviceConnectHint = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDeviceRedirects(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DeviceRedirects = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDeviceRules(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DeviceRules = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDisconnectKeepSession(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DisconnectKeepSession = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDisconnectKeepSessionTime(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DisconnectKeepSessionTime = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDiskOverload(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DiskOverload = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDisplayMode(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DisplayMode = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDomainRegisterValue(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DomainRegisterValue = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDomainResolveRule(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DomainResolveRule = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetDomainResolveRuleType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.DomainResolveRuleType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetEndUserApplyAdminCoordinate(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.EndUserApplyAdminCoordinate = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetEndUserCount(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.EndUserCount = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetEndUserGroupCoordinate(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.EndUserGroupCoordinate = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetExternalDrive(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ExternalDrive = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileMigrate(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileMigrate = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferAddress(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferAddress = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferInSize(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferInSize = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferInUnit(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferInUnit = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferOutSize(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferOutSize = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferOutUnit(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferOutUnit = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferSizeLimit(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferSizeLimit = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferSpeed(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferSpeed = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFileTransferSpeedLocation(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FileTransferSpeedLocation = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetFolderSync(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.FolderSync = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetGpuAcceleration(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.GpuAcceleration = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetHoverConfigMsg(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.HoverConfigMsg = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetHoverHibernate(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.HoverHibernate = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetHoverRestart(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.HoverRestart = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetHoverShutdown(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.HoverShutdown = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetHtml5Access(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.Html5Access = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetHtml5FileTransfer(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.Html5FileTransfer = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetInternetCommunicationProtocol(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.InternetCommunicationProtocol = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetInternetPrinter(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.InternetPrinter = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetKeyboardControl(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.KeyboardControl = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetKmsServer(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.KmsServer = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetKmsServerConfig(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.KmsServerConfig = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetLocalDrive(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.LocalDrive = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMaxReconnectTime(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MaxReconnectTime = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMcpService(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.McpService = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMemoryDownGradeDuration(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MemoryDownGradeDuration = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMemoryOverload(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MemoryOverload = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMemoryProcessors(v []*string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MemoryProcessors = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMemoryProtectedMode(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MemoryProtectedMode = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMemoryRateLimit(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MemoryRateLimit = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMemorySampleDuration(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MemorySampleDuration = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMemorySingleRateLimit(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MemorySingleRateLimit = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMobileRestart(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MobileRestart = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMobileSafeMenu(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MobileSafeMenu = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMobileShutdown(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MobileShutdown = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMobileWuyingKeeper(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MobileWuyingKeeper = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMobileWyAssistant(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MobileWyAssistant = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetModelLibrary(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ModelLibrary = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetMultiScreen(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.MultiScreen = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetName(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.Name = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetNetRedirect(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.NetRedirect = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetNetRedirectRule(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.NetRedirectRule = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetNetworkPrinter(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.NetworkPrinter = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetNoOperationDisconnect(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.NoOperationDisconnect = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetNoOperationDisconnectTime(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.NoOperationDisconnectTime = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetOrganizationCount(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.OrganizationCount = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetPolicyGroupId(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetPolicyGroupType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.PolicyGroupType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetPolicyStatus(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.PolicyStatus = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetPortProxy(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.PortProxy = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetPrinterAlert(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.PrinterAlert = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetPrinterAlertContent(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.PrinterAlertContent = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetPrinterAlertTitle(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.PrinterAlertTitle = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetPrinterRedirection(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.PrinterRedirection = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetQualityEnhancement(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.QualityEnhancement = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordContent(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordContent = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordContentExpires(v int64) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordContentExpires = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordEventDuration(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordEventDuration = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordEventFilePaths(v []*string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordEventFilePaths = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordEventRegisters(v []*string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordEventRegisters = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecording(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.Recording = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordingAudio(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordingAudio = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordingDuration(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordingDuration = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordingEndTime(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordingEndTime = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordingExpires(v int64) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordingExpires = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordingFps(v int64) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordingFps = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordingStartTime(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordingStartTime = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordingUserNotify(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordingUserNotify = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRecordingUserNotifyMessage(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RecordingUserNotifyMessage = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetRemoteCoordinate(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.RemoteCoordinate = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetResetDesktop(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ResetDesktop = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetResolutionDpi(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ResolutionDpi = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetResolutionHeight(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ResolutionHeight = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetResolutionModel(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ResolutionModel = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetResolutionWidth(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ResolutionWidth = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetResourceGroupCount(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ResourceGroupCount = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetSafeMenu(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.SafeMenu = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetScope(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.Scope = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetScopeValue(v []*string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ScopeValue = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetScreenDisplayMode(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ScreenDisplayMode = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetSmoothEnhancement(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.SmoothEnhancement = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetStatusMonitor(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.StatusMonitor = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetStreamingMode(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.StreamingMode = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetTargetFps(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.TargetFps = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetTaskbar(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.Taskbar = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetThreeScreen(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.ThreeScreen = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetUsbRedirect(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.UsbRedirect = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetUsbSupplyRedirectRule(v []*DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.UsbSupplyRedirectRule = v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetUseTime(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.UseTime = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetVideoEncAvgKbps(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.VideoEncAvgKbps = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetVideoEncMaxQP(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.VideoEncMaxQP = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetVideoEncMinQP(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.VideoEncMinQP = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetVideoEncPeakKbps(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.VideoEncPeakKbps = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetVideoEncPolicy(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.VideoEncPolicy = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetVideoRedirect(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.VideoRedirect = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetVisualQuality(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.VisualQuality = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermark(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.Watermark = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkAntiCam(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkAntiCam = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkColor(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkColor = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkCustomText(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkCustomText = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkDegree(v float64) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkDegree = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkFontSize(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkFontSize = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkFontStyle(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkFontStyle = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkPower(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkPower = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkRowAmount(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkRowAmount = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkSecurity(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkSecurity = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkShadow(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkShadow = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkTransparencyValue(v int32) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkTransparencyValue = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWatermarkType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WatermarkType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWuyingKeeper(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WuyingKeeper = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) SetWyAssistant(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroups {
	s.WyAssistant = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroups) Validate() error {
	if s.AppScreenshotBuckets != nil {
		for _, item := range s.AppScreenshotBuckets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AuthorizeAccessPolicyRules != nil {
		for _, item := range s.AuthorizeAccessPolicyRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AuthorizeSecurityPolicyRules != nil {
		for _, item := range s.AuthorizeSecurityPolicyRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ClientTypes != nil {
		for _, item := range s.ClientTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DeviceRedirects != nil {
		for _, item := range s.DeviceRedirects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DeviceRules != nil {
		for _, item := range s.DeviceRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DomainResolveRule != nil {
		for _, item := range s.DomainResolveRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NetRedirectRule != nil {
		for _, item := range s.NetRedirectRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UsbSupplyRedirectRule != nil {
		for _, item := range s.UsbSupplyRedirectRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets struct {
	// Bucket 名稱。
	//
	// example:
	//
	// my-bucket-01
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// 使用者 OSS Bucket 所在地域，必須與 Bucket 的實際地域一致。
	//
	// example:
	//
	// cn-beijing
	BucketRegionId *string `json:"BucketRegionId,omitempty" xml:"BucketRegionId,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets) GetBucketRegionId() *string {
	return s.BucketRegionId
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets) SetBucketName(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets {
	s.BucketName = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets) SetBucketRegionId(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets {
	s.BucketRegionId = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAppScreenshotBuckets) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules struct {
	// 用戶端存取 IP 位址區段。格式為 CIDR 格式的 IPv4 位址區段。
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// 用戶端存取 IP 位址區段備註資訊。
	//
	// example:
	//
	// 公司辦公網段
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) GetCidrIp() *string {
	return s.CidrIp
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) SetCidrIp(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules {
	s.CidrIp = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) SetDescription(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules {
	s.Description = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules struct {
	// 安全性群組管控規則的對象，CIDR 格式的 IPv4 位址區段。
	//
	// example:
	//
	// 47.100.XX.XX/16
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// 安全性群組管控規則的描述資訊。
	//
	// example:
	//
	// 允許存取內網研發環境
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 安全性群組管控規則的通訊協定類型。
	//
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// 安全性群組管控規則的授權原則。
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// 安全性群組管控規則的連接埠範圍。
	//
	// example:
	//
	// 22/22
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// 安全性群組管控規則的優先順序，數字越小，代表優先順序越高。
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// 安全性群組管控規則的規則方向。
	//
	// example:
	//
	// inflow
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetCidrIp() *string {
	return s.CidrIp
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetPortRange() *string {
	return s.PortRange
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetPriority() *string {
	return s.Priority
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GetType() *string {
	return s.Type
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetCidrIp(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.CidrIp = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetDescription(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Description = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetIpProtocol(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.IpProtocol = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPolicy(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Policy = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPortRange(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.PortRange = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPriority(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Priority = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Type = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes struct {
	// 用戶端類型。
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// 是否允許某一類型的用戶端登入雲端桌面。
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes) GetClientType() *string {
	return s.ClientType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes) GetStatus() *string {
	return s.Status
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes) SetClientType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes {
	s.ClientType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes) SetStatus(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes {
	s.Status = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsClientTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects struct {
	// 周邊裝置類型。
	//
	// example:
	//
	// camera
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// 重新導向類型。
	//
	// example:
	//
	// usbRedirect
	RedirectType *string `json:"RedirectType,omitempty" xml:"RedirectType,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects) GetRedirectType() *string {
	return s.RedirectType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects) SetDeviceType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects {
	s.DeviceType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects) SetRedirectType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects {
	s.RedirectType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRedirects) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules struct {
	// 裝置名稱。
	//
	// example:
	//
	// sandisk
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// 產品 ID。
	//
	// example:
	//
	// 0x55b1
	DevicePid *string `json:"DevicePid,omitempty" xml:"DevicePid,omitempty"`
	// 周邊裝置類型。
	//
	// example:
	//
	// storage
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// 生產廠商 ID。參見 [Valid USB Vendor IDs (VIDs)](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf)。
	//
	// example:
	//
	// 0x0781
	DeviceVid *string `json:"DeviceVid,omitempty" xml:"DeviceVid,omitempty"`
	// 連結最佳化指令。
	//
	// example:
	//
	// 2:0
	OptCommand *string `json:"OptCommand,omitempty" xml:"OptCommand,omitempty"`
	// 裝置規則適用的平台類型。
	//
	// example:
	//
	// Windows
	Platforms *string `json:"Platforms,omitempty" xml:"Platforms,omitempty"`
	// 重新導向類型。
	//
	// example:
	//
	// usbRedirect
	RedirectType *string `json:"RedirectType,omitempty" xml:"RedirectType,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) GetDevicePid() *string {
	return s.DevicePid
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) GetDeviceVid() *string {
	return s.DeviceVid
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) GetOptCommand() *string {
	return s.OptCommand
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) GetPlatforms() *string {
	return s.Platforms
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) GetRedirectType() *string {
	return s.RedirectType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) SetDeviceName(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules {
	s.DeviceName = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) SetDevicePid(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules {
	s.DevicePid = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) SetDeviceType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules {
	s.DeviceType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) SetDeviceVid(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules {
	s.DeviceVid = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) SetOptCommand(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules {
	s.OptCommand = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) SetPlatforms(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules {
	s.Platforms = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) SetRedirectType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules {
	s.RedirectType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDeviceRules) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule struct {
	// 原則描述。
	//
	// example:
	//
	// 測試規則
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 網域。
	//
	// example:
	//
	// *.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 解析原則。
	//
	// example:
	//
	// allow
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) GetDescription() *string {
	return s.Description
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) GetDomain() *string {
	return s.Domain
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) SetDescription(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule {
	s.Description = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) SetDomain(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule {
	s.Domain = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) SetPolicy(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule {
	s.Policy = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsDomainResolveRule) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule struct {
	// 網域。
	//
	// example:
	//
	// *.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 重新導向策略。
	//
	// example:
	//
	// allow
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// 規則類型。
	//
	// example:
	//
	// domain
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) GetDomain() *string {
	return s.Domain
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) SetDomain(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule {
	s.Domain = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) SetPolicy(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule {
	s.Policy = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) SetRuleType(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule {
	s.RuleType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsNetRedirectRule) Validate() error {
	return dara.Validate(s)
}

type DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule struct {
	// 規則描述。
	//
	// example:
	//
	// 測試規則
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 產品 ID。
	//
	// example:
	//
	// 08**
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// USB 重新導向類型。
	//
	// example:
	//
	// 1
	UsbRedirectType *int64 `json:"UsbRedirectType,omitempty" xml:"UsbRedirectType,omitempty"`
	// USB 重新導向規則類型。
	//
	// example:
	//
	// 1
	UsbRuleType *int64 `json:"UsbRuleType,omitempty" xml:"UsbRuleType,omitempty"`
	// 生產廠商 ID。參見 [Valid USB Vendor IDs (VIDs)](https://www.usb.org/sites/default/files/vendor_ids032322.pdf_1.pdf)。
	//
	// example:
	//
	// 04**
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetDescription() *string {
	return s.Description
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetProductId() *string {
	return s.ProductId
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetUsbRedirectType() *int64 {
	return s.UsbRedirectType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetUsbRuleType() *int64 {
	return s.UsbRuleType
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) GetVendorId() *string {
	return s.VendorId
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetDescription(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.Description = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetProductId(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.ProductId = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetUsbRedirectType(v int64) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.UsbRedirectType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetUsbRuleType(v int64) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.UsbRuleType = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) SetVendorId(v string) *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule {
	s.VendorId = &v
	return s
}

func (s *DescribeCenterPolicyListResponseBodyDescribePolicyGroupsUsbSupplyRedirectRule) Validate() error {
	return dara.Validate(s)
}
