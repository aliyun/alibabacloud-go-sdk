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
	// The frontend port. The value must be an integer from 1 to 65,535.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The scheduling algorithm. If you do not specify this parameter, \\`rr\\` is used. Valid values:
	//
	// - wrr: weighted round-robin. Backend servers with higher weights receive more requests.
	//
	// - rr: round-robin. Requests are distributed to backend servers in sequence.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The information about the service ports. Use this parameter to configure multiple listeners or use protocols other than TCP.
	//
	// This parameter must be a JSON array. Example:
	//
	// [{"targetPort":8080,"port":82,"loadBalancerProtocol":"TCP"},{"port":81,"certId":"1362469756373809_16c185d6fa2_1914500329_-xxxxxxx","targetPort":8181,"loadBalancerProtocol":"HTTPS"}]
	//
	// - port: Required. The frontend port. The value must be an integer from 1 to 65,535. Each port number must be unique.
	//
	// - targetPort: Required. The backend port. The value must be an integer from 1 to 65,535.
	//
	// - loadBalancerProtocol: Required. The frontend protocol. Valid values: TCP and HTTPS. For HTTP, use TCP.
	//
	// - certId: Required if you use the HTTPS protocol. You can purchase a certificate in the SLB console.
	//
	// > This parameter is used to configure multiple listeners. You must use it with the appId, clusterId, type, and slbId parameters.
	//
	// example:
	//
	// [{"targetPort":8080,"port":82,"loadBalancerProtocol":"TCP"},{"port":81,"certId":"136246975637380916c185d6fa21914500329_-988as","targetPort":8181,"lo adBalancerProtocol":"HTTPS"}]
	ServicePortInfos *string `json:"ServicePortInfos,omitempty" xml:"ServicePortInfos,omitempty"`
	// The ID of the SLB instance. If you do not specify this parameter, EDAS automatically purchases a new SLB instance.
	//
	// example:
	//
	// lb-2ze1quax9t****iz82bjt
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The frontend protocol for the SLB instance. Valid values: TCP, HTTP, and HTTPS.
	//
	// example:
	//
	// TCP
	SlbProtocol *string `json:"SlbProtocol,omitempty" xml:"SlbProtocol,omitempty"`
	// The specification of the SLB instance.
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
	// example:
	//
	// slb.s1.small
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The backend port. This port is also the service port of the application. The value must be an integer from 1 to 65,535.
	//
	// example:
	//
	// 8080
	TargetPort *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	// The type of the SLB instance.
	//
	// - internet: an internet-facing SLB instance.
	//
	// - intranet: an internal-facing SLB instance.
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
