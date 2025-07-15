// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainEdgeLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainEdgeLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainEdgeLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainEdgeLogResponseBody) *DescribeLiveDomainEdgeLogResponse
	GetBody() *DescribeLiveDomainEdgeLogResponseBody
}

type DescribeLiveDomainEdgeLogResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainEdgeLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainEdgeLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainEdgeLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainEdgeLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainEdgeLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainEdgeLogResponse) GetBody() *DescribeLiveDomainEdgeLogResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainEdgeLogResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainEdgeLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainEdgeLogResponse) SetStatusCode(v int32) *DescribeLiveDomainEdgeLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainEdgeLogResponse) SetBody(v *DescribeLiveDomainEdgeLogResponseBody) *DescribeLiveDomainEdgeLogResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainEdgeLogResponse) Validate() error {
	return dara.Validate(s)
}
