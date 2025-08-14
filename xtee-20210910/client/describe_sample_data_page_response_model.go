// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDataPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSampleDataPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSampleDataPageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSampleDataPageResponseBody) *DescribeSampleDataPageResponse
	GetBody() *DescribeSampleDataPageResponseBody
}

type DescribeSampleDataPageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSampleDataPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSampleDataPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDataPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSampleDataPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSampleDataPageResponse) GetBody() *DescribeSampleDataPageResponseBody {
	return s.Body
}

func (s *DescribeSampleDataPageResponse) SetHeaders(v map[string]*string) *DescribeSampleDataPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeSampleDataPageResponse) SetStatusCode(v int32) *DescribeSampleDataPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSampleDataPageResponse) SetBody(v *DescribeSampleDataPageResponseBody) *DescribeSampleDataPageResponse {
	s.Body = v
	return s
}

func (s *DescribeSampleDataPageResponse) Validate() error {
	return dara.Validate(s)
}
