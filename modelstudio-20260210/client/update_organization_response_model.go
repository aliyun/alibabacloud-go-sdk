// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOrganizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOrganizationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOrganizationResponseBody) *UpdateOrganizationResponse
	GetBody() *UpdateOrganizationResponseBody
}

type UpdateOrganizationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOrganizationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOrganizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOrganizationResponse) GetBody() *UpdateOrganizationResponseBody {
	return s.Body
}

func (s *UpdateOrganizationResponse) SetHeaders(v map[string]*string) *UpdateOrganizationResponse {
	s.Headers = v
	return s
}

func (s *UpdateOrganizationResponse) SetStatusCode(v int32) *UpdateOrganizationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOrganizationResponse) SetBody(v *UpdateOrganizationResponseBody) *UpdateOrganizationResponse {
	s.Body = v
	return s
}

func (s *UpdateOrganizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
