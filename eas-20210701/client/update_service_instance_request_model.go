// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsolate(v bool) *UpdateServiceInstanceRequest
	GetIsolate() *bool
}

type UpdateServiceInstanceRequest struct {
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

func (s *UpdateServiceInstanceRequest) GetIsolate() *bool {
	return s.Isolate
}

func (s *UpdateServiceInstanceRequest) SetIsolate(v bool) *UpdateServiceInstanceRequest {
	s.Isolate = &v
	return s
}

func (s *UpdateServiceInstanceRequest) Validate() error {
	return dara.Validate(s)
}
