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
	AfterStatus *string `json:"AfterStatus,omitempty" xml:"AfterStatus,omitempty"`
	// The IDs of the cloud computers to rebuild. You can specify 1 to 20 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The ID of the new image.
	//
	// example:
	//
	// m-84mztzatmlnys****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The operating system language. This parameter applies only to system images. For Linux cloud computers, only English is supported.
	//
	// example:
	//
	// en-US
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Specifies how to handle the data disk.
	//
	// > This parameter is optional.
	//
	// - If a cloud computer does not have a data disk, this parameter is ignored.<br>
	//
	// - If a cloud computer has a data disk:
	//
	//   1. If the new image has the same operating system as the original one:
	//
	//      - If you set this parameter to `replace`, the data disk is replaced.
	//
	//      - If you do not specify this parameter, the data disk is retained.
	//
	//   2. If the new image has a different operating system:
	//
	//      - If you set this parameter to `replace`, the data disk is replaced.
	//
	//      - If you do not specify this parameter, the data disk is erased.
	//
	// example:
	//
	// replace
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to find the regions where Elastic Desktop Service is available.
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
