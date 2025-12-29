// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomQAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomQAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomQAResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomQAResponseBody) *ListCustomQAResponse
	GetBody() *ListCustomQAResponseBody
}

type ListCustomQAResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomQAResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomQAResponse) GoString() string {
	return s.String()
}

func (s *ListCustomQAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomQAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomQAResponse) GetBody() *ListCustomQAResponseBody {
	return s.Body
}

func (s *ListCustomQAResponse) SetHeaders(v map[string]*string) *ListCustomQAResponse {
	s.Headers = v
	return s
}

func (s *ListCustomQAResponse) SetStatusCode(v int32) *ListCustomQAResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomQAResponse) SetBody(v *ListCustomQAResponseBody) *ListCustomQAResponse {
	s.Body = v
	return s
}

func (s *ListCustomQAResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
