// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContainerClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteContainerClusterRequest
	GetClusterId() *string
	SetForce(v bool) *DeleteContainerClusterRequest
	GetForce() *bool
}

type DeleteContainerClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-0005**********hhjw
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
}

func (s DeleteContainerClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContainerClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteContainerClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteContainerClusterRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteContainerClusterRequest) SetClusterId(v string) *DeleteContainerClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteContainerClusterRequest) SetForce(v bool) *DeleteContainerClusterRequest {
	s.Force = &v
	return s
}

func (s *DeleteContainerClusterRequest) Validate() error {
	return dara.Validate(s)
}
