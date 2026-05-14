// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBudgetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBudgetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBudgetResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBudgetResponseBody) *DescribeBudgetResponse
	GetBody() *DescribeBudgetResponseBody
}

type DescribeBudgetResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBudgetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBudgetResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetResponse) GoString() string {
	return s.String()
}

func (s *DescribeBudgetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBudgetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBudgetResponse) GetBody() *DescribeBudgetResponseBody {
	return s.Body
}

func (s *DescribeBudgetResponse) SetHeaders(v map[string]*string) *DescribeBudgetResponse {
	s.Headers = v
	return s
}

func (s *DescribeBudgetResponse) SetStatusCode(v int32) *DescribeBudgetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBudgetResponse) SetBody(v *DescribeBudgetResponseBody) *DescribeBudgetResponse {
	s.Body = v
	return s
}

func (s *DescribeBudgetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
