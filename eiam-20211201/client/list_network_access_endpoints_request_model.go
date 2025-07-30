// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkAccessEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListNetworkAccessEndpointsRequest
	GetInstanceId() *string
	SetMaxResults(v int64) *ListNetworkAccessEndpointsRequest
	GetMaxResults() *int64
	SetNetworkAccessEndpointStatus(v string) *ListNetworkAccessEndpointsRequest
	GetNetworkAccessEndpointStatus() *string
	SetNetworkAccessEndpointType(v string) *ListNetworkAccessEndpointsRequest
	GetNetworkAccessEndpointType() *string
	SetNextToken(v string) *ListNetworkAccessEndpointsRequest
	GetNextToken() *string
	SetVpcId(v string) *ListNetworkAccessEndpointsRequest
	GetVpcId() *string
	SetVpcRegionId(v string) *ListNetworkAccessEndpointsRequest
	GetVpcRegionId() *string
}

type ListNetworkAccessEndpointsRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 分页查询时每页行数。默认值为20，最大值为100。
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 专属网络端点连接的状态。NetworkAccessEndpointType取值为shared时不生效。
	//
	// example:
	//
	// running
	NetworkAccessEndpointStatus *string `json:"NetworkAccessEndpointStatus,omitempty" xml:"NetworkAccessEndpointStatus,omitempty"`
	// 专属网络端点连接的类型。取值可选范围：1. private - 专属网络端点；2. shared - 共享网络端点
	//
	// example:
	//
	// private
	NetworkAccessEndpointType *string `json:"NetworkAccessEndpointType,omitempty" xml:"NetworkAccessEndpointType,omitempty"`
	// 查询凭证（Token），取值为上一次API调用返回的NextToken参数值。
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 专属网络端点连接的Vpc ID。NetworkAccessEndpointType取值为shared时不生效。
	//
	// example:
	//
	// vpc-examplexxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 专属网络端点连接的Vpc所属地域，该地域取值必须在ListNetworkAccessEndpointAvailableRegions接口中返回。NetworkAccessEndpointType取值为shared时不生效。
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
}

func (s ListNetworkAccessEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNetworkAccessEndpointsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListNetworkAccessEndpointsRequest) GetNetworkAccessEndpointStatus() *string {
	return s.NetworkAccessEndpointStatus
}

func (s *ListNetworkAccessEndpointsRequest) GetNetworkAccessEndpointType() *string {
	return s.NetworkAccessEndpointType
}

func (s *ListNetworkAccessEndpointsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNetworkAccessEndpointsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListNetworkAccessEndpointsRequest) GetVpcRegionId() *string {
	return s.VpcRegionId
}

func (s *ListNetworkAccessEndpointsRequest) SetInstanceId(v string) *ListNetworkAccessEndpointsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNetworkAccessEndpointsRequest) SetMaxResults(v int64) *ListNetworkAccessEndpointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNetworkAccessEndpointsRequest) SetNetworkAccessEndpointStatus(v string) *ListNetworkAccessEndpointsRequest {
	s.NetworkAccessEndpointStatus = &v
	return s
}

func (s *ListNetworkAccessEndpointsRequest) SetNetworkAccessEndpointType(v string) *ListNetworkAccessEndpointsRequest {
	s.NetworkAccessEndpointType = &v
	return s
}

func (s *ListNetworkAccessEndpointsRequest) SetNextToken(v string) *ListNetworkAccessEndpointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNetworkAccessEndpointsRequest) SetVpcId(v string) *ListNetworkAccessEndpointsRequest {
	s.VpcId = &v
	return s
}

func (s *ListNetworkAccessEndpointsRequest) SetVpcRegionId(v string) *ListNetworkAccessEndpointsRequest {
	s.VpcRegionId = &v
	return s
}

func (s *ListNetworkAccessEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
