// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceShareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourceShareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourceShareResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResourceShareResponseBody) *UpdateResourceShareResponse
	GetBody() *UpdateResourceShareResponseBody
}

type UpdateResourceShareResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceShareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceShareResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceShareResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceShareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourceShareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourceShareResponse) GetBody() *UpdateResourceShareResponseBody {
	return s.Body
}

func (s *UpdateResourceShareResponse) SetHeaders(v map[string]*string) *UpdateResourceShareResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceShareResponse) SetStatusCode(v int32) *UpdateResourceShareResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceShareResponse) SetBody(v *UpdateResourceShareResponseBody) *UpdateResourceShareResponse {
	s.Body = v
	return s
}

func (s *UpdateResourceShareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
