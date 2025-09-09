// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsTasksResponseBody) *DescribeDrdsTasksResponse
	GetBody() *DescribeDrdsTasksResponseBody
}

type DescribeDrdsTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsTasksResponse) GetBody() *DescribeDrdsTasksResponseBody {
	return s.Body
}

func (s *DescribeDrdsTasksResponse) SetHeaders(v map[string]*string) *DescribeDrdsTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsTasksResponse) SetStatusCode(v int32) *DescribeDrdsTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsTasksResponse) SetBody(v *DescribeDrdsTasksResponseBody) *DescribeDrdsTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsTasksResponse) Validate() error {
	return dara.Validate(s)
}
