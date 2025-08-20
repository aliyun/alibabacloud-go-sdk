// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceVpcEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceVpcEndpointRequest
	GetInstanceId() *string
	SetModuleName(v string) *GetInstanceVpcEndpointRequest
	GetModuleName() *string
}

type GetInstanceVpcEndpointRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// 	- `Registry`: the image repository.
	//
	// 	- `Chart`: a Helm chart.
	//
	// example:
	//
	// Chart
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s GetInstanceVpcEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceVpcEndpointRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceVpcEndpointRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetInstanceVpcEndpointRequest) SetInstanceId(v string) *GetInstanceVpcEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceVpcEndpointRequest) SetModuleName(v string) *GetInstanceVpcEndpointRequest {
	s.ModuleName = &v
	return s
}

func (s *GetInstanceVpcEndpointRequest) Validate() error {
	return dara.Validate(s)
}
