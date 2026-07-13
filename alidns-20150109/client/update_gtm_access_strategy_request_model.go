// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGtmAccessStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessLines(v string) *UpdateGtmAccessStrategyRequest
	GetAccessLines() *string
	SetDefaultAddrPoolId(v string) *UpdateGtmAccessStrategyRequest
	GetDefaultAddrPoolId() *string
	SetFailoverAddrPoolId(v string) *UpdateGtmAccessStrategyRequest
	GetFailoverAddrPoolId() *string
	SetLang(v string) *UpdateGtmAccessStrategyRequest
	GetLang() *string
	SetStrategyId(v string) *UpdateGtmAccessStrategyRequest
	GetStrategyId() *string
	SetStrategyName(v string) *UpdateGtmAccessStrategyRequest
	GetStrategyName() *string
}

type UpdateGtmAccessStrategyRequest struct {
	// The source of the DNS requests.
	//
	// example:
	//
	// ["default", "mobile"]
	AccessLines *string `json:"AccessLines,omitempty" xml:"AccessLines,omitempty"`
	// The ID of the primary address pool collection.
	//
	// example:
	//
	// hrsix
	DefaultAddrPoolId *string `json:"DefaultAddrPoolId,omitempty" xml:"DefaultAddrPoolId,omitempty"`
	// The ID of the failover address pool collection.
	//
	// example:
	//
	// hrsyw
	FailoverAddrPoolId *string `json:"FailoverAddrPoolId,omitempty" xml:"FailoverAddrPoolId,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the access policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// hrmxc
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The name of the access policy.
	//
	// example:
	//
	// 访问策略测试
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
}

func (s UpdateGtmAccessStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmAccessStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateGtmAccessStrategyRequest) GetAccessLines() *string {
	return s.AccessLines
}

func (s *UpdateGtmAccessStrategyRequest) GetDefaultAddrPoolId() *string {
	return s.DefaultAddrPoolId
}

func (s *UpdateGtmAccessStrategyRequest) GetFailoverAddrPoolId() *string {
	return s.FailoverAddrPoolId
}

func (s *UpdateGtmAccessStrategyRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateGtmAccessStrategyRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *UpdateGtmAccessStrategyRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *UpdateGtmAccessStrategyRequest) SetAccessLines(v string) *UpdateGtmAccessStrategyRequest {
	s.AccessLines = &v
	return s
}

func (s *UpdateGtmAccessStrategyRequest) SetDefaultAddrPoolId(v string) *UpdateGtmAccessStrategyRequest {
	s.DefaultAddrPoolId = &v
	return s
}

func (s *UpdateGtmAccessStrategyRequest) SetFailoverAddrPoolId(v string) *UpdateGtmAccessStrategyRequest {
	s.FailoverAddrPoolId = &v
	return s
}

func (s *UpdateGtmAccessStrategyRequest) SetLang(v string) *UpdateGtmAccessStrategyRequest {
	s.Lang = &v
	return s
}

func (s *UpdateGtmAccessStrategyRequest) SetStrategyId(v string) *UpdateGtmAccessStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *UpdateGtmAccessStrategyRequest) SetStrategyName(v string) *UpdateGtmAccessStrategyRequest {
	s.StrategyName = &v
	return s
}

func (s *UpdateGtmAccessStrategyRequest) Validate() error {
	return dara.Validate(s)
}
