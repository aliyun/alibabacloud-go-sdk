// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceAttributeResponseBody) *UpdateInstanceAttributeResponse
	GetBody() *UpdateInstanceAttributeResponseBody
}

type UpdateInstanceAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceAttributeResponse) GetBody() *UpdateInstanceAttributeResponseBody {
	return s.Body
}

func (s *UpdateInstanceAttributeResponse) SetHeaders(v map[string]*string) *UpdateInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceAttributeResponse) SetStatusCode(v int32) *UpdateInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceAttributeResponse) SetBody(v *UpdateInstanceAttributeResponseBody) *UpdateInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
