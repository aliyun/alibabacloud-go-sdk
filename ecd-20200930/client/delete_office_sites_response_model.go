// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOfficeSitesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOfficeSitesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOfficeSitesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOfficeSitesResponseBody) *DeleteOfficeSitesResponse
	GetBody() *DeleteOfficeSitesResponseBody
}

type DeleteOfficeSitesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOfficeSitesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOfficeSitesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOfficeSitesResponse) GoString() string {
	return s.String()
}

func (s *DeleteOfficeSitesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOfficeSitesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOfficeSitesResponse) GetBody() *DeleteOfficeSitesResponseBody {
	return s.Body
}

func (s *DeleteOfficeSitesResponse) SetHeaders(v map[string]*string) *DeleteOfficeSitesResponse {
	s.Headers = v
	return s
}

func (s *DeleteOfficeSitesResponse) SetStatusCode(v int32) *DeleteOfficeSitesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOfficeSitesResponse) SetBody(v *DeleteOfficeSitesResponseBody) *DeleteOfficeSitesResponse {
	s.Body = v
	return s
}

func (s *DeleteOfficeSitesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
