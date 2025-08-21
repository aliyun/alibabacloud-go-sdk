// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopUserAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainTopUserAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainTopUserAgentResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainTopUserAgentResponseBody) *DescribeDomainTopUserAgentResponse
	GetBody() *DescribeDomainTopUserAgentResponseBody
}

type DescribeDomainTopUserAgentResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainTopUserAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainTopUserAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUserAgentResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUserAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainTopUserAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainTopUserAgentResponse) GetBody() *DescribeDomainTopUserAgentResponseBody {
	return s.Body
}

func (s *DescribeDomainTopUserAgentResponse) SetHeaders(v map[string]*string) *DescribeDomainTopUserAgentResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainTopUserAgentResponse) SetStatusCode(v int32) *DescribeDomainTopUserAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainTopUserAgentResponse) SetBody(v *DescribeDomainTopUserAgentResponseBody) *DescribeDomainTopUserAgentResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainTopUserAgentResponse) Validate() error {
	return dara.Validate(s)
}
