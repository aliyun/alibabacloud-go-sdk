// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMiniAppBindingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMiniAppBindingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMiniAppBindingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMiniAppBindingResponseBody) *UpdateMiniAppBindingResponse
	GetBody() *UpdateMiniAppBindingResponseBody
}

type UpdateMiniAppBindingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMiniAppBindingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMiniAppBindingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMiniAppBindingResponse) GoString() string {
	return s.String()
}

func (s *UpdateMiniAppBindingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMiniAppBindingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMiniAppBindingResponse) GetBody() *UpdateMiniAppBindingResponseBody {
	return s.Body
}

func (s *UpdateMiniAppBindingResponse) SetHeaders(v map[string]*string) *UpdateMiniAppBindingResponse {
	s.Headers = v
	return s
}

func (s *UpdateMiniAppBindingResponse) SetStatusCode(v int32) *UpdateMiniAppBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMiniAppBindingResponse) SetBody(v *UpdateMiniAppBindingResponseBody) *UpdateMiniAppBindingResponse {
	s.Body = v
	return s
}

func (s *UpdateMiniAppBindingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
