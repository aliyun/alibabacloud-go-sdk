// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkerResourceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkerResourceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkerResourceStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkerResourceStatusResponseBody) *UpdateWorkerResourceStatusResponse
	GetBody() *UpdateWorkerResourceStatusResponseBody
}

type UpdateWorkerResourceStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkerResourceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkerResourceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerResourceStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResourceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkerResourceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkerResourceStatusResponse) GetBody() *UpdateWorkerResourceStatusResponseBody {
	return s.Body
}

func (s *UpdateWorkerResourceStatusResponse) SetHeaders(v map[string]*string) *UpdateWorkerResourceStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkerResourceStatusResponse) SetStatusCode(v int32) *UpdateWorkerResourceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkerResourceStatusResponse) SetBody(v *UpdateWorkerResourceStatusResponseBody) *UpdateWorkerResourceStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkerResourceStatusResponse) Validate() error {
	return dara.Validate(s)
}
