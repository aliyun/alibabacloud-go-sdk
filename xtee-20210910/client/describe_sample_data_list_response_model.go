// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDataListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSampleDataListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSampleDataListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSampleDataListResponseBody) *DescribeSampleDataListResponse
	GetBody() *DescribeSampleDataListResponseBody
}

type DescribeSampleDataListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSampleDataListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSampleDataListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDataListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSampleDataListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSampleDataListResponse) GetBody() *DescribeSampleDataListResponseBody {
	return s.Body
}

func (s *DescribeSampleDataListResponse) SetHeaders(v map[string]*string) *DescribeSampleDataListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSampleDataListResponse) SetStatusCode(v int32) *DescribeSampleDataListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSampleDataListResponse) SetBody(v *DescribeSampleDataListResponseBody) *DescribeSampleDataListResponse {
	s.Body = v
	return s
}

func (s *DescribeSampleDataListResponse) Validate() error {
	return dara.Validate(s)
}
