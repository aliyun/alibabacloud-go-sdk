// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDGsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *DescribeSDGsShrinkRequest
	GetInstanceIdsShrink() *string
	SetSDGIdsShrink(v string) *DescribeSDGsShrinkRequest
	GetSDGIdsShrink() *string
}

type DescribeSDGsShrinkRequest struct {
	// The AIC instance ID to be queried.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The IDs of SDGs that you want to query. By default, all SDGs are queried.
	SDGIdsShrink *string `json:"SDGIds,omitempty" xml:"SDGIds,omitempty"`
}

func (s DescribeSDGsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSDGsShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *DescribeSDGsShrinkRequest) GetSDGIdsShrink() *string {
	return s.SDGIdsShrink
}

func (s *DescribeSDGsShrinkRequest) SetInstanceIdsShrink(v string) *DescribeSDGsShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DescribeSDGsShrinkRequest) SetSDGIdsShrink(v string) *DescribeSDGsShrinkRequest {
	s.SDGIdsShrink = &v
	return s
}

func (s *DescribeSDGsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
