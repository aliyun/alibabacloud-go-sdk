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
	// A collection of network endpoints.
	NetworkAccessEndpoints []*ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints `json:"NetworkAccessEndpoints,omitempty" xml:"NetworkAccessEndpoints,omitempty" type:"Repeated"`
	// The token returned for the next query.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
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
	if s.NetworkAccessEndpoints != nil {
		for _, item := range s.NetworkAccessEndpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints struct {
	// The time when the network endpoint was created. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network endpoint ID.
	//
	// example:
	//
	// nae_examplexxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// The name of the network endpoint.
	//
	// example:
	//
	// VPC access endpoint for xx service
	NetworkAccessEndpointName *string `json:"NetworkAccessEndpointName,omitempty" xml:"NetworkAccessEndpointName,omitempty"`
	// The type of the network endpoint. Valid values:
	//
	// - shared: a shared network endpoint.
	//
	// - private: a private network endpoint.
	//
	// example:
	//
	// private
	NetworkAccessEndpointType *string `json:"NetworkAccessEndpointType,omitempty" xml:"NetworkAccessEndpointType,omitempty"`
	// The ID of the security group used by the private network endpoint.
	//
	// example:
	//
	// sg-examplexxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The status of the network endpoint. Valid values:
	//
	// - pending: The endpoint is pending initialization.
	//
	// - creating: The endpoint is being created.
	//
	// - running: The endpoint is running.
	//
	// - deleting: The endpoint is being deleted.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the network endpoint was last updated. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// A list of vSwitches to which the private network endpoint is connected.
	//
	// example:
	//
	// vsw-examplexxx
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the VPC to which the private network endpoint is connected.
	//
	// example:
	//
	// vpc-examplexxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The region ID of the VPC to which the private network endpoint is connected.
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
