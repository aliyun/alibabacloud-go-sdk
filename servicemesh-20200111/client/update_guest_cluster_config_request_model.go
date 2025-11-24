// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGuestClusterConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGuestClusterId(v string) *UpdateGuestClusterConfigRequest
	GetGuestClusterId() *string
	SetSMCEnabled(v bool) *UpdateGuestClusterConfigRequest
	GetSMCEnabled() *bool
	SetServiceMeshId(v string) *UpdateGuestClusterConfigRequest
	GetServiceMeshId() *string
}

type UpdateGuestClusterConfigRequest struct {
	// example:
	//
	// c42186268a27f475c975e5667bb66****
	GuestClusterId *string `json:"GuestClusterId,omitempty" xml:"GuestClusterId,omitempty"`
	// example:
	//
	// false
	SMCEnabled *bool `json:"SMCEnabled,omitempty" xml:"SMCEnabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s UpdateGuestClusterConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGuestClusterConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateGuestClusterConfigRequest) GetGuestClusterId() *string {
	return s.GuestClusterId
}

func (s *UpdateGuestClusterConfigRequest) GetSMCEnabled() *bool {
	return s.SMCEnabled
}

func (s *UpdateGuestClusterConfigRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateGuestClusterConfigRequest) SetGuestClusterId(v string) *UpdateGuestClusterConfigRequest {
	s.GuestClusterId = &v
	return s
}

func (s *UpdateGuestClusterConfigRequest) SetSMCEnabled(v bool) *UpdateGuestClusterConfigRequest {
	s.SMCEnabled = &v
	return s
}

func (s *UpdateGuestClusterConfigRequest) SetServiceMeshId(v string) *UpdateGuestClusterConfigRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateGuestClusterConfigRequest) Validate() error {
	return dara.Validate(s)
}
