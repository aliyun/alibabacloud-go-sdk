// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumBillsByDateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SumBillsByDateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SumBillsByDateResponse
	GetStatusCode() *int32
	SetBody(v *SumBillsByDateResponseBody) *SumBillsByDateResponse
	GetBody() *SumBillsByDateResponseBody
}

type SumBillsByDateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SumBillsByDateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SumBillsByDateResponse) String() string {
	return dara.Prettify(s)
}

func (s SumBillsByDateResponse) GoString() string {
	return s.String()
}

func (s *SumBillsByDateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SumBillsByDateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SumBillsByDateResponse) GetBody() *SumBillsByDateResponseBody {
	return s.Body
}

func (s *SumBillsByDateResponse) SetHeaders(v map[string]*string) *SumBillsByDateResponse {
	s.Headers = v
	return s
}

func (s *SumBillsByDateResponse) SetStatusCode(v int32) *SumBillsByDateResponse {
	s.StatusCode = &v
	return s
}

func (s *SumBillsByDateResponse) SetBody(v *SumBillsByDateResponseBody) *SumBillsByDateResponse {
	s.Body = v
	return s
}

func (s *SumBillsByDateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
