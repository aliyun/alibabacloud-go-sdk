// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCommandAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCommandAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCommandAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCommandAttributeResponseBody) *UpdateCommandAttributeResponse
	GetBody() *UpdateCommandAttributeResponseBody
}

type UpdateCommandAttributeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCommandAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCommandAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCommandAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateCommandAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCommandAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCommandAttributeResponse) GetBody() *UpdateCommandAttributeResponseBody {
	return s.Body
}

func (s *UpdateCommandAttributeResponse) SetHeaders(v map[string]*string) *UpdateCommandAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateCommandAttributeResponse) SetStatusCode(v int32) *UpdateCommandAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCommandAttributeResponse) SetBody(v *UpdateCommandAttributeResponseBody) *UpdateCommandAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateCommandAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
