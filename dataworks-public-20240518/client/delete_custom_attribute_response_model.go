// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomAttributeResponseBody) *DeleteCustomAttributeResponse
	GetBody() *DeleteCustomAttributeResponseBody
}

type DeleteCustomAttributeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAttributeResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomAttributeResponse) GetBody() *DeleteCustomAttributeResponseBody {
	return s.Body
}

func (s *DeleteCustomAttributeResponse) SetHeaders(v map[string]*string) *DeleteCustomAttributeResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomAttributeResponse) SetStatusCode(v int32) *DeleteCustomAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomAttributeResponse) SetBody(v *DeleteCustomAttributeResponseBody) *DeleteCustomAttributeResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
