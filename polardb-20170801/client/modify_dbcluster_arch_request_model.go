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
	// The ID of the cluster.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable the hot standby storage cluster feature. Valid values:
	//
	// 	- **on**: enables hot standby storage cluster.
	//
	// 	- **equal**: Enable a peer-to-peer cluster.
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
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query information about regions.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The zone of the hot standby storage cluster. Valid values:
	//
	// 	- **auto*	- (default): The zone is automatically selected.
	//
	// >  You can use the default value when HotStandbyCluster is set to on. If HotStandbyCluster is set to equal, specify the zone of the hot standby storage cluster. You can call the [DescribeZones](https://help.aliyun.com/document_detail/98041.html) operation to query information about zones.
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
