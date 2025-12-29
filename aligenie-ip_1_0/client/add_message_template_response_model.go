// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMessageTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMessageTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMessageTemplateResponse
	GetStatusCode() *int32
	SetBody(v *AddMessageTemplateResponseBody) *AddMessageTemplateResponse
	GetBody() *AddMessageTemplateResponseBody
}

type AddMessageTemplateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMessageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMessageTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMessageTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddMessageTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMessageTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMessageTemplateResponse) GetBody() *AddMessageTemplateResponseBody {
	return s.Body
}

func (s *AddMessageTemplateResponse) SetHeaders(v map[string]*string) *AddMessageTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddMessageTemplateResponse) SetStatusCode(v int32) *AddMessageTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMessageTemplateResponse) SetBody(v *AddMessageTemplateResponseBody) *AddMessageTemplateResponse {
	s.Body = v
	return s
}

func (s *AddMessageTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
