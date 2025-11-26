// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateClusterAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateClusterAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateClusterAttributeResponseBody) *UpdateClusterAttributeResponse
	GetBody() *UpdateClusterAttributeResponseBody
}

type UpdateClusterAttributeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClusterAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateClusterAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateClusterAttributeResponse) GetBody() *UpdateClusterAttributeResponseBody {
	return s.Body
}

func (s *UpdateClusterAttributeResponse) SetHeaders(v map[string]*string) *UpdateClusterAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterAttributeResponse) SetStatusCode(v int32) *UpdateClusterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClusterAttributeResponse) SetBody(v *UpdateClusterAttributeResponseBody) *UpdateClusterAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateClusterAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
