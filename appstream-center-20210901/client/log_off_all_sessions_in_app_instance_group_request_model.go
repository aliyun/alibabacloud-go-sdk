// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogOffAllSessionsInAppInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *LogOffAllSessionsInAppInstanceGroupRequest
	GetAppInstanceGroupId() *string
	SetProductType(v string) *LogOffAllSessionsInAppInstanceGroupRequest
	GetProductType() *string
}

type LogOffAllSessionsInAppInstanceGroupRequest struct {
	// The delivery group ID.
	//
	// Note:
	//
	// - The delivery group must meet the following status requirements:
	//
	// 1. ChargeType=PostPaid
	//
	// 2. NodePool.StrategyType=NODE_SCALING_BY_SCHEDULE
	//
	// 3. Status=PUBLISHED
	//
	// 4. The current time is not within the NodePool.RecurrenceSchedules window.
	//
	// - You can call the ListAppInstanceGroup or GetAppInstanceGroup operation to query the current status of the delivery group.
	//
	// - You can call the CreateAppInstanceGroup operation to create a delivery group that meets the requirements.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The product type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s LogOffAllSessionsInAppInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s LogOffAllSessionsInAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *LogOffAllSessionsInAppInstanceGroupRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *LogOffAllSessionsInAppInstanceGroupRequest) GetProductType() *string {
	return s.ProductType
}

func (s *LogOffAllSessionsInAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *LogOffAllSessionsInAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupRequest) SetProductType(v string) *LogOffAllSessionsInAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupRequest) Validate() error {
	return dara.Validate(s)
}
