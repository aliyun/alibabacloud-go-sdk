// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ConfigInstanceWhitelistRequest
	GetInstanceId() *string
	SetWhitelist(v string) *ConfigInstanceWhitelistRequest
	GetWhitelist() *string
}

type ConfigInstanceWhitelistRequest struct {
	// The ID of the HSM.
	//
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// A list of IP addresses that you want to configure in the whitelist. Separate multiple IP addresses with spaces or commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 18.68.XX.XX,18.68.XX.XX
	Whitelist *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
}

func (s ConfigInstanceWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhitelistRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigInstanceWhitelistRequest) GetWhitelist() *string {
	return s.Whitelist
}

func (s *ConfigInstanceWhitelistRequest) SetInstanceId(v string) *ConfigInstanceWhitelistRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceWhitelistRequest) SetWhitelist(v string) *ConfigInstanceWhitelistRequest {
	s.Whitelist = &v
	return s
}

func (s *ConfigInstanceWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
