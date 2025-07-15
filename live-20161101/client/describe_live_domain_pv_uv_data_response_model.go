// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainPvUvDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainPvUvDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainPvUvDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainPvUvDataResponseBody) *DescribeLiveDomainPvUvDataResponse
	GetBody() *DescribeLiveDomainPvUvDataResponseBody
}

type DescribeLiveDomainPvUvDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainPvUvDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainPvUvDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPvUvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPvUvDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainPvUvDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainPvUvDataResponse) GetBody() *DescribeLiveDomainPvUvDataResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainPvUvDataResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainPvUvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainPvUvDataResponse) SetStatusCode(v int32) *DescribeLiveDomainPvUvDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainPvUvDataResponse) SetBody(v *DescribeLiveDomainPvUvDataResponseBody) *DescribeLiveDomainPvUvDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainPvUvDataResponse) Validate() error {
	return dara.Validate(s)
}
