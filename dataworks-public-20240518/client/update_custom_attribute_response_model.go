// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomAttributeResponseBody) *UpdateCustomAttributeResponse
	GetBody() *UpdateCustomAttributeResponseBody
}

type UpdateCustomAttributeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomAttributeResponse) GetBody() *UpdateCustomAttributeResponseBody {
	return s.Body
}

func (s *UpdateCustomAttributeResponse) SetHeaders(v map[string]*string) *UpdateCustomAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomAttributeResponse) SetStatusCode(v int32) *UpdateCustomAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomAttributeResponse) SetBody(v *UpdateCustomAttributeResponseBody) *UpdateCustomAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
