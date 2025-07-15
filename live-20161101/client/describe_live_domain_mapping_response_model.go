// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainMappingResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainMappingResponseBody) *DescribeLiveDomainMappingResponse
	GetBody() *DescribeLiveDomainMappingResponseBody
}

type DescribeLiveDomainMappingResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainMappingResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainMappingResponse) GetBody() *DescribeLiveDomainMappingResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainMappingResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainMappingResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainMappingResponse) SetStatusCode(v int32) *DescribeLiveDomainMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainMappingResponse) SetBody(v *DescribeLiveDomainMappingResponseBody) *DescribeLiveDomainMappingResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainMappingResponse) Validate() error {
	return dara.Validate(s)
}
