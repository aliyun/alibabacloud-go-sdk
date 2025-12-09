// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationSsoFormParamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationSsoFormParamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationSsoFormParamsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationSsoFormParamsResponseBody) *UpdateApplicationSsoFormParamsResponse
	GetBody() *UpdateApplicationSsoFormParamsResponseBody
}

type UpdateApplicationSsoFormParamsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationSsoFormParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationSsoFormParamsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationSsoFormParamsResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationSsoFormParamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationSsoFormParamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationSsoFormParamsResponse) GetBody() *UpdateApplicationSsoFormParamsResponseBody {
	return s.Body
}

func (s *UpdateApplicationSsoFormParamsResponse) SetHeaders(v map[string]*string) *UpdateApplicationSsoFormParamsResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationSsoFormParamsResponse) SetStatusCode(v int32) *UpdateApplicationSsoFormParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationSsoFormParamsResponse) SetBody(v *UpdateApplicationSsoFormParamsResponseBody) *UpdateApplicationSsoFormParamsResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationSsoFormParamsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
