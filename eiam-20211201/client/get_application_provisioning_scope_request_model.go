// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationProvisioningScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetApplicationProvisioningScopeRequest
	GetApplicationId() *string
	SetInstanceId(v string) *GetApplicationProvisioningScopeRequest
	GetInstanceId() *string
}

type GetApplicationProvisioningScopeRequest struct {
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

func (s GetApplicationProvisioningScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningScopeRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationProvisioningScopeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApplicationProvisioningScopeRequest) SetApplicationId(v string) *GetApplicationProvisioningScopeRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationProvisioningScopeRequest) SetInstanceId(v string) *GetApplicationProvisioningScopeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationProvisioningScopeRequest) Validate() error {
	return dara.Validate(s)
}
