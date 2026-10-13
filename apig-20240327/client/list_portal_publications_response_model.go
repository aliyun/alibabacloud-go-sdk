// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPortalPublicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPortalPublicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPortalPublicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListPortalPublicationsResponseBody) *ListPortalPublicationsResponse
	GetBody() *ListPortalPublicationsResponseBody
}

type ListPortalPublicationsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPortalPublicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPortalPublicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPortalPublicationsResponse) GoString() string {
	return s.String()
}

func (s *ListPortalPublicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPortalPublicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPortalPublicationsResponse) GetBody() *ListPortalPublicationsResponseBody {
	return s.Body
}

func (s *ListPortalPublicationsResponse) SetHeaders(v map[string]*string) *ListPortalPublicationsResponse {
	s.Headers = v
	return s
}

func (s *ListPortalPublicationsResponse) SetStatusCode(v int32) *ListPortalPublicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPortalPublicationsResponse) SetBody(v *ListPortalPublicationsResponseBody) *ListPortalPublicationsResponse {
	s.Body = v
	return s
}

func (s *ListPortalPublicationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
