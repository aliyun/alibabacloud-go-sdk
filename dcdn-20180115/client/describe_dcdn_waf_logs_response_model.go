// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafLogsResponseBody) *DescribeDcdnWafLogsResponse
	GetBody() *DescribeDcdnWafLogsResponseBody
}

type DescribeDcdnWafLogsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafLogsResponse) GetBody() *DescribeDcdnWafLogsResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafLogsResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafLogsResponse) SetStatusCode(v int32) *DescribeDcdnWafLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafLogsResponse) SetBody(v *DescribeDcdnWafLogsResponseBody) *DescribeDcdnWafLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafLogsResponse) Validate() error {
	return dara.Validate(s)
}
