// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartUserAppAsyncEnhanceInMsaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApkProtector(v bool) *StartUserAppAsyncEnhanceInMsaRequest
	GetApkProtector() *bool
	SetAppId(v string) *StartUserAppAsyncEnhanceInMsaRequest
	GetAppId() *string
	SetAssetsFileList(v string) *StartUserAppAsyncEnhanceInMsaRequest
	GetAssetsFileList() *string
	SetClasses(v string) *StartUserAppAsyncEnhanceInMsaRequest
	GetClasses() *string
	SetDalvikDebugger(v int32) *StartUserAppAsyncEnhanceInMsaRequest
	GetDalvikDebugger() *int32
	SetEmulatorEnvironment(v int32) *StartUserAppAsyncEnhanceInMsaRequest
	GetEmulatorEnvironment() *int32
	SetId(v int64) *StartUserAppAsyncEnhanceInMsaRequest
	GetId() *int64
	SetJavaHook(v int32) *StartUserAppAsyncEnhanceInMsaRequest
	GetJavaHook() *int32
	SetMemoryDump(v int32) *StartUserAppAsyncEnhanceInMsaRequest
	GetMemoryDump() *int32
	SetNativeDebugger(v int32) *StartUserAppAsyncEnhanceInMsaRequest
	GetNativeDebugger() *int32
	SetNativeHook(v int32) *StartUserAppAsyncEnhanceInMsaRequest
	GetNativeHook() *int32
	SetNewShieldConfig(v string) *StartUserAppAsyncEnhanceInMsaRequest
	GetNewShieldConfig() *string
	SetPackageTampered(v int32) *StartUserAppAsyncEnhanceInMsaRequest
	GetPackageTampered() *int32
	SetRoot(v int32) *StartUserAppAsyncEnhanceInMsaRequest
	GetRoot() *int32
	SetRunMode(v string) *StartUserAppAsyncEnhanceInMsaRequest
	GetRunMode() *string
	SetSoFileList(v string) *StartUserAppAsyncEnhanceInMsaRequest
	GetSoFileList() *string
	SetTaskType(v string) *StartUserAppAsyncEnhanceInMsaRequest
	GetTaskType() *string
	SetTenantId(v string) *StartUserAppAsyncEnhanceInMsaRequest
	GetTenantId() *string
	SetTotalSwitch(v bool) *StartUserAppAsyncEnhanceInMsaRequest
	GetTotalSwitch() *bool
	SetUseAShield(v bool) *StartUserAppAsyncEnhanceInMsaRequest
	GetUseAShield() *bool
	SetUseYShield(v bool) *StartUserAppAsyncEnhanceInMsaRequest
	GetUseYShield() *bool
	SetWorkspaceId(v string) *StartUserAppAsyncEnhanceInMsaRequest
	GetWorkspaceId() *string
}

type StartUserAppAsyncEnhanceInMsaRequest struct {
	ApkProtector *bool `json:"ApkProtector,omitempty" xml:"ApkProtector,omitempty"`
	// This parameter is required.
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AssetsFileList      *string `json:"AssetsFileList,omitempty" xml:"AssetsFileList,omitempty"`
	Classes             *string `json:"Classes,omitempty" xml:"Classes,omitempty"`
	DalvikDebugger      *int32  `json:"DalvikDebugger,omitempty" xml:"DalvikDebugger,omitempty"`
	EmulatorEnvironment *int32  `json:"EmulatorEnvironment,omitempty" xml:"EmulatorEnvironment,omitempty"`
	// This parameter is required.
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	JavaHook        *int32  `json:"JavaHook,omitempty" xml:"JavaHook,omitempty"`
	MemoryDump      *int32  `json:"MemoryDump,omitempty" xml:"MemoryDump,omitempty"`
	NativeDebugger  *int32  `json:"NativeDebugger,omitempty" xml:"NativeDebugger,omitempty"`
	NativeHook      *int32  `json:"NativeHook,omitempty" xml:"NativeHook,omitempty"`
	NewShieldConfig *string `json:"NewShieldConfig,omitempty" xml:"NewShieldConfig,omitempty"`
	PackageTampered *int32  `json:"PackageTampered,omitempty" xml:"PackageTampered,omitempty"`
	Root            *int32  `json:"Root,omitempty" xml:"Root,omitempty"`
	RunMode         *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	SoFileList      *string `json:"SoFileList,omitempty" xml:"SoFileList,omitempty"`
	TaskType        *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// This parameter is required.
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TotalSwitch *bool   `json:"TotalSwitch,omitempty" xml:"TotalSwitch,omitempty"`
	UseAShield  *bool   `json:"UseAShield,omitempty" xml:"UseAShield,omitempty"`
	UseYShield  *bool   `json:"UseYShield,omitempty" xml:"UseYShield,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s StartUserAppAsyncEnhanceInMsaRequest) String() string {
	return dara.Prettify(s)
}

func (s StartUserAppAsyncEnhanceInMsaRequest) GoString() string {
	return s.String()
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetApkProtector() *bool {
	return s.ApkProtector
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetAssetsFileList() *string {
	return s.AssetsFileList
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetClasses() *string {
	return s.Classes
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetDalvikDebugger() *int32 {
	return s.DalvikDebugger
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetEmulatorEnvironment() *int32 {
	return s.EmulatorEnvironment
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetId() *int64 {
	return s.Id
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetJavaHook() *int32 {
	return s.JavaHook
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetMemoryDump() *int32 {
	return s.MemoryDump
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetNativeDebugger() *int32 {
	return s.NativeDebugger
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetNativeHook() *int32 {
	return s.NativeHook
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetNewShieldConfig() *string {
	return s.NewShieldConfig
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetPackageTampered() *int32 {
	return s.PackageTampered
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetRoot() *int32 {
	return s.Root
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetRunMode() *string {
	return s.RunMode
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetSoFileList() *string {
	return s.SoFileList
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetTotalSwitch() *bool {
	return s.TotalSwitch
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetUseAShield() *bool {
	return s.UseAShield
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetUseYShield() *bool {
	return s.UseYShield
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetApkProtector(v bool) *StartUserAppAsyncEnhanceInMsaRequest {
	s.ApkProtector = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetAppId(v string) *StartUserAppAsyncEnhanceInMsaRequest {
	s.AppId = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetAssetsFileList(v string) *StartUserAppAsyncEnhanceInMsaRequest {
	s.AssetsFileList = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetClasses(v string) *StartUserAppAsyncEnhanceInMsaRequest {
	s.Classes = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetDalvikDebugger(v int32) *StartUserAppAsyncEnhanceInMsaRequest {
	s.DalvikDebugger = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetEmulatorEnvironment(v int32) *StartUserAppAsyncEnhanceInMsaRequest {
	s.EmulatorEnvironment = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetId(v int64) *StartUserAppAsyncEnhanceInMsaRequest {
	s.Id = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetJavaHook(v int32) *StartUserAppAsyncEnhanceInMsaRequest {
	s.JavaHook = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetMemoryDump(v int32) *StartUserAppAsyncEnhanceInMsaRequest {
	s.MemoryDump = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetNativeDebugger(v int32) *StartUserAppAsyncEnhanceInMsaRequest {
	s.NativeDebugger = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetNativeHook(v int32) *StartUserAppAsyncEnhanceInMsaRequest {
	s.NativeHook = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetNewShieldConfig(v string) *StartUserAppAsyncEnhanceInMsaRequest {
	s.NewShieldConfig = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetPackageTampered(v int32) *StartUserAppAsyncEnhanceInMsaRequest {
	s.PackageTampered = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetRoot(v int32) *StartUserAppAsyncEnhanceInMsaRequest {
	s.Root = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetRunMode(v string) *StartUserAppAsyncEnhanceInMsaRequest {
	s.RunMode = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetSoFileList(v string) *StartUserAppAsyncEnhanceInMsaRequest {
	s.SoFileList = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetTaskType(v string) *StartUserAppAsyncEnhanceInMsaRequest {
	s.TaskType = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetTenantId(v string) *StartUserAppAsyncEnhanceInMsaRequest {
	s.TenantId = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetTotalSwitch(v bool) *StartUserAppAsyncEnhanceInMsaRequest {
	s.TotalSwitch = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetUseAShield(v bool) *StartUserAppAsyncEnhanceInMsaRequest {
	s.UseAShield = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetUseYShield(v bool) *StartUserAppAsyncEnhanceInMsaRequest {
	s.UseYShield = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) SetWorkspaceId(v string) *StartUserAppAsyncEnhanceInMsaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaRequest) Validate() error {
	return dara.Validate(s)
}
