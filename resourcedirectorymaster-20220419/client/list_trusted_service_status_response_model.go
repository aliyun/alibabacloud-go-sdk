// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrustedServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrustedServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrustedServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *ListTrustedServiceStatusResponseBody) *ListTrustedServiceStatusResponse
	GetBody() *ListTrustedServiceStatusResponseBody
}

type ListTrustedServiceStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrustedServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrustedServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrustedServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrustedServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrustedServiceStatusResponse) GetBody() *ListTrustedServiceStatusResponseBody {
	return s.Body
}

func (s *ListTrustedServiceStatusResponse) SetHeaders(v map[string]*string) *ListTrustedServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *ListTrustedServiceStatusResponse) SetStatusCode(v int32) *ListTrustedServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrustedServiceStatusResponse) SetBody(v *ListTrustedServiceStatusResponseBody) *ListTrustedServiceStatusResponse {
	s.Body = v
	return s
}

func (s *ListTrustedServiceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
