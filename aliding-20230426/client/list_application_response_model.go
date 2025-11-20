// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationResponseBody) *ListApplicationResponse
	GetBody() *ListApplicationResponseBody
}

type ListApplicationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationResponse) GetBody() *ListApplicationResponseBody {
	return s.Body
}

func (s *ListApplicationResponse) SetHeaders(v map[string]*string) *ListApplicationResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationResponse) SetStatusCode(v int32) *ListApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationResponse) SetBody(v *ListApplicationResponseBody) *ListApplicationResponse {
	s.Body = v
	return s
}

func (s *ListApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
