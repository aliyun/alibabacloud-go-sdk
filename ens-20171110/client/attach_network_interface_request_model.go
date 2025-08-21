// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachNetworkInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AttachNetworkInterfaceRequest
	GetInstanceId() *string
	SetNetworkInterfaceId(v string) *AttachNetworkInterfaceRequest
	GetNetworkInterfaceId() *string
}

type AttachNetworkInterfaceRequest struct {
	// The ID of the instance
	//
	// This parameter is required.
	//
	// example:
	//
	// i-5p67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the ENI.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-58z57orgmt6d1****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
}

func (s AttachNetworkInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *AttachNetworkInterfaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachNetworkInterfaceRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *AttachNetworkInterfaceRequest) SetInstanceId(v string) *AttachNetworkInterfaceRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachNetworkInterfaceRequest) SetNetworkInterfaceId(v string) *AttachNetworkInterfaceRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *AttachNetworkInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
