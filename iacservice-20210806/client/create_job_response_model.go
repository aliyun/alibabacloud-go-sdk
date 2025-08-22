// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateJobResponseBody) *CreateJobResponse
	GetBody() *CreateJobResponseBody
}

type CreateJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateJobResponse) GoString() string {
	return s.String()
}

func (s *CreateJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateJobResponse) GetBody() *CreateJobResponseBody {
	return s.Body
}

func (s *CreateJobResponse) SetHeaders(v map[string]*string) *CreateJobResponse {
	s.Headers = v
	return s
}

func (s *CreateJobResponse) SetStatusCode(v int32) *CreateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobResponse) SetBody(v *CreateJobResponseBody) *CreateJobResponse {
	s.Body = v
	return s
}

func (s *CreateJobResponse) Validate() error {
	return dara.Validate(s)
}
