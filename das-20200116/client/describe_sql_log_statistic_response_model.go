// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlLogStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSqlLogStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSqlLogStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSqlLogStatisticResponseBody) *DescribeSqlLogStatisticResponse
	GetBody() *DescribeSqlLogStatisticResponseBody
}

type DescribeSqlLogStatisticResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSqlLogStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSqlLogStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSqlLogStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSqlLogStatisticResponse) GetBody() *DescribeSqlLogStatisticResponseBody {
	return s.Body
}

func (s *DescribeSqlLogStatisticResponse) SetHeaders(v map[string]*string) *DescribeSqlLogStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlLogStatisticResponse) SetStatusCode(v int32) *DescribeSqlLogStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSqlLogStatisticResponse) SetBody(v *DescribeSqlLogStatisticResponseBody) *DescribeSqlLogStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeSqlLogStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
