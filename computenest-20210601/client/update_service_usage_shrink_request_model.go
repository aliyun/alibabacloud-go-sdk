// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceUsageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateServiceUsageShrinkRequest
	GetClientToken() *string
	SetServiceId(v string) *UpdateServiceUsageShrinkRequest
	GetServiceId() *string
	SetUserInformationShrink(v string) *UpdateServiceUsageShrinkRequest
	GetUserInformationShrink() *string
}

type UpdateServiceUsageShrinkRequest struct {
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
	UserInformationShrink *string `json:"UserInformation,omitempty" xml:"UserInformation,omitempty"`
}

func (s UpdateServiceUsageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceUsageShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceUsageShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServiceUsageShrinkRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateServiceUsageShrinkRequest) GetUserInformationShrink() *string {
	return s.UserInformationShrink
}

func (s *UpdateServiceUsageShrinkRequest) SetClientToken(v string) *UpdateServiceUsageShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceUsageShrinkRequest) SetServiceId(v string) *UpdateServiceUsageShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateServiceUsageShrinkRequest) SetUserInformationShrink(v string) *UpdateServiceUsageShrinkRequest {
	s.UserInformationShrink = &v
	return s
}

func (s *UpdateServiceUsageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
