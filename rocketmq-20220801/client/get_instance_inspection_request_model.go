// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceInspectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetInstanceInspectionRequest
	GetRegionId() *string
}

type GetInstanceInspectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetInstanceInspectionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceInspectionRequest) SetRegionId(v string) *GetInstanceInspectionRequest {
	s.RegionId = &v
	return s
}

func (s *GetInstanceInspectionRequest) Validate() error {
	return dara.Validate(s)
}
