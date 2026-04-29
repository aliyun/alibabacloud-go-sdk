// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolarClawBindingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPolarClawBindingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPolarClawBindingsResponse
	GetStatusCode() *int32
	SetBody(v *ListPolarClawBindingsResponseBody) *ListPolarClawBindingsResponse
	GetBody() *ListPolarClawBindingsResponseBody
}

type ListPolarClawBindingsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolarClawBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolarClawBindingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPolarClawBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListPolarClawBindingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPolarClawBindingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPolarClawBindingsResponse) GetBody() *ListPolarClawBindingsResponseBody {
	return s.Body
}

func (s *ListPolarClawBindingsResponse) SetHeaders(v map[string]*string) *ListPolarClawBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListPolarClawBindingsResponse) SetStatusCode(v int32) *ListPolarClawBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolarClawBindingsResponse) SetBody(v *ListPolarClawBindingsResponseBody) *ListPolarClawBindingsResponse {
	s.Body = v
	return s
}

func (s *ListPolarClawBindingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
