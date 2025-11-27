// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrateTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMigrateTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMigrateTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMigrateTasksResponseBody) *DescribeMigrateTasksResponse
	GetBody() *DescribeMigrateTasksResponseBody
}

type DescribeMigrateTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMigrateTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMigrateTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrateTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeMigrateTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMigrateTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMigrateTasksResponse) GetBody() *DescribeMigrateTasksResponseBody {
	return s.Body
}

func (s *DescribeMigrateTasksResponse) SetHeaders(v map[string]*string) *DescribeMigrateTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeMigrateTasksResponse) SetStatusCode(v int32) *DescribeMigrateTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMigrateTasksResponse) SetBody(v *DescribeMigrateTasksResponseBody) *DescribeMigrateTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeMigrateTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
