// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDevicesResponseBody
	GetCode() *string
	SetData(v []*ListDevicesResponseBodyData) *ListDevicesResponseBody
	GetData() []*ListDevicesResponseBodyData
	SetMessage(v string) *ListDevicesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListDevicesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDevicesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDevicesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDevicesResponseBody
	GetTotalCount() *int64
}

type ListDevicesResponseBody struct {
	Code       *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDevicesResponseBody) GetData() []*ListDevicesResponseBodyData {
	return s.Data
}

func (s *ListDevicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDevicesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDevicesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDevicesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDevicesResponseBody) SetCode(v string) *ListDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListDevicesResponseBody) SetData(v []*ListDevicesResponseBodyData) *ListDevicesResponseBody {
	s.Data = v
	return s
}

func (s *ListDevicesResponseBody) SetMessage(v string) *ListDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDevicesResponseBody) SetPageNumber(v int32) *ListDevicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDevicesResponseBody) SetPageSize(v int32) *ListDevicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDevicesResponseBody) SetRequestId(v string) *ListDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevicesResponseBody) SetTotalCount(v int64) *ListDevicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDevicesResponseBodyData struct {
	ActiveTime                 *string                                           `json:"ActiveTime,omitempty" xml:"ActiveTime,omitempty"`
	Alias                      *string                                           `json:"Alias,omitempty" xml:"Alias,omitempty"`
	AutoLockScreenTime         *int32                                            `json:"AutoLockScreenTime,omitempty" xml:"AutoLockScreenTime,omitempty"`
	AutoLogin                  *int32                                            `json:"AutoLogin,omitempty" xml:"AutoLogin,omitempty"`
	AutoType                   *string                                           `json:"AutoType,omitempty" xml:"AutoType,omitempty"`
	Bluetooth                  *string                                           `json:"Bluetooth,omitempty" xml:"Bluetooth,omitempty"`
	BuildId                    *string                                           `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	ClientId                   *string                                           `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientType                 *string                                           `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ConnectConfigs             []*ListDevicesResponseBodyDataConnectConfigs      `json:"ConnectConfigs,omitempty" xml:"ConnectConfigs,omitempty" type:"Repeated"`
	CustomIdleAction           *int32                                            `json:"CustomIdleAction,omitempty" xml:"CustomIdleAction,omitempty"`
	CustomPowerOn              *int32                                            `json:"CustomPowerOn,omitempty" xml:"CustomPowerOn,omitempty"`
	CustomProperty             *string                                           `json:"CustomProperty,omitempty" xml:"CustomProperty,omitempty"`
	CustomResourcePackage      *ListDevicesResponseBodyDataCustomResourcePackage `json:"CustomResourcePackage,omitempty" xml:"CustomResourcePackage,omitempty" type:"Struct"`
	DefinePowerButton          *int32                                            `json:"DefinePowerButton,omitempty" xml:"DefinePowerButton,omitempty"`
	DeviceIpV4                 *string                                           `json:"DeviceIpV4,omitempty" xml:"DeviceIpV4,omitempty"`
	DeviceLock                 *int32                                            `json:"DeviceLock,omitempty" xml:"DeviceLock,omitempty"`
	DeviceMqttConnectionStatus *int32                                            `json:"DeviceMqttConnectionStatus,omitempty" xml:"DeviceMqttConnectionStatus,omitempty"`
	DeviceName                 *string                                           `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceOS                   *string                                           `json:"DeviceOS,omitempty" xml:"DeviceOS,omitempty"`
	DevicePlatform             *string                                           `json:"DevicePlatform,omitempty" xml:"DevicePlatform,omitempty"`
	DisplayLayout              *string                                           `json:"DisplayLayout,omitempty" xml:"DisplayLayout,omitempty"`
	DisplayResolution          *string                                           `json:"DisplayResolution,omitempty" xml:"DisplayResolution,omitempty"`
	DisplayScaleRatio          *string                                           `json:"DisplayScaleRatio,omitempty" xml:"DisplayScaleRatio,omitempty"`
	EnableAdb                  *int32                                            `json:"EnableAdb,omitempty" xml:"EnableAdb,omitempty"`
	EnableAutoLockScreen       *int32                                            `json:"EnableAutoLockScreen,omitempty" xml:"EnableAutoLockScreen,omitempty"`
	EnableBluetooth            *int32                                            `json:"EnableBluetooth,omitempty" xml:"EnableBluetooth,omitempty"`
	EnableLockScreenPassword   *int32                                            `json:"EnableLockScreenPassword,omitempty" xml:"EnableLockScreenPassword,omitempty"`
	EnableModifyPassword       *int32                                            `json:"EnableModifyPassword,omitempty" xml:"EnableModifyPassword,omitempty"`
	EnableScheduledPowerOff    *int32                                            `json:"EnableScheduledPowerOff,omitempty" xml:"EnableScheduledPowerOff,omitempty"`
	EnableUnlockPassword       *int32                                            `json:"EnableUnlockPassword,omitempty" xml:"EnableUnlockPassword,omitempty"`
	EnableWlan                 *int32                                            `json:"EnableWlan,omitempty" xml:"EnableWlan,omitempty"`
	EndUserList                []*ListDevicesResponseBodyDataEndUserList         `json:"EndUserList,omitempty" xml:"EndUserList,omitempty" type:"Repeated"`
	EtherMac                   *string                                           `json:"EtherMac,omitempty" xml:"EtherMac,omitempty"`
	GmtModified                *string                                           `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GmtSync                    *string                                           `json:"GmtSync,omitempty" xml:"GmtSync,omitempty"`
	Id                         *int64                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	IdleTime                   *int32                                            `json:"IdleTime,omitempty" xml:"IdleTime,omitempty"`
	IsActive                   *string                                           `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
	LabelList                  []*ListDevicesResponseBodyDataLabelList           `json:"LabelList,omitempty" xml:"LabelList,omitempty" type:"Repeated"`
	LastLoginUser              *string                                           `json:"LastLoginUser,omitempty" xml:"LastLoginUser,omitempty"`
	LocalUsbPrint              *int32                                            `json:"LocalUsbPrint,omitempty" xml:"LocalUsbPrint,omitempty"`
	LocationInfo               *string                                           `json:"LocationInfo,omitempty" xml:"LocationInfo,omitempty"`
	LockPassword               *string                                           `json:"LockPassword,omitempty" xml:"LockPassword,omitempty"`
	Model                      *string                                           `json:"Model,omitempty" xml:"Model,omitempty"`
	OrderId                    *string                                           `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PeripheralConfig           *ListDevicesResponseBodyDataPeripheralConfig      `json:"PeripheralConfig,omitempty" xml:"PeripheralConfig,omitempty" type:"Struct"`
	ScheduledPowerOff          *string                                           `json:"ScheduledPowerOff,omitempty" xml:"ScheduledPowerOff,omitempty"`
	SecureNetworkType          *string                                           `json:"SecureNetworkType,omitempty" xml:"SecureNetworkType,omitempty"`
	SerialNo                   *string                                           `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SleepTime                  *int32                                            `json:"SleepTime,omitempty" xml:"SleepTime,omitempty"`
	Source                     *string                                           `json:"Source,omitempty" xml:"Source,omitempty"`
	TenantId                   *string                                           `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UsbStorage                 *int32                                            `json:"UsbStorage,omitempty" xml:"UsbStorage,omitempty"`
	Uuid                       *string                                           `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Wlan                       *string                                           `json:"Wlan,omitempty" xml:"Wlan,omitempty"`
}

func (s ListDevicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDevicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyData) GetActiveTime() *string {
	return s.ActiveTime
}

func (s *ListDevicesResponseBodyData) GetAlias() *string {
	return s.Alias
}

func (s *ListDevicesResponseBodyData) GetAutoLockScreenTime() *int32 {
	return s.AutoLockScreenTime
}

func (s *ListDevicesResponseBodyData) GetAutoLogin() *int32 {
	return s.AutoLogin
}

func (s *ListDevicesResponseBodyData) GetAutoType() *string {
	return s.AutoType
}

func (s *ListDevicesResponseBodyData) GetBluetooth() *string {
	return s.Bluetooth
}

func (s *ListDevicesResponseBodyData) GetBuildId() *string {
	return s.BuildId
}

func (s *ListDevicesResponseBodyData) GetClientId() *string {
	return s.ClientId
}

func (s *ListDevicesResponseBodyData) GetClientType() *string {
	return s.ClientType
}

func (s *ListDevicesResponseBodyData) GetConnectConfigs() []*ListDevicesResponseBodyDataConnectConfigs {
	return s.ConnectConfigs
}

func (s *ListDevicesResponseBodyData) GetCustomIdleAction() *int32 {
	return s.CustomIdleAction
}

func (s *ListDevicesResponseBodyData) GetCustomPowerOn() *int32 {
	return s.CustomPowerOn
}

func (s *ListDevicesResponseBodyData) GetCustomProperty() *string {
	return s.CustomProperty
}

func (s *ListDevicesResponseBodyData) GetCustomResourcePackage() *ListDevicesResponseBodyDataCustomResourcePackage {
	return s.CustomResourcePackage
}

func (s *ListDevicesResponseBodyData) GetDefinePowerButton() *int32 {
	return s.DefinePowerButton
}

func (s *ListDevicesResponseBodyData) GetDeviceIpV4() *string {
	return s.DeviceIpV4
}

func (s *ListDevicesResponseBodyData) GetDeviceLock() *int32 {
	return s.DeviceLock
}

func (s *ListDevicesResponseBodyData) GetDeviceMqttConnectionStatus() *int32 {
	return s.DeviceMqttConnectionStatus
}

func (s *ListDevicesResponseBodyData) GetDeviceName() *string {
	return s.DeviceName
}

func (s *ListDevicesResponseBodyData) GetDeviceOS() *string {
	return s.DeviceOS
}

func (s *ListDevicesResponseBodyData) GetDevicePlatform() *string {
	return s.DevicePlatform
}

func (s *ListDevicesResponseBodyData) GetDisplayLayout() *string {
	return s.DisplayLayout
}

func (s *ListDevicesResponseBodyData) GetDisplayResolution() *string {
	return s.DisplayResolution
}

func (s *ListDevicesResponseBodyData) GetDisplayScaleRatio() *string {
	return s.DisplayScaleRatio
}

func (s *ListDevicesResponseBodyData) GetEnableAdb() *int32 {
	return s.EnableAdb
}

func (s *ListDevicesResponseBodyData) GetEnableAutoLockScreen() *int32 {
	return s.EnableAutoLockScreen
}

func (s *ListDevicesResponseBodyData) GetEnableBluetooth() *int32 {
	return s.EnableBluetooth
}

func (s *ListDevicesResponseBodyData) GetEnableLockScreenPassword() *int32 {
	return s.EnableLockScreenPassword
}

func (s *ListDevicesResponseBodyData) GetEnableModifyPassword() *int32 {
	return s.EnableModifyPassword
}

func (s *ListDevicesResponseBodyData) GetEnableScheduledPowerOff() *int32 {
	return s.EnableScheduledPowerOff
}

func (s *ListDevicesResponseBodyData) GetEnableUnlockPassword() *int32 {
	return s.EnableUnlockPassword
}

func (s *ListDevicesResponseBodyData) GetEnableWlan() *int32 {
	return s.EnableWlan
}

func (s *ListDevicesResponseBodyData) GetEndUserList() []*ListDevicesResponseBodyDataEndUserList {
	return s.EndUserList
}

func (s *ListDevicesResponseBodyData) GetEtherMac() *string {
	return s.EtherMac
}

func (s *ListDevicesResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDevicesResponseBodyData) GetGmtSync() *string {
	return s.GmtSync
}

func (s *ListDevicesResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListDevicesResponseBodyData) GetIdleTime() *int32 {
	return s.IdleTime
}

func (s *ListDevicesResponseBodyData) GetIsActive() *string {
	return s.IsActive
}

func (s *ListDevicesResponseBodyData) GetLabelList() []*ListDevicesResponseBodyDataLabelList {
	return s.LabelList
}

func (s *ListDevicesResponseBodyData) GetLastLoginUser() *string {
	return s.LastLoginUser
}

func (s *ListDevicesResponseBodyData) GetLocalUsbPrint() *int32 {
	return s.LocalUsbPrint
}

func (s *ListDevicesResponseBodyData) GetLocationInfo() *string {
	return s.LocationInfo
}

func (s *ListDevicesResponseBodyData) GetLockPassword() *string {
	return s.LockPassword
}

func (s *ListDevicesResponseBodyData) GetModel() *string {
	return s.Model
}

func (s *ListDevicesResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *ListDevicesResponseBodyData) GetPeripheralConfig() *ListDevicesResponseBodyDataPeripheralConfig {
	return s.PeripheralConfig
}

func (s *ListDevicesResponseBodyData) GetScheduledPowerOff() *string {
	return s.ScheduledPowerOff
}

func (s *ListDevicesResponseBodyData) GetSecureNetworkType() *string {
	return s.SecureNetworkType
}

func (s *ListDevicesResponseBodyData) GetSerialNo() *string {
	return s.SerialNo
}

func (s *ListDevicesResponseBodyData) GetSleepTime() *int32 {
	return s.SleepTime
}

func (s *ListDevicesResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *ListDevicesResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *ListDevicesResponseBodyData) GetUsbStorage() *int32 {
	return s.UsbStorage
}

func (s *ListDevicesResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *ListDevicesResponseBodyData) GetWlan() *string {
	return s.Wlan
}

func (s *ListDevicesResponseBodyData) SetActiveTime(v string) *ListDevicesResponseBodyData {
	s.ActiveTime = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetAlias(v string) *ListDevicesResponseBodyData {
	s.Alias = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetAutoLockScreenTime(v int32) *ListDevicesResponseBodyData {
	s.AutoLockScreenTime = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetAutoLogin(v int32) *ListDevicesResponseBodyData {
	s.AutoLogin = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetAutoType(v string) *ListDevicesResponseBodyData {
	s.AutoType = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetBluetooth(v string) *ListDevicesResponseBodyData {
	s.Bluetooth = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetBuildId(v string) *ListDevicesResponseBodyData {
	s.BuildId = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetClientId(v string) *ListDevicesResponseBodyData {
	s.ClientId = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetClientType(v string) *ListDevicesResponseBodyData {
	s.ClientType = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetConnectConfigs(v []*ListDevicesResponseBodyDataConnectConfigs) *ListDevicesResponseBodyData {
	s.ConnectConfigs = v
	return s
}

func (s *ListDevicesResponseBodyData) SetCustomIdleAction(v int32) *ListDevicesResponseBodyData {
	s.CustomIdleAction = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetCustomPowerOn(v int32) *ListDevicesResponseBodyData {
	s.CustomPowerOn = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetCustomProperty(v string) *ListDevicesResponseBodyData {
	s.CustomProperty = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetCustomResourcePackage(v *ListDevicesResponseBodyDataCustomResourcePackage) *ListDevicesResponseBodyData {
	s.CustomResourcePackage = v
	return s
}

func (s *ListDevicesResponseBodyData) SetDefinePowerButton(v int32) *ListDevicesResponseBodyData {
	s.DefinePowerButton = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDeviceIpV4(v string) *ListDevicesResponseBodyData {
	s.DeviceIpV4 = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDeviceLock(v int32) *ListDevicesResponseBodyData {
	s.DeviceLock = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDeviceMqttConnectionStatus(v int32) *ListDevicesResponseBodyData {
	s.DeviceMqttConnectionStatus = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDeviceName(v string) *ListDevicesResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDeviceOS(v string) *ListDevicesResponseBodyData {
	s.DeviceOS = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDevicePlatform(v string) *ListDevicesResponseBodyData {
	s.DevicePlatform = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDisplayLayout(v string) *ListDevicesResponseBodyData {
	s.DisplayLayout = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDisplayResolution(v string) *ListDevicesResponseBodyData {
	s.DisplayResolution = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDisplayScaleRatio(v string) *ListDevicesResponseBodyData {
	s.DisplayScaleRatio = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEnableAdb(v int32) *ListDevicesResponseBodyData {
	s.EnableAdb = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEnableAutoLockScreen(v int32) *ListDevicesResponseBodyData {
	s.EnableAutoLockScreen = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEnableBluetooth(v int32) *ListDevicesResponseBodyData {
	s.EnableBluetooth = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEnableLockScreenPassword(v int32) *ListDevicesResponseBodyData {
	s.EnableLockScreenPassword = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEnableModifyPassword(v int32) *ListDevicesResponseBodyData {
	s.EnableModifyPassword = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEnableScheduledPowerOff(v int32) *ListDevicesResponseBodyData {
	s.EnableScheduledPowerOff = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEnableUnlockPassword(v int32) *ListDevicesResponseBodyData {
	s.EnableUnlockPassword = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEnableWlan(v int32) *ListDevicesResponseBodyData {
	s.EnableWlan = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetEndUserList(v []*ListDevicesResponseBodyDataEndUserList) *ListDevicesResponseBodyData {
	s.EndUserList = v
	return s
}

func (s *ListDevicesResponseBodyData) SetEtherMac(v string) *ListDevicesResponseBodyData {
	s.EtherMac = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetGmtModified(v string) *ListDevicesResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetGmtSync(v string) *ListDevicesResponseBodyData {
	s.GmtSync = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetId(v int64) *ListDevicesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetIdleTime(v int32) *ListDevicesResponseBodyData {
	s.IdleTime = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetIsActive(v string) *ListDevicesResponseBodyData {
	s.IsActive = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetLabelList(v []*ListDevicesResponseBodyDataLabelList) *ListDevicesResponseBodyData {
	s.LabelList = v
	return s
}

func (s *ListDevicesResponseBodyData) SetLastLoginUser(v string) *ListDevicesResponseBodyData {
	s.LastLoginUser = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetLocalUsbPrint(v int32) *ListDevicesResponseBodyData {
	s.LocalUsbPrint = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetLocationInfo(v string) *ListDevicesResponseBodyData {
	s.LocationInfo = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetLockPassword(v string) *ListDevicesResponseBodyData {
	s.LockPassword = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetModel(v string) *ListDevicesResponseBodyData {
	s.Model = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetOrderId(v string) *ListDevicesResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetPeripheralConfig(v *ListDevicesResponseBodyDataPeripheralConfig) *ListDevicesResponseBodyData {
	s.PeripheralConfig = v
	return s
}

func (s *ListDevicesResponseBodyData) SetScheduledPowerOff(v string) *ListDevicesResponseBodyData {
	s.ScheduledPowerOff = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetSecureNetworkType(v string) *ListDevicesResponseBodyData {
	s.SecureNetworkType = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetSerialNo(v string) *ListDevicesResponseBodyData {
	s.SerialNo = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetSleepTime(v int32) *ListDevicesResponseBodyData {
	s.SleepTime = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetSource(v string) *ListDevicesResponseBodyData {
	s.Source = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetTenantId(v string) *ListDevicesResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetUsbStorage(v int32) *ListDevicesResponseBodyData {
	s.UsbStorage = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetUuid(v string) *ListDevicesResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetWlan(v string) *ListDevicesResponseBodyData {
	s.Wlan = &v
	return s
}

func (s *ListDevicesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDevicesResponseBodyDataConnectConfigs struct {
	ConnectScript  *string `json:"ConnectScript,omitempty" xml:"ConnectScript,omitempty"`
	PeripheralPid  *string `json:"PeripheralPid,omitempty" xml:"PeripheralPid,omitempty"`
	PeripheralVid  *string `json:"PeripheralVid,omitempty" xml:"PeripheralVid,omitempty"`
	RedirectPolicy *int32  `json:"RedirectPolicy,omitempty" xml:"RedirectPolicy,omitempty"`
}

func (s ListDevicesResponseBodyDataConnectConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListDevicesResponseBodyDataConnectConfigs) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyDataConnectConfigs) GetConnectScript() *string {
	return s.ConnectScript
}

func (s *ListDevicesResponseBodyDataConnectConfigs) GetPeripheralPid() *string {
	return s.PeripheralPid
}

func (s *ListDevicesResponseBodyDataConnectConfigs) GetPeripheralVid() *string {
	return s.PeripheralVid
}

func (s *ListDevicesResponseBodyDataConnectConfigs) GetRedirectPolicy() *int32 {
	return s.RedirectPolicy
}

func (s *ListDevicesResponseBodyDataConnectConfigs) SetConnectScript(v string) *ListDevicesResponseBodyDataConnectConfigs {
	s.ConnectScript = &v
	return s
}

func (s *ListDevicesResponseBodyDataConnectConfigs) SetPeripheralPid(v string) *ListDevicesResponseBodyDataConnectConfigs {
	s.PeripheralPid = &v
	return s
}

func (s *ListDevicesResponseBodyDataConnectConfigs) SetPeripheralVid(v string) *ListDevicesResponseBodyDataConnectConfigs {
	s.PeripheralVid = &v
	return s
}

func (s *ListDevicesResponseBodyDataConnectConfigs) SetRedirectPolicy(v int32) *ListDevicesResponseBodyDataConnectConfigs {
	s.RedirectPolicy = &v
	return s
}

func (s *ListDevicesResponseBodyDataConnectConfigs) Validate() error {
	return dara.Validate(s)
}

type ListDevicesResponseBodyDataCustomResourcePackage struct {
	ConfigAboutLogo     *string `json:"ConfigAboutLogo,omitempty" xml:"ConfigAboutLogo,omitempty"`
	DesktopWallpaper    *string `json:"DesktopWallpaper,omitempty" xml:"DesktopWallpaper,omitempty"`
	LoginPageBackground *string `json:"LoginPageBackground,omitempty" xml:"LoginPageBackground,omitempty"`
	LoginPageLogo       *string `json:"LoginPageLogo,omitempty" xml:"LoginPageLogo,omitempty"`
	PersonalCenterLogo  *string `json:"PersonalCenterLogo,omitempty" xml:"PersonalCenterLogo,omitempty"`
	StartLogo           *string `json:"StartLogo,omitempty" xml:"StartLogo,omitempty"`
	StartMenuLogo       *string `json:"StartMenuLogo,omitempty" xml:"StartMenuLogo,omitempty"`
	UpgradeLogo         *string `json:"UpgradeLogo,omitempty" xml:"UpgradeLogo,omitempty"`
}

func (s ListDevicesResponseBodyDataCustomResourcePackage) String() string {
	return dara.Prettify(s)
}

func (s ListDevicesResponseBodyDataCustomResourcePackage) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) GetConfigAboutLogo() *string {
	return s.ConfigAboutLogo
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) GetDesktopWallpaper() *string {
	return s.DesktopWallpaper
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) GetLoginPageBackground() *string {
	return s.LoginPageBackground
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) GetLoginPageLogo() *string {
	return s.LoginPageLogo
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) GetPersonalCenterLogo() *string {
	return s.PersonalCenterLogo
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) GetStartLogo() *string {
	return s.StartLogo
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) GetStartMenuLogo() *string {
	return s.StartMenuLogo
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) GetUpgradeLogo() *string {
	return s.UpgradeLogo
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) SetConfigAboutLogo(v string) *ListDevicesResponseBodyDataCustomResourcePackage {
	s.ConfigAboutLogo = &v
	return s
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) SetDesktopWallpaper(v string) *ListDevicesResponseBodyDataCustomResourcePackage {
	s.DesktopWallpaper = &v
	return s
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) SetLoginPageBackground(v string) *ListDevicesResponseBodyDataCustomResourcePackage {
	s.LoginPageBackground = &v
	return s
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) SetLoginPageLogo(v string) *ListDevicesResponseBodyDataCustomResourcePackage {
	s.LoginPageLogo = &v
	return s
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) SetPersonalCenterLogo(v string) *ListDevicesResponseBodyDataCustomResourcePackage {
	s.PersonalCenterLogo = &v
	return s
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) SetStartLogo(v string) *ListDevicesResponseBodyDataCustomResourcePackage {
	s.StartLogo = &v
	return s
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) SetStartMenuLogo(v string) *ListDevicesResponseBodyDataCustomResourcePackage {
	s.StartMenuLogo = &v
	return s
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) SetUpgradeLogo(v string) *ListDevicesResponseBodyDataCustomResourcePackage {
	s.UpgradeLogo = &v
	return s
}

func (s *ListDevicesResponseBodyDataCustomResourcePackage) Validate() error {
	return dara.Validate(s)
}

type ListDevicesResponseBodyDataEndUserList struct {
	AdDomain    *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	BindTime    *string `json:"BindTime,omitempty" xml:"BindTime,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId   *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	SerialNo    *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserType    *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s ListDevicesResponseBodyDataEndUserList) String() string {
	return dara.Prettify(s)
}

func (s ListDevicesResponseBodyDataEndUserList) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyDataEndUserList) GetAdDomain() *string {
	return s.AdDomain
}

func (s *ListDevicesResponseBodyDataEndUserList) GetBindTime() *string {
	return s.BindTime
}

func (s *ListDevicesResponseBodyDataEndUserList) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListDevicesResponseBodyDataEndUserList) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListDevicesResponseBodyDataEndUserList) GetId() *int64 {
	return s.Id
}

func (s *ListDevicesResponseBodyDataEndUserList) GetSerialNo() *string {
	return s.SerialNo
}

func (s *ListDevicesResponseBodyDataEndUserList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListDevicesResponseBodyDataEndUserList) GetUserType() *string {
	return s.UserType
}

func (s *ListDevicesResponseBodyDataEndUserList) SetAdDomain(v string) *ListDevicesResponseBodyDataEndUserList {
	s.AdDomain = &v
	return s
}

func (s *ListDevicesResponseBodyDataEndUserList) SetBindTime(v string) *ListDevicesResponseBodyDataEndUserList {
	s.BindTime = &v
	return s
}

func (s *ListDevicesResponseBodyDataEndUserList) SetDirectoryId(v string) *ListDevicesResponseBodyDataEndUserList {
	s.DirectoryId = &v
	return s
}

func (s *ListDevicesResponseBodyDataEndUserList) SetEndUserId(v string) *ListDevicesResponseBodyDataEndUserList {
	s.EndUserId = &v
	return s
}

func (s *ListDevicesResponseBodyDataEndUserList) SetId(v int64) *ListDevicesResponseBodyDataEndUserList {
	s.Id = &v
	return s
}

func (s *ListDevicesResponseBodyDataEndUserList) SetSerialNo(v string) *ListDevicesResponseBodyDataEndUserList {
	s.SerialNo = &v
	return s
}

func (s *ListDevicesResponseBodyDataEndUserList) SetTenantId(v string) *ListDevicesResponseBodyDataEndUserList {
	s.TenantId = &v
	return s
}

func (s *ListDevicesResponseBodyDataEndUserList) SetUserType(v string) *ListDevicesResponseBodyDataEndUserList {
	s.UserType = &v
	return s
}

func (s *ListDevicesResponseBodyDataEndUserList) Validate() error {
	return dara.Validate(s)
}

type ListDevicesResponseBodyDataLabelList struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	LabelId     *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDevicesResponseBodyDataLabelList) String() string {
	return dara.Prettify(s)
}

func (s ListDevicesResponseBodyDataLabelList) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyDataLabelList) GetContent() *string {
	return s.Content
}

func (s *ListDevicesResponseBodyDataLabelList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListDevicesResponseBodyDataLabelList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListDevicesResponseBodyDataLabelList) GetLabelId() *string {
	return s.LabelId
}

func (s *ListDevicesResponseBodyDataLabelList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListDevicesResponseBodyDataLabelList) SetContent(v string) *ListDevicesResponseBodyDataLabelList {
	s.Content = &v
	return s
}

func (s *ListDevicesResponseBodyDataLabelList) SetGmtCreate(v int64) *ListDevicesResponseBodyDataLabelList {
	s.GmtCreate = &v
	return s
}

func (s *ListDevicesResponseBodyDataLabelList) SetGmtModified(v int64) *ListDevicesResponseBodyDataLabelList {
	s.GmtModified = &v
	return s
}

func (s *ListDevicesResponseBodyDataLabelList) SetLabelId(v string) *ListDevicesResponseBodyDataLabelList {
	s.LabelId = &v
	return s
}

func (s *ListDevicesResponseBodyDataLabelList) SetTenantId(v string) *ListDevicesResponseBodyDataLabelList {
	s.TenantId = &v
	return s
}

func (s *ListDevicesResponseBodyDataLabelList) Validate() error {
	return dara.Validate(s)
}

type ListDevicesResponseBodyDataPeripheralConfig struct {
	DefaultPolicy        *int32 `json:"DefaultPolicy,omitempty" xml:"DefaultPolicy,omitempty"`
	PolicyStrategy       *int32 `json:"PolicyStrategy,omitempty" xml:"PolicyStrategy,omitempty"`
	UsbAndInternalCamera *int32 `json:"UsbAndInternalCamera,omitempty" xml:"UsbAndInternalCamera,omitempty"`
	UsbPrinter           *int32 `json:"UsbPrinter,omitempty" xml:"UsbPrinter,omitempty"`
	UsbStorage           *int32 `json:"UsbStorage,omitempty" xml:"UsbStorage,omitempty"`
}

func (s ListDevicesResponseBodyDataPeripheralConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDevicesResponseBodyDataPeripheralConfig) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyDataPeripheralConfig) GetDefaultPolicy() *int32 {
	return s.DefaultPolicy
}

func (s *ListDevicesResponseBodyDataPeripheralConfig) GetPolicyStrategy() *int32 {
	return s.PolicyStrategy
}

func (s *ListDevicesResponseBodyDataPeripheralConfig) GetUsbAndInternalCamera() *int32 {
	return s.UsbAndInternalCamera
}

func (s *ListDevicesResponseBodyDataPeripheralConfig) GetUsbPrinter() *int32 {
	return s.UsbPrinter
}

func (s *ListDevicesResponseBodyDataPeripheralConfig) GetUsbStorage() *int32 {
	return s.UsbStorage
}

func (s *ListDevicesResponseBodyDataPeripheralConfig) SetDefaultPolicy(v int32) *ListDevicesResponseBodyDataPeripheralConfig {
	s.DefaultPolicy = &v
	return s
}

func (s *ListDevicesResponseBodyDataPeripheralConfig) SetPolicyStrategy(v int32) *ListDevicesResponseBodyDataPeripheralConfig {
	s.PolicyStrategy = &v
	return s
}

func (s *ListDevicesResponseBodyDataPeripheralConfig) SetUsbAndInternalCamera(v int32) *ListDevicesResponseBodyDataPeripheralConfig {
	s.UsbAndInternalCamera = &v
	return s
}

func (s *ListDevicesResponseBodyDataPeripheralConfig) SetUsbPrinter(v int32) *ListDevicesResponseBodyDataPeripheralConfig {
	s.UsbPrinter = &v
	return s
}

func (s *ListDevicesResponseBodyDataPeripheralConfig) SetUsbStorage(v int32) *ListDevicesResponseBodyDataPeripheralConfig {
	s.UsbStorage = &v
	return s
}

func (s *ListDevicesResponseBodyDataPeripheralConfig) Validate() error {
	return dara.Validate(s)
}
