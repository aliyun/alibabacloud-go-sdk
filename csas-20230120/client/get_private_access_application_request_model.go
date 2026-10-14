// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrivateAccessApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetPrivateAccessApplicationRequest
	GetApplicationId() *string
}

type GetPrivateAccessApplicationRequest struct {
	// The ID of the private access application. You can obtain the value from the following operations:
	//
	// - [ListPrivateAccessApplications](~~ListPrivateAccessApplications~~): Lists private access applications.
	//
	// - [CreatePrivateAccessApplication](~~CreatePrivateAccessApplication~~): Creates a private access application.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s GetPrivateAccessApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPrivateAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetPrivateAccessApplicationRequest) SetApplicationId(v string) *GetPrivateAccessApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetPrivateAccessApplicationRequest) Validate() error {
	return dara.Validate(s)
}
