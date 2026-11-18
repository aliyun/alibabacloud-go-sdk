// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTairAIGatewayInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailCount(v int32) *CreateTairAIGatewayInstanceResponseBody
	GetFailCount() *int32
	SetInstanceId(v string) *CreateTairAIGatewayInstanceResponseBody
	GetInstanceId() *string
	SetInstanceResults(v *CreateTairAIGatewayInstanceResponseBodyInstanceResults) *CreateTairAIGatewayInstanceResponseBody
	GetInstanceResults() *CreateTairAIGatewayInstanceResponseBodyInstanceResults
	SetNodeId(v string) *CreateTairAIGatewayInstanceResponseBody
	GetNodeId() *string
	SetOrderId(v int64) *CreateTairAIGatewayInstanceResponseBody
	GetOrderId() *int64
	SetRegionId(v string) *CreateTairAIGatewayInstanceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *CreateTairAIGatewayInstanceResponseBody
	GetRequestId() *string
	SetSuccessCount(v int32) *CreateTairAIGatewayInstanceResponseBody
	GetSuccessCount() *int32
	SetTotalCount(v int32) *CreateTairAIGatewayInstanceResponseBody
	GetTotalCount() *int32
	SetVClusterId(v string) *CreateTairAIGatewayInstanceResponseBody
	GetVClusterId() *string
	SetVkName(v string) *CreateTairAIGatewayInstanceResponseBody
	GetVkName() *string
	SetZoneId(v string) *CreateTairAIGatewayInstanceResponseBody
	GetZoneId() *string
}

type CreateTairAIGatewayInstanceResponseBody struct {
	// The number of instances that failed to be created in a batch. This parameter is not returned when a single instance is created.
	//
	// example:
	//
	// true
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// tg-bp1zxszhcgatnx****
	InstanceId      *string                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceResults *CreateTairAIGatewayInstanceResponseBodyInstanceResults `json:"InstanceResults,omitempty" xml:"InstanceResults,omitempty" type:"Struct"`
	// **[Deprecated]*	- This parameter is deprecated. The ID of the virtual node.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 20881824000****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of instances successfully created in a batch. This parameter is not returned when a single instance is created.
	//
	// example:
	//
	// 2
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The total number of instances created in a batch. This parameter is not returned when a single instance is created.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. The ID of the virtual cluster.
	//
	// example:
	//
	// vc-16965a9267*****-*****
	VClusterId *string `json:"VClusterId,omitempty" xml:"VClusterId,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. The instance ID of the VCluster virtual cluster where the VNode resides.
	//
	// example:
	//
	// tk-2ze4bba3c8fe****
	VkName *string `json:"VkName,omitempty" xml:"VkName,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-beijing-l
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateTairAIGatewayInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTairAIGatewayInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTairAIGatewayInstanceResponseBody) GetFailCount() *int32 {
	return s.FailCount
}

func (s *CreateTairAIGatewayInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTairAIGatewayInstanceResponseBody) GetInstanceResults() *CreateTairAIGatewayInstanceResponseBodyInstanceResults {
	return s.InstanceResults
}

func (s *CreateTairAIGatewayInstanceResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateTairAIGatewayInstanceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateTairAIGatewayInstanceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTairAIGatewayInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTairAIGatewayInstanceResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *CreateTairAIGatewayInstanceResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *CreateTairAIGatewayInstanceResponseBody) GetVClusterId() *string {
	return s.VClusterId
}

func (s *CreateTairAIGatewayInstanceResponseBody) GetVkName() *string {
	return s.VkName
}

func (s *CreateTairAIGatewayInstanceResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateTairAIGatewayInstanceResponseBody) SetFailCount(v int32) *CreateTairAIGatewayInstanceResponseBody {
	s.FailCount = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBody) SetInstanceId(v string) *CreateTairAIGatewayInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBody) SetInstanceResults(v *CreateTairAIGatewayInstanceResponseBodyInstanceResults) *CreateTairAIGatewayInstanceResponseBody {
	s.InstanceResults = v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBody) SetNodeId(v string) *CreateTairAIGatewayInstanceResponseBody {
	s.NodeId = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBody) SetOrderId(v int64) *CreateTairAIGatewayInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBody) SetRegionId(v string) *CreateTairAIGatewayInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBody) SetRequestId(v string) *CreateTairAIGatewayInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBody) SetSuccessCount(v int32) *CreateTairAIGatewayInstanceResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBody) SetTotalCount(v int32) *CreateTairAIGatewayInstanceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBody) SetVClusterId(v string) *CreateTairAIGatewayInstanceResponseBody {
	s.VClusterId = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBody) SetVkName(v string) *CreateTairAIGatewayInstanceResponseBody {
	s.VkName = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBody) SetZoneId(v string) *CreateTairAIGatewayInstanceResponseBody {
	s.ZoneId = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBody) Validate() error {
	if s.InstanceResults != nil {
		if err := s.InstanceResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTairAIGatewayInstanceResponseBodyInstanceResults struct {
	InstanceResults []*CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults `json:"instanceResults,omitempty" xml:"instanceResults,omitempty" type:"Repeated"`
}

func (s CreateTairAIGatewayInstanceResponseBodyInstanceResults) String() string {
	return dara.Prettify(s)
}

func (s CreateTairAIGatewayInstanceResponseBodyInstanceResults) GoString() string {
	return s.String()
}

func (s *CreateTairAIGatewayInstanceResponseBodyInstanceResults) GetInstanceResults() []*CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults {
	return s.InstanceResults
}

func (s *CreateTairAIGatewayInstanceResponseBodyInstanceResults) SetInstanceResults(v []*CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults) *CreateTairAIGatewayInstanceResponseBodyInstanceResults {
	s.InstanceResults = v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBodyInstanceResults) Validate() error {
	if s.InstanceResults != nil {
		for _, item := range s.InstanceResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId      *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults) String() string {
	return dara.Prettify(s)
}

func (s CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults) GoString() string {
	return s.String()
}

func (s *CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults) SetErrorCode(v string) *CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults {
	s.ErrorCode = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults) SetErrorMessage(v string) *CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults {
	s.ErrorMessage = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults) SetInstanceId(v string) *CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults {
	s.InstanceId = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults) SetOrderId(v int64) *CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults {
	s.OrderId = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults) SetSuccess(v bool) *CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults {
	s.Success = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponseBodyInstanceResultsInstanceResults) Validate() error {
	return dara.Validate(s)
}
