// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlLogRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSqlLogRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSqlLogRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSqlLogRecordsResponseBody) *DescribeSqlLogRecordsResponse
	GetBody() *DescribeSqlLogRecordsResponseBody
}

type DescribeSqlLogRecordsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSqlLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSqlLogRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSqlLogRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSqlLogRecordsResponse) GetBody() *DescribeSqlLogRecordsResponseBody {
	return s.Body
}

func (s *DescribeSqlLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeSqlLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlLogRecordsResponse) SetStatusCode(v int32) *DescribeSqlLogRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSqlLogRecordsResponse) SetBody(v *DescribeSqlLogRecordsResponseBody) *DescribeSqlLogRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeSqlLogRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
