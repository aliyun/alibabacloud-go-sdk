// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchClusterMasterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *SwitchClusterMasterRequest
	GetClusterId() *string
	SetInstanceId(v string) *SwitchClusterMasterRequest
	GetInstanceId() *string
}

type SwitchClusterMasterRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster-w3G9vOJI2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the HSM that you want to promote to the master HSM.
	//
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SwitchClusterMasterRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchClusterMasterRequest) GoString() string {
	return s.String()
}

func (s *SwitchClusterMasterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *SwitchClusterMasterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SwitchClusterMasterRequest) SetClusterId(v string) *SwitchClusterMasterRequest {
	s.ClusterId = &v
	return s
}

func (s *SwitchClusterMasterRequest) SetInstanceId(v string) *SwitchClusterMasterRequest {
	s.InstanceId = &v
	return s
}

func (s *SwitchClusterMasterRequest) Validate() error {
	return dara.Validate(s)
}
