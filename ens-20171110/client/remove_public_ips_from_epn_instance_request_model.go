// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePublicIpsFromEpnInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEPNInstanceId(v string) *RemovePublicIpsFromEpnInstanceRequest
	GetEPNInstanceId() *string
	SetInstanceInfos(v string) *RemovePublicIpsFromEpnInstanceRequest
	GetInstanceInfos() *string
}

type RemovePublicIpsFromEpnInstanceRequest struct {
	// The edge-to-edge communication instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epn-xxxx
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	// The information about the public IP addresses to be removed.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"PublicIpAddress":"2.230.XX.XX"},{"PublicIpAddress":"2.230.XX.XX"}]
	InstanceInfos *string `json:"InstanceInfos,omitempty" xml:"InstanceInfos,omitempty"`
}

func (s RemovePublicIpsFromEpnInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RemovePublicIpsFromEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *RemovePublicIpsFromEpnInstanceRequest) GetEPNInstanceId() *string {
	return s.EPNInstanceId
}

func (s *RemovePublicIpsFromEpnInstanceRequest) GetInstanceInfos() *string {
	return s.InstanceInfos
}

func (s *RemovePublicIpsFromEpnInstanceRequest) SetEPNInstanceId(v string) *RemovePublicIpsFromEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *RemovePublicIpsFromEpnInstanceRequest) SetInstanceInfos(v string) *RemovePublicIpsFromEpnInstanceRequest {
	s.InstanceInfos = &v
	return s
}

func (s *RemovePublicIpsFromEpnInstanceRequest) Validate() error {
	return dara.Validate(s)
}
