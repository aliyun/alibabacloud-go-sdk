// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarningConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWarningConfigResponseBody
	GetCode() *string
	SetData(v *ListWarningConfigResponseBodyData) *ListWarningConfigResponseBody
	GetData() *ListWarningConfigResponseBodyData
	SetMessage(v string) *ListWarningConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListWarningConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWarningConfigResponseBody
	GetSuccess() *bool
}

type ListWarningConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListWarningConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4D55C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWarningConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWarningConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWarningConfigResponseBody) GetData() *ListWarningConfigResponseBodyData {
	return s.Data
}

func (s *ListWarningConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWarningConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWarningConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWarningConfigResponseBody) SetCode(v string) *ListWarningConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ListWarningConfigResponseBody) SetData(v *ListWarningConfigResponseBodyData) *ListWarningConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListWarningConfigResponseBody) SetMessage(v string) *ListWarningConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ListWarningConfigResponseBody) SetRequestId(v string) *ListWarningConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWarningConfigResponseBody) SetSuccess(v bool) *ListWarningConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ListWarningConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListWarningConfigResponseBodyData struct {
	WarningConfigInfo []*ListWarningConfigResponseBodyDataWarningConfigInfo `json:"WarningConfigInfo,omitempty" xml:"WarningConfigInfo,omitempty" type:"Repeated"`
}

func (s ListWarningConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWarningConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyData) GetWarningConfigInfo() []*ListWarningConfigResponseBodyDataWarningConfigInfo {
	return s.WarningConfigInfo
}

func (s *ListWarningConfigResponseBodyData) SetWarningConfigInfo(v []*ListWarningConfigResponseBodyDataWarningConfigInfo) *ListWarningConfigResponseBodyData {
	s.WarningConfigInfo = v
	return s
}

func (s *ListWarningConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListWarningConfigResponseBodyDataWarningConfigInfo struct {
	Channels *ListWarningConfigResponseBodyDataWarningConfigInfoChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Struct"`
	// example:
	//
	// 32
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// example:
	//
	// 2019-10-29T15:30Z
	CreateTime *string                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RidList    *ListWarningConfigResponseBodyDataWarningConfigInfoRidList  `json:"RidList,omitempty" xml:"RidList,omitempty" type:"Struct"`
	RuleList   *ListWarningConfigResponseBodyDataWarningConfigInfoRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2019-10-29T17:24Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfo) String() string {
	return dara.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfo) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) GetChannels() *ListWarningConfigResponseBodyDataWarningConfigInfoChannels {
	return s.Channels
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) GetConfigName() *string {
	return s.ConfigName
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) GetRidList() *ListWarningConfigResponseBodyDataWarningConfigInfoRidList {
	return s.RidList
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) GetRuleList() *ListWarningConfigResponseBodyDataWarningConfigInfoRuleList {
	return s.RuleList
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) GetStatus() *int32 {
	return s.Status
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetChannels(v *ListWarningConfigResponseBodyDataWarningConfigInfoChannels) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.Channels = v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetConfigId(v int64) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.ConfigId = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetConfigName(v string) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.ConfigName = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetCreateTime(v string) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.CreateTime = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetRidList(v *ListWarningConfigResponseBodyDataWarningConfigInfoRidList) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.RidList = v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetRuleList(v *ListWarningConfigResponseBodyDataWarningConfigInfoRuleList) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.RuleList = v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetStatus(v int32) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.Status = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetUpdateTime(v string) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.UpdateTime = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) Validate() error {
	return dara.Validate(s)
}

type ListWarningConfigResponseBodyDataWarningConfigInfoChannels struct {
	Channel []*ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Repeated"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoChannels) String() string {
	return dara.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoChannels) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoChannels) GetChannel() []*ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel {
	return s.Channel
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoChannels) SetChannel(v []*ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) *ListWarningConfigResponseBodyDataWarningConfigInfoChannels {
	s.Channel = v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoChannels) Validate() error {
	return dara.Validate(s)
}

type ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel struct {
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// oapi.dingtalk.com/robot/send?access_token=c55628f700eb9ad2a3ca
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) String() string {
	return dara.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) GetType() *int32 {
	return s.Type
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) GetUrl() *string {
	return s.Url
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) SetType(v int32) *ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel {
	s.Type = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) SetUrl(v string) *ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel {
	s.Url = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) Validate() error {
	return dara.Validate(s)
}

type ListWarningConfigResponseBodyDataWarningConfigInfoRidList struct {
	RidList []*string `json:"RidList,omitempty" xml:"RidList,omitempty" type:"Repeated"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRidList) String() string {
	return dara.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRidList) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRidList) GetRidList() []*string {
	return s.RidList
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRidList) SetRidList(v []*string) *ListWarningConfigResponseBodyDataWarningConfigInfoRidList {
	s.RidList = v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRidList) Validate() error {
	return dara.Validate(s)
}

type ListWarningConfigResponseBodyDataWarningConfigInfoRuleList struct {
	WarningRule []*ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule `json:"WarningRule,omitempty" xml:"WarningRule,omitempty" type:"Repeated"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRuleList) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRuleList) GetWarningRule() []*ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule {
	return s.WarningRule
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRuleList) SetWarningRule(v []*ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) *ListWarningConfigResponseBodyDataWarningConfigInfoRuleList {
	s.WarningRule = v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRuleList) Validate() error {
	return dara.Validate(s)
}

type ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule struct {
	// example:
	//
	// 33452
	Rid      *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) String() string {
	return dara.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) GetRid() *int64 {
	return s.Rid
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) GetRuleName() *string {
	return s.RuleName
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) SetRid(v int64) *ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule {
	s.Rid = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) SetRuleName(v string) *ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule {
	s.RuleName = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) Validate() error {
	return dara.Validate(s)
}
