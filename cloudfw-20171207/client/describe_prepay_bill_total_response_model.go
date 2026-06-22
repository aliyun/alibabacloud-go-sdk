// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrepayBillTotalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePrepayBillTotalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePrepayBillTotalResponse
	GetStatusCode() *int32
	SetBody(v *DescribePrepayBillTotalResponseBody) *DescribePrepayBillTotalResponse
	GetBody() *DescribePrepayBillTotalResponseBody
}

type DescribePrepayBillTotalResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePrepayBillTotalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePrepayBillTotalResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePrepayBillTotalResponse) GoString() string {
	return s.String()
}

func (s *DescribePrepayBillTotalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePrepayBillTotalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePrepayBillTotalResponse) GetBody() *DescribePrepayBillTotalResponseBody {
	return s.Body
}

func (s *DescribePrepayBillTotalResponse) SetHeaders(v map[string]*string) *DescribePrepayBillTotalResponse {
	s.Headers = v
	return s
}

func (s *DescribePrepayBillTotalResponse) SetStatusCode(v int32) *DescribePrepayBillTotalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrepayBillTotalResponse) SetBody(v *DescribePrepayBillTotalResponseBody) *DescribePrepayBillTotalResponse {
	s.Body = v
	return s
}

func (s *DescribePrepayBillTotalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
