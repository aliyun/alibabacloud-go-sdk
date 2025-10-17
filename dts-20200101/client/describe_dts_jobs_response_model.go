// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDtsJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDtsJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDtsJobsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDtsJobsResponseBody) *DescribeDtsJobsResponse
	GetBody() *DescribeDtsJobsResponseBody
}

type DescribeDtsJobsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDtsJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDtsJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDtsJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDtsJobsResponse) GetBody() *DescribeDtsJobsResponseBody {
	return s.Body
}

func (s *DescribeDtsJobsResponse) SetHeaders(v map[string]*string) *DescribeDtsJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDtsJobsResponse) SetStatusCode(v int32) *DescribeDtsJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDtsJobsResponse) SetBody(v *DescribeDtsJobsResponseBody) *DescribeDtsJobsResponse {
	s.Body = v
	return s
}

func (s *DescribeDtsJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
