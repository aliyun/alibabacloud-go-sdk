// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMockRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddMockRuleRequest
	GetAcceptLanguage() *string
	SetConsumerAppIds(v string) *AddMockRuleRequest
	GetConsumerAppIds() *string
	SetDubboMockItems(v string) *AddMockRuleRequest
	GetDubboMockItems() *string
	SetEnable(v bool) *AddMockRuleRequest
	GetEnable() *bool
	SetExtraJson(v string) *AddMockRuleRequest
	GetExtraJson() *string
	SetMockType(v int64) *AddMockRuleRequest
	GetMockType() *int64
	SetName(v string) *AddMockRuleRequest
	GetName() *string
	SetProviderAppId(v string) *AddMockRuleRequest
	GetProviderAppId() *string
	SetProviderAppName(v string) *AddMockRuleRequest
	GetProviderAppName() *string
	SetRegion(v string) *AddMockRuleRequest
	GetRegion() *string
	SetScMockItems(v string) *AddMockRuleRequest
	GetScMockItems() *string
	SetSource(v string) *AddMockRuleRequest
	GetSource() *string
}

type AddMockRuleRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the custom application.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"appName\\":\\"provide\\",\\"appId\\":\\"bst8l6o735@f6d8aaf6e56e67d\\"}]
	ConsumerAppIds *string `json:"ConsumerAppIds,omitempty" xml:"ConsumerAppIds,omitempty"`
	// The items in the recycle bin.
	//
	// This parameter is required.
	//
	// example:
	//
	// []
	DubboMockItems *string `json:"DubboMockItems,omitempty" xml:"DubboMockItems,omitempty"`
	// Specifies whether to enable the alert rule. Valid values:
	//
	// 	- `true`: enables the alert rule.
	//
	// 	- `false`: disables the alert rule.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The description.
	//
	// This parameter is required.
	//
	// example:
	//
	// {}
	ExtraJson *string `json:"ExtraJson,omitempty" xml:"ExtraJson,omitempty"`
	// The response time (RT) threshold of slow calls. Valid values:
	//
	// 	- \\- 15: 15 ms
	//
	// 	- \\- 30: 30 ms
	//
	// 	- \\- 60: 60 ms
	//
	// 	- \\- 120: 120 ms
	//
	// example:
	//
	// 1
	MockType *int64 `json:"MockType,omitempty" xml:"MockType,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-auto-test-sc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the service provider application.
	//
	// example:
	//
	// dcqtkuhnc4@66e5235415****
	ProviderAppId *string `json:"ProviderAppId,omitempty" xml:"ProviderAppId,omitempty"`
	// The name of the service provider application.
	//
	// example:
	//
	// demo-cartservice
	ProviderAppName *string `json:"ProviderAppName,omitempty" xml:"ProviderAppName,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The input parameters. The JSON format is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// []
	ScMockItems *string `json:"ScMockItems,omitempty" xml:"ScMockItems,omitempty"`
	// The rule source.
	//
	// This parameter is required.
	//
	// example:
	//
	// edasmsc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s AddMockRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMockRuleRequest) GoString() string {
	return s.String()
}

func (s *AddMockRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddMockRuleRequest) GetConsumerAppIds() *string {
	return s.ConsumerAppIds
}

func (s *AddMockRuleRequest) GetDubboMockItems() *string {
	return s.DubboMockItems
}

func (s *AddMockRuleRequest) GetEnable() *bool {
	return s.Enable
}

func (s *AddMockRuleRequest) GetExtraJson() *string {
	return s.ExtraJson
}

func (s *AddMockRuleRequest) GetMockType() *int64 {
	return s.MockType
}

func (s *AddMockRuleRequest) GetName() *string {
	return s.Name
}

func (s *AddMockRuleRequest) GetProviderAppId() *string {
	return s.ProviderAppId
}

func (s *AddMockRuleRequest) GetProviderAppName() *string {
	return s.ProviderAppName
}

func (s *AddMockRuleRequest) GetRegion() *string {
	return s.Region
}

func (s *AddMockRuleRequest) GetScMockItems() *string {
	return s.ScMockItems
}

func (s *AddMockRuleRequest) GetSource() *string {
	return s.Source
}

func (s *AddMockRuleRequest) SetAcceptLanguage(v string) *AddMockRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddMockRuleRequest) SetConsumerAppIds(v string) *AddMockRuleRequest {
	s.ConsumerAppIds = &v
	return s
}

func (s *AddMockRuleRequest) SetDubboMockItems(v string) *AddMockRuleRequest {
	s.DubboMockItems = &v
	return s
}

func (s *AddMockRuleRequest) SetEnable(v bool) *AddMockRuleRequest {
	s.Enable = &v
	return s
}

func (s *AddMockRuleRequest) SetExtraJson(v string) *AddMockRuleRequest {
	s.ExtraJson = &v
	return s
}

func (s *AddMockRuleRequest) SetMockType(v int64) *AddMockRuleRequest {
	s.MockType = &v
	return s
}

func (s *AddMockRuleRequest) SetName(v string) *AddMockRuleRequest {
	s.Name = &v
	return s
}

func (s *AddMockRuleRequest) SetProviderAppId(v string) *AddMockRuleRequest {
	s.ProviderAppId = &v
	return s
}

func (s *AddMockRuleRequest) SetProviderAppName(v string) *AddMockRuleRequest {
	s.ProviderAppName = &v
	return s
}

func (s *AddMockRuleRequest) SetRegion(v string) *AddMockRuleRequest {
	s.Region = &v
	return s
}

func (s *AddMockRuleRequest) SetScMockItems(v string) *AddMockRuleRequest {
	s.ScMockItems = &v
	return s
}

func (s *AddMockRuleRequest) SetSource(v string) *AddMockRuleRequest {
	s.Source = &v
	return s
}

func (s *AddMockRuleRequest) Validate() error {
	return dara.Validate(s)
}
