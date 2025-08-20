// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointType(v string) *GetInstanceEndpointRequest
	GetEndpointType() *string
	SetInstanceId(v string) *GetInstanceEndpointRequest
	GetInstanceId() *string
	SetModuleName(v string) *GetInstanceEndpointRequest
	GetModuleName() *string
}

type GetInstanceEndpointRequest struct {
	// The type of the endpoint. Set the value to Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// internet
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// 	- `Registry`: the image repository.
	//
	// 	- `Chart`: a Helm chart.
	//
	// example:
	//
	// Registry
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s GetInstanceEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceEndpointRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceEndpointRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *GetInstanceEndpointRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceEndpointRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetInstanceEndpointRequest) SetEndpointType(v string) *GetInstanceEndpointRequest {
	s.EndpointType = &v
	return s
}

func (s *GetInstanceEndpointRequest) SetInstanceId(v string) *GetInstanceEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceEndpointRequest) SetModuleName(v string) *GetInstanceEndpointRequest {
	s.ModuleName = &v
	return s
}

func (s *GetInstanceEndpointRequest) Validate() error {
	return dara.Validate(s)
}
