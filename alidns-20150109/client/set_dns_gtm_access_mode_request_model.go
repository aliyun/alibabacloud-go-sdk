// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDnsGtmAccessModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessMode(v string) *SetDnsGtmAccessModeRequest
	GetAccessMode() *string
	SetLang(v string) *SetDnsGtmAccessModeRequest
	GetLang() *string
	SetStrategyId(v string) *SetDnsGtmAccessModeRequest
	GetStrategyId() *string
}

type SetDnsGtmAccessModeRequest struct {
	// The switchover policy for the address pool collection:
	//
	// - AUTO: Automatic switchover
	//
	// - DEFAULT: The primary address pool collection
	//
	// - FAILOVER: The secondary address pool collection
	//
	// This parameter is required.
	//
	// example:
	//
	// AUTO
	AccessMode *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	// The language of certain response parameters. Default: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the access policy. Call [DescribeDnsGtmAccessStrategies](https://help.aliyun.com/document_detail/2357191.html) to obtain the policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hr**zb
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s SetDnsGtmAccessModeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDnsGtmAccessModeRequest) GoString() string {
	return s.String()
}

func (s *SetDnsGtmAccessModeRequest) GetAccessMode() *string {
	return s.AccessMode
}

func (s *SetDnsGtmAccessModeRequest) GetLang() *string {
	return s.Lang
}

func (s *SetDnsGtmAccessModeRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *SetDnsGtmAccessModeRequest) SetAccessMode(v string) *SetDnsGtmAccessModeRequest {
	s.AccessMode = &v
	return s
}

func (s *SetDnsGtmAccessModeRequest) SetLang(v string) *SetDnsGtmAccessModeRequest {
	s.Lang = &v
	return s
}

func (s *SetDnsGtmAccessModeRequest) SetStrategyId(v string) *SetDnsGtmAccessModeRequest {
	s.StrategyId = &v
	return s
}

func (s *SetDnsGtmAccessModeRequest) Validate() error {
	return dara.Validate(s)
}
