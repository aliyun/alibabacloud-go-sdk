// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCalculationJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCalculationJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCalculationJobsResponse
	GetStatusCode() *int32
	SetBody(v *CreateCalculationJobsResponseBody) *CreateCalculationJobsResponse
	GetBody() *CreateCalculationJobsResponseBody
}

type CreateCalculationJobsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCalculationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCalculationJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCalculationJobsResponse) GoString() string {
	return s.String()
}

func (s *CreateCalculationJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCalculationJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCalculationJobsResponse) GetBody() *CreateCalculationJobsResponseBody {
	return s.Body
}

func (s *CreateCalculationJobsResponse) SetHeaders(v map[string]*string) *CreateCalculationJobsResponse {
	s.Headers = v
	return s
}

func (s *CreateCalculationJobsResponse) SetStatusCode(v int32) *CreateCalculationJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCalculationJobsResponse) SetBody(v *CreateCalculationJobsResponseBody) *CreateCalculationJobsResponse {
	s.Body = v
	return s
}

func (s *CreateCalculationJobsResponse) Validate() error {
	return dara.Validate(s)
}
