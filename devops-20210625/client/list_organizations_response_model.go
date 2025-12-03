// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOrganizationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOrganizationsResponse
	GetStatusCode() *int32
	SetBody(v *ListOrganizationsResponseBody) *ListOrganizationsResponse
	GetBody() *ListOrganizationsResponseBody
}

type ListOrganizationsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrganizationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrganizationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationsResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOrganizationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOrganizationsResponse) GetBody() *ListOrganizationsResponseBody {
	return s.Body
}

func (s *ListOrganizationsResponse) SetHeaders(v map[string]*string) *ListOrganizationsResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationsResponse) SetStatusCode(v int32) *ListOrganizationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationsResponse) SetBody(v *ListOrganizationsResponseBody) *ListOrganizationsResponse {
	s.Body = v
	return s
}

func (s *ListOrganizationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
