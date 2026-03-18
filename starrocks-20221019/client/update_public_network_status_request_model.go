// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublicNetworkStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentType(v string) *UpdatePublicNetworkStatusRequest
	GetComponentType() *string
	SetInstanceId(v string) *UpdatePublicNetworkStatusRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *UpdatePublicNetworkStatusRequest
	GetNodeGroupId() *string
	SetPublicNetworkEnabled(v bool) *UpdatePublicNetworkStatusRequest
	GetPublicNetworkEnabled() *bool
}

type UpdatePublicNetworkStatusRequest struct {
	// example:
	//
	// BE
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// true
	PublicNetworkEnabled *bool `json:"PublicNetworkEnabled,omitempty" xml:"PublicNetworkEnabled,omitempty"`
}

func (s UpdatePublicNetworkStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicNetworkStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdatePublicNetworkStatusRequest) GetComponentType() *string {
	return s.ComponentType
}

func (s *UpdatePublicNetworkStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdatePublicNetworkStatusRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *UpdatePublicNetworkStatusRequest) GetPublicNetworkEnabled() *bool {
	return s.PublicNetworkEnabled
}

func (s *UpdatePublicNetworkStatusRequest) SetComponentType(v string) *UpdatePublicNetworkStatusRequest {
	s.ComponentType = &v
	return s
}

func (s *UpdatePublicNetworkStatusRequest) SetInstanceId(v string) *UpdatePublicNetworkStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdatePublicNetworkStatusRequest) SetNodeGroupId(v string) *UpdatePublicNetworkStatusRequest {
	s.NodeGroupId = &v
	return s
}

func (s *UpdatePublicNetworkStatusRequest) SetPublicNetworkEnabled(v bool) *UpdatePublicNetworkStatusRequest {
	s.PublicNetworkEnabled = &v
	return s
}

func (s *UpdatePublicNetworkStatusRequest) Validate() error {
	return dara.Validate(s)
}
