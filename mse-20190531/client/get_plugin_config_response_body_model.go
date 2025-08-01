// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPluginConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetPluginConfigResponseBody
	GetCode() *int32
	SetData(v *GetPluginConfigResponseBodyData) *GetPluginConfigResponseBody
	GetData() *GetPluginConfigResponseBodyData
	SetDynamicCode(v string) *GetPluginConfigResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetPluginConfigResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *GetPluginConfigResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *GetPluginConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPluginConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPluginConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPluginConfigResponseBody
	GetSuccess() *bool
}

type GetPluginConfigResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetPluginConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// message
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// 500
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 03A3E2F4-6804-5663-9D5D-2EC47A1*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPluginConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPluginConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetPluginConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetPluginConfigResponseBody) GetData() *GetPluginConfigResponseBodyData {
	return s.Data
}

func (s *GetPluginConfigResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetPluginConfigResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetPluginConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPluginConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPluginConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPluginConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPluginConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPluginConfigResponseBody) SetCode(v int32) *GetPluginConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetPluginConfigResponseBody) SetData(v *GetPluginConfigResponseBodyData) *GetPluginConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetPluginConfigResponseBody) SetDynamicCode(v string) *GetPluginConfigResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetPluginConfigResponseBody) SetDynamicMessage(v string) *GetPluginConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetPluginConfigResponseBody) SetErrorCode(v string) *GetPluginConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPluginConfigResponseBody) SetHttpStatusCode(v int32) *GetPluginConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPluginConfigResponseBody) SetMessage(v string) *GetPluginConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetPluginConfigResponseBody) SetRequestId(v string) *GetPluginConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPluginConfigResponseBody) SetSuccess(v bool) *GetPluginConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetPluginConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPluginConfigResponseBodyData struct {
	// The category of the plug-in. Valid values:
	//
	// 0: user-defined
	//
	// 1: permission authentication
	//
	// 2: security protection
	//
	// 3: transmission protocol
	//
	// 4: traffic control
	//
	// 5: traffic observation
	//
	// example:
	//
	// 0
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The information about the plug-in configuration used for checking.
	//
	// example:
	//
	// \\# The configuration includes the fields required for checking, such as name, age, and friends. Sample configuration: name: John age: 18 friends: - David - Anne
	ConfigCheck   *string `json:"ConfigCheck,omitempty" xml:"ConfigCheck,omitempty"`
	ConfigExample *string `json:"ConfigExample,omitempty" xml:"ConfigExample,omitempty"`
	// example:
	//
	// 5
	DomainConfigStartIndex *int32 `json:"DomainConfigStartIndex,omitempty" xml:"DomainConfigStartIndex,omitempty"`
	// The list of gateway plug-in configurations.
	GatewayConfigList []*GetPluginConfigResponseBodyDataGatewayConfigList `json:"GatewayConfigList,omitempty" xml:"GatewayConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// 7
	GatewayConfigStartIndex *int32 `json:"GatewayConfigStartIndex,omitempty" xml:"GatewayConfigStartIndex,omitempty"`
	// The ID of the plug-in.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// name
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The mode.
	//
	// example:
	//
	// 0
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The name of the plug-in.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The execution stage of the plug-in. Valid values:
	//
	// 0: default stage
	//
	// 1: authorization stage
	//
	// 2: authentication stage
	//
	// 3: statistics stage
	//
	// example:
	//
	// 0
	Phase *int32 `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The ID of the creator.
	//
	// example:
	//
	// 123
	PrimaryUser *string `json:"PrimaryUser,omitempty" xml:"PrimaryUser,omitempty"`
	// The execution priority of the plug-in. A larger value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The publish status.
	//
	// example:
	//
	// 1
	PublishState *int32 `json:"PublishState,omitempty" xml:"PublishState,omitempty"`
	// The description of the README file.
	//
	// example:
	//
	// read me
	Readme *string `json:"Readme,omitempty" xml:"Readme,omitempty"`
	// The description of the README file that is edited in English.
	//
	// example:
	//
	// read me
	ReadmeEn *string `json:"ReadmeEn,omitempty" xml:"ReadmeEn,omitempty"`
	// example:
	//
	// 0
	RouteConfigStartIndex *int32 `json:"RouteConfigStartIndex,omitempty" xml:"RouteConfigStartIndex,omitempty"`
	// Indicates whether the plug-in is enabled. Valid values:
	//
	// 0: disabled
	//
	// 1: enabled
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The summary of the plug-in.
	//
	// example:
	//
	// This is a plug-in.
	Summary   *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	SummaryEn *string `json:"SummaryEn,omitempty" xml:"SummaryEn,omitempty"`
	// The type.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the plug-in.
	//
	// example:
	//
	// v1
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionJson *string `json:"VersionJson,omitempty" xml:"VersionJson,omitempty"`
	// The WebAssembly language. Valid values:
	//
	// 0: C++
	//
	// 1: TinyGo
	//
	// 2: Rust
	//
	// 3: AssemblyScript
	//
	// 4: Zig
	//
	// example:
	//
	// 0
	WasmLang *int32 `json:"WasmLang,omitempty" xml:"WasmLang,omitempty"`
}

func (s GetPluginConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPluginConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPluginConfigResponseBodyData) GetCategory() *int32 {
	return s.Category
}

func (s *GetPluginConfigResponseBodyData) GetConfigCheck() *string {
	return s.ConfigCheck
}

func (s *GetPluginConfigResponseBodyData) GetConfigExample() *string {
	return s.ConfigExample
}

func (s *GetPluginConfigResponseBodyData) GetDomainConfigStartIndex() *int32 {
	return s.DomainConfigStartIndex
}

func (s *GetPluginConfigResponseBodyData) GetGatewayConfigList() []*GetPluginConfigResponseBodyDataGatewayConfigList {
	return s.GatewayConfigList
}

func (s *GetPluginConfigResponseBodyData) GetGatewayConfigStartIndex() *int32 {
	return s.GatewayConfigStartIndex
}

func (s *GetPluginConfigResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetPluginConfigResponseBodyData) GetImageName() *string {
	return s.ImageName
}

func (s *GetPluginConfigResponseBodyData) GetMode() *int32 {
	return s.Mode
}

func (s *GetPluginConfigResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetPluginConfigResponseBodyData) GetPhase() *int32 {
	return s.Phase
}

func (s *GetPluginConfigResponseBodyData) GetPrimaryUser() *string {
	return s.PrimaryUser
}

func (s *GetPluginConfigResponseBodyData) GetPriority() *int32 {
	return s.Priority
}

func (s *GetPluginConfigResponseBodyData) GetPublishState() *int32 {
	return s.PublishState
}

func (s *GetPluginConfigResponseBodyData) GetReadme() *string {
	return s.Readme
}

func (s *GetPluginConfigResponseBodyData) GetReadmeEn() *string {
	return s.ReadmeEn
}

func (s *GetPluginConfigResponseBodyData) GetRouteConfigStartIndex() *int32 {
	return s.RouteConfigStartIndex
}

func (s *GetPluginConfigResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetPluginConfigResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *GetPluginConfigResponseBodyData) GetSummaryEn() *string {
	return s.SummaryEn
}

func (s *GetPluginConfigResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *GetPluginConfigResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetPluginConfigResponseBodyData) GetVersionJson() *string {
	return s.VersionJson
}

func (s *GetPluginConfigResponseBodyData) GetWasmLang() *int32 {
	return s.WasmLang
}

func (s *GetPluginConfigResponseBodyData) SetCategory(v int32) *GetPluginConfigResponseBodyData {
	s.Category = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetConfigCheck(v string) *GetPluginConfigResponseBodyData {
	s.ConfigCheck = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetConfigExample(v string) *GetPluginConfigResponseBodyData {
	s.ConfigExample = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetDomainConfigStartIndex(v int32) *GetPluginConfigResponseBodyData {
	s.DomainConfigStartIndex = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetGatewayConfigList(v []*GetPluginConfigResponseBodyDataGatewayConfigList) *GetPluginConfigResponseBodyData {
	s.GatewayConfigList = v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetGatewayConfigStartIndex(v int32) *GetPluginConfigResponseBodyData {
	s.GatewayConfigStartIndex = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetId(v int64) *GetPluginConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetImageName(v string) *GetPluginConfigResponseBodyData {
	s.ImageName = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetMode(v int32) *GetPluginConfigResponseBodyData {
	s.Mode = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetName(v string) *GetPluginConfigResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetPhase(v int32) *GetPluginConfigResponseBodyData {
	s.Phase = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetPrimaryUser(v string) *GetPluginConfigResponseBodyData {
	s.PrimaryUser = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetPriority(v int32) *GetPluginConfigResponseBodyData {
	s.Priority = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetPublishState(v int32) *GetPluginConfigResponseBodyData {
	s.PublishState = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetReadme(v string) *GetPluginConfigResponseBodyData {
	s.Readme = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetReadmeEn(v string) *GetPluginConfigResponseBodyData {
	s.ReadmeEn = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetRouteConfigStartIndex(v int32) *GetPluginConfigResponseBodyData {
	s.RouteConfigStartIndex = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetStatus(v string) *GetPluginConfigResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetSummary(v string) *GetPluginConfigResponseBodyData {
	s.Summary = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetSummaryEn(v string) *GetPluginConfigResponseBodyData {
	s.SummaryEn = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetType(v int32) *GetPluginConfigResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetVersion(v string) *GetPluginConfigResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetVersionJson(v string) *GetPluginConfigResponseBodyData {
	s.VersionJson = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) SetWasmLang(v int32) *GetPluginConfigResponseBodyData {
	s.WasmLang = &v
	return s
}

func (s *GetPluginConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetPluginConfigResponseBodyDataGatewayConfigList struct {
	// The plug-in configuration.
	//
	// example:
	//
	// \\# Configure a check for the required fields of the plug-in, such as name, age, and friends. Sample configuration: name: John age: 18 friends: - David - Anne
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The application scope of the plug-in. Valid values:
	//
	// 0: global
	//
	// 1: domain names
	//
	// 2: routes
	//
	// example:
	//
	// 0
	ConfigLevel *int32 `json:"ConfigLevel,omitempty" xml:"ConfigLevel,omitempty"`
	// Indicates whether the plug-in is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 1
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-ubuwqygbq4783gqb2y3f87q****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1667309705000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1667309705000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the plug-in configuration.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the gateway plug-in.
	//
	// example:
	//
	// 1
	PluginId     *int64                                                          `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	ResourceList []*GetPluginConfigResponseBodyDataGatewayConfigListResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
}

func (s GetPluginConfigResponseBodyDataGatewayConfigList) String() string {
	return dara.Prettify(s)
}

func (s GetPluginConfigResponseBodyDataGatewayConfigList) GoString() string {
	return s.String()
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) GetConfig() *string {
	return s.Config
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) GetConfigLevel() *int32 {
	return s.ConfigLevel
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) GetEnable() *bool {
	return s.Enable
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) GetId() *int64 {
	return s.Id
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) GetPluginId() *int64 {
	return s.PluginId
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) GetResourceList() []*GetPluginConfigResponseBodyDataGatewayConfigListResourceList {
	return s.ResourceList
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) SetConfig(v string) *GetPluginConfigResponseBodyDataGatewayConfigList {
	s.Config = &v
	return s
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) SetConfigLevel(v int32) *GetPluginConfigResponseBodyDataGatewayConfigList {
	s.ConfigLevel = &v
	return s
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) SetEnable(v bool) *GetPluginConfigResponseBodyDataGatewayConfigList {
	s.Enable = &v
	return s
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) SetGatewayId(v int64) *GetPluginConfigResponseBodyDataGatewayConfigList {
	s.GatewayId = &v
	return s
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) SetGatewayUniqueId(v string) *GetPluginConfigResponseBodyDataGatewayConfigList {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) SetGmtCreate(v string) *GetPluginConfigResponseBodyDataGatewayConfigList {
	s.GmtCreate = &v
	return s
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) SetGmtModified(v string) *GetPluginConfigResponseBodyDataGatewayConfigList {
	s.GmtModified = &v
	return s
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) SetId(v int64) *GetPluginConfigResponseBodyDataGatewayConfigList {
	s.Id = &v
	return s
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) SetPluginId(v int64) *GetPluginConfigResponseBodyDataGatewayConfigList {
	s.PluginId = &v
	return s
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) SetResourceList(v []*GetPluginConfigResponseBodyDataGatewayConfigListResourceList) *GetPluginConfigResponseBodyDataGatewayConfigList {
	s.ResourceList = v
	return s
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigList) Validate() error {
	return dara.Validate(s)
}

type GetPluginConfigResponseBodyDataGatewayConfigListResourceList struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test-route
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPluginConfigResponseBodyDataGatewayConfigListResourceList) String() string {
	return dara.Prettify(s)
}

func (s GetPluginConfigResponseBodyDataGatewayConfigListResourceList) GoString() string {
	return s.String()
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigListResourceList) GetId() *int64 {
	return s.Id
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigListResourceList) GetName() *string {
	return s.Name
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigListResourceList) SetId(v int64) *GetPluginConfigResponseBodyDataGatewayConfigListResourceList {
	s.Id = &v
	return s
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigListResourceList) SetName(v string) *GetPluginConfigResponseBodyDataGatewayConfigListResourceList {
	s.Name = &v
	return s
}

func (s *GetPluginConfigResponseBodyDataGatewayConfigListResourceList) Validate() error {
	return dara.Validate(s)
}
