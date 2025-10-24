// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePeakTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePeakTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePeakTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribePeakTrendResponseBody) *DescribePeakTrendResponse
	GetBody() *DescribePeakTrendResponseBody
}

type DescribePeakTrendResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePeakTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePeakTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePeakTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribePeakTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePeakTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePeakTrendResponse) GetBody() *DescribePeakTrendResponseBody {
	return s.Body
}

func (s *DescribePeakTrendResponse) SetHeaders(v map[string]*string) *DescribePeakTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribePeakTrendResponse) SetStatusCode(v int32) *DescribePeakTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePeakTrendResponse) SetBody(v *DescribePeakTrendResponseBody) *DescribePeakTrendResponse {
	s.Body = v
	return s
}

func (s *DescribePeakTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
