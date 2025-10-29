// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComponentIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateComponentIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateComponentIndexResponse
	GetStatusCode() *int32
	SetBody(v *UpdateComponentIndexResponseBody) *UpdateComponentIndexResponse
	GetBody() *UpdateComponentIndexResponseBody
}

type UpdateComponentIndexResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComponentIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComponentIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateComponentIndexResponse) GoString() string {
	return s.String()
}

func (s *UpdateComponentIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateComponentIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateComponentIndexResponse) GetBody() *UpdateComponentIndexResponseBody {
	return s.Body
}

func (s *UpdateComponentIndexResponse) SetHeaders(v map[string]*string) *UpdateComponentIndexResponse {
	s.Headers = v
	return s
}

func (s *UpdateComponentIndexResponse) SetStatusCode(v int32) *UpdateComponentIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComponentIndexResponse) SetBody(v *UpdateComponentIndexResponseBody) *UpdateComponentIndexResponse {
	s.Body = v
	return s
}

func (s *UpdateComponentIndexResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
