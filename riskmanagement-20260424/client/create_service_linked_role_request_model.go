// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateServiceLinkedRoleRequest
	GetRegionId() *string
	SetSdkRequest(v *CreateServiceLinkedRoleRequestSdkRequest) *CreateServiceLinkedRoleRequest
	GetSdkRequest() *CreateServiceLinkedRoleRequestSdkRequest
}

type CreateServiceLinkedRoleRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId   *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SdkRequest *CreateServiceLinkedRoleRequestSdkRequest `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty" type:"Struct"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceLinkedRoleRequest) GetSdkRequest() *CreateServiceLinkedRoleRequestSdkRequest {
	return s.SdkRequest
}

func (s *CreateServiceLinkedRoleRequest) SetRegionId(v string) *CreateServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetSdkRequest(v *CreateServiceLinkedRoleRequestSdkRequest) *CreateServiceLinkedRoleRequest {
	s.SdkRequest = v
	return s
}

func (s *CreateServiceLinkedRoleRequest) Validate() error {
	if s.SdkRequest != nil {
		if err := s.SdkRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServiceLinkedRoleRequestSdkRequest struct {
	// example:
	//
	// AliyunServiceRoleForWebsiteBuildPublish
	ServiceLinkedRole *string `json:"ServiceLinkedRole,omitempty" xml:"ServiceLinkedRole,omitempty"`
}

func (s CreateServiceLinkedRoleRequestSdkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleRequestSdkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequestSdkRequest) GetServiceLinkedRole() *string {
	return s.ServiceLinkedRole
}

func (s *CreateServiceLinkedRoleRequestSdkRequest) SetServiceLinkedRole(v string) *CreateServiceLinkedRoleRequestSdkRequest {
	s.ServiceLinkedRole = &v
	return s
}

func (s *CreateServiceLinkedRoleRequestSdkRequest) Validate() error {
	return dara.Validate(s)
}
