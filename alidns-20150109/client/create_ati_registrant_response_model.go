// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAtiRegistrantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAtiRegistrantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAtiRegistrantResponse
	GetStatusCode() *int32
	SetBody(v *CreateAtiRegistrantResponseBody) *CreateAtiRegistrantResponse
	GetBody() *CreateAtiRegistrantResponseBody
}

type CreateAtiRegistrantResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAtiRegistrantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAtiRegistrantResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAtiRegistrantResponse) GoString() string {
	return s.String()
}

func (s *CreateAtiRegistrantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAtiRegistrantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAtiRegistrantResponse) GetBody() *CreateAtiRegistrantResponseBody {
	return s.Body
}

func (s *CreateAtiRegistrantResponse) SetHeaders(v map[string]*string) *CreateAtiRegistrantResponse {
	s.Headers = v
	return s
}

func (s *CreateAtiRegistrantResponse) SetStatusCode(v int32) *CreateAtiRegistrantResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAtiRegistrantResponse) SetBody(v *CreateAtiRegistrantResponseBody) *CreateAtiRegistrantResponse {
	s.Body = v
	return s
}

func (s *CreateAtiRegistrantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
