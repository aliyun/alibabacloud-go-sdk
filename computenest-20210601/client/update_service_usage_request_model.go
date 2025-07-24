// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateServiceUsageRequest
	GetClientToken() *string
	SetServiceId(v string) *UpdateServiceUsageRequest
	GetServiceId() *string
	SetUserInformation(v map[string]*string) *UpdateServiceUsageRequest
	GetUserInformation() map[string]*string
}

type UpdateServiceUsageRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// AAAAAYChudnQUoBH+mGWFpb6oP0=
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-39f4f251e94843xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The information about the applicant.
	UserInformation map[string]*string `json:"UserInformation,omitempty" xml:"UserInformation,omitempty"`
}

func (s UpdateServiceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceUsageRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceUsageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServiceUsageRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateServiceUsageRequest) GetUserInformation() map[string]*string {
	return s.UserInformation
}

func (s *UpdateServiceUsageRequest) SetClientToken(v string) *UpdateServiceUsageRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceUsageRequest) SetServiceId(v string) *UpdateServiceUsageRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateServiceUsageRequest) SetUserInformation(v map[string]*string) *UpdateServiceUsageRequest {
	s.UserInformation = v
	return s
}

func (s *UpdateServiceUsageRequest) Validate() error {
	return dara.Validate(s)
}
