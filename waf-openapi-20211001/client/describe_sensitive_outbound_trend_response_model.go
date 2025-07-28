// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveOutboundTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSensitiveOutboundTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSensitiveOutboundTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSensitiveOutboundTrendResponseBody) *DescribeSensitiveOutboundTrendResponse
	GetBody() *DescribeSensitiveOutboundTrendResponseBody
}

type DescribeSensitiveOutboundTrendResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSensitiveOutboundTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSensitiveOutboundTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveOutboundTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveOutboundTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSensitiveOutboundTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSensitiveOutboundTrendResponse) GetBody() *DescribeSensitiveOutboundTrendResponseBody {
	return s.Body
}

func (s *DescribeSensitiveOutboundTrendResponse) SetHeaders(v map[string]*string) *DescribeSensitiveOutboundTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeSensitiveOutboundTrendResponse) SetStatusCode(v int32) *DescribeSensitiveOutboundTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSensitiveOutboundTrendResponse) SetBody(v *DescribeSensitiveOutboundTrendResponseBody) *DescribeSensitiveOutboundTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeSensitiveOutboundTrendResponse) Validate() error {
	return dara.Validate(s)
}
