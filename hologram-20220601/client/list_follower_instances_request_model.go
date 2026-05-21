// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFollowerInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListFollowerInstancesRequest
	GetRegionId() *string
}

type ListFollowerInstancesRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListFollowerInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFollowerInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListFollowerInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListFollowerInstancesRequest) SetRegionId(v string) *ListFollowerInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListFollowerInstancesRequest) Validate() error {
	return dara.Validate(s)
}
