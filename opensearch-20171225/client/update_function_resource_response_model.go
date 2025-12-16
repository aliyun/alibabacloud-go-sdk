// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunctionResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFunctionResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFunctionResourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFunctionResourceResponseBody) *UpdateFunctionResourceResponse
	GetBody() *UpdateFunctionResourceResponseBody
}

type UpdateFunctionResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFunctionResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFunctionResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFunctionResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFunctionResourceResponse) GetBody() *UpdateFunctionResourceResponseBody {
	return s.Body
}

func (s *UpdateFunctionResourceResponse) SetHeaders(v map[string]*string) *UpdateFunctionResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionResourceResponse) SetStatusCode(v int32) *UpdateFunctionResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFunctionResourceResponse) SetBody(v *UpdateFunctionResourceResponseBody) *UpdateFunctionResourceResponse {
	s.Body = v
	return s
}

func (s *UpdateFunctionResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
