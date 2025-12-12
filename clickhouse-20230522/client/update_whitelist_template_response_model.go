// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhitelistTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWhitelistTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWhitelistTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWhitelistTemplateResponseBody) *UpdateWhitelistTemplateResponse
	GetBody() *UpdateWhitelistTemplateResponseBody
}

type UpdateWhitelistTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWhitelistTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWhitelistTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhitelistTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateWhitelistTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWhitelistTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWhitelistTemplateResponse) GetBody() *UpdateWhitelistTemplateResponseBody {
	return s.Body
}

func (s *UpdateWhitelistTemplateResponse) SetHeaders(v map[string]*string) *UpdateWhitelistTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateWhitelistTemplateResponse) SetStatusCode(v int32) *UpdateWhitelistTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWhitelistTemplateResponse) SetBody(v *UpdateWhitelistTemplateResponseBody) *UpdateWhitelistTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateWhitelistTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
