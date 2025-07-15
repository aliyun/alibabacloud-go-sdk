// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainPublishErrorCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainPublishErrorCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainPublishErrorCodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainPublishErrorCodeResponseBody) *DescribeLiveDomainPublishErrorCodeResponse
	GetBody() *DescribeLiveDomainPublishErrorCodeResponseBody
}

type DescribeLiveDomainPublishErrorCodeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainPublishErrorCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainPublishErrorCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPublishErrorCodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPublishErrorCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainPublishErrorCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainPublishErrorCodeResponse) GetBody() *DescribeLiveDomainPublishErrorCodeResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainPublishErrorCodeResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainPublishErrorCodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeResponse) SetStatusCode(v int32) *DescribeLiveDomainPublishErrorCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeResponse) SetBody(v *DescribeLiveDomainPublishErrorCodeResponseBody) *DescribeLiveDomainPublishErrorCodeResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeResponse) Validate() error {
	return dara.Validate(s)
}
