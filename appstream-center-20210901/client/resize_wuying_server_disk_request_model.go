// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeWuyingServerDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataDiskList(v []*ResizeWuyingServerDiskRequestDataDiskList) *ResizeWuyingServerDiskRequest
	GetDataDiskList() []*ResizeWuyingServerDiskRequestDataDiskList
	SetProductType(v string) *ResizeWuyingServerDiskRequest
	GetProductType() *string
	SetSystemDiskNewSize(v int32) *ResizeWuyingServerDiskRequest
	GetSystemDiskNewSize() *int32
	SetWuyingServerId(v string) *ResizeWuyingServerDiskRequest
	GetWuyingServerId() *string
}

type ResizeWuyingServerDiskRequest struct {
	// The size and specifications of the data cloud disks.
	DataDiskList []*ResizeWuyingServerDiskRequestDataDiskList `json:"DataDiskList,omitempty" xml:"DataDiskList,omitempty" type:"Repeated"`
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

func (s ResizeWuyingServerDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s ResizeWuyingServerDiskRequest) GoString() string {
	return s.String()
}

func (s *ResizeWuyingServerDiskRequest) GetDataDiskList() []*ResizeWuyingServerDiskRequestDataDiskList {
	return s.DataDiskList
}

func (s *ResizeWuyingServerDiskRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ResizeWuyingServerDiskRequest) GetSystemDiskNewSize() *int32 {
	return s.SystemDiskNewSize
}

func (s *ResizeWuyingServerDiskRequest) GetWuyingServerId() *string {
	return s.WuyingServerId
}

func (s *ResizeWuyingServerDiskRequest) SetDataDiskList(v []*ResizeWuyingServerDiskRequestDataDiskList) *ResizeWuyingServerDiskRequest {
	s.DataDiskList = v
	return s
}

func (s *ResizeWuyingServerDiskRequest) SetProductType(v string) *ResizeWuyingServerDiskRequest {
	s.ProductType = &v
	return s
}

func (s *ResizeWuyingServerDiskRequest) SetSystemDiskNewSize(v int32) *ResizeWuyingServerDiskRequest {
	s.SystemDiskNewSize = &v
	return s
}

func (s *ResizeWuyingServerDiskRequest) SetWuyingServerId(v string) *ResizeWuyingServerDiskRequest {
	s.WuyingServerId = &v
	return s
}

func (s *ResizeWuyingServerDiskRequest) Validate() error {
	if s.DataDiskList != nil {
		for _, item := range s.DataDiskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ResizeWuyingServerDiskRequestDataDiskList struct {
	// The disk number.
	//
	// example:
	//
	// 1
	DiskNo *string `json:"DiskNo,omitempty" xml:"DiskNo,omitempty"`
	// The disk subtype.
	//
	// example:
	//
	// USER_DATA
	DiskSubType *string `json:"DiskSubType,omitempty" xml:"DiskSubType,omitempty"`
	// The type of disk data included in the image. By default, the system cloud disk and data cloud disks of the instance are included.
	//
	// example:
	//
	// SYSTEM
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The new size.
	//
	// example:
	//
	// 220
	NewSize *int32 `json:"NewSize,omitempty" xml:"NewSize,omitempty"`
}

func (s ResizeWuyingServerDiskRequestDataDiskList) String() string {
	return dara.Prettify(s)
}

func (s ResizeWuyingServerDiskRequestDataDiskList) GoString() string {
	return s.String()
}

func (s *ResizeWuyingServerDiskRequestDataDiskList) GetDiskNo() *string {
	return s.DiskNo
}

func (s *ResizeWuyingServerDiskRequestDataDiskList) GetDiskSubType() *string {
	return s.DiskSubType
}

func (s *ResizeWuyingServerDiskRequestDataDiskList) GetDiskType() *string {
	return s.DiskType
}

func (s *ResizeWuyingServerDiskRequestDataDiskList) GetNewSize() *int32 {
	return s.NewSize
}

func (s *ResizeWuyingServerDiskRequestDataDiskList) SetDiskNo(v string) *ResizeWuyingServerDiskRequestDataDiskList {
	s.DiskNo = &v
	return s
}

func (s *ResizeWuyingServerDiskRequestDataDiskList) SetDiskSubType(v string) *ResizeWuyingServerDiskRequestDataDiskList {
	s.DiskSubType = &v
	return s
}

func (s *ResizeWuyingServerDiskRequestDataDiskList) SetDiskType(v string) *ResizeWuyingServerDiskRequestDataDiskList {
	s.DiskType = &v
	return s
}

func (s *ResizeWuyingServerDiskRequestDataDiskList) SetNewSize(v int32) *ResizeWuyingServerDiskRequestDataDiskList {
	s.NewSize = &v
	return s
}

func (s *ResizeWuyingServerDiskRequestDataDiskList) Validate() error {
	return dara.Validate(s)
}
