// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBudgetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBudgetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBudgetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBudgetsResponseBody) *DescribeBudgetsResponse
	GetBody() *DescribeBudgetsResponseBody
}

type DescribeBudgetsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBudgetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBudgetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBudgetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBudgetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBudgetsResponse) GetBody() *DescribeBudgetsResponseBody {
	return s.Body
}

func (s *DescribeBudgetsResponse) SetHeaders(v map[string]*string) *DescribeBudgetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBudgetsResponse) SetStatusCode(v int32) *DescribeBudgetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBudgetsResponse) SetBody(v *DescribeBudgetsResponseBody) *DescribeBudgetsResponse {
	s.Body = v
	return s
}

func (s *DescribeBudgetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
