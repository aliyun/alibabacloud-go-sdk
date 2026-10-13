// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductPublicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProductPublicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProductPublicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListProductPublicationsResponseBody) *ListProductPublicationsResponse
	GetBody() *ListProductPublicationsResponseBody
}

type ListProductPublicationsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductPublicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductPublicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProductPublicationsResponse) GoString() string {
	return s.String()
}

func (s *ListProductPublicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProductPublicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProductPublicationsResponse) GetBody() *ListProductPublicationsResponseBody {
	return s.Body
}

func (s *ListProductPublicationsResponse) SetHeaders(v map[string]*string) *ListProductPublicationsResponse {
	s.Headers = v
	return s
}

func (s *ListProductPublicationsResponse) SetStatusCode(v int32) *ListProductPublicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductPublicationsResponse) SetBody(v *ListProductPublicationsResponseBody) *ListProductPublicationsResponse {
	s.Body = v
	return s
}

func (s *ListProductPublicationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
