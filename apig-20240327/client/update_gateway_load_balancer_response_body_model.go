// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayLoadBalancerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateGatewayLoadBalancerResponseBody
	GetCode() *string
	SetData(v *UpdateGatewayLoadBalancerResponseBodyData) *UpdateGatewayLoadBalancerResponseBody
	GetData() *UpdateGatewayLoadBalancerResponseBodyData
	SetMessage(v string) *UpdateGatewayLoadBalancerResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayLoadBalancerResponseBody
	GetRequestId() *string
}

type UpdateGatewayLoadBalancerResponseBody struct {
	// example:
	//
	// 200
	Code *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateGatewayLoadBalancerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CEB8F71F-F889-599E-9D03-250978412350
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateGatewayLoadBalancerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayLoadBalancerResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateGatewayLoadBalancerResponseBody) GetData() *UpdateGatewayLoadBalancerResponseBodyData {
	return s.Data
}

func (s *UpdateGatewayLoadBalancerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayLoadBalancerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayLoadBalancerResponseBody) SetCode(v string) *UpdateGatewayLoadBalancerResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBody) SetData(v *UpdateGatewayLoadBalancerResponseBodyData) *UpdateGatewayLoadBalancerResponseBody {
	s.Data = v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBody) SetMessage(v string) *UpdateGatewayLoadBalancerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBody) SetRequestId(v string) *UpdateGatewayLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGatewayLoadBalancerResponseBodyData struct {
	// example:
	//
	// true
	EditEnable *bool `json:"editEnable,omitempty" xml:"editEnable,omitempty"`
	// example:
	//
	// 47.x.x.x
	LoadBalancerAddress *string `json:"loadBalancerAddress,omitempty" xml:"loadBalancerAddress,omitempty"`
	// example:
	//
	// lb-bp1xxxx
	LoadBalancerId *string `json:"loadBalancerId,omitempty" xml:"loadBalancerId,omitempty"`
	// example:
	//
	// my-clb
	LoadBalancerName *string `json:"loadBalancerName,omitempty" xml:"loadBalancerName,omitempty"`
	// example:
	//
	// CLB
	LoadBalancerType *string `json:"loadBalancerType,omitempty" xml:"loadBalancerType,omitempty"`
	// example:
	//
	// Internet
	NetworkType *string  `json:"networkType,omitempty" xml:"networkType,omitempty"`
	Ports       []*int32 `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ServiceWeight *int64 `json:"serviceWeight,omitempty" xml:"serviceWeight,omitempty"`
	// example:
	//
	// 状态描述
	StatusDescription  *string                                                        `json:"statusDescription,omitempty" xml:"statusDescription,omitempty"`
	VirtualServiceList []*UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList `json:"virtualServiceList,omitempty" xml:"virtualServiceList,omitempty" type:"Repeated"`
}

func (s UpdateGatewayLoadBalancerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayLoadBalancerResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) GetEditEnable() *bool {
	return s.EditEnable
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) GetLoadBalancerAddress() *string {
	return s.LoadBalancerAddress
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) GetLoadBalancerType() *string {
	return s.LoadBalancerType
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) GetPorts() []*int32 {
	return s.Ports
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) GetServiceWeight() *int64 {
	return s.ServiceWeight
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) GetStatusDescription() *string {
	return s.StatusDescription
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) GetVirtualServiceList() []*UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList {
	return s.VirtualServiceList
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) SetEditEnable(v bool) *UpdateGatewayLoadBalancerResponseBodyData {
	s.EditEnable = &v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) SetLoadBalancerAddress(v string) *UpdateGatewayLoadBalancerResponseBodyData {
	s.LoadBalancerAddress = &v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) SetLoadBalancerId(v string) *UpdateGatewayLoadBalancerResponseBodyData {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) SetLoadBalancerName(v string) *UpdateGatewayLoadBalancerResponseBodyData {
	s.LoadBalancerName = &v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) SetLoadBalancerType(v string) *UpdateGatewayLoadBalancerResponseBodyData {
	s.LoadBalancerType = &v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) SetNetworkType(v string) *UpdateGatewayLoadBalancerResponseBodyData {
	s.NetworkType = &v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) SetPorts(v []*int32) *UpdateGatewayLoadBalancerResponseBodyData {
	s.Ports = v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) SetServiceWeight(v int64) *UpdateGatewayLoadBalancerResponseBodyData {
	s.ServiceWeight = &v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) SetStatusDescription(v string) *UpdateGatewayLoadBalancerResponseBodyData {
	s.StatusDescription = &v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) SetVirtualServiceList(v []*UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList) *UpdateGatewayLoadBalancerResponseBodyData {
	s.VirtualServiceList = v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBodyData) Validate() error {
	if s.VirtualServiceList != nil {
		for _, item := range s.VirtualServiceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList struct {
	// example:
	//
	// 80
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// http
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// rsp-xxxx
	VirtualServiceGroupId *string `json:"virtualServiceGroupId,omitempty" xml:"virtualServiceGroupId,omitempty"`
	// example:
	//
	// 80-tcp
	VirtualServiceGroupName *string `json:"virtualServiceGroupName,omitempty" xml:"virtualServiceGroupName,omitempty"`
}

func (s UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList) GoString() string {
	return s.String()
}

func (s *UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList) GetPort() *string {
	return s.Port
}

func (s *UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList) GetVirtualServiceGroupId() *string {
	return s.VirtualServiceGroupId
}

func (s *UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList) GetVirtualServiceGroupName() *string {
	return s.VirtualServiceGroupName
}

func (s *UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList) SetPort(v string) *UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList {
	s.Port = &v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList) SetProtocol(v string) *UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList {
	s.Protocol = &v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList) SetVirtualServiceGroupId(v string) *UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList {
	s.VirtualServiceGroupId = &v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList) SetVirtualServiceGroupName(v string) *UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList {
	s.VirtualServiceGroupName = &v
	return s
}

func (s *UpdateGatewayLoadBalancerResponseBodyDataVirtualServiceList) Validate() error {
	return dara.Validate(s)
}
