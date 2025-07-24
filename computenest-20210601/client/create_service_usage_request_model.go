// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateServiceUsageRequest
	GetClientToken() *string
	SetServiceId(v string) *CreateServiceUsageRequest
	GetServiceId() *string
	SetUserInformation(v map[string]*string) *CreateServiceUsageRequest
	GetUserInformation() map[string]*string
}

type CreateServiceUsageRequest struct {
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
	UserInformation map[string]*string `json:"UserInformation,omitempty" xml:"UserInformation,omitempty"`
}

func (s CreateServiceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceUsageRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceUsageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServiceUsageRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateServiceUsageRequest) GetUserInformation() map[string]*string {
	return s.UserInformation
}

func (s *CreateServiceUsageRequest) SetClientToken(v string) *CreateServiceUsageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceUsageRequest) SetServiceId(v string) *CreateServiceUsageRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceUsageRequest) SetUserInformation(v map[string]*string) *CreateServiceUsageRequest {
	s.UserInformation = v
	return s
}

func (s *CreateServiceUsageRequest) Validate() error {
	return dara.Validate(s)
}
