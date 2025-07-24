// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivyComputeTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteLivyComputeTokenRequest
	GetRegionId() *string
}

type DeleteLivyComputeTokenRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DeleteLivyComputeTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivyComputeTokenRequest) GoString() string {
	return s.String()
}

func (s *DeleteLivyComputeTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLivyComputeTokenRequest) SetRegionId(v string) *DeleteLivyComputeTokenRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLivyComputeTokenRequest) Validate() error {
	return dara.Validate(s)
}
