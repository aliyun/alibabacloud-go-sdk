// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySandboxTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySandboxTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySandboxTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ModifySandboxTemplateResponseBody) *ModifySandboxTemplateResponse
	GetBody() *ModifySandboxTemplateResponseBody
}

type ModifySandboxTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySandboxTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySandboxTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySandboxTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifySandboxTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySandboxTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySandboxTemplateResponse) GetBody() *ModifySandboxTemplateResponseBody {
	return s.Body
}

func (s *ModifySandboxTemplateResponse) SetHeaders(v map[string]*string) *ModifySandboxTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifySandboxTemplateResponse) SetStatusCode(v int32) *ModifySandboxTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySandboxTemplateResponse) SetBody(v *ModifySandboxTemplateResponseBody) *ModifySandboxTemplateResponse {
	s.Body = v
	return s
}

func (s *ModifySandboxTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
