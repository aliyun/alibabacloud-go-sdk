// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceModels(v []*ListAppInstancesResponseBodyAppInstanceModels) *ListAppInstancesResponseBody
	GetAppInstanceModels() []*ListAppInstancesResponseBodyAppInstanceModels
	SetPageNumber(v int32) *ListAppInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAppInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAppInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAppInstancesResponseBody
	GetTotalCount() *int32
}

type ListAppInstancesResponseBody struct {
	// The app instances.
	AppInstanceModels []*ListAppInstancesResponseBodyAppInstanceModels `json:"AppInstanceModels,omitempty" xml:"AppInstanceModels,omitempty" type:"Repeated"`
	// The page number of the returned page. We recommend that you configure this parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page. The value cannot be greater than `100`. We recommend that you configure this parameter.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 18
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppInstancesResponseBody) GetAppInstanceModels() []*ListAppInstancesResponseBodyAppInstanceModels {
	return s.AppInstanceModels
}

func (s *ListAppInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAppInstancesResponseBody) SetAppInstanceModels(v []*ListAppInstancesResponseBodyAppInstanceModels) *ListAppInstancesResponseBody {
	s.AppInstanceModels = v
	return s
}

func (s *ListAppInstancesResponseBody) SetPageNumber(v int32) *ListAppInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetPageSize(v int32) *ListAppInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetRequestId(v string) *ListAppInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetTotalCount(v int32) *ListAppInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAppInstancesResponseBody) Validate() error {
	if s.AppInstanceModels != nil {
		for _, item := range s.AppInstanceModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppInstancesResponseBodyAppInstanceModels struct {
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-dk8p95irqfst9****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the application instance.
	//
	// example:
	//
	// ai-8dl7dzchklmka****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The information about the binding between the application instance and end users.
	BindInfo *ListAppInstancesResponseBodyAppInstanceModelsBindInfo `json:"BindInfo,omitempty" xml:"BindInfo,omitempty" type:"Struct"`
	// The billing method of the app instance. Valid values:
	//
	// 	- **PrePaid**: subscription.
	//
	// 	- **PostPaid**: pay-as-you-go
	//
	// >  This parameter is returned only if the ChargeResourceMode parameter of the delivery group to which the app instance belongs is set to Node.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the application instance was created.
	//
	// example:
	//
	// 2023-03-07T20:29:19.000+08:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the application instance was updated.
	//
	// example:
	//
	// 2023-03-07T20:29:19.000+08:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The public IP address associated with the primary NIC. This value is returned only if `StrategyType` is set to `Mixed`.
	//
	// example:
	//
	// 10.13.13.211
	MainEthPublicIp    *string `json:"MainEthPublicIp,omitempty" xml:"MainEthPublicIp,omitempty"`
	NetworkInterfaceIp *string `json:"NetworkInterfaceIp,omitempty" xml:"NetworkInterfaceIp,omitempty"`
	// The ID of the node on which the app instance runs.
	//
	// >  This parameter is returned only if the ChargeResourceMode parameter of the delivery group to which the app instance belongs is set to Node.
	//
	// example:
	//
	// i-bp13********
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The session status. This parameter is returned only if the application instance is in the `RUNNING` state.
	//
	// Valid values:
	//
	// 	- disconnect: disconnected
	//
	// 	- connect: connected
	//
	// example:
	//
	// connect
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	// The status of the application instance.
	//
	// example:
	//
	// BOUND
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAppInstancesResponseBodyAppInstanceModels) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstancesResponseBodyAppInstanceModels) GoString() string {
	return s.String()
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) GetBindInfo() *ListAppInstancesResponseBodyAppInstanceModelsBindInfo {
	return s.BindInfo
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) GetMainEthPublicIp() *string {
	return s.MainEthPublicIp
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) GetNetworkInterfaceIp() *string {
	return s.NetworkInterfaceIp
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) GetNodeId() *string {
	return s.NodeId
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) GetStatus() *string {
	return s.Status
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetAppInstanceGroupId(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetAppInstanceId(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.AppInstanceId = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetBindInfo(v *ListAppInstancesResponseBodyAppInstanceModelsBindInfo) *ListAppInstancesResponseBodyAppInstanceModels {
	s.BindInfo = v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetChargeType(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.ChargeType = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetGmtCreate(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.GmtCreate = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetGmtModified(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.GmtModified = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetMainEthPublicIp(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.MainEthPublicIp = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetNetworkInterfaceIp(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.NetworkInterfaceIp = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetNodeId(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.NodeId = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetSessionStatus(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.SessionStatus = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetStatus(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.Status = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) Validate() error {
	if s.BindInfo != nil {
		if err := s.BindInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppInstancesResponseBodyAppInstanceModelsBindInfo struct {
	// The ID of the end user that is bound to the application instance.
	//
	// example:
	//
	// app.test
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The use duration of the application instance. Unit: seconds.
	//
	// example:
	//
	// 2000
	UsageDuration *int64 `json:"UsageDuration,omitempty" xml:"UsageDuration,omitempty"`
}

func (s ListAppInstancesResponseBodyAppInstanceModelsBindInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstancesResponseBodyAppInstanceModelsBindInfo) GoString() string {
	return s.String()
}

func (s *ListAppInstancesResponseBodyAppInstanceModelsBindInfo) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListAppInstancesResponseBodyAppInstanceModelsBindInfo) GetUsageDuration() *int64 {
	return s.UsageDuration
}

func (s *ListAppInstancesResponseBodyAppInstanceModelsBindInfo) SetEndUserId(v string) *ListAppInstancesResponseBodyAppInstanceModelsBindInfo {
	s.EndUserId = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModelsBindInfo) SetUsageDuration(v int64) *ListAppInstancesResponseBodyAppInstanceModelsBindInfo {
	s.UsageDuration = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModelsBindInfo) Validate() error {
	return dara.Validate(s)
}
