// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlPatternResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSqlPatternResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSqlPatternResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSqlPatternResponseBody) *DescribeSqlPatternResponse
	GetBody() *DescribeSqlPatternResponseBody
}

type DescribeSqlPatternResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSqlPatternResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSqlPatternResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlPatternResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSqlPatternResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSqlPatternResponse) GetBody() *DescribeSqlPatternResponseBody {
	return s.Body
}

func (s *DescribeSqlPatternResponse) SetHeaders(v map[string]*string) *DescribeSqlPatternResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlPatternResponse) SetStatusCode(v int32) *DescribeSqlPatternResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSqlPatternResponse) SetBody(v *DescribeSqlPatternResponseBody) *DescribeSqlPatternResponse {
	s.Body = v
	return s
}

func (s *DescribeSqlPatternResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
