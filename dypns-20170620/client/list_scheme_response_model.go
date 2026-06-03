// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSchemeResponse
	GetStatusCode() *int32
	SetBody(v *ListSchemeResponseBody) *ListSchemeResponse
	GetBody() *ListSchemeResponseBody
}

type ListSchemeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeResponse) GoString() string {
	return s.String()
}

func (s *ListSchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSchemeResponse) GetBody() *ListSchemeResponseBody {
	return s.Body
}

func (s *ListSchemeResponse) SetHeaders(v map[string]*string) *ListSchemeResponse {
	s.Headers = v
	return s
}

func (s *ListSchemeResponse) SetStatusCode(v int32) *ListSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSchemeResponse) SetBody(v *ListSchemeResponseBody) *ListSchemeResponse {
	s.Body = v
	return s
}

func (s *ListSchemeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
