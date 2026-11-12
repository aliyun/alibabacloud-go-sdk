// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseSlowLogRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDatabaseSlowLogRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDatabaseSlowLogRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDatabaseSlowLogRecordsResponseBody) *DescribeDatabaseSlowLogRecordsResponse
	GetBody() *DescribeDatabaseSlowLogRecordsResponseBody
}

type DescribeDatabaseSlowLogRecordsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDatabaseSlowLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDatabaseSlowLogRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseSlowLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseSlowLogRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDatabaseSlowLogRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDatabaseSlowLogRecordsResponse) GetBody() *DescribeDatabaseSlowLogRecordsResponseBody {
	return s.Body
}

func (s *DescribeDatabaseSlowLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeDatabaseSlowLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponse) SetStatusCode(v int32) *DescribeDatabaseSlowLogRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponse) SetBody(v *DescribeDatabaseSlowLogRecordsResponseBody) *DescribeDatabaseSlowLogRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
