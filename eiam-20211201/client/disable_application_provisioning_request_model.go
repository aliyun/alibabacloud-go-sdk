// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationProvisioningRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DisableApplicationProvisioningRequest
	GetApplicationId() *string
	SetInstanceId(v string) *DisableApplicationProvisioningRequest
	GetInstanceId() *string
}

type DisableApplicationProvisioningRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableApplicationProvisioningRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationProvisioningRequest) GoString() string {
	return s.String()
}

func (s *DisableApplicationProvisioningRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DisableApplicationProvisioningRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableApplicationProvisioningRequest) SetApplicationId(v string) *DisableApplicationProvisioningRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisableApplicationProvisioningRequest) SetInstanceId(v string) *DisableApplicationProvisioningRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableApplicationProvisioningRequest) Validate() error {
	return dara.Validate(s)
}
