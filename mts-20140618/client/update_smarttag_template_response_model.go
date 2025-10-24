// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmarttagTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmarttagTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmarttagTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmarttagTemplateResponseBody) *UpdateSmarttagTemplateResponse
	GetBody() *UpdateSmarttagTemplateResponseBody
}

type UpdateSmarttagTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmarttagTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmarttagTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmarttagTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmarttagTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmarttagTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmarttagTemplateResponse) GetBody() *UpdateSmarttagTemplateResponseBody {
	return s.Body
}

func (s *UpdateSmarttagTemplateResponse) SetHeaders(v map[string]*string) *UpdateSmarttagTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmarttagTemplateResponse) SetStatusCode(v int32) *UpdateSmarttagTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmarttagTemplateResponse) SetBody(v *UpdateSmarttagTemplateResponseBody) *UpdateSmarttagTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateSmarttagTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
