// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstanceLevelTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsInstanceLevelTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsInstanceLevelTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsInstanceLevelTasksResponseBody) *DescribeDrdsInstanceLevelTasksResponse
	GetBody() *DescribeDrdsInstanceLevelTasksResponseBody
}

type DescribeDrdsInstanceLevelTasksResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsInstanceLevelTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsInstanceLevelTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsInstanceLevelTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsInstanceLevelTasksResponse) GetBody() *DescribeDrdsInstanceLevelTasksResponseBody {
	return s.Body
}

func (s *DescribeDrdsInstanceLevelTasksResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstanceLevelTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponse) SetStatusCode(v int32) *DescribeDrdsInstanceLevelTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponse) SetBody(v *DescribeDrdsInstanceLevelTasksResponseBody) *DescribeDrdsInstanceLevelTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponse) Validate() error {
	return dara.Validate(s)
}
