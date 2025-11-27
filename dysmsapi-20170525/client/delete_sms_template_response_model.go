// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmsTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSmsTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSmsTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSmsTemplateResponseBody) *DeleteSmsTemplateResponse
	GetBody() *DeleteSmsTemplateResponseBody
}

type DeleteSmsTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSmsTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteSmsTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSmsTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSmsTemplateResponse) GetBody() *DeleteSmsTemplateResponseBody {
	return s.Body
}

func (s *DeleteSmsTemplateResponse) SetHeaders(v map[string]*string) *DeleteSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteSmsTemplateResponse) SetStatusCode(v int32) *DeleteSmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSmsTemplateResponse) SetBody(v *DeleteSmsTemplateResponseBody) *DeleteSmsTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteSmsTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
