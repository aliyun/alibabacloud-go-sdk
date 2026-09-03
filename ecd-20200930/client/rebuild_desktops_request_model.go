// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfterStatus(v string) *RebuildDesktopsRequest
	GetAfterStatus() *string
	SetDesktopId(v []*string) *RebuildDesktopsRequest
	GetDesktopId() []*string
	SetImageId(v string) *RebuildDesktopsRequest
	GetImageId() *string
	SetLanguage(v string) *RebuildDesktopsRequest
	GetLanguage() *string
	SetOperateType(v string) *RebuildDesktopsRequest
	GetOperateType() *string
	SetRegionId(v string) *RebuildDesktopsRequest
	GetRegionId() *string
}

type RebuildDesktopsRequest struct {
	// The target status of the cloud computer after the rebuild is complete.
	//
	// example:
	//
	// Running
	AfterStatus *string `json:"AfterStatus,omitempty" xml:"AfterStatus,omitempty"`
	// The cloud computer ID. You can specify 1 to 20 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The ID of the new image to use after the change.
	//
	// example:
	//
	// m-84mztzatmlnys****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The operating system language. Currently, only system images are supported, and Linux computers only support English.
	//
	// example:
	//
	// en-US
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The operation type for the data cloud disk.
	//
	// > Regardless of whether the cloud computer has a data cloud disk, no field value is passed in by default when you call this operation.
	//
	// - If the cloud computer has no data cloud disk:
	//
	//         No data cloud disk operation is performed regardless of the field value passed in.
	//
	// - If the cloud computer has a data cloud disk:
	//
	//     1. When the operating system of the cloud computer is the same as that of the target image:
	//
	//         - If the field value is `replace`, the data cloud disk of the cloud computer is replaced.
	//
	//         - If no field value is passed in, the original data cloud disk of the cloud computer is retained.
	//
	//     2. When the operating system of the cloud computer is different from that of the target image:
	//
	//         - If the field value is `replace`, the data cloud disk of the cloud computer is replaced.
	//
	//         - If no field value is passed in, the data cloud disk of the cloud computer is cleared.
	//
	// example:
	//
	// replace
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the list of regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RebuildDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s RebuildDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsRequest) GetAfterStatus() *string {
	return s.AfterStatus
}

func (s *RebuildDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *RebuildDesktopsRequest) GetImageId() *string {
	return s.ImageId
}

func (s *RebuildDesktopsRequest) GetLanguage() *string {
	return s.Language
}

func (s *RebuildDesktopsRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *RebuildDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RebuildDesktopsRequest) SetAfterStatus(v string) *RebuildDesktopsRequest {
	s.AfterStatus = &v
	return s
}

func (s *RebuildDesktopsRequest) SetDesktopId(v []*string) *RebuildDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *RebuildDesktopsRequest) SetImageId(v string) *RebuildDesktopsRequest {
	s.ImageId = &v
	return s
}

func (s *RebuildDesktopsRequest) SetLanguage(v string) *RebuildDesktopsRequest {
	s.Language = &v
	return s
}

func (s *RebuildDesktopsRequest) SetOperateType(v string) *RebuildDesktopsRequest {
	s.OperateType = &v
	return s
}

func (s *RebuildDesktopsRequest) SetRegionId(v string) *RebuildDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *RebuildDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
