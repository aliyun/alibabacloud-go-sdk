// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *GetServiceEndpointRequest
	GetEndpointId() *string
	SetInstanceId(v string) *GetServiceEndpointRequest
	GetInstanceId() *string
}

type GetServiceEndpointRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mep-test0001
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetServiceEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointRequest) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *GetServiceEndpointRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetServiceEndpointRequest) SetEndpointId(v string) *GetServiceEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *GetServiceEndpointRequest) SetInstanceId(v string) *GetServiceEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *GetServiceEndpointRequest) Validate() error {
	return dara.Validate(s)
}
