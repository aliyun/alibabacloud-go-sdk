// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceInspectionRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetInstanceInspectionRunRequest
	GetRegionId() *string
}

type GetInstanceInspectionRunRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetInstanceInspectionRunRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceInspectionRunRequest) SetRegionId(v string) *GetInstanceInspectionRunRequest {
	s.RegionId = &v
	return s
}

func (s *GetInstanceInspectionRunRequest) Validate() error {
	return dara.Validate(s)
}
