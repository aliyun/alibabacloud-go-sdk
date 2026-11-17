// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebFilterTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWebFilterTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWebFilterTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWebFilterTemplateResponseBody) *DeleteWebFilterTemplateResponse
	GetBody() *DeleteWebFilterTemplateResponseBody
}

type DeleteWebFilterTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWebFilterTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWebFilterTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebFilterTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebFilterTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWebFilterTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWebFilterTemplateResponse) GetBody() *DeleteWebFilterTemplateResponseBody {
	return s.Body
}

func (s *DeleteWebFilterTemplateResponse) SetHeaders(v map[string]*string) *DeleteWebFilterTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebFilterTemplateResponse) SetStatusCode(v int32) *DeleteWebFilterTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWebFilterTemplateResponse) SetBody(v *DeleteWebFilterTemplateResponseBody) *DeleteWebFilterTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteWebFilterTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
