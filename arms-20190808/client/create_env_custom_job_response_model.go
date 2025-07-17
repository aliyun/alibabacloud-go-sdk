// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnvCustomJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEnvCustomJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEnvCustomJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateEnvCustomJobResponseBody) *CreateEnvCustomJobResponse
	GetBody() *CreateEnvCustomJobResponseBody
}

type CreateEnvCustomJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEnvCustomJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnvCustomJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvCustomJobResponse) GoString() string {
	return s.String()
}

func (s *CreateEnvCustomJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEnvCustomJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEnvCustomJobResponse) GetBody() *CreateEnvCustomJobResponseBody {
	return s.Body
}

func (s *CreateEnvCustomJobResponse) SetHeaders(v map[string]*string) *CreateEnvCustomJobResponse {
	s.Headers = v
	return s
}

func (s *CreateEnvCustomJobResponse) SetStatusCode(v int32) *CreateEnvCustomJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnvCustomJobResponse) SetBody(v *CreateEnvCustomJobResponseBody) *CreateEnvCustomJobResponse {
	s.Body = v
	return s
}

func (s *CreateEnvCustomJobResponse) Validate() error {
	return dara.Validate(s)
}
