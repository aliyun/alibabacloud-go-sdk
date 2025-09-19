// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSuspEventDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSuspEventDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSuspEventDetailResponseBody) *DescribeSuspEventDetailResponse
	GetBody() *DescribeSuspEventDetailResponseBody
}

type DescribeSuspEventDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSuspEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSuspEventDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSuspEventDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSuspEventDetailResponse) GetBody() *DescribeSuspEventDetailResponseBody {
	return s.Body
}

func (s *DescribeSuspEventDetailResponse) SetHeaders(v map[string]*string) *DescribeSuspEventDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspEventDetailResponse) SetStatusCode(v int32) *DescribeSuspEventDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSuspEventDetailResponse) SetBody(v *DescribeSuspEventDetailResponseBody) *DescribeSuspEventDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeSuspEventDetailResponse) Validate() error {
	return dara.Validate(s)
}
