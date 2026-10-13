// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPortalsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPortalsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPortalsResponse
	GetStatusCode() *int32
	SetBody(v *ListPortalsResponseBody) *ListPortalsResponse
	GetBody() *ListPortalsResponseBody
}

type ListPortalsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPortalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPortalsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPortalsResponse) GoString() string {
	return s.String()
}

func (s *ListPortalsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPortalsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPortalsResponse) GetBody() *ListPortalsResponseBody {
	return s.Body
}

func (s *ListPortalsResponse) SetHeaders(v map[string]*string) *ListPortalsResponse {
	s.Headers = v
	return s
}

func (s *ListPortalsResponse) SetStatusCode(v int32) *ListPortalsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPortalsResponse) SetBody(v *ListPortalsResponseBody) *ListPortalsResponse {
	s.Body = v
	return s
}

func (s *ListPortalsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
