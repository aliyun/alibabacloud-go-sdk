// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCapabilitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCapabilitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCapabilitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListCapabilitiesResponseBody) *ListCapabilitiesResponse
	GetBody() *ListCapabilitiesResponseBody
}

type ListCapabilitiesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCapabilitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCapabilitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCapabilitiesResponse) GoString() string {
	return s.String()
}

func (s *ListCapabilitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCapabilitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCapabilitiesResponse) GetBody() *ListCapabilitiesResponseBody {
	return s.Body
}

func (s *ListCapabilitiesResponse) SetHeaders(v map[string]*string) *ListCapabilitiesResponse {
	s.Headers = v
	return s
}

func (s *ListCapabilitiesResponse) SetStatusCode(v int32) *ListCapabilitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCapabilitiesResponse) SetBody(v *ListCapabilitiesResponseBody) *ListCapabilitiesResponse {
	s.Body = v
	return s
}

func (s *ListCapabilitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
