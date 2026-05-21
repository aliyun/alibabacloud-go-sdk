// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstancePortRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPort(v int32) *UpdateInstancePortRequest
	GetPort() *int32
}

type UpdateInstancePortRequest struct {
	// example:
	//
	// 443
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
}

func (s UpdateInstancePortRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstancePortRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstancePortRequest) GetPort() *int32 {
	return s.Port
}

func (s *UpdateInstancePortRequest) SetPort(v int32) *UpdateInstancePortRequest {
	s.Port = &v
	return s
}

func (s *UpdateInstancePortRequest) Validate() error {
	return dara.Validate(s)
}
