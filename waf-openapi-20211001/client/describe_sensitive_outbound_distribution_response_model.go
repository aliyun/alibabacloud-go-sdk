// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveOutboundDistributionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSensitiveOutboundDistributionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSensitiveOutboundDistributionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSensitiveOutboundDistributionResponseBody) *DescribeSensitiveOutboundDistributionResponse
	GetBody() *DescribeSensitiveOutboundDistributionResponseBody
}

type DescribeSensitiveOutboundDistributionResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSensitiveOutboundDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSensitiveOutboundDistributionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveOutboundDistributionResponse) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveOutboundDistributionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSensitiveOutboundDistributionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSensitiveOutboundDistributionResponse) GetBody() *DescribeSensitiveOutboundDistributionResponseBody {
	return s.Body
}

func (s *DescribeSensitiveOutboundDistributionResponse) SetHeaders(v map[string]*string) *DescribeSensitiveOutboundDistributionResponse {
	s.Headers = v
	return s
}

func (s *DescribeSensitiveOutboundDistributionResponse) SetStatusCode(v int32) *DescribeSensitiveOutboundDistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSensitiveOutboundDistributionResponse) SetBody(v *DescribeSensitiveOutboundDistributionResponseBody) *DescribeSensitiveOutboundDistributionResponse {
	s.Body = v
	return s
}

func (s *DescribeSensitiveOutboundDistributionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
