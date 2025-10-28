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
	// The ID of the application. You can query the application ID by calling the ListApplication operation. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a166fbd-****-****-a286-781659d9f54c
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the cluster. You can query the cluster ID by calling the GetK8sCluster operation. For more information, see [GetK8sCluster](https://help.aliyun.com/document_detail/181437.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 712082c3-****-****-9217-a947b5cde6ee
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to disable listener configuration overriding.
	//
	// 	- true: disables listener configuration overriding.
	//
	// 	- false: enables listener configuration overriding.
	//
	// example:
	//
	// true
	DisableForceOverride *bool `json:"DisableForceOverride,omitempty" xml:"DisableForceOverride,omitempty"`
	// The frontend port. Valid values: 1 to 65535.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The scheduling algorithm for the SLB instance. If you do not specify this parameter, the default value rr is used. SLB supports the following scheduling algorithms: round-robin and weighted round-robin. Valid values:
	//
	// 	- wrr: weighted round-robin scheduling. Backend servers that have higher weights receive more requests than those that have lower weights.
	//
	// 	- rr: round-robin scheduling. Requests are sequentially distributed to backend servers.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The information about the ports. This parameter is required if you want to configure multi-port mappings or use a protocol other than TCP. You must set this parameter to a JSON array. Example: [{"targetPort":8080,"port":82,"loadBalancerProtocol":"TCP"},{"port":81,"certId":"1362469756373809_16c185d6fa2_1914500329_-xxxxxxx","targetPort":8181,"lo adBalancerProtocol":"HTTPS"}]
	//
	// 	- port: required. The frontend port. Valid values: 1 to 65535. Each port must be unique.
	//
	// 	- targetPort: required. The backend port. Valid values: 1 to 65535.
	//
	// 	- loadBalancerProtocol: required. Valid values: TCP and HTTPS. If the HTTP protocol is used, set this parameter to TCP.
	//
	// 	- certId: the ID of the certificate. This parameter is required if the HTTPS protocol is used. You can purchase an SLB instance in the SLB console.
	//
	// 	- Note: The ServicePortInfos parameter is specified to support multi-port mappings. If you want this parameter to take effect, make sure that you specify the AppId, ClusterId, Type, and SlbId parameters.
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
	// The protocol used by the SLB instance. Set the value to TCP.
	//
	// example:
	//
	// TCP
	SlbProtocol *string `json:"SlbProtocol,omitempty" xml:"SlbProtocol,omitempty"`
	// The specifications of the SLB instance.
	//
	// 	- slb.s1.small
	//
	// 	- slb.s2.small
	//
	// 	- slb.s2.medium
	//
	// 	- slb.s3.small
	//
	// 	- slb.s3.medium
	//
	// 	- slb.s3.large
	//
	// If you do not specify this parameter, the default value slb.s1.small is used.
	//
	// example:
	//
	// slb.s1.small
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The backend port, which is also the service port of the application. Valid values: 1 to 65535.
	//
	// example:
	//
	// 8082
	TargetPort *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	// The type of the SLB instance. Valid values:
	//
	// 	- Internet: an Internet-facing SLB instance
	//
	// 	- Intranet: an internal-facing SLB instance
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
