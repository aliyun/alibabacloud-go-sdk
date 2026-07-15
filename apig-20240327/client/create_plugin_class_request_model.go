// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePluginClassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *CreatePluginClassRequest
	GetAlias() *string
	SetDescription(v string) *CreatePluginClassRequest
	GetDescription() *string
	SetExecutePriority(v int32) *CreatePluginClassRequest
	GetExecutePriority() *int32
	SetExecuteStage(v string) *CreatePluginClassRequest
	GetExecuteStage() *string
	SetName(v string) *CreatePluginClassRequest
	GetName() *string
	SetSupportedMinGatewayVersion(v string) *CreatePluginClassRequest
	GetSupportedMinGatewayVersion() *string
	SetVersion(v string) *CreatePluginClassRequest
	GetVersion() *string
	SetVersionDescription(v string) *CreatePluginClassRequest
	GetVersionDescription() *string
	SetWasmLanguage(v string) *CreatePluginClassRequest
	GetWasmLanguage() *string
	SetWasmUrl(v string) *CreatePluginClassRequest
	GetWasmUrl() *string
}

type CreatePluginClassRequest struct {
	// The alias of the plugin.
	//
	// example:
	//
	// My Wasm Plugin
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// The description of the plugin.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom authentication plugin for validating tokens in request headers
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution priority of the plugin.
	//
	// example:
	//
	// 100
	ExecutePriority *int32 `json:"executePriority,omitempty" xml:"executePriority,omitempty"`
	// The execution stage of the plugin.
	//
	// This parameter is required.
	//
	// example:
	//
	// AUTHN
	ExecuteStage *string `json:"executeStage,omitempty" xml:"executeStage,omitempty"`
	// The name of the plugin class.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-wasm-plugin
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The minimum gateway version that the plugin is compatible with.
	//
	// example:
	//
	// 2.0.0
	SupportedMinGatewayVersion *string `json:"supportedMinGatewayVersion,omitempty" xml:"supportedMinGatewayVersion,omitempty"`
	// The version number of the plugin.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The description of the current version.
	//
	// This parameter is required.
	//
	// example:
	//
	// Initial version with basic token validation
	VersionDescription *string `json:"versionDescription,omitempty" xml:"versionDescription,omitempty"`
	// The programming language used to develop the WASM plugin.
	//
	// This parameter is required.
	//
	// example:
	//
	// Rust
	WasmLanguage *string `json:"wasmLanguage,omitempty" xml:"wasmLanguage,omitempty"`
	// The download URL of the WASM plugin binary file.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/plugins/my-plugin.wasm
	WasmUrl *string `json:"wasmUrl,omitempty" xml:"wasmUrl,omitempty"`
}

func (s CreatePluginClassRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginClassRequest) GoString() string {
	return s.String()
}

func (s *CreatePluginClassRequest) GetAlias() *string {
	return s.Alias
}

func (s *CreatePluginClassRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePluginClassRequest) GetExecutePriority() *int32 {
	return s.ExecutePriority
}

func (s *CreatePluginClassRequest) GetExecuteStage() *string {
	return s.ExecuteStage
}

func (s *CreatePluginClassRequest) GetName() *string {
	return s.Name
}

func (s *CreatePluginClassRequest) GetSupportedMinGatewayVersion() *string {
	return s.SupportedMinGatewayVersion
}

func (s *CreatePluginClassRequest) GetVersion() *string {
	return s.Version
}

func (s *CreatePluginClassRequest) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *CreatePluginClassRequest) GetWasmLanguage() *string {
	return s.WasmLanguage
}

func (s *CreatePluginClassRequest) GetWasmUrl() *string {
	return s.WasmUrl
}

func (s *CreatePluginClassRequest) SetAlias(v string) *CreatePluginClassRequest {
	s.Alias = &v
	return s
}

func (s *CreatePluginClassRequest) SetDescription(v string) *CreatePluginClassRequest {
	s.Description = &v
	return s
}

func (s *CreatePluginClassRequest) SetExecutePriority(v int32) *CreatePluginClassRequest {
	s.ExecutePriority = &v
	return s
}

func (s *CreatePluginClassRequest) SetExecuteStage(v string) *CreatePluginClassRequest {
	s.ExecuteStage = &v
	return s
}

func (s *CreatePluginClassRequest) SetName(v string) *CreatePluginClassRequest {
	s.Name = &v
	return s
}

func (s *CreatePluginClassRequest) SetSupportedMinGatewayVersion(v string) *CreatePluginClassRequest {
	s.SupportedMinGatewayVersion = &v
	return s
}

func (s *CreatePluginClassRequest) SetVersion(v string) *CreatePluginClassRequest {
	s.Version = &v
	return s
}

func (s *CreatePluginClassRequest) SetVersionDescription(v string) *CreatePluginClassRequest {
	s.VersionDescription = &v
	return s
}

func (s *CreatePluginClassRequest) SetWasmLanguage(v string) *CreatePluginClassRequest {
	s.WasmLanguage = &v
	return s
}

func (s *CreatePluginClassRequest) SetWasmUrl(v string) *CreatePluginClassRequest {
	s.WasmUrl = &v
	return s
}

func (s *CreatePluginClassRequest) Validate() error {
	return dara.Validate(s)
}
