// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeTotalDeviceCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetEdgeTotalDeviceCountRequest
	GetRegionId() *string
}

type GetEdgeTotalDeviceCountRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetEdgeTotalDeviceCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeTotalDeviceCountRequest) GoString() string {
	return s.String()
}

func (s *GetEdgeTotalDeviceCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEdgeTotalDeviceCountRequest) SetRegionId(v string) *GetEdgeTotalDeviceCountRequest {
	s.RegionId = &v
	return s
}

func (s *GetEdgeTotalDeviceCountRequest) Validate() error {
	return dara.Validate(s)
}
