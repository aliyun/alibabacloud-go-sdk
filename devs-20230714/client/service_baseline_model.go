// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceBaseline interface {
	dara.Model
	String() string
	GoString() string
	SetServiceInstance(v *ServiceInstance) *ServiceBaseline
	GetServiceInstance() *ServiceInstance
}

type ServiceBaseline struct {
	ServiceInstance *ServiceInstance `json:"serviceInstance,omitempty" xml:"serviceInstance,omitempty"`
}

func (s ServiceBaseline) String() string {
	return dara.Prettify(s)
}

func (s ServiceBaseline) GoString() string {
	return s.String()
}

func (s *ServiceBaseline) GetServiceInstance() *ServiceInstance {
	return s.ServiceInstance
}

func (s *ServiceBaseline) SetServiceInstance(v *ServiceInstance) *ServiceBaseline {
	s.ServiceInstance = v
	return s
}

func (s *ServiceBaseline) Validate() error {
	return dara.Validate(s)
}
