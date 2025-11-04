// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnoseReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiagnoseReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiagnoseReportsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiagnoseReportsResponseBody) *DescribeDiagnoseReportsResponse
	GetBody() *DescribeDiagnoseReportsResponseBody
}

type DescribeDiagnoseReportsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnoseReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnoseReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnoseReportsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiagnoseReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiagnoseReportsResponse) GetBody() *DescribeDiagnoseReportsResponseBody {
	return s.Body
}

func (s *DescribeDiagnoseReportsResponse) SetHeaders(v map[string]*string) *DescribeDiagnoseReportsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnoseReportsResponse) SetStatusCode(v int32) *DescribeDiagnoseReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnoseReportsResponse) SetBody(v *DescribeDiagnoseReportsResponseBody) *DescribeDiagnoseReportsResponse {
	s.Body = v
	return s
}

func (s *DescribeDiagnoseReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
