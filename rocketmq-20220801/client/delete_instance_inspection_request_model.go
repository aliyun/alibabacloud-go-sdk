// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceInspectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteInstanceInspectionRequest
	GetRegionId() *string
}

type DeleteInstanceInspectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DeleteInstanceInspectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceInspectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceInspectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteInstanceInspectionRequest) SetRegionId(v string) *DeleteInstanceInspectionRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteInstanceInspectionRequest) Validate() error {
	return dara.Validate(s)
}
