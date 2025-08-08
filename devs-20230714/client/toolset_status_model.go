// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToolsetStatus interface {
	dara.Model
	String() string
	GoString() string
	SetObservedGeneration(v int64) *ToolsetStatus
	GetObservedGeneration() *int64
	SetObservedTime(v string) *ToolsetStatus
	GetObservedTime() *string
	SetOutputs(v *ToolsetStatusOutputs) *ToolsetStatus
	GetOutputs() *ToolsetStatusOutputs
	SetPhase(v string) *ToolsetStatus
	GetPhase() *string
}

type ToolsetStatus struct {
	ObservedGeneration *int64                `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
	ObservedTime       *string               `json:"observedTime,omitempty" xml:"observedTime,omitempty"`
	Outputs            *ToolsetStatusOutputs `json:"outputs,omitempty" xml:"outputs,omitempty" type:"Struct"`
	// example:
	//
	// Installed
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
}

func (s ToolsetStatus) String() string {
	return dara.Prettify(s)
}

func (s ToolsetStatus) GoString() string {
	return s.String()
}

func (s *ToolsetStatus) GetObservedGeneration() *int64 {
	return s.ObservedGeneration
}

func (s *ToolsetStatus) GetObservedTime() *string {
	return s.ObservedTime
}

func (s *ToolsetStatus) GetOutputs() *ToolsetStatusOutputs {
	return s.Outputs
}

func (s *ToolsetStatus) GetPhase() *string {
	return s.Phase
}

func (s *ToolsetStatus) SetObservedGeneration(v int64) *ToolsetStatus {
	s.ObservedGeneration = &v
	return s
}

func (s *ToolsetStatus) SetObservedTime(v string) *ToolsetStatus {
	s.ObservedTime = &v
	return s
}

func (s *ToolsetStatus) SetOutputs(v *ToolsetStatusOutputs) *ToolsetStatus {
	s.Outputs = v
	return s
}

func (s *ToolsetStatus) SetPhase(v string) *ToolsetStatus {
	s.Phase = &v
	return s
}

func (s *ToolsetStatus) Validate() error {
	return dara.Validate(s)
}

type ToolsetStatusOutputs struct {
	FunctionArn     *string                              `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	McpServerConfig *ToolsetStatusOutputsMcpServerConfig `json:"mcpServerConfig,omitempty" xml:"mcpServerConfig,omitempty" type:"Struct"`
	OpenApiTools    []*OpenAPIToolMeta                   `json:"openApiTools,omitempty" xml:"openApiTools,omitempty" type:"Repeated"`
	Tools           []*MCPToolMeta                       `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	Urls            *ToolsetStatusOutputsUrls            `json:"urls,omitempty" xml:"urls,omitempty" type:"Struct"`
}

func (s ToolsetStatusOutputs) String() string {
	return dara.Prettify(s)
}

func (s ToolsetStatusOutputs) GoString() string {
	return s.String()
}

func (s *ToolsetStatusOutputs) GetFunctionArn() *string {
	return s.FunctionArn
}

func (s *ToolsetStatusOutputs) GetMcpServerConfig() *ToolsetStatusOutputsMcpServerConfig {
	return s.McpServerConfig
}

func (s *ToolsetStatusOutputs) GetOpenApiTools() []*OpenAPIToolMeta {
	return s.OpenApiTools
}

func (s *ToolsetStatusOutputs) GetTools() []*MCPToolMeta {
	return s.Tools
}

func (s *ToolsetStatusOutputs) GetUrls() *ToolsetStatusOutputsUrls {
	return s.Urls
}

func (s *ToolsetStatusOutputs) SetFunctionArn(v string) *ToolsetStatusOutputs {
	s.FunctionArn = &v
	return s
}

func (s *ToolsetStatusOutputs) SetMcpServerConfig(v *ToolsetStatusOutputsMcpServerConfig) *ToolsetStatusOutputs {
	s.McpServerConfig = v
	return s
}

func (s *ToolsetStatusOutputs) SetOpenApiTools(v []*OpenAPIToolMeta) *ToolsetStatusOutputs {
	s.OpenApiTools = v
	return s
}

func (s *ToolsetStatusOutputs) SetTools(v []*MCPToolMeta) *ToolsetStatusOutputs {
	s.Tools = v
	return s
}

func (s *ToolsetStatusOutputs) SetUrls(v *ToolsetStatusOutputsUrls) *ToolsetStatusOutputs {
	s.Urls = v
	return s
}

func (s *ToolsetStatusOutputs) Validate() error {
	return dara.Validate(s)
}

type ToolsetStatusOutputsMcpServerConfig struct {
	Headers       map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	TransportType *string            `json:"transportType,omitempty" xml:"transportType,omitempty"`
	Url           *string            `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ToolsetStatusOutputsMcpServerConfig) String() string {
	return dara.Prettify(s)
}

func (s ToolsetStatusOutputsMcpServerConfig) GoString() string {
	return s.String()
}

func (s *ToolsetStatusOutputsMcpServerConfig) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ToolsetStatusOutputsMcpServerConfig) GetTransportType() *string {
	return s.TransportType
}

func (s *ToolsetStatusOutputsMcpServerConfig) GetUrl() *string {
	return s.Url
}

func (s *ToolsetStatusOutputsMcpServerConfig) SetHeaders(v map[string]*string) *ToolsetStatusOutputsMcpServerConfig {
	s.Headers = v
	return s
}

func (s *ToolsetStatusOutputsMcpServerConfig) SetTransportType(v string) *ToolsetStatusOutputsMcpServerConfig {
	s.TransportType = &v
	return s
}

func (s *ToolsetStatusOutputsMcpServerConfig) SetUrl(v string) *ToolsetStatusOutputsMcpServerConfig {
	s.Url = &v
	return s
}

func (s *ToolsetStatusOutputsMcpServerConfig) Validate() error {
	return dara.Validate(s)
}

type ToolsetStatusOutputsUrls struct {
	InternetUrl *string `json:"internetUrl,omitempty" xml:"internetUrl,omitempty"`
	IntranetUrl *string `json:"intranetUrl,omitempty" xml:"intranetUrl,omitempty"`
}

func (s ToolsetStatusOutputsUrls) String() string {
	return dara.Prettify(s)
}

func (s ToolsetStatusOutputsUrls) GoString() string {
	return s.String()
}

func (s *ToolsetStatusOutputsUrls) GetInternetUrl() *string {
	return s.InternetUrl
}

func (s *ToolsetStatusOutputsUrls) GetIntranetUrl() *string {
	return s.IntranetUrl
}

func (s *ToolsetStatusOutputsUrls) SetInternetUrl(v string) *ToolsetStatusOutputsUrls {
	s.InternetUrl = &v
	return s
}

func (s *ToolsetStatusOutputsUrls) SetIntranetUrl(v string) *ToolsetStatusOutputsUrls {
	s.IntranetUrl = &v
	return s
}

func (s *ToolsetStatusOutputsUrls) Validate() error {
	return dara.Validate(s)
}
