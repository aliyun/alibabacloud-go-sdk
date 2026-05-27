// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumBillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SumBillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SumBillsResponse
	GetStatusCode() *int32
	SetBody(v *SumBillsResponseBody) *SumBillsResponse
	GetBody() *SumBillsResponseBody
}

type SumBillsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SumBillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SumBillsResponse) String() string {
	return dara.Prettify(s)
}

func (s SumBillsResponse) GoString() string {
	return s.String()
}

func (s *SumBillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SumBillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SumBillsResponse) GetBody() *SumBillsResponseBody {
	return s.Body
}

func (s *SumBillsResponse) SetHeaders(v map[string]*string) *SumBillsResponse {
	s.Headers = v
	return s
}

func (s *SumBillsResponse) SetStatusCode(v int32) *SumBillsResponse {
	s.StatusCode = &v
	return s
}

func (s *SumBillsResponse) SetBody(v *SumBillsResponseBody) *SumBillsResponse {
	s.Body = v
	return s
}

func (s *SumBillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
