// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSlsStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainSlsStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainSlsStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainSlsStatusResponseBody) *DescribeDomainSlsStatusResponse
	GetBody() *DescribeDomainSlsStatusResponseBody
}

type DescribeDomainSlsStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainSlsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainSlsStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSlsStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSlsStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainSlsStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainSlsStatusResponse) GetBody() *DescribeDomainSlsStatusResponseBody {
	return s.Body
}

func (s *DescribeDomainSlsStatusResponse) SetHeaders(v map[string]*string) *DescribeDomainSlsStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSlsStatusResponse) SetStatusCode(v int32) *DescribeDomainSlsStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainSlsStatusResponse) SetBody(v *DescribeDomainSlsStatusResponseBody) *DescribeDomainSlsStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainSlsStatusResponse) Validate() error {
	return dara.Validate(s)
}
