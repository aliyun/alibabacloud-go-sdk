// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *RestartInstanceRequest
	GetClusterId() *string
	SetComponents(v string) *RestartInstanceRequest
	GetComponents() *string
}

type RestartInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// THRIFT
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
}

func (s RestartInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *RestartInstanceRequest) GetComponents() *string {
	return s.Components
}

func (s *RestartInstanceRequest) SetClusterId(v string) *RestartInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *RestartInstanceRequest) SetComponents(v string) *RestartInstanceRequest {
	s.Components = &v
	return s
}

func (s *RestartInstanceRequest) Validate() error {
	return dara.Validate(s)
}
