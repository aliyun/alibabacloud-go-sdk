// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBruteForceRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBruteForceRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBruteForceRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBruteForceRecordsResponseBody) *DescribeBruteForceRecordsResponse
	GetBody() *DescribeBruteForceRecordsResponseBody
}

type DescribeBruteForceRecordsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBruteForceRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBruteForceRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBruteForceRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBruteForceRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBruteForceRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBruteForceRecordsResponse) GetBody() *DescribeBruteForceRecordsResponseBody {
	return s.Body
}

func (s *DescribeBruteForceRecordsResponse) SetHeaders(v map[string]*string) *DescribeBruteForceRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBruteForceRecordsResponse) SetStatusCode(v int32) *DescribeBruteForceRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBruteForceRecordsResponse) SetBody(v *DescribeBruteForceRecordsResponseBody) *DescribeBruteForceRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeBruteForceRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
