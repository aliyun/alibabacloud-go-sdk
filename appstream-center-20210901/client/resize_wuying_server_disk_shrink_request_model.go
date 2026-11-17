// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeWuyingServerDiskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataDiskListShrink(v string) *ResizeWuyingServerDiskShrinkRequest
	GetDataDiskListShrink() *string
	SetProductType(v string) *ResizeWuyingServerDiskShrinkRequest
	GetProductType() *string
	SetSystemDiskNewSize(v int32) *ResizeWuyingServerDiskShrinkRequest
	GetSystemDiskNewSize() *int32
	SetWuyingServerId(v string) *ResizeWuyingServerDiskShrinkRequest
	GetWuyingServerId() *string
}

type ResizeWuyingServerDiskShrinkRequest struct {
	// The size and specifications of the data cloud disks.
	DataDiskListShrink *string `json:"DataDiskList,omitempty" xml:"DataDiskList,omitempty"`
	// The product type.
	//
	// example:
	//
	// WuyingServer
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The new size of the system cloud disk, in GB.
	//
	// example:
	//
	// 220
	SystemDiskNewSize *int32 `json:"SystemDiskNewSize,omitempty" xml:"SystemDiskNewSize,omitempty"`
	// The ID of the WUYING Workspace.
	//
	// example:
	//
	// ws-0bw2f11****dial
	WuyingServerId *string `json:"WuyingServerId,omitempty" xml:"WuyingServerId,omitempty"`
}

func (s ResizeWuyingServerDiskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ResizeWuyingServerDiskShrinkRequest) GoString() string {
	return s.String()
}

func (s *ResizeWuyingServerDiskShrinkRequest) GetDataDiskListShrink() *string {
	return s.DataDiskListShrink
}

func (s *ResizeWuyingServerDiskShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ResizeWuyingServerDiskShrinkRequest) GetSystemDiskNewSize() *int32 {
	return s.SystemDiskNewSize
}

func (s *ResizeWuyingServerDiskShrinkRequest) GetWuyingServerId() *string {
	return s.WuyingServerId
}

func (s *ResizeWuyingServerDiskShrinkRequest) SetDataDiskListShrink(v string) *ResizeWuyingServerDiskShrinkRequest {
	s.DataDiskListShrink = &v
	return s
}

func (s *ResizeWuyingServerDiskShrinkRequest) SetProductType(v string) *ResizeWuyingServerDiskShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *ResizeWuyingServerDiskShrinkRequest) SetSystemDiskNewSize(v int32) *ResizeWuyingServerDiskShrinkRequest {
	s.SystemDiskNewSize = &v
	return s
}

func (s *ResizeWuyingServerDiskShrinkRequest) SetWuyingServerId(v string) *ResizeWuyingServerDiskShrinkRequest {
	s.WuyingServerId = &v
	return s
}

func (s *ResizeWuyingServerDiskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
