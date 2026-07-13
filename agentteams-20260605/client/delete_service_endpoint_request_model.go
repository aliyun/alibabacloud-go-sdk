// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *DeleteServiceEndpointRequest
	GetEndpointId() *string
	SetInstanceId(v string) *DeleteServiceEndpointRequest
	GetInstanceId() *string
}

type DeleteServiceEndpointRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ep-xxx
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteServiceEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceEndpointRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DeleteServiceEndpointRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteServiceEndpointRequest) SetEndpointId(v string) *DeleteServiceEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *DeleteServiceEndpointRequest) SetInstanceId(v string) *DeleteServiceEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteServiceEndpointRequest) Validate() error {
	return dara.Validate(s)
}
