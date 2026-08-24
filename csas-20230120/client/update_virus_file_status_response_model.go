// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVirusFileStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVirusFileStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVirusFileStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVirusFileStatusResponseBody) *UpdateVirusFileStatusResponse
	GetBody() *UpdateVirusFileStatusResponseBody
}

type UpdateVirusFileStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVirusFileStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVirusFileStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirusFileStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateVirusFileStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVirusFileStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVirusFileStatusResponse) GetBody() *UpdateVirusFileStatusResponseBody {
	return s.Body
}

func (s *UpdateVirusFileStatusResponse) SetHeaders(v map[string]*string) *UpdateVirusFileStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateVirusFileStatusResponse) SetStatusCode(v int32) *UpdateVirusFileStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVirusFileStatusResponse) SetBody(v *UpdateVirusFileStatusResponseBody) *UpdateVirusFileStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateVirusFileStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
