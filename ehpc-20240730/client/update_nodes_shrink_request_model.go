// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateNodesShrinkRequest
	GetClusterId() *string
	SetInstancesShrink(v string) *UpdateNodesShrinkRequest
	GetInstancesShrink() *string
}

type UpdateNodesShrinkRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about the compute nodes that you want to update.
	InstancesShrink *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
}

func (s UpdateNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodesShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateNodesShrinkRequest) GetInstancesShrink() *string {
	return s.InstancesShrink
}

func (s *UpdateNodesShrinkRequest) SetClusterId(v string) *UpdateNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateNodesShrinkRequest) SetInstancesShrink(v string) *UpdateNodesShrinkRequest {
	s.InstancesShrink = &v
	return s
}

func (s *UpdateNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
