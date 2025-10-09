// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNacosMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetNacosMcpServerResponseBodyData) *GetNacosMcpServerResponseBody
	GetData() *GetNacosMcpServerResponseBodyData
	SetRequestId(v string) *GetNacosMcpServerResponseBody
	GetRequestId() *string
}

type GetNacosMcpServerResponseBody struct {
	Data *GetNacosMcpServerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 72FC625E-9629-591B-9C01-3F0BFDAB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNacosMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNacosMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *GetNacosMcpServerResponseBody) GetData() *GetNacosMcpServerResponseBodyData {
	return s.Data
}

func (s *GetNacosMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNacosMcpServerResponseBody) SetData(v *GetNacosMcpServerResponseBodyData) *GetNacosMcpServerResponseBody {
	s.Data = v
	return s
}

func (s *GetNacosMcpServerResponseBody) SetRequestId(v string) *GetNacosMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNacosMcpServerResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetNacosMcpServerResponseBodyData struct {
	AllVersions      []*GetNacosMcpServerResponseBodyDataAllVersions      `json:"AllVersions,omitempty" xml:"AllVersions,omitempty" type:"Repeated"`
	BackendEndpoints []*GetNacosMcpServerResponseBodyDataBackendEndpoints `json:"BackendEndpoints,omitempty" xml:"BackendEndpoints,omitempty" type:"Repeated"`
	Capabilities     []*string                                            `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Repeated"`
	Description      *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// mcp-sse
	FrontProtocol *string `json:"FrontProtocol,omitempty" xml:"FrontProtocol,omitempty"`
	// ID。
	//
	// example:
	//
	// 5e3c5211-d365-4013-8540-c4106ec2a5dc
	Id                *string                `json:"Id,omitempty" xml:"Id,omitempty"`
	LocalServerConfig map[string]interface{} `json:"LocalServerConfig,omitempty" xml:"LocalServerConfig,omitempty"`
	// example:
	//
	// mcp-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// mcp-sse
	Protocol           *string                                              `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RemoteServerConfig *GetNacosMcpServerResponseBodyDataRemoteServerConfig `json:"RemoteServerConfig,omitempty" xml:"RemoteServerConfig,omitempty" type:"Struct"`
	ToolSpec           *GetNacosMcpServerResponseBodyDataToolSpec           `json:"ToolSpec,omitempty" xml:"ToolSpec,omitempty" type:"Struct"`
	VersionDetail      *GetNacosMcpServerResponseBodyDataVersionDetail      `json:"VersionDetail,omitempty" xml:"VersionDetail,omitempty" type:"Struct"`
	// example:
	//
	// allowTools:
	//
	// - demo-tool
	//
	// securityScheme: {}
	//
	// server:
	//
	//   name: mcp-demo
	//
	// tools:
	//
	// - args:
	//
	//   - name: name
	//
	//     description: name
	//
	//     type: string
	//
	//   description: a demo tool
	//
	//   name: demo-tool
	YamlConfig *string `json:"YamlConfig,omitempty" xml:"YamlConfig,omitempty"`
}

func (s GetNacosMcpServerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNacosMcpServerResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNacosMcpServerResponseBodyData) GetAllVersions() []*GetNacosMcpServerResponseBodyDataAllVersions {
	return s.AllVersions
}

func (s *GetNacosMcpServerResponseBodyData) GetBackendEndpoints() []*GetNacosMcpServerResponseBodyDataBackendEndpoints {
	return s.BackendEndpoints
}

func (s *GetNacosMcpServerResponseBodyData) GetCapabilities() []*string {
	return s.Capabilities
}

func (s *GetNacosMcpServerResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetNacosMcpServerResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetNacosMcpServerResponseBodyData) GetFrontProtocol() *string {
	return s.FrontProtocol
}

func (s *GetNacosMcpServerResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetNacosMcpServerResponseBodyData) GetLocalServerConfig() map[string]interface{} {
	return s.LocalServerConfig
}

func (s *GetNacosMcpServerResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetNacosMcpServerResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *GetNacosMcpServerResponseBodyData) GetRemoteServerConfig() *GetNacosMcpServerResponseBodyDataRemoteServerConfig {
	return s.RemoteServerConfig
}

func (s *GetNacosMcpServerResponseBodyData) GetToolSpec() *GetNacosMcpServerResponseBodyDataToolSpec {
	return s.ToolSpec
}

func (s *GetNacosMcpServerResponseBodyData) GetVersionDetail() *GetNacosMcpServerResponseBodyDataVersionDetail {
	return s.VersionDetail
}

func (s *GetNacosMcpServerResponseBodyData) GetYamlConfig() *string {
	return s.YamlConfig
}

func (s *GetNacosMcpServerResponseBodyData) SetAllVersions(v []*GetNacosMcpServerResponseBodyDataAllVersions) *GetNacosMcpServerResponseBodyData {
	s.AllVersions = v
	return s
}

func (s *GetNacosMcpServerResponseBodyData) SetBackendEndpoints(v []*GetNacosMcpServerResponseBodyDataBackendEndpoints) *GetNacosMcpServerResponseBodyData {
	s.BackendEndpoints = v
	return s
}

func (s *GetNacosMcpServerResponseBodyData) SetCapabilities(v []*string) *GetNacosMcpServerResponseBodyData {
	s.Capabilities = v
	return s
}

func (s *GetNacosMcpServerResponseBodyData) SetDescription(v string) *GetNacosMcpServerResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyData) SetEnabled(v bool) *GetNacosMcpServerResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyData) SetFrontProtocol(v string) *GetNacosMcpServerResponseBodyData {
	s.FrontProtocol = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyData) SetId(v string) *GetNacosMcpServerResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyData) SetLocalServerConfig(v map[string]interface{}) *GetNacosMcpServerResponseBodyData {
	s.LocalServerConfig = v
	return s
}

func (s *GetNacosMcpServerResponseBodyData) SetName(v string) *GetNacosMcpServerResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyData) SetProtocol(v string) *GetNacosMcpServerResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyData) SetRemoteServerConfig(v *GetNacosMcpServerResponseBodyDataRemoteServerConfig) *GetNacosMcpServerResponseBodyData {
	s.RemoteServerConfig = v
	return s
}

func (s *GetNacosMcpServerResponseBodyData) SetToolSpec(v *GetNacosMcpServerResponseBodyDataToolSpec) *GetNacosMcpServerResponseBodyData {
	s.ToolSpec = v
	return s
}

func (s *GetNacosMcpServerResponseBodyData) SetVersionDetail(v *GetNacosMcpServerResponseBodyDataVersionDetail) *GetNacosMcpServerResponseBodyData {
	s.VersionDetail = v
	return s
}

func (s *GetNacosMcpServerResponseBodyData) SetYamlConfig(v string) *GetNacosMcpServerResponseBodyData {
	s.YamlConfig = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetNacosMcpServerResponseBodyDataAllVersions struct {
	// example:
	//
	// true
	IsLatest *bool `json:"Is_latest,omitempty" xml:"Is_latest,omitempty"`
	// example:
	//
	// 2025-07-16
	ReleaseDate *string `json:"Release_date,omitempty" xml:"Release_date,omitempty"`
	// example:
	//
	// 1.0.9
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetNacosMcpServerResponseBodyDataAllVersions) String() string {
	return dara.Prettify(s)
}

func (s GetNacosMcpServerResponseBodyDataAllVersions) GoString() string {
	return s.String()
}

func (s *GetNacosMcpServerResponseBodyDataAllVersions) GetIsLatest() *bool {
	return s.IsLatest
}

func (s *GetNacosMcpServerResponseBodyDataAllVersions) GetReleaseDate() *string {
	return s.ReleaseDate
}

func (s *GetNacosMcpServerResponseBodyDataAllVersions) GetVersion() *string {
	return s.Version
}

func (s *GetNacosMcpServerResponseBodyDataAllVersions) SetIsLatest(v bool) *GetNacosMcpServerResponseBodyDataAllVersions {
	s.IsLatest = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataAllVersions) SetReleaseDate(v string) *GetNacosMcpServerResponseBodyDataAllVersions {
	s.ReleaseDate = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataAllVersions) SetVersion(v string) *GetNacosMcpServerResponseBodyDataAllVersions {
	s.Version = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataAllVersions) Validate() error {
	return dara.Validate(s)
}

type GetNacosMcpServerResponseBodyDataBackendEndpoints struct {
	// example:
	//
	// 127.0.0.1
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// /sse
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s GetNacosMcpServerResponseBodyDataBackendEndpoints) String() string {
	return dara.Prettify(s)
}

func (s GetNacosMcpServerResponseBodyDataBackendEndpoints) GoString() string {
	return s.String()
}

func (s *GetNacosMcpServerResponseBodyDataBackendEndpoints) GetAddress() *string {
	return s.Address
}

func (s *GetNacosMcpServerResponseBodyDataBackendEndpoints) GetPath() *string {
	return s.Path
}

func (s *GetNacosMcpServerResponseBodyDataBackendEndpoints) GetPort() *int32 {
	return s.Port
}

func (s *GetNacosMcpServerResponseBodyDataBackendEndpoints) SetAddress(v string) *GetNacosMcpServerResponseBodyDataBackendEndpoints {
	s.Address = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataBackendEndpoints) SetPath(v string) *GetNacosMcpServerResponseBodyDataBackendEndpoints {
	s.Path = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataBackendEndpoints) SetPort(v int32) *GetNacosMcpServerResponseBodyDataBackendEndpoints {
	s.Port = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataBackendEndpoints) Validate() error {
	return dara.Validate(s)
}

type GetNacosMcpServerResponseBodyDataRemoteServerConfig struct {
	// example:
	//
	// /sse
	ExportPath *string                                                        `json:"ExportPath,omitempty" xml:"ExportPath,omitempty"`
	ServiceRef *GetNacosMcpServerResponseBodyDataRemoteServerConfigServiceRef `json:"ServiceRef,omitempty" xml:"ServiceRef,omitempty" type:"Struct"`
}

func (s GetNacosMcpServerResponseBodyDataRemoteServerConfig) String() string {
	return dara.Prettify(s)
}

func (s GetNacosMcpServerResponseBodyDataRemoteServerConfig) GoString() string {
	return s.String()
}

func (s *GetNacosMcpServerResponseBodyDataRemoteServerConfig) GetExportPath() *string {
	return s.ExportPath
}

func (s *GetNacosMcpServerResponseBodyDataRemoteServerConfig) GetServiceRef() *GetNacosMcpServerResponseBodyDataRemoteServerConfigServiceRef {
	return s.ServiceRef
}

func (s *GetNacosMcpServerResponseBodyDataRemoteServerConfig) SetExportPath(v string) *GetNacosMcpServerResponseBodyDataRemoteServerConfig {
	s.ExportPath = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataRemoteServerConfig) SetServiceRef(v *GetNacosMcpServerResponseBodyDataRemoteServerConfigServiceRef) *GetNacosMcpServerResponseBodyDataRemoteServerConfig {
	s.ServiceRef = v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataRemoteServerConfig) Validate() error {
	return dara.Validate(s)
}

type GetNacosMcpServerResponseBodyDataRemoteServerConfigServiceRef struct {
	// example:
	//
	// DEFAULT
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// public
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// mcp-service
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetNacosMcpServerResponseBodyDataRemoteServerConfigServiceRef) String() string {
	return dara.Prettify(s)
}

func (s GetNacosMcpServerResponseBodyDataRemoteServerConfigServiceRef) GoString() string {
	return s.String()
}

func (s *GetNacosMcpServerResponseBodyDataRemoteServerConfigServiceRef) GetGroupName() *string {
	return s.GroupName
}

func (s *GetNacosMcpServerResponseBodyDataRemoteServerConfigServiceRef) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetNacosMcpServerResponseBodyDataRemoteServerConfigServiceRef) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetNacosMcpServerResponseBodyDataRemoteServerConfigServiceRef) SetGroupName(v string) *GetNacosMcpServerResponseBodyDataRemoteServerConfigServiceRef {
	s.GroupName = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataRemoteServerConfigServiceRef) SetNamespaceId(v string) *GetNacosMcpServerResponseBodyDataRemoteServerConfigServiceRef {
	s.NamespaceId = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataRemoteServerConfigServiceRef) SetServiceName(v string) *GetNacosMcpServerResponseBodyDataRemoteServerConfigServiceRef {
	s.ServiceName = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataRemoteServerConfigServiceRef) Validate() error {
	return dara.Validate(s)
}

type GetNacosMcpServerResponseBodyDataToolSpec struct {
	SecuritySchemes   interface{}                                       `json:"SecuritySchemes,omitempty" xml:"SecuritySchemes,omitempty"`
	SpecificationType *string                                           `json:"SpecificationType,omitempty" xml:"SpecificationType,omitempty"`
	Tools             []*GetNacosMcpServerResponseBodyDataToolSpecTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	ToolsMeta         map[string]*DataToolSpecToolsMetaValue            `json:"ToolsMeta,omitempty" xml:"ToolsMeta,omitempty"`
}

func (s GetNacosMcpServerResponseBodyDataToolSpec) String() string {
	return dara.Prettify(s)
}

func (s GetNacosMcpServerResponseBodyDataToolSpec) GoString() string {
	return s.String()
}

func (s *GetNacosMcpServerResponseBodyDataToolSpec) GetSecuritySchemes() interface{} {
	return s.SecuritySchemes
}

func (s *GetNacosMcpServerResponseBodyDataToolSpec) GetSpecificationType() *string {
	return s.SpecificationType
}

func (s *GetNacosMcpServerResponseBodyDataToolSpec) GetTools() []*GetNacosMcpServerResponseBodyDataToolSpecTools {
	return s.Tools
}

func (s *GetNacosMcpServerResponseBodyDataToolSpec) GetToolsMeta() map[string]*DataToolSpecToolsMetaValue {
	return s.ToolsMeta
}

func (s *GetNacosMcpServerResponseBodyDataToolSpec) SetSecuritySchemes(v interface{}) *GetNacosMcpServerResponseBodyDataToolSpec {
	s.SecuritySchemes = v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataToolSpec) SetSpecificationType(v string) *GetNacosMcpServerResponseBodyDataToolSpec {
	s.SpecificationType = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataToolSpec) SetTools(v []*GetNacosMcpServerResponseBodyDataToolSpecTools) *GetNacosMcpServerResponseBodyDataToolSpec {
	s.Tools = v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataToolSpec) SetToolsMeta(v map[string]*DataToolSpecToolsMetaValue) *GetNacosMcpServerResponseBodyDataToolSpec {
	s.ToolsMeta = v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataToolSpec) Validate() error {
	return dara.Validate(s)
}

type GetNacosMcpServerResponseBodyDataToolSpecTools struct {
	// example:
	//
	// a weather query tool
	Description *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	InputSchema map[string]interface{} `json:"InputSchema,omitempty" xml:"InputSchema,omitempty"`
	// example:
	//
	// get_weather
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetNacosMcpServerResponseBodyDataToolSpecTools) String() string {
	return dara.Prettify(s)
}

func (s GetNacosMcpServerResponseBodyDataToolSpecTools) GoString() string {
	return s.String()
}

func (s *GetNacosMcpServerResponseBodyDataToolSpecTools) GetDescription() *string {
	return s.Description
}

func (s *GetNacosMcpServerResponseBodyDataToolSpecTools) GetInputSchema() map[string]interface{} {
	return s.InputSchema
}

func (s *GetNacosMcpServerResponseBodyDataToolSpecTools) GetName() *string {
	return s.Name
}

func (s *GetNacosMcpServerResponseBodyDataToolSpecTools) SetDescription(v string) *GetNacosMcpServerResponseBodyDataToolSpecTools {
	s.Description = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataToolSpecTools) SetInputSchema(v map[string]interface{}) *GetNacosMcpServerResponseBodyDataToolSpecTools {
	s.InputSchema = v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataToolSpecTools) SetName(v string) *GetNacosMcpServerResponseBodyDataToolSpecTools {
	s.Name = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataToolSpecTools) Validate() error {
	return dara.Validate(s)
}

type GetNacosMcpServerResponseBodyDataVersionDetail struct {
	// example:
	//
	// true
	IsLatest *bool `json:"IsLatest,omitempty" xml:"IsLatest,omitempty"`
	// example:
	//
	// 2025-07-16
	ReleaseDate *string `json:"ReleaseDate,omitempty" xml:"ReleaseDate,omitempty"`
	// example:
	//
	// 1.11.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetNacosMcpServerResponseBodyDataVersionDetail) String() string {
	return dara.Prettify(s)
}

func (s GetNacosMcpServerResponseBodyDataVersionDetail) GoString() string {
	return s.String()
}

func (s *GetNacosMcpServerResponseBodyDataVersionDetail) GetIsLatest() *bool {
	return s.IsLatest
}

func (s *GetNacosMcpServerResponseBodyDataVersionDetail) GetReleaseDate() *string {
	return s.ReleaseDate
}

func (s *GetNacosMcpServerResponseBodyDataVersionDetail) GetVersion() *string {
	return s.Version
}

func (s *GetNacosMcpServerResponseBodyDataVersionDetail) SetIsLatest(v bool) *GetNacosMcpServerResponseBodyDataVersionDetail {
	s.IsLatest = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataVersionDetail) SetReleaseDate(v string) *GetNacosMcpServerResponseBodyDataVersionDetail {
	s.ReleaseDate = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataVersionDetail) SetVersion(v string) *GetNacosMcpServerResponseBodyDataVersionDetail {
	s.Version = &v
	return s
}

func (s *GetNacosMcpServerResponseBodyDataVersionDetail) Validate() error {
	return dara.Validate(s)
}
