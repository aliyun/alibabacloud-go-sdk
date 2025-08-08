// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProjectStatus interface {
	dara.Model
	String() string
	GoString() string
	SetServices(v []*ServiceMeta) *ProjectStatus
	GetServices() []*ServiceMeta
}

type ProjectStatus struct {
	Services []*ServiceMeta `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s ProjectStatus) String() string {
	return dara.Prettify(s)
}

func (s ProjectStatus) GoString() string {
	return s.String()
}

func (s *ProjectStatus) GetServices() []*ServiceMeta {
	return s.Services
}

func (s *ProjectStatus) SetServices(v []*ServiceMeta) *ProjectStatus {
	s.Services = v
	return s
}

func (s *ProjectStatus) Validate() error {
	return dara.Validate(s)
}
