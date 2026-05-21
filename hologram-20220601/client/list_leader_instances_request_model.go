// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLeaderInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListLeaderInstancesRequest
	GetRegionId() *string
}

type ListLeaderInstancesRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListLeaderInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLeaderInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListLeaderInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLeaderInstancesRequest) SetRegionId(v string) *ListLeaderInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListLeaderInstancesRequest) Validate() error {
	return dara.Validate(s)
}
