// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnBindLeaderInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UnBindLeaderInstanceRequest
	GetRegionId() *string
}

type UnBindLeaderInstanceRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UnBindLeaderInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UnBindLeaderInstanceRequest) GoString() string {
	return s.String()
}

func (s *UnBindLeaderInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnBindLeaderInstanceRequest) SetRegionId(v string) *UnBindLeaderInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UnBindLeaderInstanceRequest) Validate() error {
	return dara.Validate(s)
}
