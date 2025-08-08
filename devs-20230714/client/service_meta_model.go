// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceMeta interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ServiceMeta
	GetName() *string
	SetType(v string) *ServiceMeta
	GetType() *string
}

type ServiceMeta struct {
	// example:
	//
	// my-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// fc3
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ServiceMeta) String() string {
	return dara.Prettify(s)
}

func (s ServiceMeta) GoString() string {
	return s.String()
}

func (s *ServiceMeta) GetName() *string {
	return s.Name
}

func (s *ServiceMeta) GetType() *string {
	return s.Type
}

func (s *ServiceMeta) SetName(v string) *ServiceMeta {
	s.Name = &v
	return s
}

func (s *ServiceMeta) SetType(v string) *ServiceMeta {
	s.Type = &v
	return s
}

func (s *ServiceMeta) Validate() error {
	return dara.Validate(s)
}
