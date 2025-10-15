// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOrganizationalUnitChildrenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOrganizationalUnitChildrenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOrganizationalUnitChildrenResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOrganizationalUnitChildrenResponseBody) *DeleteOrganizationalUnitChildrenResponse
	GetBody() *DeleteOrganizationalUnitChildrenResponseBody
}

type DeleteOrganizationalUnitChildrenResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOrganizationalUnitChildrenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOrganizationalUnitChildrenResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOrganizationalUnitChildrenResponse) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationalUnitChildrenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOrganizationalUnitChildrenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOrganizationalUnitChildrenResponse) GetBody() *DeleteOrganizationalUnitChildrenResponseBody {
	return s.Body
}

func (s *DeleteOrganizationalUnitChildrenResponse) SetHeaders(v map[string]*string) *DeleteOrganizationalUnitChildrenResponse {
	s.Headers = v
	return s
}

func (s *DeleteOrganizationalUnitChildrenResponse) SetStatusCode(v int32) *DeleteOrganizationalUnitChildrenResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOrganizationalUnitChildrenResponse) SetBody(v *DeleteOrganizationalUnitChildrenResponseBody) *DeleteOrganizationalUnitChildrenResponse {
	s.Body = v
	return s
}

func (s *DeleteOrganizationalUnitChildrenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
