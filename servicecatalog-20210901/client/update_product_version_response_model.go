// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProductVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProductVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProductVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProductVersionResponseBody) *UpdateProductVersionResponse
	GetBody() *UpdateProductVersionResponseBody
}

type UpdateProductVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProductVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProductVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProductVersionResponse) GetBody() *UpdateProductVersionResponseBody {
	return s.Body
}

func (s *UpdateProductVersionResponse) SetHeaders(v map[string]*string) *UpdateProductVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateProductVersionResponse) SetStatusCode(v int32) *UpdateProductVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProductVersionResponse) SetBody(v *UpdateProductVersionResponseBody) *UpdateProductVersionResponse {
	s.Body = v
	return s
}

func (s *UpdateProductVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
