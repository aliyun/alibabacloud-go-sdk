// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopGroupId(v string) *ResetDesktopsRequest
	GetDesktopGroupId() *string
	SetDesktopGroupIds(v []*string) *ResetDesktopsRequest
	GetDesktopGroupIds() []*string
	SetDesktopId(v []*string) *ResetDesktopsRequest
	GetDesktopId() []*string
	SetImageId(v string) *ResetDesktopsRequest
	GetImageId() *string
	SetLastRetryTime(v int64) *ResetDesktopsRequest
	GetLastRetryTime() *int64
	SetPayType(v string) *ResetDesktopsRequest
	GetPayType() *string
	SetRegionId(v string) *ResetDesktopsRequest
	GetRegionId() *string
	SetResetScope(v string) *ResetDesktopsRequest
	GetResetScope() *string
	SetResetType(v string) *ResetDesktopsRequest
	GetResetType() *string
}

type ResetDesktopsRequest struct {
	// The ID of the shared cloud desktop.
	//
	// - If you specify `DesktopId`, the system ignores `DesktopGroupId`.
	//
	// - If `DesktopId` is empty, the system uses `DesktopGroupId` to retrieve the `DesktopId` of all cloud desktops in the shared cloud desktop group.
	//
	// example:
	//
	// dg-07if7qsxoxkb6****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// A list of shared cloud desktop group IDs.
	DesktopGroupIds []*string `json:"DesktopGroupIds,omitempty" xml:"DesktopGroupIds,omitempty" type:"Repeated"`
	// A list of cloud desktop IDs. You can specify 1 to 100 IDs.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The image ID.
	//
	// example:
	//
	// m-4zfb6zj728hhr****
	ImageId       *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	LastRetryTime *int64  `json:"LastRetryTime,omitempty" xml:"LastRetryTime,omitempty"`
	// The billing method.
	//
	// > This parameter applies only when resetting shared cloud desktops. If you leave it empty, the system resets all cloud desktops in the shared cloud desktop group, regardless of their billing method.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID. Call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to list regions that support WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scope of the reset operation. Set this parameter to reset either the image or the cloud desktop.
	//
	// example:
	//
	// ALL
	ResetScope *string `json:"ResetScope,omitempty" xml:"ResetScope,omitempty"`
	// The reset type. This determines whether to reset and which disks to reset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ResetType *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
}

func (s ResetDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetDesktopsRequest) GoString() string {
	return s.String()
}

func (s *ResetDesktopsRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *ResetDesktopsRequest) GetDesktopGroupIds() []*string {
	return s.DesktopGroupIds
}

func (s *ResetDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *ResetDesktopsRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ResetDesktopsRequest) GetLastRetryTime() *int64 {
	return s.LastRetryTime
}

func (s *ResetDesktopsRequest) GetPayType() *string {
	return s.PayType
}

func (s *ResetDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetDesktopsRequest) GetResetScope() *string {
	return s.ResetScope
}

func (s *ResetDesktopsRequest) GetResetType() *string {
	return s.ResetType
}

func (s *ResetDesktopsRequest) SetDesktopGroupId(v string) *ResetDesktopsRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *ResetDesktopsRequest) SetDesktopGroupIds(v []*string) *ResetDesktopsRequest {
	s.DesktopGroupIds = v
	return s
}

func (s *ResetDesktopsRequest) SetDesktopId(v []*string) *ResetDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *ResetDesktopsRequest) SetImageId(v string) *ResetDesktopsRequest {
	s.ImageId = &v
	return s
}

func (s *ResetDesktopsRequest) SetLastRetryTime(v int64) *ResetDesktopsRequest {
	s.LastRetryTime = &v
	return s
}

func (s *ResetDesktopsRequest) SetPayType(v string) *ResetDesktopsRequest {
	s.PayType = &v
	return s
}

func (s *ResetDesktopsRequest) SetRegionId(v string) *ResetDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *ResetDesktopsRequest) SetResetScope(v string) *ResetDesktopsRequest {
	s.ResetScope = &v
	return s
}

func (s *ResetDesktopsRequest) SetResetType(v string) *ResetDesktopsRequest {
	s.ResetType = &v
	return s
}

func (s *ResetDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
