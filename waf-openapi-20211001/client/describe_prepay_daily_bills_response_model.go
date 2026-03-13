// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrepayDailyBillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePrepayDailyBillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePrepayDailyBillsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePrepayDailyBillsResponseBody) *DescribePrepayDailyBillsResponse
	GetBody() *DescribePrepayDailyBillsResponseBody
}

type DescribePrepayDailyBillsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePrepayDailyBillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePrepayDailyBillsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePrepayDailyBillsResponse) GoString() string {
	return s.String()
}

func (s *DescribePrepayDailyBillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePrepayDailyBillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePrepayDailyBillsResponse) GetBody() *DescribePrepayDailyBillsResponseBody {
	return s.Body
}

func (s *DescribePrepayDailyBillsResponse) SetHeaders(v map[string]*string) *DescribePrepayDailyBillsResponse {
	s.Headers = v
	return s
}

func (s *DescribePrepayDailyBillsResponse) SetStatusCode(v int32) *DescribePrepayDailyBillsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrepayDailyBillsResponse) SetBody(v *DescribePrepayDailyBillsResponseBody) *DescribePrepayDailyBillsResponse {
	s.Body = v
	return s
}

func (s *DescribePrepayDailyBillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
