// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogstashDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateLogstashDescriptionRequest
	GetDescription() *string
	SetClientToken(v string) *UpdateLogstashDescriptionRequest
	GetClientToken() *string
}

type UpdateLogstashDescriptionRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateLogstashDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogstashDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateLogstashDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLogstashDescriptionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateLogstashDescriptionRequest) SetDescription(v string) *UpdateLogstashDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateLogstashDescriptionRequest) SetClientToken(v string) *UpdateLogstashDescriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLogstashDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
