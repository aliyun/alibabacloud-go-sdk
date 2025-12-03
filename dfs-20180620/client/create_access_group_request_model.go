// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupName(v string) *CreateAccessGroupRequest
	GetAccessGroupName() *string
	SetDescription(v string) *CreateAccessGroupRequest
	GetDescription() *string
	SetInputRegionId(v string) *CreateAccessGroupRequest
	GetInputRegionId() *string
	SetNetworkType(v string) *CreateAccessGroupRequest
	GetNetworkType() *string
}

type CreateAccessGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// my-online-cluster-policy
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s CreateAccessGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessGroupRequest) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *CreateAccessGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAccessGroupRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *CreateAccessGroupRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateAccessGroupRequest) SetAccessGroupName(v string) *CreateAccessGroupRequest {
	s.AccessGroupName = &v
	return s
}

func (s *CreateAccessGroupRequest) SetDescription(v string) *CreateAccessGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateAccessGroupRequest) SetInputRegionId(v string) *CreateAccessGroupRequest {
	s.InputRegionId = &v
	return s
}

func (s *CreateAccessGroupRequest) SetNetworkType(v string) *CreateAccessGroupRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateAccessGroupRequest) Validate() error {
	return dara.Validate(s)
}
