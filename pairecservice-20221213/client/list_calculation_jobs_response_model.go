// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalculationJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCalculationJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCalculationJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListCalculationJobsResponseBody) *ListCalculationJobsResponse
	GetBody() *ListCalculationJobsResponseBody
}

type ListCalculationJobsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCalculationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCalculationJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCalculationJobsResponse) GoString() string {
	return s.String()
}

func (s *ListCalculationJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCalculationJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCalculationJobsResponse) GetBody() *ListCalculationJobsResponseBody {
	return s.Body
}

func (s *ListCalculationJobsResponse) SetHeaders(v map[string]*string) *ListCalculationJobsResponse {
	s.Headers = v
	return s
}

func (s *ListCalculationJobsResponse) SetStatusCode(v int32) *ListCalculationJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCalculationJobsResponse) SetBody(v *ListCalculationJobsResponseBody) *ListCalculationJobsResponse {
	s.Body = v
	return s
}

func (s *ListCalculationJobsResponse) Validate() error {
	return dara.Validate(s)
}
