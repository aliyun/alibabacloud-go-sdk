// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVersionDescResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVersionDescResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVersionDescResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVersionDescResponseBody) *UpdateVersionDescResponse
	GetBody() *UpdateVersionDescResponseBody
}

type UpdateVersionDescResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVersionDescResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVersionDescResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVersionDescResponse) GoString() string {
	return s.String()
}

func (s *UpdateVersionDescResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVersionDescResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVersionDescResponse) GetBody() *UpdateVersionDescResponseBody {
	return s.Body
}

func (s *UpdateVersionDescResponse) SetHeaders(v map[string]*string) *UpdateVersionDescResponse {
	s.Headers = v
	return s
}

func (s *UpdateVersionDescResponse) SetStatusCode(v int32) *UpdateVersionDescResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVersionDescResponse) SetBody(v *UpdateVersionDescResponseBody) *UpdateVersionDescResponse {
	s.Body = v
	return s
}

func (s *UpdateVersionDescResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
