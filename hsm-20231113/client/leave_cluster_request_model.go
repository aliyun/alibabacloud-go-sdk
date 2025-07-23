// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLeaveClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *LeaveClusterRequest
	GetClusterId() *string
	SetInstanceId(v string) *LeaveClusterRequest
	GetInstanceId() *string
}

type LeaveClusterRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster-729dm40FG****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the HSM.
	//
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-mp90fxef****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s LeaveClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s LeaveClusterRequest) GoString() string {
	return s.String()
}

func (s *LeaveClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *LeaveClusterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *LeaveClusterRequest) SetClusterId(v string) *LeaveClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *LeaveClusterRequest) SetInstanceId(v string) *LeaveClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *LeaveClusterRequest) Validate() error {
	return dara.Validate(s)
}
