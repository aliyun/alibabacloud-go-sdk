// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyNetworkAttributeRequest
	GetDescription() *string
	SetNetworkId(v string) *ModifyNetworkAttributeRequest
	GetNetworkId() *string
	SetNetworkName(v string) *ModifyNetworkAttributeRequest
	GetNetworkName() *string
}

type ModifyNetworkAttributeRequest struct {
	// The description of the network.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
	//
	// example:
	//
	// this is my first network
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the network.
	//
	// This parameter is required.
	//
	// example:
	//
	// n-****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The name of the network. The name must meet the following requirements:
	//
	// 	- The name must be 2 to 128 characters in length
	//
	// 	- It must start with a letter but cannot start with http:// or https://.
	//
	// 	- The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// abc
	NetworkName *string `json:"NetworkName,omitempty" xml:"NetworkName,omitempty"`
}

func (s ModifyNetworkAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyNetworkAttributeRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *ModifyNetworkAttributeRequest) GetNetworkName() *string {
	return s.NetworkName
}

func (s *ModifyNetworkAttributeRequest) SetDescription(v string) *ModifyNetworkAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyNetworkAttributeRequest) SetNetworkId(v string) *ModifyNetworkAttributeRequest {
	s.NetworkId = &v
	return s
}

func (s *ModifyNetworkAttributeRequest) SetNetworkName(v string) *ModifyNetworkAttributeRequest {
	s.NetworkName = &v
	return s
}

func (s *ModifyNetworkAttributeRequest) Validate() error {
	return dara.Validate(s)
}
