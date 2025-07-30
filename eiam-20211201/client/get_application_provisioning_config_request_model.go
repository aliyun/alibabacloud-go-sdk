// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationProvisioningConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetApplicationProvisioningConfigRequest
	GetApplicationId() *string
	SetInstanceId(v string) *GetApplicationProvisioningConfigRequest
	GetInstanceId() *string
}

type GetApplicationProvisioningConfigRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetApplicationProvisioningConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningConfigRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningConfigRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationProvisioningConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApplicationProvisioningConfigRequest) SetApplicationId(v string) *GetApplicationProvisioningConfigRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationProvisioningConfigRequest) SetInstanceId(v string) *GetApplicationProvisioningConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationProvisioningConfigRequest) Validate() error {
	return dara.Validate(s)
}
