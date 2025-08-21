// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterVersion(v string) *CreateClusterRequest
	GetClusterVersion() *string
	SetName(v string) *CreateClusterRequest
	GetName() *string
}

type CreateClusterRequest struct {
	// The version of the cluster.
	//
	// example:
	//
	// 1.18.8
	ClusterVersion *string `json:"ClusterVersion,omitempty" xml:"ClusterVersion,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// mycluster-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) GetClusterVersion() *string {
	return s.ClusterVersion
}

func (s *CreateClusterRequest) GetName() *string {
	return s.Name
}

func (s *CreateClusterRequest) SetClusterVersion(v string) *CreateClusterRequest {
	s.ClusterVersion = &v
	return s
}

func (s *CreateClusterRequest) SetName(v string) *CreateClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateClusterRequest) Validate() error {
	return dara.Validate(s)
}
