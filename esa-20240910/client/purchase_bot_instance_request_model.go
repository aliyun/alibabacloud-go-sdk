// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseBotInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBotInstanceLevel(v string) *PurchaseBotInstanceRequest
	GetBotInstanceLevel() *string
	SetSiteInstanceId(v string) *PurchaseBotInstanceRequest
	GetSiteInstanceId() *string
}

type PurchaseBotInstanceRequest struct {
	// The bot instance specifications.
	//
	// This parameter is required.
	//
	// example:
	//
	// enterprise_bot
	BotInstanceLevel *string `json:"BotInstanceLevel,omitempty" xml:"BotInstanceLevel,omitempty"`
	// The site instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// esa-site-23kde*****
	SiteInstanceId *string `json:"SiteInstanceId,omitempty" xml:"SiteInstanceId,omitempty"`
}

func (s PurchaseBotInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s PurchaseBotInstanceRequest) GoString() string {
	return s.String()
}

func (s *PurchaseBotInstanceRequest) GetBotInstanceLevel() *string {
	return s.BotInstanceLevel
}

func (s *PurchaseBotInstanceRequest) GetSiteInstanceId() *string {
	return s.SiteInstanceId
}

func (s *PurchaseBotInstanceRequest) SetBotInstanceLevel(v string) *PurchaseBotInstanceRequest {
	s.BotInstanceLevel = &v
	return s
}

func (s *PurchaseBotInstanceRequest) SetSiteInstanceId(v string) *PurchaseBotInstanceRequest {
	s.SiteInstanceId = &v
	return s
}

func (s *PurchaseBotInstanceRequest) Validate() error {
	return dara.Validate(s)
}
