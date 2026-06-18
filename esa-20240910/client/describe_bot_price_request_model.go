// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBotPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBotInstanceLevel(v string) *DescribeBotPriceRequest
	GetBotInstanceLevel() *string
}

type DescribeBotPriceRequest struct {
	// The bot instance type.
	//
	// This parameter is required.
	//
	// example:
	//
	// enterprise_bot
	BotInstanceLevel *string `json:"BotInstanceLevel,omitempty" xml:"BotInstanceLevel,omitempty"`
}

func (s DescribeBotPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBotPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeBotPriceRequest) GetBotInstanceLevel() *string {
	return s.BotInstanceLevel
}

func (s *DescribeBotPriceRequest) SetBotInstanceLevel(v string) *DescribeBotPriceRequest {
	s.BotInstanceLevel = &v
	return s
}

func (s *DescribeBotPriceRequest) Validate() error {
	return dara.Validate(s)
}
