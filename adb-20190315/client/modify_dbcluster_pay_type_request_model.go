// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterPayTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbClusterId(v string) *ModifyDBClusterPayTypeRequest
	GetDbClusterId() *string
	SetPayType(v string) *ModifyDBClusterPayTypeRequest
	GetPayType() *string
	SetPeriod(v string) *ModifyDBClusterPayTypeRequest
	GetPeriod() *string
	SetRegionId(v string) *ModifyDBClusterPayTypeRequest
	GetRegionId() *string
	SetUsedTime(v string) *ModifyDBClusterPayTypeRequest
	GetUsedTime() *string
}

type ModifyDBClusterPayTypeRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp11q28kvl688****
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription type of the subscription cluster. Valid values:
	//
	// 	- **Year**: subscription on a yearly basis.
	//
	// 	- **Month**: subscription on a monthly basis.
	//
	// > This parameter must be specified when PayType is set to Prepaid.
	//
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The subscription duration of the subscription cluster.
	//
	// 	- Valid values when Period is set to Year: 1, 2, 3, and 5 (integer).
	//
	// 	- Valid values when Period is set to Month: 1 to 11 (integer).
	//
	// >
	//
	// 	- This parameter must be specified when PayType is set to Prepaid.
	//
	// 	- Longer subscription durations offer more savings. Purchasing a cluster for one year is more cost-effective than purchasing the cluster for 10 or 11 months.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s ModifyDBClusterPayTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterPayTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterPayTypeRequest) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *ModifyDBClusterPayTypeRequest) GetPayType() *string {
	return s.PayType
}

func (s *ModifyDBClusterPayTypeRequest) GetPeriod() *string {
	return s.Period
}

func (s *ModifyDBClusterPayTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBClusterPayTypeRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *ModifyDBClusterPayTypeRequest) SetDbClusterId(v string) *ModifyDBClusterPayTypeRequest {
	s.DbClusterId = &v
	return s
}

func (s *ModifyDBClusterPayTypeRequest) SetPayType(v string) *ModifyDBClusterPayTypeRequest {
	s.PayType = &v
	return s
}

func (s *ModifyDBClusterPayTypeRequest) SetPeriod(v string) *ModifyDBClusterPayTypeRequest {
	s.Period = &v
	return s
}

func (s *ModifyDBClusterPayTypeRequest) SetRegionId(v string) *ModifyDBClusterPayTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterPayTypeRequest) SetUsedTime(v string) *ModifyDBClusterPayTypeRequest {
	s.UsedTime = &v
	return s
}

func (s *ModifyDBClusterPayTypeRequest) Validate() error {
	return dara.Validate(s)
}
