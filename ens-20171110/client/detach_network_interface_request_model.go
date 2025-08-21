// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachNetworkInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInterfaceId(v string) *DetachNetworkInterfaceRequest
	GetNetworkInterfaceId() *string
}

type DetachNetworkInterfaceRequest struct {
	// The ID of the ENI.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-58z57orgmt6d1****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
}

func (s DetachNetworkInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *DetachNetworkInterfaceRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DetachNetworkInterfaceRequest) SetNetworkInterfaceId(v string) *DetachNetworkInterfaceRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DetachNetworkInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
