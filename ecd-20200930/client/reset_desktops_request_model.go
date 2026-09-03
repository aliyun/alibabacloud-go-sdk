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
	// The shared cloud computer ID.
	//
	// - If DesktopId is specified, DesktopGroupId is ignored.
	//
	// - If DesktopId is empty, the system retrieves the DesktopId of all cloud computers within the shared cloud computer based on DesktopGroupId.
	//
	// example:
	//
	// dg-07if7qsxoxkb6****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The shared cloud computer IDs.
	DesktopGroupIds []*string `json:"DesktopGroupIds,omitempty" xml:"DesktopGroupIds,omitempty" type:"Repeated"`
	// The cloud computer IDs. You can specify 1 to 100 IDs.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The image ID.
	//
	// example:
	//
	// m-4zfb6zj728hhr****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The timestamp of the last retry. Unit: milliseconds.
	//
	// example:
	//
	// 1699960800000
	LastRetryTime *int64 `json:"LastRetryTime,omitempty" xml:"LastRetryTime,omitempty"`
	// The billing method.
	//
	// > This parameter takes effect only for resetting shared cloud computers. If this parameter is left empty, all cloud computers of all billing methods within the shared cloud computer are reset.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID. Call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The reset scope. You can configure this parameter to specify whether to reset the image or the cloud computer.
	//
	// example:
	//
	// ALL
	ResetScope *string `json:"ResetScope,omitempty" xml:"ResetScope,omitempty"`
	// The reset type, which determines whether to reset and the scope of cloud disks to reset.
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
