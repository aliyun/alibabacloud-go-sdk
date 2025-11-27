// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmsTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySmsTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySmsTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ModifySmsTemplateResponseBody) *ModifySmsTemplateResponse
	GetBody() *ModifySmsTemplateResponseBody
}

type ModifySmsTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySmsTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifySmsTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySmsTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySmsTemplateResponse) GetBody() *ModifySmsTemplateResponseBody {
	return s.Body
}

func (s *ModifySmsTemplateResponse) SetHeaders(v map[string]*string) *ModifySmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifySmsTemplateResponse) SetStatusCode(v int32) *ModifySmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySmsTemplateResponse) SetBody(v *ModifySmsTemplateResponseBody) *ModifySmsTemplateResponse {
	s.Body = v
	return s
}

func (s *ModifySmsTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
