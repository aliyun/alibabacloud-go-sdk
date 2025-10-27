// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLPatternsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSQLPatternsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSQLPatternsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSQLPatternsResponseBody) *DescribeSQLPatternsResponse
	GetBody() *DescribeSQLPatternsResponseBody
}

type DescribeSQLPatternsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQLPatternsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQLPatternsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLPatternsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSQLPatternsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSQLPatternsResponse) GetBody() *DescribeSQLPatternsResponseBody {
	return s.Body
}

func (s *DescribeSQLPatternsResponse) SetHeaders(v map[string]*string) *DescribeSQLPatternsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLPatternsResponse) SetStatusCode(v int32) *DescribeSQLPatternsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLPatternsResponse) SetBody(v *DescribeSQLPatternsResponseBody) *DescribeSQLPatternsResponse {
	s.Body = v
	return s
}

func (s *DescribeSQLPatternsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
