// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOrganizationalUnitsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOrganizationalUnitsResponse
	GetStatusCode() *int32
	SetBody(v *ListOrganizationalUnitsResponseBody) *ListOrganizationalUnitsResponse
	GetBody() *ListOrganizationalUnitsResponseBody
}

type ListOrganizationalUnitsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrganizationalUnitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrganizationalUnitsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOrganizationalUnitsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOrganizationalUnitsResponse) GetBody() *ListOrganizationalUnitsResponseBody {
	return s.Body
}

func (s *ListOrganizationalUnitsResponse) SetHeaders(v map[string]*string) *ListOrganizationalUnitsResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationalUnitsResponse) SetStatusCode(v int32) *ListOrganizationalUnitsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationalUnitsResponse) SetBody(v *ListOrganizationalUnitsResponseBody) *ListOrganizationalUnitsResponse {
	s.Body = v
	return s
}

func (s *ListOrganizationalUnitsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
