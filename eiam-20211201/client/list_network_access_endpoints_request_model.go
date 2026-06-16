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
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return on each page. The maximum value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
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
	// This parameter does not take effect when NetworkAccessEndpointType is set to shared.
	//
	// example:
	//
	// running
	NetworkAccessEndpointStatus *string `json:"NetworkAccessEndpointStatus,omitempty" xml:"NetworkAccessEndpointStatus,omitempty"`
	// The type of the network endpoint. Valid values:
	//
	// - shared: a shared network endpoint.
	//
	// - private: a private network endpoint.
	//
	// The default value is private.
	//
	// example:
	//
	// private
	NetworkAccessEndpointType *string `json:"NetworkAccessEndpointType,omitempty" xml:"NetworkAccessEndpointType,omitempty"`
	// The token used for the next query. Set this parameter to the NextToken value returned from the previous API call. Leave this parameter empty for the first query.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the VPC to which the private network endpoint is connected. This parameter does not take effect when NetworkAccessEndpointType is set to shared.
	//
	// example:
	//
	// vpc-examplexxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The region ID of the VPC to which the private network endpoint is connected. The value of this parameter must be a region returned by the ListNetworkAccessEndpointAvailableRegions operation. This parameter does not take effect when NetworkAccessEndpointType is set to shared.
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
