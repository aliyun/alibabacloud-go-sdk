// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppFromTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppFromTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppFromTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppFromTemplateResponseBody) *CreateAppFromTemplateResponse
	GetBody() *CreateAppFromTemplateResponseBody
}

type CreateAppFromTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppFromTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppFromTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppFromTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateAppFromTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppFromTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppFromTemplateResponse) GetBody() *CreateAppFromTemplateResponseBody {
	return s.Body
}

func (s *CreateAppFromTemplateResponse) SetHeaders(v map[string]*string) *CreateAppFromTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateAppFromTemplateResponse) SetStatusCode(v int32) *CreateAppFromTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppFromTemplateResponse) SetBody(v *CreateAppFromTemplateResponseBody) *CreateAppFromTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateAppFromTemplateResponse) Validate() error {
	return dara.Validate(s)
}
