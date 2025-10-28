// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConcurrencyConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConcurrencyConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConcurrencyConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListConcurrencyConfigsOutput) *ListConcurrencyConfigsResponse
	GetBody() *ListConcurrencyConfigsOutput
}

type ListConcurrencyConfigsResponse struct {
	Headers    map[string]*string            `json:"headers" xml:"headers"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConcurrencyConfigsOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConcurrencyConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConcurrencyConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListConcurrencyConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConcurrencyConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConcurrencyConfigsResponse) GetBody() *ListConcurrencyConfigsOutput {
	return s.Body
}

func (s *ListConcurrencyConfigsResponse) SetHeaders(v map[string]*string) *ListConcurrencyConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListConcurrencyConfigsResponse) SetStatusCode(v int32) *ListConcurrencyConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConcurrencyConfigsResponse) SetBody(v *ListConcurrencyConfigsOutput) *ListConcurrencyConfigsResponse {
	s.Body = v
	return s
}

func (s *ListConcurrencyConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
