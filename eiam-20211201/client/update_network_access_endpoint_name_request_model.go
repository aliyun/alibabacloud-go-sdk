// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkAccessEndpointNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateNetworkAccessEndpointNameRequest
	GetInstanceId() *string
	SetNetworkAccessEndpointId(v string) *UpdateNetworkAccessEndpointNameRequest
	GetNetworkAccessEndpointId() *string
	SetNetworkAccessEndpointName(v string) *UpdateNetworkAccessEndpointNameRequest
	GetNetworkAccessEndpointName() *string
}

type UpdateNetworkAccessEndpointNameRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 专属网络端点ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// nae_examplexxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// 专属网络端点名称。
	//
	// This parameter is required.
	//
	// example:
	//
	// xx业务VPC访问端点
	NetworkAccessEndpointName *string `json:"NetworkAccessEndpointName,omitempty" xml:"NetworkAccessEndpointName,omitempty"`
}

func (s UpdateNetworkAccessEndpointNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkAccessEndpointNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateNetworkAccessEndpointNameRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateNetworkAccessEndpointNameRequest) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *UpdateNetworkAccessEndpointNameRequest) GetNetworkAccessEndpointName() *string {
	return s.NetworkAccessEndpointName
}

func (s *UpdateNetworkAccessEndpointNameRequest) SetInstanceId(v string) *UpdateNetworkAccessEndpointNameRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNetworkAccessEndpointNameRequest) SetNetworkAccessEndpointId(v string) *UpdateNetworkAccessEndpointNameRequest {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *UpdateNetworkAccessEndpointNameRequest) SetNetworkAccessEndpointName(v string) *UpdateNetworkAccessEndpointNameRequest {
	s.NetworkAccessEndpointName = &v
	return s
}

func (s *UpdateNetworkAccessEndpointNameRequest) Validate() error {
	return dara.Validate(s)
}
