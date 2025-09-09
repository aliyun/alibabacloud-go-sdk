// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceRegistrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetServiceRegistrationRequest
	GetRegionId() *string
	SetRegistrationId(v string) *GetServiceRegistrationRequest
	GetRegistrationId() *string
}

type GetServiceRegistrationRequest struct {
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service registration ID.
	//
	// example:
	//
	// sr-1b4aabc1c9eb4109****
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
}

func (s GetServiceRegistrationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceRegistrationRequest) GoString() string {
	return s.String()
}

func (s *GetServiceRegistrationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceRegistrationRequest) GetRegistrationId() *string {
	return s.RegistrationId
}

func (s *GetServiceRegistrationRequest) SetRegionId(v string) *GetServiceRegistrationRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceRegistrationRequest) SetRegistrationId(v string) *GetServiceRegistrationRequest {
	s.RegistrationId = &v
	return s
}

func (s *GetServiceRegistrationRequest) Validate() error {
	return dara.Validate(s)
}
