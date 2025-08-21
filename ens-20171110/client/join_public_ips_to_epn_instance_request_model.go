// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinPublicIpsToEpnInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEPNInstanceId(v string) *JoinPublicIpsToEpnInstanceRequest
	GetEPNInstanceId() *string
	SetInstanceInfos(v string) *JoinPublicIpsToEpnInstanceRequest
	GetInstanceInfos() *string
}

type JoinPublicIpsToEpnInstanceRequest struct {
	// The ID of the EPN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// epn-xxxx
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	// The information about the public IP address that you want to add to the EPN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// InstanceInfos=[{"PublicIpAddress":"2.230.XX.XX"},{"PublicIpAddress":"2.230.XX.XX"}]
	InstanceInfos *string `json:"InstanceInfos,omitempty" xml:"InstanceInfos,omitempty"`
}

func (s JoinPublicIpsToEpnInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s JoinPublicIpsToEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *JoinPublicIpsToEpnInstanceRequest) GetEPNInstanceId() *string {
	return s.EPNInstanceId
}

func (s *JoinPublicIpsToEpnInstanceRequest) GetInstanceInfos() *string {
	return s.InstanceInfos
}

func (s *JoinPublicIpsToEpnInstanceRequest) SetEPNInstanceId(v string) *JoinPublicIpsToEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *JoinPublicIpsToEpnInstanceRequest) SetInstanceInfos(v string) *JoinPublicIpsToEpnInstanceRequest {
	s.InstanceInfos = &v
	return s
}

func (s *JoinPublicIpsToEpnInstanceRequest) Validate() error {
	return dara.Validate(s)
}
