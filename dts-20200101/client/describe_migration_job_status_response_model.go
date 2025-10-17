// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrationJobStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMigrationJobStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMigrationJobStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMigrationJobStatusResponseBody) *DescribeMigrationJobStatusResponse
	GetBody() *DescribeMigrationJobStatusResponseBody
}

type DescribeMigrationJobStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMigrationJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMigrationJobStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMigrationJobStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMigrationJobStatusResponse) GetBody() *DescribeMigrationJobStatusResponseBody {
	return s.Body
}

func (s *DescribeMigrationJobStatusResponse) SetHeaders(v map[string]*string) *DescribeMigrationJobStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeMigrationJobStatusResponse) SetStatusCode(v int32) *DescribeMigrationJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMigrationJobStatusResponse) SetBody(v *DescribeMigrationJobStatusResponseBody) *DescribeMigrationJobStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeMigrationJobStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
