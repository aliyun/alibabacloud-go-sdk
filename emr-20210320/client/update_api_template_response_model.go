// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApiTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApiTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApiTemplateResponseBody) *UpdateApiTemplateResponse
	GetBody() *UpdateApiTemplateResponseBody
}

type UpdateApiTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApiTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApiTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateApiTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApiTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApiTemplateResponse) GetBody() *UpdateApiTemplateResponseBody {
	return s.Body
}

func (s *UpdateApiTemplateResponse) SetHeaders(v map[string]*string) *UpdateApiTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateApiTemplateResponse) SetStatusCode(v int32) *UpdateApiTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApiTemplateResponse) SetBody(v *UpdateApiTemplateResponseBody) *UpdateApiTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateApiTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
