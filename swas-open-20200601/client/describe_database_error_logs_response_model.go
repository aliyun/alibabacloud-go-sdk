// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseErrorLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDatabaseErrorLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDatabaseErrorLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDatabaseErrorLogsResponseBody) *DescribeDatabaseErrorLogsResponse
	GetBody() *DescribeDatabaseErrorLogsResponseBody
}

type DescribeDatabaseErrorLogsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDatabaseErrorLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDatabaseErrorLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseErrorLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseErrorLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDatabaseErrorLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDatabaseErrorLogsResponse) GetBody() *DescribeDatabaseErrorLogsResponseBody {
	return s.Body
}

func (s *DescribeDatabaseErrorLogsResponse) SetHeaders(v map[string]*string) *DescribeDatabaseErrorLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabaseErrorLogsResponse) SetStatusCode(v int32) *DescribeDatabaseErrorLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponse) SetBody(v *DescribeDatabaseErrorLogsResponseBody) *DescribeDatabaseErrorLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeDatabaseErrorLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
