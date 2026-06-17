// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawPluginsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawPluginsResponseBody
	GetApplicationId() *string
	SetCode(v int32) *DescribePolarClawPluginsResponseBody
	GetCode() *int32
	SetDiagnostics(v []*DescribePolarClawPluginsResponseBodyDiagnostics) *DescribePolarClawPluginsResponseBody
	GetDiagnostics() []*DescribePolarClawPluginsResponseBodyDiagnostics
	SetMessage(v string) *DescribePolarClawPluginsResponseBody
	GetMessage() *string
	SetPlugins(v []*DescribePolarClawPluginsResponseBodyPlugins) *DescribePolarClawPluginsResponseBody
	GetPlugins() []*DescribePolarClawPluginsResponseBodyPlugins
	SetRequestId(v string) *DescribePolarClawPluginsResponseBody
	GetRequestId() *string
}

type DescribePolarClawPluginsResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// An array of diagnostic objects.
	Diagnostics []*DescribePolarClawPluginsResponseBodyDiagnostics `json:"Diagnostics,omitempty" xml:"Diagnostics,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// An array of plugin objects.
	Plugins []*DescribePolarClawPluginsResponseBodyPlugins `json:"Plugins,omitempty" xml:"Plugins,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2281C6C9-CBAB-1AFD-8400-670750CF6025_2212
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePolarClawPluginsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarClawPluginsResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawPluginsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribePolarClawPluginsResponseBody) GetDiagnostics() []*DescribePolarClawPluginsResponseBodyDiagnostics {
	return s.Diagnostics
}

func (s *DescribePolarClawPluginsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePolarClawPluginsResponseBody) GetPlugins() []*DescribePolarClawPluginsResponseBodyPlugins {
	return s.Plugins
}

func (s *DescribePolarClawPluginsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarClawPluginsResponseBody) SetApplicationId(v string) *DescribePolarClawPluginsResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBody) SetCode(v int32) *DescribePolarClawPluginsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBody) SetDiagnostics(v []*DescribePolarClawPluginsResponseBodyDiagnostics) *DescribePolarClawPluginsResponseBody {
	s.Diagnostics = v
	return s
}

func (s *DescribePolarClawPluginsResponseBody) SetMessage(v string) *DescribePolarClawPluginsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBody) SetPlugins(v []*DescribePolarClawPluginsResponseBodyPlugins) *DescribePolarClawPluginsResponseBody {
	s.Plugins = v
	return s
}

func (s *DescribePolarClawPluginsResponseBody) SetRequestId(v string) *DescribePolarClawPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBody) Validate() error {
	if s.Diagnostics != nil {
		for _, item := range s.Diagnostics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Plugins != nil {
		for _, item := range s.Plugins {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarClawPluginsResponseBodyDiagnostics struct {
	// The severity level. Valid values: `error` and `warn`.
	//
	// example:
	//
	// warn
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The detailed diagnostic message.
	//
	// example:
	//
	// loaded without install/load-path provenance; treat as untracked local code
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the associated plugin.
	//
	// example:
	//
	// openclaw-lark
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The source file path associated with the diagnostic.
	//
	// example:
	//
	// /home/node/.openclaw/extensions/openclaw-lark/index.js
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribePolarClawPluginsResponseBodyDiagnostics) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawPluginsResponseBodyDiagnostics) GoString() string {
	return s.String()
}

func (s *DescribePolarClawPluginsResponseBodyDiagnostics) GetLevel() *string {
	return s.Level
}

func (s *DescribePolarClawPluginsResponseBodyDiagnostics) GetMessage() *string {
	return s.Message
}

func (s *DescribePolarClawPluginsResponseBodyDiagnostics) GetPluginId() *string {
	return s.PluginId
}

func (s *DescribePolarClawPluginsResponseBodyDiagnostics) GetSource() *string {
	return s.Source
}

func (s *DescribePolarClawPluginsResponseBodyDiagnostics) SetLevel(v string) *DescribePolarClawPluginsResponseBodyDiagnostics {
	s.Level = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBodyDiagnostics) SetMessage(v string) *DescribePolarClawPluginsResponseBodyDiagnostics {
	s.Message = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBodyDiagnostics) SetPluginId(v string) *DescribePolarClawPluginsResponseBodyDiagnostics {
	s.PluginId = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBodyDiagnostics) SetSource(v string) *DescribePolarClawPluginsResponseBodyDiagnostics {
	s.Source = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBodyDiagnostics) Validate() error {
	return dara.Validate(s)
}

type DescribePolarClawPluginsResponseBodyPlugins struct {
	// A list of channel IDs, which can be empty.
	ChannelIds []*string `json:"ChannelIds,omitempty" xml:"ChannelIds,omitempty" type:"Repeated"`
	// The description of the plugin.
	//
	// example:
	//
	// Lark/Feishu channel plugin with im/doc/wiki/drive/task/calendar tools
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The error message, or `null` if no error occurred.
	//
	// example:
	//
	// null
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The format of the plugin, which can be an empty string.
	//
	// example:
	//
	// openclaw
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The plugin ID.
	//
	// example:
	//
	// openclaw-lark
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The display name of the plugin.
	//
	// example:
	//
	// Feishu
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The origin of the plugin. Valid values: `bundled`, `global`, and `user-install`.
	//
	// example:
	//
	// global
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// A list of provider IDs, which can be empty.
	ProviderIds []*string `json:"ProviderIds,omitempty" xml:"ProviderIds,omitempty" type:"Repeated"`
	// The file path to the plugin\\"s entry point.
	//
	// example:
	//
	// /home/node/.openclaw/extensions/openclaw-lark/index.js
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the plugin. Valid values: `loaded`, `disabled`, and `error`.
	//
	// example:
	//
	// loaded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version number of the plugin.
	//
	// example:
	//
	// 2026.4.7
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribePolarClawPluginsResponseBodyPlugins) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawPluginsResponseBodyPlugins) GoString() string {
	return s.String()
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) GetChannelIds() []*string {
	return s.ChannelIds
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) GetDescription() *string {
	return s.Description
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) GetError() *string {
	return s.Error
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) GetFormat() *string {
	return s.Format
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) GetId() *string {
	return s.Id
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) GetName() *string {
	return s.Name
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) GetOrigin() *string {
	return s.Origin
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) GetProviderIds() []*string {
	return s.ProviderIds
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) GetSource() *string {
	return s.Source
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) GetStatus() *string {
	return s.Status
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) GetVersion() *string {
	return s.Version
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) SetChannelIds(v []*string) *DescribePolarClawPluginsResponseBodyPlugins {
	s.ChannelIds = v
	return s
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) SetDescription(v string) *DescribePolarClawPluginsResponseBodyPlugins {
	s.Description = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) SetError(v string) *DescribePolarClawPluginsResponseBodyPlugins {
	s.Error = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) SetFormat(v string) *DescribePolarClawPluginsResponseBodyPlugins {
	s.Format = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) SetId(v string) *DescribePolarClawPluginsResponseBodyPlugins {
	s.Id = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) SetName(v string) *DescribePolarClawPluginsResponseBodyPlugins {
	s.Name = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) SetOrigin(v string) *DescribePolarClawPluginsResponseBodyPlugins {
	s.Origin = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) SetProviderIds(v []*string) *DescribePolarClawPluginsResponseBodyPlugins {
	s.ProviderIds = v
	return s
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) SetSource(v string) *DescribePolarClawPluginsResponseBodyPlugins {
	s.Source = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) SetStatus(v string) *DescribePolarClawPluginsResponseBodyPlugins {
	s.Status = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) SetVersion(v string) *DescribePolarClawPluginsResponseBodyPlugins {
	s.Version = &v
	return s
}

func (s *DescribePolarClawPluginsResponseBodyPlugins) Validate() error {
	return dara.Validate(s)
}
