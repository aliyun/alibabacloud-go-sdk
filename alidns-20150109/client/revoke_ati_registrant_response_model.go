// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAtiRegistrantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeAtiRegistrantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeAtiRegistrantResponse
	GetStatusCode() *int32
	SetBody(v *RevokeAtiRegistrantResponseBody) *RevokeAtiRegistrantResponse
	GetBody() *RevokeAtiRegistrantResponseBody
}

type RevokeAtiRegistrantResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeAtiRegistrantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeAtiRegistrantResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeAtiRegistrantResponse) GoString() string {
	return s.String()
}

func (s *RevokeAtiRegistrantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeAtiRegistrantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeAtiRegistrantResponse) GetBody() *RevokeAtiRegistrantResponseBody {
	return s.Body
}

func (s *RevokeAtiRegistrantResponse) SetHeaders(v map[string]*string) *RevokeAtiRegistrantResponse {
	s.Headers = v
	return s
}

func (s *RevokeAtiRegistrantResponse) SetStatusCode(v int32) *RevokeAtiRegistrantResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeAtiRegistrantResponse) SetBody(v *RevokeAtiRegistrantResponseBody) *RevokeAtiRegistrantResponse {
	s.Body = v
	return s
}

func (s *RevokeAtiRegistrantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
