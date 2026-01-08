// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatappTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteChatappTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteChatappTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteChatappTemplateResponseBody) *DeleteChatappTemplateResponse
	GetBody() *DeleteChatappTemplateResponseBody
}

type DeleteChatappTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChatappTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatappTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteChatappTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteChatappTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteChatappTemplateResponse) GetBody() *DeleteChatappTemplateResponseBody {
	return s.Body
}

func (s *DeleteChatappTemplateResponse) SetHeaders(v map[string]*string) *DeleteChatappTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteChatappTemplateResponse) SetStatusCode(v int32) *DeleteChatappTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChatappTemplateResponse) SetBody(v *DeleteChatappTemplateResponseBody) *DeleteChatappTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteChatappTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
