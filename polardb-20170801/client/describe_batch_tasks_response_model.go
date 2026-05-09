// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBatchTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBatchTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBatchTasksResponseBody) *DescribeBatchTasksResponse
	GetBody() *DescribeBatchTasksResponseBody
}

type DescribeBatchTasksResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBatchTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBatchTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeBatchTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBatchTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBatchTasksResponse) GetBody() *DescribeBatchTasksResponseBody {
	return s.Body
}

func (s *DescribeBatchTasksResponse) SetHeaders(v map[string]*string) *DescribeBatchTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeBatchTasksResponse) SetStatusCode(v int32) *DescribeBatchTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBatchTasksResponse) SetBody(v *DescribeBatchTasksResponseBody) *DescribeBatchTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeBatchTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
