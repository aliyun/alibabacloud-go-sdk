// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveOutboundStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSensitiveOutboundStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSensitiveOutboundStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSensitiveOutboundStatisticResponseBody) *DescribeSensitiveOutboundStatisticResponse
	GetBody() *DescribeSensitiveOutboundStatisticResponseBody
}

type DescribeSensitiveOutboundStatisticResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSensitiveOutboundStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSensitiveOutboundStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveOutboundStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveOutboundStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSensitiveOutboundStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSensitiveOutboundStatisticResponse) GetBody() *DescribeSensitiveOutboundStatisticResponseBody {
	return s.Body
}

func (s *DescribeSensitiveOutboundStatisticResponse) SetHeaders(v map[string]*string) *DescribeSensitiveOutboundStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeSensitiveOutboundStatisticResponse) SetStatusCode(v int32) *DescribeSensitiveOutboundStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticResponse) SetBody(v *DescribeSensitiveOutboundStatisticResponseBody) *DescribeSensitiveOutboundStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeSensitiveOutboundStatisticResponse) Validate() error {
	return dara.Validate(s)
}
