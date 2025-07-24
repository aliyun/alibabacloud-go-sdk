// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceUsageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateServiceUsageShrinkRequest
	GetClientToken() *string
	SetServiceId(v string) *CreateServiceUsageShrinkRequest
	GetServiceId() *string
	SetUserInformationShrink(v string) *CreateServiceUsageShrinkRequest
	GetUserInformationShrink() *string
}

type CreateServiceUsageShrinkRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-d6fc5f949a9246xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The application information.
	UserInformationShrink *string `json:"UserInformation,omitempty" xml:"UserInformation,omitempty"`
}

func (s CreateServiceUsageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceUsageShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceUsageShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServiceUsageShrinkRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateServiceUsageShrinkRequest) GetUserInformationShrink() *string {
	return s.UserInformationShrink
}

func (s *CreateServiceUsageShrinkRequest) SetClientToken(v string) *CreateServiceUsageShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceUsageShrinkRequest) SetServiceId(v string) *CreateServiceUsageShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceUsageShrinkRequest) SetUserInformationShrink(v string) *CreateServiceUsageShrinkRequest {
	s.UserInformationShrink = &v
	return s
}

func (s *CreateServiceUsageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
