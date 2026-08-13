// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBillingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBillingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBillingResponse
	GetStatusCode() *int32
	SetBody(v *ListBillingResponseBody) *ListBillingResponse
	GetBody() *ListBillingResponseBody
}

type ListBillingResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBillingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBillingResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBillingResponse) GoString() string {
	return s.String()
}

func (s *ListBillingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBillingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBillingResponse) GetBody() *ListBillingResponseBody {
	return s.Body
}

func (s *ListBillingResponse) SetHeaders(v map[string]*string) *ListBillingResponse {
	s.Headers = v
	return s
}

func (s *ListBillingResponse) SetStatusCode(v int32) *ListBillingResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBillingResponse) SetBody(v *ListBillingResponseBody) *ListBillingResponse {
	s.Body = v
	return s
}

func (s *ListBillingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
