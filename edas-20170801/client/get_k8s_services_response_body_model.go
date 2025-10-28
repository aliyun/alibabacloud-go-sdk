// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetK8sServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetK8sServicesResponseBody
	GetCode() *int32
	SetMessage(v string) *GetK8sServicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetK8sServicesResponseBody
	GetRequestId() *string
	SetServices(v []*GetK8sServicesResponseBodyServices) *GetK8sServicesResponseBody
	GetServices() []*GetK8sServicesResponseBodyServices
}

type GetK8sServicesResponseBody struct {
	// The ID of the change process.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4823-bhjf-23u4-eiufh
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of services in the Kubernetes cluster.
	Services []*GetK8sServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s GetK8sServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetK8sServicesResponseBody) GoString() string {
	return s.String()
}

func (s *GetK8sServicesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetK8sServicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetK8sServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetK8sServicesResponseBody) GetServices() []*GetK8sServicesResponseBodyServices {
	return s.Services
}

func (s *GetK8sServicesResponseBody) SetCode(v int32) *GetK8sServicesResponseBody {
	s.Code = &v
	return s
}

func (s *GetK8sServicesResponseBody) SetMessage(v string) *GetK8sServicesResponseBody {
	s.Message = &v
	return s
}

func (s *GetK8sServicesResponseBody) SetRequestId(v string) *GetK8sServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetK8sServicesResponseBody) SetServices(v []*GetK8sServicesResponseBodyServices) *GetK8sServicesResponseBody {
	s.Services = v
	return s
}

func (s *GetK8sServicesResponseBody) Validate() error {
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetK8sServicesResponseBodyServices struct {
	// The IP address of the service in the Kubernetes cluster.
	//
	// example:
	//
	// 104.23.xx.xx
	ClusterIP *string `json:"ClusterIP,omitempty" xml:"ClusterIP,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// service-http
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The mapping of service ports.
	ServicePorts []*GetK8sServicesResponseBodyServicesServicePorts `json:"ServicePorts,omitempty" xml:"ServicePorts,omitempty" type:"Repeated"`
	// The type of the service.
	//
	// example:
	//
	// ClusterIP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetK8sServicesResponseBodyServices) String() string {
	return dara.Prettify(s)
}

func (s GetK8sServicesResponseBodyServices) GoString() string {
	return s.String()
}

func (s *GetK8sServicesResponseBodyServices) GetClusterIP() *string {
	return s.ClusterIP
}

func (s *GetK8sServicesResponseBodyServices) GetName() *string {
	return s.Name
}

func (s *GetK8sServicesResponseBodyServices) GetServicePorts() []*GetK8sServicesResponseBodyServicesServicePorts {
	return s.ServicePorts
}

func (s *GetK8sServicesResponseBodyServices) GetType() *string {
	return s.Type
}

func (s *GetK8sServicesResponseBodyServices) SetClusterIP(v string) *GetK8sServicesResponseBodyServices {
	s.ClusterIP = &v
	return s
}

func (s *GetK8sServicesResponseBodyServices) SetName(v string) *GetK8sServicesResponseBodyServices {
	s.Name = &v
	return s
}

func (s *GetK8sServicesResponseBodyServices) SetServicePorts(v []*GetK8sServicesResponseBodyServicesServicePorts) *GetK8sServicesResponseBodyServices {
	s.ServicePorts = v
	return s
}

func (s *GetK8sServicesResponseBodyServices) SetType(v string) *GetK8sServicesResponseBodyServices {
	s.Type = &v
	return s
}

func (s *GetK8sServicesResponseBodyServices) Validate() error {
	if s.ServicePorts != nil {
		for _, item := range s.ServicePorts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetK8sServicesResponseBodyServicesServicePorts struct {
	// The port of the node.
	//
	// example:
	//
	// 0
	NodePort *int32 `json:"NodePort,omitempty" xml:"NodePort,omitempty"`
	// The frontend service port.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol of the service.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The backend container port.
	//
	// example:
	//
	// 8080
	TargetPort *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s GetK8sServicesResponseBodyServicesServicePorts) String() string {
	return dara.Prettify(s)
}

func (s GetK8sServicesResponseBodyServicesServicePorts) GoString() string {
	return s.String()
}

func (s *GetK8sServicesResponseBodyServicesServicePorts) GetNodePort() *int32 {
	return s.NodePort
}

func (s *GetK8sServicesResponseBodyServicesServicePorts) GetPort() *int32 {
	return s.Port
}

func (s *GetK8sServicesResponseBodyServicesServicePorts) GetProtocol() *string {
	return s.Protocol
}

func (s *GetK8sServicesResponseBodyServicesServicePorts) GetTargetPort() *string {
	return s.TargetPort
}

func (s *GetK8sServicesResponseBodyServicesServicePorts) SetNodePort(v int32) *GetK8sServicesResponseBodyServicesServicePorts {
	s.NodePort = &v
	return s
}

func (s *GetK8sServicesResponseBodyServicesServicePorts) SetPort(v int32) *GetK8sServicesResponseBodyServicesServicePorts {
	s.Port = &v
	return s
}

func (s *GetK8sServicesResponseBodyServicesServicePorts) SetProtocol(v string) *GetK8sServicesResponseBodyServicesServicePorts {
	s.Protocol = &v
	return s
}

func (s *GetK8sServicesResponseBodyServicesServicePorts) SetTargetPort(v string) *GetK8sServicesResponseBodyServicesServicePorts {
	s.TargetPort = &v
	return s
}

func (s *GetK8sServicesResponseBodyServicesServicePorts) Validate() error {
	return dara.Validate(s)
}
