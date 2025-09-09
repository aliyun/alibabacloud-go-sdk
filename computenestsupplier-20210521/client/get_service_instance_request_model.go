// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetServiceInstanceRequest
	GetRegionId() *string
	SetServiceInstanceId(v string) *GetServiceInstanceRequest
	GetServiceInstanceId() *string
}

type GetServiceInstanceRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s GetServiceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceInstanceRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *GetServiceInstanceRequest) SetRegionId(v string) *GetServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceInstanceRequest) SetServiceInstanceId(v string) *GetServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceRequest) Validate() error {
	return dara.Validate(s)
}
