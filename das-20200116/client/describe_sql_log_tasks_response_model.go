// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlLogTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSqlLogTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSqlLogTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSqlLogTasksResponseBody) *DescribeSqlLogTasksResponse
	GetBody() *DescribeSqlLogTasksResponseBody
}

type DescribeSqlLogTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSqlLogTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSqlLogTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSqlLogTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSqlLogTasksResponse) GetBody() *DescribeSqlLogTasksResponseBody {
	return s.Body
}

func (s *DescribeSqlLogTasksResponse) SetHeaders(v map[string]*string) *DescribeSqlLogTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlLogTasksResponse) SetStatusCode(v int32) *DescribeSqlLogTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSqlLogTasksResponse) SetBody(v *DescribeSqlLogTasksResponseBody) *DescribeSqlLogTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeSqlLogTasksResponse) Validate() error {
	return dara.Validate(s)
}
