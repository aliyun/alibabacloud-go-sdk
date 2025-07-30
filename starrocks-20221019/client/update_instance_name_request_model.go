// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *UpdateInstanceNameRequest
	GetClusterName() *string
	SetInstanceId(v string) *UpdateInstanceNameRequest
	GetInstanceId() *string
}

type UpdateInstanceNameRequest struct {
	// The new name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// new_name
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateInstanceNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *UpdateInstanceNameRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceNameRequest) SetClusterName(v string) *UpdateInstanceNameRequest {
	s.ClusterName = &v
	return s
}

func (s *UpdateInstanceNameRequest) SetInstanceId(v string) *UpdateInstanceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceNameRequest) Validate() error {
	return dara.Validate(s)
}
