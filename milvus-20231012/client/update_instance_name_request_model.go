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
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// milvus-test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-123xxx
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
