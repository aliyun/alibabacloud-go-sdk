// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRumFileStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRumFileStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRumFileStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRumFileStatusResponseBody) *UpdateRumFileStatusResponse
	GetBody() *UpdateRumFileStatusResponseBody
}

type UpdateRumFileStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRumFileStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRumFileStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRumFileStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateRumFileStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRumFileStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRumFileStatusResponse) GetBody() *UpdateRumFileStatusResponseBody {
	return s.Body
}

func (s *UpdateRumFileStatusResponse) SetHeaders(v map[string]*string) *UpdateRumFileStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateRumFileStatusResponse) SetStatusCode(v int32) *UpdateRumFileStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRumFileStatusResponse) SetBody(v *UpdateRumFileStatusResponseBody) *UpdateRumFileStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateRumFileStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
