// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityProviderSCIMRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateIdentityProviderSCIMRequestBody) *UpdateIdentityProviderSCIMRequest
	GetBody() *UpdateIdentityProviderSCIMRequestBody
}

type UpdateIdentityProviderSCIMRequest struct {
	// Submits the business parameters as described in this API. The body parameters of the public API are mapped to the backend wrapper by the gateway. Do not mix them with the backend direct connection body wrapper.
	//
	// This parameter is required.
	Body *UpdateIdentityProviderSCIMRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UpdateIdentityProviderSCIMRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderSCIMRequest) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderSCIMRequest) GetBody() *UpdateIdentityProviderSCIMRequestBody {
	return s.Body
}

func (s *UpdateIdentityProviderSCIMRequest) SetBody(v *UpdateIdentityProviderSCIMRequestBody) *UpdateIdentityProviderSCIMRequest {
	s.Body = v
	return s
}

func (s *UpdateIdentityProviderSCIMRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIdentityProviderSCIMRequestBody struct {
	// Specifies whether to enable synchronization. Do not set this parameter to true during CustomSSO initialization. You can only update this parameter by calling UpdateIdentityProviderSCIM later. Determine the result based on the operationStatus field.
	//
	// This parameter is required.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
}

func (s UpdateIdentityProviderSCIMRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderSCIMRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderSCIMRequestBody) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *UpdateIdentityProviderSCIMRequestBody) SetSyncEnabled(v bool) *UpdateIdentityProviderSCIMRequestBody {
	s.SyncEnabled = &v
	return s
}

func (s *UpdateIdentityProviderSCIMRequestBody) Validate() error {
	return dara.Validate(s)
}
