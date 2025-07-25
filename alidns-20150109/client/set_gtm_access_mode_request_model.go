// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetGtmAccessModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessMode(v string) *SetGtmAccessModeRequest
	GetAccessMode() *string
	SetLang(v string) *SetGtmAccessModeRequest
	GetLang() *string
	SetStrategyId(v string) *SetGtmAccessModeRequest
	GetStrategyId() *string
}

type SetGtmAccessModeRequest struct {
	// The desired access policy. Valid values:
	//
	// 	- **AUTO: performs automatic switchover between the primary and secondary address pool sets upon failures.**
	//
	// 	- **DEFAULT: specifies the primary address pool set.**
	//
	// 	- **FAILOVER: specifies the secondary address pool set.
	//
	// This parameter is required.
	//
	// example:
	//
	// AUTO
	AccessMode *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
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
	// hra0hx
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s SetGtmAccessModeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetGtmAccessModeRequest) GoString() string {
	return s.String()
}

func (s *SetGtmAccessModeRequest) GetAccessMode() *string {
	return s.AccessMode
}

func (s *SetGtmAccessModeRequest) GetLang() *string {
	return s.Lang
}

func (s *SetGtmAccessModeRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *SetGtmAccessModeRequest) SetAccessMode(v string) *SetGtmAccessModeRequest {
	s.AccessMode = &v
	return s
}

func (s *SetGtmAccessModeRequest) SetLang(v string) *SetGtmAccessModeRequest {
	s.Lang = &v
	return s
}

func (s *SetGtmAccessModeRequest) SetStrategyId(v string) *SetGtmAccessModeRequest {
	s.StrategyId = &v
	return s
}

func (s *SetGtmAccessModeRequest) Validate() error {
	return dara.Validate(s)
}
