// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisablePolarClawChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DisablePolarClawChannelRequest
	GetApplicationId() *string
	SetChannelId(v string) *DisablePolarClawChannelRequest
	GetChannelId() *string
	SetRestart(v bool) *DisablePolarClawChannelRequest
	GetRestart() *bool
}

type DisablePolarClawChannelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// feishu
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s DisablePolarClawChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DisablePolarClawChannelRequest) GoString() string {
	return s.String()
}

func (s *DisablePolarClawChannelRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DisablePolarClawChannelRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DisablePolarClawChannelRequest) GetRestart() *bool {
	return s.Restart
}

func (s *DisablePolarClawChannelRequest) SetApplicationId(v string) *DisablePolarClawChannelRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisablePolarClawChannelRequest) SetChannelId(v string) *DisablePolarClawChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *DisablePolarClawChannelRequest) SetRestart(v bool) *DisablePolarClawChannelRequest {
	s.Restart = &v
	return s
}

func (s *DisablePolarClawChannelRequest) Validate() error {
	return dara.Validate(s)
}
