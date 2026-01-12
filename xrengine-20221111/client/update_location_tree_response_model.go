// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLocationTreeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLocationTreeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLocationTreeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLocationTreeResponseBody) *UpdateLocationTreeResponse
	GetBody() *UpdateLocationTreeResponseBody
}

type UpdateLocationTreeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLocationTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLocationTreeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocationTreeResponse) GoString() string {
	return s.String()
}

func (s *UpdateLocationTreeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLocationTreeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLocationTreeResponse) GetBody() *UpdateLocationTreeResponseBody {
	return s.Body
}

func (s *UpdateLocationTreeResponse) SetHeaders(v map[string]*string) *UpdateLocationTreeResponse {
	s.Headers = v
	return s
}

func (s *UpdateLocationTreeResponse) SetStatusCode(v int32) *UpdateLocationTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLocationTreeResponse) SetBody(v *UpdateLocationTreeResponseBody) *UpdateLocationTreeResponse {
	s.Body = v
	return s
}

func (s *UpdateLocationTreeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
