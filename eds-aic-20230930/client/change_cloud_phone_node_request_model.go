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
	SetUpBandwidthLimit(v int32) *ChangeCloudPhoneNodeRequest
	GetUpBandwidthLimit() *int32
}

type ChangeCloudPhoneNodeRequest struct {
	AutoPay            *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	DisplayConfig      *string `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty"`
	DownBandwidthLimit *int32  `json:"DownBandwidthLimit,omitempty" xml:"DownBandwidthLimit,omitempty"`
	// example:
	//
	// ac.max
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// cpn-0ugbptfu473fy****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 20
	PhoneCount       *int32  `json:"PhoneCount,omitempty" xml:"PhoneCount,omitempty"`
	PhoneDataVolume  *int32  `json:"PhoneDataVolume,omitempty" xml:"PhoneDataVolume,omitempty"`
	PromotionId      *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	UpBandwidthLimit *int32  `json:"UpBandwidthLimit,omitempty" xml:"UpBandwidthLimit,omitempty"`
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

func (s *ChangeCloudPhoneNodeRequest) SetUpBandwidthLimit(v int32) *ChangeCloudPhoneNodeRequest {
	s.UpBandwidthLimit = &v
	return s
}

func (s *ChangeCloudPhoneNodeRequest) Validate() error {
	return dara.Validate(s)
}
