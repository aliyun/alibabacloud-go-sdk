// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestModelProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestModelProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestModelProviderResponse
	GetStatusCode() *int32
	SetBody(v *TestModelProviderResponseBody) *TestModelProviderResponse
	GetBody() *TestModelProviderResponseBody
}

type TestModelProviderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestModelProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestModelProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s TestModelProviderResponse) GoString() string {
	return s.String()
}

func (s *TestModelProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestModelProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestModelProviderResponse) GetBody() *TestModelProviderResponseBody {
	return s.Body
}

func (s *TestModelProviderResponse) SetHeaders(v map[string]*string) *TestModelProviderResponse {
	s.Headers = v
	return s
}

func (s *TestModelProviderResponse) SetStatusCode(v int32) *TestModelProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *TestModelProviderResponse) SetBody(v *TestModelProviderResponseBody) *TestModelProviderResponse {
	s.Body = v
	return s
}

func (s *TestModelProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
