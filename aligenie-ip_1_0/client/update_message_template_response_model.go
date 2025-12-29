// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMessageTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMessageTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMessageTemplateResponseBody) *UpdateMessageTemplateResponse
	GetBody() *UpdateMessageTemplateResponseBody
}

type UpdateMessageTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMessageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMessageTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateMessageTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMessageTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMessageTemplateResponse) GetBody() *UpdateMessageTemplateResponseBody {
	return s.Body
}

func (s *UpdateMessageTemplateResponse) SetHeaders(v map[string]*string) *UpdateMessageTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateMessageTemplateResponse) SetStatusCode(v int32) *UpdateMessageTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMessageTemplateResponse) SetBody(v *UpdateMessageTemplateResponseBody) *UpdateMessageTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateMessageTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
