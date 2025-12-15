// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateNodesRequest
	GetClusterId() *string
	SetInstances(v []*UpdateNodesRequestInstances) *UpdateNodesRequest
	GetInstances() []*UpdateNodesRequestInstances
}

type UpdateNodesRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about the compute nodes that you want to update.
	Instances []*UpdateNodesRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
}

func (s UpdateNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodesRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateNodesRequest) GetInstances() []*UpdateNodesRequestInstances {
	return s.Instances
}

func (s *UpdateNodesRequest) SetClusterId(v string) *UpdateNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateNodesRequest) SetInstances(v []*UpdateNodesRequestInstances) *UpdateNodesRequest {
	s.Instances = v
	return s
}

func (s *UpdateNodesRequest) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateNodesRequestInstances struct {
	// The instance ID of the compute node.
	//
	// example:
	//
	// i-bp1bzqq1ddeemuddn****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to enable deletion protection for the compute node.
	//
	// example:
	//
	// true
	KeepAlive *bool `json:"KeepAlive,omitempty" xml:"KeepAlive,omitempty"`
}

func (s UpdateNodesRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodesRequestInstances) GoString() string {
	return s.String()
}

func (s *UpdateNodesRequestInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateNodesRequestInstances) GetKeepAlive() *bool {
	return s.KeepAlive
}

func (s *UpdateNodesRequestInstances) SetInstanceId(v string) *UpdateNodesRequestInstances {
	s.InstanceId = &v
	return s
}

func (s *UpdateNodesRequestInstances) SetKeepAlive(v bool) *UpdateNodesRequestInstances {
	s.KeepAlive = &v
	return s
}

func (s *UpdateNodesRequestInstances) Validate() error {
	return dara.Validate(s)
}
