// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDeviceConfigsResponseBody
	GetCode() *string
	SetData(v *GetDeviceConfigsResponseBodyData) *GetDeviceConfigsResponseBody
	GetData() *GetDeviceConfigsResponseBodyData
	SetMessage(v string) *GetDeviceConfigsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeviceConfigsResponseBody
	GetRequestId() *string
}

type GetDeviceConfigsResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDeviceConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDeviceConfigsResponseBody) GetData() *GetDeviceConfigsResponseBodyData {
	return s.Data
}

func (s *GetDeviceConfigsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeviceConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceConfigsResponseBody) SetCode(v string) *GetDeviceConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceConfigsResponseBody) SetData(v *GetDeviceConfigsResponseBodyData) *GetDeviceConfigsResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceConfigsResponseBody) SetMessage(v string) *GetDeviceConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceConfigsResponseBody) SetRequestId(v string) *GetDeviceConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDeviceConfigsResponseBodyData struct {
	AutoLockScreenTime       *int32                                                 `json:"AutoLockScreenTime,omitempty" xml:"AutoLockScreenTime,omitempty"`
	AutoLogin                *int32                                                 `json:"AutoLogin,omitempty" xml:"AutoLogin,omitempty"`
	AutoUpdate               *int32                                                 `json:"AutoUpdate,omitempty" xml:"AutoUpdate,omitempty"`
	CustomIdleAction         *int32                                                 `json:"CustomIdleAction,omitempty" xml:"CustomIdleAction,omitempty"`
	CustomPowerOn            *int32                                                 `json:"CustomPowerOn,omitempty" xml:"CustomPowerOn,omitempty"`
	CustomResourcePackage    *GetDeviceConfigsResponseBodyDataCustomResourcePackage `json:"CustomResourcePackage,omitempty" xml:"CustomResourcePackage,omitempty" type:"Struct"`
	DefinePowerButton        *int32                                                 `json:"DefinePowerButton,omitempty" xml:"DefinePowerButton,omitempty"`
	DeviceLock               *int32                                                 `json:"DeviceLock,omitempty" xml:"DeviceLock,omitempty"`
	DisplayLayout            *string                                                `json:"DisplayLayout,omitempty" xml:"DisplayLayout,omitempty"`
	DisplayResolution        *string                                                `json:"DisplayResolution,omitempty" xml:"DisplayResolution,omitempty"`
	DisplayScaleRatio        *string                                                `json:"DisplayScaleRatio,omitempty" xml:"DisplayScaleRatio,omitempty"`
	EnableAdb                *int32                                                 `json:"EnableAdb,omitempty" xml:"EnableAdb,omitempty"`
	EnableAutoLockScreen     *int32                                                 `json:"EnableAutoLockScreen,omitempty" xml:"EnableAutoLockScreen,omitempty"`
	EnableBluetooth          *int32                                                 `json:"EnableBluetooth,omitempty" xml:"EnableBluetooth,omitempty"`
	EnableLockScreenPassword *int32                                                 `json:"EnableLockScreenPassword,omitempty" xml:"EnableLockScreenPassword,omitempty"`
	EnableModifyPassword     *int32                                                 `json:"EnableModifyPassword,omitempty" xml:"EnableModifyPassword,omitempty"`
	EnableScheduledPowerOff  *int32                                                 `json:"EnableScheduledPowerOff,omitempty" xml:"EnableScheduledPowerOff,omitempty"`
	EnableUnlockPassword     *int32                                                 `json:"EnableUnlockPassword,omitempty" xml:"EnableUnlockPassword,omitempty"`
	EnableWlan               *int32                                                 `json:"EnableWlan,omitempty" xml:"EnableWlan,omitempty"`
	IdleTime                 *int32                                                 `json:"IdleTime,omitempty" xml:"IdleTime,omitempty"`
	LocalUsbPrint            *int32                                                 `json:"LocalUsbPrint,omitempty" xml:"LocalUsbPrint,omitempty"`
	LockPassword             *string                                                `json:"LockPassword,omitempty" xml:"LockPassword,omitempty"`
	ScheduledPowerOff        *string                                                `json:"ScheduledPowerOff,omitempty" xml:"ScheduledPowerOff,omitempty"`
	SecureNetworkType        *string                                                `json:"SecureNetworkType,omitempty" xml:"SecureNetworkType,omitempty"`
	SerialNo                 *string                                                `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SleepTime                *int32                                                 `json:"SleepTime,omitempty" xml:"SleepTime,omitempty"`
	Urcl                     *string                                                `json:"Urcl,omitempty" xml:"Urcl,omitempty"`
	UsbStorage               *int32                                                 `json:"UsbStorage,omitempty" xml:"UsbStorage,omitempty"`
	UserCustomId             *string                                                `json:"UserCustomId,omitempty" xml:"UserCustomId,omitempty"`
	Uuid                     *string                                                `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetDeviceConfigsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigsResponseBodyData) GetAutoLockScreenTime() *int32 {
	return s.AutoLockScreenTime
}

func (s *GetDeviceConfigsResponseBodyData) GetAutoLogin() *int32 {
	return s.AutoLogin
}

func (s *GetDeviceConfigsResponseBodyData) GetAutoUpdate() *int32 {
	return s.AutoUpdate
}

func (s *GetDeviceConfigsResponseBodyData) GetCustomIdleAction() *int32 {
	return s.CustomIdleAction
}

func (s *GetDeviceConfigsResponseBodyData) GetCustomPowerOn() *int32 {
	return s.CustomPowerOn
}

func (s *GetDeviceConfigsResponseBodyData) GetCustomResourcePackage() *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	return s.CustomResourcePackage
}

func (s *GetDeviceConfigsResponseBodyData) GetDefinePowerButton() *int32 {
	return s.DefinePowerButton
}

func (s *GetDeviceConfigsResponseBodyData) GetDeviceLock() *int32 {
	return s.DeviceLock
}

func (s *GetDeviceConfigsResponseBodyData) GetDisplayLayout() *string {
	return s.DisplayLayout
}

func (s *GetDeviceConfigsResponseBodyData) GetDisplayResolution() *string {
	return s.DisplayResolution
}

func (s *GetDeviceConfigsResponseBodyData) GetDisplayScaleRatio() *string {
	return s.DisplayScaleRatio
}

func (s *GetDeviceConfigsResponseBodyData) GetEnableAdb() *int32 {
	return s.EnableAdb
}

func (s *GetDeviceConfigsResponseBodyData) GetEnableAutoLockScreen() *int32 {
	return s.EnableAutoLockScreen
}

func (s *GetDeviceConfigsResponseBodyData) GetEnableBluetooth() *int32 {
	return s.EnableBluetooth
}

func (s *GetDeviceConfigsResponseBodyData) GetEnableLockScreenPassword() *int32 {
	return s.EnableLockScreenPassword
}

func (s *GetDeviceConfigsResponseBodyData) GetEnableModifyPassword() *int32 {
	return s.EnableModifyPassword
}

func (s *GetDeviceConfigsResponseBodyData) GetEnableScheduledPowerOff() *int32 {
	return s.EnableScheduledPowerOff
}

func (s *GetDeviceConfigsResponseBodyData) GetEnableUnlockPassword() *int32 {
	return s.EnableUnlockPassword
}

func (s *GetDeviceConfigsResponseBodyData) GetEnableWlan() *int32 {
	return s.EnableWlan
}

func (s *GetDeviceConfigsResponseBodyData) GetIdleTime() *int32 {
	return s.IdleTime
}

func (s *GetDeviceConfigsResponseBodyData) GetLocalUsbPrint() *int32 {
	return s.LocalUsbPrint
}

func (s *GetDeviceConfigsResponseBodyData) GetLockPassword() *string {
	return s.LockPassword
}

func (s *GetDeviceConfigsResponseBodyData) GetScheduledPowerOff() *string {
	return s.ScheduledPowerOff
}

func (s *GetDeviceConfigsResponseBodyData) GetSecureNetworkType() *string {
	return s.SecureNetworkType
}

func (s *GetDeviceConfigsResponseBodyData) GetSerialNo() *string {
	return s.SerialNo
}

func (s *GetDeviceConfigsResponseBodyData) GetSleepTime() *int32 {
	return s.SleepTime
}

func (s *GetDeviceConfigsResponseBodyData) GetUrcl() *string {
	return s.Urcl
}

func (s *GetDeviceConfigsResponseBodyData) GetUsbStorage() *int32 {
	return s.UsbStorage
}

func (s *GetDeviceConfigsResponseBodyData) GetUserCustomId() *string {
	return s.UserCustomId
}

func (s *GetDeviceConfigsResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *GetDeviceConfigsResponseBodyData) SetAutoLockScreenTime(v int32) *GetDeviceConfigsResponseBodyData {
	s.AutoLockScreenTime = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetAutoLogin(v int32) *GetDeviceConfigsResponseBodyData {
	s.AutoLogin = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetAutoUpdate(v int32) *GetDeviceConfigsResponseBodyData {
	s.AutoUpdate = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetCustomIdleAction(v int32) *GetDeviceConfigsResponseBodyData {
	s.CustomIdleAction = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetCustomPowerOn(v int32) *GetDeviceConfigsResponseBodyData {
	s.CustomPowerOn = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetCustomResourcePackage(v *GetDeviceConfigsResponseBodyDataCustomResourcePackage) *GetDeviceConfigsResponseBodyData {
	s.CustomResourcePackage = v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetDefinePowerButton(v int32) *GetDeviceConfigsResponseBodyData {
	s.DefinePowerButton = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetDeviceLock(v int32) *GetDeviceConfigsResponseBodyData {
	s.DeviceLock = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetDisplayLayout(v string) *GetDeviceConfigsResponseBodyData {
	s.DisplayLayout = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetDisplayResolution(v string) *GetDeviceConfigsResponseBodyData {
	s.DisplayResolution = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetDisplayScaleRatio(v string) *GetDeviceConfigsResponseBodyData {
	s.DisplayScaleRatio = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetEnableAdb(v int32) *GetDeviceConfigsResponseBodyData {
	s.EnableAdb = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetEnableAutoLockScreen(v int32) *GetDeviceConfigsResponseBodyData {
	s.EnableAutoLockScreen = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetEnableBluetooth(v int32) *GetDeviceConfigsResponseBodyData {
	s.EnableBluetooth = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetEnableLockScreenPassword(v int32) *GetDeviceConfigsResponseBodyData {
	s.EnableLockScreenPassword = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetEnableModifyPassword(v int32) *GetDeviceConfigsResponseBodyData {
	s.EnableModifyPassword = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetEnableScheduledPowerOff(v int32) *GetDeviceConfigsResponseBodyData {
	s.EnableScheduledPowerOff = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetEnableUnlockPassword(v int32) *GetDeviceConfigsResponseBodyData {
	s.EnableUnlockPassword = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetEnableWlan(v int32) *GetDeviceConfigsResponseBodyData {
	s.EnableWlan = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetIdleTime(v int32) *GetDeviceConfigsResponseBodyData {
	s.IdleTime = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetLocalUsbPrint(v int32) *GetDeviceConfigsResponseBodyData {
	s.LocalUsbPrint = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetLockPassword(v string) *GetDeviceConfigsResponseBodyData {
	s.LockPassword = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetScheduledPowerOff(v string) *GetDeviceConfigsResponseBodyData {
	s.ScheduledPowerOff = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetSecureNetworkType(v string) *GetDeviceConfigsResponseBodyData {
	s.SecureNetworkType = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetSerialNo(v string) *GetDeviceConfigsResponseBodyData {
	s.SerialNo = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetSleepTime(v int32) *GetDeviceConfigsResponseBodyData {
	s.SleepTime = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetUrcl(v string) *GetDeviceConfigsResponseBodyData {
	s.Urcl = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetUsbStorage(v int32) *GetDeviceConfigsResponseBodyData {
	s.UsbStorage = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetUserCustomId(v string) *GetDeviceConfigsResponseBodyData {
	s.UserCustomId = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) SetUuid(v string) *GetDeviceConfigsResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDeviceConfigsResponseBodyDataCustomResourcePackage struct {
	ConfigAboutLogo     *string `json:"ConfigAboutLogo,omitempty" xml:"ConfigAboutLogo,omitempty"`
	DesktopWallpaper    *string `json:"DesktopWallpaper,omitempty" xml:"DesktopWallpaper,omitempty"`
	LoginPageBackground *string `json:"LoginPageBackground,omitempty" xml:"LoginPageBackground,omitempty"`
	LoginPageLogo       *string `json:"LoginPageLogo,omitempty" xml:"LoginPageLogo,omitempty"`
	PersonalCenterLogo  *string `json:"PersonalCenterLogo,omitempty" xml:"PersonalCenterLogo,omitempty"`
	StartLogo           *string `json:"StartLogo,omitempty" xml:"StartLogo,omitempty"`
	StartMenuLogo       *string `json:"StartMenuLogo,omitempty" xml:"StartMenuLogo,omitempty"`
	UpgradeLogo         *string `json:"UpgradeLogo,omitempty" xml:"UpgradeLogo,omitempty"`
}

func (s GetDeviceConfigsResponseBodyDataCustomResourcePackage) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceConfigsResponseBodyDataCustomResourcePackage) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) GetConfigAboutLogo() *string {
	return s.ConfigAboutLogo
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) GetDesktopWallpaper() *string {
	return s.DesktopWallpaper
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) GetLoginPageBackground() *string {
	return s.LoginPageBackground
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) GetLoginPageLogo() *string {
	return s.LoginPageLogo
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) GetPersonalCenterLogo() *string {
	return s.PersonalCenterLogo
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) GetStartLogo() *string {
	return s.StartLogo
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) GetStartMenuLogo() *string {
	return s.StartMenuLogo
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) GetUpgradeLogo() *string {
	return s.UpgradeLogo
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) SetConfigAboutLogo(v string) *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	s.ConfigAboutLogo = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) SetDesktopWallpaper(v string) *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	s.DesktopWallpaper = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) SetLoginPageBackground(v string) *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	s.LoginPageBackground = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) SetLoginPageLogo(v string) *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	s.LoginPageLogo = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) SetPersonalCenterLogo(v string) *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	s.PersonalCenterLogo = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) SetStartLogo(v string) *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	s.StartLogo = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) SetStartMenuLogo(v string) *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	s.StartMenuLogo = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) SetUpgradeLogo(v string) *GetDeviceConfigsResponseBodyDataCustomResourcePackage {
	s.UpgradeLogo = &v
	return s
}

func (s *GetDeviceConfigsResponseBodyDataCustomResourcePackage) Validate() error {
	return dara.Validate(s)
}
