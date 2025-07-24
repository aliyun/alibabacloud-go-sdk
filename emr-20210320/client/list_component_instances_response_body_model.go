// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponentInstances(v []*ListComponentInstancesResponseBodyComponentInstances) *ListComponentInstancesResponseBody
	GetComponentInstances() []*ListComponentInstancesResponseBodyComponentInstances
	SetMaxResults(v int32) *ListComponentInstancesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListComponentInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListComponentInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListComponentInstancesResponseBody
	GetTotalCount() *int32
}

type ListComponentInstancesResponseBody struct {
	ComponentInstances []*ListComponentInstancesResponseBodyComponentInstances `json:"ComponentInstances,omitempty" xml:"ComponentInstances,omitempty" type:"Repeated"`
	// 本次请求所返回的最大记录条数。
	//
	// example:
	//
	// 2
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 返回读取到的数据位置，空代表数据已经读取完毕。
	//
	// example:
	//
	// “”
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// 7345241A-014C-17D2-A3AC-C72771188F46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 本次请求条件下的数据总量。
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListComponentInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComponentInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListComponentInstancesResponseBody) GetComponentInstances() []*ListComponentInstancesResponseBodyComponentInstances {
	return s.ComponentInstances
}

func (s *ListComponentInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListComponentInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListComponentInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComponentInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListComponentInstancesResponseBody) SetComponentInstances(v []*ListComponentInstancesResponseBodyComponentInstances) *ListComponentInstancesResponseBody {
	s.ComponentInstances = v
	return s
}

func (s *ListComponentInstancesResponseBody) SetMaxResults(v int32) *ListComponentInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListComponentInstancesResponseBody) SetNextToken(v string) *ListComponentInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListComponentInstancesResponseBody) SetRequestId(v string) *ListComponentInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComponentInstancesResponseBody) SetTotalCount(v int32) *ListComponentInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListComponentInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListComponentInstancesResponseBodyComponentInstances struct {
	// 应用名称。
	//
	// example:
	//
	// KNOX
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// 组件服务状态，取值如下：
	//
	// - active：主服务
	//
	// - standby：备用服务。
	//
	// example:
	//
	// active
	BizState *string `json:"BizState,omitempty" xml:"BizState,omitempty"`
	// Commission状态，取值如下：
	//
	// - COMMISSIONED：已上线
	//
	// - COMMISSIONING：上线中
	//
	// - DECOMMISSIONED：已下线
	//
	// - DECOMMISSIONINPROGRESS：下线进程中
	//
	// - DECOMMISSIONFAILED：下线失败
	//
	// - INSERVICE：服务中
	//
	// - UNKNOWN：未知状态。
	//
	// <p>
	//
	// example:
	//
	// INSERVICE
	CommissionState *string `json:"CommissionState,omitempty" xml:"CommissionState,omitempty"`
	// 组件实例操作状态，取值如下：
	//
	// - WAITING：等待中
	//
	// - INSTALLING：安装中
	//
	// - INSTALLED：已安装
	//
	// - INSTALL_FAILED：安装失败
	//
	// - STARTING：启动中
	//
	// - STARTED：已启动
	//
	// - START_FAILED：启动失败
	//
	// - STOPPING：停止中
	//
	// - STOPPED：已停止
	//
	// - STOP_FAILED：停止失败
	//
	// example:
	//
	// STARTED
	ComponentInstanceState *string `json:"ComponentInstanceState,omitempty" xml:"ComponentInstanceState,omitempty"`
	// 组件名称。
	//
	// example:
	//
	// KNOX
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// 安装时间戳。
	//
	// example:
	//
	// 1628248947000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 期望状态，取值如下：
	//
	// - WAITING：等待中
	//
	// - INSTALLING：安装中
	//
	// - INSTALLED：已安装
	//
	// - INSTALL_FAILED：安装失败
	//
	// - STARTING：启动中
	//
	// - STARTED：已启动
	//
	// - START_FAILED：启动失败
	//
	// - STOPPING：停止中
	//
	// - STOPPED：已停止
	//
	// - STOP_FAILED：停止失败。
	//
	// example:
	//
	// STARTED
	DesiredState *string `json:"DesiredState,omitempty" xml:"DesiredState,omitempty"`
	// 节点ID。
	//
	// example:
	//
	// i-bp17yy050pxo01m2****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// 节点名称。
	//
	// example:
	//
	// emr-worker-1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListComponentInstancesResponseBodyComponentInstances) String() string {
	return dara.Prettify(s)
}

func (s ListComponentInstancesResponseBodyComponentInstances) GoString() string {
	return s.String()
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetBizState() *string {
	return s.BizState
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetCommissionState() *string {
	return s.CommissionState
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetComponentInstanceState() *string {
	return s.ComponentInstanceState
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetComponentName() *string {
	return s.ComponentName
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetDesiredState() *string {
	return s.DesiredState
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetNodeId() *string {
	return s.NodeId
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetNodeName() *string {
	return s.NodeName
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetApplicationName(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.ApplicationName = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetBizState(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.BizState = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetCommissionState(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.CommissionState = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetComponentInstanceState(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.ComponentInstanceState = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetComponentName(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.ComponentName = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetCreateTime(v int64) *ListComponentInstancesResponseBodyComponentInstances {
	s.CreateTime = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetDesiredState(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.DesiredState = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetNodeId(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.NodeId = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetNodeName(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.NodeName = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetZoneId(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.ZoneId = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) Validate() error {
	return dara.Validate(s)
}
