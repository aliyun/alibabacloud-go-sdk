// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccessApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateAccessApplicationRequestBody) *UpdateAccessApplicationRequest
	GetBody() *UpdateAccessApplicationRequestBody
}

type UpdateAccessApplicationRequest struct {
	// The request body for updating the access application.
	//
	// This parameter is required.
	Body *UpdateAccessApplicationRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UpdateAccessApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccessApplicationRequest) GetBody() *UpdateAccessApplicationRequestBody {
	return s.Body
}

func (s *UpdateAccessApplicationRequest) SetBody(v *UpdateAccessApplicationRequestBody) *UpdateAccessApplicationRequest {
	s.Body = v
	return s
}

func (s *UpdateAccessApplicationRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAccessApplicationRequestBody struct {
	// The name of the access application.
	//
	// example:
	//
	// my-access-application
	ApplicationName *string `json:"applicationName,omitempty" xml:"applicationName,omitempty"`
	// The list of redirect URIs.
	RedirectUris []*string `json:"redirectUris,omitempty" xml:"redirectUris,omitempty" type:"Repeated"`
}

func (s UpdateAccessApplicationRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessApplicationRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateAccessApplicationRequestBody) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *UpdateAccessApplicationRequestBody) GetRedirectUris() []*string {
	return s.RedirectUris
}

func (s *UpdateAccessApplicationRequestBody) SetApplicationName(v string) *UpdateAccessApplicationRequestBody {
	s.ApplicationName = &v
	return s
}

func (s *UpdateAccessApplicationRequestBody) SetRedirectUris(v []*string) *UpdateAccessApplicationRequestBody {
	s.RedirectUris = v
	return s
}

func (s *UpdateAccessApplicationRequestBody) Validate() error {
	return dara.Validate(s)
}
