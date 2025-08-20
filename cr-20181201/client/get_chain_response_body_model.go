// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChainConfig(v *GetChainResponseBodyChainConfig) *GetChainResponseBody
	GetChainConfig() *GetChainResponseBodyChainConfig
	SetChainId(v string) *GetChainResponseBody
	GetChainId() *string
	SetCode(v string) *GetChainResponseBody
	GetCode() *string
	SetCreateTime(v int64) *GetChainResponseBody
	GetCreateTime() *int64
	SetDescription(v string) *GetChainResponseBody
	GetDescription() *string
	SetInstanceId(v string) *GetChainResponseBody
	GetInstanceId() *string
	SetIsSuccess(v bool) *GetChainResponseBody
	GetIsSuccess() *bool
	SetModifiedTime(v int64) *GetChainResponseBody
	GetModifiedTime() *int64
	SetName(v string) *GetChainResponseBody
	GetName() *string
	SetRequestId(v string) *GetChainResponseBody
	GetRequestId() *string
	SetScopeExclude(v []*string) *GetChainResponseBody
	GetScopeExclude() []*string
	SetScopeId(v string) *GetChainResponseBody
	GetScopeId() *string
	SetScopeType(v string) *GetChainResponseBody
	GetScopeType() *string
}

type GetChainResponseBody struct {
	ChainConfig *GetChainResponseBodyChainConfig `json:"ChainConfig,omitempty" xml:"ChainConfig,omitempty" type:"Struct"`
	// example:
	//
	// chi-0ops0gsmw5x2****
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1638255427000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// cri-4cdrlqmhn4gm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// 1638259914000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// C87993B5-7D61-5CAC-8D64-1AC732DD69FF
	RequestId    *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScopeExclude []*string `json:"ScopeExclude,omitempty" xml:"ScopeExclude,omitempty" type:"Repeated"`
	// example:
	//
	// crr-nyrh2oko32xb****
	ScopeId *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	// example:
	//
	// REPOSITORY
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s GetChainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChainResponseBody) GoString() string {
	return s.String()
}

func (s *GetChainResponseBody) GetChainConfig() *GetChainResponseBodyChainConfig {
	return s.ChainConfig
}

func (s *GetChainResponseBody) GetChainId() *string {
	return s.ChainId
}

func (s *GetChainResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChainResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetChainResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetChainResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetChainResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetChainResponseBody) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *GetChainResponseBody) GetName() *string {
	return s.Name
}

func (s *GetChainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChainResponseBody) GetScopeExclude() []*string {
	return s.ScopeExclude
}

func (s *GetChainResponseBody) GetScopeId() *string {
	return s.ScopeId
}

func (s *GetChainResponseBody) GetScopeType() *string {
	return s.ScopeType
}

func (s *GetChainResponseBody) SetChainConfig(v *GetChainResponseBodyChainConfig) *GetChainResponseBody {
	s.ChainConfig = v
	return s
}

func (s *GetChainResponseBody) SetChainId(v string) *GetChainResponseBody {
	s.ChainId = &v
	return s
}

func (s *GetChainResponseBody) SetCode(v string) *GetChainResponseBody {
	s.Code = &v
	return s
}

func (s *GetChainResponseBody) SetCreateTime(v int64) *GetChainResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetChainResponseBody) SetDescription(v string) *GetChainResponseBody {
	s.Description = &v
	return s
}

func (s *GetChainResponseBody) SetInstanceId(v string) *GetChainResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetChainResponseBody) SetIsSuccess(v bool) *GetChainResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetChainResponseBody) SetModifiedTime(v int64) *GetChainResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetChainResponseBody) SetName(v string) *GetChainResponseBody {
	s.Name = &v
	return s
}

func (s *GetChainResponseBody) SetRequestId(v string) *GetChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChainResponseBody) SetScopeExclude(v []*string) *GetChainResponseBody {
	s.ScopeExclude = v
	return s
}

func (s *GetChainResponseBody) SetScopeId(v string) *GetChainResponseBody {
	s.ScopeId = &v
	return s
}

func (s *GetChainResponseBody) SetScopeType(v string) *GetChainResponseBody {
	s.ScopeType = &v
	return s
}

func (s *GetChainResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetChainResponseBodyChainConfig struct {
	// example:
	//
	// cci-lz3ycgo69ukt****
	ChainConfigId *string `json:"ChainConfigId,omitempty" xml:"ChainConfigId,omitempty"`
	// example:
	//
	// true
	IsActive *bool                                     `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
	Nodes    []*GetChainResponseBodyChainConfigNodes   `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Routers  []*GetChainResponseBodyChainConfigRouters `json:"Routers,omitempty" xml:"Routers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetChainResponseBodyChainConfig) String() string {
	return dara.Prettify(s)
}

func (s GetChainResponseBodyChainConfig) GoString() string {
	return s.String()
}

func (s *GetChainResponseBodyChainConfig) GetChainConfigId() *string {
	return s.ChainConfigId
}

func (s *GetChainResponseBodyChainConfig) GetIsActive() *bool {
	return s.IsActive
}

func (s *GetChainResponseBodyChainConfig) GetNodes() []*GetChainResponseBodyChainConfigNodes {
	return s.Nodes
}

func (s *GetChainResponseBodyChainConfig) GetRouters() []*GetChainResponseBodyChainConfigRouters {
	return s.Routers
}

func (s *GetChainResponseBodyChainConfig) GetVersion() *string {
	return s.Version
}

func (s *GetChainResponseBodyChainConfig) SetChainConfigId(v string) *GetChainResponseBodyChainConfig {
	s.ChainConfigId = &v
	return s
}

func (s *GetChainResponseBodyChainConfig) SetIsActive(v bool) *GetChainResponseBodyChainConfig {
	s.IsActive = &v
	return s
}

func (s *GetChainResponseBodyChainConfig) SetNodes(v []*GetChainResponseBodyChainConfigNodes) *GetChainResponseBodyChainConfig {
	s.Nodes = v
	return s
}

func (s *GetChainResponseBodyChainConfig) SetRouters(v []*GetChainResponseBodyChainConfigRouters) *GetChainResponseBodyChainConfig {
	s.Routers = v
	return s
}

func (s *GetChainResponseBodyChainConfig) SetVersion(v string) *GetChainResponseBodyChainConfig {
	s.Version = &v
	return s
}

func (s *GetChainResponseBodyChainConfig) Validate() error {
	return dara.Validate(s)
}

type GetChainResponseBodyChainConfigNodes struct {
	// example:
	//
	// true
	Enable     *bool                                           `json:"Enable,omitempty" xml:"Enable,omitempty"`
	NodeConfig *GetChainResponseBodyChainConfigNodesNodeConfig `json:"NodeConfig,omitempty" xml:"NodeConfig,omitempty" type:"Struct"`
	// example:
	//
	// VULNERABILITY_SCANNING
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s GetChainResponseBodyChainConfigNodes) String() string {
	return dara.Prettify(s)
}

func (s GetChainResponseBodyChainConfigNodes) GoString() string {
	return s.String()
}

func (s *GetChainResponseBodyChainConfigNodes) GetEnable() *bool {
	return s.Enable
}

func (s *GetChainResponseBodyChainConfigNodes) GetNodeConfig() *GetChainResponseBodyChainConfigNodesNodeConfig {
	return s.NodeConfig
}

func (s *GetChainResponseBodyChainConfigNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *GetChainResponseBodyChainConfigNodes) SetEnable(v bool) *GetChainResponseBodyChainConfigNodes {
	s.Enable = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodes) SetNodeConfig(v *GetChainResponseBodyChainConfigNodesNodeConfig) *GetChainResponseBodyChainConfigNodes {
	s.NodeConfig = v
	return s
}

func (s *GetChainResponseBodyChainConfigNodes) SetNodeName(v string) *GetChainResponseBodyChainConfigNodes {
	s.NodeName = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodes) Validate() error {
	return dara.Validate(s)
}

type GetChainResponseBodyChainConfigNodesNodeConfig struct {
	DenyPolicy *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy `json:"DenyPolicy,omitempty" xml:"DenyPolicy,omitempty" type:"Struct"`
	// example:
	//
	// 3
	Retry *int32 `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// example:
	//
	// ACR_SCAN_SERVICE
	ScanEngine *string `json:"ScanEngine,omitempty" xml:"ScanEngine,omitempty"`
	Timeout    *int64  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s GetChainResponseBodyChainConfigNodesNodeConfig) String() string {
	return dara.Prettify(s)
}

func (s GetChainResponseBodyChainConfigNodesNodeConfig) GoString() string {
	return s.String()
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfig) GetDenyPolicy() *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy {
	return s.DenyPolicy
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfig) GetRetry() *int32 {
	return s.Retry
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfig) GetScanEngine() *string {
	return s.ScanEngine
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfig) GetTimeout() *int64 {
	return s.Timeout
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfig) SetDenyPolicy(v *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) *GetChainResponseBodyChainConfigNodesNodeConfig {
	s.DenyPolicy = v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfig) SetRetry(v int32) *GetChainResponseBodyChainConfigNodesNodeConfig {
	s.Retry = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfig) SetScanEngine(v string) *GetChainResponseBodyChainConfigNodesNodeConfig {
	s.ScanEngine = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfig) SetTimeout(v int64) *GetChainResponseBodyChainConfigNodesNodeConfig {
	s.Timeout = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfig) Validate() error {
	return dara.Validate(s)
}

type GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy struct {
	// example:
	//
	// BLOCK
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// identification,hc_image_exploit
	BaselineList *string `json:"BaselineList,omitempty" xml:"BaselineList,omitempty"`
	// example:
	//
	// 10
	IssueCount *string `json:"IssueCount,omitempty" xml:"IssueCount,omitempty"`
	// example:
	//
	// HIGH
	IssueLevel *string `json:"IssueLevel,omitempty" xml:"IssueLevel,omitempty"`
	// example:
	//
	// CVE-2020-8286,CVE-2020-8285
	IssueList *string `json:"IssueList,omitempty" xml:"IssueList,omitempty"`
	// example:
	//
	// AND
	Logic *string `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// example:
	//
	// mutate_cockhorse,abnormal_program
	MaliciousList *string `json:"MaliciousList,omitempty" xml:"MaliciousList,omitempty"`
}

func (s GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) GoString() string {
	return s.String()
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) GetAction() *string {
	return s.Action
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) GetBaselineList() *string {
	return s.BaselineList
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) GetIssueCount() *string {
	return s.IssueCount
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) GetIssueLevel() *string {
	return s.IssueLevel
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) GetIssueList() *string {
	return s.IssueList
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) GetLogic() *string {
	return s.Logic
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) GetMaliciousList() *string {
	return s.MaliciousList
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) SetAction(v string) *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy {
	s.Action = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) SetBaselineList(v string) *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy {
	s.BaselineList = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) SetIssueCount(v string) *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy {
	s.IssueCount = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) SetIssueLevel(v string) *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy {
	s.IssueLevel = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) SetIssueList(v string) *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy {
	s.IssueList = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) SetLogic(v string) *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy {
	s.Logic = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) SetMaliciousList(v string) *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy {
	s.MaliciousList = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) Validate() error {
	return dara.Validate(s)
}

type GetChainResponseBodyChainConfigRouters struct {
	From *GetChainResponseBodyChainConfigRoutersFrom `json:"From,omitempty" xml:"From,omitempty" type:"Struct"`
	To   *GetChainResponseBodyChainConfigRoutersTo   `json:"To,omitempty" xml:"To,omitempty" type:"Struct"`
}

func (s GetChainResponseBodyChainConfigRouters) String() string {
	return dara.Prettify(s)
}

func (s GetChainResponseBodyChainConfigRouters) GoString() string {
	return s.String()
}

func (s *GetChainResponseBodyChainConfigRouters) GetFrom() *GetChainResponseBodyChainConfigRoutersFrom {
	return s.From
}

func (s *GetChainResponseBodyChainConfigRouters) GetTo() *GetChainResponseBodyChainConfigRoutersTo {
	return s.To
}

func (s *GetChainResponseBodyChainConfigRouters) SetFrom(v *GetChainResponseBodyChainConfigRoutersFrom) *GetChainResponseBodyChainConfigRouters {
	s.From = v
	return s
}

func (s *GetChainResponseBodyChainConfigRouters) SetTo(v *GetChainResponseBodyChainConfigRoutersTo) *GetChainResponseBodyChainConfigRouters {
	s.To = v
	return s
}

func (s *GetChainResponseBodyChainConfigRouters) Validate() error {
	return dara.Validate(s)
}

type GetChainResponseBodyChainConfigRoutersFrom struct {
	// example:
	//
	// DOCKER_IMAGE_BUILD
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s GetChainResponseBodyChainConfigRoutersFrom) String() string {
	return dara.Prettify(s)
}

func (s GetChainResponseBodyChainConfigRoutersFrom) GoString() string {
	return s.String()
}

func (s *GetChainResponseBodyChainConfigRoutersFrom) GetNodeName() *string {
	return s.NodeName
}

func (s *GetChainResponseBodyChainConfigRoutersFrom) SetNodeName(v string) *GetChainResponseBodyChainConfigRoutersFrom {
	s.NodeName = &v
	return s
}

func (s *GetChainResponseBodyChainConfigRoutersFrom) Validate() error {
	return dara.Validate(s)
}

type GetChainResponseBodyChainConfigRoutersTo struct {
	// example:
	//
	// DOCKER_IMAGE_PUSH
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s GetChainResponseBodyChainConfigRoutersTo) String() string {
	return dara.Prettify(s)
}

func (s GetChainResponseBodyChainConfigRoutersTo) GoString() string {
	return s.String()
}

func (s *GetChainResponseBodyChainConfigRoutersTo) GetNodeName() *string {
	return s.NodeName
}

func (s *GetChainResponseBodyChainConfigRoutersTo) SetNodeName(v string) *GetChainResponseBodyChainConfigRoutersTo {
	s.NodeName = &v
	return s
}

func (s *GetChainResponseBodyChainConfigRoutersTo) Validate() error {
	return dara.Validate(s)
}
