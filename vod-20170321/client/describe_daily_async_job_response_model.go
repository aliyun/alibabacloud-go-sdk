// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDailyAsyncJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDailyAsyncJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDailyAsyncJobResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDailyAsyncJobResponseBody) *DescribeDailyAsyncJobResponse
	GetBody() *DescribeDailyAsyncJobResponseBody
}

type DescribeDailyAsyncJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDailyAsyncJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDailyAsyncJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDailyAsyncJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeDailyAsyncJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDailyAsyncJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDailyAsyncJobResponse) GetBody() *DescribeDailyAsyncJobResponseBody {
	return s.Body
}

func (s *DescribeDailyAsyncJobResponse) SetHeaders(v map[string]*string) *DescribeDailyAsyncJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeDailyAsyncJobResponse) SetStatusCode(v int32) *DescribeDailyAsyncJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDailyAsyncJobResponse) SetBody(v *DescribeDailyAsyncJobResponseBody) *DescribeDailyAsyncJobResponse {
	s.Body = v
	return s
}

func (s *DescribeDailyAsyncJobResponse) Validate() error {
	return dara.Validate(s)
}
