// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVccResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVccResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVccResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVccResponseBody) *UpdateVccResponse
	GetBody() *UpdateVccResponseBody
}

type UpdateVccResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVccResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVccResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVccResponse) GoString() string {
	return s.String()
}

func (s *UpdateVccResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVccResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVccResponse) GetBody() *UpdateVccResponseBody {
	return s.Body
}

func (s *UpdateVccResponse) SetHeaders(v map[string]*string) *UpdateVccResponse {
	s.Headers = v
	return s
}

func (s *UpdateVccResponse) SetStatusCode(v int32) *UpdateVccResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVccResponse) SetBody(v *UpdateVccResponseBody) *UpdateVccResponse {
	s.Body = v
	return s
}

func (s *UpdateVccResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
