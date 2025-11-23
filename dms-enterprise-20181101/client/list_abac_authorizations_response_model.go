// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAbacAuthorizationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAbacAuthorizationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAbacAuthorizationsResponse
	GetStatusCode() *int32
	SetBody(v *ListAbacAuthorizationsResponseBody) *ListAbacAuthorizationsResponse
	GetBody() *ListAbacAuthorizationsResponseBody
}

type ListAbacAuthorizationsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAbacAuthorizationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAbacAuthorizationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAbacAuthorizationsResponse) GoString() string {
	return s.String()
}

func (s *ListAbacAuthorizationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAbacAuthorizationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAbacAuthorizationsResponse) GetBody() *ListAbacAuthorizationsResponseBody {
	return s.Body
}

func (s *ListAbacAuthorizationsResponse) SetHeaders(v map[string]*string) *ListAbacAuthorizationsResponse {
	s.Headers = v
	return s
}

func (s *ListAbacAuthorizationsResponse) SetStatusCode(v int32) *ListAbacAuthorizationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAbacAuthorizationsResponse) SetBody(v *ListAbacAuthorizationsResponseBody) *ListAbacAuthorizationsResponse {
	s.Body = v
	return s
}

func (s *ListAbacAuthorizationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
