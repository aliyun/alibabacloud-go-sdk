// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokePageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokePageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokePageResponse
	GetStatusCode() *int32
	SetBody(v *InvokePageResponseBody) *InvokePageResponse
	GetBody() *InvokePageResponseBody
}

type InvokePageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokePageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokePageResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokePageResponse) GoString() string {
	return s.String()
}

func (s *InvokePageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokePageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokePageResponse) GetBody() *InvokePageResponseBody {
	return s.Body
}

func (s *InvokePageResponse) SetHeaders(v map[string]*string) *InvokePageResponse {
	s.Headers = v
	return s
}

func (s *InvokePageResponse) SetStatusCode(v int32) *InvokePageResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokePageResponse) SetBody(v *InvokePageResponseBody) *InvokePageResponse {
	s.Body = v
	return s
}

func (s *InvokePageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
