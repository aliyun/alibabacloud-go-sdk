// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListEdgeDeviceRequest
	GetName() *string
	SetNum(v int32) *ListEdgeDeviceRequest
	GetNum() *int32
	SetProductKey(v string) *ListEdgeDeviceRequest
	GetProductKey() *string
	SetRegionId(v string) *ListEdgeDeviceRequest
	GetRegionId() *string
	SetSize(v int32) *ListEdgeDeviceRequest
	GetSize() *int32
}

type ListEdgeDeviceRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Num *int32 `json:"Num,omitempty" xml:"Num,omitempty"`
	// This parameter is required.
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListEdgeDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeDeviceRequest) GoString() string {
	return s.String()
}

func (s *ListEdgeDeviceRequest) GetName() *string {
	return s.Name
}

func (s *ListEdgeDeviceRequest) GetNum() *int32 {
	return s.Num
}

func (s *ListEdgeDeviceRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *ListEdgeDeviceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEdgeDeviceRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListEdgeDeviceRequest) SetName(v string) *ListEdgeDeviceRequest {
	s.Name = &v
	return s
}

func (s *ListEdgeDeviceRequest) SetNum(v int32) *ListEdgeDeviceRequest {
	s.Num = &v
	return s
}

func (s *ListEdgeDeviceRequest) SetProductKey(v string) *ListEdgeDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *ListEdgeDeviceRequest) SetRegionId(v string) *ListEdgeDeviceRequest {
	s.RegionId = &v
	return s
}

func (s *ListEdgeDeviceRequest) SetSize(v int32) *ListEdgeDeviceRequest {
	s.Size = &v
	return s
}

func (s *ListEdgeDeviceRequest) Validate() error {
	return dara.Validate(s)
}
