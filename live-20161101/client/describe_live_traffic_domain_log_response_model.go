// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveTrafficDomainLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveTrafficDomainLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveTrafficDomainLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveTrafficDomainLogResponseBody) *DescribeLiveTrafficDomainLogResponse
	GetBody() *DescribeLiveTrafficDomainLogResponseBody
}

type DescribeLiveTrafficDomainLogResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveTrafficDomainLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveTrafficDomainLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTrafficDomainLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveTrafficDomainLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveTrafficDomainLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveTrafficDomainLogResponse) GetBody() *DescribeLiveTrafficDomainLogResponseBody {
	return s.Body
}

func (s *DescribeLiveTrafficDomainLogResponse) SetHeaders(v map[string]*string) *DescribeLiveTrafficDomainLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponse) SetStatusCode(v int32) *DescribeLiveTrafficDomainLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponse) SetBody(v *DescribeLiveTrafficDomainLogResponseBody) *DescribeLiveTrafficDomainLogResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponse) Validate() error {
	return dara.Validate(s)
}
