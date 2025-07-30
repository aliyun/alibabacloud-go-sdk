// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkAccessEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAccessEndpoints(v []*ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) *ListNetworkAccessEndpointsResponseBody
	GetNetworkAccessEndpoints() []*ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints
	SetNextToken(v string) *ListNetworkAccessEndpointsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListNetworkAccessEndpointsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListNetworkAccessEndpointsResponseBody
	GetTotalCount() *int64
}

type ListNetworkAccessEndpointsResponseBody struct {
	NetworkAccessEndpoints []*ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints `json:"NetworkAccessEndpoints,omitempty" xml:"NetworkAccessEndpoints,omitempty" type:"Repeated"`
	// 本次调用返回的查询凭证（Token）值，用于下一次翻页查询。
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNetworkAccessEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointsResponseBody) GetNetworkAccessEndpoints() []*ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	return s.NetworkAccessEndpoints
}

func (s *ListNetworkAccessEndpointsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNetworkAccessEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNetworkAccessEndpointsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListNetworkAccessEndpointsResponseBody) SetNetworkAccessEndpoints(v []*ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) *ListNetworkAccessEndpointsResponseBody {
	s.NetworkAccessEndpoints = v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBody) SetNextToken(v string) *ListNetworkAccessEndpointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBody) SetRequestId(v string) *ListNetworkAccessEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBody) SetTotalCount(v int64) *ListNetworkAccessEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints struct {
	// 专属网络端点创建时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 专属网络端点ID。
	//
	// example:
	//
	// nae_examplexxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// 专属网络端点名称。
	//
	// example:
	//
	// xx业务VPC访问端点
	NetworkAccessEndpointName *string `json:"NetworkAccessEndpointName,omitempty" xml:"NetworkAccessEndpointName,omitempty"`
	// 专属网络端点连接的类型。
	//
	// example:
	//
	// private
	NetworkAccessEndpointType *string `json:"NetworkAccessEndpointType,omitempty" xml:"NetworkAccessEndpointType,omitempty"`
	// 专属网络端点使用的安全组ID。
	//
	// example:
	//
	// sg-examplexxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 专属网络端点状态。
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 专属网络端点最近更新时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 专属网络端点连接的指定vSwitch列表。
	//
	// example:
	//
	// vsw-examplexxx
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// 专属网络端点连接的VpcID。
	//
	// example:
	//
	// vpc-examplexxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 专属网络端点连接的Vpc所属地域。
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
}

func (s ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetNetworkAccessEndpointName() *string {
	return s.NetworkAccessEndpointName
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetNetworkAccessEndpointType() *string {
	return s.NetworkAccessEndpointType
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetStatus() *string {
	return s.Status
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetVpcId() *string {
	return s.VpcId
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetVpcRegionId() *string {
	return s.VpcRegionId
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetCreateTime(v int64) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.CreateTime = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetInstanceId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.InstanceId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetNetworkAccessEndpointId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetNetworkAccessEndpointName(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.NetworkAccessEndpointName = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetNetworkAccessEndpointType(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.NetworkAccessEndpointType = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetSecurityGroupId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.SecurityGroupId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetStatus(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.Status = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetUpdateTime(v int64) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.UpdateTime = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetVSwitchIds(v []*string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.VSwitchIds = v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetVpcId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.VpcId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetVpcRegionId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.VpcRegionId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) Validate() error {
	return dara.Validate(s)
}
