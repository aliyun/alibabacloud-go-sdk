// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindK8sSlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *BindK8sSlbRequest
	GetAppId() *string
	SetClusterId(v string) *BindK8sSlbRequest
	GetClusterId() *string
	SetPort(v string) *BindK8sSlbRequest
	GetPort() *string
	SetScheduler(v string) *BindK8sSlbRequest
	GetScheduler() *string
	SetServicePortInfos(v string) *BindK8sSlbRequest
	GetServicePortInfos() *string
	SetSlbId(v string) *BindK8sSlbRequest
	GetSlbId() *string
	SetSlbProtocol(v string) *BindK8sSlbRequest
	GetSlbProtocol() *string
	SetSpecification(v string) *BindK8sSlbRequest
	GetSpecification() *string
	SetTargetPort(v string) *BindK8sSlbRequest
	GetTargetPort() *string
	SetType(v string) *BindK8sSlbRequest
	GetType() *string
}

type BindK8sSlbRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a166fbd-****-****-a286-781659d9f54c
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// 712082c3-f554-****-****-a947b5cde6ee
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The frontend port. Valid values: 1 to 65535.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The scheduling algorithm for the SLB instance. If you do not specify this parameter, the default value rr is used. Valid values:
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
	// 	- port: The frontend port. Valid values: 1 to 65535. This parameter is required. Each port must be unique.
	//
	// 	- targetPort: The backend port. Valid values: 1 to 65535. This parameter is required.
	//
	// 	- loadBalancerProtocol: This parameter is required. Valid values: TCP and HTTPS. If the HTTP protocol is used, set this parameter to TCP.
	//
	// 	- certId: the ID of the certificate. This parameter is required if the HTTPS protocol is used. You can purchase an SLB instance in the SLB console.
	//
	// > The ServicePortInfos parameter is specified to support multi-port mappings. If you want this parameter to take effect, make sure that you have set the AppId, ClusterId, Type, and SlbId parameters.
	//
	// example:
	//
	// [{"targetPort":8080,"port":82,"loadBalancerProtocol":"TCP"},{"port":81,"certId":"136246975637380916c185d6fa21914500329_-988as","targetPort":8181,"lo adBalancerProtocol":"HTTPS"}]
	ServicePortInfos *string `json:"ServicePortInfos,omitempty" xml:"ServicePortInfos,omitempty"`
	// The ID of the SLB instance. If you leave this parameter empty, Enterprise Distributed Application Service (EDAS) automatically purchases an SLB instance.
	//
	// example:
	//
	// lb-2ze1quax9t****iz82bjt
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The protocol used by the SLB instance. Valid values: TCP, HTTP, and HTTPS.
	//
	// example:
	//
	// TCP
	SlbProtocol *string `json:"SlbProtocol,omitempty" xml:"SlbProtocol,omitempty"`
	// The instance type of the SLB instance. Valid values:
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
	// example:
	//
	// slb.s1.small
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The backend port, which is also the service port of the application. Valid values: 1 to 65535.
	//
	// example:
	//
	// 8080
	TargetPort *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	// The type of the SLB instance. Valid values:
	//
	// 	- internet: Internet-facing SLB instance
	//
	// 	- intranet: internal-facing SLB instance
	//
	// This parameter is required.
	//
	// example:
	//
	// internet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BindK8sSlbRequest) String() string {
	return dara.Prettify(s)
}

func (s BindK8sSlbRequest) GoString() string {
	return s.String()
}

func (s *BindK8sSlbRequest) GetAppId() *string {
	return s.AppId
}

func (s *BindK8sSlbRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *BindK8sSlbRequest) GetPort() *string {
	return s.Port
}

func (s *BindK8sSlbRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *BindK8sSlbRequest) GetServicePortInfos() *string {
	return s.ServicePortInfos
}

func (s *BindK8sSlbRequest) GetSlbId() *string {
	return s.SlbId
}

func (s *BindK8sSlbRequest) GetSlbProtocol() *string {
	return s.SlbProtocol
}

func (s *BindK8sSlbRequest) GetSpecification() *string {
	return s.Specification
}

func (s *BindK8sSlbRequest) GetTargetPort() *string {
	return s.TargetPort
}

func (s *BindK8sSlbRequest) GetType() *string {
	return s.Type
}

func (s *BindK8sSlbRequest) SetAppId(v string) *BindK8sSlbRequest {
	s.AppId = &v
	return s
}

func (s *BindK8sSlbRequest) SetClusterId(v string) *BindK8sSlbRequest {
	s.ClusterId = &v
	return s
}

func (s *BindK8sSlbRequest) SetPort(v string) *BindK8sSlbRequest {
	s.Port = &v
	return s
}

func (s *BindK8sSlbRequest) SetScheduler(v string) *BindK8sSlbRequest {
	s.Scheduler = &v
	return s
}

func (s *BindK8sSlbRequest) SetServicePortInfos(v string) *BindK8sSlbRequest {
	s.ServicePortInfos = &v
	return s
}

func (s *BindK8sSlbRequest) SetSlbId(v string) *BindK8sSlbRequest {
	s.SlbId = &v
	return s
}

func (s *BindK8sSlbRequest) SetSlbProtocol(v string) *BindK8sSlbRequest {
	s.SlbProtocol = &v
	return s
}

func (s *BindK8sSlbRequest) SetSpecification(v string) *BindK8sSlbRequest {
	s.Specification = &v
	return s
}

func (s *BindK8sSlbRequest) SetTargetPort(v string) *BindK8sSlbRequest {
	s.TargetPort = &v
	return s
}

func (s *BindK8sSlbRequest) SetType(v string) *BindK8sSlbRequest {
	s.Type = &v
	return s
}

func (s *BindK8sSlbRequest) Validate() error {
	return dara.Validate(s)
}
