// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAtiRegistrantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAtiRegistrantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAtiRegistrantResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAtiRegistrantResponseBody) *DeleteAtiRegistrantResponse
	GetBody() *DeleteAtiRegistrantResponseBody
}

type DeleteAtiRegistrantResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAtiRegistrantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAtiRegistrantResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAtiRegistrantResponse) GoString() string {
	return s.String()
}

func (s *DeleteAtiRegistrantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAtiRegistrantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAtiRegistrantResponse) GetBody() *DeleteAtiRegistrantResponseBody {
	return s.Body
}

func (s *DeleteAtiRegistrantResponse) SetHeaders(v map[string]*string) *DeleteAtiRegistrantResponse {
	s.Headers = v
	return s
}

func (s *DeleteAtiRegistrantResponse) SetStatusCode(v int32) *DeleteAtiRegistrantResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAtiRegistrantResponse) SetBody(v *DeleteAtiRegistrantResponseBody) *DeleteAtiRegistrantResponse {
	s.Body = v
	return s
}

func (s *DeleteAtiRegistrantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
