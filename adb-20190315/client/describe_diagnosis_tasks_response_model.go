// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiagnosisTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiagnosisTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiagnosisTasksResponseBody) *DescribeDiagnosisTasksResponse
	GetBody() *DescribeDiagnosisTasksResponseBody
}

type DescribeDiagnosisTasksResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosisTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosisTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiagnosisTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiagnosisTasksResponse) GetBody() *DescribeDiagnosisTasksResponseBody {
	return s.Body
}

func (s *DescribeDiagnosisTasksResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisTasksResponse) SetStatusCode(v int32) *DescribeDiagnosisTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisTasksResponse) SetBody(v *DescribeDiagnosisTasksResponseBody) *DescribeDiagnosisTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeDiagnosisTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
