// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkerResponseBody) *UpdateWorkerResponse
	GetBody() *UpdateWorkerResponseBody
}

type UpdateWorkerResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkerResponse) GetBody() *UpdateWorkerResponseBody {
	return s.Body
}

func (s *UpdateWorkerResponse) SetHeaders(v map[string]*string) *UpdateWorkerResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkerResponse) SetStatusCode(v int32) *UpdateWorkerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkerResponse) SetBody(v *UpdateWorkerResponseBody) *UpdateWorkerResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
