// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContainerClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateContainerClusterRequest
	GetClusterId() *string
	SetDescription(v string) *UpdateContainerClusterRequest
	GetDescription() *string
	SetName(v string) *UpdateContainerClusterRequest
	GetName() *string
	SetNetworkType(v string) *UpdateContainerClusterRequest
	GetNetworkType() *string
	SetRenewToken(v bool) *UpdateContainerClusterRequest
	GetRenewToken() *bool
}

type UpdateContainerClusterRequest struct {
	// Cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-000**************134
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster description.
	//
	// example:
	//
	// description ack pv backup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Cluster name.
	//
	// example:
	//
	// ack_pv_backup_location
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Network type, with possible values including:
	//
	// 	- **CLASSIC**: Classic Network.
	//
	// 	- **VPC**: Virtual Private Cloud.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// Whether to regenerate the token.
	//
	// example:
	//
	// false
	RenewToken *bool `json:"RenewToken,omitempty" xml:"RenewToken,omitempty"`
}

func (s UpdateContainerClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContainerClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateContainerClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateContainerClusterRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateContainerClusterRequest) GetName() *string {
	return s.Name
}

func (s *UpdateContainerClusterRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateContainerClusterRequest) GetRenewToken() *bool {
	return s.RenewToken
}

func (s *UpdateContainerClusterRequest) SetClusterId(v string) *UpdateContainerClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateContainerClusterRequest) SetDescription(v string) *UpdateContainerClusterRequest {
	s.Description = &v
	return s
}

func (s *UpdateContainerClusterRequest) SetName(v string) *UpdateContainerClusterRequest {
	s.Name = &v
	return s
}

func (s *UpdateContainerClusterRequest) SetNetworkType(v string) *UpdateContainerClusterRequest {
	s.NetworkType = &v
	return s
}

func (s *UpdateContainerClusterRequest) SetRenewToken(v bool) *UpdateContainerClusterRequest {
	s.RenewToken = &v
	return s
}

func (s *UpdateContainerClusterRequest) Validate() error {
	return dara.Validate(s)
}
