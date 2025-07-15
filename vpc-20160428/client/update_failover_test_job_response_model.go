// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFailoverTestJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFailoverTestJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFailoverTestJobResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFailoverTestJobResponseBody) *UpdateFailoverTestJobResponse
	GetBody() *UpdateFailoverTestJobResponseBody
}

type UpdateFailoverTestJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFailoverTestJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFailoverTestJobResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFailoverTestJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateFailoverTestJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFailoverTestJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFailoverTestJobResponse) GetBody() *UpdateFailoverTestJobResponseBody {
	return s.Body
}

func (s *UpdateFailoverTestJobResponse) SetHeaders(v map[string]*string) *UpdateFailoverTestJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateFailoverTestJobResponse) SetStatusCode(v int32) *UpdateFailoverTestJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFailoverTestJobResponse) SetBody(v *UpdateFailoverTestJobResponseBody) *UpdateFailoverTestJobResponse {
	s.Body = v
	return s
}

func (s *UpdateFailoverTestJobResponse) Validate() error {
	return dara.Validate(s)
}
