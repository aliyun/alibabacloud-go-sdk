// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOrganizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOrganizationResponse
	GetStatusCode() *int32
	SetBody(v *GetOrganizationResponseBody) *GetOrganizationResponse
	GetBody() *GetOrganizationResponseBody
}

type GetOrganizationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrganizationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOrganizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOrganizationResponse) GetBody() *GetOrganizationResponseBody {
	return s.Body
}

func (s *GetOrganizationResponse) SetHeaders(v map[string]*string) *GetOrganizationResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationResponse) SetStatusCode(v int32) *GetOrganizationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrganizationResponse) SetBody(v *GetOrganizationResponseBody) *GetOrganizationResponse {
	s.Body = v
	return s
}

func (s *GetOrganizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
