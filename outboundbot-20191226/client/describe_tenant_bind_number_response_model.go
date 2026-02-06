// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTenantBindNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTenantBindNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTenantBindNumberResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTenantBindNumberResponseBody) *DescribeTenantBindNumberResponse
	GetBody() *DescribeTenantBindNumberResponseBody
}

type DescribeTenantBindNumberResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTenantBindNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTenantBindNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTenantBindNumberResponse) GoString() string {
	return s.String()
}

func (s *DescribeTenantBindNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTenantBindNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTenantBindNumberResponse) GetBody() *DescribeTenantBindNumberResponseBody {
	return s.Body
}

func (s *DescribeTenantBindNumberResponse) SetHeaders(v map[string]*string) *DescribeTenantBindNumberResponse {
	s.Headers = v
	return s
}

func (s *DescribeTenantBindNumberResponse) SetStatusCode(v int32) *DescribeTenantBindNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTenantBindNumberResponse) SetBody(v *DescribeTenantBindNumberResponseBody) *DescribeTenantBindNumberResponse {
	s.Body = v
	return s
}

func (s *DescribeTenantBindNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
