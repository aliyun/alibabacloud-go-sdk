// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RestartNodesRequest
	GetInstanceId() *string
	SetRestartNodeGroups(v []*RestartNodesRequestRestartNodeGroups) *RestartNodesRequest
	GetRestartNodeGroups() []*RestartNodesRequestRestartNodeGroups
}

type RestartNodesRequest struct {
	// example:
	//
	// c-b25e21e24388****
	InstanceId        *string                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RestartNodeGroups []*RestartNodesRequestRestartNodeGroups `json:"RestartNodeGroups,omitempty" xml:"RestartNodeGroups,omitempty" type:"Repeated"`
}

func (s RestartNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartNodesRequest) GoString() string {
	return s.String()
}

func (s *RestartNodesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RestartNodesRequest) GetRestartNodeGroups() []*RestartNodesRequestRestartNodeGroups {
	return s.RestartNodeGroups
}

func (s *RestartNodesRequest) SetInstanceId(v string) *RestartNodesRequest {
	s.InstanceId = &v
	return s
}

func (s *RestartNodesRequest) SetRestartNodeGroups(v []*RestartNodesRequestRestartNodeGroups) *RestartNodesRequest {
	s.RestartNodeGroups = v
	return s
}

func (s *RestartNodesRequest) Validate() error {
	if s.RestartNodeGroups != nil {
		for _, item := range s.RestartNodeGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RestartNodesRequestRestartNodeGroups struct {
	// example:
	//
	// false
	FastMode *bool `json:"FastMode,omitempty" xml:"FastMode,omitempty"`
	// example:
	//
	// ng-dcc7450e06a271b9
	NodeGroupId *string   `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeIds     []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
}

func (s RestartNodesRequestRestartNodeGroups) String() string {
	return dara.Prettify(s)
}

func (s RestartNodesRequestRestartNodeGroups) GoString() string {
	return s.String()
}

func (s *RestartNodesRequestRestartNodeGroups) GetFastMode() *bool {
	return s.FastMode
}

func (s *RestartNodesRequestRestartNodeGroups) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *RestartNodesRequestRestartNodeGroups) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *RestartNodesRequestRestartNodeGroups) SetFastMode(v bool) *RestartNodesRequestRestartNodeGroups {
	s.FastMode = &v
	return s
}

func (s *RestartNodesRequestRestartNodeGroups) SetNodeGroupId(v string) *RestartNodesRequestRestartNodeGroups {
	s.NodeGroupId = &v
	return s
}

func (s *RestartNodesRequestRestartNodeGroups) SetNodeIds(v []*string) *RestartNodesRequestRestartNodeGroups {
	s.NodeIds = v
	return s
}

func (s *RestartNodesRequestRestartNodeGroups) Validate() error {
	return dara.Validate(s)
}
