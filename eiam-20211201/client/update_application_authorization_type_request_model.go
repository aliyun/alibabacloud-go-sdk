// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationAuthorizationTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateApplicationAuthorizationTypeRequest
	GetApplicationId() *string
	SetAuthorizationType(v string) *UpdateApplicationAuthorizationTypeRequest
	GetAuthorizationType() *string
	SetInstanceId(v string) *UpdateApplicationAuthorizationTypeRequest
	GetInstanceId() *string
}

type UpdateApplicationAuthorizationTypeRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The access authorization type of the application. Valid values:
	//
	// - authorize_required: Explicit authorization is required for access.
	//
	// - default_all: All members have access permissions by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// authorize_required
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateApplicationAuthorizationTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationAuthorizationTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationAuthorizationTypeRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationAuthorizationTypeRequest) GetAuthorizationType() *string {
	return s.AuthorizationType
}

func (s *UpdateApplicationAuthorizationTypeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateApplicationAuthorizationTypeRequest) SetApplicationId(v string) *UpdateApplicationAuthorizationTypeRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationAuthorizationTypeRequest) SetAuthorizationType(v string) *UpdateApplicationAuthorizationTypeRequest {
	s.AuthorizationType = &v
	return s
}

func (s *UpdateApplicationAuthorizationTypeRequest) SetInstanceId(v string) *UpdateApplicationAuthorizationTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateApplicationAuthorizationTypeRequest) Validate() error {
	return dara.Validate(s)
}
