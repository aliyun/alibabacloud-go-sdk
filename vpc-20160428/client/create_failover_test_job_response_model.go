// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFailoverTestJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFailoverTestJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFailoverTestJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateFailoverTestJobResponseBody) *CreateFailoverTestJobResponse
	GetBody() *CreateFailoverTestJobResponseBody
}

type CreateFailoverTestJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFailoverTestJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFailoverTestJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFailoverTestJobResponse) GoString() string {
	return s.String()
}

func (s *CreateFailoverTestJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFailoverTestJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFailoverTestJobResponse) GetBody() *CreateFailoverTestJobResponseBody {
	return s.Body
}

func (s *CreateFailoverTestJobResponse) SetHeaders(v map[string]*string) *CreateFailoverTestJobResponse {
	s.Headers = v
	return s
}

func (s *CreateFailoverTestJobResponse) SetStatusCode(v int32) *CreateFailoverTestJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFailoverTestJobResponse) SetBody(v *CreateFailoverTestJobResponseBody) *CreateFailoverTestJobResponse {
	s.Body = v
	return s
}

func (s *CreateFailoverTestJobResponse) Validate() error {
	return dara.Validate(s)
}
