// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMobileAgentPackagePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMobileAgentPackagePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMobileAgentPackagePriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMobileAgentPackagePriceResponseBody) *DescribeMobileAgentPackagePriceResponse
	GetBody() *DescribeMobileAgentPackagePriceResponseBody
}

type DescribeMobileAgentPackagePriceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMobileAgentPackagePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMobileAgentPackagePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackagePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackagePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMobileAgentPackagePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMobileAgentPackagePriceResponse) GetBody() *DescribeMobileAgentPackagePriceResponseBody {
	return s.Body
}

func (s *DescribeMobileAgentPackagePriceResponse) SetHeaders(v map[string]*string) *DescribeMobileAgentPackagePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponse) SetStatusCode(v int32) *DescribeMobileAgentPackagePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponse) SetBody(v *DescribeMobileAgentPackagePriceResponseBody) *DescribeMobileAgentPackagePriceResponse {
	s.Body = v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
