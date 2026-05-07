// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetJobTemplateDefaultVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetJobTemplateDefaultVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetJobTemplateDefaultVersionResponse
	GetStatusCode() *int32
	SetBody(v *SetJobTemplateDefaultVersionResponseBody) *SetJobTemplateDefaultVersionResponse
	GetBody() *SetJobTemplateDefaultVersionResponseBody
}

type SetJobTemplateDefaultVersionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetJobTemplateDefaultVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetJobTemplateDefaultVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetJobTemplateDefaultVersionResponse) GoString() string {
	return s.String()
}

func (s *SetJobTemplateDefaultVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetJobTemplateDefaultVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetJobTemplateDefaultVersionResponse) GetBody() *SetJobTemplateDefaultVersionResponseBody {
	return s.Body
}

func (s *SetJobTemplateDefaultVersionResponse) SetHeaders(v map[string]*string) *SetJobTemplateDefaultVersionResponse {
	s.Headers = v
	return s
}

func (s *SetJobTemplateDefaultVersionResponse) SetStatusCode(v int32) *SetJobTemplateDefaultVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetJobTemplateDefaultVersionResponse) SetBody(v *SetJobTemplateDefaultVersionResponseBody) *SetJobTemplateDefaultVersionResponse {
	s.Body = v
	return s
}

func (s *SetJobTemplateDefaultVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
