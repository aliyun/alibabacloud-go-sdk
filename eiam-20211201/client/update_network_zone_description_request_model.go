// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkZoneDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateNetworkZoneDescriptionRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateNetworkZoneDescriptionRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateNetworkZoneDescriptionRequest
	GetInstanceId() *string
	SetNetworkZoneId(v string) *UpdateNetworkZoneDescriptionRequest
	GetNetworkZoneId() *string
}

type UpdateNetworkZoneDescriptionRequest struct {
	// example:
	//
	// client-token-examplexxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 网络区域描述
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// IDaaS的网络区域主键id
	//
	// This parameter is required.
	//
	// example:
	//
	// network_11111
	NetworkZoneId *string `json:"NetworkZoneId,omitempty" xml:"NetworkZoneId,omitempty"`
}

func (s UpdateNetworkZoneDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkZoneDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateNetworkZoneDescriptionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateNetworkZoneDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateNetworkZoneDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateNetworkZoneDescriptionRequest) GetNetworkZoneId() *string {
	return s.NetworkZoneId
}

func (s *UpdateNetworkZoneDescriptionRequest) SetClientToken(v string) *UpdateNetworkZoneDescriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateNetworkZoneDescriptionRequest) SetDescription(v string) *UpdateNetworkZoneDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateNetworkZoneDescriptionRequest) SetInstanceId(v string) *UpdateNetworkZoneDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNetworkZoneDescriptionRequest) SetNetworkZoneId(v string) *UpdateNetworkZoneDescriptionRequest {
	s.NetworkZoneId = &v
	return s
}

func (s *UpdateNetworkZoneDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
