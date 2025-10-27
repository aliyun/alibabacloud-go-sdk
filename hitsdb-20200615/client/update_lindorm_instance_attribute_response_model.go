// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLindormInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLindormInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLindormInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLindormInstanceAttributeResponseBody) *UpdateLindormInstanceAttributeResponse
	GetBody() *UpdateLindormInstanceAttributeResponseBody
}

type UpdateLindormInstanceAttributeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLindormInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLindormInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLindormInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateLindormInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLindormInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLindormInstanceAttributeResponse) GetBody() *UpdateLindormInstanceAttributeResponseBody {
	return s.Body
}

func (s *UpdateLindormInstanceAttributeResponse) SetHeaders(v map[string]*string) *UpdateLindormInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateLindormInstanceAttributeResponse) SetStatusCode(v int32) *UpdateLindormInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLindormInstanceAttributeResponse) SetBody(v *UpdateLindormInstanceAttributeResponseBody) *UpdateLindormInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateLindormInstanceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
