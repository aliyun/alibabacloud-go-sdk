// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWhitelistIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyWhitelistIpsRequest
	GetInstanceId() *string
	SetIpWhitelist(v string) *ModifyWhitelistIpsRequest
	GetIpWhitelist() *string
}

type ModifyWhitelistIpsRequest struct {
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 127.0.0.1,192.168.1.0/24
	IpWhitelist *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
}

func (s ModifyWhitelistIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWhitelistIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifyWhitelistIpsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyWhitelistIpsRequest) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *ModifyWhitelistIpsRequest) SetInstanceId(v string) *ModifyWhitelistIpsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyWhitelistIpsRequest) SetIpWhitelist(v string) *ModifyWhitelistIpsRequest {
	s.IpWhitelist = &v
	return s
}

func (s *ModifyWhitelistIpsRequest) Validate() error {
	return dara.Validate(s)
}
