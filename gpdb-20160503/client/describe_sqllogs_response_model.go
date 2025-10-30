// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSQLLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSQLLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSQLLogsResponseBody) *DescribeSQLLogsResponse
	GetBody() *DescribeSQLLogsResponseBody
}

type DescribeSQLLogsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQLLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQLLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSQLLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSQLLogsResponse) GetBody() *DescribeSQLLogsResponseBody {
	return s.Body
}

func (s *DescribeSQLLogsResponse) SetHeaders(v map[string]*string) *DescribeSQLLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogsResponse) SetStatusCode(v int32) *DescribeSQLLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLLogsResponse) SetBody(v *DescribeSQLLogsResponseBody) *DescribeSQLLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeSQLLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
