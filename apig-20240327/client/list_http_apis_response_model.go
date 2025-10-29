// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHttpApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHttpApisResponse
	GetStatusCode() *int32
	SetBody(v *ListHttpApisResponseBody) *ListHttpApisResponse
	GetBody() *ListHttpApisResponseBody
}

type ListHttpApisResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHttpApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHttpApisResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHttpApisResponse) GoString() string {
	return s.String()
}

func (s *ListHttpApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHttpApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHttpApisResponse) GetBody() *ListHttpApisResponseBody {
	return s.Body
}

func (s *ListHttpApisResponse) SetHeaders(v map[string]*string) *ListHttpApisResponse {
	s.Headers = v
	return s
}

func (s *ListHttpApisResponse) SetStatusCode(v int32) *ListHttpApisResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHttpApisResponse) SetBody(v *ListHttpApisResponseBody) *ListHttpApisResponse {
	s.Body = v
	return s
}

func (s *ListHttpApisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
