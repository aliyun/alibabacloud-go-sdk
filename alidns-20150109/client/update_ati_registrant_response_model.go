// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAtiRegistrantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAtiRegistrantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAtiRegistrantResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAtiRegistrantResponseBody) *UpdateAtiRegistrantResponse
	GetBody() *UpdateAtiRegistrantResponseBody
}

type UpdateAtiRegistrantResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAtiRegistrantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAtiRegistrantResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAtiRegistrantResponse) GoString() string {
	return s.String()
}

func (s *UpdateAtiRegistrantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAtiRegistrantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAtiRegistrantResponse) GetBody() *UpdateAtiRegistrantResponseBody {
	return s.Body
}

func (s *UpdateAtiRegistrantResponse) SetHeaders(v map[string]*string) *UpdateAtiRegistrantResponse {
	s.Headers = v
	return s
}

func (s *UpdateAtiRegistrantResponse) SetStatusCode(v int32) *UpdateAtiRegistrantResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAtiRegistrantResponse) SetBody(v *UpdateAtiRegistrantResponseBody) *UpdateAtiRegistrantResponse {
	s.Body = v
	return s
}

func (s *UpdateAtiRegistrantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
