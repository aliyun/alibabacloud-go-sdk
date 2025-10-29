// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListVpcEndpointsResponseBody
	GetRequestId() *string
	SetResult(v []*ListVpcEndpointsResponseBodyResult) *ListVpcEndpointsResponseBody
	GetResult() []*ListVpcEndpointsResponseBodyResult
}

type ListVpcEndpointsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC47D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the endpoints.
	Result []*ListVpcEndpointsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcEndpointsResponseBody) GetResult() []*ListVpcEndpointsResponseBodyResult {
	return s.Result
}

func (s *ListVpcEndpointsResponseBody) SetRequestId(v string) *ListVpcEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointsResponseBody) SetResult(v []*ListVpcEndpointsResponseBodyResult) *ListVpcEndpointsResponseBody {
	s.Result = v
	return s
}

func (s *ListVpcEndpointsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVpcEndpointsResponseBodyResult struct {
	// The status of the endpoint connection. Valid values:
	//
	// 	- Pending
	//
	// 	- Connecting
	//
	// 	- Connected
	//
	// 	- Disconnecting
	//
	// 	- Disconnected
	//
	// 	- Deleting
	//
	// 	- ServiceDeleted
	//
	// example:
	//
	// Disconnected
	ConnectionStatus *string `json:"connectionStatus,omitempty" xml:"connectionStatus,omitempty"`
	// The time when the endpoint was created.
	//
	// example:
	//
	// 2021-07-22T01:19:24Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The business status of the endpoint. Valid values:
	//
	// 	- Normal
	//
	// 	- FinancialLocked
	//
	// example:
	//
	// Normal
	EndpointBusinessStatus *string `json:"endpointBusinessStatus,omitempty" xml:"endpointBusinessStatus,omitempty"`
	// The domain name of the endpoint. The domain name is used for connection configuration.
	//
	// example:
	//
	// ep-bp18s6wy9420wdi4****.epsrv-bp1bz3efowa4kc0****.cn-hangzhou.privatelink.aliyuncs.com
	EndpointDomain *string `json:"endpointDomain,omitempty" xml:"endpointDomain,omitempty"`
	// The ID of the endpoint.
	//
	// example:
	//
	// ep-bp1tah7zbrwmkjef****
	EndpointId *string `json:"endpointId,omitempty" xml:"endpointId,omitempty"`
	// The name of the endpoint.
	//
	// example:
	//
	// test
	EndpointName *string `json:"endpointName,omitempty" xml:"endpointName,omitempty"`
	// The status of the endpoint. Valid values:
	//
	// 	- Creating
	//
	// 	- Active
	//
	// 	- Pending
	//
	// 	- Deleting
	//
	// example:
	//
	// Active
	EndpointStatus *string `json:"endpointStatus,omitempty" xml:"endpointStatus,omitempty"`
	// The ID of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// epsrv-bp1w0p3jdirbfmt6****
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The name of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-hangzhou.epsrv-bp1w0p3jdirbfmt6****
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s ListVpcEndpointsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsResponseBodyResult) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *ListVpcEndpointsResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListVpcEndpointsResponseBodyResult) GetEndpointBusinessStatus() *string {
	return s.EndpointBusinessStatus
}

func (s *ListVpcEndpointsResponseBodyResult) GetEndpointDomain() *string {
	return s.EndpointDomain
}

func (s *ListVpcEndpointsResponseBodyResult) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListVpcEndpointsResponseBodyResult) GetEndpointName() *string {
	return s.EndpointName
}

func (s *ListVpcEndpointsResponseBodyResult) GetEndpointStatus() *string {
	return s.EndpointStatus
}

func (s *ListVpcEndpointsResponseBodyResult) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListVpcEndpointsResponseBodyResult) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListVpcEndpointsResponseBodyResult) SetConnectionStatus(v string) *ListVpcEndpointsResponseBodyResult {
	s.ConnectionStatus = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) SetCreateTime(v string) *ListVpcEndpointsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) SetEndpointBusinessStatus(v string) *ListVpcEndpointsResponseBodyResult {
	s.EndpointBusinessStatus = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) SetEndpointDomain(v string) *ListVpcEndpointsResponseBodyResult {
	s.EndpointDomain = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) SetEndpointId(v string) *ListVpcEndpointsResponseBodyResult {
	s.EndpointId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) SetEndpointName(v string) *ListVpcEndpointsResponseBodyResult {
	s.EndpointName = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) SetEndpointStatus(v string) *ListVpcEndpointsResponseBodyResult {
	s.EndpointStatus = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) SetServiceId(v string) *ListVpcEndpointsResponseBodyResult {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) SetServiceName(v string) *ListVpcEndpointsResponseBodyResult {
	s.ServiceName = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
