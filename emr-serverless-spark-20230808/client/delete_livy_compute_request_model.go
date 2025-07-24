// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivyComputeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteLivyComputeRequest
	GetRegionId() *string
}

type DeleteLivyComputeRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DeleteLivyComputeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivyComputeRequest) GoString() string {
	return s.String()
}

func (s *DeleteLivyComputeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLivyComputeRequest) SetRegionId(v string) *DeleteLivyComputeRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLivyComputeRequest) Validate() error {
	return dara.Validate(s)
}
