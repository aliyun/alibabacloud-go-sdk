// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSampleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSampleListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSampleListResponseBody) *DescribeSampleListResponse
	GetBody() *DescribeSampleListResponseBody
}

type DescribeSampleListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSampleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSampleListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSampleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSampleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSampleListResponse) GetBody() *DescribeSampleListResponseBody {
	return s.Body
}

func (s *DescribeSampleListResponse) SetHeaders(v map[string]*string) *DescribeSampleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSampleListResponse) SetStatusCode(v int32) *DescribeSampleListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSampleListResponse) SetBody(v *DescribeSampleListResponseBody) *DescribeSampleListResponse {
	s.Body = v
	return s
}

func (s *DescribeSampleListResponse) Validate() error {
	return dara.Validate(s)
}
