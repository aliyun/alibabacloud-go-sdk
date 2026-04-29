// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawChannelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawChannelsResponseBody
	GetApplicationId() *string
	SetChannels(v []*DescribePolarClawChannelsResponseBodyChannels) *DescribePolarClawChannelsResponseBody
	GetChannels() []*DescribePolarClawChannelsResponseBodyChannels
	SetCode(v int32) *DescribePolarClawChannelsResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribePolarClawChannelsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePolarClawChannelsResponseBody
	GetRequestId() *string
}

type DescribePolarClawChannelsResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string                                          `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Channels      []*DescribePolarClawChannelsResponseBodyChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2281C6C9-CBAB-1AFD-8400-670750CF6025_2212
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePolarClawChannelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarClawChannelsResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawChannelsResponseBody) GetChannels() []*DescribePolarClawChannelsResponseBodyChannels {
	return s.Channels
}

func (s *DescribePolarClawChannelsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribePolarClawChannelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePolarClawChannelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarClawChannelsResponseBody) SetApplicationId(v string) *DescribePolarClawChannelsResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawChannelsResponseBody) SetChannels(v []*DescribePolarClawChannelsResponseBodyChannels) *DescribePolarClawChannelsResponseBody {
	s.Channels = v
	return s
}

func (s *DescribePolarClawChannelsResponseBody) SetCode(v int32) *DescribePolarClawChannelsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePolarClawChannelsResponseBody) SetMessage(v string) *DescribePolarClawChannelsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePolarClawChannelsResponseBody) SetRequestId(v string) *DescribePolarClawChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarClawChannelsResponseBody) Validate() error {
	if s.Channels != nil {
		for _, item := range s.Channels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarClawChannelsResponseBodyChannels struct {
	Accounts []*DescribePolarClawChannelsResponseBodyChannelsAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// example:
	//
	// feishu
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// true
	Configured *bool `json:"Configured,omitempty" xml:"Configured,omitempty"`
	// example:
	//
	// default
	DefaultAccountId *string `json:"DefaultAccountId,omitempty" xml:"DefaultAccountId,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s DescribePolarClawChannelsResponseBodyChannels) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawChannelsResponseBodyChannels) GoString() string {
	return s.String()
}

func (s *DescribePolarClawChannelsResponseBodyChannels) GetAccounts() []*DescribePolarClawChannelsResponseBodyChannelsAccounts {
	return s.Accounts
}

func (s *DescribePolarClawChannelsResponseBodyChannels) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribePolarClawChannelsResponseBodyChannels) GetConfigured() *bool {
	return s.Configured
}

func (s *DescribePolarClawChannelsResponseBodyChannels) GetDefaultAccountId() *string {
	return s.DefaultAccountId
}

func (s *DescribePolarClawChannelsResponseBodyChannels) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribePolarClawChannelsResponseBodyChannels) SetAccounts(v []*DescribePolarClawChannelsResponseBodyChannelsAccounts) *DescribePolarClawChannelsResponseBodyChannels {
	s.Accounts = v
	return s
}

func (s *DescribePolarClawChannelsResponseBodyChannels) SetChannelId(v string) *DescribePolarClawChannelsResponseBodyChannels {
	s.ChannelId = &v
	return s
}

func (s *DescribePolarClawChannelsResponseBodyChannels) SetConfigured(v bool) *DescribePolarClawChannelsResponseBodyChannels {
	s.Configured = &v
	return s
}

func (s *DescribePolarClawChannelsResponseBodyChannels) SetDefaultAccountId(v string) *DescribePolarClawChannelsResponseBodyChannels {
	s.DefaultAccountId = &v
	return s
}

func (s *DescribePolarClawChannelsResponseBodyChannels) SetEnabled(v bool) *DescribePolarClawChannelsResponseBodyChannels {
	s.Enabled = &v
	return s
}

func (s *DescribePolarClawChannelsResponseBodyChannels) Validate() error {
	if s.Accounts != nil {
		for _, item := range s.Accounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarClawChannelsResponseBodyChannelsAccounts struct {
	// example:
	//
	// default
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// true
	Configured *bool `json:"Configured,omitempty" xml:"Configured,omitempty"`
	// example:
	//
	// true
	Connected *bool `json:"Connected,omitempty" xml:"Connected,omitempty"`
	// example:
	//
	// true
	Enabled *bool     `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Issues  []*string `json:"Issues,omitempty" xml:"Issues,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	LastInboundAt *int64 `json:"LastInboundAt,omitempty" xml:"LastInboundAt,omitempty"`
	// example:
	//
	// 0
	LastOutboundAt *int64 `json:"LastOutboundAt,omitempty" xml:"LastOutboundAt,omitempty"`
}

func (s DescribePolarClawChannelsResponseBodyChannelsAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawChannelsResponseBodyChannelsAccounts) GoString() string {
	return s.String()
}

func (s *DescribePolarClawChannelsResponseBodyChannelsAccounts) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribePolarClawChannelsResponseBodyChannelsAccounts) GetConfigured() *bool {
	return s.Configured
}

func (s *DescribePolarClawChannelsResponseBodyChannelsAccounts) GetConnected() *bool {
	return s.Connected
}

func (s *DescribePolarClawChannelsResponseBodyChannelsAccounts) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribePolarClawChannelsResponseBodyChannelsAccounts) GetIssues() []*string {
	return s.Issues
}

func (s *DescribePolarClawChannelsResponseBodyChannelsAccounts) GetLastInboundAt() *int64 {
	return s.LastInboundAt
}

func (s *DescribePolarClawChannelsResponseBodyChannelsAccounts) GetLastOutboundAt() *int64 {
	return s.LastOutboundAt
}

func (s *DescribePolarClawChannelsResponseBodyChannelsAccounts) SetAccountId(v string) *DescribePolarClawChannelsResponseBodyChannelsAccounts {
	s.AccountId = &v
	return s
}

func (s *DescribePolarClawChannelsResponseBodyChannelsAccounts) SetConfigured(v bool) *DescribePolarClawChannelsResponseBodyChannelsAccounts {
	s.Configured = &v
	return s
}

func (s *DescribePolarClawChannelsResponseBodyChannelsAccounts) SetConnected(v bool) *DescribePolarClawChannelsResponseBodyChannelsAccounts {
	s.Connected = &v
	return s
}

func (s *DescribePolarClawChannelsResponseBodyChannelsAccounts) SetEnabled(v bool) *DescribePolarClawChannelsResponseBodyChannelsAccounts {
	s.Enabled = &v
	return s
}

func (s *DescribePolarClawChannelsResponseBodyChannelsAccounts) SetIssues(v []*string) *DescribePolarClawChannelsResponseBodyChannelsAccounts {
	s.Issues = v
	return s
}

func (s *DescribePolarClawChannelsResponseBodyChannelsAccounts) SetLastInboundAt(v int64) *DescribePolarClawChannelsResponseBodyChannelsAccounts {
	s.LastInboundAt = &v
	return s
}

func (s *DescribePolarClawChannelsResponseBodyChannelsAccounts) SetLastOutboundAt(v int64) *DescribePolarClawChannelsResponseBodyChannelsAccounts {
	s.LastOutboundAt = &v
	return s
}

func (s *DescribePolarClawChannelsResponseBodyChannelsAccounts) Validate() error {
	return dara.Validate(s)
}
