// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteClusterRequest
	GetClusterId() *string
	SetRetainResources(v bool) *DeleteClusterRequest
	GetRetainResources() *bool
}

type DeleteClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RetainResources *bool   `json:"RetainResources,omitempty" xml:"RetainResources,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteClusterRequest) GetRetainResources() *bool {
	return s.RetainResources
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterRequest) SetRetainResources(v bool) *DeleteClusterRequest {
	s.RetainResources = &v
	return s
}

func (s *DeleteClusterRequest) Validate() error {
	return dara.Validate(s)
}
