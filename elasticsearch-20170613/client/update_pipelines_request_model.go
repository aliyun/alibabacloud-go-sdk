// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdatePipelinesRequest
	GetBody() *string
	SetClientToken(v string) *UpdatePipelinesRequest
	GetClientToken() *string
	SetTrigger(v bool) *UpdatePipelinesRequest
	GetTrigger() *bool
}

type UpdatePipelinesRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// Specifies whether to deploy the pipeline immediately.
	//
	// example:
	//
	// false
	Trigger *bool `json:"trigger,omitempty" xml:"trigger,omitempty"`
}

func (s UpdatePipelinesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelinesRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelinesRequest) GetBody() *string {
	return s.Body
}

func (s *UpdatePipelinesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdatePipelinesRequest) GetTrigger() *bool {
	return s.Trigger
}

func (s *UpdatePipelinesRequest) SetBody(v string) *UpdatePipelinesRequest {
	s.Body = &v
	return s
}

func (s *UpdatePipelinesRequest) SetClientToken(v string) *UpdatePipelinesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePipelinesRequest) SetTrigger(v bool) *UpdatePipelinesRequest {
	s.Trigger = &v
	return s
}

func (s *UpdatePipelinesRequest) Validate() error {
	return dara.Validate(s)
}
