// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultCustomTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDefaultCustomTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDefaultCustomTemplateResponse
	GetStatusCode() *int32
	SetBody(v *SetDefaultCustomTemplateResponseBody) *SetDefaultCustomTemplateResponse
	GetBody() *SetDefaultCustomTemplateResponseBody
}

type SetDefaultCustomTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultCustomTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultCustomTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultCustomTemplateResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultCustomTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDefaultCustomTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDefaultCustomTemplateResponse) GetBody() *SetDefaultCustomTemplateResponseBody {
	return s.Body
}

func (s *SetDefaultCustomTemplateResponse) SetHeaders(v map[string]*string) *SetDefaultCustomTemplateResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultCustomTemplateResponse) SetStatusCode(v int32) *SetDefaultCustomTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultCustomTemplateResponse) SetBody(v *SetDefaultCustomTemplateResponseBody) *SetDefaultCustomTemplateResponse {
	s.Body = v
	return s
}

func (s *SetDefaultCustomTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
