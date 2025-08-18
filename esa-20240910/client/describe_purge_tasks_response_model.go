// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurgeTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePurgeTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePurgeTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribePurgeTasksResponseBody) *DescribePurgeTasksResponse
	GetBody() *DescribePurgeTasksResponseBody
}

type DescribePurgeTasksResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePurgeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePurgeTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePurgeTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribePurgeTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePurgeTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePurgeTasksResponse) GetBody() *DescribePurgeTasksResponseBody {
	return s.Body
}

func (s *DescribePurgeTasksResponse) SetHeaders(v map[string]*string) *DescribePurgeTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribePurgeTasksResponse) SetStatusCode(v int32) *DescribePurgeTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePurgeTasksResponse) SetBody(v *DescribePurgeTasksResponseBody) *DescribePurgeTasksResponse {
	s.Body = v
	return s
}

func (s *DescribePurgeTasksResponse) Validate() error {
	return dara.Validate(s)
}
