// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePerformanceViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePerformanceViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePerformanceViewResponse
	GetStatusCode() *int32
	SetBody(v *CreatePerformanceViewResponseBody) *CreatePerformanceViewResponse
	GetBody() *CreatePerformanceViewResponseBody
}

type CreatePerformanceViewResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePerformanceViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePerformanceViewResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePerformanceViewResponse) GoString() string {
	return s.String()
}

func (s *CreatePerformanceViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePerformanceViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePerformanceViewResponse) GetBody() *CreatePerformanceViewResponseBody {
	return s.Body
}

func (s *CreatePerformanceViewResponse) SetHeaders(v map[string]*string) *CreatePerformanceViewResponse {
	s.Headers = v
	return s
}

func (s *CreatePerformanceViewResponse) SetStatusCode(v int32) *CreatePerformanceViewResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePerformanceViewResponse) SetBody(v *CreatePerformanceViewResponseBody) *CreatePerformanceViewResponse {
	s.Body = v
	return s
}

func (s *CreatePerformanceViewResponse) Validate() error {
	return dara.Validate(s)
}
