// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDtsJobDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDtsJobDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDtsJobDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDtsJobDetailResponseBody) *DescribeDtsJobDetailResponse
	GetBody() *DescribeDtsJobDetailResponseBody
}

type DescribeDtsJobDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDtsJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDtsJobDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDtsJobDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDtsJobDetailResponse) GetBody() *DescribeDtsJobDetailResponseBody {
	return s.Body
}

func (s *DescribeDtsJobDetailResponse) SetHeaders(v map[string]*string) *DescribeDtsJobDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDtsJobDetailResponse) SetStatusCode(v int32) *DescribeDtsJobDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDtsJobDetailResponse) SetBody(v *DescribeDtsJobDetailResponseBody) *DescribeDtsJobDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeDtsJobDetailResponse) Validate() error {
	return dara.Validate(s)
}
