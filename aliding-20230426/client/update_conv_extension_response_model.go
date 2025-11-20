// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConvExtensionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConvExtensionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConvExtensionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConvExtensionResponseBody) *UpdateConvExtensionResponse
	GetBody() *UpdateConvExtensionResponseBody
}

type UpdateConvExtensionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConvExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConvExtensionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConvExtensionResponse) GoString() string {
	return s.String()
}

func (s *UpdateConvExtensionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConvExtensionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConvExtensionResponse) GetBody() *UpdateConvExtensionResponseBody {
	return s.Body
}

func (s *UpdateConvExtensionResponse) SetHeaders(v map[string]*string) *UpdateConvExtensionResponse {
	s.Headers = v
	return s
}

func (s *UpdateConvExtensionResponse) SetStatusCode(v int32) *UpdateConvExtensionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConvExtensionResponse) SetBody(v *UpdateConvExtensionResponseBody) *UpdateConvExtensionResponse {
	s.Body = v
	return s
}

func (s *UpdateConvExtensionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
