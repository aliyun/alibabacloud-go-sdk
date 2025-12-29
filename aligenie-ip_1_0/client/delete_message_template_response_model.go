// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMessageTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMessageTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMessageTemplateResponseBody) *DeleteMessageTemplateResponse
	GetBody() *DeleteMessageTemplateResponseBody
}

type DeleteMessageTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMessageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMessageTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteMessageTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMessageTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMessageTemplateResponse) GetBody() *DeleteMessageTemplateResponseBody {
	return s.Body
}

func (s *DeleteMessageTemplateResponse) SetHeaders(v map[string]*string) *DeleteMessageTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteMessageTemplateResponse) SetStatusCode(v int32) *DeleteMessageTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMessageTemplateResponse) SetBody(v *DeleteMessageTemplateResponseBody) *DeleteMessageTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteMessageTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
