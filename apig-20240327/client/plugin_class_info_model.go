// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPluginClassInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *PluginClassInfo
	GetAlias() *string
	SetConfigExample(v string) *PluginClassInfo
	GetConfigExample() *string
	SetDescription(v string) *PluginClassInfo
	GetDescription() *string
	SetExecutePriority(v int32) *PluginClassInfo
	GetExecutePriority() *int32
	SetExecuteStage(v string) *PluginClassInfo
	GetExecuteStage() *string
	SetImageName(v string) *PluginClassInfo
	GetImageName() *string
	SetInnerPlugin(v bool) *PluginClassInfo
	GetInnerPlugin() *bool
	SetMode(v string) *PluginClassInfo
	GetMode() *string
	SetName(v string) *PluginClassInfo
	GetName() *string
	SetPluginClassId(v string) *PluginClassInfo
	GetPluginClassId() *string
	SetSource(v string) *PluginClassInfo
	GetSource() *string
	SetSupportedMinGatewayVersion(v string) *PluginClassInfo
	GetSupportedMinGatewayVersion() *string
	SetType(v string) *PluginClassInfo
	GetType() *string
	SetVersion(v string) *PluginClassInfo
	GetVersion() *string
	SetVersionDescription(v string) *PluginClassInfo
	GetVersionDescription() *string
	SetWasmLanguage(v string) *PluginClassInfo
	GetWasmLanguage() *string
	SetWasmUrl(v string) *PluginClassInfo
	GetWasmUrl() *string
}

type PluginClassInfo struct {
	// The plugin class alias.
	//
	// example:
	//
	// AI安全护栏
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// The configuration example (JSON string).
	//
	// example:
	//
	// {"checkRequest":true}
	ConfigExample *string `json:"configExample,omitempty" xml:"configExample,omitempty"`
	// The plugin class description.
	//
	// example:
	//
	// AI请求内容安全检测插件
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution priority.
	//
	// example:
	//
	// 100
	ExecutePriority *int32 `json:"executePriority,omitempty" xml:"executePriority,omitempty"`
	// The execution stage.
	//
	// example:
	//
	// AUTHN
	ExecuteStage *string `json:"executeStage,omitempty" xml:"executeStage,omitempty"`
	// The Wasm image name.
	//
	// example:
	//
	// higress-registry.tencentcloudcr.com/ai-security-guard
	ImageName *string `json:"imageName,omitempty" xml:"imageName,omitempty"`
	// Indicates whether the plugin is a built-in plugin.
	//
	// example:
	//
	// false
	InnerPlugin *bool `json:"innerPlugin,omitempty" xml:"innerPlugin,omitempty"`
	// The plugin running mode.
	//
	// example:
	//
	// Wasm
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The plugin class name.
	//
	// example:
	//
	// ai-security-guard
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The plugin class ID.
	//
	// example:
	//
	// cls-xxx
	PluginClassId *string `json:"pluginClassId,omitempty" xml:"pluginClassId,omitempty"`
	// The plugin source.
	//
	// example:
	//
	// HigressOfficial
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The minimum supported gateway DPI engine version.
	//
	// example:
	//
	// 2.1.11
	SupportedMinGatewayVersion *string `json:"supportedMinGatewayVersion,omitempty" xml:"supportedMinGatewayVersion,omitempty"`
	// The plugin type category.
	//
	// example:
	//
	// AI
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The plugin version number.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The version description.
	//
	// example:
	//
	// 初始版本
	VersionDescription *string `json:"versionDescription,omitempty" xml:"versionDescription,omitempty"`
	// The Wasm programming language.
	//
	// example:
	//
	// Rust
	WasmLanguage *string `json:"wasmLanguage,omitempty" xml:"wasmLanguage,omitempty"`
	// The Wasm file download URL.
	//
	// example:
	//
	// https://...
	WasmUrl *string `json:"wasmUrl,omitempty" xml:"wasmUrl,omitempty"`
}

func (s PluginClassInfo) String() string {
	return dara.Prettify(s)
}

func (s PluginClassInfo) GoString() string {
	return s.String()
}

func (s *PluginClassInfo) GetAlias() *string {
	return s.Alias
}

func (s *PluginClassInfo) GetConfigExample() *string {
	return s.ConfigExample
}

func (s *PluginClassInfo) GetDescription() *string {
	return s.Description
}

func (s *PluginClassInfo) GetExecutePriority() *int32 {
	return s.ExecutePriority
}

func (s *PluginClassInfo) GetExecuteStage() *string {
	return s.ExecuteStage
}

func (s *PluginClassInfo) GetImageName() *string {
	return s.ImageName
}

func (s *PluginClassInfo) GetInnerPlugin() *bool {
	return s.InnerPlugin
}

func (s *PluginClassInfo) GetMode() *string {
	return s.Mode
}

func (s *PluginClassInfo) GetName() *string {
	return s.Name
}

func (s *PluginClassInfo) GetPluginClassId() *string {
	return s.PluginClassId
}

func (s *PluginClassInfo) GetSource() *string {
	return s.Source
}

func (s *PluginClassInfo) GetSupportedMinGatewayVersion() *string {
	return s.SupportedMinGatewayVersion
}

func (s *PluginClassInfo) GetType() *string {
	return s.Type
}

func (s *PluginClassInfo) GetVersion() *string {
	return s.Version
}

func (s *PluginClassInfo) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *PluginClassInfo) GetWasmLanguage() *string {
	return s.WasmLanguage
}

func (s *PluginClassInfo) GetWasmUrl() *string {
	return s.WasmUrl
}

func (s *PluginClassInfo) SetAlias(v string) *PluginClassInfo {
	s.Alias = &v
	return s
}

func (s *PluginClassInfo) SetConfigExample(v string) *PluginClassInfo {
	s.ConfigExample = &v
	return s
}

func (s *PluginClassInfo) SetDescription(v string) *PluginClassInfo {
	s.Description = &v
	return s
}

func (s *PluginClassInfo) SetExecutePriority(v int32) *PluginClassInfo {
	s.ExecutePriority = &v
	return s
}

func (s *PluginClassInfo) SetExecuteStage(v string) *PluginClassInfo {
	s.ExecuteStage = &v
	return s
}

func (s *PluginClassInfo) SetImageName(v string) *PluginClassInfo {
	s.ImageName = &v
	return s
}

func (s *PluginClassInfo) SetInnerPlugin(v bool) *PluginClassInfo {
	s.InnerPlugin = &v
	return s
}

func (s *PluginClassInfo) SetMode(v string) *PluginClassInfo {
	s.Mode = &v
	return s
}

func (s *PluginClassInfo) SetName(v string) *PluginClassInfo {
	s.Name = &v
	return s
}

func (s *PluginClassInfo) SetPluginClassId(v string) *PluginClassInfo {
	s.PluginClassId = &v
	return s
}

func (s *PluginClassInfo) SetSource(v string) *PluginClassInfo {
	s.Source = &v
	return s
}

func (s *PluginClassInfo) SetSupportedMinGatewayVersion(v string) *PluginClassInfo {
	s.SupportedMinGatewayVersion = &v
	return s
}

func (s *PluginClassInfo) SetType(v string) *PluginClassInfo {
	s.Type = &v
	return s
}

func (s *PluginClassInfo) SetVersion(v string) *PluginClassInfo {
	s.Version = &v
	return s
}

func (s *PluginClassInfo) SetVersionDescription(v string) *PluginClassInfo {
	s.VersionDescription = &v
	return s
}

func (s *PluginClassInfo) SetWasmLanguage(v string) *PluginClassInfo {
	s.WasmLanguage = &v
	return s
}

func (s *PluginClassInfo) SetWasmUrl(v string) *PluginClassInfo {
	s.WasmUrl = &v
	return s
}

func (s *PluginClassInfo) Validate() error {
	return dara.Validate(s)
}
