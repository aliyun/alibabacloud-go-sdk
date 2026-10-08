// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sSlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateK8sSlbRequest
	GetAppId() *string
	SetClusterId(v string) *UpdateK8sSlbRequest
	GetClusterId() *string
	SetDisableForceOverride(v bool) *UpdateK8sSlbRequest
	GetDisableForceOverride() *bool
	SetPort(v string) *UpdateK8sSlbRequest
	GetPort() *string
	SetScheduler(v string) *UpdateK8sSlbRequest
	GetScheduler() *string
	SetServicePortInfos(v string) *UpdateK8sSlbRequest
	GetServicePortInfos() *string
	SetSlbName(v string) *UpdateK8sSlbRequest
	GetSlbName() *string
	SetSlbProtocol(v string) *UpdateK8sSlbRequest
	GetSlbProtocol() *string
	SetSpecification(v string) *UpdateK8sSlbRequest
	GetSpecification() *string
	SetTargetPort(v string) *UpdateK8sSlbRequest
	GetTargetPort() *string
	SetType(v string) *UpdateK8sSlbRequest
	GetType() *string
}

type UpdateK8sSlbRequest struct {
	// The ID of the application. Call [ListApplication](https://help.aliyun.com/document_detail/149390.html) to get this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a166fbd-****-****-a286-781659d9f54c
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the cluster. Call [GetK8sCluster](https://help.aliyun.com/document_detail/181437.html) to get this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 712082c3-****-****-9217-a947b5cde6ee
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to disable overwriting the SLB listener configuration.
	//
	// - true: Disables overwriting.
	//
	// - false: Allows overwriting.
	//
	// example:
	//
	// true
	DisableForceOverride *bool `json:"DisableForceOverride,omitempty" xml:"DisableForceOverride,omitempty"`
	// The frontend port. The value ranges from 1 to 65535.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The scheduling algorithm of the SLB instance. If you do not set this parameter, rr is used. The supported algorithms are round-robin (rr) and weighted round-robin (wrr).
	//
	// - Weighted round-robin (wrr): Backend servers with higher weights receive more requests.
	//
	// - Round-robin (rr): Requests are distributed to backend servers in sequence.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// This parameter is used for scenarios that involve multiple ports or protocols other than TCP. The value must be a JSON array. For example:
	//
	// [{"targetPort":8080,"port":82,"loadBalancerProtocol":"TCP"},{"port":81,"certId":"1362469756373809_16c185d6fa2_1914500329_-xxxxxxx","targetPort":8181,"loadBalancerProtocol":"HTTPS"}]
	//
	// - port: Required. The frontend port. The value ranges from 1 to 65535. Each port number must be unique.
	//
	// - targetPort: Required. The backend port. The value ranges from 1 to 65535.
	//
	// - loadBalancerProtocol: Required. Only TCP and HTTPS are supported. For HTTP listeners, set this parameter to TCP.
	//
	// - certId: This parameter is required for HTTPS listeners. It specifies the ID of a certificate that you can purchase in the SLB console.
	//
	// - Note: This parameter is used to support multiple ports and must be used with the appId, clusterId, type, and slbId parameters.
	//
	// example:
	//
	// {"targetPort":8080,"port":82,"loadBalancerProtocol":"TCP"},{"port":81,"certId":"136246975637380916c185d6fa21914500329_-xxxxxxx","targetPort":8181,"lo adBalancerProtocol":"HTTPS"}
	ServicePortInfos *string `json:"ServicePortInfos,omitempty" xml:"ServicePortInfos,omitempty"`
	// The name of the SLB instance.
	//
	// example:
	//
	// SLB_doctest
	SlbName *string `json:"SlbName,omitempty" xml:"SlbName,omitempty"`
	// The protocol of the SLB instance. Currently, only TCP is supported.
	//
	// example:
	//
	// TCP
	SlbProtocol *string `json:"SlbProtocol,omitempty" xml:"SlbProtocol,omitempty"`
	// The specification of the SLB instance. The following specifications are supported:
	//
	// - slb.s1.small
	//
	// - slb.s2.small
	//
	// - slb.s2.medium
	//
	// - slb.s3.small
	//
	// - slb.s3.medium
	//
	// - slb.s3.large
	//
	// If you do not set this parameter, the default value is slb.s1.small.
	//
	// example:
	//
	// slb.s1.small
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The backend port, which is the service port of the application. The value ranges from 1 to 65535.
	//
	// example:
	//
	// 8082
	TargetPort *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	// The type of the SLB instance.
	//
	// - Internet: An Internet-facing instance.
	//
	// - Intranet: An internal-facing instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// Internet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateK8sSlbRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sSlbRequest) GoString() string {
	return s.String()
}

func (s *UpdateK8sSlbRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateK8sSlbRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateK8sSlbRequest) GetDisableForceOverride() *bool {
	return s.DisableForceOverride
}

func (s *UpdateK8sSlbRequest) GetPort() *string {
	return s.Port
}

func (s *UpdateK8sSlbRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *UpdateK8sSlbRequest) GetServicePortInfos() *string {
	return s.ServicePortInfos
}

func (s *UpdateK8sSlbRequest) GetSlbName() *string {
	return s.SlbName
}

func (s *UpdateK8sSlbRequest) GetSlbProtocol() *string {
	return s.SlbProtocol
}

func (s *UpdateK8sSlbRequest) GetSpecification() *string {
	return s.Specification
}

func (s *UpdateK8sSlbRequest) GetTargetPort() *string {
	return s.TargetPort
}

func (s *UpdateK8sSlbRequest) GetType() *string {
	return s.Type
}

func (s *UpdateK8sSlbRequest) SetAppId(v string) *UpdateK8sSlbRequest {
	s.AppId = &v
	return s
}

func (s *UpdateK8sSlbRequest) SetClusterId(v string) *UpdateK8sSlbRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateK8sSlbRequest) SetDisableForceOverride(v bool) *UpdateK8sSlbRequest {
	s.DisableForceOverride = &v
	return s
}

func (s *UpdateK8sSlbRequest) SetPort(v string) *UpdateK8sSlbRequest {
	s.Port = &v
	return s
}

func (s *UpdateK8sSlbRequest) SetScheduler(v string) *UpdateK8sSlbRequest {
	s.Scheduler = &v
	return s
}

func (s *UpdateK8sSlbRequest) SetServicePortInfos(v string) *UpdateK8sSlbRequest {
	s.ServicePortInfos = &v
	return s
}

func (s *UpdateK8sSlbRequest) SetSlbName(v string) *UpdateK8sSlbRequest {
	s.SlbName = &v
	return s
}

func (s *UpdateK8sSlbRequest) SetSlbProtocol(v string) *UpdateK8sSlbRequest {
	s.SlbProtocol = &v
	return s
}

func (s *UpdateK8sSlbRequest) SetSpecification(v string) *UpdateK8sSlbRequest {
	s.Specification = &v
	return s
}

func (s *UpdateK8sSlbRequest) SetTargetPort(v string) *UpdateK8sSlbRequest {
	s.TargetPort = &v
	return s
}

func (s *UpdateK8sSlbRequest) SetType(v string) *UpdateK8sSlbRequest {
	s.Type = &v
	return s
}

func (s *UpdateK8sSlbRequest) Validate() error {
	return dara.Validate(s)
}
