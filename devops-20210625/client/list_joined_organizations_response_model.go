// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJoinedOrganizationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJoinedOrganizationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJoinedOrganizationsResponse
	GetStatusCode() *int32
	SetBody(v *ListJoinedOrganizationsResponseBody) *ListJoinedOrganizationsResponse
	GetBody() *ListJoinedOrganizationsResponseBody
}

type ListJoinedOrganizationsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJoinedOrganizationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJoinedOrganizationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJoinedOrganizationsResponse) GoString() string {
	return s.String()
}

func (s *ListJoinedOrganizationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJoinedOrganizationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJoinedOrganizationsResponse) GetBody() *ListJoinedOrganizationsResponseBody {
	return s.Body
}

func (s *ListJoinedOrganizationsResponse) SetHeaders(v map[string]*string) *ListJoinedOrganizationsResponse {
	s.Headers = v
	return s
}

func (s *ListJoinedOrganizationsResponse) SetStatusCode(v int32) *ListJoinedOrganizationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJoinedOrganizationsResponse) SetBody(v *ListJoinedOrganizationsResponseBody) *ListJoinedOrganizationsResponse {
	s.Body = v
	return s
}

func (s *ListJoinedOrganizationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
