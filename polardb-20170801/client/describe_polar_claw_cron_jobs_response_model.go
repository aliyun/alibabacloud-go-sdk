// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawCronJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarClawCronJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarClawCronJobsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarClawCronJobsResponseBody) *DescribePolarClawCronJobsResponse
	GetBody() *DescribePolarClawCronJobsResponseBody
}

type DescribePolarClawCronJobsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarClawCronJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarClawCronJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawCronJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarClawCronJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarClawCronJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarClawCronJobsResponse) GetBody() *DescribePolarClawCronJobsResponseBody {
	return s.Body
}

func (s *DescribePolarClawCronJobsResponse) SetHeaders(v map[string]*string) *DescribePolarClawCronJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarClawCronJobsResponse) SetStatusCode(v int32) *DescribePolarClawCronJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarClawCronJobsResponse) SetBody(v *DescribePolarClawCronJobsResponseBody) *DescribePolarClawCronJobsResponse {
	s.Body = v
	return s
}

func (s *DescribePolarClawCronJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
