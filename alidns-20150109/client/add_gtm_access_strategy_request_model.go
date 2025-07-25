// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGtmAccessStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessLines(v string) *AddGtmAccessStrategyRequest
	GetAccessLines() *string
	SetDefaultAddrPoolId(v string) *AddGtmAccessStrategyRequest
	GetDefaultAddrPoolId() *string
	SetFailoverAddrPoolId(v string) *AddGtmAccessStrategyRequest
	GetFailoverAddrPoolId() *string
	SetInstanceId(v string) *AddGtmAccessStrategyRequest
	GetInstanceId() *string
	SetLang(v string) *AddGtmAccessStrategyRequest
	GetLang() *string
	SetStrategyName(v string) *AddGtmAccessStrategyRequest
	GetStrategyName() *string
}

type AddGtmAccessStrategyRequest struct {
	// The line codes of access regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["default", "drpeng"]
	AccessLines *string `json:"AccessLines,omitempty" xml:"AccessLines,omitempty"`
	// The ID of the default address pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// hrsix
	DefaultAddrPoolId *string `json:"DefaultAddrPoolId,omitempty" xml:"DefaultAddrPoolId,omitempty"`
	// The ID of the failover address pool.
	//
	// If the failover address pool is not set, pass the **Empty*	- value.
	//
	// This parameter is required.
	//
	// example:
	//
	// hrsyw
	FailoverAddrPoolId *string `json:"FailoverAddrPoolId,omitempty" xml:"FailoverAddrPoolId,omitempty"`
	// The ID of the GTM instance for which you want to create an access policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language used by the user.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the access policy.
	//
	// This parameter is required.
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
}

func (s AddGtmAccessStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGtmAccessStrategyRequest) GoString() string {
	return s.String()
}

func (s *AddGtmAccessStrategyRequest) GetAccessLines() *string {
	return s.AccessLines
}

func (s *AddGtmAccessStrategyRequest) GetDefaultAddrPoolId() *string {
	return s.DefaultAddrPoolId
}

func (s *AddGtmAccessStrategyRequest) GetFailoverAddrPoolId() *string {
	return s.FailoverAddrPoolId
}

func (s *AddGtmAccessStrategyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddGtmAccessStrategyRequest) GetLang() *string {
	return s.Lang
}

func (s *AddGtmAccessStrategyRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *AddGtmAccessStrategyRequest) SetAccessLines(v string) *AddGtmAccessStrategyRequest {
	s.AccessLines = &v
	return s
}

func (s *AddGtmAccessStrategyRequest) SetDefaultAddrPoolId(v string) *AddGtmAccessStrategyRequest {
	s.DefaultAddrPoolId = &v
	return s
}

func (s *AddGtmAccessStrategyRequest) SetFailoverAddrPoolId(v string) *AddGtmAccessStrategyRequest {
	s.FailoverAddrPoolId = &v
	return s
}

func (s *AddGtmAccessStrategyRequest) SetInstanceId(v string) *AddGtmAccessStrategyRequest {
	s.InstanceId = &v
	return s
}

func (s *AddGtmAccessStrategyRequest) SetLang(v string) *AddGtmAccessStrategyRequest {
	s.Lang = &v
	return s
}

func (s *AddGtmAccessStrategyRequest) SetStrategyName(v string) *AddGtmAccessStrategyRequest {
	s.StrategyName = &v
	return s
}

func (s *AddGtmAccessStrategyRequest) Validate() error {
	return dara.Validate(s)
}
