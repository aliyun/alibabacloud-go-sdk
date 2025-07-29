// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAddonsUpgradeStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentIdsShrink(v string) *DescribeClusterAddonsUpgradeStatusShrinkRequest
	GetComponentIdsShrink() *string
}

type DescribeClusterAddonsUpgradeStatusShrinkRequest struct {
	// The list of component names.
	//
	// This parameter is required.
	ComponentIdsShrink *string `json:"componentIds,omitempty" xml:"componentIds,omitempty"`
}

func (s DescribeClusterAddonsUpgradeStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAddonsUpgradeStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsUpgradeStatusShrinkRequest) GetComponentIdsShrink() *string {
	return s.ComponentIdsShrink
}

func (s *DescribeClusterAddonsUpgradeStatusShrinkRequest) SetComponentIdsShrink(v string) *DescribeClusterAddonsUpgradeStatusShrinkRequest {
	s.ComponentIdsShrink = &v
	return s
}

func (s *DescribeClusterAddonsUpgradeStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
