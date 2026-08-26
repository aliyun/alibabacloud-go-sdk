// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrustedOriginsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrustedOriginsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrustedOriginsResponse
	GetStatusCode() *int32
	SetBody(v *ListTrustedOriginsResponseBody) *ListTrustedOriginsResponse
	GetBody() *ListTrustedOriginsResponseBody
}

type ListTrustedOriginsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrustedOriginsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrustedOriginsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrustedOriginsResponse) GoString() string {
	return s.String()
}

func (s *ListTrustedOriginsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrustedOriginsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrustedOriginsResponse) GetBody() *ListTrustedOriginsResponseBody {
	return s.Body
}

func (s *ListTrustedOriginsResponse) SetHeaders(v map[string]*string) *ListTrustedOriginsResponse {
	s.Headers = v
	return s
}

func (s *ListTrustedOriginsResponse) SetStatusCode(v int32) *ListTrustedOriginsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrustedOriginsResponse) SetBody(v *ListTrustedOriginsResponseBody) *ListTrustedOriginsResponse {
	s.Body = v
	return s
}

func (s *ListTrustedOriginsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
