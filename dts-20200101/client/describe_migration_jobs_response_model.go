// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrationJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMigrationJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMigrationJobsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMigrationJobsResponseBody) *DescribeMigrationJobsResponse
	GetBody() *DescribeMigrationJobsResponseBody
}

type DescribeMigrationJobsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMigrationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMigrationJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMigrationJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMigrationJobsResponse) GetBody() *DescribeMigrationJobsResponseBody {
	return s.Body
}

func (s *DescribeMigrationJobsResponse) SetHeaders(v map[string]*string) *DescribeMigrationJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMigrationJobsResponse) SetStatusCode(v int32) *DescribeMigrationJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMigrationJobsResponse) SetBody(v *DescribeMigrationJobsResponseBody) *DescribeMigrationJobsResponse {
	s.Body = v
	return s
}

func (s *DescribeMigrationJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
