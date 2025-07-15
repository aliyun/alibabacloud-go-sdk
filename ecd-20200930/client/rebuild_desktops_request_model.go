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
	// The cloud computer IDs. You can specify the IDs of 1 to 20 cloud computers.
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
	// The OS language. Only system images are supported, and Linux cloud computers support only English.
	//
	// Valid values:
	//
	// 	- en-US: English
	//
	// 	- zh-HK: Traditional Chinese (Hong Kong, China)
	//
	// 	- zh-CN: Simplified Chinese
	//
	// 	- ja-JP: Japanese
	//
	// example:
	//
	// en-US
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The operation type on the data disk.
	//
	// >  This parameter is empty by default regardless of whether data disks are attached to the cloud computer.
	//
	// 	- No data disks are attached to the cloud computer:\\
	//
	//     No operation is performed on the data disks of the cloud computer regardless of the value of this parameter.
	//
	// 	- Data disks are attached to the cloud computer:
	//
	//     1.  The OS of the cloud computer is the same as the OS of the destination image:
	//
	//         	- If you set the OperateType parameter to `replace`, the data in the data disks of the cloud computer is replaced.
	//
	//         	- If you leave the OperateType parameter empty, the data in the data disks of the cloud computer is retained.
	//
	//     2.  The OS of the cloud computer is different from the OS of the destination image:
	//
	//         	- If you set the OperateType parameter to `replace`, the data in the data disks of the cloud computer is replaced.
	//
	//         	- If you leave the OperateType parameter empty, the data in the data disks of the cloud computer is cleared.
	//
	// example:
	//
	// replace
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the regions supported by Elastic Desktop Service (EDS).
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
