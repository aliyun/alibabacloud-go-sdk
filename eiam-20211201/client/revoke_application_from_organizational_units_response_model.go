// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeApplicationFromOrganizationalUnitsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeApplicationFromOrganizationalUnitsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeApplicationFromOrganizationalUnitsResponse
	GetStatusCode() *int32
	SetBody(v *RevokeApplicationFromOrganizationalUnitsResponseBody) *RevokeApplicationFromOrganizationalUnitsResponse
	GetBody() *RevokeApplicationFromOrganizationalUnitsResponseBody
}

type RevokeApplicationFromOrganizationalUnitsResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeApplicationFromOrganizationalUnitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeApplicationFromOrganizationalUnitsResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeApplicationFromOrganizationalUnitsResponse) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromOrganizationalUnitsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeApplicationFromOrganizationalUnitsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeApplicationFromOrganizationalUnitsResponse) GetBody() *RevokeApplicationFromOrganizationalUnitsResponseBody {
	return s.Body
}

func (s *RevokeApplicationFromOrganizationalUnitsResponse) SetHeaders(v map[string]*string) *RevokeApplicationFromOrganizationalUnitsResponse {
	s.Headers = v
	return s
}

func (s *RevokeApplicationFromOrganizationalUnitsResponse) SetStatusCode(v int32) *RevokeApplicationFromOrganizationalUnitsResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeApplicationFromOrganizationalUnitsResponse) SetBody(v *RevokeApplicationFromOrganizationalUnitsResponseBody) *RevokeApplicationFromOrganizationalUnitsResponse {
	s.Body = v
	return s
}

func (s *RevokeApplicationFromOrganizationalUnitsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
