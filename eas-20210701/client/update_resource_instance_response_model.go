// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourceInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourceInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResourceInstanceResponseBody) *UpdateResourceInstanceResponse
	GetBody() *UpdateResourceInstanceResponseBody
}

type UpdateResourceInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourceInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourceInstanceResponse) GetBody() *UpdateResourceInstanceResponseBody {
	return s.Body
}

func (s *UpdateResourceInstanceResponse) SetHeaders(v map[string]*string) *UpdateResourceInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceInstanceResponse) SetStatusCode(v int32) *UpdateResourceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceInstanceResponse) SetBody(v *UpdateResourceInstanceResponseBody) *UpdateResourceInstanceResponse {
	s.Body = v
	return s
}

func (s *UpdateResourceInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
