// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecallManagementConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateRecallManagementConfigRequest
	GetInstanceId() *string
	SetNetworkConfigs(v []*CreateRecallManagementConfigRequestNetworkConfigs) *CreateRecallManagementConfigRequest
	GetNetworkConfigs() []*CreateRecallManagementConfigRequestNetworkConfigs
	SetPassword(v string) *CreateRecallManagementConfigRequest
	GetPassword() *string
	SetUserName(v string) *CreateRecallManagementConfigRequest
	GetUserName() *string
}

type CreateRecallManagementConfigRequest struct {
	// example:
	//
	// learn-pairec-xxx
	InstanceId     *string                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NetworkConfigs []*CreateRecallManagementConfigRequestNetworkConfigs `json:"NetworkConfigs,omitempty" xml:"NetworkConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// 12345
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// zhhangsan
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateRecallManagementConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRecallManagementConfigRequest) GetNetworkConfigs() []*CreateRecallManagementConfigRequestNetworkConfigs {
	return s.NetworkConfigs
}

func (s *CreateRecallManagementConfigRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateRecallManagementConfigRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateRecallManagementConfigRequest) SetInstanceId(v string) *CreateRecallManagementConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRecallManagementConfigRequest) SetNetworkConfigs(v []*CreateRecallManagementConfigRequestNetworkConfigs) *CreateRecallManagementConfigRequest {
	s.NetworkConfigs = v
	return s
}

func (s *CreateRecallManagementConfigRequest) SetPassword(v string) *CreateRecallManagementConfigRequest {
	s.Password = &v
	return s
}

func (s *CreateRecallManagementConfigRequest) SetUserName(v string) *CreateRecallManagementConfigRequest {
	s.UserName = &v
	return s
}

func (s *CreateRecallManagementConfigRequest) Validate() error {
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

type CreateRecallManagementConfigRequestNetworkConfigs struct {
	VSwitchIds map[string]*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// Vpc id
	//
	// example:
	//
	// vpc-xxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateRecallManagementConfigRequestNetworkConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementConfigRequestNetworkConfigs) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementConfigRequestNetworkConfigs) GetVSwitchIds() map[string]*string {
	return s.VSwitchIds
}

func (s *CreateRecallManagementConfigRequestNetworkConfigs) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateRecallManagementConfigRequestNetworkConfigs) SetVSwitchIds(v map[string]*string) *CreateRecallManagementConfigRequestNetworkConfigs {
	s.VSwitchIds = v
	return s
}

func (s *CreateRecallManagementConfigRequestNetworkConfigs) SetVpcId(v string) *CreateRecallManagementConfigRequestNetworkConfigs {
	s.VpcId = &v
	return s
}

func (s *CreateRecallManagementConfigRequestNetworkConfigs) Validate() error {
	return dara.Validate(s)
}
