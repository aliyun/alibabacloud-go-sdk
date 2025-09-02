// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDIJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDIJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDIJobResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDIJobResponseBody) *UpdateDIJobResponse
	GetBody() *UpdateDIJobResponseBody
}

type UpdateDIJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDIJobResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateDIJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDIJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDIJobResponse) GetBody() *UpdateDIJobResponseBody {
	return s.Body
}

func (s *UpdateDIJobResponse) SetHeaders(v map[string]*string) *UpdateDIJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateDIJobResponse) SetStatusCode(v int32) *UpdateDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDIJobResponse) SetBody(v *UpdateDIJobResponseBody) *UpdateDIJobResponse {
	s.Body = v
	return s
}

func (s *UpdateDIJobResponse) Validate() error {
	return dara.Validate(s)
}
