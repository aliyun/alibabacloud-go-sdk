// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCardTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCardTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCardTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetCardTemplateResponseBody) *GetCardTemplateResponse
	GetBody() *GetCardTemplateResponseBody
}

type GetCardTemplateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCardTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCardTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCardTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetCardTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCardTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCardTemplateResponse) GetBody() *GetCardTemplateResponseBody {
	return s.Body
}

func (s *GetCardTemplateResponse) SetHeaders(v map[string]*string) *GetCardTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetCardTemplateResponse) SetStatusCode(v int32) *GetCardTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCardTemplateResponse) SetBody(v *GetCardTemplateResponseBody) *GetCardTemplateResponse {
	s.Body = v
	return s
}

func (s *GetCardTemplateResponse) Validate() error {
	return dara.Validate(s)
}
