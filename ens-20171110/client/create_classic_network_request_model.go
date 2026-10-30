// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClassicNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *CreateClassicNetworkRequest
	GetCidrBlock() *string
	SetDescription(v string) *CreateClassicNetworkRequest
	GetDescription() *string
	SetEnsRegionId(v string) *CreateClassicNetworkRequest
	GetEnsRegionId() *string
	SetNetworkName(v string) *CreateClassicNetworkRequest
	GetNetworkName() *string
}

type CreateClassicNetworkRequest struct {
	// The CIDR block of the network. You can use the following CIDR blocks or their subsets as the CIDR block of the network:
	//
	// - 10.0.0.0/8 (default).
	//
	// - 172.16.0.0/12.
	//
	// - 192.168.0.0/16.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/16
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The description of the network. The description must meet the following rules:
	//
	// 	- It must be 2 to 256 characters in length. It must start with a letter or a Chinese character, and cannot start with http:// or https://.
	//
	// example:
	//
	// This is my vswitch.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The edge zone ID of ENS.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu-xxxx-4
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The name of the network. The name must meet the following rules:
	//
	// - It must be 2 to 128 characters in length, and can contain English and Chinese characters.
	//
	// - It must start with an uppercase or lowercase letter or a Chinese character, and cannot start with http:// or https://.
	//
	// - It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// example
	NetworkName *string `json:"NetworkName,omitempty" xml:"NetworkName,omitempty"`
}

func (s CreateClassicNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClassicNetworkRequest) GoString() string {
	return s.String()
}

func (s *CreateClassicNetworkRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *CreateClassicNetworkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateClassicNetworkRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateClassicNetworkRequest) GetNetworkName() *string {
	return s.NetworkName
}

func (s *CreateClassicNetworkRequest) SetCidrBlock(v string) *CreateClassicNetworkRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateClassicNetworkRequest) SetDescription(v string) *CreateClassicNetworkRequest {
	s.Description = &v
	return s
}

func (s *CreateClassicNetworkRequest) SetEnsRegionId(v string) *CreateClassicNetworkRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateClassicNetworkRequest) SetNetworkName(v string) *CreateClassicNetworkRequest {
	s.NetworkName = &v
	return s
}

func (s *CreateClassicNetworkRequest) Validate() error {
	return dara.Validate(s)
}
