// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHibernate(v bool) *UpdateServiceInstanceRequest
	GetHibernate() *bool
	SetIsolate(v bool) *UpdateServiceInstanceRequest
	GetIsolate() *bool
}

type UpdateServiceInstanceRequest struct {
	Hibernate *bool `json:"Hibernate,omitempty" xml:"Hibernate,omitempty"`
	// Specifies whether to isolate the service instance. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Isolate *bool `json:"Isolate,omitempty" xml:"Isolate,omitempty"`
}

func (s UpdateServiceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceRequest) GetHibernate() *bool {
	return s.Hibernate
}

func (s *UpdateServiceInstanceRequest) GetIsolate() *bool {
	return s.Isolate
}

func (s *UpdateServiceInstanceRequest) SetHibernate(v bool) *UpdateServiceInstanceRequest {
	s.Hibernate = &v
	return s
}

func (s *UpdateServiceInstanceRequest) SetIsolate(v bool) *UpdateServiceInstanceRequest {
	s.Isolate = &v
	return s
}

func (s *UpdateServiceInstanceRequest) Validate() error {
	return dara.Validate(s)
}
