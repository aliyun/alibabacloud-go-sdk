// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsAccountSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePdnsAccountSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePdnsAccountSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribePdnsAccountSummaryResponseBody) *DescribePdnsAccountSummaryResponse
	GetBody() *DescribePdnsAccountSummaryResponseBody
}

type DescribePdnsAccountSummaryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePdnsAccountSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePdnsAccountSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsAccountSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribePdnsAccountSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePdnsAccountSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePdnsAccountSummaryResponse) GetBody() *DescribePdnsAccountSummaryResponseBody {
	return s.Body
}

func (s *DescribePdnsAccountSummaryResponse) SetHeaders(v map[string]*string) *DescribePdnsAccountSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribePdnsAccountSummaryResponse) SetStatusCode(v int32) *DescribePdnsAccountSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePdnsAccountSummaryResponse) SetBody(v *DescribePdnsAccountSummaryResponseBody) *DescribePdnsAccountSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribePdnsAccountSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
