// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateClusterRequest
	GetClusterId() *string
	SetClusterName(v string) *UpdateClusterRequest
	GetClusterName() *string
	SetIpWhitelist(v string) *UpdateClusterRequest
	GetIpWhitelist() *string
}

type UpdateClusterRequest struct {
	// This parameter is required.
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// 192.168.1.0/24
	IpWhitelist *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
}

func (s UpdateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *UpdateClusterRequest) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *UpdateClusterRequest) SetClusterId(v string) *UpdateClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterRequest) SetClusterName(v string) *UpdateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *UpdateClusterRequest) SetIpWhitelist(v string) *UpdateClusterRequest {
	s.IpWhitelist = &v
	return s
}

func (s *UpdateClusterRequest) Validate() error {
	return dara.Validate(s)
}
