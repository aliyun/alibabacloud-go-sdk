// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGtmAccessStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteGtmAccessStrategyRequest
	GetLang() *string
	SetStrategyId(v string) *DeleteGtmAccessStrategyRequest
	GetStrategyId() *string
}

type DeleteGtmAccessStrategyRequest struct {
	// The language used by the user.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the access policy that you want to delete.
	//
	// example:
	//
	// hrskc
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DeleteGtmAccessStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGtmAccessStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteGtmAccessStrategyRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteGtmAccessStrategyRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *DeleteGtmAccessStrategyRequest) SetLang(v string) *DeleteGtmAccessStrategyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteGtmAccessStrategyRequest) SetStrategyId(v string) *DeleteGtmAccessStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *DeleteGtmAccessStrategyRequest) Validate() error {
	return dara.Validate(s)
}
