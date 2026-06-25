// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCloudPhoneNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ChangeCloudPhoneNodeRequest
	GetAutoPay() *bool
	SetDisplayConfig(v string) *ChangeCloudPhoneNodeRequest
	GetDisplayConfig() *string
	SetDownBandwidthLimit(v int32) *ChangeCloudPhoneNodeRequest
	GetDownBandwidthLimit() *int32
	SetInstanceType(v string) *ChangeCloudPhoneNodeRequest
	GetInstanceType() *string
	SetNodeId(v string) *ChangeCloudPhoneNodeRequest
	GetNodeId() *string
	SetPhoneCount(v int32) *ChangeCloudPhoneNodeRequest
	GetPhoneCount() *int32
	SetPhoneDataVolume(v int32) *ChangeCloudPhoneNodeRequest
	GetPhoneDataVolume() *int32
	SetPromotionId(v string) *ChangeCloudPhoneNodeRequest
	GetPromotionId() *string
	SetShareDataVolume(v int32) *ChangeCloudPhoneNodeRequest
	GetShareDataVolume() *int32
	SetSwapSize(v int32) *ChangeCloudPhoneNodeRequest
	GetSwapSize() *int32
	SetUpBandwidthLimit(v int32) *ChangeCloudPhoneNodeRequest
	GetUpBandwidthLimit() *int32
}

type ChangeCloudPhoneNodeRequest struct {
	// Indicates if automatic payment is enabled. Default: false.
	//
	// example:
	//
	// false
	AutoPay       *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	DisplayConfig *string `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty"`
	// The downstream bandwidth limit, in Mbps.
	//
	// example:
	//
	// 50
	DownBandwidthLimit *int32 `json:"DownBandwidthLimit,omitempty" xml:"DownBandwidthLimit,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ac.max
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the cloud phone matrix.
	//
	// example:
	//
	// cpn-0ugbptfu473fy****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The number of cloud phone instances. Call the [](t2729804.xdita#)operation to query the minimum and maximum number of allowed instances.
	//
	// example:
	//
	// 20
	PhoneCount *int32 `json:"PhoneCount,omitempty" xml:"PhoneCount,omitempty"`
	// The capacity of the internal storage, in GiB. Valid values: 10 to 4000. If you do not specify this parameter, the current capacity is retained.
	//
	// example:
	//
	// 10
	PhoneDataVolume *int32 `json:"PhoneDataVolume,omitempty" xml:"PhoneDataVolume,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 50003308011****
	PromotionId     *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	ShareDataVolume *int32  `json:"ShareDataVolume,omitempty" xml:"ShareDataVolume,omitempty"`
	SwapSize        *int32  `json:"SwapSize,omitempty" xml:"SwapSize,omitempty"`
	// The upstream bandwidth limit, in Mbps.
	//
	// example:
	//
	// 50
	UpBandwidthLimit *int32 `json:"UpBandwidthLimit,omitempty" xml:"UpBandwidthLimit,omitempty"`
}

func (s ChangeCloudPhoneNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeCloudPhoneNodeRequest) GoString() string {
	return s.String()
}

func (s *ChangeCloudPhoneNodeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ChangeCloudPhoneNodeRequest) GetDisplayConfig() *string {
	return s.DisplayConfig
}

func (s *ChangeCloudPhoneNodeRequest) GetDownBandwidthLimit() *int32 {
	return s.DownBandwidthLimit
}

func (s *ChangeCloudPhoneNodeRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ChangeCloudPhoneNodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ChangeCloudPhoneNodeRequest) GetPhoneCount() *int32 {
	return s.PhoneCount
}

func (s *ChangeCloudPhoneNodeRequest) GetPhoneDataVolume() *int32 {
	return s.PhoneDataVolume
}

func (s *ChangeCloudPhoneNodeRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *ChangeCloudPhoneNodeRequest) GetShareDataVolume() *int32 {
	return s.ShareDataVolume
}

func (s *ChangeCloudPhoneNodeRequest) GetSwapSize() *int32 {
	return s.SwapSize
}

func (s *ChangeCloudPhoneNodeRequest) GetUpBandwidthLimit() *int32 {
	return s.UpBandwidthLimit
}

func (s *ChangeCloudPhoneNodeRequest) SetAutoPay(v bool) *ChangeCloudPhoneNodeRequest {
	s.AutoPay = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) SetDisplayConfig(v string) *ChangeCloudPhoneNodeRequest {
	s.DisplayConfig = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) SetDownBandwidthLimit(v int32) *ChangeCloudPhoneNodeRequest {
	s.DownBandwidthLimit = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) SetInstanceType(v string) *ChangeCloudPhoneNodeRequest {
	s.InstanceType = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) SetNodeId(v string) *ChangeCloudPhoneNodeRequest {
	s.NodeId = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) SetPhoneCount(v int32) *ChangeCloudPhoneNodeRequest {
	s.PhoneCount = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) SetPhoneDataVolume(v int32) *ChangeCloudPhoneNodeRequest {
	s.PhoneDataVolume = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) SetPromotionId(v string) *ChangeCloudPhoneNodeRequest {
	s.PromotionId = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) SetShareDataVolume(v int32) *ChangeCloudPhoneNodeRequest {
	s.ShareDataVolume = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) SetSwapSize(v int32) *ChangeCloudPhoneNodeRequest {
	s.SwapSize = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) SetUpBandwidthLimit(v int32) *ChangeCloudPhoneNodeRequest {
	s.UpBandwidthLimit = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) Validate() error {
	return dara.Validate(s)
}
