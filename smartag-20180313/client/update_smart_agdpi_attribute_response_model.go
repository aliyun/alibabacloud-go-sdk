// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAGDpiAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmartAGDpiAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmartAGDpiAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmartAGDpiAttributeResponseBody) *UpdateSmartAGDpiAttributeResponse
	GetBody() *UpdateSmartAGDpiAttributeResponseBody
}

type UpdateSmartAGDpiAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmartAGDpiAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmartAGDpiAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAGDpiAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmartAGDpiAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmartAGDpiAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmartAGDpiAttributeResponse) GetBody() *UpdateSmartAGDpiAttributeResponseBody {
	return s.Body
}

func (s *UpdateSmartAGDpiAttributeResponse) SetHeaders(v map[string]*string) *UpdateSmartAGDpiAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmartAGDpiAttributeResponse) SetStatusCode(v int32) *UpdateSmartAGDpiAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmartAGDpiAttributeResponse) SetBody(v *UpdateSmartAGDpiAttributeResponseBody) *UpdateSmartAGDpiAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateSmartAGDpiAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
