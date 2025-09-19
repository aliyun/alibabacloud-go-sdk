// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContainerDefenseRuleDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetContainerDefenseRuleDetailResponseBody
	GetCode() *string
	SetData(v *GetContainerDefenseRuleDetailResponseBodyData) *GetContainerDefenseRuleDetailResponseBody
	GetData() *GetContainerDefenseRuleDetailResponseBodyData
	SetHttpStatusCode(v int32) *GetContainerDefenseRuleDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetContainerDefenseRuleDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetContainerDefenseRuleDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetContainerDefenseRuleDetailResponseBody
	GetSuccess() *bool
}

type GetContainerDefenseRuleDetailResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the rule.
	Data *GetContainerDefenseRuleDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 77546BF4-CCE8-5F8D-B42B-5FD3306B43B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetContainerDefenseRuleDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContainerDefenseRuleDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetContainerDefenseRuleDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetContainerDefenseRuleDetailResponseBody) GetData() *GetContainerDefenseRuleDetailResponseBodyData {
	return s.Data
}

func (s *GetContainerDefenseRuleDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetContainerDefenseRuleDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetContainerDefenseRuleDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContainerDefenseRuleDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetContainerDefenseRuleDetailResponseBody) SetCode(v string) *GetContainerDefenseRuleDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBody) SetData(v *GetContainerDefenseRuleDetailResponseBodyData) *GetContainerDefenseRuleDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBody) SetHttpStatusCode(v int32) *GetContainerDefenseRuleDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBody) SetMessage(v string) *GetContainerDefenseRuleDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBody) SetRequestId(v string) *GetContainerDefenseRuleDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBody) SetSuccess(v bool) *GetContainerDefenseRuleDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetContainerDefenseRuleDetailResponseBodyData struct {
	// The user ID.
	//
	// example:
	//
	// 1766185894104***
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// Custom defense configuration
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The alert name. Valid values:
	//
	// 	- **Non-image Program Startup**
	//
	// example:
	//
	// EventName
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The alert type. Valid values:
	//
	// 	- **Proactive Defense for Containers**
	//
	// example:
	//
	// EventType
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 1948
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The action specified in the rule. Valid values:
	//
	// 	- **1**: alert
	//
	// 	- **2**: block
	//
	// example:
	//
	// 1
	RuleAction *int32 `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test-000
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 0
	RuleSwitch *int32 `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **1**: system rule
	//
	// 	- **2**: custom rule
	//
	// example:
	//
	// 1
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The effective scope of the rule.
	Scope []*GetContainerDefenseRuleDetailResponseBodyDataScope `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
	// The whitelist.
	Whitelist *GetContainerDefenseRuleDetailResponseBodyDataWhitelist `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Struct"`
}

func (s GetContainerDefenseRuleDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetContainerDefenseRuleDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) GetEventName() *string {
	return s.EventName
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) GetEventType() *string {
	return s.EventType
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) GetRuleAction() *int32 {
	return s.RuleAction
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) GetRuleName() *string {
	return s.RuleName
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) GetRuleSwitch() *int32 {
	return s.RuleSwitch
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) GetRuleType() *string {
	return s.RuleType
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) GetScope() []*GetContainerDefenseRuleDetailResponseBodyDataScope {
	return s.Scope
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) GetWhitelist() *GetContainerDefenseRuleDetailResponseBodyDataWhitelist {
	return s.Whitelist
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) SetAliUid(v int64) *GetContainerDefenseRuleDetailResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) SetDescription(v string) *GetContainerDefenseRuleDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) SetEventName(v string) *GetContainerDefenseRuleDetailResponseBodyData {
	s.EventName = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) SetEventType(v string) *GetContainerDefenseRuleDetailResponseBodyData {
	s.EventType = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) SetId(v int64) *GetContainerDefenseRuleDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) SetRuleAction(v int32) *GetContainerDefenseRuleDetailResponseBodyData {
	s.RuleAction = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) SetRuleName(v string) *GetContainerDefenseRuleDetailResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) SetRuleSwitch(v int32) *GetContainerDefenseRuleDetailResponseBodyData {
	s.RuleSwitch = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) SetRuleType(v string) *GetContainerDefenseRuleDetailResponseBodyData {
	s.RuleType = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) SetScope(v []*GetContainerDefenseRuleDetailResponseBodyDataScope) *GetContainerDefenseRuleDetailResponseBodyData {
	s.Scope = v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) SetWhitelist(v *GetContainerDefenseRuleDetailResponseBodyDataWhitelist) *GetContainerDefenseRuleDetailResponseBodyData {
	s.Whitelist = v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetContainerDefenseRuleDetailResponseBodyDataScope struct {
	// Indicates whether all namespaces are included. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	AllNamespace *int32 `json:"AllNamespace,omitempty" xml:"AllNamespace,omitempty"`
	// The ID of the container cluster.
	//
	// example:
	//
	// c9bea04*2b25**
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// An array that consists of queried namespaces.
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
}

func (s GetContainerDefenseRuleDetailResponseBodyDataScope) String() string {
	return dara.Prettify(s)
}

func (s GetContainerDefenseRuleDetailResponseBodyDataScope) GoString() string {
	return s.String()
}

func (s *GetContainerDefenseRuleDetailResponseBodyDataScope) GetAllNamespace() *int32 {
	return s.AllNamespace
}

func (s *GetContainerDefenseRuleDetailResponseBodyDataScope) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetContainerDefenseRuleDetailResponseBodyDataScope) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *GetContainerDefenseRuleDetailResponseBodyDataScope) SetAllNamespace(v int32) *GetContainerDefenseRuleDetailResponseBodyDataScope {
	s.AllNamespace = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyDataScope) SetClusterId(v string) *GetContainerDefenseRuleDetailResponseBodyDataScope {
	s.ClusterId = &v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyDataScope) SetNamespaces(v []*string) *GetContainerDefenseRuleDetailResponseBodyDataScope {
	s.Namespaces = v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyDataScope) Validate() error {
	return dara.Validate(s)
}

type GetContainerDefenseRuleDetailResponseBodyDataWhitelist struct {
	// Deprecated
	//
	// The hash values of the files that are added to the whitelist.
	//
	// >  This parameter is not supported.
	Hash []*string `json:"Hash,omitempty" xml:"Hash,omitempty" type:"Repeated"`
	// An array consisting of images that are added to the whitelist.
	Image []*string `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
	// The paths to the files that are added to the whitelist.
	Path []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
}

func (s GetContainerDefenseRuleDetailResponseBodyDataWhitelist) String() string {
	return dara.Prettify(s)
}

func (s GetContainerDefenseRuleDetailResponseBodyDataWhitelist) GoString() string {
	return s.String()
}

func (s *GetContainerDefenseRuleDetailResponseBodyDataWhitelist) GetHash() []*string {
	return s.Hash
}

func (s *GetContainerDefenseRuleDetailResponseBodyDataWhitelist) GetImage() []*string {
	return s.Image
}

func (s *GetContainerDefenseRuleDetailResponseBodyDataWhitelist) GetPath() []*string {
	return s.Path
}

func (s *GetContainerDefenseRuleDetailResponseBodyDataWhitelist) SetHash(v []*string) *GetContainerDefenseRuleDetailResponseBodyDataWhitelist {
	s.Hash = v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyDataWhitelist) SetImage(v []*string) *GetContainerDefenseRuleDetailResponseBodyDataWhitelist {
	s.Image = v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyDataWhitelist) SetPath(v []*string) *GetContainerDefenseRuleDetailResponseBodyDataWhitelist {
	s.Path = v
	return s
}

func (s *GetContainerDefenseRuleDetailResponseBodyDataWhitelist) Validate() error {
	return dara.Validate(s)
}
