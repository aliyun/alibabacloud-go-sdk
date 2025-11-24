// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshProxyStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeServiceMeshProxyStatusResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeServiceMeshProxyStatusResponseBody
	GetMessage() *string
	SetProxyStatus(v []*DescribeServiceMeshProxyStatusResponseBodyProxyStatus) *DescribeServiceMeshProxyStatusResponseBody
	GetProxyStatus() []*DescribeServiceMeshProxyStatusResponseBodyProxyStatus
	SetRequestId(v string) *DescribeServiceMeshProxyStatusResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeServiceMeshProxyStatusResponseBody
	GetSuccess() *string
}

type DescribeServiceMeshProxyStatusResponseBody struct {
	// The status code. Valid values:
	//
	// `200`: The operation is successful.
	//
	// 	- `403`: You are not authorized to perform this operation.
	//
	// 	- `503`: A backend server error occurs.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The information about the status of the proxies on the data plane.
	ProxyStatus []*DescribeServiceMeshProxyStatusResponseBodyProxyStatus `json:"ProxyStatus,omitempty" xml:"ProxyStatus,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// success
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeServiceMeshProxyStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshProxyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshProxyStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeServiceMeshProxyStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeServiceMeshProxyStatusResponseBody) GetProxyStatus() []*DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	return s.ProxyStatus
}

func (s *DescribeServiceMeshProxyStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceMeshProxyStatusResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeServiceMeshProxyStatusResponseBody) SetCode(v string) *DescribeServiceMeshProxyStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBody) SetMessage(v string) *DescribeServiceMeshProxyStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBody) SetProxyStatus(v []*DescribeServiceMeshProxyStatusResponseBodyProxyStatus) *DescribeServiceMeshProxyStatusResponseBody {
	s.ProxyStatus = v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBody) SetRequestId(v string) *DescribeServiceMeshProxyStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBody) SetSuccess(v string) *DescribeServiceMeshProxyStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBody) Validate() error {
	if s.ProxyStatus != nil {
		for _, item := range s.ProxyStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeServiceMeshProxyStatusResponseBodyProxyStatus struct {
	// The update status of the proxy. Valid values:
	//
	// 	- `SYNCED`: The status of the proxy is updated.
	//
	// 	- `NOT SENT`: The status of the proxy is not updated.
	//
	// 	- `STALE (Never Acknowledged)`: Istiod has sent multiple requests to the Envoy proxy to update the status of the proxy but receives no response.
	//
	// 	- `STALE`: Istiod has sent a request to the Envoy proxy to update the status of the proxy but receives no response.
	//
	// example:
	//
	// SYNCED
	ClusterSynced *string `json:"ClusterSynced,omitempty" xml:"ClusterSynced,omitempty"`
	// The percentage of the updated endpoints.
	//
	// example:
	//
	// 1
	EndpointPercent *string `json:"EndpointPercent,omitempty" xml:"EndpointPercent,omitempty"`
	// The update status of the endpoint. Valid values:
	//
	// 	- `SYNCED`: The status of the endpoint is updated.
	//
	// 	- `NOT SENT`: The status of the endpoint is not updated.
	//
	// 	- `STALE (Never Acknowledged)`: Istiod has sent multiple requests to the Envoy proxy to update the status of the endpoint but receives no response.
	//
	// 	- `STALE`: Istiod has sent a request to the Envoy proxy to update the status of the endpoint but receives no response.
	//
	// example:
	//
	// SYNCED
	EndpointSynced *string `json:"EndpointSynced,omitempty" xml:"EndpointSynced,omitempty"`
	// The version of Istiod.
	//
	// example:
	//
	// 1.9.7
	IstioVersion *string `json:"IstioVersion,omitempty" xml:"IstioVersion,omitempty"`
	// The update status of the listener. Valid values:
	//
	// 	- `SYNCED`: The status of the listener is updated.
	//
	// 	- `NOT SENT`: The status of the listener is not updated.
	//
	// 	- `STALE (Never Acknowledged)`: Istiod has sent multiple requests to the Envoy proxy to update the status of the listener but receives no response.
	//
	// 	- `STALE`: Istiod has sent a request to the Envoy proxy to update the status of the listener but receives no response.
	//
	// example:
	//
	// SYNCED
	ListenerSynced *string `json:"ListenerSynced,omitempty" xml:"ListenerSynced,omitempty"`
	// The ID of the proxy on the data plane.
	//
	// example:
	//
	// 119q****
	ProxyId *string `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// The version number of a proxy on the data plane.
	//
	// example:
	//
	// 1.9.7
	ProxyVersion *string `json:"ProxyVersion,omitempty" xml:"ProxyVersion,omitempty"`
	// The update status of the route. Valid values:
	//
	// 	- `SYNCED`: The status of the route is updated.
	//
	// 	- `NOT SENT`: The status of the route is not updated.
	//
	// 	- `STALE (Never Acknowledged)`: Istiod has sent multiple requests to the Envoy proxy to update the status of the route but receives no response.
	//
	// 	- `STALE`: Istiod has sent a request to the Envoy proxy to update the status of the route but receives no response.
	//
	// example:
	//
	// SYNCED
	RouteSynced *string `json:"RouteSynced,omitempty" xml:"RouteSynced,omitempty"`
}

func (s DescribeServiceMeshProxyStatusResponseBodyProxyStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshProxyStatusResponseBodyProxyStatus) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) GetClusterSynced() *string {
	return s.ClusterSynced
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) GetEndpointPercent() *string {
	return s.EndpointPercent
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) GetEndpointSynced() *string {
	return s.EndpointSynced
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) GetIstioVersion() *string {
	return s.IstioVersion
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) GetListenerSynced() *string {
	return s.ListenerSynced
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) GetProxyId() *string {
	return s.ProxyId
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) GetProxyVersion() *string {
	return s.ProxyVersion
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) GetRouteSynced() *string {
	return s.RouteSynced
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) SetClusterSynced(v string) *DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	s.ClusterSynced = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) SetEndpointPercent(v string) *DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	s.EndpointPercent = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) SetEndpointSynced(v string) *DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	s.EndpointSynced = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) SetIstioVersion(v string) *DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	s.IstioVersion = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) SetListenerSynced(v string) *DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	s.ListenerSynced = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) SetProxyId(v string) *DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	s.ProxyId = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) SetProxyVersion(v string) *DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	s.ProxyVersion = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) SetRouteSynced(v string) *DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	s.RouteSynced = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) Validate() error {
	return dara.Validate(s)
}
