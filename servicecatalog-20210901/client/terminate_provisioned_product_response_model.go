// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateProvisionedProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TerminateProvisionedProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TerminateProvisionedProductResponse
	GetStatusCode() *int32
	SetBody(v *TerminateProvisionedProductResponseBody) *TerminateProvisionedProductResponse
	GetBody() *TerminateProvisionedProductResponseBody
}

type TerminateProvisionedProductResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminateProvisionedProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminateProvisionedProductResponse) String() string {
	return dara.Prettify(s)
}

func (s TerminateProvisionedProductResponse) GoString() string {
	return s.String()
}

func (s *TerminateProvisionedProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TerminateProvisionedProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TerminateProvisionedProductResponse) GetBody() *TerminateProvisionedProductResponseBody {
	return s.Body
}

func (s *TerminateProvisionedProductResponse) SetHeaders(v map[string]*string) *TerminateProvisionedProductResponse {
	s.Headers = v
	return s
}

func (s *TerminateProvisionedProductResponse) SetStatusCode(v int32) *TerminateProvisionedProductResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminateProvisionedProductResponse) SetBody(v *TerminateProvisionedProductResponseBody) *TerminateProvisionedProductResponse {
	s.Body = v
	return s
}

func (s *TerminateProvisionedProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
