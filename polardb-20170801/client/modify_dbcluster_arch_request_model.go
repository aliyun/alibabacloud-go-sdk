// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterArchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUseCoupon(v bool) *ModifyDBClusterArchRequest
	GetAutoUseCoupon() *bool
	SetDBClusterId(v string) *ModifyDBClusterArchRequest
	GetDBClusterId() *string
	SetHotStandbyCluster(v string) *ModifyDBClusterArchRequest
	GetHotStandbyCluster() *string
	SetPromotionCode(v string) *ModifyDBClusterArchRequest
	GetPromotionCode() *string
	SetRegionId(v string) *ModifyDBClusterArchRequest
	GetRegionId() *string
	SetStandbyAZ(v string) *ModifyDBClusterArchRequest
	GetStandbyAZ() *string
}

type ModifyDBClusterArchRequest struct {
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable a hot standby cluster. Valid values:
	//
	// - **on**: Enables a hot standby cluster.
	//
	// - **equal**: Enables a peer cluster.
	//
	// example:
	//
	// on
	HotStandbyCluster *string `json:"HotStandbyCluster,omitempty" xml:"HotStandbyCluster,omitempty"`
	// example:
	//
	// 727xxxxxx934
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The region ID.
	//
	// > For more information, see [DescribeRegions](https://help.aliyun.com/document_detail/98041.html).
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The zone for the hot standby storage cluster. Valid values:
	//
	// - **auto*	- (default): The system automatically selects a zone.
	//
	// > The default value is valid only when \\`HotStandbyCluster\\` is set to \\`on\\`. A specific zone is required when \\`HotStandbyCluster\\` is set to \\`equal\\`. For more information about zones, see [DescribeZones](https://help.aliyun.com/document_detail/98041.html).
	//
	// example:
	//
	// cn-beijing-i
	StandbyAZ *string `json:"StandbyAZ,omitempty" xml:"StandbyAZ,omitempty"`
}

func (s ModifyDBClusterArchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterArchRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterArchRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *ModifyDBClusterArchRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterArchRequest) GetHotStandbyCluster() *string {
	return s.HotStandbyCluster
}

func (s *ModifyDBClusterArchRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyDBClusterArchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBClusterArchRequest) GetStandbyAZ() *string {
	return s.StandbyAZ
}

func (s *ModifyDBClusterArchRequest) SetAutoUseCoupon(v bool) *ModifyDBClusterArchRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *ModifyDBClusterArchRequest) SetDBClusterId(v string) *ModifyDBClusterArchRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterArchRequest) SetHotStandbyCluster(v string) *ModifyDBClusterArchRequest {
	s.HotStandbyCluster = &v
	return s
}

func (s *ModifyDBClusterArchRequest) SetPromotionCode(v string) *ModifyDBClusterArchRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyDBClusterArchRequest) SetRegionId(v string) *ModifyDBClusterArchRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterArchRequest) SetStandbyAZ(v string) *ModifyDBClusterArchRequest {
	s.StandbyAZ = &v
	return s
}

func (s *ModifyDBClusterArchRequest) Validate() error {
	return dara.Validate(s)
}
