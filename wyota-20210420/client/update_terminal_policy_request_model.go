// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTerminalPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowManualLockScreen(v int32) *UpdateTerminalPolicyRequest
	GetAllowManualLockScreen() *int32
	SetBackgroundModeTitle(v string) *UpdateTerminalPolicyRequest
	GetBackgroundModeTitle() *string
	SetCustomScreenCastRes(v bool) *UpdateTerminalPolicyRequest
	GetCustomScreenCastRes() *bool
	SetDisplayLayout(v string) *UpdateTerminalPolicyRequest
	GetDisplayLayout() *string
	SetDisplayResolution(v string) *UpdateTerminalPolicyRequest
	GetDisplayResolution() *string
	SetDisplayScaleRatio(v string) *UpdateTerminalPolicyRequest
	GetDisplayScaleRatio() *string
	SetEnableAutoLockScreen(v int32) *UpdateTerminalPolicyRequest
	GetEnableAutoLockScreen() *int32
	SetEnableAutoLogin(v int32) *UpdateTerminalPolicyRequest
	GetEnableAutoLogin() *int32
	SetEnableBackgroundMode(v int32) *UpdateTerminalPolicyRequest
	GetEnableBackgroundMode() *int32
	SetEnableBluetooth(v int32) *UpdateTerminalPolicyRequest
	GetEnableBluetooth() *int32
	SetEnableControlPanel(v int32) *UpdateTerminalPolicyRequest
	GetEnableControlPanel() *int32
	SetEnableImmersiveMode(v int32) *UpdateTerminalPolicyRequest
	GetEnableImmersiveMode() *int32
	SetEnableLockScreenHotKey(v int32) *UpdateTerminalPolicyRequest
	GetEnableLockScreenHotKey() *int32
	SetEnableModifyPassword(v int32) *UpdateTerminalPolicyRequest
	GetEnableModifyPassword() *int32
	SetEnableScanLogin(v int32) *UpdateTerminalPolicyRequest
	GetEnableScanLogin() *int32
	SetEnableScheduledReboot(v int32) *UpdateTerminalPolicyRequest
	GetEnableScheduledReboot() *int32
	SetEnableScheduledShutdown(v int32) *UpdateTerminalPolicyRequest
	GetEnableScheduledShutdown() *int32
	SetEnableSmsLogin(v int32) *UpdateTerminalPolicyRequest
	GetEnableSmsLogin() *int32
	SetEnableSwitchPersonal(v int32) *UpdateTerminalPolicyRequest
	GetEnableSwitchPersonal() *int32
	SetEnableWlan(v int32) *UpdateTerminalPolicyRequest
	GetEnableWlan() *int32
	SetFollowCloudReboot(v int32) *UpdateTerminalPolicyRequest
	GetFollowCloudReboot() *int32
	SetFollowCloudShutdown(v int32) *UpdateTerminalPolicyRequest
	GetFollowCloudShutdown() *int32
	SetFollowTerminalReboot(v int32) *UpdateTerminalPolicyRequest
	GetFollowTerminalReboot() *int32
	SetFollowTerminalShutdown(v int32) *UpdateTerminalPolicyRequest
	GetFollowTerminalShutdown() *int32
	SetForceSetPinCode(v int32) *UpdateTerminalPolicyRequest
	GetForceSetPinCode() *int32
	SetIdleTimeout(v int32) *UpdateTerminalPolicyRequest
	GetIdleTimeout() *int32
	SetIdleTimeoutAction(v int32) *UpdateTerminalPolicyRequest
	GetIdleTimeoutAction() *int32
	SetLockScreenPasswordRequired(v int32) *UpdateTerminalPolicyRequest
	GetLockScreenPasswordRequired() *int32
	SetLockScreenTimeout(v int32) *UpdateTerminalPolicyRequest
	GetLockScreenTimeout() *int32
	SetMainBizType(v string) *UpdateTerminalPolicyRequest
	GetMainBizType() *string
	SetName(v string) *UpdateTerminalPolicyRequest
	GetName() *string
	SetPowerButtonDefine(v int32) *UpdateTerminalPolicyRequest
	GetPowerButtonDefine() *int32
	SetPowerButtonDefineForAs(v int32) *UpdateTerminalPolicyRequest
	GetPowerButtonDefineForAs() *int32
	SetPowerButtonDefineForNs(v int32) *UpdateTerminalPolicyRequest
	GetPowerButtonDefineForNs() *int32
	SetPowerOnBehavior(v int32) *UpdateTerminalPolicyRequest
	GetPowerOnBehavior() *int32
	SetRunningMode(v string) *UpdateTerminalPolicyRequest
	GetRunningMode() *string
	SetScheduledReboot(v string) *UpdateTerminalPolicyRequest
	GetScheduledReboot() *string
	SetScheduledShutdown(v string) *UpdateTerminalPolicyRequest
	GetScheduledShutdown() *string
	SetScreenCastResPaths(v []*string) *UpdateTerminalPolicyRequest
	GetScreenCastResPaths() []*string
	SetSettingLock(v int32) *UpdateTerminalPolicyRequest
	GetSettingLock() *int32
	SetTerminalPolicyId(v string) *UpdateTerminalPolicyRequest
	GetTerminalPolicyId() *string
	SetUnlockMethod(v int32) *UpdateTerminalPolicyRequest
	GetUnlockMethod() *int32
}

type UpdateTerminalPolicyRequest struct {
	AllowManualLockScreen      *int32    `json:"AllowManualLockScreen,omitempty" xml:"AllowManualLockScreen,omitempty"`
	BackgroundModeTitle        *string   `json:"BackgroundModeTitle,omitempty" xml:"BackgroundModeTitle,omitempty"`
	CustomScreenCastRes        *bool     `json:"CustomScreenCastRes,omitempty" xml:"CustomScreenCastRes,omitempty"`
	DisplayLayout              *string   `json:"DisplayLayout,omitempty" xml:"DisplayLayout,omitempty"`
	DisplayResolution          *string   `json:"DisplayResolution,omitempty" xml:"DisplayResolution,omitempty"`
	DisplayScaleRatio          *string   `json:"DisplayScaleRatio,omitempty" xml:"DisplayScaleRatio,omitempty"`
	EnableAutoLockScreen       *int32    `json:"EnableAutoLockScreen,omitempty" xml:"EnableAutoLockScreen,omitempty"`
	EnableAutoLogin            *int32    `json:"EnableAutoLogin,omitempty" xml:"EnableAutoLogin,omitempty"`
	EnableBackgroundMode       *int32    `json:"EnableBackgroundMode,omitempty" xml:"EnableBackgroundMode,omitempty"`
	EnableBluetooth            *int32    `json:"EnableBluetooth,omitempty" xml:"EnableBluetooth,omitempty"`
	EnableControlPanel         *int32    `json:"EnableControlPanel,omitempty" xml:"EnableControlPanel,omitempty"`
	EnableImmersiveMode        *int32    `json:"EnableImmersiveMode,omitempty" xml:"EnableImmersiveMode,omitempty"`
	EnableLockScreenHotKey     *int32    `json:"EnableLockScreenHotKey,omitempty" xml:"EnableLockScreenHotKey,omitempty"`
	EnableModifyPassword       *int32    `json:"EnableModifyPassword,omitempty" xml:"EnableModifyPassword,omitempty"`
	EnableScanLogin            *int32    `json:"EnableScanLogin,omitempty" xml:"EnableScanLogin,omitempty"`
	EnableScheduledReboot      *int32    `json:"EnableScheduledReboot,omitempty" xml:"EnableScheduledReboot,omitempty"`
	EnableScheduledShutdown    *int32    `json:"EnableScheduledShutdown,omitempty" xml:"EnableScheduledShutdown,omitempty"`
	EnableSmsLogin             *int32    `json:"EnableSmsLogin,omitempty" xml:"EnableSmsLogin,omitempty"`
	EnableSwitchPersonal       *int32    `json:"EnableSwitchPersonal,omitempty" xml:"EnableSwitchPersonal,omitempty"`
	EnableWlan                 *int32    `json:"EnableWlan,omitempty" xml:"EnableWlan,omitempty"`
	FollowCloudReboot          *int32    `json:"FollowCloudReboot,omitempty" xml:"FollowCloudReboot,omitempty"`
	FollowCloudShutdown        *int32    `json:"FollowCloudShutdown,omitempty" xml:"FollowCloudShutdown,omitempty"`
	FollowTerminalReboot       *int32    `json:"FollowTerminalReboot,omitempty" xml:"FollowTerminalReboot,omitempty"`
	FollowTerminalShutdown     *int32    `json:"FollowTerminalShutdown,omitempty" xml:"FollowTerminalShutdown,omitempty"`
	ForceSetPinCode            *int32    `json:"ForceSetPinCode,omitempty" xml:"ForceSetPinCode,omitempty"`
	IdleTimeout                *int32    `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	IdleTimeoutAction          *int32    `json:"IdleTimeoutAction,omitempty" xml:"IdleTimeoutAction,omitempty"`
	LockScreenPasswordRequired *int32    `json:"LockScreenPasswordRequired,omitempty" xml:"LockScreenPasswordRequired,omitempty"`
	LockScreenTimeout          *int32    `json:"LockScreenTimeout,omitempty" xml:"LockScreenTimeout,omitempty"`
	MainBizType                *string   `json:"MainBizType,omitempty" xml:"MainBizType,omitempty"`
	Name                       *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PowerButtonDefine          *int32    `json:"PowerButtonDefine,omitempty" xml:"PowerButtonDefine,omitempty"`
	PowerButtonDefineForAs     *int32    `json:"PowerButtonDefineForAs,omitempty" xml:"PowerButtonDefineForAs,omitempty"`
	PowerButtonDefineForNs     *int32    `json:"PowerButtonDefineForNs,omitempty" xml:"PowerButtonDefineForNs,omitempty"`
	PowerOnBehavior            *int32    `json:"PowerOnBehavior,omitempty" xml:"PowerOnBehavior,omitempty"`
	RunningMode                *string   `json:"RunningMode,omitempty" xml:"RunningMode,omitempty"`
	ScheduledReboot            *string   `json:"ScheduledReboot,omitempty" xml:"ScheduledReboot,omitempty"`
	ScheduledShutdown          *string   `json:"ScheduledShutdown,omitempty" xml:"ScheduledShutdown,omitempty"`
	ScreenCastResPaths         []*string `json:"ScreenCastResPaths,omitempty" xml:"ScreenCastResPaths,omitempty" type:"Repeated"`
	SettingLock                *int32    `json:"SettingLock,omitempty" xml:"SettingLock,omitempty"`
	TerminalPolicyId           *string   `json:"TerminalPolicyId,omitempty" xml:"TerminalPolicyId,omitempty"`
	UnlockMethod               *int32    `json:"UnlockMethod,omitempty" xml:"UnlockMethod,omitempty"`
}

func (s UpdateTerminalPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTerminalPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateTerminalPolicyRequest) GetAllowManualLockScreen() *int32 {
	return s.AllowManualLockScreen
}

func (s *UpdateTerminalPolicyRequest) GetBackgroundModeTitle() *string {
	return s.BackgroundModeTitle
}

func (s *UpdateTerminalPolicyRequest) GetCustomScreenCastRes() *bool {
	return s.CustomScreenCastRes
}

func (s *UpdateTerminalPolicyRequest) GetDisplayLayout() *string {
	return s.DisplayLayout
}

func (s *UpdateTerminalPolicyRequest) GetDisplayResolution() *string {
	return s.DisplayResolution
}

func (s *UpdateTerminalPolicyRequest) GetDisplayScaleRatio() *string {
	return s.DisplayScaleRatio
}

func (s *UpdateTerminalPolicyRequest) GetEnableAutoLockScreen() *int32 {
	return s.EnableAutoLockScreen
}

func (s *UpdateTerminalPolicyRequest) GetEnableAutoLogin() *int32 {
	return s.EnableAutoLogin
}

func (s *UpdateTerminalPolicyRequest) GetEnableBackgroundMode() *int32 {
	return s.EnableBackgroundMode
}

func (s *UpdateTerminalPolicyRequest) GetEnableBluetooth() *int32 {
	return s.EnableBluetooth
}

func (s *UpdateTerminalPolicyRequest) GetEnableControlPanel() *int32 {
	return s.EnableControlPanel
}

func (s *UpdateTerminalPolicyRequest) GetEnableImmersiveMode() *int32 {
	return s.EnableImmersiveMode
}

func (s *UpdateTerminalPolicyRequest) GetEnableLockScreenHotKey() *int32 {
	return s.EnableLockScreenHotKey
}

func (s *UpdateTerminalPolicyRequest) GetEnableModifyPassword() *int32 {
	return s.EnableModifyPassword
}

func (s *UpdateTerminalPolicyRequest) GetEnableScanLogin() *int32 {
	return s.EnableScanLogin
}

func (s *UpdateTerminalPolicyRequest) GetEnableScheduledReboot() *int32 {
	return s.EnableScheduledReboot
}

func (s *UpdateTerminalPolicyRequest) GetEnableScheduledShutdown() *int32 {
	return s.EnableScheduledShutdown
}

func (s *UpdateTerminalPolicyRequest) GetEnableSmsLogin() *int32 {
	return s.EnableSmsLogin
}

func (s *UpdateTerminalPolicyRequest) GetEnableSwitchPersonal() *int32 {
	return s.EnableSwitchPersonal
}

func (s *UpdateTerminalPolicyRequest) GetEnableWlan() *int32 {
	return s.EnableWlan
}

func (s *UpdateTerminalPolicyRequest) GetFollowCloudReboot() *int32 {
	return s.FollowCloudReboot
}

func (s *UpdateTerminalPolicyRequest) GetFollowCloudShutdown() *int32 {
	return s.FollowCloudShutdown
}

func (s *UpdateTerminalPolicyRequest) GetFollowTerminalReboot() *int32 {
	return s.FollowTerminalReboot
}

func (s *UpdateTerminalPolicyRequest) GetFollowTerminalShutdown() *int32 {
	return s.FollowTerminalShutdown
}

func (s *UpdateTerminalPolicyRequest) GetForceSetPinCode() *int32 {
	return s.ForceSetPinCode
}

func (s *UpdateTerminalPolicyRequest) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *UpdateTerminalPolicyRequest) GetIdleTimeoutAction() *int32 {
	return s.IdleTimeoutAction
}

func (s *UpdateTerminalPolicyRequest) GetLockScreenPasswordRequired() *int32 {
	return s.LockScreenPasswordRequired
}

func (s *UpdateTerminalPolicyRequest) GetLockScreenTimeout() *int32 {
	return s.LockScreenTimeout
}

func (s *UpdateTerminalPolicyRequest) GetMainBizType() *string {
	return s.MainBizType
}

func (s *UpdateTerminalPolicyRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTerminalPolicyRequest) GetPowerButtonDefine() *int32 {
	return s.PowerButtonDefine
}

func (s *UpdateTerminalPolicyRequest) GetPowerButtonDefineForAs() *int32 {
	return s.PowerButtonDefineForAs
}

func (s *UpdateTerminalPolicyRequest) GetPowerButtonDefineForNs() *int32 {
	return s.PowerButtonDefineForNs
}

func (s *UpdateTerminalPolicyRequest) GetPowerOnBehavior() *int32 {
	return s.PowerOnBehavior
}

func (s *UpdateTerminalPolicyRequest) GetRunningMode() *string {
	return s.RunningMode
}

func (s *UpdateTerminalPolicyRequest) GetScheduledReboot() *string {
	return s.ScheduledReboot
}

func (s *UpdateTerminalPolicyRequest) GetScheduledShutdown() *string {
	return s.ScheduledShutdown
}

func (s *UpdateTerminalPolicyRequest) GetScreenCastResPaths() []*string {
	return s.ScreenCastResPaths
}

func (s *UpdateTerminalPolicyRequest) GetSettingLock() *int32 {
	return s.SettingLock
}

func (s *UpdateTerminalPolicyRequest) GetTerminalPolicyId() *string {
	return s.TerminalPolicyId
}

func (s *UpdateTerminalPolicyRequest) GetUnlockMethod() *int32 {
	return s.UnlockMethod
}

func (s *UpdateTerminalPolicyRequest) SetAllowManualLockScreen(v int32) *UpdateTerminalPolicyRequest {
	s.AllowManualLockScreen = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetBackgroundModeTitle(v string) *UpdateTerminalPolicyRequest {
	s.BackgroundModeTitle = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetCustomScreenCastRes(v bool) *UpdateTerminalPolicyRequest {
	s.CustomScreenCastRes = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetDisplayLayout(v string) *UpdateTerminalPolicyRequest {
	s.DisplayLayout = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetDisplayResolution(v string) *UpdateTerminalPolicyRequest {
	s.DisplayResolution = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetDisplayScaleRatio(v string) *UpdateTerminalPolicyRequest {
	s.DisplayScaleRatio = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableAutoLockScreen(v int32) *UpdateTerminalPolicyRequest {
	s.EnableAutoLockScreen = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableAutoLogin(v int32) *UpdateTerminalPolicyRequest {
	s.EnableAutoLogin = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableBackgroundMode(v int32) *UpdateTerminalPolicyRequest {
	s.EnableBackgroundMode = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableBluetooth(v int32) *UpdateTerminalPolicyRequest {
	s.EnableBluetooth = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableControlPanel(v int32) *UpdateTerminalPolicyRequest {
	s.EnableControlPanel = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableImmersiveMode(v int32) *UpdateTerminalPolicyRequest {
	s.EnableImmersiveMode = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableLockScreenHotKey(v int32) *UpdateTerminalPolicyRequest {
	s.EnableLockScreenHotKey = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableModifyPassword(v int32) *UpdateTerminalPolicyRequest {
	s.EnableModifyPassword = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableScanLogin(v int32) *UpdateTerminalPolicyRequest {
	s.EnableScanLogin = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableScheduledReboot(v int32) *UpdateTerminalPolicyRequest {
	s.EnableScheduledReboot = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableScheduledShutdown(v int32) *UpdateTerminalPolicyRequest {
	s.EnableScheduledShutdown = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableSmsLogin(v int32) *UpdateTerminalPolicyRequest {
	s.EnableSmsLogin = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableSwitchPersonal(v int32) *UpdateTerminalPolicyRequest {
	s.EnableSwitchPersonal = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetEnableWlan(v int32) *UpdateTerminalPolicyRequest {
	s.EnableWlan = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetFollowCloudReboot(v int32) *UpdateTerminalPolicyRequest {
	s.FollowCloudReboot = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetFollowCloudShutdown(v int32) *UpdateTerminalPolicyRequest {
	s.FollowCloudShutdown = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetFollowTerminalReboot(v int32) *UpdateTerminalPolicyRequest {
	s.FollowTerminalReboot = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetFollowTerminalShutdown(v int32) *UpdateTerminalPolicyRequest {
	s.FollowTerminalShutdown = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetForceSetPinCode(v int32) *UpdateTerminalPolicyRequest {
	s.ForceSetPinCode = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetIdleTimeout(v int32) *UpdateTerminalPolicyRequest {
	s.IdleTimeout = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetIdleTimeoutAction(v int32) *UpdateTerminalPolicyRequest {
	s.IdleTimeoutAction = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetLockScreenPasswordRequired(v int32) *UpdateTerminalPolicyRequest {
	s.LockScreenPasswordRequired = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetLockScreenTimeout(v int32) *UpdateTerminalPolicyRequest {
	s.LockScreenTimeout = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetMainBizType(v string) *UpdateTerminalPolicyRequest {
	s.MainBizType = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetName(v string) *UpdateTerminalPolicyRequest {
	s.Name = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetPowerButtonDefine(v int32) *UpdateTerminalPolicyRequest {
	s.PowerButtonDefine = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetPowerButtonDefineForAs(v int32) *UpdateTerminalPolicyRequest {
	s.PowerButtonDefineForAs = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetPowerButtonDefineForNs(v int32) *UpdateTerminalPolicyRequest {
	s.PowerButtonDefineForNs = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetPowerOnBehavior(v int32) *UpdateTerminalPolicyRequest {
	s.PowerOnBehavior = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetRunningMode(v string) *UpdateTerminalPolicyRequest {
	s.RunningMode = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetScheduledReboot(v string) *UpdateTerminalPolicyRequest {
	s.ScheduledReboot = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetScheduledShutdown(v string) *UpdateTerminalPolicyRequest {
	s.ScheduledShutdown = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetScreenCastResPaths(v []*string) *UpdateTerminalPolicyRequest {
	s.ScreenCastResPaths = v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetSettingLock(v int32) *UpdateTerminalPolicyRequest {
	s.SettingLock = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetTerminalPolicyId(v string) *UpdateTerminalPolicyRequest {
	s.TerminalPolicyId = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) SetUnlockMethod(v int32) *UpdateTerminalPolicyRequest {
	s.UnlockMethod = &v
	return s
}

func (s *UpdateTerminalPolicyRequest) Validate() error {
	return dara.Validate(s)
}
