// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeDeviceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListEdgeDeviceGroupRequest
	GetName() *string
	SetNum(v int32) *ListEdgeDeviceGroupRequest
	GetNum() *int32
	SetRegionId(v string) *ListEdgeDeviceGroupRequest
	GetRegionId() *string
	SetSize(v int32) *ListEdgeDeviceGroupRequest
	GetSize() *int32
	SetStatus(v string) *ListEdgeDeviceGroupRequest
	GetStatus() *string
}

type ListEdgeDeviceGroupRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Num      *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Size   *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListEdgeDeviceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *ListEdgeDeviceGroupRequest) GetName() *string {
	return s.Name
}

func (s *ListEdgeDeviceGroupRequest) GetNum() *int32 {
	return s.Num
}

func (s *ListEdgeDeviceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEdgeDeviceGroupRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListEdgeDeviceGroupRequest) GetStatus() *string {
	return s.Status
}

func (s *ListEdgeDeviceGroupRequest) SetName(v string) *ListEdgeDeviceGroupRequest {
	s.Name = &v
	return s
}

func (s *ListEdgeDeviceGroupRequest) SetNum(v int32) *ListEdgeDeviceGroupRequest {
	s.Num = &v
	return s
}

func (s *ListEdgeDeviceGroupRequest) SetRegionId(v string) *ListEdgeDeviceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListEdgeDeviceGroupRequest) SetSize(v int32) *ListEdgeDeviceGroupRequest {
	s.Size = &v
	return s
}

func (s *ListEdgeDeviceGroupRequest) SetStatus(v string) *ListEdgeDeviceGroupRequest {
	s.Status = &v
	return s
}

func (s *ListEdgeDeviceGroupRequest) Validate() error {
	return dara.Validate(s)
}
