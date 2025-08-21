// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReadWritePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateReadWritePolicyRequest
	GetClientToken() *string
	SetBody(v string) *UpdateReadWritePolicyRequest
	GetBody() *string
}

type UpdateReadWritePolicyRequest struct {
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Body        *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateReadWritePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateReadWritePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateReadWritePolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateReadWritePolicyRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateReadWritePolicyRequest) SetClientToken(v string) *UpdateReadWritePolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateReadWritePolicyRequest) SetBody(v string) *UpdateReadWritePolicyRequest {
	s.Body = &v
	return s
}

func (s *UpdateReadWritePolicyRequest) Validate() error {
	return dara.Validate(s)
}
