// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspiciousOverallConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSuspiciousOverallConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSuspiciousOverallConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSuspiciousOverallConfigResponseBody) *DescribeSuspiciousOverallConfigResponse
	GetBody() *DescribeSuspiciousOverallConfigResponseBody
}

type DescribeSuspiciousOverallConfigResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSuspiciousOverallConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSuspiciousOverallConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspiciousOverallConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspiciousOverallConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSuspiciousOverallConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSuspiciousOverallConfigResponse) GetBody() *DescribeSuspiciousOverallConfigResponseBody {
	return s.Body
}

func (s *DescribeSuspiciousOverallConfigResponse) SetHeaders(v map[string]*string) *DescribeSuspiciousOverallConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspiciousOverallConfigResponse) SetStatusCode(v int32) *DescribeSuspiciousOverallConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSuspiciousOverallConfigResponse) SetBody(v *DescribeSuspiciousOverallConfigResponseBody) *DescribeSuspiciousOverallConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeSuspiciousOverallConfigResponse) Validate() error {
	return dara.Validate(s)
}
