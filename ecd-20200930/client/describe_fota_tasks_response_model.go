// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFotaTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFotaTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFotaTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFotaTasksResponseBody) *DescribeFotaTasksResponse
	GetBody() *DescribeFotaTasksResponseBody
}

type DescribeFotaTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFotaTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFotaTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFotaTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeFotaTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFotaTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFotaTasksResponse) GetBody() *DescribeFotaTasksResponseBody {
	return s.Body
}

func (s *DescribeFotaTasksResponse) SetHeaders(v map[string]*string) *DescribeFotaTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeFotaTasksResponse) SetStatusCode(v int32) *DescribeFotaTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFotaTasksResponse) SetBody(v *DescribeFotaTasksResponseBody) *DescribeFotaTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeFotaTasksResponse) Validate() error {
	return dara.Validate(s)
}
