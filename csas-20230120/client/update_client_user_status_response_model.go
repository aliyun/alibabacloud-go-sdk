// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientUserStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateClientUserStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateClientUserStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateClientUserStatusResponseBody) *UpdateClientUserStatusResponse
	GetBody() *UpdateClientUserStatusResponseBody
}

type UpdateClientUserStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClientUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClientUserStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientUserStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateClientUserStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateClientUserStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateClientUserStatusResponse) GetBody() *UpdateClientUserStatusResponseBody {
	return s.Body
}

func (s *UpdateClientUserStatusResponse) SetHeaders(v map[string]*string) *UpdateClientUserStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateClientUserStatusResponse) SetStatusCode(v int32) *UpdateClientUserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClientUserStatusResponse) SetBody(v *UpdateClientUserStatusResponseBody) *UpdateClientUserStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateClientUserStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
