// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecallManagementConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateRecallManagementConfigRequest
	GetInstanceId() *string
	SetNetworkConfigs(v []*UpdateRecallManagementConfigRequestNetworkConfigs) *UpdateRecallManagementConfigRequest
	GetNetworkConfigs() []*UpdateRecallManagementConfigRequestNetworkConfigs
	SetPassword(v string) *UpdateRecallManagementConfigRequest
	GetPassword() *string
}

type UpdateRecallManagementConfigRequest struct {
	// example:
	//
	// 1
	InstanceId     *string                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NetworkConfigs []*UpdateRecallManagementConfigRequestNetworkConfigs `json:"NetworkConfigs,omitempty" xml:"NetworkConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// 12345
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s UpdateRecallManagementConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateRecallManagementConfigRequest) GetNetworkConfigs() []*UpdateRecallManagementConfigRequestNetworkConfigs {
	return s.NetworkConfigs
}

func (s *UpdateRecallManagementConfigRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateRecallManagementConfigRequest) SetInstanceId(v string) *UpdateRecallManagementConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRecallManagementConfigRequest) SetNetworkConfigs(v []*UpdateRecallManagementConfigRequestNetworkConfigs) *UpdateRecallManagementConfigRequest {
	s.NetworkConfigs = v
	return s
}

func (s *UpdateRecallManagementConfigRequest) SetPassword(v string) *UpdateRecallManagementConfigRequest {
	s.Password = &v
	return s
}

func (s *UpdateRecallManagementConfigRequest) Validate() error {
	if s.NetworkConfigs != nil {
		for _, item := range s.NetworkConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateRecallManagementConfigRequestNetworkConfigs struct {
	// example:
	//
	// vpc-xxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou
	VswitchIds map[string]*string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty"`
}

func (s UpdateRecallManagementConfigRequestNetworkConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementConfigRequestNetworkConfigs) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementConfigRequestNetworkConfigs) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateRecallManagementConfigRequestNetworkConfigs) GetVswitchIds() map[string]*string {
	return s.VswitchIds
}

func (s *UpdateRecallManagementConfigRequestNetworkConfigs) SetVpcId(v string) *UpdateRecallManagementConfigRequestNetworkConfigs {
	s.VpcId = &v
	return s
}

func (s *UpdateRecallManagementConfigRequestNetworkConfigs) SetVswitchIds(v map[string]*string) *UpdateRecallManagementConfigRequestNetworkConfigs {
	s.VswitchIds = v
	return s
}

func (s *UpdateRecallManagementConfigRequestNetworkConfigs) Validate() error {
	return dara.Validate(s)
}
