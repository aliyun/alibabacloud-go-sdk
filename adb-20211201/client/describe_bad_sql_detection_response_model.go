// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBadSqlDetectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBadSqlDetectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBadSqlDetectionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBadSqlDetectionResponseBody) *DescribeBadSqlDetectionResponse
	GetBody() *DescribeBadSqlDetectionResponseBody
}

type DescribeBadSqlDetectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBadSqlDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBadSqlDetectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBadSqlDetectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeBadSqlDetectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBadSqlDetectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBadSqlDetectionResponse) GetBody() *DescribeBadSqlDetectionResponseBody {
	return s.Body
}

func (s *DescribeBadSqlDetectionResponse) SetHeaders(v map[string]*string) *DescribeBadSqlDetectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeBadSqlDetectionResponse) SetStatusCode(v int32) *DescribeBadSqlDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBadSqlDetectionResponse) SetBody(v *DescribeBadSqlDetectionResponseBody) *DescribeBadSqlDetectionResponse {
	s.Body = v
	return s
}

func (s *DescribeBadSqlDetectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
