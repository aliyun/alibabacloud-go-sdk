// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSampleDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSampleDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSampleDataResponseBody) *DescribeSampleDataResponse
	GetBody() *DescribeSampleDataResponseBody
}

type DescribeSampleDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSampleDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSampleDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSampleDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSampleDataResponse) GetBody() *DescribeSampleDataResponseBody {
	return s.Body
}

func (s *DescribeSampleDataResponse) SetHeaders(v map[string]*string) *DescribeSampleDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeSampleDataResponse) SetStatusCode(v int32) *DescribeSampleDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSampleDataResponse) SetBody(v *DescribeSampleDataResponseBody) *DescribeSampleDataResponse {
	s.Body = v
	return s
}

func (s *DescribeSampleDataResponse) Validate() error {
	return dara.Validate(s)
}
