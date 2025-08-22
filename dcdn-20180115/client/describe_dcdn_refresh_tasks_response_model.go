// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnRefreshTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnRefreshTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnRefreshTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnRefreshTasksResponseBody) *DescribeDcdnRefreshTasksResponse
	GetBody() *DescribeDcdnRefreshTasksResponseBody
}

type DescribeDcdnRefreshTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnRefreshTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnRefreshTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRefreshTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnRefreshTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnRefreshTasksResponse) GetBody() *DescribeDcdnRefreshTasksResponseBody {
	return s.Body
}

func (s *DescribeDcdnRefreshTasksResponse) SetHeaders(v map[string]*string) *DescribeDcdnRefreshTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnRefreshTasksResponse) SetStatusCode(v int32) *DescribeDcdnRefreshTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponse) SetBody(v *DescribeDcdnRefreshTasksResponseBody) *DescribeDcdnRefreshTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnRefreshTasksResponse) Validate() error {
	return dara.Validate(s)
}
