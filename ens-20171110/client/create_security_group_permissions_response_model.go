// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityGroupPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSecurityGroupPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSecurityGroupPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *CreateSecurityGroupPermissionsResponseBody) *CreateSecurityGroupPermissionsResponse
	GetBody() *CreateSecurityGroupPermissionsResponseBody
}

type CreateSecurityGroupPermissionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecurityGroupPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecurityGroupPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityGroupPermissionsResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityGroupPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSecurityGroupPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSecurityGroupPermissionsResponse) GetBody() *CreateSecurityGroupPermissionsResponseBody {
	return s.Body
}

func (s *CreateSecurityGroupPermissionsResponse) SetHeaders(v map[string]*string) *CreateSecurityGroupPermissionsResponse {
	s.Headers = v
	return s
}

func (s *CreateSecurityGroupPermissionsResponse) SetStatusCode(v int32) *CreateSecurityGroupPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecurityGroupPermissionsResponse) SetBody(v *CreateSecurityGroupPermissionsResponseBody) *CreateSecurityGroupPermissionsResponse {
	s.Body = v
	return s
}

func (s *CreateSecurityGroupPermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
