// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstancesToNodePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AttachInstancesToNodePoolRequest
	GetClusterId() *string
	SetInstances(v []*string) *AttachInstancesToNodePoolRequest
	GetInstances() []*string
	SetNodepoolId(v string) *AttachInstancesToNodePoolRequest
	GetNodepoolId() *string
}

type AttachInstancesToNodePoolRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	Instances []*string `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// np68mi5y1dd748ky37ojo2kqdrz
	NodepoolId *string `json:"NodepoolId,omitempty" xml:"NodepoolId,omitempty"`
}

func (s AttachInstancesToNodePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachInstancesToNodePoolRequest) GoString() string {
	return s.String()
}

func (s *AttachInstancesToNodePoolRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AttachInstancesToNodePoolRequest) GetInstances() []*string {
	return s.Instances
}

func (s *AttachInstancesToNodePoolRequest) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *AttachInstancesToNodePoolRequest) SetClusterId(v string) *AttachInstancesToNodePoolRequest {
	s.ClusterId = &v
	return s
}

func (s *AttachInstancesToNodePoolRequest) SetInstances(v []*string) *AttachInstancesToNodePoolRequest {
	s.Instances = v
	return s
}

func (s *AttachInstancesToNodePoolRequest) SetNodepoolId(v string) *AttachInstancesToNodePoolRequest {
	s.NodepoolId = &v
	return s
}

func (s *AttachInstancesToNodePoolRequest) Validate() error {
	return dara.Validate(s)
}
