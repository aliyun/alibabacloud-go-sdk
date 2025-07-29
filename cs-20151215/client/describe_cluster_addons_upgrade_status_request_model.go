// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAddonsUpgradeStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentIds(v []*string) *DescribeClusterAddonsUpgradeStatusRequest
	GetComponentIds() []*string
}

type DescribeClusterAddonsUpgradeStatusRequest struct {
	// The list of component names.
	//
	// This parameter is required.
	ComponentIds []*string `json:"componentIds,omitempty" xml:"componentIds,omitempty" type:"Repeated"`
}

func (s DescribeClusterAddonsUpgradeStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAddonsUpgradeStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsUpgradeStatusRequest) GetComponentIds() []*string {
	return s.ComponentIds
}

func (s *DescribeClusterAddonsUpgradeStatusRequest) SetComponentIds(v []*string) *DescribeClusterAddonsUpgradeStatusRequest {
	s.ComponentIds = v
	return s
}

func (s *DescribeClusterAddonsUpgradeStatusRequest) Validate() error {
	return dara.Validate(s)
}
