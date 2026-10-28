// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfsTenantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePxfsTenantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePxfsTenantResponse
	GetStatusCode() *int32
	SetBody(v *DescribePxfsTenantResponseBody) *DescribePxfsTenantResponse
	GetBody() *DescribePxfsTenantResponseBody
}

type DescribePxfsTenantResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePxfsTenantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePxfsTenantResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsTenantResponse) GoString() string {
	return s.String()
}

func (s *DescribePxfsTenantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePxfsTenantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePxfsTenantResponse) GetBody() *DescribePxfsTenantResponseBody {
	return s.Body
}

func (s *DescribePxfsTenantResponse) SetHeaders(v map[string]*string) *DescribePxfsTenantResponse {
	s.Headers = v
	return s
}

func (s *DescribePxfsTenantResponse) SetStatusCode(v int32) *DescribePxfsTenantResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePxfsTenantResponse) SetBody(v *DescribePxfsTenantResponseBody) *DescribePxfsTenantResponse {
	s.Body = v
	return s
}

func (s *DescribePxfsTenantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
