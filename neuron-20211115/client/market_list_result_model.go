// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMarketListResult interface {
	dara.Model
	String() string
	GoString() string
	SetMarkets(v []*MarketInfo) *MarketListResult
	GetMarkets() []*MarketInfo
}

type MarketListResult struct {
	Markets []*MarketInfo `json:"markets,omitempty" xml:"markets,omitempty" type:"Repeated"`
}

func (s MarketListResult) String() string {
	return dara.Prettify(s)
}

func (s MarketListResult) GoString() string {
	return s.String()
}

func (s *MarketListResult) GetMarkets() []*MarketInfo {
	return s.Markets
}

func (s *MarketListResult) SetMarkets(v []*MarketInfo) *MarketListResult {
	s.Markets = v
	return s
}

func (s *MarketListResult) Validate() error {
	if s.Markets != nil {
		for _, item := range s.Markets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
