// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllDeviceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListAllDeviceGroupRequest
	GetRegionId() *string
}

type ListAllDeviceGroupRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAllDeviceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *ListAllDeviceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAllDeviceGroupRequest) SetRegionId(v string) *ListAllDeviceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListAllDeviceGroupRequest) Validate() error {
	return dara.Validate(s)
}
