// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWhitelistTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWhitelistTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWhitelistTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateWhitelistTemplateResponseBody) *CreateWhitelistTemplateResponse
	GetBody() *CreateWhitelistTemplateResponseBody
}

type CreateWhitelistTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWhitelistTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWhitelistTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWhitelistTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateWhitelistTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWhitelistTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWhitelistTemplateResponse) GetBody() *CreateWhitelistTemplateResponseBody {
	return s.Body
}

func (s *CreateWhitelistTemplateResponse) SetHeaders(v map[string]*string) *CreateWhitelistTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateWhitelistTemplateResponse) SetStatusCode(v int32) *CreateWhitelistTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWhitelistTemplateResponse) SetBody(v *CreateWhitelistTemplateResponseBody) *CreateWhitelistTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateWhitelistTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
