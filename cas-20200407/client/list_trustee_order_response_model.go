// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrusteeOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrusteeOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrusteeOrderResponse
	GetStatusCode() *int32
	SetBody(v *ListTrusteeOrderResponseBody) *ListTrusteeOrderResponse
	GetBody() *ListTrusteeOrderResponseBody
}

type ListTrusteeOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrusteeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrusteeOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrusteeOrderResponse) GoString() string {
	return s.String()
}

func (s *ListTrusteeOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrusteeOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrusteeOrderResponse) GetBody() *ListTrusteeOrderResponseBody {
	return s.Body
}

func (s *ListTrusteeOrderResponse) SetHeaders(v map[string]*string) *ListTrusteeOrderResponse {
	s.Headers = v
	return s
}

func (s *ListTrusteeOrderResponse) SetStatusCode(v int32) *ListTrusteeOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrusteeOrderResponse) SetBody(v *ListTrusteeOrderResponseBody) *ListTrusteeOrderResponse {
	s.Body = v
	return s
}

func (s *ListTrusteeOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
