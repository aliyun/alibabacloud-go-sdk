// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *DeleteClusterNodesShrinkRequest
	GetBodyShrink() *string
	SetClusterId(v string) *DeleteClusterNodesShrinkRequest
	GetClusterId() *string
}

type DeleteClusterNodesShrinkRequest struct {
	// This parameter is required.
	BodyShrink *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DeleteClusterNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodesShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *DeleteClusterNodesShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteClusterNodesShrinkRequest) SetBodyShrink(v string) *DeleteClusterNodesShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *DeleteClusterNodesShrinkRequest) SetClusterId(v string) *DeleteClusterNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
