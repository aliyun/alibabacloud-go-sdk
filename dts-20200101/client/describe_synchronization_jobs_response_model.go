// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSynchronizationJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSynchronizationJobsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSynchronizationJobsResponseBody) *DescribeSynchronizationJobsResponse
	GetBody() *DescribeSynchronizationJobsResponseBody
}

type DescribeSynchronizationJobsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSynchronizationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSynchronizationJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSynchronizationJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSynchronizationJobsResponse) GetBody() *DescribeSynchronizationJobsResponseBody {
	return s.Body
}

func (s *DescribeSynchronizationJobsResponse) SetHeaders(v map[string]*string) *DescribeSynchronizationJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynchronizationJobsResponse) SetStatusCode(v int32) *DescribeSynchronizationJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynchronizationJobsResponse) SetBody(v *DescribeSynchronizationJobsResponseBody) *DescribeSynchronizationJobsResponse {
	s.Body = v
	return s
}

func (s *DescribeSynchronizationJobsResponse) Validate() error {
	return dara.Validate(s)
}
