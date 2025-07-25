// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDnsGtmAccessStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteDnsGtmAccessStrategyRequest
	GetLang() *string
	SetStrategyId(v string) *DeleteDnsGtmAccessStrategyRequest
	GetStrategyId() *string
}

type DeleteDnsGtmAccessStrategyRequest struct {
	// The language to return some response parameters. Default value: en. Valid values: en, zh, and ja.
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
	// testStrategyId1
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DeleteDnsGtmAccessStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDnsGtmAccessStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteDnsGtmAccessStrategyRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDnsGtmAccessStrategyRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *DeleteDnsGtmAccessStrategyRequest) SetLang(v string) *DeleteDnsGtmAccessStrategyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDnsGtmAccessStrategyRequest) SetStrategyId(v string) *DeleteDnsGtmAccessStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *DeleteDnsGtmAccessStrategyRequest) Validate() error {
	return dara.Validate(s)
}
