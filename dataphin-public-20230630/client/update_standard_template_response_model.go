// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStandardTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStandardTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStandardTemplateResponseBody) *UpdateStandardTemplateResponse
	GetBody() *UpdateStandardTemplateResponseBody
}

type UpdateStandardTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStandardTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStandardTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStandardTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStandardTemplateResponse) GetBody() *UpdateStandardTemplateResponseBody {
	return s.Body
}

func (s *UpdateStandardTemplateResponse) SetHeaders(v map[string]*string) *UpdateStandardTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateStandardTemplateResponse) SetStatusCode(v int32) *UpdateStandardTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStandardTemplateResponse) SetBody(v *UpdateStandardTemplateResponseBody) *UpdateStandardTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateStandardTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
